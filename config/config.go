package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
	"github.com/rss3-network/serving-node/internal/database"
	"github.com/rss3-network/serving-node/internal/stream"
	"github.com/rss3-network/serving-node/schema/filter"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

const (
	Environment = "environment"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

type File struct {
	Environment   string     `yaml:"environment" validate:"required" default:"development"`
	Discovery     *Discovery `yaml:"discovery" validate:"required"`
	Node          *Node      `yaml:"component" validate:"required"`
	Database      *Database  `yaml:"database" validate:"required"`
	Stream        *Stream    `yaml:"stream" validate:"required"`
	Observability *Telemetry `yaml:"observability" validate:"required"`
}

type Discovery struct {
	Maintainer *Maintainer `yaml:"maintainer"`
	Server     *Server     `yaml:"server"`
}

type Maintainer struct {
	EvmAddress common.Address `yaml:"evm_address"`
	Signature  string         `yaml:"signature"`
}

type Server struct {
	Endpoint              string `yaml:"endpoint"`
	GlobalIndexerEndpoint string `yaml:"global_indexer_endpoint"`
}

type Node struct {
	RSS           []*Module `yaml:"rss" validate:"dive"`
	Federated     []*Module `yaml:"federated" validate:"dive"`
	Decentralized []*Module `yaml:"decentralized" validate:"dive"`
}

type Module struct {
	Network      filter.Network `yaml:"network" validate:"required"`
	Endpoint     string         `yaml:"endpoint" validate:"required"`
	IPFSGateways []string       `yaml:"ipfs_gateways"`
	Worker       filter.Name    `yaml:"worker"`
	Parameters   *Options       `yaml:"parameters"`
}

type Database struct {
	Driver    database.Driver `mapstructure:"driver" validate:"required" default:"cockroachdb"`
	Partition *bool           `mapstructure:"partition" validate:"required" default:"true"`
	URI       string          `mapstructure:"uri" validate:"required" default:"postgres://root@localhost:26257/defaultdb"`
}

type Stream struct {
	Enable *bool         `mapstructure:"enable" validate:"required" default:"false"`
	Driver stream.Driver `mapstructure:"driver" validate:"required" default:"kafka"`
	Topic  string        `mapstructure:"topic" validate:"required" default:"rss3.node.feeds"`
	URI    string        `mapstructure:"uri" validate:"required" default:"localhost:9092"`
}

type Telemetry struct {
	OpenTelemetry *OpenTelemetryConfig `yaml:"opentelemetry" validate:"required"`
}

type OpenTelemetryConfig struct {
	Metrics *OpenTelemetryMetricsConfig `yaml:"metrics" validate:"required"`
	Traces  *OpenTelemetryTracesConfig  `yaml:"traces" validate:"required"`
}

type OpenTelemetryMetricsConfig struct {
	Enable   bool   `yaml:"enable" default:"false"`
	Endpoint string `yaml:"endpoint" default:"0.0.0.0:9090"`
}

type OpenTelemetryTracesConfig struct {
	Enable   bool   `yaml:"enable" default:"false"`
	Insecure bool   `yaml:"insecure" default:"false"`
	Endpoint string `yaml:"endpoint" validate:"required" default:"0.0.0.0:4318"`
}

// ID returns the unique identifier of the configuration.
func (c *Module) ID() string {
	id := fmt.Sprintf("%s.%s", c.Network, c.Worker)

	if c.Parameters != nil {
		var buffer map[string]any

		lo.Must0(c.Parameters.Decode(&buffer))

		if buffer != nil {
			id = fmt.Sprintf("%s.%s", id, string(lo.Must(json.Marshal(buffer))))
		}
	}

	return id
}

var _ fmt.Stringer = (*Options)(nil)

type Options struct {
	*yaml.Node
}

func (o *Options) UnmarshalYAML(node *yaml.Node) error {
	o.Node = node

	return nil
}

func (o *Options) String() string {
	var buffer map[string]any

	lo.Must0(o.Decode(&buffer))

	if buffer == nil {
		return "{}"
	}

	for key, value := range buffer {
		if value == nil {
			delete(buffer, key)
		}
	}

	return string(lo.Must(json.Marshal(buffer)))
}

func Setup(configName string) (*File, error) {
	var (
		config []byte
		err    error
	)

	configPaths := []string{"/etc/serving-node/", os.Getenv("HOME") + "/.serving-node/", "./config/", "./deploy/"}

	// Read config file.
	for _, configPath := range configPaths {
		configFilePath := filepath.Join(configPath, configName)

		config, err = os.ReadFile(configFilePath)
		if err == nil {
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	// Unmarshal config file.
	var configFile File
	if err := yaml.Unmarshal(config, &configFile); err != nil {
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

	return &configFile, nil
}
