package worker

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/node/v2/internal/engine/worker/federated/bluesky"
	"github.com/rss3-network/node/v2/internal/engine/worker/federated/mastodon"
	"github.com/rss3-network/node/v2/schema/worker/federated"
)

func New(config *config.Module, databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	switch config.Worker.(federated.Worker) {
	case federated.Mastodon:
		return mastodon.NewWorker(databaseClient, redisClient)
	case federated.Bluesky:
		return bluesky.NewWorker(databaseClient)
	default:
		return nil, fmt.Errorf("[federated/factory.go] unsupported worker %s", config.Worker)
	}
}
