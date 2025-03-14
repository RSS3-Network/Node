package worker

import (
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
)

func New(_ *config.Module) (engine.Worker, error) {
	// Not implemented
	// If the rsshub worker exists, it will be started by default, and there will not be a separate worker instance
	return nil, nil
}
