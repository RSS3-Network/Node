package worker

import (
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
)

func New(_ *config.Module) (engine.Worker, error) {

	// Not implemented
	// If the rsshub worker exists, it will be started by default, and there will not be a separate worker instance

	return nil, nil
}
