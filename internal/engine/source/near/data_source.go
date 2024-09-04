package near

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/near"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

const (
	// The block time in Near mainnet is designed to be approximately 1 min.
	defaultBlockTime = 1 * time.Second
)

// Ensure that dataSource implements DataSource.
var _ engine.DataSource = (*dataSource)(nil)

type dataSource struct {
	config      *config.Module
	option      *Option
	filter      *Filter
	nearClient  near.Client
	redisClient rueidis.Client
	state       State
}

func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	// Initialize dataSource.
	if err := s.initialize(ctx); err != nil {
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)

		return
	}

	// Start a goroutine to poll blocks.
	go func() {
		retryableFunc := func() error {
			if err := s.pollBlocks(ctx, tasksChan, s.filter); err != nil {
				return fmt.Errorf("poll blocks: %w", err)
			}

			return nil
		}

		err := retry.Do(retryableFunc,
			retry.Attempts(0),
			retry.Delay(time.Second),
			retry.DelayType(retry.BackOffDelay),
			retry.MaxDelay(5*time.Minute),
			retry.OnRetry(func(n uint, err error) {
				zap.L().Error("retry near dataSource start", zap.Uint("retry", n), zap.Error(err))
			}),
		)
		if err != nil {
			errorChan <- err
		}
	}()
}

// initialize initializes the dataSource.
func (s *dataSource) initialize(ctx context.Context) (err error) {
	// Initialize near client.
	if s.nearClient, err = near.Dial(ctx, s.config.Endpoint.URL); err != nil {
		return fmt.Errorf("create near client: %w", err)
	}

	return nil
}

// initializeBlockHeights initializes block heights.
func (s *dataSource) initializeBlockHeights() {
	if s.option.BlockStart != nil && s.option.BlockStart.Uint64() > s.state.BlockHeight {
		s.state.BlockHeight = s.option.BlockStart.Uint64()
	}
}

// updateBlockHeight checks and updates the dataSource's block height
// if the remote start block is greater than the current block height.
func (s *dataSource) updateBlockHeight(ctx context.Context) {
	remoteBlockStart, err := parameter.GetNetworkBlockStart(ctx, s.redisClient, s.config.Network.String())
	if err != nil {
		zap.L().Error("get network block start from cache", zap.Error(err))
		return
	}

	if remoteBlockStart > s.state.BlockHeight {
		s.state.BlockHeight = remoteBlockStart
		zap.L().Info("Updated block height from remote", zap.Uint64("newBlockHeight", s.state.BlockHeight))
	}
}

