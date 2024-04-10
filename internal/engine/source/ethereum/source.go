package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
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
	// The latest block number of the remote RPC endpoint.
	var blockNumberLatestRemote uint64

	// Set the block number to the start block number if it is greater than the current block number.
	if s.option.BlockNumberStart != nil && s.option.BlockNumberStart.Uint64() > s.state.BlockNumber {
		s.state.BlockNumber = s.option.BlockNumberStart.Uint64()
	}

	for {
		// Stop the source if the block number target is not nil and the current block number is greater than the target
		// block number. This is useful when the source is used to index a specific range of blocks.
		if s.option.BlockNumberTarget != nil && s.option.BlockNumberTarget.Uint64() < s.state.BlockNumber {
			break
		}

		blockNumberStart := s.state.BlockNumber + 1

		// Check if the local block number is greater than or equal to the remote block number, or the remote block
		// number is zero. If so, sync the remote block number and wait for the new block.
		if blockNumberStart > blockNumberLatestRemote || blockNumberLatestRemote == 0 {
			// Refresh the remote block number.
			blockNumber, err := s.ethereumClient.BlockNumber(ctx)
			if err != nil {
				return fmt.Errorf("get latest block number: %w", err)
			}

			blockNumberLatestRemote = blockNumber.Uint64()

			// RPC providers may incorrectly shunt the request to a lagging node.
			if blockNumberStart > blockNumberLatestRemote {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestRemote), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			}

			continue
		}

		blockNumbers := lo.Map(lo.Range(int(*s.option.RPCThreadBlocks)), func(item int, _ int) *big.Int {
			return new(big.Int).SetUint64(blockNumberStart + uint64(item))
		})

		// Filter block numbers that are less than or equal to the latest remote block number.
		blockNumbers = lo.Filter(blockNumbers, func(blockNumber *big.Int, _ int) bool {
			return blockNumber.Uint64() <= blockNumberLatestRemote
		})

		blocks, err := s.getBlocks(ctx, blockNumbers)
		if err != nil {
			return fmt.Errorf("get blocks by block numbers: %w", err)
		}

		receipts, err := s.getReceipts(ctx, blocks)
		if err != nil {
			return fmt.Errorf("get receipts: %w", err)
		}

		// Build tasks for each block.
		var tasks []*Task

		for _, block := range blocks {
			block := block

			blockTasks, err := s.buildTasks(block, receipts)
			if err != nil {
				return fmt.Errorf("build tasks for block hash: %s: %w", block.Hash, err)
			}

			tasks = append(tasks, blockTasks...)
		}

		// Push tasks to the source.
		tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })

		latestBlock := lo.Must(lo.Last(blocks))

		s.state.BlockHash = latestBlock.Hash
		s.state.BlockNumber = latestBlock.Number.Uint64()
	}

	return nil
}

func (s *source) pollLogs(ctx context.Context, tasksChan chan<- []engine.Task) error {
	// The latest block number of the remote RPC endpoint.
	var blockNumberLatestRemote uint64

	// Set the block number to the start block number if it is greater than the current block number.
	if s.option.BlockNumberStart != nil && s.option.BlockNumberStart.Uint64() > s.state.BlockNumber {
		s.state.BlockNumber = s.option.BlockNumberStart.Uint64()
	}

	for {
		// Stop the source if the block number target is not nil and the current block number is greater than the target
		// block number. This is useful when the source is used to index a specific range of blocks.
		if s.option.BlockNumberTarget != nil && s.option.BlockNumberTarget.Uint64() < s.state.BlockNumber {
			break
		}

		blockNumberStart := s.state.BlockNumber + 1

		// Check if the local block number is greater than or equal to the remote block number, or the remote block
		// number is zero. If so, sync the remote block number and wait for the new block.
		if blockNumberStart > blockNumberLatestRemote || blockNumberLatestRemote == 0 {
			// Refresh the remote block number.
			blockNumber, err := s.ethereumClient.BlockNumber(ctx)
			if err != nil {
				return fmt.Errorf("get latest block number: %w", err)
			}

			blockNumberLatestRemote = blockNumber.Uint64()

			// RPC providers may incorrectly shunt the request to a lagging node.
			if blockNumberStart > blockNumberLatestRemote {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestRemote), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			}

			continue
		}

		blockNumberEnd := min(blockNumberStart+*s.option.RPCThreadBlocks-1, blockNumberLatestRemote)

		// Build log filter by the filter config.
		logFilter := ethereum.Filter{
			FromBlock: new(big.Int).SetUint64(blockNumberStart),
			ToBlock:   new(big.Int).SetUint64(blockNumberEnd),
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

		var latestBlock *ethereum.Block

		if len(logs) == 0 {
			// If there are no logs in the block range, update the latest block by the block number only.
			if latestBlock, err = s.ethereumClient.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumberEnd)); err != nil {
				return fmt.Errorf("get block by number %d: %w", s.state.BlockNumber, err)
			}

			// Push an empty task slice to the channel to update the block number.
			tasksChan <- make([]engine.Task, 0)
		} else {
			transactionHashes := lo.Map(logs, func(log *ethereum.Log, _ int) common.Hash {
				return log.TransactionHash
			})

			// Filter unique transaction hashes.
			transactionHashes = lo.UniqBy(transactionHashes, func(transactionHash common.Hash) common.Hash {
				return transactionHash
			})

			blockNumbers := lo.Map(logs, func(log *ethereum.Log, _ int) *big.Int {
				return log.BlockNumber
			})

			// Filter unique block numbers.
			blockNumbers = lo.UniqBy(blockNumbers, func(blockNumber *big.Int) uint64 {
				return blockNumber.Uint64()
			})

			blocks, err := s.getBlocks(ctx, blockNumbers)
			if err != nil {
				return fmt.Errorf("get blocks: %w", err)
			}

			// Filter blocks by transaction hashes of logs.
			blocks = lo.Map(blocks, func(block *ethereum.Block, _ int) *ethereum.Block {
				block.Transactions = lo.Filter(block.Transactions, func(transaction *ethereum.Transaction, _ int) bool {
					return lo.Contains(transactionHashes, transaction.Hash)
				})

				return block
			})

			if latestBlock, err = lo.Last(blocks); err != nil {
				return fmt.Errorf("get latest block: %w", err)
			}

			receipts, err := s.getReceiptsByTransactionHashes(ctx, transactionHashes)
			if err != nil {
				return fmt.Errorf("get receipts: %w", err)
			}

			tasks := make([]*Task, 0)

			for _, block := range blocks {
				blockTasks, err := s.buildTasks(block, receipts)
				if err != nil {
					return err
				}

				tasks = append(tasks, blockTasks...)
			}

			tasksChan <- lo.Map(tasks, func(task *Task, _ int) engine.Task { return task })
		}

		s.state.BlockHash = latestBlock.Hash
		s.state.BlockNumber = latestBlock.Number.Uint64()
	}

	return nil
}

