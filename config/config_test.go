package config

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	configName        = "config.yaml"
	configExampleYaml = `environment: development
discovery:
    operator:
      evm_address: 0x111222333444555666777888999aaabbbcccddde
      signature: 0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee
    server:
      endpoint: https://node.mydomain.com/
      global_indexer_endpoint: https://gi.rss3.dev/
      access_token: test
endpoints:
    ethereum:
      url: https://rpc.ankr.com/eth
      http_headers:
        user-agent: rss3-node
database:
  driver: postgres
  partition: true
  uri: postgres://postgres@localhost:5432/postgres
stream:
  enable: false
  driver: kafka
  topic: rss3.node.activities
  uri: localhost:9092
redis:
  endpoint: localhost:6379
  username:
  password:
  disable_cache: true
observability:
  opentelemetry:
    metrics:
      enable: true
      endpoint: 0.0.0.0:9090
    traces:
      enable: true
      insecure: true
      endpoint: localhost:4318
component:
  rss:
    network: rss
    worker: rsshub
    endpoint: https://rsshub.app/
    parameters:
      authentication:
       username: user
       password: pass
       access_key: abc
       access_code: def
  decentralized:
    - network: ethereum
      worker: core
      endpoint: ethereum
      parameters:
        block_start: 47370106
        block_target: 456
    - network: ethereum
      worker: rss3
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_start: 123
        concurrent_block_requests: 2
`
	configExampleJSON = `{
  "environment": "development",
  "endpoints": {
     "ethereum": {
      "url": "https://rpc.ankr.com/eth",
      "http_headers": {
        "user-agent": "rss3-node"
      }
    }
  },
  "discovery": {
    "operator": {
      "evm_address": "0x111222333444555666777888999aaabbbcccddde",
      "signature": "0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee"
    },
    "server": {
      "endpoint": "https://node.mydomain.com/",
      "global_indexer_endpoint": "https://gi.rss3.dev/",
      "access_token": "test"
    }
  },
  "database": {
	"coverage_period": 3,
    "driver": "postgres",
    "partition": true,
    "uri": "postgres://postgres@localhost:5432/postgres"
  },
  "stream": {
    "enable": false,
    "driver": "kafka",
    "topic": "rss3.node.activities",
    "uri": "localhost:9092"
  },
  "redis": {
    "endpoint": "localhost:6379",
    "username": "",
    "password": "",
    "disable_cache": true
  },
  "observability": {
    "opentelemetry": {
      "metrics": {
        "enable": true,
        "endpoint": "0.0.0.0:9090"
      },
      "traces": {
        "enable": true,
        "insecure": true,
        "endpoint": "localhost:4318"
      }
    }
  },
  "component": {
    "rss":
      {
        "network": "rss",
        "worker": "rsshub",
        "endpoint": "https://rsshub.app/",
        "parameters": {
          "authentication": {
            "username": "user",
            "password": "pass",
            "access_key": "abc",
            "access_code": "def"
          }
        }
      },
    "decentralized": [
      {
        "network": "ethereum",
        "worker": "core",
        "endpoint": "ethereum",
        "parameters": {
          "block_start": 47370106,
          "block_target": 456
        }
      },
      {
        "network": "ethereum",
        "worker": "rss3",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_start": 123,
          "concurrent_block_requests": 2
        }
      }
    ]
  }
}`
	configExampleToml = `environment = "development"

[discovery.operator]
evm_address = "0x111222333444555666777888999aaabbbcccddde"
signature = "0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee"

[endpoints.ethereum]
url = "https://rpc.ankr.com/eth"

	[endpoints.ethereum.http_headers]
	user-agent = "rss3-node"

[discovery.server]
endpoint = "https://node.mydomain.com/"
global_indexer_endpoint = "https://gi.rss3.dev/"
access_token = "test"

[database]
coverage_period = 3
driver = "postgres"
partition = true
uri = "postgres://postgres@localhost:5432/postgres"

[stream]
enable = false
driver = "kafka"
topic = "rss3.node.activities"
uri = "localhost:9092"

[redis]
endpoint = "localhost:6379"
username = ""
password = ""
disable_cache = true

[observability.opentelemetry.metrics]
enable = true
endpoint = "0.0.0.0:9090"

[observability.opentelemetry.traces]
enable = true
insecure = true
endpoint = "localhost:4318"

[component.rss]
network = "rss"
worker = "rsshub"
endpoint = "https://rsshub.app/"

[component.rss.parameters.authentication]
username = "user"
password = "pass"
access_key = "abc"
access_code = "def"

[[component.decentralized]]
network = "ethereum"
worker = "core"
endpoint = "ethereum"

  [component.decentralized.parameters]
  block_start = 47370106
  block_target = 456

[[component.decentralized]]
network = "ethereum"
worker = "rss3"
endpoint = "https://rpc.ankr.com/eth"

  [component.decentralized.parameters]
  block_start = 123
  concurrent_block_requests = 2

`
)

