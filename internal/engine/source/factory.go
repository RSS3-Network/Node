package source

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/arweave"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

// New creates a new source.
func New(config *engine.Config, sourceFilter engine.SourceFilter, checkpoint *engine.Checkpoint) (engine.Source, error) {
	switch config.Chain.Network() {
	case filter.NetworkEthereum:
		return ethereum.NewSource(config, sourceFilter, checkpoint)
	case filter.NetworkArweave:
		return arweave.NewSource(config, checkpoint)
	default:
		return nil, fmt.Errorf("unsupported source %s", config.Chain.Network())
	}
}
