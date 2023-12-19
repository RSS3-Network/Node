package fallback

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/arweave"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

// NewWorker creates a new fallback worker.
func NewWorker(config *engine.Config) (engine.Worker, error) {
	switch config.Network.Source() {
	case filter.NetworkEthereumSource:
		return ethereum.NewWorker(config)
	case filter.NetworkArweaveSource:
		return arweave.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Network)
	}
}
