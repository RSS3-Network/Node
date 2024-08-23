package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const (
	// The block time in Ethereum mainnet is designed to be approximately 15 seconds.
	defaultBlockTime = 15 * time.Second
)

var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config         *config.Module
	option         *Option
	filter         *Filter
	ethereumClient ethereum.Client
	redisClient    rueidis.Client
	state          State
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	if err := s.initialize(ctx); err != nil {
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)

		return
	}

	go func() {
		retryableFunc := func() error {
			switch {
			case s.filter.LogAddresses != nil || s.filter.LogTopics != nil:
				if err := s.pollLogs(ctx, tasksChan); err != nil {
					return fmt.Errorf("poll logs: %w", err)
				}
			default:
				if err := s.pollBlocks(ctx, tasksChan); err != nil {
					return fmt.Errorf("poll blocks: %w", err)
				}
			}

			return nil
		}

		err := retry.Do(retryableFunc,
			retry.Attempts(0),
			retry.Delay(time.Second),            // Set initial delay to 1 second.
			retry.DelayType(retry.BackOffDelay), // Use backoff delay type, increasing delay on each retry.
			retry.MaxDelay(5*time.Minute),
			retry.OnRetry(func(n uint, err error) {
				zap.L().Error("retry ethereum dataSource start", zap.Uint("retry", n), zap.Error(err))
			}),
		)
		if err != nil {
			errorChan <- err
		}
	}()
}

func (s *dataSource) initialize(ctx context.Context) (err error) {
	if s.ethereumClient, err = ethereum.Dial(ctx, s.config.Endpoint.URL, s.config.Endpoint.BuildEthereumOptions()...); err != nil {
		return fmt.Errorf("dial to ethereum rpc endpoint: %w", err)
	}

	return nil
}

func (s *dataSource) pollBlocks(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	// The latest block number of the remote RPC endpoint.
	var blockNumberLatestRemote uint64

	// Set the block number to the start block number if it is greater than the current block number.
	if s.option.BlockStart != nil && s.option.BlockStart.Uint64() > s.state.BlockNumber {
		s.state.BlockNumber = s.option.BlockStart.Uint64()
	}

	for {
		ctx, span := otel.Tracer("").Start(ctx, "DataSource pollBlocks", trace.WithSpanKind(trace.SpanKindProducer))

		// Stop the dataSource if the block number target is not nil and the current block number is greater than the target
		// block number. This is useful when the dataSource is used to index a specific range of blocks.
		if s.option.BlockTarget != nil && s.option.BlockTarget.Uint64() < s.state.BlockNumber {
			zap.L().Info("dataSource has indexed the specified block range", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.target", s.option.BlockTarget.Uint64()))

			break
		}

		remoteBlockStart, err := parameter.GetNetworkBlockStart(ctx, s.redisClient, s.config.Network.String())
		if err != nil {
			return fmt.Errorf("get network block start from cache: %w", err)
		}

		if remoteBlockStart > s.state.BlockNumber {
			s.state.BlockNumber = remoteBlockStart
			zap.L().Info("Updated block number from remote", zap.Uint64("newBlockNumber", s.state.BlockNumber))
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

			// No new block has been mined, or the RPC node is lagging.
			if blockNumberStart > blockNumberLatestRemote {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestRemote), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			}

			continue
		}

		blockNumbers := lo.Map(lo.Range(int(*s.option.ConcurrentBlockRequests)), func(item int, _ int) *big.Int {
			return new(big.Int).SetUint64(blockNumberStart + uint64(item))
		})

		// Filter block numbers that are less than or equal to the latest remote block number.
		blockNumbers = lo.Filter(blockNumbers, func(blockNumber *big.Int, _ int) bool {
			return blockNumber.Uint64() <= blockNumberLatestRemote
		})

		span.SetAttributes(
			attribute.Stringer("block.number.start", blockNumbers[0]),
			attribute.Stringer("block.number.end", blockNumbers[len(blockNumbers)-1]),
		)

		blocks, err := s.getBlocks(ctx, blockNumbers)
		if err != nil {
			return fmt.Errorf("get blocks by block numbers: %w", err)
		}

		receipts, err := s.getReceipts(ctx, blocks)
		if err != nil {
			return fmt.Errorf("get receipts: %w", err)
		}

		// Build tasks for each block.
		var tasks engine.Tasks

		for _, block := range blocks {
			block := block

			blockTasks, err := s.buildTasks(block, receipts)
			if err != nil {
				return fmt.Errorf("build tasks for block hash: %s: %w", block.Hash, err)
			}

			tasks.Tasks = append(tasks.Tasks, lo.Map(blockTasks, func(blockTask *Task, _ int) engine.Task { return blockTask })...)
		}

		span.End()

		// Push tasks to the dataSource.
		s.pushTasks(ctx, tasksChan, &tasks)

		latestBlock := lo.Must(lo.Last(blocks))

		s.state.BlockHash = latestBlock.Hash
		s.state.BlockNumber = latestBlock.Number.Uint64()
	}

	return nil
}