var configFileExpected = &File{
	Environment: "development",
	Endpoints: map[string]Endpoint{
		"ethereum": {
			URL: "https://rpc.ankr.com/eth",
			HTTPHeaders: map[string]string{
				"user-agent": "rss3-node",
			},
		},
	},
	Discovery: &Discovery{
		Operator: &Operator{
			EvmAddress: common.HexToAddress("0x111222333444555666777888999aaabbbcccddde"),
			Signature:  "0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee",
		},
		Server: &Server{
			Endpoint:              "https://node.mydomain.com/",
			GlobalIndexerEndpoint: "https://gi.rss3.dev/",
			AccessToken:           "test",
		},
	},
	Component: &Component{
		RSS: &Module{
			Network:    network.RSS,
			EndpointID: "https://rsshub.app/",
			Endpoint: Endpoint{
				URL: "https://rsshub.app/",
			},
			Worker: rss.RSSHub,
			Parameters: &Parameters{
				"authentication": map[string]any{
					"access_code": "def",
					"access_key":  "abc",
					"password":    "pass",
					"username":    "user",
				},
			},
		},
		Federated: nil,
		Decentralized: []*Module{
			{
				Network:    network.Ethereum,
				Worker:     decentralized.Core,
				EndpointID: "ethereum",
				Endpoint: Endpoint{
					URL: "https://rpc.ankr.com/eth",
					HTTPHeaders: map[string]string{
						"user-agent": "rss3-node",
					},
				},
				Parameters: &Parameters{
					"block_start":  47370106,
					"block_target": 456,
				},
			},
			{
				Network:    network.Ethereum,
				Worker:     decentralized.RSS3,
				EndpointID: "https://rpc.ankr.com/eth",
				Endpoint: Endpoint{
					URL: "https://rpc.ankr.com/eth",
				},
				Parameters: &Parameters{
					"block_start":               123,
					"concurrent_block_requests": 2,
				},
			},
		},
	},
	Database: &Database{
		CoveragePeriod: 3,
		Driver:         "postgres",
		Partition:      lo.ToPtr(true),
		URI:            "postgres://postgres@localhost:5432/postgres",
	},
	Stream: &Stream{
		Enable: lo.ToPtr(false),
		Driver: "kafka",
		Topic:  "rss3.node.activities",
		URI:    "localhost:9092",
	},
	Redis: &Redis{
		Endpoint:     "localhost:6379",
		Username:     "",
		Password:     "",
		DisableCache: true,
	},
	Observability: &Telemetry{
		OpenTelemetry: &OpenTelemetryConfig{
			Metrics: &OpenTelemetryMetricsConfig{
				Enable:   true,
				Endpoint: "0.0.0.0:9090",
			},
			Traces: &OpenTelemetryTracesConfig{
				Enable:   true,
				Insecure: true,
				Endpoint: "localhost:4318",
			},
		},
	},
}

