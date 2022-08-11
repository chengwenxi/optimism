package derivation

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/l2"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/beacon"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

const maxTrackedPayloads = 10

// payloadQueueItem represents an id->payload tuple to store until it's retrieved
// or evicted.
type payloadQueueItem struct {
	id   beacon.PayloadID
	data *eth.ExecutionPayload
}

// payloadQueue tracks the latest handful of constructed payloads to be retrieved
// by the beacon chain if block production is requested.
type payloadQueue struct {
	payloads []*payloadQueueItem
	lock     sync.RWMutex
}

// newPayloadQueue creates a pre-initialized queue with a fixed number of slots
// all containing empty items.
func newPayloadQueue() *payloadQueue {
	return &payloadQueue{
		payloads: make([]*payloadQueueItem, maxTrackedPayloads),
	}
}

// put inserts a new payload into the queue at the given id.
func (q *payloadQueue) put(id beacon.PayloadID, data *eth.ExecutionPayload) {
	q.lock.Lock()
	defer q.lock.Unlock()

	copy(q.payloads[1:], q.payloads)
	q.payloads[0] = &payloadQueueItem{
		id:   id,
		data: data,
	}
}

// get retrieves a previously stored payload item or nil if it does not exist.
func (q *payloadQueue) get(id beacon.PayloadID) *eth.ExecutionPayload {
	q.lock.RLock()
	defer q.lock.RUnlock()

	for _, item := range q.payloads {
		if item == nil {
			return nil // no more items
		}
		if item.id == id {
			return item.data
		}
	}
	return nil
}

// Engine is an in-memory implementation of the Engine API,
// without support for snap-sync, and no concurrency or background processes.
type Engine struct {
	rollupCfg *rollup.Config
	log       log.Logger
	chain     *core.BlockChain
	db        ethdb.Database
	payloads  *payloadQueue
}

// TODO init engine

var _ derive.Engine = (*Engine)(nil)

func (e *Engine) GetPayload(ctx context.Context, payloadId eth.PayloadID) (*eth.ExecutionPayload, error) {
	e.log.Trace("Engine API request received", "method", "GetPayload", "id", payloadId)
	data := e.payloads.get(payloadId)
	if data == nil {
		return nil, beacon.UnknownPayload
	}
	return data, nil
}

var (
	STATUS_INVALID         = &eth.ForkchoiceUpdatedResult{PayloadStatus: eth.PayloadStatusV1{Status: eth.ExecutionInvalid}, PayloadID: nil}
	STATUS_SYNCING         = &eth.ForkchoiceUpdatedResult{PayloadStatus: eth.PayloadStatusV1{Status: eth.ExecutionSyncing}, PayloadID: nil}
	INVALID_TERMINAL_BLOCK = eth.PayloadStatusV1{Status: eth.ExecutionInvalid, LatestValidHash: &common.Hash{}}
)

// computePayloadId computes a pseudo-random payloadid, based on the parameters.
func computePayloadId(headBlockHash common.Hash, params *eth.PayloadAttributes) beacon.PayloadID {
	// Hash
	hasher := sha256.New()
	hasher.Write(headBlockHash[:])
	_ = binary.Write(hasher, binary.BigEndian, params.Timestamp)
	hasher.Write(params.PrevRandao[:])
	hasher.Write(params.SuggestedFeeRecipient[:])
	var out beacon.PayloadID
	copy(out[:], hasher.Sum(nil)[:8])
	return out
}

func (e *Engine) getSealingBlock(parent common.Hash, timestamp uint64, coinbase common.Address, random common.Hash, noTxs bool, forceTxs types.Transactions) (*types.Block, error) {
	// TODO
	return nil, nil
}

