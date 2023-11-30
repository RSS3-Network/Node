package config

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var configFile File

	// parse config file into struct
	if err := viper.Unmarshal(&configFile); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	// set default values
	if err := defaults.Set(&configFile); err != nil {
		return nil, fmt.Errorf("set default values: %w", err)
	}

	// validate config file
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&configFile); err != nil {
		return nil, fmt.Errorf("validate config file: %w", err)
	}

	zap.L().Info("config file loaded", zap.Any("config", configFile))

	return &configFile, nil
}

func init() {
	viper.AutomaticEnv()
}
