package arweave

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/samber/lo"
)

const (
	defaultRPCThreadBlocks = uint64(1)
)

type Option struct {
	BlockHeightStart  *big.Int `json:"block_height_start" mapstructure:"block_height_start"`
	BlockHeightTarget *big.Int `json:"block_height_target" mapstructure:"block_height_target"`

	// RPCThreadBlocks is the number of concurrent RPC requests associated with the blocks. Changes to this field don't
	// cause data inconsistencies and therefore don't need to be stored in the id field of the checkpoint.
	RPCThreadBlocks *uint64 `json:"rpc_thread_blocks" mapstructure:"rpc_thread_blocks"`
}

func NewOption(options *config.Options) (*Option, error) {
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
