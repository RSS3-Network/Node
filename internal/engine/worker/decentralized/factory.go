package worker

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/internal/engine"
	oneinch "github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/1inch"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/aave"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/aavegotchi"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/arbitrum"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/base"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/benddao"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/cow"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/crossbell"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/curve"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/ens"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/highlight"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/iqwiki"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/kiwistand"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/lens"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/lido"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/linea"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/linear"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/looksrare"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/matters"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/mirror"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/momoka"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/nearsocial"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/nouns"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/opensea"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/optimism"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/paragraph"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/paraswap"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/polymarket"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/rainbow"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/rss3"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/savm"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/stargate"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/uniswap"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/vsl"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/zerion"
	"github.com/rss3-network/node/v2/internal/engine/worker/decentralized/core"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
)

func New(config *config.Module, databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	if config.Worker == decentralized.Core {
		return core.NewWorker(config, redisClient)
	}

	return newNonCoreWorker(config, databaseClient, redisClient)
}

//nolint:gocyclo
func newNonCoreWorker(config *config.Module, databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	switch config.Worker {
	case decentralized.Mirror:
		return mirror.NewWorker(config, databaseClient)
	case decentralized.RSS3:
		return rss3.NewWorker(config)
	case decentralized.Paragraph:
		return paragraph.NewWorker(config)
	case decentralized.OpenSea:
		return opensea.NewWorker(config)
	case decentralized.Uniswap:
		return uniswap.NewWorker(config)
	case decentralized.Arbitrum:
		return arbitrum.NewWorker(config)
	case decentralized.Optimism:
		return optimism.NewWorker(config)
	case decentralized.Aavegotchi:
		return aavegotchi.NewWorker(config)
	case decentralized.Lens:
		return lens.NewWorker(config)
	case decentralized.Looksrare:
		return looksrare.NewWorker(config)
	case decentralized.Matters:
		return matters.NewWorker(config)
	case decentralized.Momoka:
		return momoka.NewWorker(config)
	case decentralized.Nouns:
		return nouns.NewWorker(config)
	case decentralized.Highlight:
		return highlight.NewWorker(config)
	case decentralized.Aave:
		return aave.NewWorker(config)
	case decentralized.IQWiki:
		return iqwiki.NewWorker(config)
	case decentralized.Lido:
		return lido.NewWorker(config)
	case decentralized.Crossbell:
		return crossbell.NewWorker(config)
	case decentralized.ENS:
		return ens.NewWorker(config, databaseClient)
	case decentralized.Oneinch:
		return oneinch.NewWorker(config)
	case decentralized.KiwiStand:
		return kiwistand.NewWorker(config)
	case decentralized.SAVM:
		return savm.NewWorker(config)
	case decentralized.VSL:
		return vsl.NewWorker(config)
	case decentralized.Stargate:
		return stargate.NewWorker(config)
	case decentralized.Curve:
		return curve.NewWorker(config, redisClient)
	case decentralized.Paraswap:
		return paraswap.NewWorker(config)
	case decentralized.Cow:
		return cow.NewWorker(config)
	case decentralized.BendDAO:
		return benddao.NewWorker(config)
	case decentralized.Base:
		return base.NewWorker(config)
	case decentralized.Linea:
		return linea.NewWorker(config)
	case decentralized.Polymarket:
		return polymarket.NewWorker(config)
	case decentralized.LiNEAR:
		return linear.NewWorker(config)
	case decentralized.Zerion:
		return zerion.NewWorker(config)
	case decentralized.Rainbow:
		return rainbow.NewWorker(config)
	case decentralized.NearSocial:
		return nearsocial.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Worker)
	}
}
