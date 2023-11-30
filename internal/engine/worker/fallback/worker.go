package fallback

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

// NewWorker creates a new fallback worker.
func NewWorker(config *engine.Config) (engine.Worker, error) {
	switch config.Network {
	case filter.NetworkEthereum:
		return ethereum.NewWorker(config)
	}

	return nil, fmt.Errorf("unsupported worker %s", config.Network)
}
