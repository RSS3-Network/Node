package ethereum

import (
	"github.com/naturalselectionlabs/rss3-node/config"
	"math/big"
)

type Option struct {
	BlockNumberStart  *big.Int `yaml:"block_number_start"`
	BlockNumberTarget *big.Int `yaml:"block_number_target"`
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
