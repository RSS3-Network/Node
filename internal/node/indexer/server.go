package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/source"
	"github.com/rss3-network/node/internal/engine/worker"
	"github.com/rss3-network/node/internal/stream"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Server struct {
	id             string
	config         *config.Module
	source         engine.Source
	worker         engine.Worker
	databaseClient database.Client
	streamClient   stream.Client
	ethereumClient ethereum.Client
	arweaveClient  arweave.Client
	// meterTasksCounter is a counter of the number of tasks processed.
	// Deprecated: use meterTasksHistogram instead.
	meterTasksCounter   metric.Int64Counter
	meterTasksHistogram metric.Float64Histogram
	meterCurrentBlock   metric.Int64ObservableGauge
	meterLatestBlock    metric.Int64ObservableGauge
}

func (s *Server) Run(ctx context.Context) error {
	var (
		// TODO Develop a more effective solution to implement back pressure.
		tasksChan = make(chan *engine.Tasks)
		errorChan = make(chan error)
	)

	zap.L().Info("start node", zap.String("version", constant.BuildVersion()))

	s.source.Start(ctx, tasksChan, errorChan)

	for {
		select {
		case tasks := <-tasksChan:
			retryableFunc := func() error {
				if err := s.handleTasks(ctx, tasks); err != nil {
					return fmt.Errorf("handle tasks: %w", err)
				}

				return nil
			}

			err := retry.Do(retryableFunc,
				retry.Attempts(0),
				retry.Delay(time.Second),            // Set initial delay to 1 second.
				retry.DelayType(retry.BackOffDelay), // Use backoff delay type, increasing delay on each retry.
				retry.MaxDelay(5*time.Minute),
				retry.OnRetry(func(n uint, err error) {
					zap.L().Error("retry handle tasks", zap.Uint("retry", n), zap.Error(err))
				}),
			)
			if err != nil {
				return fmt.Errorf("retry handle tasks: %w", err)
			}
		case err := <-errorChan:
			if err != nil {
				return fmt.Errorf("an error occurred in the source: %w", err)
			}

			return nil
		}
	}
}

func (s *Server) handleTasks(ctx context.Context, tasks *engine.Tasks) error {
	// Initialize the attributes of the meter.
	meterTasksCounterAttributes := metric.WithAttributes(
		attribute.String("service", constant.Name),
		attribute.String("worker", s.worker.Name()),
		attribute.Int("tasks", tasks.Len()),
	)

	// Start a new timer to record the time it takes to handle tasks.
	taskTimer := time.Now()

	checkpoint := engine.Checkpoint{
		ID:      s.id,
		Network: s.source.Network(),
		Worker:  s.worker.Name(),
		State:   s.source.State(),
	}

	// Extract the OpenTelemetry context from the tasks.
	ctx = otel.GetTextMapPropagator().Extract(ctx, tasks)

	ctx, span := otel.Tracer("").Start(ctx, "Indexer handleTasks", trace.WithSpanKind(trace.SpanKindConsumer))
	defer span.End()

	span.SetAttributes(
		attribute.String("service", constant.Name),
		attribute.String("worker", s.worker.Name()),
		attribute.Int("tasks", tasks.Len()),
		attribute.String("state", string(checkpoint.State)),
	)

	// If no tasks are returned, only save the checkpoint to the database.
	if tasks.Len() == 0 {
		zap.L().Info("save checkpoint", zap.Any("checkpoint", checkpoint))

		if err := s.databaseClient.SaveCheckpoint(ctx, &checkpoint); err != nil {
			return fmt.Errorf("save checkpoint: %w", err)
		}

		return nil
	}

	resultPool := pool.NewWithResults[*activityx.Activity]().WithMaxGoroutines(lo.Ternary(tasks.Len() < 20*runtime.NumCPU(), tasks.Len(), 20*runtime.NumCPU()))

	for _, task := range tasks.Tasks {
		task := task

		resultPool.Go(func() *activityx.Activity {
			zap.L().Debug("start match task", zap.String("task.id", task.ID()))

			matched, err := s.worker.Match(ctx, task)
			if err != nil {
				zap.L().Error("match task", zap.String("task.id", task.ID()), zap.Error(err))

				return nil
			}

			// If the task does not meet the filter conditions, it will be discarded.
			if !matched {
				zap.L().Warn("unmatched task", zap.String("task.id", task.ID()))

				return nil
			}

			zap.L().Debug("start transform task", zap.String("task.id", task.ID()))

			activity, err := s.worker.Transform(ctx, task)
			if err != nil {
				zap.L().Error("transform task", zap.String("task.id", task.ID()), zap.Error(err))

				return nil
			}

			return activity
		})
	}

	// Filter failed activities.
	activities := lo.Filter(resultPool.Wait(), func(activity *activityx.Activity, _ int) bool {
		return activity != nil
	})

	// Deprecated: use meterTasksHistogram instead.
	s.meterTasksCounter.Add(ctx, int64(tasks.Len()), meterTasksCounterAttributes)
	checkpoint.IndexCount = int64(len(activities))

	// Save activities and checkpoint to the database.
	if err := s.databaseClient.SaveActivities(ctx, activities); err != nil {
		return fmt.Errorf("save %d activities: %w", len(activities), err)
	}

	zap.L().Info("save checkpoint", zap.Any("checkpoint", checkpoint))

	if err := s.databaseClient.SaveCheckpoint(ctx, &checkpoint); err != nil {
		return fmt.Errorf("save checkpoint: %w", err)
	}

	// Record the time it takes to handle tasks.
	duration := time.Since(taskTimer).Seconds()
	s.meterTasksHistogram.Record(ctx, duration, meterTasksCounterAttributes)

	// Push activities to the stream.
	if s.streamClient != nil && len(activities) > 0 {
		if err := s.streamClient.PushActivities(ctx, activities); err != nil {
			return fmt.Errorf("publish %d activities: %w", len(activities), err)
		}
	}

	return nil
}

