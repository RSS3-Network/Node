package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	// The block time in Ethereum mainnet is designed to be approximately 15 seconds.
	defaultBlockTime = 15 * time.Second
)

var _ engine.Source = (*source)(nil)

type source struct {
	config         *engine.Config
	ethereumClient *ethclient.Client
	state          State
	pendingState   State
}

func (s *source) Chain() filter.Chain {
	return filter.ChainEthereumMainnet
}

func (s *source) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *source) Start(ctx context.Context, tasksChan chan<- []engine.Task, errorChan chan<- error) {
	if err := s.initialize(ctx); err != nil {
		errorChan <- fmt.Errorf("initialize source: %w", err)

		return
	}

	go func() {
		errorChan <- s.pollBlocks(ctx, tasksChan)
	}()
}

func (s *source) initialize(ctx context.Context) (err error) {
	if s.ethereumClient, err = ethclient.Dial(s.config.Endpoint); err != nil {
		return fmt.Errorf("dial rpc endpoint: %w", err)
	}

	if s.ethereumClient, err = ethclient.Dial(s.config.Endpoint); err != nil {
		return fmt.Errorf("dial ethereum endpoint: %w", err)
	}

	chainID, err := s.ethereumClient.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("get chain id: %w", err)
	}

	if s.Chain().ID() != chainID.Uint64() {
		return fmt.Errorf("mismatch between local chain id %d and remote chain id %d", s.Chain().ID(), chainID.Uint64())
	}

	return nil
}

func (s *source) pollBlocks(ctx context.Context, tasksChan chan<- []engine.Task) error {
	var blockNumberLatestLocal uint64

	// TODO End polling at completion if target block height is specified.
	for {
		// The local block number is equal than the remote block number.
		if s.state.BlockNumber >= blockNumberLatestLocal || blockNumberLatestLocal == 0 {
			// Refresh the remote block number.
			blockNumberLatestRemote, err := s.ethereumClient.BlockNumber(ctx)
			if err != nil {
				return fmt.Errorf("get latest block number: %w", err)
			}

			// RPC providers may incorrectly shunt the request to a lagging node.
			if blockNumberLatestRemote <= blockNumberLatestLocal {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestLocal), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			} else {
				// TODO Need to handle block reorganization.
				blockNumberLatestLocal = blockNumberLatestRemote
			}

			continue
		}

		block, err := s.ethereumClient.BlockByNumber(ctx, new(big.Int).SetUint64(s.state.BlockNumber))
		if err != nil {
			return fmt.Errorf("get block by number %d: %w", s.state.BlockNumber, err)
		}

		// Before being able to handle block reorganization, correctly handle canonical.
		receipts, err := s.ethereumClient.BlockReceipts(ctx, rpc.BlockNumberOrHashWithHash(block.Hash(), true))
		if err != nil {
			return fmt.Errorf("get receipts by block hash %s: %w", block.Hash(), err)
		}

		tasks, err := s.buildTasks(block, receipts)
		if err != nil {
			return fmt.Errorf("build tasks for block hash: %s: %w", block.Hash(), err)
		}

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		// Update state by two phase commit to avoid data inconsistency.
		s.state = s.pendingState

		s.pendingState.BlockHash = block.Hash()
		s.pendingState.BlockNumber++
	}
}

func (s *source) buildTasks(block *types.Block, receipts types.Receipts) ([]*Task, error) {
	tasks := make([]*Task, len(block.Transactions()))

	for index, transaction := range block.Transactions() {
		// There is no guarantee that the receipts provided by RPC will be in the same order as the transactions,
		// so instead of using a transaction index, we can match the hash.
		receipt, exists := lo.Find(receipts, func(receipt *types.Receipt) bool {
			return receipt.TxHash == transaction.Hash()
		})

		// TODO Often this is also caused by a lagging RPC node failing to get the latest data,
		// and it's best to fix this before the build task.
		if !exists {
			return nil, fmt.Errorf("no receipt matched to transaction hash %s", transaction.Hash())
		}

		task := Task{
			Chain:       filter.ChainEthereum(s.state.ChainID),
			Header:      block.Header(),
			Transaction: transaction,
			Receipt:     receipt,
		}

		tasks[index] = &task
	}

	return tasks, nil
}

func NewSource(config *engine.Config, checkpoint *engine.Checkpoint) (engine.Source, error) {
	var state State

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := source{
		config: config,
		state:  state,
	}

	return &instance, nil
}
