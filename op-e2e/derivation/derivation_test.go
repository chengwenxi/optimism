package derivation

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var InvalidActionErr = errors.New("invalid action")

type Action uint8

const (
	actL2PipelineStep        = iota // run L2 derivation pipeline
	actL2BatchBuffer                // add next L2 block to batch buffer
	actL2BatchSubmit                // construct batch tx from L2 buffer content, submit to L1
	actL1Deposit                    // add rollup deposit to L1 tx queue
	actL1AddTx                      // add regular tx to L1 tx queue
	actL2AddTx                      // add regular tx to L2 tx queue
	actL2MoveL1Origin               // move to next L1 origin
	actL1RewindToParent             // rewind L1 chain to parent block of head
	actL1StartBlock                 // start new L1 block on top of head
	actL1IncludeTx                  // include next tx from L1 tx queue
	actL1EndBlock                   // finish new L1 block, apply to chain as unsafe block
	actL1Sync                       // process next canonical L1 block (may reorg)
	actL2StartBlock                 // start new L2 block on top of head
	actL2IncludeTx                  // add next tx from L2 tx queue
	actL2EndBlock                   // finish new L2 block, apply to chain as unsafe block
	actL1RPCFail                    // make next L1 request fail
	actL2RPCFail                    // make next L2 request fail
	actL2UnsafeGossipFail           // make next gossip receive fail
	actL2UnsafeGossipReceive        // process payload from gossip
)

type BatcherCfg struct {
	// Limit the size of txs
	MinL1TxSize uint64
	MaxL1TxSize uint64

	// Where to send the batch txs to.
	BatchInboxAddress common.Address

	BatcherKey *ecdsa.PrivateKey
}

type State struct {
	log            log.Logger
	l1TxQueue      []*types.Transaction
	l2TxQueue      []*types.Transaction
	l2Pipeline     *derive.DerivationPipeline
	l2PipelineIdle bool

	seqNextOrigin bool // move to next L1 origin when sequencing a block

	rollupCfg *rollup.Config

	l2ChannelOut     *derive.ChannelOut
	l2Submitting     bool // when the channel out is being submitted, and not safe to write to without resetting
	l2BufferedBlock  eth.BlockID
	l2SubmittedBlock eth.BlockID
	l2BatcherCfg     BatcherCfg

	// L1 evm / chain
	l1Chain     *core.BlockChain
	l1Database  ethdb.Database
	l1Consensus consensus.Engine
	l1Cfg       *core.Genesis
	l1Signer    types.Signer

	// L2 evm / chain
	l2Chain     *core.BlockChain
	l2Database  ethdb.Database
	l2Consensus consensus.Engine
	l2Cfg       *core.Genesis
	l2Signer    types.Signer
	l2Engine    *Engine

	// Used to sync L1 from another node.
	// The other node always has the canonical chain.
	// May be nil if there is nothing to sync from
	canonL1 func(num uint64) *types.Block

	// Mock errs, to be returned on next usage of the call/function
	failL1RPC               error
	failL2RPC               error
	failL2GossipUnsafeBlock error
}

func NewState(canonL1 func(num uint64) *types.Block) {
	// todo types.LatestSignerForChainID(1234)
}

func (s *State) RunAction(action Action) error {
	switch action {
	case actL2PipelineStep:
		return s.actL2PipelineStep()
	case actL2BatchBuffer:
		return s.actL2BatchBuffer()
	case actL2BatchSubmit:
		return s.actL2BatchSubmit()
	case actL1Deposit:
		return s.actL1Deposit()
	case actL1AddTx:
		return s.actL1AddTx()
	case actL2AddTx:
		return s.actL2AddTx()
	case actL2MoveL1Origin:
		return s.actL2MoveL1Origin()
	case actL1RewindToParent:
		return s.actL1RewindToParent()
	case actL1StartBlock:
		return s.actL1StartBlock()
	case actL1IncludeTx:
		return s.actL1IncludeTx()
	case actL1EndBlock:
		return s.actL1EndBlock()
	case actL1Sync:
		return s.actL1Sync()
	case actL2StartBlock:
		return s.actL2StartBlock()
	case actL2IncludeTx:
		return s.actL2IncludeTx()
	case actL2EndBlock:
		return s.actL2EndBlock()
	case actL1RPCFail:
		return s.actL1RPCFail()
	case actL2RPCFail:
		return s.actL2RPCFail()
	case actL2UnsafeGossipFail:
		return s.actL2UnsafeGossipFail()
	case actL2UnsafeGossipReceive:
		return s.actL2UnsafeGossipReceive()
	default:
		return InvalidActionErr
	}
}

