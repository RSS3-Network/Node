package arweave

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

const (
	defaultConcurrentBlockRequests = uint64(1)
)

type Option struct {
	// BlockStart is the starting block height on arweave chain.
	BlockStart *big.Int `json:"block_start" mapstructure:"block_start"`
	// BlockTarget is the target block height on arweave chain.
	BlockTarget *big.Int `json:"block_target" mapstructure:"block_target"`
	// ConcurrentBlockRequests is the number of concurrent RPC requests associated with the blocks.
	ConcurrentBlockRequests *uint64 `json:"concurrent_block_requests" mapstructure:"concurrent_block_requests"`
}

func NewOption(n network.Network, options *config.Parameters) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	// Set default values.
	if instance.ConcurrentBlockRequests == nil {
		instance.ConcurrentBlockRequests = lo.ToPtr(defaultConcurrentBlockRequests)
	}

	if instance.BlockStart == nil {
		instance.BlockStart = parameter.NetworkStartBlock[n]
	}

	return &instance, nil
}
