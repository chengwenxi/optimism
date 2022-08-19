package derivation

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-node/l2"
	"github.com/ethereum-optimism/optimism/op-node/testutils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/trie"
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
	actL2PipelineStep   = iota // run L2 derivation pipeline
	actL2BatchBuffer           // add next L2 block to batch buffer
	actL2BatchSubmit           // construct batch tx from L2 buffer content, submit to L1
	actL1Deposit               // add rollup deposit to L1 tx queue
	actL1AddTx                 // add regular tx to L1 tx queue
	actL2AddTx                 // add regular tx to L2 tx queue
	actL2MoveL1Origin          // move to next L1 origin
	actL1RewindToParent        // rewind L1 chain to parent block of head
	actL1StartBlock            // start new L1 block on top of head
	actL1IncludeTx             // include next tx from L1 tx queue
	actL1EndBlock              // finish new L1 block, apply to chain as unsafe block
	actL1FinalizeNext
	actL1SafeNext
	actL1Sync                // process next canonical L1 block (may reorg)
	actL2StartBlock          // start new L2 block on top of head
	actL2IncludeTx           // add next tx from L2 tx queue
	actL2EndBlock            // finish new L2 block, apply to chain as unsafe block
	actL1RPCFail             // make next L1 request fail
	actL2RPCFail             // make next L2 request fail
	actL2UnsafeGossipFail    // make next gossip receive fail
	actL2UnsafeGossipReceive // process payload from gossip
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

	// L1 block building data
	l1VMconf         vm.Config            // vm config used during ongoing building (includes trace options)
	l1BuildingHeader *types.Header        // block header that we add txs to for block building
	l1BuildingState  *state.StateDB       // state used for block building
	l1GasPool        *core.GasPool        // track gas used of ongoing building
	l1Transactions   []*types.Transaction // collects txs that were successfully included into current block build
	l1Receipts       []*types.Receipt     // collect receipts of ongoing building
	l1TimeDelta      uint64               // how time to add to next block timestamp. Minimum of 1.
	l1Building       bool
	l1TxFailed       []*types.Transaction // log of failed transactions which could not be included

	// L2 evm / chain
	l2Chain     *core.BlockChain
	l2Database  ethdb.Database
	l2Consensus consensus.Engine
	l2Cfg       *core.Genesis
	l2Signer    types.Signer
	l2Engine    *Engine

	// L2 block building data
	l2VMconf         vm.Config            // vm config used during ongoing building (includes trace options)
	l2BuildingHeader *types.Header        // block header that we add txs to for block building
	l2BuildingState  *state.StateDB       // state used for block building
	l2GasPool        *core.GasPool        // track gas used of ongoing building
	l2Transactions   []*types.Transaction // collects txs that were successfully included into current block build
	l2Receipts       []*types.Receipt     // collect receipts of ongoing building
	l2Building       bool
	l2TxFailed       []*types.Transaction // log of failed transactions which could not be included

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
	case actL1FinalizeNext:
		return s.actL1FinalizeNext()
	case actL1SafeNext:
		return s.actL1SafeNext()
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
	if s.l2Building {
		return InvalidActionErr
	}

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
	if s.l1Building {
		// not valid if we already started building a block
		return InvalidActionErr
	}
	if s.l1TimeDelta == 0 || s.l1TimeDelta > 60*60*24 {
		return fmt.Errorf("invalid time delta: %d", s.l1TimeDelta)
	}

	parent := s.l1Chain.CurrentHeader()
	parentHash := parent.Hash()
	statedb, err := state.New(parent.Root, state.NewDatabase(s.l1Database), nil)
	if err != nil {
		return fmt.Errorf("failed to init state db around block %s (state %s): %w", parentHash, parent.Root, err)
	}
	header := &types.Header{
		ParentHash: parentHash,
		Coinbase:   parent.Coinbase,
		Difficulty: common.Big0,
		Number:     new(big.Int).Add(parent.Number, common.Big1),
		GasLimit:   parent.GasLimit,
		Time:       parent.Time + s.l1TimeDelta,
		Extra:      []byte("L1 was here"),
		MixDigest:  common.Hash{}, // TODO: maybe randomize this (prev-randao value)
	}
	if s.l1Cfg.Config.IsLondon(header.Number) {
		header.BaseFee = misc.CalcBaseFee(s.l1Cfg.Config, parent)
		// At the transition, double the gas limit so the gas target is equal to the old gas limit.
		if !s.l1Cfg.Config.IsLondon(parent.Number) {
			header.GasLimit = parent.GasLimit * params.ElasticityMultiplier
		}
	}

	s.l1Building = true
	s.l1BuildingHeader = header
	s.l1BuildingState = statedb
	s.l1Receipts = make([]*types.Receipt, 0)
	s.l1Transactions = make([]*types.Transaction, 0)
	s.l1VMconf = vm.Config{}
	// TODO: maybe add tracer to vm config here

	s.l1GasPool = new(core.GasPool).AddGas(header.GasLimit)
	return nil
}

