package source

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

func New(config *engine.Config) (engine.Source, error) {
	switch config.Network {
	case filter.NetworkEthereum:
		return ethereum.NewSource(config)
	default:
		return nil, fmt.Errorf("unsupported source %s", config.Network)
	}
}
