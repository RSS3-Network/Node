package worker

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/worker/contract/1inch"
	"github.com/rss3-network/node/internal/engine/worker/contract/aave"
	"github.com/rss3-network/node/internal/engine/worker/contract/aavegotchi"
	"github.com/rss3-network/node/internal/engine/worker/contract/crossbell"
	"github.com/rss3-network/node/internal/engine/worker/contract/curve"
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
	"github.com/rss3-network/node/internal/engine/worker/contract/savm"
	"github.com/rss3-network/node/internal/engine/worker/contract/stargate"
	"github.com/rss3-network/node/internal/engine/worker/contract/uniswap"
	"github.com/rss3-network/node/internal/engine/worker/contract/vsl"
	"github.com/rss3-network/node/internal/engine/worker/core"
	"github.com/rss3-network/node/schema/worker"
)

func New(config *config.Module, databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	switch config.Worker {
	case worker.Core:
		return core.NewWorker(config, redisClient)
	case worker.Mirror:
		return mirror.NewWorker(config, databaseClient)
	case worker.RSS3:
		return rss3.NewWorker(config)
	case worker.Paragraph:
		return paragraph.NewWorker(config)
	case worker.OpenSea:
		return opensea.NewWorker(config)
	case worker.Uniswap:
		return uniswap.NewWorker(config)
	case worker.Optimism:
		return optimism.NewWorker(config)
	case worker.Aavegotchi:
		return aavegotchi.NewWorker(config)
	case worker.Lens:
		return lens.NewWorker(config)
	case worker.Looksrare:
		return looksrare.NewWorker(config)
	case worker.Matters:
		return matters.NewWorker(config)
	case worker.Momoka:
		return momoka.NewWorker(config)
	case worker.Highlight:
		return highlight.NewWorker(config)
	case worker.Aave:
		return aave.NewWorker(config)
	case worker.IQWiki:
		return iqwiki.NewWorker(config)
	case worker.Lido:
		return lido.NewWorker(config)
	case worker.Crossbell:
		return crossbell.NewWorker(config)
	case worker.ENS:
		return ens.NewWorker(config, databaseClient)
	case worker.Oneinch:
		return oneinch.NewWorker(config)
	case worker.KiwiStand:
		return kiwistand.NewWorker(config)
	case worker.SAVM:
		return savm.NewWorker(config)
	case worker.VSL:
		return vsl.NewWorker(config)
	case worker.Stargate:
		return stargate.NewWorker(config)
	case worker.Curve:
		return curve.NewWorker(config, redisClient)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