// pollBlocks polls blocks from arweave network.
func (s *dataSource) pollBlocks(ctx context.Context, tasksChan chan<- *engine.Tasks, _ *Filter) error {
	var (
		blockHeightLatestRemote int64
		err                     error
	)

	// Get start block height from config
	// if not set, use default value 0
	s.initializeBlockHeights()

	// Get target block height from config
	// if not set, use the latest block height from arweave network
	if s.option.BlockTarget != nil {
		zap.L().Info("block height target", zap.Uint64("block.height.target", s.option.BlockTarget.Uint64()))
		blockHeightLatestRemote = int64(s.option.BlockTarget.Uint64())
	} else {
		// Get remote block height from arweave network.
		blockHeightLatestRemote, err = s.nearClient.GetBlockHeight(ctx)
		if err != nil {
			return fmt.Errorf("get latest block height: %w", err)
		}

		zap.L().Info("get latest block height", zap.Int64("block.height", blockHeightLatestRemote))
	}

	for {
		if s.option.BlockTarget != nil && s.option.BlockTarget.Uint64() <= s.state.BlockHeight {
			break
		}

		// Check if remote block start is larger than current one
		// if so, update the current block height to remote block start
		s.updateBlockHeight(ctx)

		// Check if block height is latest.
		if s.state.BlockHeight >= uint64(blockHeightLatestRemote) {
			// Get the latest block height from arweave network for reconfirming.
			if blockHeightLatestRemote, err = s.nearClient.GetBlockHeight(ctx); err != nil {
				return fmt.Errorf("get latest block height: %w", err)
			}

			zap.L().Info("get latest block height", zap.Int64("block.height", blockHeightLatestRemote))

			if s.state.BlockHeight >= uint64(blockHeightLatestRemote) {
				// Wait for the next block on arweave network.
				time.Sleep(defaultBlockTime)
			}

			continue
		}

		// If the block height of state is 0, force to start from 0.
		blockHeightStart := lo.Ternary(s.state.BlockHeight == 0, 0, s.state.BlockHeight+1)

		// Pull blocks
		blockHeightEnd := lo.Min([]uint64{
			uint64(blockHeightLatestRemote),
			blockHeightStart + *s.option.ConcurrentBlockRequests - 1,
		})
		// Pull blocks by range.
		blocks, err := s.batchPullBlocksByRange(ctx, blockHeightStart, blockHeightEnd)
		if err != nil {
			return fmt.Errorf("batch pull blocks: %w", err)
		}

		tasks := &engine.Tasks{}

		// Get chunks from blocks
		for _, block := range blocks {
			for _, chunkHash := range block.Chunks {
				chunk, err := s.nearClient.ChunkByHash(ctx, chunkHash.ChunkHash)
				if err != nil {
					return fmt.Errorf("get chunk by hash: %w", err)
				}

				for _, transaction := range chunk.Transactions {
					// Build tasks from chunks
					tasks.Tasks = append(tasks.Tasks, &Task{
						Network:     s.Network(),
						Block:       *block,
						Transaction: transaction,
					})
				}
			}

			// Update BlockTimestamp in state
			s.state.BlockTimestamp = uint64(time.Duration(block.Header.Timestamp).Seconds())
		}

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- tasks

		// Update block height to state.
		s.state.BlockHeight = blockHeightEnd
	}

	return nil
}

// batchPullBlocksByRange pulls blocks by range, from local state block height to remote block height.
func (s *dataSource) batchPullBlocksByRange(ctx context.Context, blockHeightStart, blockHeightEnd uint64) ([]*near.Block, error) {
	zap.L().Info("begin to batch pull blocks by range", zap.Uint64("block.height.start", blockHeightStart), zap.Uint64("block.height.end", blockHeightEnd))

	// Generate block heights
	blockHeights := lo.Map(lo.RangeWithSteps(blockHeightStart, blockHeightEnd+1, 1), func(blockHeight uint64, _ int) *big.Int {
		return new(big.Int).SetUint64(blockHeight)
	})

	return s.batchPullBlocks(ctx, blockHeights)
}

// batchPullBlocks pulls blocks by block heights.
func (s *dataSource) batchPullBlocks(ctx context.Context, blockHeights []*big.Int) ([]*near.Block, error) {
	zap.L().Info("begin to pull blocks", zap.Int("blocks", len(blockHeights)))

	resultPool := pool.NewWithResults[*near.Block]().
		WithContext(ctx).
		WithCancelOnError().
		WithMaxGoroutines(int(lo.FromPtr(s.option.ConcurrentBlockRequests)))

	for _, blockHeight := range blockHeights {
		blockHeight := blockHeight

		resultPool.Go(func(ctx context.Context) (*near.Block, error) {
			retryableFunc := func() (*near.Block, error) {
				return s.nearClient.BlockByHeight(ctx, blockHeight)
			}

			return retry.DoWithData(
				retryableFunc,
				retry.Attempts(defaultRetryAttempts),
				retry.Delay(defaultRetryDelay),
				retry.DelayType(retry.BackOffDelay),
				retry.OnRetry(func(attempt uint, err error) {
					zap.L().Error("retry pull block", zap.Stringer("block.height", blockHeight), zap.Uint("attempt", attempt), zap.Error(err))
				}),
			)
		})
	}

	return resultPool.Wait()
}

// NewSource creates a new arweave dataSource.
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

	return &instance, nil
}
