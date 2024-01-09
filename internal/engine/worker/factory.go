package worker

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/aavegotchi"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/lens"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/opensea"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/optimism"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/paragraph"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/rss3"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/uniswap"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/farcaster"
)

func New(config *config.Module, databaseClient database.Client) (engine.Worker, error) {
	switch config.Worker {
	case engine.Fallback:
		return fallback.NewWorker(config)
	case engine.Mirror:
		return mirror.NewWorker(config, databaseClient)
	case engine.Farcaster:
		return farcaster.NewWorker()
	case engine.RSS3:
		return rss3.NewWorker(config)
	case engine.Paragraph:
		return paragraph.NewWorker(config)
	case engine.OpenSea:
		return opensea.NewWorker(config)
	case engine.Uniswap:
		return uniswap.NewWorker(config)
	case engine.Optimism:
		return optimism.NewWorker(config)
	case engine.Aavegotchi:
		return aavegotchi.NewWorker(config)
	case engine.Lens:
		return lens.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
