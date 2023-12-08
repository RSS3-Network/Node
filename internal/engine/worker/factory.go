package worker

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback"
)

func New(config *engine.Config) (engine.Worker, error) {
	switch config.Worker {
	case engine.Fallback:
		return fallback.NewWorker(config)
	case engine.Mirror:
		return mirror.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
