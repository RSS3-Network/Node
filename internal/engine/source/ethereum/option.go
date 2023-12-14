package ethereum

import (
	"math/big"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
)

type Option struct {
	BlockNumberStart  *big.Int `yaml:"block_number_start"`
	BlockNumberTarget *big.Int `yaml:"block_number_target"`
}

func NewOption(options engine.Options) (*Option, error) {
	var instance Option

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	return nil, nil
}
