package arweave

import (
	"math/big"

	"github.com/rss3-network/node/config"
)

type Option struct {
	BlockHeightStart  *big.Int `yaml:"block_height_start"`
	BlockHeightTarget *big.Int `yaml:"block_height_target"`
	RPCThreadBlocks   uint64   `yaml:"rpc_thread_blocks"`
}

func NewOption(options *config.Options) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	return &instance, nil
}