func (s *State) actL1IncludeTx() error {
	if !s.l1Building {
		return InvalidActionErr
	}
	if len(s.l1TxQueue) == 0 {
		return InvalidActionErr
	}
	tx := s.l1TxQueue[0]
	if tx.Gas() > s.l1BuildingHeader.GasLimit {
		return fmt.Errorf("tx consumes %d gas, more than available in L1 block %d", tx.Gas(), s.l1BuildingHeader.GasLimit)
	}
	if uint64(*s.l1GasPool) < tx.Gas() {
		return InvalidActionErr // insufficient gas to include the tx
	}
	s.l1TxQueue = s.l1TxQueue[1:] // won't retry the tx
	receipt, err := core.ApplyTransaction(s.l1Cfg.Config, s.l1Chain, &s.l1BuildingHeader.Coinbase,
		s.l1GasPool, s.l1BuildingState, s.l1BuildingHeader, tx, &s.l1BuildingHeader.GasUsed, s.l1VMconf)
	if err != nil {
		s.l1TxFailed = append(s.l1TxFailed, tx)
		return fmt.Errorf("failed to apply transaction to L1 block (tx %d): %w", len(s.l1Transactions), err)
	}
	s.l1Receipts = append(s.l1Receipts, receipt)
	s.l1Transactions = append(s.l1Transactions, tx)
	return nil
}

func (s *State) actL1EndBlock() error {
	if !s.l1Building {
		// not valid if we are not building a block currently
		return InvalidActionErr
	}

	s.l1Building = false
	s.l1BuildingHeader.GasUsed = s.l1BuildingHeader.GasLimit - uint64(*s.l1GasPool)
	s.l1BuildingHeader.Root = s.l1BuildingState.IntermediateRoot(s.l1Cfg.Config.IsEIP158(s.l1BuildingHeader.Number))
	block := types.NewBlock(s.l1BuildingHeader, s.l1Transactions, nil, s.l1Receipts, trie.NewStackTrie(nil))

	// Write state changes to db
	root, err := s.l1BuildingState.Commit(s.l1Cfg.Config.IsEIP158(s.l1BuildingHeader.Number))
	if err != nil {
		return fmt.Errorf("l1 state write error: %v", err)
	}
	if err := s.l1BuildingState.Database().TrieDB().Commit(root, false, nil); err != nil {
		return fmt.Errorf("l1 trie write error: %v", err)
	}

	_, err = s.l1Chain.InsertChain(types.Blocks{block})
	if err != nil {
		return fmt.Errorf("failed to insert block into l1 chain")
	}
	return nil
}

func (s *State) actL1FinalizeNext() error {
	safe := s.l1Chain.CurrentSafeBlock()
	finalizedNum := s.l1Chain.CurrentFinalizedBlock().NumberU64()
	if safe.NumberU64() <= finalizedNum {
		return InvalidActionErr // need to move forward safe block before moving finalized block
	}
	next := s.l1Chain.GetBlockByNumber(finalizedNum + 1)
	if next == nil {
		return fmt.Errorf("expected next block after finalized L1 block %d, safe head is ahead", finalizedNum)
	}
	s.l1Chain.SetFinalized(next)
	return nil
}
func (s *State) actL1SafeNext() error {
	safe := s.l1Chain.CurrentSafeBlock()
	next := s.l1Chain.GetBlockByNumber(safe.NumberU64() + 1)
	if next == nil {
		return InvalidActionErr // if head of chain is marked as safe then there's no next block
	}
	s.l1Chain.SetSafe(next)
	return nil
}

func (s *State) actL1Sync() error {
	if s.canonL1 != nil {
		// TODO: implement basic sync
		return InvalidActionErr
	}
	return InvalidActionErr
}