// getBlocks is used to concurrently get blocks by block number.
func (s *source) getBlocks(ctx context.Context, blockNumbers []*big.Int) ([]*ethereum.Block, error) {
	resultPool := pool.NewWithResults[*ethereum.Block]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, blockNumber := range blockNumbers {
		blockNumber := blockNumber

		resultPool.Go(func(ctx context.Context) (*ethereum.Block, error) {
			return s.ethereumClient.BlockByNumber(ctx, blockNumber)
		})
	}

	blocks, err := resultPool.Wait()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(blocks, func(left, right int) bool {
		// Sort blocks by block number in ascending order.
		return blocks[left].Number.Cmp(blocks[right].Number) == -1
	})

	return blocks, nil
}

func (s *source) getReceipts(ctx context.Context, blocks []*ethereum.Block) ([]*ethereum.Receipt, error) {
	switch s.config.Network {
	case
		filter.NetworkCrossbell,
		filter.NetworkArbitrum,
		filter.NetworkSatoshiVM,
		filter.NetworkLinea,
		filter.NetworkBinanceSmartChain:
		transactionHashes := lo.Map(blocks, func(block *ethereum.Block, _ int) []common.Hash {
			return lo.Map(block.Transactions, func(transaction *ethereum.Transaction, _ int) common.Hash {
				return transaction.Hash
			})
		})

		return s.getReceiptsByTransactionHashes(ctx, lo.Flatten(transactionHashes))
	default:
		blockNumbers := lo.Map(blocks, func(block *ethereum.Block, _ int) *big.Int {
			return block.Number
		})

		return s.getReceiptsByBlockNumbers(ctx, blockNumbers)
	}
}

func (s *source) getReceiptsByBlockNumbers(ctx context.Context, blockNumbers []*big.Int) ([]*ethereum.Receipt, error) {
	resultPool := pool.NewWithResults[[]*ethereum.Receipt]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, blockNumber := range blockNumbers {
		blockNumber := blockNumber

		resultPool.Go(func(ctx context.Context) ([]*ethereum.Receipt, error) {
			return s.ethereumClient.BlockReceipts(ctx, blockNumber)
		})
	}

	receipts, err := resultPool.Wait()
	if err != nil {
		return nil, err
	}

	return lo.Flatten(receipts), nil
}

func (s *source) getReceiptsByTransactionHashes(ctx context.Context, transactionHashes []common.Hash) ([]*ethereum.Receipt, error) {
	resultPool := pool.NewWithResults[*ethereum.Receipt]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, transactionHash := range transactionHashes {
		transactionHash := transactionHash

		resultPool.Go(func(ctx context.Context) (*ethereum.Receipt, error) {
			return s.ethereumClient.TransactionReceipt(ctx, transactionHash)
		})
	}

	return resultPool.Wait()
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
		config: config,
		filter: new(Filter), // Set a default filter for the source.
		state:  state,
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
