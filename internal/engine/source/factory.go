package source

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/farcaster"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

// New creates a new source.
func New(config *engine.Config, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.Source, error) {
	switch config.Chain.Network() {
	case filter.NetworkEthereum:
		return ethereum.NewSource(config, checkpoint)
	case filter.NetworkFarcaster:
		return farcaster.NewSource(config, checkpoint, databaseClient)
	default:
		return nil, fmt.Errorf("unsupported source %s", config.Chain.Network())
	}
}