func TestSetupConfig(t *testing.T) {
	t.Parallel()

	configDir := "/etc/rss3/node"
	fs := afero.NewMemMapFs()

	err := fs.Mkdir(configDir, 0o777)
	assert.NoError(t, err)

	file, err := fs.Create(path.Join(configDir, configName))
	assert.NoError(t, err)

	_, err = file.WriteString(configExampleYaml)
	require.NoError(t, err)

	v := viper.New()
	v.SetFs(fs)

	f, err := _Setup(configName, "yaml", v)
	assert.NoError(t, err)

	AssertConfig(t, f, configFileExpected)
}

//nolint:paralleltest
func TestConfigEnvOverride(t *testing.T) {
	expectEnvironment := "testing"
	expectDatabaseURI := "postgres://mock@localhost:5432/postgres"
	expectMetricsEndpoint := "127.0.0.1:9000"

	t.Setenv("NODE_ENVIRONMENT", expectEnvironment)
	t.Setenv("NODE_DATABASE_URI", expectDatabaseURI)
	t.Setenv("NODE_OBSERVABILITY_OPENTELEMETRY_METRICS_ENDPOINT", expectMetricsEndpoint)

	configDir := "/etc/rss3/node"
	fs := afero.NewMemMapFs()

	err := fs.Mkdir(configDir, 0o777)
	assert.NoError(t, err)

	file, err := fs.Create(path.Join(configDir, configName))
	assert.NoError(t, err)

	_, err = file.WriteString(configExampleYaml)
	require.NoError(t, err)

	v := viper.New()
	v.SetFs(fs)

	f, err := _Setup(configName, "yaml", v)
	assert.NoError(t, err)

	assert.Equal(t, expectEnvironment, f.Environment)
	assert.Equal(t, expectDatabaseURI, f.Database.URI)
	assert.Equal(t, expectMetricsEndpoint, f.Observability.OpenTelemetry.Metrics.Endpoint)
}

//nolint:paralleltest
func TestConfigEnvOverrideAccessToken(t *testing.T) {
	configDir := "/etc/rss3/node"
	fs := afero.NewMemMapFs()

	err := fs.Mkdir(configDir, 0o777)
	assert.NoError(t, err)

	// Create a config file without access_token
	configWithoutToken := strings.Replace(configExampleYaml,
		"access_token: test",
		"# access_token: test", 1)

	file, err := fs.Create(path.Join(configDir, configName))
	assert.NoError(t, err)

	_, err = file.WriteString(configWithoutToken)
	require.NoError(t, err)

	v := viper.New()
	v.SetFs(fs)

	// Set the environment variable
	envToken := "env_access_token"
	t.Setenv("NODE_DISCOVERY_SERVER_ACCESS_TOKEN", envToken)

	f, err := _Setup(configName, "yaml", v)
	assert.NoError(t, err)

	// Check if the access token is set from the environment variable
	assert.Equal(t, envToken, f.Discovery.Server.AccessToken)

	// Check other config values to ensure they're still correctly set
	assert.Equal(t, "development", f.Environment)
	assert.Equal(t, "postgres://postgres@localhost:5432/postgres", f.Database.URI)
}

func TestConfigFilePath(t *testing.T) {
	t.Parallel()

	currentDir, err := os.Getwd()
	assert.NoError(t, err)

	configPaths := []string{
		"/etc/rss3/node/",
		path.Join(os.Getenv("HOME"), ".rss3", "node"),
		path.Join(currentDir, "config"),
		path.Join(currentDir, "deploy"),
	}

	for _, configPath := range configPaths {
		func(_path string) {
			t.Run(_path, func(t *testing.T) {
				t.Parallel()

				fs := afero.NewMemMapFs()

				err := fs.Mkdir(_path, 0o777)
				assert.NoError(t, err)

				file, err := fs.Create(path.Join(_path, configName))
				assert.NoError(t, err)

				_, err = file.WriteString(configExampleYaml)
				require.NoError(t, err)

				v := viper.New()
				v.SetFs(fs)

				f, err := _Setup(configName, "yaml", v)
				assert.NoError(t, err)

				assert.Equal(t, configFileExpected, f)
			})
		}(configPath)
	}
}

