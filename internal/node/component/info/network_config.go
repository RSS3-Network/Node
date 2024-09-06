package info

// This file includes the relations between networks and their associated workers.
// Results are exclusively provided to external users and do not influence internal operations within the Node.
// Each worker has a default configuration, which can be customized based on various factors.

import (
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
)

type ConfigDetailValueType string

const (
	BooleanType   ConfigDetailValueType = "bool"
	StringType    ConfigDetailValueType = "string"
	StringMapType ConfigDetailValueType = "map[string]string"
	UintType      ConfigDetailValueType = "uint"
	URLType       ConfigDetailValueType = "url"
	URLArrayType  ConfigDetailValueType = "[]url"
)

type ConfigDetail struct {
	IsRequired  bool                  `json:"is_required"`
	Type        ConfigDetailValueType `json:"type"`
	Value       interface{}           `json:"value"`
	Description string                `json:"description"`
}

type Endpoint struct {
	URL           *ConfigDetail `json:"url"`
	HTTPHeaders   *ConfigDetail `json:"http_headers"`
	HTTP2Disabled *ConfigDetail `json:"http2_disabled"`
}

type Authentication struct {
	AccessKey *ConfigDetail `json:"access_key"`
}

type Parameters struct {
	BlockStart              *ConfigDetail   `json:"block_start,omitempty"`
	BlockTarget             *ConfigDetail   `json:"block_target,omitempty"`
	ConcurrentBlockRequests *ConfigDetail   `json:"concurrent_block_requests,omitempty"`
	BlockBatchSize          *ConfigDetail   `json:"block_batch_size,omitempty"`
	ReceiptsBatchSize       *ConfigDetail   `json:"receipts_batch_size,omitempty"`
	BlockReceiptBatchSize   *ConfigDetail   `json:"block_receipts_batch_size,omitempty"`
	APIKey                  *ConfigDetail   `json:"api_key,omitempty"`
	Authentication          *Authentication `json:"authentication,omitempty"`
}

type workerConfig struct {
	ID              ConfigDetail    `json:"id"`
	Network         ConfigDetail    `json:"network"`
	Worker          ConfigDetail    `json:"worker"`
	EndpointID      *ConfigDetail   `json:"endpoint,omitempty"`
	IPFSGateways    *ConfigDetail   `json:"ipfs_gateways,omitempty"`
	Parameters      *Parameters     `json:"parameters,omitempty"`
	MinimumResource MinimumResource `json:"minimum_resource"`
}

var defaultNetworkParameters = map[network.Source]*Parameters{
	network.EthereumSource: {
		// unnecessary to expose
		//BlockStart: &ConfigDetail{
		//	IsRequired:  false,
		//	Type:        BigIntType,
		//	Description: "The block number where your worker will begin indexing. Each version of Node has a different starting block number.",
		// },
		// BlockTarget: &ConfigDetail{
		//	IsRequired:  false,
		//	Type:        BigIntType,
		//	Description: "The block number where your worker will stop indexing",
		// },
		ConcurrentBlockRequests: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(8),
			Description: "The number of concurrent RPC requests to the blockchain rpc. Default: 8",
		},
		BlockBatchSize: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(8),
			Description: "The number of blocks to fetch in a single RPC request. Default: 8",
		},
		ReceiptsBatchSize: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(200),
			Description: "The number of receipts to fetch in a single RPC request. Default: 200",
		},
		BlockReceiptBatchSize: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(8),
			Description: "The number of blocks to fetch receipts in a single RPC request. Default: 8",
		},
	},
	network.ArweaveSource: {
		// unnecessary to expose
		//BlockStart: &ConfigDetail{
		//	IsRequired:  false,
		//	Type:        BigIntType,
		//	Description: "The block height where your worker will begin indexing. Each version of Node has a different starting block height",
		// },
		// BlockTarget: &ConfigDetail{
		//	IsRequired:  false,
		//	Type:        BigIntType,
		//	Description: "The block height where your worker will stop indexing",
		// },
		ConcurrentBlockRequests: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(1),
			Description: "The number of concurrent RPC requests to the Arweave gateway. Default: 1",
		},
	},
}

