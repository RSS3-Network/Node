package ethereum

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/samber/lo"
)

const (
	defaultRPCThreadBlocks = uint64(8)
)

type Option struct {
	BlockNumberStart  *big.Int `json:"block_number_start" mapstructure:"block_number_start"`
	BlockNumberTarget *big.Int `json:"block_number_target" mapstructure:"block_number_target"`

	// RPCThreadBlocks is the number of concurrent RPC requests associated with the blocks.
	RPCThreadBlocks *uint64 `json:"rpc_thread_blocks" mapstructure:"rpc_thread_blocks"`
}

func NewOption(options *config.Parameters) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	// Set default values.
	if instance.RPCThreadBlocks == nil {
		instance.RPCThreadBlocks = lo.ToPtr(defaultRPCThreadBlocks)
	}

	return &instance, nil
}
