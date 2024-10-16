package core

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/arweave"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/ethereum"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/farcaster"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/near"
	"github.com/rss3-network/protocol-go/schema/network"
)

// NewWorker creates a new core worker.
func NewWorker(config *config.Module, databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	switch config.Network.Source() {
	case network.EthereumSource:
		return ethereum.NewWorker(config, redisClient)
	case network.ArweaveSource:
		return arweave.NewWorker(config)
	case network.FarcasterSource:
		return farcaster.NewWorker()
	case network.NearSource:
		return near.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Network)
	}
}