func (e *Engine) ForkchoiceUpdate(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error) {
	e.log.Trace("Engine API request received", "method", "ForkchoiceUpdated", "head", state.HeadBlockHash, "finalized", state.FinalizedBlockHash, "safe", state.SafeBlockHash)
	if state.HeadBlockHash == (common.Hash{}) {
		e.log.Warn("Forkchoice requested update to zero hash")
		return STATUS_INVALID, nil
	}
	// Check whether we have the block yet in our database or not. If not, we'll
	// need to either trigger a sync, or to reject this forkchoice update for a
	// reason.
	block := e.chain.GetBlockByHash(state.HeadBlockHash)
	if block == nil {
		// TODO: syncing not supported yet
		return STATUS_SYNCING, nil
	}
	// Block is known locally, just sanity check that the beacon client does not
	// attempt to push us back to before the merge.
	if block.Difficulty().BitLen() > 0 || block.NumberU64() == 0 {
		var (
			td  = e.chain.GetTd(state.HeadBlockHash, block.NumberU64())
			ptd = e.chain.GetTd(block.ParentHash(), block.NumberU64()-1)
			ttd = e.chain.Config().TerminalTotalDifficulty
		)
		if td == nil || (block.NumberU64() > 0 && ptd == nil) {
			e.log.Error("TDs unavailable for TTD check", "number", block.NumberU64(), "hash", state.HeadBlockHash, "td", td, "parent", block.ParentHash(), "ptd", ptd)
			return STATUS_INVALID, errors.New("TDs unavailable for TDD check")
		}
		if td.Cmp(ttd) < 0 {
			e.log.Error("Refusing beacon update to pre-merge", "number", block.NumberU64(), "hash", state.HeadBlockHash, "diff", block.Difficulty(), "age", common.PrettyAge(time.Unix(int64(block.Time()), 0)))
			return &eth.ForkchoiceUpdatedResult{PayloadStatus: INVALID_TERMINAL_BLOCK, PayloadID: nil}, nil
		}
		if block.NumberU64() > 0 && ptd.Cmp(ttd) >= 0 {
			e.log.Error("Parent block is already post-ttd", "number", block.NumberU64(), "hash", state.HeadBlockHash, "diff", block.Difficulty(), "age", common.PrettyAge(time.Unix(int64(block.Time()), 0)))
			return &eth.ForkchoiceUpdatedResult{PayloadStatus: INVALID_TERMINAL_BLOCK, PayloadID: nil}, nil
		}
	}
	valid := func(id *beacon.PayloadID) *eth.ForkchoiceUpdatedResult {
		return &eth.ForkchoiceUpdatedResult{
			PayloadStatus: eth.PayloadStatusV1{Status: eth.ExecutionValid, LatestValidHash: &state.HeadBlockHash},
			PayloadID:     id,
		}
	}
	if rawdb.ReadCanonicalHash(e.db, block.NumberU64()) != state.HeadBlockHash {
		// Block is not canonical, set head.
		if latestValid, err := e.chain.SetCanonical(block); err != nil {
			return &eth.ForkchoiceUpdatedResult{PayloadStatus: eth.PayloadStatusV1{Status: eth.ExecutionInvalid, LatestValidHash: &latestValid}}, err
		}
	} else if e.chain.CurrentBlock().Hash() == state.HeadBlockHash {
		// If the specified head matches with our local head, do nothing and keep
		// generating the payload. It's a special corner case that a few slots are
		// missing and we are requested to generate the payload in slot.
	} else if e.chain.Config().Optimism == nil { // minor Engine API divergence: allow proposers to reorg their own chain
		panic("engine not configured as optimism engine")
	}

	// If the beacon client also advertised a finalized block, mark the local
	// chain final and completely in PoS mode.
	if state.FinalizedBlockHash != (common.Hash{}) {
		// If the finalized block is not in our canonical tree, somethings wrong
		finalBlock := e.chain.GetBlockByHash(state.FinalizedBlockHash)
		if finalBlock == nil {
			e.log.Warn("Final block not available in database", "hash", state.FinalizedBlockHash)
			return STATUS_INVALID, beacon.InvalidForkChoiceState.With(errors.New("final block not available in database"))
		} else if rawdb.ReadCanonicalHash(e.db, finalBlock.NumberU64()) != state.FinalizedBlockHash {
			e.log.Warn("Final block not in canonical chain", "number", block.NumberU64(), "hash", state.HeadBlockHash)
			return STATUS_INVALID, beacon.InvalidForkChoiceState.With(errors.New("final block not in canonical chain"))
		}
		// Set the finalized block
		e.chain.SetFinalized(finalBlock)
	}
	// Check if the safe block hash is in our canonical tree, if not somethings wrong
	if state.SafeBlockHash != (common.Hash{}) {
		safeBlock := e.chain.GetBlockByHash(state.SafeBlockHash)
		if safeBlock == nil {
			e.log.Warn("Safe block not available in database")
			return STATUS_INVALID, beacon.InvalidForkChoiceState.With(errors.New("safe block not available in database"))
		}
		if rawdb.ReadCanonicalHash(e.db, safeBlock.NumberU64()) != state.SafeBlockHash {
			e.log.Warn("Safe block not in canonical chain")
			return STATUS_INVALID, beacon.InvalidForkChoiceState.With(errors.New("safe block not in canonical chain"))
		}
		// Set the safe block
		e.chain.SetSafe(safeBlock)
	}
	// If payload generation was requested, create a new block to be potentially
	// sealed by the beacon client. The payload will be requested later, and we
	// might replace it arbitrarily many times in between.
	if attr != nil {
		// Decode forceTxs. TODO: How to handle tx validity.
		forceTxs := make(types.Transactions, 0, len(attr.Transactions))
		for i, otx := range attr.Transactions {
			var tx types.Transaction
			if err := tx.UnmarshalBinary(otx); err != nil {
				return STATUS_INVALID, fmt.Errorf("transaction %d is not valid: %v", i, err)
			}
			forceTxs = append(forceTxs, &tx)
		}
		// Create an empty block first which can be used as a fallback
		block, err := e.getSealingBlock(state.HeadBlockHash, uint64(attr.Timestamp), attr.SuggestedFeeRecipient, common.Hash(attr.PrevRandao), attr.NoTxPool, forceTxs)
		if err != nil {
			e.log.Error("Failed to create payload", "err", err, "noTxPool", attr.NoTxPool, "txs", len(forceTxs), "timestamp", attr.Timestamp)
			return STATUS_INVALID, beacon.InvalidPayloadAttributes.With(err)
		}
		id := computePayloadId(state.HeadBlockHash, attr)
		payload, err := eth.BlockAsPayload(block)
		if err != nil {
			panic(fmt.Errorf("failed to convert block into execution payload: %w", err))
		}
		e.payloads.put(id, payload)
		return valid(&id), nil
	}
	return valid(nil), nil
}

