package atproto

import (
	"fmt"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/worker/atproto/core"
	"github.com/rss3-network/node/schema/worker/atproto"
)

func New(config *config.Module, databaseClient database.Client) (engine.Worker, error) {
	switch config.Worker {
	case atproto.Core:
		return core.NewWorker(databaseClient)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
