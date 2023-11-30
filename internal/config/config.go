package config

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/spf13/viper"
)

const (
	Environment = "environment"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

type File struct {
	Environment string           `mapstructure:"environment" validate:"required" default:"development"`
	Node        []*engine.Config `mapstructure:"node" validate:"required,gt=0,dive,required"`
	Database    *database.Config `mapstructure:"database" validate:"required"`
}

func Setup(configFilePath string) (*File, error) {
	viper.SetConfigFile(configFilePath)

	// Read config file.
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var configFile File

	// Parse config file.
	if err := viper.Unmarshal(&configFile); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	// Set default values.
	if err := defaults.Set(&configFile); err != nil {
		return nil, fmt.Errorf("set default values: %w", err)
	}

	// Validate config values.
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&configFile); err != nil {
		return nil, fmt.Errorf("validate config file: %w", err)
	}

	// Parse node config.
	for _, nodeConfig := range configFile.Node {
		if err := nodeConfig.Parse(); err != nil {
			return nil, fmt.Errorf("parse node config: %w", err)
		}
	}

	return &configFile, nil
}

func init() {
	viper.AutomaticEnv()
}
