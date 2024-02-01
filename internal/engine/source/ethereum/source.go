package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	// The block time in Ethereum mainnet is designed to be approximately 15 seconds.
	defaultBlockTime = 15 * time.Second
)

var _ engine.Source = (*source)(nil)

type source struct {
	config         *config.Module
	option         *Option
	filter         *Filter
	ethereumClient ethereum.Client
	state          State
	pendingState   State
}

func (s *source) Network() filter.Network {
	return s.config.Network
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
		switch {
		case s.filter.LogAddresses != nil || s.filter.LogTopics != nil:
			errorChan <- s.pollLogs(ctx, tasksChan)
		default:
			errorChan <- s.pollBlocks(ctx, tasksChan)
		}
	}()
}

func (s *source) initialize(ctx context.Context) (err error) {
	if s.ethereumClient, err = ethereum.Dial(ctx, s.config.Endpoint); err != nil {
		return fmt.Errorf("dial to ethereum rpc endpoint: %w", err)
	}

	return nil
}

func (s *source) pollBlocks(ctx context.Context, tasksChan chan<- []engine.Task) error {
	var blockNumberLatestLocal uint64

	if s.option.BlockNumberStart != nil && s.option.BlockNumberStart.Uint64() > s.state.BlockNumber {
		s.pendingState.BlockNumber = s.option.BlockNumberStart.Uint64()
		s.state.BlockNumber = s.option.BlockNumberStart.Uint64()
	}

	for {
		if s.option.BlockNumberTarget != nil && s.option.BlockNumberTarget.Uint64() < s.state.BlockNumber {
			break
		}

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

		// nolint:prealloc
		var receipts []*ethereum.Receipt

		// handle crossbell, fantom and arbitrum blockchain cause lack of `eth_getBlockReceipts` implementation
		switch s.config.Network {
		case filter.NetworkCrossbell, filter.NetworkFantom, filter.NetworkArbitrum:
			receipts, err = s.getBlockReceipts(ctx, block)
			if err != nil {
				return fmt.Errorf("get receipts by block number %d: %w", block.Number, err)
			}
		default:
			// Before being able to handle block reorganization, correctly handle canonical.
			receipts, err = s.ethereumClient.BlockReceipts(ctx, block.Number)
			if err != nil {
				return fmt.Errorf("get receipts by block number %d: %w", block.Number, err)
			}
		}

		tasks, err := s.buildTasks(block, receipts)
		if err != nil {
			return fmt.Errorf("build tasks for block hash: %s: %w", block.Hash, err)
		}

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		// Update state by two phase commit to avoid data inconsistency.
		s.state = s.pendingState

		s.pendingState.BlockHash = block.Hash
		s.pendingState.BlockNumber++
	}

	return nil
}

