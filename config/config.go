package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/stream"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

const (
	EnvPrefix = "NODE"

	Environment = "environment"

	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"

	EnvironmentPyroscopeEndpoint = EnvPrefix + "_PYROSCOPE_ENDPOINT"
)

type File struct {
	Environment   string              `mapstructure:"environment" validate:"required" default:"development"`
	Endpoints     map[string]Endpoint `mapstructure:"endpoints"`
	Discovery     *Discovery          `mapstructure:"discovery" validate:"required"`
	Component     *Component          `mapstructure:"component" validate:"required"`
	Database      *Database           `mapstructure:"database" validate:"required"`
	Stream        *Stream             `mapstructure:"stream"`
	Redis         *Redis              `mapstructure:"redis" validate:"required"`
	Observability *Telemetry          `mapstructure:"observability"`
}

// LoadModulesEndpoint loads the endpoint url and headers for each module.
// If the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.
func (f *File) LoadModulesEndpoint() error {
	assignEndpoint := func(modules []*Module) {
		for index, module := range modules {
			endpoint, found := f.Endpoints[module.EndpointID]

			if !found {
				endpoint = Endpoint{
					URL: module.EndpointID,
				}
			}

			modules[index].Endpoint = endpoint
		}
	}

	if f.Component.RSS != nil {
		assignEndpoint([]*Module{f.Component.RSS})
	}

	assignEndpoint(f.Component.Decentralized)
	assignEndpoint(f.Component.Federated)

	return nil
}

type Discovery struct {
	Operator *Operator `mapstructure:"operator"`
	Server   *Server   `mapstructure:"server"`
}

type Operator struct {
	EvmAddress common.Address `mapstructure:"evm_address"`
	Signature  string         `mapstructure:"signature"`
}

type Server struct {
	Endpoint              string `mapstructure:"endpoint"`
	GlobalIndexerEndpoint string `mapstructure:"global_indexer_endpoint"`
	AccessToken           string `mapstructure:"access_token"`
}

type Component struct {
	RSS           *Module   `mapstructure:"rss"`
	Federated     []*Module `mapstructure:"federated" validate:"dive"`
	Decentralized []*Module `mapstructure:"decentralized" validate:"dive"`
}

type Module struct {
	ID           string          `mapstructure:"id"`
	Network      network.Network `mapstructure:"network" validate:"required"`
	EndpointID   string          `mapstructure:"endpoint"`
	IPFSGateways []string        `mapstructure:"ipfs_gateways"`
	Worker       worker.Worker   `mapstructure:"worker"`
	Parameters   *Parameters     `mapstructure:"parameters"`
	Endpoint     Endpoint        `mapstructure:"-"`
}

type Endpoint struct {
	URL           string            `mapstructure:"url"`
	HTTPHeaders   map[string]string `mapstructure:"http_headers"`
	HTTP2Disabled bool              `mapstructure:"http2_disabled"`
}

// BuildEthereumOptions builds the custom options to be supplied to an ethereum client.
func (e Endpoint) BuildEthereumOptions() []ethereum.Option {
	options := make([]ethereum.Option, 0)

	if e.HTTP2Disabled {
		options = append(options, ethereum.WithHTTP2Disabled())
	}

	if len(e.HTTPHeaders) > 0 {
		options = append(options, ethereum.WithHTTPHeader(e.HTTPHeaders))
	}

	return options
}

type Database struct {
	CoveragePeriod int             `mapstructure:"coverage_period" validate:"min=3,max=12" default:"3"`
	Driver         database.Driver `mapstructure:"driver" validate:"required" default:"postgres"`
	Partition      *bool           `mapstructure:"partition" validate:"required" default:"true"`
	URI            string          `mapstructure:"uri" validate:"required" default:"postgres://postgres@localhost:5432/postgres"`
}

type Stream struct {
	Enable *bool         `mapstructure:"enable" validate:"required" default:"false"`
	Driver stream.Driver `mapstructure:"driver" validate:"required" default:"kafka"`
	Topic  string        `mapstructure:"topic" validate:"required" default:"rss3.node.activities"`
	URI    string        `mapstructure:"uri" validate:"required" default:"localhost:9092"`
}

