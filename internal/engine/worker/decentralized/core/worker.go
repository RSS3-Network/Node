package core

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/arweave"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/ethereum"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/farcaster"
	"github.com/rss3-network/node/internal/engine/worker/decentralized/core/near"
	"github.com/rss3-network/protocol-go/schema/network"
)

// NewWorker creates a new core worker.
func NewWorker(config *config.Module, redisClient rueidis.Client) (engine.Worker, error) {
	switch config.Network.Protocol() {
	case network.EthereumProtocol:
		return ethereum.NewWorker(config, redisClient)
	case network.ArweaveProtocol:
		return arweave.NewWorker(config)
	case network.FarcasterProtocol:
		return farcaster.NewWorker()
	case network.NearProtocol:
		return near.NewWorker(config)
	default:
		return nil, fmt.Errorf("unsupported worker %s", config.Network)
	}
}
