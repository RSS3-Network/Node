package contract

import (
	"github.com/rss3-network/serving-node/config"
)

type Option struct {
	ParseTokenMetadata bool `yaml:"parse_token_metadata"`
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
