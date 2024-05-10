package rss

import (
	"github.com/rss3-network/node/config"
)

type Option struct {
	Authentication OptionAuthentication `json:"authentication" mapstructure:"authentication"`
}

type OptionAuthentication struct {
	AccessKey string `json:"access_key" mapstructure:"access_key"`
}

func NewOption(options *config.Parameters) (*Option, error) {
	var instance Option

	if options == nil {
		return &instance, nil
	}

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	return &instance, nil
}
