package worker

import (
	"fmt"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/worker/contract/1inch"
	"github.com/rss3-network/node/internal/engine/worker/contract/aave"
	"github.com/rss3-network/node/internal/engine/worker/contract/aavegotchi"
	"github.com/rss3-network/node/internal/engine/worker/contract/crossbell"
	"github.com/rss3-network/node/internal/engine/worker/contract/ens"
	"github.com/rss3-network/node/internal/engine/worker/contract/highlight"
	"github.com/rss3-network/node/internal/engine/worker/contract/iqwiki"
	"github.com/rss3-network/node/internal/engine/worker/contract/kiwistand"
	"github.com/rss3-network/node/internal/engine/worker/contract/lens"
	"github.com/rss3-network/node/internal/engine/worker/contract/lido"
	"github.com/rss3-network/node/internal/engine/worker/contract/looksrare"
	"github.com/rss3-network/node/internal/engine/worker/contract/matters"
	"github.com/rss3-network/node/internal/engine/worker/contract/mirror"
	"github.com/rss3-network/node/internal/engine/worker/contract/momoka"
	"github.com/rss3-network/node/internal/engine/worker/contract/opensea"
	"github.com/rss3-network/node/internal/engine/worker/contract/optimism"
	"github.com/rss3-network/node/internal/engine/worker/contract/paragraph"
	"github.com/rss3-network/node/internal/engine/worker/contract/rss3"
	"github.com/rss3-network/node/internal/engine/worker/contract/uniswap"
	"github.com/rss3-network/node/internal/engine/worker/fallback"
	"github.com/rss3-network/node/internal/engine/worker/farcaster"
	"github.com/rss3-network/protocol-go/schema/filter"
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
	case filter.Momoka:
		return momoka.NewWorker(config)
	case filter.Highlight:
		return highlight.NewWorker(config)
	case filter.Aave:
		return aave.NewWorker(config)
	case filter.IQWiki:
		return iqwiki.NewWorker(config)
	case filter.Lido:
		return lido.NewWorker(config)
	case filter.Crossbell:
		return crossbell.NewWorker(config)
	case filter.ENS:
		return ens.NewWorker(config, databaseClient)
	case filter.Oneinch:
		return oneinch.NewWorker(config)
	case filter.KiwiStand:
		return kiwistand.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