func (s *Server) initializeMeter() (err error) {
	// init ethereum client
	switch s.config.Network {
	case network.Arweave:
		arweaveClient, err := arweave.NewClient()
		if err != nil {
			return fmt.Errorf("new arweave client: %w", err)
		}

		s.arweaveClient = arweaveClient
	case network.Farcaster:
		break
	default:
		ethereumClient, err := ethereum.Dial(context.Background(), s.config.Endpoint)
		if err != nil {
			return fmt.Errorf("dial ethereum: %w", err)
		}

		s.ethereumClient = ethereumClient
	}

	// init meter
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if s.meterTasksCounter, err = meter.Int64Counter("rss3_node_tasks"); err != nil {
		return fmt.Errorf("create meter of tasks counter: %w", err)
	}

	if s.meterTasksHistogram, err = meter.Float64Histogram("rss3_node_task_handle_duration_seconds", metric.WithUnit("s")); err != nil {
		return fmt.Errorf("create meter of tasks histogram: %w", err)
	}

	if s.meterCurrentBlock, err = meter.Int64ObservableGauge("rss3_node_current_block", metric.WithInt64Callback(s.currentBlockMetricHandler)); err != nil {
		return fmt.Errorf("failed to observe meter CurrentBlock: %w", err)
	}

	if s.meterLatestBlock, err = meter.Int64ObservableGauge("rss3_node_latest_block", metric.WithInt64Callback(s.latestBlockMetricHandler)); err != nil {
		return fmt.Errorf("failed to observe meter LatestBlock: %w", err)
	}

	return nil
}

func (s *Server) currentBlockMetricHandler(ctx context.Context, observer metric.Int64Observer) error {
	go func() {
		// get current block height state
		latestCheckpoint, err := s.databaseClient.LoadCheckpoint(ctx, s.id, s.source.Network(), s.worker.Name())
		if err != nil {
			zap.L().Error("find latest checkpoint", zap.Error(err))
			return
		}

		if latestCheckpoint != nil {
			// Get the current block height/block number from the checkpoint state.
			type CheckpointState struct {
				BlockHeight uint64 `json:"block_height"`
				BlockNumber uint64 `json:"block_number"`
			}

			var state CheckpointState
			if err := json.Unmarshal(latestCheckpoint.State, &state); err != nil {
				zap.L().Error("unmarshal checkpoint state", zap.Error(err))
				return
			}

			currentBlockHeight := state.BlockNumber
			if currentBlockHeight == 0 {
				currentBlockHeight = state.BlockHeight
			}

			observer.Observe(int64(currentBlockHeight), metric.WithAttributes(
				attribute.String("service", constant.Name),
				attribute.String("worker", s.worker.Name()),
			))
		}
	}()

	return nil
}

func (s *Server) latestBlockMetricHandler(ctx context.Context, observer metric.Int64Observer) error {
	go func() {
		// get latest block height
		var latestBlockHeight int64

		switch s.config.Network {
		case network.Arweave:
			// get arweave client
			latestHeight, err := s.arweaveClient.GetBlockHeight(ctx)
			if err != nil {
				zap.L().Error("get latest block height", zap.Error(err))
				return
			}

			latestBlockHeight = latestHeight
		case network.Farcaster:
			break
		default:
			latestBlock, err := s.ethereumClient.BlockNumber(ctx)
			if err != nil {
				zap.L().Error("get latest block number", zap.Error(err))
				return
			}

			latestBlockHeight = latestBlock.Int64()
		}

		observer.Observe(latestBlockHeight, metric.WithAttributes(
			attribute.String("service", constant.Name),
			attribute.String("worker", s.worker.Name())))
	}()

	return nil
}

func NewServer(ctx context.Context, config *config.Module, databaseClient database.Client, streamClient stream.Client, redisClient rueidis.Client) (server *Server, err error) {
	instance := Server{
		id:             config.ID(),
		config:         config,
		databaseClient: databaseClient,
		streamClient:   streamClient,
	}

	// Initialize worker.
	if instance.worker, err = worker.New(instance.config, databaseClient, redisClient); err != nil {
		return nil, fmt.Errorf("new worker: %w", err)
	}

	if err := instance.initializeMeter(); err != nil {
		return nil, fmt.Errorf("initialize meter: %w", err)
	}

	// Load checkpoint for initialize the source.
	checkpoint, err := instance.databaseClient.LoadCheckpoint(ctx, instance.id, config.Network, instance.worker.Name())
	if err != nil {
		return nil, fmt.Errorf("loca checkpoint: %w", err)
	}

	// Unmarshal checkpoint state to map for print it in log.
	var state map[string]any
	if err := json.Unmarshal(checkpoint.State, &state); err != nil {
		return nil, fmt.Errorf("unmarshal checkpoint state: %w", err)
	}

	zap.L().Info("load checkpoint", zap.String("checkpoint.id", checkpoint.ID), zap.String("checkpoint.network", checkpoint.Network.String()), zap.String("checkpoint.worker", checkpoint.Worker), zap.Any("checkpoint.state", state))

	// Initialize source.
	if instance.source, err = source.New(instance.config, instance.worker.Filter(), checkpoint, databaseClient); err != nil {
		return nil, fmt.Errorf("new source: %w", err)
	}

	return &instance, nil
}
