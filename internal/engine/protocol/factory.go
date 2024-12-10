package protocol

import (
	"fmt"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/protocol/activitypub"
	"github.com/rss3-network/node/internal/engine/protocol/arweave"
	"github.com/rss3-network/node/internal/engine/protocol/atproto"
	"github.com/rss3-network/node/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/internal/engine/protocol/farcaster"
	"github.com/rss3-network/node/internal/engine/protocol/near"
	"github.com/rss3-network/protocol-go/schema/network"
)

// New creates a new protocol.
func New(config *config.Module, sourceFilter engine.DataSourceFilter, checkpoint *engine.Checkpoint, databaseClient database.Client, redisClient rueidis.Client) (engine.DataSource, error) {
	switch config.Network.Protocol() {
	case network.EthereumProtocol:
		return ethereum.NewSource(config, sourceFilter, checkpoint, redisClient)
	case network.ArweaveProtocol:
		return arweave.NewSource(config, sourceFilter, checkpoint, redisClient)
	case network.FarcasterProtocol:
		return farcaster.NewSource(config, checkpoint, databaseClient)
	case network.ActivityPubProtocol:
		return activitypub.NewSource(config, checkpoint, databaseClient)
	case network.NearProtocol:
		return near.NewSource(config, sourceFilter, checkpoint, redisClient)
	case network.ATProtocol:
		return atproto.NewSource(config, sourceFilter, checkpoint, databaseClient)
	default:
		return nil, fmt.Errorf("unsupported network protocol %s", config.Network)
	}
}
