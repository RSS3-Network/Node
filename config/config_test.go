package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/serving-node/schema/filter"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

const (
	configName        = "config.development.yaml"
	configExampleYaml = `environment: development
discovery:
    maintainer:
      evm_address:
      signature:
    server:
      endpoint:
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
          username:
          password:
          access_key:
          access_code:
  decentralized:
    # ethereum
    - network: ethereum
      worker: fallback
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
        block_number_target:
    - network: ethereum
      worker: rss3
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
        block_number_target:
    - network: polygon
      worker: lens
      endpoint: https://rpc.ankr.com/polygon
      ipfs_gateways:
      parameters:
        block_number_start:
    - network: ethereum
      worker: opensea
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
    - network: ethereum
      worker: uniswap
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
    - network: ethereum
      worker: optimism
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
    - network: polygon
      worker: aavegotchi
      endpoint: https://rpc.ankr.com/polygon
      parameters:
          block_number_start: 48300005
    - network: ethereum
      worker: highlight
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
    # farcaster
    - network: farcaster
      worker: farcaster
      endpoint: https://nemes.farcaster.xyz:2281
    # arweave
    - network: arweave
      worker: mirror
      endpoint: https://arweave.net/
      ipfs_gateways:
      parameters:
        block_height_start:
        rpc_thread_blocks:
    - network: arweave
      worker: paragraph
      endpoint: https://arweave.net
      ipfs_gateways:
      parameters:
        block_height_start:
        rpc_thread_blocks:
    - network: ethereum
      worker: looksrare
      endpoint: https://rpc.ankr.com/eth
      parameters:
        block_number_start:
    - network: polygon
      worker: matters
      endpoint: https://rpc.ankr.com/polygon
      parameters:
        block_number_start: 52398940
    - network: arweave
      worker: momoka
      endpoint: https://rpc.ankr.com/polygon
      parameters:
        block_height_start:
        rpc_thread_blocks:
`
	configExampleJson = `{
  "environment": "development",
  "discovery": {
    "maintainer": {
      "evm_address": null,
      "signature": null
    },
    "server": {
      "endpoint": null,
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
            "username": null,
            "password": null,
            "access_key": null,
            "access_code": null
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
          "block_number_start": null,
          "block_number_target": null
        }
      },
      {
        "network": "ethereum",
        "worker": "rss3",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": null,
          "block_number_target": null
        }
      },
      {
        "network": "polygon",
        "worker": "lens",
        "endpoint": "https://rpc.ankr.com/polygon",
        "ipfs_gateways": null,
        "parameters": {
          "block_number_start": null
        }
      },
      {
        "network": "ethereum",
        "worker": "opensea",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": null
        }
      },
      {
        "network": "ethereum",
        "worker": "uniswap",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": null
        }
      },
      {
        "network": "ethereum",
        "worker": "optimism",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": null
        }
      },
      {
        "network": "polygon",
        "worker": "aavegotchi",
        "endpoint": "https://rpc.ankr.com/polygon",
        "parameters": {
          "block_number_start": 48300005
        }
      },
      {
        "network": "ethereum",
        "worker": "highlight",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": null
        }
      },
      {
        "network": "farcaster",
        "worker": "farcaster",
        "endpoint": "https://nemes.farcaster.xyz:2281"
      },
      {
        "network": "arweave",
        "worker": "mirror",
        "endpoint": "https://arweave.net/",
        "ipfs_gateways": null,
        "parameters": {
          "block_height_start": null,
          "rpc_thread_blocks": null
        }
      },
      {
        "network": "arweave",
        "worker": "paragraph",
        "endpoint": "https://arweave.net",
        "ipfs_gateways": null,
        "parameters": {
          "block_height_start": null,
          "rpc_thread_blocks": null
        }
      },
      {
        "network": "ethereum",
        "worker": "looksrare",
        "endpoint": "https://rpc.ankr.com/eth",
        "parameters": {
          "block_number_start": null
        }
      },
      {
        "network": "polygon",
        "worker": "matters",
        "endpoint": "https://rpc.ankr.com/polygon",
        "parameters": {
          "block_number_start": 52398940
        }
      },
      {
        "network": "arweave",
        "worker": "momoka",
        "endpoint": "https://rpc.ankr.com/polygon",
        "parameters": {
          "block_height_start": null,
          "rpc_thread_blocks": null
        }
      }
    ]
  }
}`
	configExampleToml = `environment = "development"

[discovery.maintainer]
evm_address = ""
signature = ""

[discovery.server]
endpoint = ""
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

[[component.decentralized]]
network = "ethereum"
worker = "fallback"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "ethereum"
worker = "rss3"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "polygon"
worker = "lens"
endpoint = "https://rpc.ankr.com/polygon"

[component.decentralized.parameters]

[[component.decentralized]]
network = "ethereum"
worker = "opensea"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "ethereum"
worker = "uniswap"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "ethereum"
worker = "optimism"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "polygon"
worker = "aavegotchi"
endpoint = "https://rpc.ankr.com/polygon"

[component.decentralized.parameters]
block_number_start = 48_300_005

[[component.decentralized]]
network = "ethereum"
worker = "highlight"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "farcaster"
worker = "farcaster"
endpoint = "https://nemes.farcaster.xyz:2281"

[[component.decentralized]]
network = "arweave"
worker = "mirror"
endpoint = "https://arweave.net/"

[component.decentralized.parameters]

[[component.decentralized]]
network = "arweave"
worker = "paragraph"
endpoint = "https://arweave.net"

[component.decentralized.parameters]

[[component.decentralized]]
network = "ethereum"
worker = "looksrare"
endpoint = "https://rpc.ankr.com/eth"

[component.decentralized.parameters]

[[component.decentralized]]
network = "polygon"
worker = "matters"
endpoint = "https://rpc.ankr.com/polygon"

[component.decentralized.parameters]
block_number_start = 52_398_940

[[component.decentralized]]
network = "arweave"
worker = "momoka"
endpoint = "https://rpc.ankr.com/polygon"

[component.decentralized.parameters]
block_height_start = ""
rpc_thread_blocks = ""
`
)

