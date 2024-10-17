package worker

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/worker/federated/core/mastodon"
	"github.com/rss3-network/node/schema/worker/federated"
)

func New(config *config.Module, databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	switch config.Worker.(federated.Worker) {
	case federated.Core:
		return mastodon.NewWorker(databaseClient, redisClient)
	default:
		return nil, fmt.Errorf("[federated/factory.go] unsupported worker %s", config.Worker)
	}
}