func TestConfigFileType(t *testing.T) {
	t.Parallel()

	configDir := "/etc/rss3/node"
	configFiles := map[string]string{
		"yaml": configExampleYaml,
		"yml":  configExampleYaml,
		"json": configExampleJSON,
		"toml": configExampleToml,
	}

	for configType, configContext := range configFiles {
		func(_type, context string) {
			t.Run(_type, func(t *testing.T) {
				t.Parallel()

				fs := afero.NewMemMapFs()

				err := fs.Mkdir(configDir, 0o777)
				assert.NoError(t, err)

				configName := fmt.Sprintf("config.%s", _type)

				file, err := fs.Create(path.Join(configDir, configName))
				assert.NoError(t, err)

				_, err = file.WriteString(context)
				require.NoError(t, err)

				v := viper.New()
				v.SetFs(fs)

				f, err := _Setup(configName, _type, v)
				assert.NoError(t, err)

				AssertConfig(t, configFileExpected, f)
			})
		}(configType, configContext)
	}
}

func AssertConfig(t *testing.T, expect, got *File) {
	t.Run("environment", func(t *testing.T) {
		assert.Equal(t, expect.Environment, got.Environment)
	})
	t.Run("database", func(t *testing.T) {
		assert.Equal(t, expect.Database, got.Database)
	})
	t.Run("stream", func(t *testing.T) {
		assert.Equal(t, expect.Stream, got.Stream)
	})
	t.Run("discovery", func(t *testing.T) {
		assert.Equal(t, expect.Discovery, got.Discovery)
	})

	t.Run("decentralized", func(t *testing.T) {
		func(_expect, got *Module) {
			t.Run("rss", func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, _expect, got)
			})
		}(expect.Component.RSS, got.Component.RSS)

		for i, indexer := range got.Component.Decentralized {
			func(_expect, got *Module) {
				t.Run(fmt.Sprintf("%s-%s", indexer.Network, indexer.Worker), func(t *testing.T) {
					t.Parallel()
					AssertIndexer(t, _expect, got)
				})
			}(configFileExpected.Component.Decentralized[i], indexer)
		}
	})
}

func AssertIndexer(t *testing.T, got, expect *Module) {
	t.Run("network", func(t *testing.T) {
		assert.Equal(t, expect.Network, got.Network)
	})
	t.Run("worker", func(t *testing.T) {
		assert.Equal(t, expect.Worker, got.Worker)
	})
	t.Run("endpoint", func(t *testing.T) {
		assert.Equal(t, expect.Endpoint, got.Endpoint)
	})
	t.Run("gateway", func(t *testing.T) {
		assert.Equal(t, expect.IPFSGateways, got.IPFSGateways)
	})
	t.Run("parameters", func(t *testing.T) {
		AssertParameters(t, *expect.Parameters, *got.Parameters)
	})
}

func AssertParameters(t *testing.T, got, expect map[string]any) {
	if expect == nil && got == nil {
		return
	}

	if expect == nil || got == nil || len(expect) != len(got) {
		t.Errorf("Maps length expect %d but got %d \n", len(expect), len(got))
		return
	}

	for key, expectedValue := range expect {
		actualValue, ok := got[key]
		if !ok {
			t.Errorf("Key %s is missing\n", key)
		}

		switch reflect.TypeOf(expectedValue).Kind() {
		case reflect.String:
			assert.Equal(t, expectedValue, actualValue)
		case reflect.Float64, reflect.Float32:
			assert.Equal(t, expectedValue.(float64), actualValue.(float64))
		case reflect.Int64:
			assert.Equal(t, expectedValue.(int64), actualValue.(int64))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			assert.Equal(t, expectedValue.(uint64), actualValue.(uint64))
		case reflect.Map:
			AssertParameters(t, expectedValue.(map[string]any), actualValue.(map[string]any))
		}
	}
}