func (s *State) actL2PipelineStep() error {
	s.l2PipelineIdle = false
	err := s.l2Pipeline.Step(context.Background())
	if err == io.EOF {
		s.l2PipelineIdle = true
		return nil
	} else if err != nil && errors.Is(err, derive.ErrReset) {
		s.log.Warn("Derivation pipeline is reset", "err", err)
		s.l2Pipeline.Reset()
		return nil
	} else if err != nil && errors.Is(err, derive.ErrTemporary) {
		s.log.Warn("Derivation process temporary error", "err", err)
		return nil
	} else if err != nil && errors.Is(err, derive.ErrCritical) {
		return fmt.Errorf("derivation failed critically: %w", err)
	} else {
		return nil
	}
}

func (s *State) actL2BatchBuffer() error {
	if s.l2Submitting { // break ongoing submitting work if necessary
		s.l2ChannelOut = nil
		s.l2Submitting = false
	}
	safeHead := s.l2Pipeline.SafeL2Head()
	// If we just started, start at safe-head
	if s.l2SubmittedBlock == (eth.BlockID{}) {
		s.log.Info("Starting batch-submitter work at safe-head", "safe", safeHead)
		s.l2SubmittedBlock = safeHead.ID()
		s.l2BufferedBlock = safeHead.ID()
		s.l2ChannelOut = nil
	}
	// If it's lagging behind, catch it up.
	if s.l2SubmittedBlock.Number < safeHead.Number {
		s.log.Warn("last submitted block lagged behind L2 safe head: batch submission will continue from the safe head now", "last", s.l2SubmittedBlock, "safe", safeHead)
		s.l2SubmittedBlock = safeHead.ID()
		s.l2BufferedBlock = safeHead.ID()
		s.l2ChannelOut = nil
	}
	// Create channel if we don't have one yet
	if s.l2ChannelOut == nil {
		ch, err := derive.NewChannelOut(s.l1Chain.CurrentHeader().Time)
		if err != nil { // should always succeed
			return fmt.Errorf("failed to create channel: %w", err)
		}
		s.l2ChannelOut = ch
	}
	// Add the next unsafe block to the channel
	if s.l2BufferedBlock.Number >= s.l2Pipeline.UnsafeL2Head().Number {
		return nil
	}
	block := s.l2Chain.GetBlockByNumber(s.l2BufferedBlock.Number + 1)
	if block.ParentHash() != s.l2BufferedBlock.Hash {
		s.log.Error("detected a reorg in L2 chain vs previous submitted information, resetting to safe head now", "safe_head", safeHead)
		s.l2SubmittedBlock = safeHead.ID()
		s.l2BufferedBlock = safeHead.ID()
		s.l2ChannelOut = nil
	}
	if err := s.l2ChannelOut.AddBlock(block); err != nil { // should always succeed
		return fmt.Errorf("failed to add block to channel: %w", err)
	}
	// if too much data
	return InvalidActionErr
}

// Returns the current nonce for the given address, including in-flight txs.
// Add +1 to get the nonce to use for the next tx.
func (s *State) l1NonceFor(addr common.Address) (uint64, error) {
	// iterate in reverse order to find any in-flight txs
	for i := len(s.l1TxQueue) - 1; i >= 0; i-- {
		tx := s.l1TxQueue[i]
		sender, err := s.l1Signer.Sender(tx) // cached, doesn't hurt
		if err != nil {
			return 0, fmt.Errorf("failed to get tx sender: %w", err)
		}
		if sender == addr {
			return tx.Nonce(), nil
		}
	}
	statedb, err := s.l1Chain.State()
	if err != nil {
		return 0, fmt.Errorf("failed to get state db: %v", err)
	}
	return statedb.GetNonce(addr), nil
}

