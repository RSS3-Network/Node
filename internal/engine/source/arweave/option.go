package arweave

import (
	"math/big"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
)

type Option struct {
	BlockHeightStart  *big.Int
	BlockHeightTarget *big.Int
	RPCThreadBlocks   uint64
}

func NewOption(options *engine.Options) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	return &instance, nil
}