func (s *dataSource) pollLogs(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	var blockNumberLatestRemote uint64

	if s.option.BlockStart != nil && s.option.BlockStart.Uint64() > s.state.BlockNumber {
		s.state.BlockNumber = s.option.BlockStart.Uint64()
	}

	for {
		ctx, span := otel.GetTracerProvider().Tracer("").Start(ctx, "DataSource pollLogs", trace.WithSpanKind(trace.SpanKindProducer))

		// Stop the dataSource if the block number target is not nil and the current block number is greater than the target
		// block number. This is useful when the dataSource is used to index a specific range of blocks.
		if s.option.BlockTarget != nil && s.option.BlockTarget.Uint64() < s.state.BlockNumber {
			zap.L().Info("dataSource has indexed the specified block range", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.target", s.option.BlockTarget.Uint64()))

			break
		}

		remoteBlockStart, err := parameter.GetNetworkBlockStart(ctx, s.redisClient, s.config.Network.String())
		if err != nil {
			return fmt.Errorf("get network block start from cache: %w", err)
		}

		if remoteBlockStart > s.state.BlockNumber {
			s.state.BlockNumber = remoteBlockStart
			zap.L().Info("Updated block number from remote", zap.Uint64("newBlockNumber", s.state.BlockNumber))
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

			// No new block has been mined, or the RPC node is lagging.
			if blockNumberStart > blockNumberLatestRemote {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestRemote), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			}

			continue
		}

		// The block number end is the start block number plus the number of blocks to be processed in parallel.
		blockNumberEnd := min(blockNumberStart+*s.option.ConcurrentBlockRequests-1, blockNumberStart)

		span.SetAttributes(
			attribute.String("block.number.start", strconv.FormatUint(blockNumberStart, 10)),
			attribute.String("block.number.end", strconv.FormatUint(blockNumberEnd, 10)),
		)

		logFilter := ethereum.Filter{
			FromBlock: new(big.Int).SetUint64(blockNumberStart),
			ToBlock:   new(big.Int).SetUint64(blockNumberEnd),
			Addresses: s.filter.LogAddresses,
			Topics: [][]common.Hash{
				s.filter.LogTopics,
			},
		}

		logs, err := s.ethereumClient.FilterLogs(ctx, logFilter)
		if err != nil {
			return fmt.Errorf("get logs by filter: %w", err)
		}

		var latestBlock *ethereum.Block

		if len(logs) == 0 {
			latestBlock, err = s.updateLatestBlock(ctx, blockNumberEnd)
			if err != nil {
				return err
			}

			span.End()

			s.pushTasks(ctx, tasksChan, new(engine.Tasks))
		} else {
			latestBlock, err = s.processLogs(ctx, logs, tasksChan)
			if err != nil {
				return err
			}
		}

		s.state.BlockHash = latestBlock.Hash
		s.state.BlockNumber = latestBlock.Number.Uint64()
	}

	return nil
}

func (s *dataSource) updateLatestBlock(ctx context.Context, blockNumberEnd uint64) (*ethereum.Block, error) {
	latestBlock, err := s.ethereumClient.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumberEnd))
	if err != nil {
		return nil, fmt.Errorf("get block by number %d: %w", s.state.BlockNumber, err)
	}

	return latestBlock, nil
}

func (s *dataSource) processLogs(ctx context.Context, logs []*ethereum.Log, tasksChan chan<- *engine.Tasks) (*ethereum.Block, error) {
	transactionHashes := lo.Map(logs, func(log *ethereum.Log, _ int) common.Hash {
		return log.TransactionHash
	})

	transactionHashes = lo.UniqBy(transactionHashes, func(transactionHash common.Hash) common.Hash {
		return transactionHash
	})

	blockNumbers := lo.Map(logs, func(log *ethereum.Log, _ int) *big.Int {
		return log.BlockNumber
	})

	blockNumbers = lo.UniqBy(blockNumbers, func(blockNumber *big.Int) uint64 {
		return blockNumber.Uint64()
	})

	blocks, err := s.getBlocks(ctx, blockNumbers)
	if err != nil {
		return nil, fmt.Errorf("get blocks: %w", err)
	}

	blocks = lo.Map(blocks, func(block *ethereum.Block, _ int) *ethereum.Block {
		block.Transactions = lo.Filter(block.Transactions, func(transaction *ethereum.Transaction, _ int) bool {
			return lo.Contains(transactionHashes, transaction.Hash)
		})

		return block
	})

	latestBlock, exist := lo.Last(blocks)
	if !exist {
		return nil, fmt.Errorf("empty blocks")
	}

	receipts, err := s.getReceiptsByTransactionHashes(ctx, transactionHashes)
	if err != nil {
		return nil, fmt.Errorf("get receipts: %w", err)
	}

	var tasks engine.Tasks

	for _, block := range blocks {
		blockTasks, err := s.buildTasks(block, receipts)
		if err != nil {
			return nil, err
		}

		tasks.Tasks = append(tasks.Tasks, lo.Map(blockTasks, func(blockTask *Task, _ int) engine.Task { return blockTask })...)
	}

	s.pushTasks(ctx, tasksChan, &tasks)

	return latestBlock, nil
}

