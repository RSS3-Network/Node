package source

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
)

const (
	NameEthereum = "ethereum"
)

func New(name string, config *engine.Config) (engine.Source, error) {
	switch name {
	case NameEthereum:
		return ethereum.NewSource(config)
	default:
		return nil, fmt.Errorf("unsupported source %s", name)
	}
}