// getDefaultParametersByNetwork returns the default parameters based on the network.
func getDefaultParametersByNetwork(network network.Source) *Parameters {
	return defaultNetworkParameters[network]
}

// defaultWorkerConfig returns the default worker config based on the worker and network.
// If parameters are supplied, use them instead of the default parameters.
func defaultWorkerConfig(worker worker.Worker, network network.Source, parameters *Parameters) workerConfig {
	// generate default parameters only if parameters are not provided
	if parameters == nil {
		parameters = getDefaultParametersByNetwork(network)
	}

	config := workerConfig{
		ID: ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Description: "Worker's id, must be unique, for example '[network]-[worker]'",
		},
		Network: ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Description: "The network where the worker operates on",
		},
		Worker: ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Value:       worker.Name(),
			Description: "Name of the worker",
		},
		EndpointID: &ConfigDetail{
			IsRequired:  true,
			Type:        URLType,
			Description: "An external endpoint to fetch data from, for example, a blockchain RPC endpoint or a Farcaster api",
		},
		Parameters: parameters,
	}

	return config
}

// customWorkerConfigWithoutEndpoint generates a config with custom fields and no endpoint.
func customWorkerConfigWithoutEndpoint(worker worker.Worker, network network.Source, parameters *Parameters, requireIPFS bool) workerConfig {
	config := defaultWorkerConfig(worker, network, parameters)

	config.EndpointID = nil

	if requireIPFS {
		setIPFSGateways(&config)
	}

	return config
}

// customWorkerConfigWithIPFS generates a config with IPFS and custom fields.
func customWorkerConfigWithIPFS(worker worker.Worker, network network.Source, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, network, nil)

	setIPFSGateways(&config)

	if len(endpointDescription) > 0 {
		config.EndpointID.Description = endpointDescription
	}

	return config
}

// customWorkerConfig generates a config with custom fields.
func customWorkerConfig(worker worker.Worker, network network.Source, parameters *Parameters, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, network, parameters)

	// Update the EndpointID description based on the provided custom description
	config.EndpointID.Description = endpointDescription

	return config
}

func getEndpointConfig() Endpoint {
	return Endpoint{
		URL: &ConfigDetail{
			IsRequired:  true,
			Type:        URLType,
			Description: "The URL of the endpoint.",
		},
		HTTPHeaders: &ConfigDetail{
			IsRequired:  false,
			Type:        StringMapType,
			Description: "HTTP headers to be sent with requests to this endpoint.",
		},
		HTTP2Disabled: &ConfigDetail{
			IsRequired:  false,
			Type:        BooleanType,
			Description: "Some endpoints may not support HTTP2, set this to true to disable HTTP2.",
		},
	}
}

func setIPFSGateways(config *workerConfig) {
	config.IPFSGateways = &ConfigDetail{
		IsRequired:  true,
		Type:        URLArrayType,
		Description: "A list of IPFS gateways to fetch data from, multiple gateways may improve the availability of IPFS data",
	}
}

// NetworkToWorkersMap is a map of Network-has-Workers.
var NetworkToWorkersMap = map[network.Network][]worker.Worker{
	network.Ethereum: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.ENS,
		decentralized.Highlight,
		decentralized.Lido,
		decentralized.Looksrare,
		decentralized.Oneinch,
		decentralized.OpenSea,
		decentralized.Optimism,
		decentralized.RSS3,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.VSL,
	},
	network.Arweave: {
		decentralized.Mirror,
		decentralized.Momoka,
		decentralized.Paragraph,
	},
	network.Farcaster: {
		decentralized.Core,
	},
	network.Polygon: {
		decentralized.Aave,
		decentralized.Aavegotchi,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.IQWiki,
		decentralized.Lens,
		decentralized.Stargate,
	},
	network.Crossbell: {
		decentralized.Core,
		decentralized.Crossbell,
	},
	network.Avalanche: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Stargate,
	},
	network.Base: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Stargate,
	},
	network.Optimism: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.KiwiStand,
		decentralized.Matters,
		decentralized.Optimism,
		decentralized.Stargate,
	},
	network.Arbitrum: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.Stargate,
	},
	network.VSL: {
		decentralized.Core,
	},
	network.SatoshiVM: {
		decentralized.Core,
		decentralized.SAVM,
		decentralized.Uniswap,
	},
	network.BinanceSmartChain: {
		decentralized.Core,
		decentralized.Stargate,
	},
	network.Gnosis: {
		decentralized.Core,
		decentralized.Curve,
	},
	network.Linea: {
		decentralized.Core,
		decentralized.Stargate,
		decentralized.Uniswap,
	},
	network.RSS: {
		rss.RSSHub,
	},
	network.XLayer: {
		decentralized.Core,
	},
}