func (s *State) actL2BatchSubmit() error {
	// Don't run this action if there's no data to submit
	if s.l2ChannelOut == nil || s.l2ChannelOut.ReadyBytes() == 0 {
		return InvalidActionErr
	}
	// Collect the output frame
	data := new(bytes.Buffer)
	data.WriteByte(derive.DerivationVersion0)
	// subtract one, to account for the version byte
	if err := s.l2ChannelOut.OutputFrame(data, s.l2BatcherCfg.MaxL1TxSize-1); err == io.EOF {
		s.l2Submitting = false
		// there may still be some data to submit
	} else if err != nil {
		s.l2Submitting = false
		return fmt.Errorf("failed to output channel data to frame: %w", err)
	}

	nonce, err := s.l1NonceFor(s.rollupCfg.BatchSenderAddress)
	if err != nil {
		return fmt.Errorf("failed to get batcher nonce: %v", err)
	}

	gasTipCap := big.NewInt(2 * params.GWei)
	head := s.l1Chain.CurrentHeader()
	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(head.BaseFee, big.NewInt(2)))

	rawTx := &types.DynamicFeeTx{
		ChainID:   s.l1Cfg.Config.ChainID,
		Nonce:     nonce + 1,
		To:        &s.rollupCfg.BatchInboxAddress,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Data:      data.Bytes(),
	}
	gas, err := core.IntrinsicGas(rawTx.Data, nil, false, true, true)
	if err != nil {
		return fmt.Errorf("failed to compute intrinsic gas: %w", err)
	}
	rawTx.Gas = gas

	tx, err := types.SignNewTx(s.l2BatcherCfg.BatcherKey, s.l1Signer, rawTx)
	if err != nil {
		return fmt.Errorf("failed to sign tx: %w", err)
	}

	s.l1TxQueue = append(s.l1TxQueue, tx)
	return nil
}

func (s *State) actL1Deposit() error {
	// TODO create a deposit tx on L1 to L2, append to L1 tx queue
	return InvalidActionErr
}

func (s *State) actL1AddTx() error {
	// TODO create a regular random tx on L1, append to L1 tx queue
	return InvalidActionErr
}

func (s *State) actL2AddTx() error {
	// TODO create a regular random tx on L2, append to L2 tx queue
	return InvalidActionErr
}

func (s *State) actL2MoveL1Origin() error {
	if s.seqNextOrigin { // don't do this twice
		return InvalidActionErr
	}
	s.seqNextOrigin = true
	return nil
}

func (s *State) actL1RewindToParent() error {
	head := s.l1Chain.CurrentHeader().Number.Uint64()
	if head == 0 {
		return InvalidActionErr
	}
	if err := s.l1Chain.SetHead(head - 1); err != nil {
		return fmt.Errorf("failed to rewind L1 chain to nr %d: %v", head-1, err)
	}
	return nil
}

func (s *State) actL1StartBlock() error {
	// if already started
	return InvalidActionErr
}

func (s *State) actL1IncludeTx() error {
	// if insufficient gas
	return InvalidActionErr
}

func (s *State) actL1EndBlock() error {
	// if not started
	return InvalidActionErr
}

func (s *State) actL1Sync() error {
	if s.canonL1 != nil {
		// TODO: implement basic sync
		return InvalidActionErr
	}
	return InvalidActionErr
}

func (s *State) actL2StartBlock() error {
	s.seqNextOrigin = false
	// if already started
	return InvalidActionErr
}

func (s *State) actL2IncludeTx() error {
	// if insufficient gas, or forced to produce empty block due to sequencer drift
	return InvalidActionErr
}

func (s *State) actL2EndBlock() error {
	// if not started
	return InvalidActionErr
}

func (s *State) actL1RPCFail() error {
	if s.failL1RPC != nil { // already set to fail?
		return InvalidActionErr
	}
	s.failL1RPC = errors.New("mock L1 RPC error")
	return nil
}

func (s *State) actL2RPCFail() error {
	if s.failL2RPC != nil { // already set to fail?
		return InvalidActionErr
	}
	s.failL2RPC = errors.New("mock L2 RPC error")
	return nil
}

func (s *State) actL2UnsafeGossipFail() error {
	return InvalidActionErr
}

func (s *State) actL2UnsafeGossipReceive() error {
	return InvalidActionErr
}

func TestDerivation(t *testing.T) {

}
