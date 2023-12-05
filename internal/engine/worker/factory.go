package worker

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/rss3"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback"
)

// New creates a new worker by the given config.
func New(config *engine.Config) (engine.Worker, error) {
	switch config.Worker {
	case engine.Fallback:
		return fallback.NewWorker(config)
	case engine.RSS3:
		return rss3.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
