package worker

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/aavegotchi"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/lens"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/looksrare"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/matters"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/opensea"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/optimism"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/paragraph"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/rss3"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/uniswap"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/farcaster"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

func New(config *config.Module, databaseClient database.Client) (engine.Worker, error) {
	switch config.Worker {
	case filter.Fallback:
		return fallback.NewWorker(config)
	case filter.Mirror:
		return mirror.NewWorker(config, databaseClient)
	case filter.Farcaster:
		return farcaster.NewWorker()
	case filter.RSS3:
		return rss3.NewWorker(config)
	case filter.Paragraph:
		return paragraph.NewWorker(config)
	case filter.OpenSea:
		return opensea.NewWorker(config)
	case filter.Uniswap:
		return uniswap.NewWorker(config)
	case filter.Optimism:
		return optimism.NewWorker(config)
	case filter.Aavegotchi:
		return aavegotchi.NewWorker(config)
	case filter.Lens:
		return lens.NewWorker(config)
	case filter.Looksrare:
		return looksrare.NewWorker(config)
	case filter.Matters:
		return matters.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
