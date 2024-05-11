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
	URL           *ConfigDetail `json:"network"`
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

// NetworkArchToConfigMap is a map of network arch to config.
var NetworkArchToConfigMap = map[network.Source]workerConfig{
	network.EthereumSource: {
		WorkerID: ConfigDetail{
			IsRequired:  false,
			Type:        "string",
			Description: "You can define your own worker id, it will run as an independent service",
		},
		Network: ConfigDetail{
			IsRequired:  true,
			Type:        "string",
			Description: "Network where the worker is running on",
		},
		Worker: ConfigDetail{
			IsRequired:  true,
			Type:        "string",
			Description: "Your arweave worker name",
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
	network.ArweaveSource: {
		WorkerID: ConfigDetail{
			IsRequired:  false,
			Type:        "string",
			Description: "You can define your own worker id, it will run as an independent service",
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
	network.FarcasterSource: {
		WorkerID: ConfigDetail{
			IsRequired:  false,
			Type:        "string",
			Description: "You can define your own worker id, it will run as an independent service",
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
}
