package source

import (
	"fmt"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/source/arweave"
	"github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/internal/engine/source/farcaster"
	"github.com/rss3-network/protocol-go/schema/filter"
)

// New creates a new source.
func New(config *config.Module, sourceFilter engine.SourceFilter, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.Source, error) {
	switch config.Network.Source() {
	case filter.NetworkEthereumSource:
		return ethereum.NewSource(config, sourceFilter, checkpoint)
	case filter.NetworkArweaveSource:
		return arweave.NewSource(config, sourceFilter, checkpoint)
	case filter.NetworkFarcasterSource:
		return farcaster.NewSource(config, checkpoint, databaseClient)
	default:
		return nil, fmt.Errorf("unsupported network source %s", config.Network)
	}
}
