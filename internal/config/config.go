package config

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/spf13/viper"
)

const (
	Environment = "environment"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

type File struct {
	Environment string         `mapstructure:"environment"`
	Node        *engine.Config `mapstructure:"node"`
}

func Setup(configFilePath string) (*File, error) {
	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var configFile File
	if err := viper.Unmarshal(&configFile); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	return &configFile, nil
}

func init() {
	viper.AutomaticEnv()
}