type Telemetry struct {
	OpenTelemetry *OpenTelemetryConfig `mapstructure:"opentelemetry" validate:"required"`
}

type OpenTelemetryConfig struct {
	Metrics *OpenTelemetryMetricsConfig `mapstructure:"metrics" validate:"required"`
	Traces  *OpenTelemetryTracesConfig  `mapstructure:"traces" validate:"required"`
}

type OpenTelemetryMetricsConfig struct {
	Enable   bool   `mapstructure:"enable" default:"false"`
	Endpoint string `mapstructure:"endpoint" default:"0.0.0.0:9090"`
}

type OpenTelemetryTracesConfig struct {
	Enable   bool   `mapstructure:"enable" default:"false"`
	Insecure bool   `mapstructure:"insecure" default:"false"`
	Endpoint string `mapstructure:"endpoint" validate:"required" default:"0.0.0.0:4318"`
}

type Redis struct {
	Endpoint     string   `mapstructure:"endpoint" default:"localhost:6379" validate:"required"`
	Username     string   `mapstructure:"username"`
	Password     string   `mapstructure:"password"`
	DisableCache bool     `mapstructure:"disable_cache" default:"true"`
	TLS          RedisTLS `mapstructure:"tls"`
}

type RedisTLS struct {
	Enabled            bool   `mapstructure:"enabled" default:"false"`
	CAFile             string `mapstructure:"ca_file"`
	CertFile           string `mapstructure:"cert_file"`
	KeyFile            string `mapstructure:"key_file"`
	InsecureSkipVerify bool   `mapstructure:"insecure_skip_verify" default:"false"`
}

// var _ fmt.Stringer = (*Parameters)(nil)

type Parameters map[string]any

func (p *Parameters) String() string {
	var buffer map[string]any

	lo.Must0(p.Decode(&buffer))

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

func (p *Parameters) Decode(v interface{}) error {
	jsonStr, err := json.Marshal(*p)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonStr, v)
}

func Setup(configName string) (*File, error) {
	configType := path.Ext(configName)[1:]

	switch configType {
	case "yaml", "yml":
		return _Setup(configName, "yaml", viper.GetViper())
	case "json", "hcl", "toml":
		return _Setup(configName, configType, viper.GetViper())
	}

	return nil, fmt.Errorf("unsupported config type: %s", configType)
}

func _Setup(configName, configType string, v *viper.Viper) (*File, error) {
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	v.AddConfigPath("/etc/rss3/node/")
	v.AddConfigPath(path.Join(os.Getenv("HOME"), ".rss3", "node"))

	if currentDir, err := os.Getwd(); err == nil {
		v.AddConfigPath(path.Join(currentDir, "config"))
		v.AddConfigPath(path.Join(currentDir, "deploy"))
	}

	v.SetEnvPrefix(EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	// Explicitly bind environment variables
	err := v.BindEnv("discovery.server.access_token")
	if err != nil {
		return nil, err
	}

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal config file.
	var configFile File
	if err := v.Unmarshal(&configFile, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		network.HookFunc(),
		worker.HookFunc(),
		EvmAddressHookFunc(),
	))); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	// Use a function to load the endpoint for each module, because mapstructure doesn't support the use of custom unmarshaler.
	// Reference https://github.com/mitchellh/mapstructure/issues/115.
	if err := configFile.LoadModulesEndpoint(); err != nil {
		return nil, fmt.Errorf("build endpoint for modules: %w", err)
	}

	// Set default values.
	if err := defaults.Set(&configFile); err != nil {
		return nil, fmt.Errorf("set default values: %w", err)
	}

	// Explicitly set the access token from the environment if it exists
	if envAccessToken := v.GetString("discovery.server.access_token"); envAccessToken != "" {
		configFile.Discovery.Server.AccessToken = envAccessToken
	}

	// validate config values.
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&configFile); err != nil {
		return nil, fmt.Errorf("validate config file: %w", err)
	}

	return &configFile, nil
}

func EvmAddressHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		// data type
		f reflect.Type,
		// target data type
		t reflect.Type,
		// raw data
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t.Kind() != reflect.TypeOf(common.Address{}).Kind() {
			return data, nil
		}

		return common.HexToAddress(data.(string)), nil
	}
}
