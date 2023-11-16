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
	Discovery     Discovery     `mapstructure:"discovery"`
	Config        Config        `mapstructure:"config"`
	Observability Observability `mapstructure:"observability"`
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
	Database      DatabaseConfig      `mapstructure:"database"`
	Redis         RedisConfig         `mapstructure:"redis"`
	ObjectStorage ObjectStorageConfig `mapstructure:"object_storage"`
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

// RedisConfig holds Redis configuration.
type RedisConfig struct {
	External     bool   `mapstructure:"external"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DatabaseName int    `mapstructure:"database_name"`
}

// ObjectStorageConfig holds object storage configuration.
type ObjectStorageConfig struct {
	External   bool             `mapstructure:"external"`
	Type       string           `mapstructure:"type"`
	Endpoint   EndpointConfig   `mapstructure:"endpoint"`
	Bucket     string           `mapstructure:"bucket"`
	Credential CredentialConfig `mapstructure:"credential"`
}

// EndpointConfig holds endpoints for object storage.
type EndpointConfig struct {
	Private string `mapstructure:"private"`
	Public  string `mapstructure:"public"`
}

// CredentialConfig holds credentials for object storage.
type CredentialConfig struct {
	ID     string `mapstructure:"id"`
	Secret string `mapstructure:"secret"`
}

// ComponentConfig holds component configuration.
type ComponentConfig struct {
	Data   DataComponentConfig `mapstructure:"data"`
	Search interface{}         `mapstructure:"search"`
}

// DataComponentConfig holds data component configuration.
type DataComponentConfig struct {
	IndexerModule IndexerModuleConfig `mapstructure:"indexer_module"`
}

// IndexerModuleConfig holds indexer module configuration.
type IndexerModuleConfig struct {
	Decentralized DecentralizedConfig `mapstructure:"decentralized"`
	RSS           RSSConfig           `mapstructure:"RSS"`
	Federated     interface{}         `mapstructure:"federated"`
}

// DecentralizedConfig holds decentralized configuration.
type DecentralizedConfig struct {
	Network NetworkConfig   `mapstructure:"network"`
	Storage []StorageConfig `mapstructure:"storage"`
}

// NetworkConfig holds network specific configuration.
type NetworkConfig struct {
	Indexer []IndexerConfig `mapstructure:"indexer"`
}

// IndexerConfig holds indexer specific configuration.
type IndexerConfig struct {
	Type      string                   `mapstructure:"type"`
	Args      []map[string]interface{} `mapstructure:"args"`
	DependsOn DependsOnConfig          `mapstructure:"depends_on"`
}

// DependsOnConfig holds dependency information.
type DependsOnConfig struct {
	Storage []string `mapstructure:"storage"`
}

// StorageConfig holds storage configuration.
type StorageConfig struct {
	ID               string `mapstructure:"id"`
	UpstreamEndpoint string `mapstructure:"upstream_endpoint"`
}

// RSSConfig holds RSS module configuration.
type RSSConfig struct {
	RSSHub RSSHubConfig `mapstructure:"RSSHub"`
}

// RSSHubConfig holds RSSHub specific configuration.
type RSSHubConfig struct {
	InstanceURL    string               `mapstructure:"instance_url"`
	Authentication AuthenticationConfig `mapstructure:"authentication"`
}

// AuthenticationConfig holds authentication configuration.
type AuthenticationConfig struct {
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	AccessKey string `mapstructure:"access_key"`
}

// Observability holds observability related configuration.
type Observability struct {
	Opentelemetry OpentelemetryConfig `mapstructure:"opentelemetry"`
}

// OpentelemetryConfig holds OpenTelemetry configuration.
type OpentelemetryConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
	Logs    LogsConfig    `mapstructure:"logs"`
	Traces  TracesConfig  `mapstructure:"traces"`
}

// MetricsConfig holds metrics configuration.
type MetricsConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

// LogsConfig holds logs configuration.
type LogsConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

// TracesConfig holds traces configuration.
type TracesConfig struct {
	Endpoint string `mapstructure:"endpoint"`
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