func (s *State) actL2StartBlock() error {
	if !s.l2PipelineIdle {
		return InvalidActionErr
	}

	s.seqNextOrigin = false

	parent := s.l2Chain.CurrentBlock()
	parentHeader := parent.Header()
	statedb, err := state.New(parent.Root(), state.NewDatabase(s.l2Database), nil)
	if err != nil {
		return fmt.Errorf("failed to init state db around block %s (state %s): %w", parent.Hash(), parent.Root(), err)
	}
	parentRef, err := l2.BlockToBlockRef(parent, &s.rollupCfg.Genesis)
	if err != nil {
		return fmt.Errorf("failed to get L2 block ref: %w", err)
	}
	currentOriginBlock := s.l1Chain.GetBlockByHash(parentRef.L1Origin.Hash)
	if currentOriginBlock == nil {
		return fmt.Errorf("origin block of L2 %s does not exist: %s", parentRef, parentRef.L1Origin)
	}
	l2Timestamp := parentHeader.Time + s.rollupCfg.BlockTime

	// findL1Origin test equivalent
	nextOriginBlock := s.l1Chain.GetBlockByNumber(currentOriginBlock.NumberU64() + 1)
	originBlock := currentOriginBlock
	if nextOriginBlock != nil && l2Timestamp >= nextOriginBlock.Time() {
		originBlock = nextOriginBlock
	}

	// PreparePayloadAttributes test equivalent
	var depositTxs []hexutil.Bytes
	var seqNumber uint64
	if parentRef.L1Origin.Number != originBlock.NumberU64() {
		if parentRef.L1Origin.Hash != originBlock.ParentHash() {
			return fmt.Errorf("cannot create new block with L1 origin %s (parent %s) on top of L1 origin %s",
				originBlock.Hash(), originBlock.ParentHash(), parentRef.L1Origin)
		}
		receipts := s.l1Chain.GetReceiptsByHash(originBlock.Hash())
		deposits, err := derive.DeriveDeposits(receipts, s.rollupCfg.DepositContractAddress)
		if err != nil {
			return fmt.Errorf("failed to derive some deposits: %w", err)
		}
		depositTxs = deposits
		seqNumber = 0
	} else {
		if parentRef.L1Origin.Hash != originBlock.Hash() {
			return fmt.Errorf("cannot create new block with L1 origin %s in conflict with L1 origin %s", originBlock.Hash(), parentRef.L1Origin)
		}
		depositTxs = nil
		seqNumber = parentRef.SequenceNumber + 1
	}

	header := &types.Header{
		ParentHash: parent.Hash(),
		Coinbase:   s.rollupCfg.FeeRecipientAddress,
		Difficulty: common.Big0,
		Number:     new(big.Int).Add(parentHeader.Number, common.Big1),
		GasLimit:   parentHeader.GasLimit,
		Time:       l2Timestamp,
		Extra:      nil,
		MixDigest:  originBlock.MixDigest(),
	}
	header.BaseFee = misc.CalcBaseFee(s.l2Cfg.Config, parentHeader)

	s.l2Building = true
	s.l2BuildingHeader = header
	s.l2BuildingState = statedb
	s.l2Receipts = make([]*types.Receipt, 0)
	s.l2Transactions = make([]*types.Transaction, 0)
	s.l2VMconf = vm.Config{}
	// TODO: maybe add tracer to vm config here

	s.l2GasPool = new(core.GasPool).AddGas(header.GasLimit)

	l1Info := &testutils.MockL1Info{
		InfoHash:           originBlock.Hash(),
		InfoParentHash:     originBlock.ParentHash(),
		InfoRoot:           originBlock.Root(),
		InfoNum:            originBlock.NumberU64(),
		InfoTime:           originBlock.Time(),
		InfoMixDigest:      originBlock.MixDigest(),
		InfoBaseFee:        originBlock.BaseFee(),
		InfoReceiptRoot:    originBlock.ReceiptHash(),
		InfoSequenceNumber: seqNumber,
	}
	l1InfoTx, err := derive.L1InfoDepositBytes(seqNumber, l1Info)
	if err != nil {
		return fmt.Errorf("failed to create l1InfoTx: %w", err)
	}

	txs := make([]hexutil.Bytes, 0, 1+len(depositTxs))
	txs = append(txs, l1InfoTx)
	txs = append(txs, depositTxs...)

	// pre-process the deposits

	// if already started
	return InvalidActionErr
}

func (s *State) actL2IncludeTx() error {
	if !s.l2PipelineIdle {
		return InvalidActionErr
	}

	// if insufficient gas, or forced to produce empty block due to sequencer drift

	return InvalidActionErr
}

func (s *State) actL2EndBlock() error {
	if !s.l2PipelineIdle {
		return InvalidActionErr
	}

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