// getBlocks is used to concurrently get blocks by block number.
func (s *dataSource) getBlocks(ctx context.Context, blockNumbers []*big.Int) ([]*ethereum.Block, error) {
	resultPool := pool.NewWithResults[[]*ethereum.Block]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, blockNumbers := range lo.Chunk(blockNumbers, int(*s.option.BlockBatchSize)) {
		blockNumbers := blockNumbers

		resultPool.Go(func(ctx context.Context) ([]*ethereum.Block, error) {
			return s.ethereumClient.BatchBlockByNumbers(ctx, blockNumbers)
		})
	}

	batchResults, err := resultPool.Wait()
	if err != nil {
		return nil, err
	}

	blocks := lo.Flatten(batchResults)

	// Sort blocks by block number in ascending order.
	sort.SliceStable(blocks, func(left, right int) bool {
		return blocks[left].Number.Cmp(blocks[right].Number) == -1
	})

	return blocks, nil
}

func (s *dataSource) getReceipts(ctx context.Context, blocks []*ethereum.Block) ([]*ethereum.Receipt, error) {
	switch s.config.Network {
	case
		network.Crossbell,
		network.Arbitrum,
		network.SatoshiVM,
		network.Linea,
		network.XLayer:
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

func (s *dataSource) getReceiptsByBlockNumbers(ctx context.Context, blockNumbers []*big.Int) ([]*ethereum.Receipt, error) {
	resultPool := pool.NewWithResults[[]*ethereum.Receipt]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, blockNumbers := range lo.Chunk(blockNumbers, int(*s.option.BlockReceiptsBatchSize)) {
		blockNumbers := blockNumbers

		resultPool.Go(func(ctx context.Context) ([]*ethereum.Receipt, error) {
			batchReceipts, err := s.ethereumClient.BatchBlockReceipts(ctx, blockNumbers)
			if err != nil {
				return nil, err
			}

			return lo.Flatten(batchReceipts), nil
		})
	}

	batchResults, err := resultPool.Wait()
	if err != nil {
		return nil, err
	}

	return lo.Flatten(batchResults), nil
}

func (s *dataSource) getReceiptsByTransactionHashes(ctx context.Context, transactionHashes []common.Hash) ([]*ethereum.Receipt, error) {
	resultPool := pool.NewWithResults[[]*ethereum.Receipt]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, transactionHashes := range lo.Chunk(transactionHashes, int(*s.option.BlockReceiptsBatchSize)) {
		transactionHashes := transactionHashes

		resultPool.Go(func(ctx context.Context) ([]*ethereum.Receipt, error) {
			return s.ethereumClient.BatchTransactionReceipt(ctx, transactionHashes)
		})
	}

	batchResults, err := resultPool.Wait()
	if err != nil {
		return nil, err
	}

	return lo.Flatten(batchResults), nil
}

func (s *dataSource) buildTasks(block *ethereum.Block, receipts []*ethereum.Receipt) ([]*Task, error) {
	var (
		tasks  = make([]*Task, len(block.Transactions))
		header = block.Header()
	)

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

		chain, err := network.EthereumChainIDString(s.Network().String())
		if err != nil {
			return nil, fmt.Errorf("unsupported chain %s", s.Network())
		}

		task := Task{
			Network:     s.Network(),
			ChainID:     uint64(chain),
			Header:      header,
			Transaction: transaction,
			Receipt:     receipt,
		}

		tasks[index] = &task
	}

	return tasks, nil
}

func (s *dataSource) pushTasks(ctx context.Context, tasksChan chan<- *engine.Tasks, tasks *engine.Tasks) {
	otel.GetTextMapPropagator().Inject(ctx, tasks)

	_, span := otel.Tracer("").Start(ctx, "DataSource pushTasks", trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	tasksChan <- tasks
}

func NewSource(config *config.Module, sourceFilter engine.DataSourceFilter, checkpoint *engine.Checkpoint, redisClient rueidis.Client) (engine.DataSource, error) {
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

	instance := dataSource{
		config:      config,
		filter:      new(Filter), // Set a default filter for the dataSource.
		state:       state,
		redisClient: redisClient,
	}

	// Initialize filter.
	if sourceFilter != nil {
		var ok bool
		if instance.filter, ok = sourceFilter.(*Filter); !ok {
			return nil, fmt.Errorf("invalid dataSource filter type %T", sourceFilter)
		}
	}

	if instance.option, err = NewOption(config.Network, config.Parameters); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	zap.L().Info("apply option", zap.Any("option", instance.option))

	return &instance, nil
}