func (s *source) pollLogs(ctx context.Context, tasksChan chan<- []engine.Task) error {
	var blockNumberLatestLocal uint64

	if s.option.BlockNumberStart != nil && s.option.BlockNumberStart.Uint64() > s.state.BlockNumber {
		s.pendingState.BlockNumber = s.option.BlockNumberStart.Uint64()
		s.state.BlockNumber = s.option.BlockNumberStart.Uint64()
	}

	for {
		if s.option.BlockNumberTarget != nil && s.option.BlockNumberTarget.Uint64() < s.state.BlockNumber {
			break
		}

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

		// Build log filter by the filter config.
		logFilter := ethereum.Filter{
			FromBlock: new(big.Int).SetUint64(s.state.BlockNumber),
			ToBlock:   new(big.Int).SetUint64(s.state.BlockNumber),
			Addresses: s.filter.LogAddresses,
			Topics: [][]common.Hash{
				s.filter.LogTopics,
			},
		}

		// Get logs by filter.
		logs, err := s.ethereumClient.FilterLogs(ctx, logFilter)
		if err != nil {
			return fmt.Errorf("get logs by filter: %w", err)
		}

		transactionHashes := lo.Map(logs, func(log *ethereum.Log, _ int) common.Hash {
			return log.TransactionHash
		})

		var block *ethereum.Block

		if len(logs) != 0 {
			blockHash := logs[0].BlockHash

			if block, err = s.ethereumClient.BlockByHash(ctx, blockHash); err != nil {
				return fmt.Errorf("get block by hash %s: %w", blockHash, err)
			}

			// nolint:prealloc
			var receipts []*ethereum.Receipt

			// handle crossbell, fantom and arbitrum blockchain cause lack of `eth_getBlockReceipts` implementation
			switch s.config.Network {
			case filter.NetworkCrossbell, filter.NetworkFantom, filter.NetworkArbitrum:
				receipts, err = s.getBlockReceipts(ctx, block)
				if err != nil {
					return fmt.Errorf("get receipts by block number %d: %w", block.Number, err)
				}
			default:
				// Before being able to handle block reorganization, correctly handle canonical.
				receipts, err = s.ethereumClient.BlockReceipts(ctx, block.Number)
				if err != nil {
					return fmt.Errorf("get receipts by block number %d: %w", block.Number, err)
				}
			}

			// Filter receipts by transaction hashes of logs.
			receipts = lo.Filter(receipts, func(receipt *ethereum.Receipt, _ int) bool {
				return lo.Contains(transactionHashes, receipt.TransactionHash)
			})

			// Remove transactions for the block if the receipt has been filtered.
			block.Transactions = lo.Filter(block.Transactions, func(transaction *ethereum.Transaction, _ int) bool {
				return lo.ContainsBy(receipts, func(receipt *ethereum.Receipt) bool {
					return receipt.TransactionHash == transaction.Hash
				})
			})

			tasks, err := s.buildTasks(block, receipts)
			if err != nil {
				return fmt.Errorf("build tasks for block hash: %s: %w", block.Hash, err)
			}

			// TODO It might be possible to use generics to avoid manual type assertions.
			tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })
		}

		// If there are no logs, only update the block number.
		if block == nil {
			if block, err = s.ethereumClient.BlockByNumber(ctx, new(big.Int).SetUint64(s.state.BlockNumber)); err != nil {
				return fmt.Errorf("get block by number %d: %w", s.state.BlockNumber, err)
			}

			s.state = s.pendingState

			s.pendingState.BlockHash = block.Hash
			s.pendingState.BlockNumber++

			// Push an empty task slice to the channel to update the block number.
			tasksChan <- make([]engine.Task, 0)

			continue
		}

		// Update state by two phase commit to avoid data inconsistency.
		s.state = s.pendingState

		s.pendingState.BlockHash = block.Hash
		s.pendingState.BlockNumber++
	}

	return nil
}

// use `eth_getTransactionReceipt`
func (s *source) getBlockReceipts(ctx context.Context, block *ethereum.Block) ([]*ethereum.Receipt, error) {
	// nolint:prealloc
	var receipts []*ethereum.Receipt

	for _, tx := range block.Transactions {
		receipt, err := s.ethereumClient.TransactionReceipt(ctx, tx.Hash)
		if err != nil {
			return nil, fmt.Errorf("get receipt by transaction hash %s: %w", tx.Hash, err)
		}

		receipts = append(receipts, receipt)
	}

	return receipts, nil
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

		chain, err := filter.EthereumChainIDString(s.Network().String())
		if err != nil {
			return nil, fmt.Errorf("unsupported chain %s", s.Network())
		}

		task := Task{
			Network:     s.Network(),
			ChainID:     uint64(chain),
			Header:      block.Header(),
			Transaction: transaction,
			Receipt:     receipt,
		}

		tasks[index] = &task
	}

	return tasks, nil
}

func NewSource(config *config.Module, sourceFilter engine.SourceFilter, checkpoint *engine.Checkpoint) (engine.Source, error) {
	var (
		state State
		err   error
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := source{
		config:       config,
		filter:       new(Filter), // Set a default filter for the source.
		state:        state,
		pendingState: state, // Default pending state is equal to the current state.
	}

	// Initialize filter.
	if sourceFilter != nil {
		var ok bool
		if instance.filter, ok = sourceFilter.(*Filter); !ok {
			return nil, fmt.Errorf("invalid source filter type %T", sourceFilter)
		}
	}

	if instance.option, err = NewOption(config.Parameters); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	zap.L().Info("apply option", zap.Any("option", instance.option))

	return &instance, nil
}
