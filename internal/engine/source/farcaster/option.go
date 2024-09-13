package farcaster

import (
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
)

type Option struct {
	APIKey *string `json:"api_key" mapstructure:"api_key"`
	// TimestampStart is the Farcaster seconds timestamp that the worker should start from.
	TimestampStart *big.Int `json:"timestamp_start" mapstructure:"timestamp_start"`
}

func NewOption(n network.Network, parameters *config.Parameters) (*Option, error) {
	var option Option

	if parameters == nil {
		return &Option{
			TimestampStart: parameter.CurrentNetworkStartBlock[n].Block,
		}, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	if option.TimestampStart == nil {
		option.TimestampStart = parameter.CurrentNetworkStartBlock[n].Block
	}

	return &option, nil
}
