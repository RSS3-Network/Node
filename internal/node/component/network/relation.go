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
	WorkerID     ConfigDetail  `json:"worker_id"`
	Network      ConfigDetail  `json:"network"`
	Worker       ConfigDetail  `json:"worker"`
	EndpointID   *ConfigDetail `json:"endpoint_id,omitempty"`
	Endpoint     *Endpoint     `json:"endpoint,omitempty"`
	IPFSGateways *ConfigDetail `json:"ipfs_gateways,omitempty"`
	Parameters   *Parameters   `json:"parameters,omitempty"`
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
		worker.Aave: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Aave,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Aavegotchi: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Aavegotchi,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Core: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Core,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Crossbell: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Crossbell,
				Description: "Your evm worker name",
			},
			IPFSGateways: &ConfigDetail{
				IsRequired:  false,
				Type:        "[]string",
				Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Curve: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Curve,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.ENS: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Aave,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Highlight: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Highlight,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.IQWiki: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.IQWiki,
				Description: "Your evm worker name",
			},
			IPFSGateways: &ConfigDetail{
				IsRequired:  false,
				Type:        "[]string",
				Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.KiwiStand: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.KiwiStand,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Lens: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Lens,
				Description: "Your evm worker name",
			},
			IPFSGateways: &ConfigDetail{
				IsRequired:  false,
				Type:        "[]string",
				Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service. You are recommended to use your own ipfs gateways because the default ones are not stable for lens",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Lido: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Lido,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Looksrare: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Looksrare,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Matters: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Matters,
				Description: "Your evm worker name",
			},
			IPFSGateways: &ConfigDetail{
				IsRequired:  false,
				Type:        "[]string",
				Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.OpenSea: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.OpenSea,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Optimism: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Optimism,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.RSS3: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.RSS3,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.SAVM: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.SAVM,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Stargate: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Stargate,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.Uniswap: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Uniswap,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
		worker.VSL: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.VSL,
				Description: "Your evm worker name",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "If you don't define the endpoint id, you should set the endpoint url here",
				},
				HTTPHeaders: &ConfigDetail{
					IsRequired:  false,
					Type:        "map[string]string",
					Description: "HTTP headers for network client",
				},
				HTTP2Disabled: &ConfigDetail{
					IsRequired:  false,
					Type:        "bool",
					Description: "You can disable HTTP/2 for the network client",
				},
			},
			Parameters: &Parameters{
				BlockNumberStart: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will start from this block number, if it's not defined, we will use the recommended one(~3 months)"},
				BlockNumberTarget: &ConfigDetail{
					IsRequired:  false,
					Type:        "big.Int",
					Description: "Your worker will stop at this block number, if it's not defined, your worker will always indexing",
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
		},
	},
	network.ArweaveSource: {
		worker.Mirror: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       network.Arweave,
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Mirror,
				Description: "Your arweave worker name",
			},
			IPFSGateways: &ConfigDetail{
				IsRequired:  false,
				Type:        "[]string",
				Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service",
			},
			Parameters: &Parameters{
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
		},
		worker.Momoka: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       network.Arweave,
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Momoka,
				Description: "Your arweave worker name",
			},
			IPFSGateways: &ConfigDetail{
				IsRequired:  false,
				Type:        "[]string",
				Description: "You can define your own ipfs gateways instead of using the default ones if your worker heavily depends on ipfs service. You are recommended to use your own ipfs gateways because the default ones are not stable for momoka",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "You should set polygon endpoint url for momoka worker because it depends on lens contract",
				},
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You should set polygon endpoint id for momoka worker because it depends on lens contract",
			},
			Parameters: &Parameters{
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
		},
		worker.Paragraph: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       network.Arweave,
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Paragraph,
				Description: "Your arweave worker name",
			},
			Parameters: &Parameters{
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
		},
	},
	network.FarcasterSource: {
		worker.Core: {
			WorkerID: ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define your own worker id, it will run as an independent service, the default is `network.worker`",
			},
			Network: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       network.Farcaster,
				Description: "Network where the worker is running on",
			},
			Worker: ConfigDetail{
				IsRequired:  true,
				Type:        "string",
				Value:       worker.Core,
				Description: "Your farcaster worker name, currently only support core",
			},
			EndpointID: &ConfigDetail{
				IsRequired:  false,
				Type:        "string",
				Description: "You can define a global endpoint id for later use, and if the endpoint id is not found in the endpoints list, it will use the endpoint id as the url.",
			},
			Endpoint: &Endpoint{
				URL: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "You should run your own farcaster hub and set the endpoint url here",
				},
			},
			Parameters: &Parameters{
				APIKey: &ConfigDetail{
					IsRequired:  false,
					Type:        "string",
					Description: "API key for your own farcaster hub",
				},
			},
		},
	},
}