var configFileTest = &File{
	Environment: "development",
	Discovery: &Discovery{
		Maintainer: &Maintainer{
			EvmAddress: common.Address{},
			Signature:  "",
		},
		Server: &Server{
			Endpoint:              "",
			GlobalIndexerEndpoint: "https://gi.rss3.dev/",
		},
	},
	Node: &Node{
		RSS: []*Module{
			{
				Network:    filter.NetworkRSS,
				Endpoint:   "https://rsshub.app/",
				Worker:     0,
				Parameters: &Options{},
			},
		},
		Federated: nil,
		Decentralized: []*Module{
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.Fallback,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.RSS3,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkPolygon,
				Worker:     filter.Lens,
				Endpoint:   "https://rpc.ankr.com/polygon",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.OpenSea,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.Uniswap,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.Optimism,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkPolygon,
				Worker:     filter.Aavegotchi,
				Endpoint:   "https://rpc.ankr.com/polygon",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.Highlight,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:  filter.NetworkFarcaster,
				Worker:   filter.Farcaster,
				Endpoint: "https://nemes.farcaster.xyz:2281",
			},
			{
				Network:    filter.NetworkArweave,
				Worker:     filter.Mirror,
				Endpoint:   "https://arweave.net/",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkArweave,
				Worker:     filter.Paragraph,
				Endpoint:   "https://arweave.net",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkEthereum,
				Worker:     filter.Looksrare,
				Endpoint:   "https://rpc.ankr.com/eth",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkPolygon,
				Worker:     filter.Matters,
				Endpoint:   "https://rpc.ankr.com/polygon",
				Parameters: &Options{},
			},
			{
				Network:    filter.NetworkArweave,
				Worker:     filter.Momoka,
				Endpoint:   "https://rpc.ankr.com/polygon",
				Parameters: &Options{},
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
	configDir := "/etc/serving-node"
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

	AssertConfig(t, f, configFileTest)

}

//func TestConfigEnvOverride(t *testing.T) {
//	t.Setenv("SN_ENVIRONMENT", "testing")
//	t.Setenv("SN_DATABASE_URI", "postgres://mock@localhost:26257/defaultdb")
//
//	configDir := "/etc/serving-node"
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
//	t.Log(v.Get("database.uri"))
//	t.Log(v.Get("DATABASE_URI"))
//
//	assert.Equal(t, "testing", f.Environment)
//	assert.Equal(t, "postgres://mock@localhost:26257/defaultdb", f.Database.URI)
//}

func TestConfigFilePath(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.NoError(t, err)

	configPaths := []string{
		"/etc/serving-node/",
		path.Join(os.Getenv("HOME"), ".serving-node"),
		path.Join(currentDir, "config"),
		path.Join(currentDir, "deploy"),
	}

	for _, configPath := range configPaths {
		t.Run(configPath, func(t *testing.T) {
			t.Parallel()
			fs := afero.NewMemMapFs()

			err := fs.Mkdir(configPath, 0o777)
			assert.NoError(t, err)

			file, err := fs.Create(path.Join(configPath, configName))
			assert.NoError(t, err)

			_, err = file.WriteString(configExampleYaml)
			require.NoError(t, err)

			v := viper.New()
			v.SetFs(fs)

			f, err := _Setup(configName, "yaml", v)
			assert.NoError(t, err)

			assert.Equal(t, configFileTest, f)
		})
	}
}

func TestConfigFileType(t *testing.T) {
	configDir := "/etc/serving-node"

	configFiles := map[string]string{
		"yaml": configExampleYaml,
		"yml":  configExampleYaml,
		"json": configExampleJson,
		"toml": configExampleToml,
	}

	for configType, configFile := range configFiles {
		t.Run(configType, func(t *testing.T) {
			t.Parallel()
			fs := afero.NewMemMapFs()

			err := fs.Mkdir(configDir, 0o777)
			assert.NoError(t, err)

			configName := fmt.Sprintf("config.development.%s", configType)

			file, err := fs.Create(path.Join(configDir, configName))
			assert.NoError(t, err)

			_, err = file.WriteString(configFile)
			require.NoError(t, err)

			v := viper.New()
			v.SetFs(fs)

			f, err := _Setup(configName, configType, v)
			assert.NoError(t, err)

			assert.Equal(t, f, configFileTest)
		})
	}
}

func AssertConfig(t *testing.T, got, expect *File) {
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
		for index, rss := range expect.Node.RSS {
			t.Run(fmt.Sprintf("rss-%d", index), func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, rss, got.Node.RSS[index])
			})
		}

		for key, indexer := range got.Node.Decentralized {
			t.Run(fmt.Sprintf("%s-%s", indexer.Network, indexer.Worker), func(t *testing.T) {
				t.Parallel()
				AssertIndexer(t, configFileTest.Node.Decentralized[key], indexer)
			})
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
		assert.Equal(t, expect.Parameters, got.Parameters)
	})
}
