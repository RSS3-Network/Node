package config

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/protocol-go/schema/filter"
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
    maintainer:
      evm_address: 0x111222333444555666777888999aaabbbcccddde
      signature: 0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee
    server:
      endpoint: https://node.mydomain.com/
      global_indexer_endpoint: https://gi.rss3.dev/
database:
  driver: cockroachdb
  partition: true
  uri: postgres://root@localhost:26257/defaultdb
stream:
  enable: false
  driver: kafka
  topic: rss3.node.feeds
  uri: localhost:9092
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
    - network: rss
      endpoint: https://rsshub.app/
      parameters:
        authentication:
          username: user
          password: pass
          access_key: abc
          access_code: def
  decentralized:
    - network: ethereum
      worker: fallback
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start: 47370106
        block_number_target: 456
    - network: ethereum
      worker: rss3
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_height_start: 123
        rpc_thread_blocks: 2
`
	configExampleJSON = `{
  "environment": "development",
  "discovery": {
    "maintainer": {
      "evm_address": "0x111222333444555666777888999aaabbbcccddde",
      "signature": "0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee"
    },
    "server": {
      "endpoint": "https://node.mydomain.com/",
      "global_indexer_endpoint": "https://gi.rss3.dev/"
    }
  },
  "database": {
    "driver": "cockroachdb",
    "partition": true,
    "uri": "postgres://root@localhost:26257/defaultdb"
  },
  "stream": {
    "enable": false,
    "driver": "kafka",
    "topic": "rss3.node.feeds",
    "uri": "localhost:9092"
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
    "rss": [
      {
        "network": "rss",
        "endpoint": "https://rsshub.app/",
        "parameters": {
          "authentication": {
            "username": "user",
            "password": "pass",
            "access_key": "abc",
            "access_code": "def"
          }
        }
      }
    ],
    "decentralized": [
      {
        "network": "ethereum",
        "worker": "fallback",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": 47370106,
          "block_number_target": 456
        }
      },
      {
        "network": "ethereum",
        "worker": "rss3",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_height_start": 123,
          "rpc_thread_blocks": 2
        }
      }
    ]
  }
}`
	configExampleToml = `environment = "development"

[discovery.maintainer]
evm_address = "0x111222333444555666777888999aaabbbcccddde"
signature = "0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee"

[discovery.server]
endpoint = "https://node.mydomain.com/"
global_indexer_endpoint = "https://gi.rss3.dev/"

[database]
driver = "cockroachdb"
partition = true
uri = "postgres://root@localhost:26257/defaultdb"

[stream]
enable = false
driver = "kafka"
topic = "rss3.node.feeds"
uri = "localhost:9092"

[observability.opentelemetry.metrics]
enable = true
endpoint = "0.0.0.0:9090"

[observability.opentelemetry.traces]
enable = true
insecure = true
endpoint = "localhost:4318"

[[component.rss]]
network = "rss"
endpoint = "https://rsshub.app/"

[component.rss.parameters.authentication]
username = "user"
password = "pass"
access_key = "abc"
access_code = "def"

[[component.decentralized]]
network = "ethereum"
worker = "fallback"
endpoint = "https://rpc.ankr.com/eth"

  [component.decentralized.parameters]
  block_number_start = 47370106
  block_number_target = 456

[[component.decentralized]]
network = "ethereum"
worker = "rss3"
endpoint = "https://rpc.ankr.com/eth"

  [component.decentralized.parameters]
  block_height_start = 123
  rpc_thread_blocks = 2

`
)

var configFileExcept = &File{
	Environment: "development",
	Discovery: &Discovery{
		Maintainer: &Maintainer{
			EvmAddress: common.HexToAddress("0x111222333444555666777888999aaabbbcccddde"),
			Signature:  "0x000000000111111111222222222333333333444444444555555555666666666777777777888888888999999999aaaaaaaaabbbbbbbbbcccccccccdddddddddeeee",
		},
		Server: &Server{
			Endpoint:              "https://node.mydomain.com/",
			GlobalIndexerEndpoint: "https://gi.rss3.dev/",
		},
	},
	Node: &Node{
		RSS: []*Module{
			{
				Network:  filter.NetworkRSS,
				Endpoint: "https://rsshub.app/",
				Worker:   0,
				Parameters: &Options{
					"authentication": map[string]any{
						"access_code": "def",
						"access_key":  "abc",
						"password":    "pass",
						"username":    "user",
					},
				},
			},
		},
		Federated: nil,
		Decentralized: []*Module{
			{
				Network:  filter.NetworkEthereum,
				Worker:   filter.Fallback,
				Endpoint: "https://rpc.ankr.com/eth",
				Parameters: &Options{
					"block_number_start":  47370106,
					"block_number_target": 456,
				},
			},
			{
				Network:  filter.NetworkEthereum,
				Worker:   filter.RSS3,
				Endpoint: "https://rpc.ankr.com/eth",
				Parameters: &Options{
					"block_height_start": 123,
					"rpc_thread_blocks":  2,
				},
			},
		},
	},
	Database: &Database{
		Driver:    "cockroachdb",
		Partition: lo.ToPtr(true),
		URI:       "postgres://root@localhost:26257/defaultdb",
	},
	Stream: &Stream{
		Enable: lo.ToPtr(false),
		Driver: "kafka",
		Topic:  "rss3.node.feeds",
		URI:    "localhost:9092",
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

	AssertConfig(t, f, configFileExcept)
}

//func TestConfigEnvOverride(t *testing.T) {
//	t.Parallel()
//
//	exceptEnvironment := "testing"
//	exceptDatabaseURI := "postgres://mock@localhost:26257/defaultdb"
//	exceptMetricsEndpoint := "127.0.0.1:9000"
//
//	t.Setenv("NODE_ENVIRONMENT", exceptEnvironment)
//	t.Setenv("NODE_DATABASE_URI", exceptDatabaseURI)
//	t.Setenv("NODE_OBSERVABILITY_OPENTELEMETRY_METRICS_ENDPOINT", exceptMetricsEndpoint)
//
//	configDir := "/etc/rss3/node"
//	fs := afero.NewMemMapFs()
//
//	err := fs.Mkdir(configDir, 0o777)
//	assert.NoError(t, err)
//
//	file, err := fs.Create(path.Join(configDir, configName))
//	assert.NoError(t, err)
//
//	_, err = file.WriteString(configExampleYaml)
//	require.NoError(t, err)
//
//	v := viper.New()
//	v.SetFs(fs)
//
//	f, err := _Setup(configName, "yaml", v)
//	assert.NoError(t, err)
//
//	assert.Equal(t, exceptEnvironment, f.Environment)
//	assert.Equal(t, exceptDatabaseURI, f.Database.URI)
//	assert.Equal(t, exceptMetricsEndpoint, f.Observability.OpenTelemetry.Metrics.Endpoint)
//}

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

				assert.Equal(t, configFileExcept, f)
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

				AssertConfig(t, configFileExcept, f)
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
		for i, rss := range expect.Node.RSS {
			func(_except, got *Module) {
				t.Run(fmt.Sprintf("rss-%d", i), func(t *testing.T) {
					t.Parallel()
					assert.Equal(t, _except, got)
				})
			}(rss, got.Node.RSS[i])
		}

		for i, indexer := range got.Node.Decentralized {
			func(_except, got *Module) {
				t.Run(fmt.Sprintf("%s-%s", indexer.Network, indexer.Worker), func(t *testing.T) {
					t.Parallel()
					AssertIndexer(t, _except, got)
				})
			}(configFileExcept.Node.Decentralized[i], indexer)
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
