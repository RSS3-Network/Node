package network

// This file includes the relations between networks and their associated workers.
// Results are exclusively provided to external users and do not influence internal operations within the Node.
// Each worker has a default configuration, which can be customized based on various factors.

import (
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
)

type ConfigDetailValueType string

const (
	BigIntType      ConfigDetailValueType = "big.Int"
	StringArrayType ConfigDetailValueType = "[]string"
	StringType      ConfigDetailValueType = "string"
	UintType        ConfigDetailValueType = "uint"
)

type ConfigDetail struct {
	IsRequired  bool                  `json:"isRequired"`
	Type        ConfigDetailValueType `json:"type"`
	Value       interface{}           `json:"value"`
	Description string                `json:"description"`
}

type Endpoint struct {
	URL           *ConfigDetail `json:"url"`
	HTTPHeaders   *ConfigDetail `json:"http_headers"`
	HTTP2Disabled *ConfigDetail `json:"http2_disabled"`
}

type Parameters struct {
	BlockStart              *ConfigDetail `json:"block_start,omitempty"`
	BlockTarget             *ConfigDetail `json:"block_target,omitempty"`
	ConcurrentBlockRequests *ConfigDetail `json:"concurrent_block_requests,omitempty"`
	BlockBatchSize          *ConfigDetail `json:"block_batch_size,omitempty"`
	ReceiptsBatchSize       *ConfigDetail `json:"receipts_batch_size,omitempty"`
	BlockReceiptBatchSize   *ConfigDetail `json:"block_receipts_batch_size,omitempty"`
	APIKey                  *ConfigDetail `json:"api_key,omitempty"`
}

type workerConfig struct {
	ID           ConfigDetail  `json:"id"`
	Network      ConfigDetail  `json:"network"`
	Worker       ConfigDetail  `json:"worker"`
	EndpointID   *ConfigDetail `json:"endpoint,omitempty"`
	IPFSGateways *ConfigDetail `json:"ipfs_gateways,omitempty"`
	Parameters   *Parameters   `json:"parameters,omitempty"`
}

var defaultNetworkParameters = map[network.Source]*Parameters{
	network.EthereumSource: {
		BlockStart: &ConfigDetail{
			IsRequired:  false,
			Type:        BigIntType,
			Description: "The block number where your worker will begin indexing. Each version of Node has a different starting block number.",
		},
		// unnecessary to expose
		// BlockNumberTarget: &ConfigDetail{
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
		BlockStart: &ConfigDetail{
			IsRequired:  false,
			Type:        BigIntType,
			Description: "The block height where your worker will begin indexing. Each version of Node has a different starting block height",
		},
		// unnecessary to expose
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
			Value:       worker.String(),
			Description: "Name of the worker",
		},
		EndpointID: &ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Description: "An external endpoint to fetch data from, for example, a blockchain RPC endpoint or an Arweave gateway",
		},
		Parameters: parameters,
	}

	return config
}

// customWorkerConfigWithIPFS generates a config with IPFS and custom fields.
func customWorkerConfigWithIPFS(worker worker.Worker, network network.Source, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, network, nil)

	config.IPFSGateways = &ConfigDetail{
		IsRequired:  false,
		Type:        StringArrayType,
		Description: "A list of IPFS gateways to fetch data from, multiple gateways may improve the availability of IPFS data",
	}

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

