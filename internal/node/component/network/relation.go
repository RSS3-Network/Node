package network

import (
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
)

type ConfigDetail struct {
	IsRequired  bool        `json:"isRequired"`
	Type        string      `json:"type"`
	Value       interface{} `json:"value"`
	Description string      `json:"description"`
}

type Endpoint struct {
	URL           *ConfigDetail `json:"url"`
	HTTPHeaders   *ConfigDetail `json:"http_headers"`
	HTTP2Disabled *ConfigDetail `json:"http2_disabled"`
}

type Parameters struct {
	BlockHeightStart      *ConfigDetail `json:"block_height_start,omitempty"`
	BlockHeightTarget     *ConfigDetail `json:"block_height_target,omitempty"`
	BlockNumberStart      *ConfigDetail `json:"block_number_start,omitempty"`
	BlockNumberTarget     *ConfigDetail `json:"block_number_target,omitempty"`
	RPCThreadBlocks       *ConfigDetail `json:"rpc_thread_blocks,omitempty"`
	RPCBatchBlocks        *ConfigDetail `json:"rpc_batch_blocks,omitempty"`
	RPCBatchReceipts      *ConfigDetail `json:"rpc_batch_receipts,omitempty"`
	RPCBatchBlockReceipts *ConfigDetail `json:"rpc_batch_block_receipts,omitempty"`
	APIKey                *ConfigDetail `json:"api_key,omitempty"`
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
		BlockNumberStart: &ConfigDetail{
			IsRequired:  false,
			Type:        "big.Int",
			Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
		BlockNumberTarget: &ConfigDetail{
			IsRequired:  false,
			Type:        "big.Int",
			Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing new blocks",
		},
		RPCThreadBlocks: &ConfigDetail{
			IsRequired:  false,
			Type:        "uint",
			Value:       uint(8),
			Description: "The number of concurrent RPC requests to the blockchain rpc, the default value is 8, you can adjust it based on your server performance",
		},
		RPCBatchBlocks: &ConfigDetail{
			IsRequired:  false,
			Type:        "uint",
			Value:       uint(8),
			Description: "The number of blocks to fetch in a single RPC request, the default value is 8, you can adjust it based on your server performance",
		},
		RPCBatchReceipts: &ConfigDetail{
			IsRequired:  false,
			Type:        "uint",
			Value:       uint(200),
			Description: "The number of receipts to fetch in a single RPC request, the default value is 200, you can adjust it based on your server performance",
		},
		RPCBatchBlockReceipts: &ConfigDetail{
			IsRequired:  false,
			Type:        "uint",
			Value:       uint(8),
			Description: "The number of blocks to fetch receipts in a single RPC request, the default value is 8, you can adjust it based on your server performance",
		},
	},
	network.ArweaveSource: {
		BlockHeightStart: &ConfigDetail{
			IsRequired:  false,
			Type:        "big.Int",
			Description: "Your arweave worker will start from this block height, if it's not defined, we will use the recommended one(~3 months)",
		},
		BlockHeightTarget: &ConfigDetail{
			IsRequired:  false,
			Type:        "big.Int",
			Description: "Your arweave worker will stop at this block height, if it's not defined, your worker will always indexing",
		},
		RPCThreadBlocks: &ConfigDetail{
			IsRequired:  false,
			Type:        "uint",
			Value:       uint(1),
			Description: "The number of concurrent RPC requests to the arweave endpoints or gateways, it's recommended to set it to 1 because the strict rate limit of arweave",
		},
	},
}

func getDefaultParameters(network network.Source) *Parameters {
	return defaultNetworkParameters[network]
}

func defaultWorkerConfig(worker worker.Worker, network network.Source, parameters *Parameters) workerConfig {
	// generate default parameters only if parameters are not provided
	if parameters == nil {
		parameters = getDefaultParameters(network)
	}

	config := workerConfig{
		ID: ConfigDetail{
			IsRequired:  true,
			Type:        "string",
			Description: "Worker's id, must be unique, for example '[network]-[worker]'",
		},
		Network: ConfigDetail{
			IsRequired:  true,
			Type:        "string",
			Description: "The network where the worker operates on",
		},
		Worker: ConfigDetail{
			IsRequired:  true,
			Type:        "string",
			Value:       worker.String(),
			Description: "Name of the worker",
		},
		EndpointID: &ConfigDetail{
			IsRequired:  true,
			Type:        "string",
			Description: "An external endpoint to fetch data from, for example, a blockchain RPC endpoint",
		},
		Parameters: parameters,
	}

	return config
}

