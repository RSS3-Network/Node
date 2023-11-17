package node

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/constant"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	config *engine.Config
	source engine.Source
	worker engine.Worker
}

func (s *Server) Run(ctx context.Context, config config.File) error {
	var (
		// TODO Develop a more effective solution to implement back pressure.
		tasksChan = make(chan []engine.Task, 1)
		errorChan = make(chan error)
	)

	zap.L().Info("register node")

	err := s.registerNode(config)
	if err != nil {
		zap.L().Error("register node", zap.Error(err))
	}

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
	for _, feed := range feeds {
		zap.L().Debug("feed", zap.Any("feed", feed))
	}

	return nil
}

func NewServer(config *engine.Config) (server *Server, err error) {
	instance := Server{
		config: config,
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

func (s *Server) registerNode(c config.File) error {
	maintainer := schema.Maintainer{
		OwnerEVMAddress: c.Discovery.Maintainer.OwnerEVMAddress,
		Signature:       c.Discovery.Maintainer.Signature,
		Website:         c.Discovery.Maintainer.Website,
	}

	node := schema.Node{
		Endpoint:  c.Discovery.Node.Endpoint,
		AccessKey: c.Discovery.Node.AccessKey,
	}

	var rss schema.RSS

	if c.Config.Component.Data.IndexerModule.RSS.RSSHub.InstanceURL != "" {
		rss = schema.RSS{
			RSSHub: schema.RSSHub{
				InstanceURL: c.Config.Component.Data.IndexerModule.RSS.RSSHub.InstanceURL,
			},
		}
	}

	indexerModule := schema.IndexerModule{
		RSS: rss,
	}

	component := schema.Component{
		Data: schema.Data{
			IndexerModule: indexerModule,
		},
	}

	data := schema.RegistrationData{
		Version:    viper.GetString("discovery.version"),
		Maintainer: maintainer,
		Node:       node,
		Component:  component,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal data: %w", err)
	}

	requestURL := viper.GetString("config.hub.endpoint") + "node/register"

	request, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, requestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("send register request to hub error: %w", err)
	}

	request.Header.Set("Content-Type", "application/json")

	//nolint:bodyclose
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("send register request to hub error: %w", err)
	}

	lo.Try(response.Body.Close)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("send register request to hub error, unexpected status code: %d", response.StatusCode)
	}

	return nil
}