// NetworkToWorkersMap is a map of Network-has-Workers.
var NetworkToWorkersMap = map[network.Network][]worker.Worker{
	network.Ethereum: {
		worker.Aave,
		worker.Core,
		worker.Curve,
		worker.ENS,
		worker.Highlight,
		worker.Lido,
		worker.Looksrare,
		worker.Oneinch,
		worker.OpenSea,
		worker.Optimism,
		worker.RSS3,
		worker.Stargate,
		worker.Uniswap,
	},
	network.Arweave: {
		worker.Mirror,
		worker.Momoka,
		worker.Paragraph,
	},
	network.Farcaster: {
		worker.Core,
	},
	network.Polygon: {
		worker.Aave,
		worker.Aavegotchi,
		worker.Core,
		worker.Curve,
		worker.Highlight,
		worker.IQWiki,
		worker.Lens,
		worker.Stargate,
	},
	network.Crossbell: {
		worker.Crossbell,
	},
	network.Avalanche: {
		worker.Aave,
		worker.Curve,
		worker.Stargate,
	},
	network.Base: {
		worker.Aave,
		worker.Core,
		worker.Stargate,
	},
	network.Optimism: {
		worker.Aave,
		worker.Core,
		worker.Curve,
		worker.Highlight,
		worker.KiwiStand,
		worker.Matters,
		worker.Stargate,
	},
	network.Arbitrum: {
		worker.Aave,
		worker.Core,
		worker.Curve,
		worker.Highlight,
		worker.Stargate,
	},
	network.VSL: {
		worker.Core,
		worker.RSS3,
		worker.VSL,
	},
	network.SatoshiVM: {
		worker.Core,
		worker.SAVM,
		worker.Uniswap,
	},
	network.BinanceSmartChain: {
		worker.Core,
		worker.Stargate,
	},
	network.Gnosis: {
		worker.Core,
		worker.Curve,
	},
	network.Linea: {
		worker.Core,
		worker.Stargate,
		worker.Uniswap,
	},
}

// WorkerToConfigMap is a map of worker to config.
var WorkerToConfigMap = map[network.Source]map[worker.Worker]workerConfig{
	network.EthereumSource: {
		worker.Aave:       defaultWorkerConfig(worker.Aave, network.EthereumSource, nil),
		worker.Aavegotchi: defaultWorkerConfig(worker.Aavegotchi, network.EthereumSource, nil),
		worker.Core:       defaultWorkerConfig(worker.Core, network.EthereumSource, nil),
		worker.Crossbell:  customWorkerConfigWithIPFS(worker.Crossbell, network.EthereumSource, ""),
		worker.Curve:      defaultWorkerConfig(worker.Curve, network.EthereumSource, nil),
		worker.ENS:        defaultWorkerConfig(worker.ENS, network.EthereumSource, nil),
		worker.Highlight:  defaultWorkerConfig(worker.Highlight, network.EthereumSource, nil),
		worker.IQWiki:     customWorkerConfigWithIPFS(worker.IQWiki, network.EthereumSource, ""),
		worker.KiwiStand:  defaultWorkerConfig(worker.KiwiStand, network.EthereumSource, nil),
		worker.Lens:       customWorkerConfigWithIPFS(worker.Lens, network.EthereumSource, ""),
		worker.Lido:       defaultWorkerConfig(worker.Lido, network.EthereumSource, nil),
		worker.Looksrare:  defaultWorkerConfig(worker.Looksrare, network.EthereumSource, nil),
		worker.Matters:    customWorkerConfigWithIPFS(worker.Matters, network.EthereumSource, ""),
		worker.OpenSea:    defaultWorkerConfig(worker.OpenSea, network.EthereumSource, nil),
		worker.Optimism:   defaultWorkerConfig(worker.Optimism, network.EthereumSource, nil),
		worker.RSS3:       defaultWorkerConfig(worker.RSS3, network.EthereumSource, nil),
		worker.SAVM:       defaultWorkerConfig(worker.SAVM, network.EthereumSource, nil),
		worker.Stargate:   defaultWorkerConfig(worker.Stargate, network.EthereumSource, nil),
		worker.Uniswap:    defaultWorkerConfig(worker.Uniswap, network.EthereumSource, nil),
		worker.VSL:        defaultWorkerConfig(worker.VSL, network.EthereumSource, nil),
	},
	network.ArweaveSource: {
		worker.Mirror:    customWorkerConfigWithIPFS(worker.Mirror, network.ArweaveSource, ""),
		worker.Momoka:    customWorkerConfigWithIPFS(worker.Momoka, network.ArweaveSource, "A Polygon RPC is required for Momoka"),
		worker.Paragraph: defaultWorkerConfig(worker.Paragraph, network.ArweaveSource, nil),
	},
	network.FarcasterSource: {
		worker.Core: customWorkerConfig(worker.Core, network.FarcasterSource, &Parameters{
			APIKey: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Description: "API key to access your Farcaster Hubble on Neynar",
			},
		}, "A Farcaster Hubble is required"),
	},
}
