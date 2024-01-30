package rss

import (
	"github.com/rss3-network/node/config"
)

type Option struct {
	Authentication OptionAuthentication `yaml:"authentication"`
}

type OptionAuthentication struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	AccessKey  string `yaml:"access_key"`
	AccessCode string `yaml:"access_code"`
}

func NewOption(options *config.Options) (*Option, error) {
	var instance Option

	if err := options.Decode(&instance); err != nil {
		return nil, err
	}

	return &instance, nil
}
