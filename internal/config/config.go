package config

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/spf13/viper"
)

const (
	Environment = "config.environment"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

// Config represents the top-level configuration structure.
type File struct {
	Discovery Discovery `mapstructure:"discovery"`
	Config    Config    `mapstructure:"config"`
}

// Discovery represents the discovery configuration.
type Discovery struct {
	Version    string     `mapstructure:"version"`
	Maintainer Maintainer `mapstructure:"maintainer"`
	Node       NodeConfig `mapstructure:"node"`
}

// Maintainer holds maintainer specific configuration.
type Maintainer struct {
	OwnerEVMAddress string `mapstructure:"owner_evm_address"`
	Signature       string `mapstructure:"signature"`
	Website         string `mapstructure:"website"`
}

// NodeConfig holds node specific configuration.
type NodeConfig struct {
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
}

// Config holds application level configuration.
type Config struct {
	Environment string          `mapstructure:"environment"`
	Hub         HubConfig       `mapstructure:"hub"`
	Node        *engine.Config  `mapstructure:"node"`
	Data        DataConfig      `mapstructure:"data"`
	Component   ComponentConfig `mapstructure:"component"`
}

type HubConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

// DataConfig holds data related configuration.
type DataConfig struct {
	Database DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig holds database configuration.
type DatabaseConfig struct {
	Type         string `mapstructure:"type"`
	External     bool   `mapstructure:"external"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"database_name"`
}

type ComponentConfig struct {
	Data   DataComponentConfig `mapstructure:"data"`
	Search interface{}         `mapstructure:"search"`
}

type DataComponentConfig struct {
	IndexerModule IndexerModuleConfig `mapstructure:"indexer_module"`
}

type IndexerModuleConfig struct {
	RSS RSSConfig `mapstructure:"RSS"`
}

type RSSConfig struct {
	RSSHub RSSHubConfig `mapstructure:"RSSHub"`
}

type RSSHubConfig struct {
	InstanceURL    string               `mapstructure:"instance_url"`
	Authentication AuthenticationConfig `mapstructure:"authentication"`
}

type AuthenticationConfig struct {
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	AccessKey string `mapstructure:"access_key"`
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
