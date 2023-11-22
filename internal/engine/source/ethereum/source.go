package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
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
	ethereumClient ethereum.Client
	state          *State
}

func (s *source) Chain() filter.Chain {
	return filter.ChainEthereumMainnet
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
	// TODO Load state from checkpoint.
	s.state = &State{
		ChainID:     filter.ChainEthereumMainnet.ID(),
		BlockHash:   ethereum.HashGenesis,
		BlockNumber: 0,
	}

	if s.ethereumClient, err = ethereum.Dial(ctx, s.config.Endpoint); err != nil {
		return fmt.Errorf("dial to an ethereum rpc node: %w", err)
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
			if blockNumberLatestRemote.Uint64() <= blockNumberLatestLocal {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestLocal), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			} else {
				// TODO Need to handle block reorganization.
				blockNumberLatestLocal = blockNumberLatestRemote.Uint64()
			}

			continue
		}

		block, err := s.ethereumClient.BlockByNumber(ctx, new(big.Int).SetUint64(s.state.BlockNumber))
		if err != nil {
			return fmt.Errorf("get block by number %d: %w", s.state.BlockNumber, err)
		}

		// Before being able to handle block reorganization, correctly handle canonical.
		receipts, err := s.ethereumClient.BlockReceipts(ctx, block.Number)
		if err != nil {
			return fmt.Errorf("get receipts by block number %d: %w", block.Number, err)
		}

		tasks, err := s.buildTasks(block, receipts)
		if err != nil {
			return fmt.Errorf("build tasks for block hash: %s: %w", block.Hash, err)
		}

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		// Update local block number.
		s.state.BlockNumber++
	}
}

func (s *source) buildTasks(block *ethereum.Block, receipts []*ethereum.Receipt) ([]*Task, error) {
	tasks := make([]*Task, len(block.Transactions))

	for index, transaction := range block.Transactions {
		// There is no guarantee that the receipts provided by RPC will be in the same order as the transactions,
		// so instead of using a transaction index, we can match the hash.
		receipt, exists := lo.Find(receipts, func(receipt *ethereum.Receipt) bool {
			return receipt.TransactionHash == transaction.Hash
		})

		// TODO Often this is also caused by a lagging RPC node failing to get the latest data,
		// and it's best to fix this before the build task.
		if !exists {
			return nil, fmt.Errorf("no receipt matched to transaction hash %s", transaction.Hash)
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

func NewSource(config *engine.Config) (engine.Source, error) {
	instance := source{
		config: config,
	}

	return &instance, nil
}
