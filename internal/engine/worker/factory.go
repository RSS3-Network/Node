package worker

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/ethereum"
)

const (
	NameFallbackEthereum = "fallback.ethereum"
)

func New(name string, config *engine.Config) (engine.Worker, error) {
	switch name {
	case NameFallbackEthereum:
		return ethereum.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported source %s", name)
	}
}
