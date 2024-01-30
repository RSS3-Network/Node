package ethereum

import (
	"math/big"

	"github.com/rss3-network/node/config"
)

type Option struct {
	BlockNumberStart  *big.Int `json:"block_number_start" mapstructure:"block_number_start"`
	BlockNumberTarget *big.Int `json:"block_number_target" mapstructure:"block_number_target"`
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