func (e *Engine) NewPayload(ctx context.Context, payload *eth.ExecutionPayload) (*eth.PayloadStatusV1, error) {
	e.log.Trace("Engine API request received", "method", "ExecutePayload", "number", payload.BlockNumber, "hash", payload.BlockHash)
	txs := make([][]byte, len(payload.Transactions))
	for i, tx := range payload.Transactions {
		txs[i] = tx
	}
	block, err := beacon.ExecutableDataToBlock(beacon.ExecutableDataV1{
		ParentHash:    payload.ParentHash,
		FeeRecipient:  payload.FeeRecipient,
		StateRoot:     common.Hash(payload.StateRoot),
		ReceiptsRoot:  common.Hash(payload.ReceiptsRoot),
		LogsBloom:     payload.LogsBloom[:],
		Random:        common.Hash(payload.PrevRandao),
		Number:        uint64(payload.BlockNumber),
		GasLimit:      uint64(payload.GasLimit),
		GasUsed:       uint64(payload.GasUsed),
		Timestamp:     uint64(payload.Timestamp),
		ExtraData:     payload.ExtraData,
		BaseFeePerGas: payload.BaseFeePerGas.ToBig(),
		BlockHash:     payload.BlockHash,
		Transactions:  txs,
	})
	if err != nil {
		log.Debug("Invalid NewPayload params", "params", payload, "error", err)
		return &eth.PayloadStatusV1{Status: eth.ExecutionInvalidBlockHash}, nil
	}
	// If we already have the block locally, ignore the entire execution and just
	// return a fake success.
	if block := e.chain.GetBlockByHash(payload.BlockHash); block != nil {
		e.log.Warn("Ignoring already known beacon payload", "number", payload.BlockNumber, "hash", payload.BlockHash, "age", common.PrettyAge(time.Unix(int64(block.Time()), 0)))
		hash := block.Hash()
		return &eth.PayloadStatusV1{Status: eth.ExecutionValid, LatestValidHash: &hash}, nil
	}

	// TODO: skipping invalid ancestor check (i.e. not remembering previously failed blocks)

	parent := e.chain.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		// TODO: hack, saying we accepted if we don't know the parent block. Might want to return critical error if we can't actually sync.
		return &eth.PayloadStatusV1{Status: eth.ExecutionAccepted, LatestValidHash: nil}, nil
	}
	if !e.chain.HasBlockAndState(block.ParentHash(), block.NumberU64()-1) {
		e.log.Warn("State not available, ignoring new payload")
		return &eth.PayloadStatusV1{Status: eth.ExecutionAccepted}, nil
	}
	if err := e.chain.InsertBlockWithoutSetHead(block); err != nil {
		e.log.Warn("NewPayloadV1: inserting block failed", "error", err)
		// TODO not remembering the payload as invalid
		return e.invalid(err, parent.Header()), nil
	}
	hash := block.Hash()
	return &eth.PayloadStatusV1{Status: eth.ExecutionValid, LatestValidHash: &hash}, nil
}

