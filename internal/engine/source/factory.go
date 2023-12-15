package source

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/arweave"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

// New creates a new source.
func New(config *engine.Config, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.Source, error) {
	switch config.Network.Source() {
	case filter.NetworkEthereumSource:
		return ethereum.NewSource(config, checkpoint)
	case filter.NetworkArweaveSource:
		return arweave.NewSource(config, checkpoint)
	case filter.NetworkFarcasterSource:
		return farcaster.NewSource(config, checkpoint, databaseClient)
	default:
		return nil, fmt.Errorf("unsupported network source %s", config.Network)
	}
}