// workerConfigWithIPFS returns the default worker config with IPFS.
// If parameters are supplied, use them to override the default parameters.
func workerConfigWithIPFS(worker worker.Worker, network network.Source, parameters *Parameters) workerConfig {
	config := defaultWorkerConfig(worker, network, parameters)

	config.IPFSGateways = &ConfigDetail{
		IsRequired:  false,
		Type:        "[]string",
		Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service",
	}

	return config
}

// customWorkerConfigWithIPFS generates a config with IPFS and custom fields.
func customWorkerConfigWithIPFS(worker worker.Worker, network network.Source, parameters *Parameters, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, network, parameters)

	config.IPFSGateways = &ConfigDetail{
		IsRequired:  false,
		Type:        "[]string",
		Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service",
	}

	config.EndpointID.Description = endpointDescription

	return config
}

// customWorkerConfig generates a config with custom fields.
func customWorkerConfig(worker worker.Worker, network network.Source, parameters *Parameters, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, network, parameters)

	// Update the EndpointID description based on the provided custom description
	config.EndpointID.Description = endpointDescription

	return config
}

// NetworkToWorkersMap is a map of network to workers.
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
		worker.Matters,
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
		worker.Aave:       defaultWorkerConfig(worker.Aave, network.EthereumSource, &Parameters{}),
		worker.Aavegotchi: defaultWorkerConfig(worker.Aavegotchi, network.EthereumSource, &Parameters{}),
		worker.Core:       defaultWorkerConfig(worker.Core, network.EthereumSource, &Parameters{}),
		worker.Crossbell:  workerConfigWithIPFS(worker.Crossbell, network.EthereumSource, &Parameters{}),
		worker.Curve:      defaultWorkerConfig(worker.Curve, network.EthereumSource, &Parameters{}),
		worker.ENS:        defaultWorkerConfig(worker.ENS, network.EthereumSource, &Parameters{}),
		worker.Highlight:  defaultWorkerConfig(worker.Highlight, network.EthereumSource, &Parameters{}),
		worker.IQWiki:     workerConfigWithIPFS(worker.IQWiki, network.EthereumSource, &Parameters{}),
		worker.KiwiStand:  defaultWorkerConfig(worker.KiwiStand, network.EthereumSource, &Parameters{}),
		worker.Lens:       workerConfigWithIPFS(worker.Lens, network.EthereumSource, &Parameters{}),
		worker.Lido:       defaultWorkerConfig(worker.Lido, network.EthereumSource, &Parameters{}),
		worker.Looksrare:  defaultWorkerConfig(worker.Looksrare, network.EthereumSource, &Parameters{}),
		worker.Matters:    workerConfigWithIPFS(worker.Matters, network.EthereumSource, &Parameters{}),
		worker.OpenSea:    defaultWorkerConfig(worker.OpenSea, network.EthereumSource, &Parameters{}),
		worker.Optimism:   defaultWorkerConfig(worker.Optimism, network.EthereumSource, &Parameters{}),
		worker.RSS3:       defaultWorkerConfig(worker.RSS3, network.EthereumSource, &Parameters{}),
		worker.SAVM:       defaultWorkerConfig(worker.SAVM, network.EthereumSource, &Parameters{}),
		worker.Stargate:   defaultWorkerConfig(worker.Stargate, network.EthereumSource, &Parameters{}),
		worker.Uniswap:    defaultWorkerConfig(worker.Uniswap, network.EthereumSource, &Parameters{}),
		worker.VSL:        defaultWorkerConfig(worker.VSL, network.EthereumSource, &Parameters{}),
	},
	network.ArweaveSource: {
		worker.Mirror:    workerConfigWithIPFS(worker.Mirror, network.ArweaveSource, &Parameters{}),
		worker.Momoka:    customWorkerConfigWithIPFS(worker.Mirror, network.ArweaveSource, &Parameters{}, "A Polygon RPC is required for Momoka"),
		worker.Paragraph: defaultWorkerConfig(worker.Mirror, network.ArweaveSource, &Parameters{}),
	},
	network.FarcasterSource: {
		worker.Core: customWorkerConfig(worker.Core, network.FarcasterSource, &Parameters{
			APIKey: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "API key to access your Farcaster Hubble",
			},
		}, "If your Farcaster Hubble requires authentication"),
	},
}
