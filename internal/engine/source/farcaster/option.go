package farcaster

import (
	"github.com/rss3-network/node/config"
)

type Option struct {
	APIKey *string `json:"api_key" mapstructure:"api_key"`
}

func NewOption(parameters *config.Parameters) (*Option, error) {
	var option Option

	if parameters == nil {
		return &option, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	return &option, nil
}