// WorkerToConfigMap is a map of worker to config.
var WorkerToConfigMap = map[network.Source]map[worker.Worker]workerConfig{
	network.EthereumSource: {
		decentralized.Aave:       defaultWorkerConfig(decentralized.Aave, network.EthereumSource, nil),
		decentralized.Aavegotchi: defaultWorkerConfig(decentralized.Aavegotchi, network.EthereumSource, nil),
		decentralized.Core:       defaultWorkerConfig(decentralized.Core, network.EthereumSource, nil),
		decentralized.Crossbell:  customWorkerConfigWithIPFS(decentralized.Crossbell, network.EthereumSource, ""),
		decentralized.Curve:      defaultWorkerConfig(decentralized.Curve, network.EthereumSource, nil),
		decentralized.ENS:        defaultWorkerConfig(decentralized.ENS, network.EthereumSource, nil),
		decentralized.Highlight:  defaultWorkerConfig(decentralized.Highlight, network.EthereumSource, nil),
		decentralized.IQWiki:     customWorkerConfigWithIPFS(decentralized.IQWiki, network.EthereumSource, ""),
		decentralized.KiwiStand:  defaultWorkerConfig(decentralized.KiwiStand, network.EthereumSource, nil),
		decentralized.Lens:       customWorkerConfigWithIPFS(decentralized.Lens, network.EthereumSource, ""),
		decentralized.Lido:       defaultWorkerConfig(decentralized.Lido, network.EthereumSource, nil),
		decentralized.Looksrare:  defaultWorkerConfig(decentralized.Looksrare, network.EthereumSource, nil),
		decentralized.Matters:    customWorkerConfigWithIPFS(decentralized.Matters, network.EthereumSource, ""),
		decentralized.Oneinch:    defaultWorkerConfig(decentralized.Oneinch, network.EthereumSource, nil),
		decentralized.OpenSea:    defaultWorkerConfig(decentralized.OpenSea, network.EthereumSource, nil),
		decentralized.Optimism:   defaultWorkerConfig(decentralized.Optimism, network.EthereumSource, nil),
		decentralized.RSS3:       defaultWorkerConfig(decentralized.RSS3, network.EthereumSource, nil),
		decentralized.SAVM:       defaultWorkerConfig(decentralized.SAVM, network.EthereumSource, nil),
		decentralized.Stargate:   defaultWorkerConfig(decentralized.Stargate, network.EthereumSource, nil),
		decentralized.Uniswap:    defaultWorkerConfig(decentralized.Uniswap, network.EthereumSource, nil),
		decentralized.VSL:        defaultWorkerConfig(decentralized.VSL, network.EthereumSource, nil),
	},
	network.ArweaveSource: {
		decentralized.Mirror:    customWorkerConfigWithoutEndpoint(decentralized.Mirror, network.ArweaveSource, nil, true),
		decentralized.Momoka:    customWorkerConfigWithIPFS(decentralized.Momoka, network.ArweaveSource, "A Polygon RPC is required for Momoka"),
		decentralized.Paragraph: customWorkerConfigWithoutEndpoint(decentralized.Paragraph, network.ArweaveSource, nil, false),
	},
	network.FarcasterSource: {
		decentralized.Core: customWorkerConfig(decentralized.Core, network.FarcasterSource, &Parameters{
			APIKey: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Description: "API key to access your Farcaster Hubble on Neynar",
			},
		}, "A Farcaster Hubble is required"),
	},
	network.RSSSource: {
		rss.RSSHub: customWorkerConfig(rss.RSSHub, network.RSSSource, &Parameters{
			Authentication: &Authentication{
				AccessKey: &ConfigDetail{
					IsRequired:  false,
					Type:        StringType,
					Description: "A key to access the RSS Feed",
				},
			},
		}, ""),
	},
}
