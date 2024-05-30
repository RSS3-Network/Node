package arweave

import (
	"fmt"
	"math/big"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

const (
	defaultConcurrentBlockRequests = uint64(1)
	defaultRetryAttempts           = uint(10)
	defaultRetryDelay              = 500 * time.Millisecond
)

type Option struct {
	// BlockStart is the starting block height on arweave chain.
	BlockStart *big.Int `json:"block_start" mapstructure:"block_start"`
	// BlockTarget is the target block height on arweave chain.
	BlockTarget *big.Int `json:"block_target" mapstructure:"block_target"`
	// ConcurrentBlockRequests is the number of concurrent RPC requests associated with the blocks.
	ConcurrentBlockRequests *uint64 `json:"concurrent_block_requests" mapstructure:"concurrent_block_requests"`
}

func NewOption(n network.Network, parameters *config.Parameters) (*Option, error) {
	var option Option

	if parameters == nil {
		return &Option{
			BlockStart:              parameter.NetworkStartBlock[n],
			ConcurrentBlockRequests: lo.ToPtr(defaultConcurrentBlockRequests),
		}, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	// Set default values.
	if option.ConcurrentBlockRequests == nil {
		option.ConcurrentBlockRequests = lo.ToPtr(defaultConcurrentBlockRequests)
	}

	if *option.ConcurrentBlockRequests == 0 {
		return nil, fmt.Errorf("concurrent block requests must be greater than 0")
	}

	if option.BlockStart == nil {
		option.BlockStart = parameter.NetworkStartBlock[n]
	}

	return &option, nil
}
