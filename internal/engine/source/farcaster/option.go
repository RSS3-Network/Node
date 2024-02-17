package farcaster

import (
	"github.com/rss3-network/node/config"
)

type Option struct {
	APIKey *string `json:"api_key" mapstructure:"api_key"`
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
