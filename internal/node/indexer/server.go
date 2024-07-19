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
	"github.com/rss3-network/node/internal/engine/worker/decentralized"
	"github.com/rss3-network/node/internal/node/monitor"
	"github.com/rss3-network/node/internal/stream"
	decentralizedx "github.com/rss3-network/node/schema/worker/decentralized"
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
	source         engine.DataSource
	worker         engine.Worker
	databaseClient database.Client
	streamClient   stream.Client
	monitorClient  monitor.Client
	redisClient    rueidis.Client
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

// currentBlockMetricHandler gets the current block height/number from the checkpoint state and get latest block height/number from the network rpc.
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
			var state monitor.CheckpointState
			if err := json.Unmarshal(latestCheckpoint.State, &state); err != nil {
				zap.L().Error("unmarshal checkpoint state", zap.Error(err))
				return
			}

			var current uint64

			currentBlockHeight, currentBlockTimestamp := s.monitorClient.CurrentState(state)

			if s.worker.Name() == decentralizedx.Momoka.String() {
				current = currentBlockTimestamp
			} else {
				current = currentBlockHeight
			}

			observer.Observe(int64(current), metric.WithAttributes(
				attribute.String("service", constant.Name),
				attribute.String("worker", s.worker.Name()),
			))
		}
	}()

	return nil
}

// latestBlockMetricHandler gets the latest block height/number from the network rpc.
func (s *Server) latestBlockMetricHandler(ctx context.Context, observer metric.Int64Observer) error {
	go func() {
		var latest uint64

		// get latest block height
		latestBlockHeight, latestBlockTimestamp, err := s.monitorClient.LatestState(ctx)
		if err != nil {
			zap.L().Error("get latest block height", zap.Error(err))
			return
		}

		if s.worker.Name() == decentralizedx.Momoka.String() {
			latest = latestBlockTimestamp
		} else {
			latest = latestBlockHeight
		}

		observer.Observe(int64(latest), metric.WithAttributes(
			attribute.String("service", constant.Name),
			attribute.String("worker", s.worker.Name())))
	}()

	return nil
}

func NewServer(ctx context.Context, config *config.Module, databaseClient database.Client, streamClient stream.Client, redisClient rueidis.Client) (server *Server, err error) {
	instance := Server{
		id:             config.ID,
		config:         config,
		databaseClient: databaseClient,
		streamClient:   streamClient,
		redisClient:    redisClient,
	}

	// Initialize worker.
	if instance.worker, err = worker.New(instance.config, databaseClient, instance.redisClient); err != nil {
		return nil, fmt.Errorf("new worker: %w", err)
	}

	switch config.Network.Source() {
	case network.ArweaveSource:
		instance.monitorClient, err = monitor.NewArweaveClient()
		if err != nil {
			return nil, fmt.Errorf("new arweave monitorClient: %w", err)
		}
	case network.FarcasterSource:
		instance.monitorClient, err = monitor.NewArweaveClient()
		if err != nil {
			return nil, fmt.Errorf("new arweave monitorClient: %w", err)
		}
	case network.EthereumSource:
		instance.monitorClient, err = monitor.NewEthereumClient(config.Endpoint)
		if err != nil {
			return nil, fmt.Errorf("new ethereum monitorClient: %w", err)
		}
	case network.RSSSource:
		instance.monitorClient, err = monitor.NewRssClient(config.EndpointID, config.Parameters)
		if err != nil {
			return nil, fmt.Errorf("new rss monitorClient: %w", err)
		}
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
	if instance.source, err = source.New(instance.config, instance.worker.Filter(), checkpoint, databaseClient, redisClient); err != nil {
		return nil, fmt.Errorf("new source: %w", err)
	}

	return &instance, nil
}
