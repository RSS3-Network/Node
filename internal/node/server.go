package node

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/constant"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

type Server struct {
	config         *engine.Config
	source         engine.Source
	worker         engine.Worker
	databaseClient database.Client
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
	resultPool := pool.NewWithResults[*schema.Feed]()

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

	// Save feeds to database.
	if err := s.databaseClient.SaveFeeds(ctx, feeds); err != nil {
		return fmt.Errorf("save %d feeds: %w", len(feeds), err)
	}

	return nil
}

func NewServer(config *engine.Config, databaseClient database.Client) (server *Server, err error) {
	instance := Server{
		config:         config,
		databaseClient: databaseClient,
	}

	// TODO Implement support for sources and workers from non Ethereum networks.
	if instance.source, err = source.New(source.NameEthereum, instance.config); err != nil {
		return nil, fmt.Errorf("new source: %w", err)
	}

	if instance.worker, err = worker.New(worker.NameFallbackEthereum, instance.config); err != nil {
		return nil, fmt.Errorf("new worker: %w", err)
	}

	return &instance, nil
}