func (e *Engine) invalid(err error, latestValid *types.Header) *eth.PayloadStatusV1 {
	currentHash := e.chain.CurrentBlock().Hash()
	if latestValid != nil {
		// Set latest valid hash to 0x0 if parent is PoW block
		currentHash = common.Hash{}
		if latestValid.Difficulty.BitLen() == 0 {
			// Otherwise set latest valid hash to parent hash
			currentHash = latestValid.Hash()
		}
	}
	errorMsg := err.Error()
	return &eth.PayloadStatusV1{Status: eth.ExecutionInvalid, LatestValidHash: &currentHash, ValidationError: &errorMsg}
}

func (e *Engine) PayloadByHash(ctx context.Context, hash common.Hash) (*eth.ExecutionPayload, error) {
	block := e.chain.GetBlockByHash(hash)
	if block == nil {
		return nil, ethereum.NotFound
	}
	payload, err := eth.BlockAsPayload(block)
	if err != nil {
		panic(fmt.Errorf("failed to convert block to payload: %w", err))
	}
	return payload, nil
}

func (e *Engine) PayloadByNumber(ctx context.Context, u uint64) (*eth.ExecutionPayload, error) {
	block := e.chain.GetBlockByNumber(u)
	if block == nil {
		return nil, ethereum.NotFound
	}
	payload, err := eth.BlockAsPayload(block)
	if err != nil {
		panic(fmt.Errorf("failed to convert block to payload: %w", err))
	}
	return payload, nil
}

func (e *Engine) L2BlockRefHead(ctx context.Context) (eth.L2BlockRef, error) {
	block := e.chain.CurrentBlock()
	return l2.BlockToBlockRef(block, &e.rollupCfg.Genesis)
}

func (e *Engine) L2BlockRefByHash(ctx context.Context, l2Hash common.Hash) (eth.L2BlockRef, error) {
	block := e.chain.GetBlockByHash(l2Hash)
	if block == nil {
		return eth.L2BlockRef{}, ethereum.NotFound
	}
	return l2.BlockToBlockRef(block, &e.rollupCfg.Genesis)
}
