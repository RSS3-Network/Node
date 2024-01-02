package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/naturalselectionlabs/rss3-node/internal/constant"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Server struct {
	id                string
	config            *engine.Config
	source            engine.Source
	worker            engine.Worker
	databaseClient    database.Client
	tracer            trace.Tracer
	meterTasksCounter metric.Int64Counter
}

func (s *Server) Run(ctx context.Context) error {
	var (
		// TODO Develop a more effective solution to implement back pressure.
		tasksChan = make(chan []engine.Task, 1)
		errorChan = make(chan error)
	)

	zap.L().Info("start node", zap.String("version", constant.BuildVersion()))

	s.source.Start(ctx, tasksChan, errorChan)

	for {
		select {
		case tasks := <-tasksChan:
			if err := s.handleTasks(ctx, tasks); err != nil {
				return fmt.Errorf("handle tasks: %w", err)
			}
		case err := <-errorChan:
			if err != nil {
				return fmt.Errorf("an error occurred in the source: %w", err)
			}

			return nil
		}
	}
}

func (s *Server) handleTasks(ctx context.Context, tasks []engine.Task) error {
	checkpoint := engine.Checkpoint{
		ID:      s.id,
		Network: s.source.Network(),
		Worker:  s.worker.Name(),
		State:   s.source.State(),
	}

	ctx, span := s.tracer.Start(ctx, "handleTasks")
	defer span.End()

	span.SetAttributes(
		attribute.String("service", constant.Name),
		attribute.String("worker", s.worker.Name()),
		attribute.Int("records", len(tasks)),
		attribute.String("state", string(checkpoint.State)),
	)

	// If no tasks are returned, only save the checkpoint to the database.
	if len(tasks) == 0 {
		zap.L().Info("save checkpoint", zap.Any("checkpoint", checkpoint))

		if err := s.databaseClient.SaveCheckpoint(ctx, &checkpoint); err != nil {
			return fmt.Errorf("save checkpoint: %w", err)
		}

		return nil
	}

	resultPool := pool.NewWithResults[*schema.Feed]().WithMaxGoroutines(lo.Ternary(len(tasks) < 20*runtime.NumCPU(), len(tasks), 20*runtime.NumCPU()))

	for _, task := range tasks {
		task := task

		resultPool.Go(func() *schema.Feed {
			zap.L().Debug("start match task", zap.String("task.id", task.ID()))

			matched, err := s.worker.Match(ctx, task)
			if err != nil {
				zap.L().Error("matched task", zap.String("task.id", task.ID()))

				return nil
			}

			if !matched {
				zap.L().Warn("unmatched task", zap.String("task.id", task.ID()))

				return nil
			}

			zap.L().Debug("start transform task", zap.String("task.id", task.ID()))

			feed, err := s.worker.Transform(ctx, task)
			if err != nil {
				zap.L().Error("transform task", zap.String("task.id", task.ID()))

				return nil
			}

			return feed
		})
	}

	// Filter failed feeds.
	feeds := lo.Filter(resultPool.Wait(), func(feed *schema.Feed, _ int) bool {
		return feed != nil
	})

	meterTasksCounterAttributes := metric.WithAttributes(
		attribute.String("service", constant.Name),
		attribute.String("worker", s.worker.Name()),
		attribute.Int("records", len(tasks)),
		attribute.String("state", string(checkpoint.State)),
	)

	s.meterTasksCounter.Add(ctx, int64(len(tasks)), meterTasksCounterAttributes)

	// Save feeds and checkpoint to the database.
	return s.databaseClient.WithTransaction(ctx, func(ctx context.Context, client database.Client) error {
		if err := client.SaveFeeds(ctx, feeds); err != nil {
			return fmt.Errorf("save %d feeds: %w", len(feeds), err)
		}

		zap.L().Info("save checkpoint", zap.Any("checkpoint", checkpoint))

		if err := client.SaveCheckpoint(ctx, &checkpoint); err != nil {
			return fmt.Errorf("save checkpoint: %w", err)
		}

		return nil
	})
}

func (s *Server) initializeMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if s.meterTasksCounter, err = meter.Int64Counter("rss3_node_tasks"); err != nil {
		return fmt.Errorf("create meter of tasks counter: %w", err)
	}

	return nil
}

func NewServer(ctx context.Context, config *engine.Config, databaseClient database.Client) (server *Server, err error) {
	instance := Server{
		id:             config.ID(),
		config:         config,
		databaseClient: databaseClient,
		tracer:         otel.Tracer(""),
	}

	// Initialize worker.
	if instance.worker, err = worker.New(instance.config, databaseClient); err != nil {
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
