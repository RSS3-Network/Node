package info

// This file includes the relations between networks and their associated workers.
// Results are exclusively provided to external users and do not influence internal operations within the Node.
// Each worker has a default configuration, which can be customized based on various factors.

import (
	"github.com/rss3-network/node/v2/provider/ethereum/endpoint"
	"github.com/rss3-network/node/v2/schema/worker"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/node/v2/schema/worker/federated"
	"github.com/rss3-network/node/v2/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
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
	Title       string                `json:"title"`
	Key         string                `json:"key"`
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
	TimestampStart          *ConfigDetail   `json:"timestamp_start,omitempty"`
	RelayURLList            *ConfigDetail   `json:"relay_url_list,omitempty"`
	Port                    *ConfigDetail   `json:"port,omitempty"`
	AgentdataDBURL          *ConfigDetail   `json:"agentdata_db_url,omitempty"`
	OpenAIAPIKey            *ConfigDetail   `json:"openai_api_key,omitempty"`
	OllamaHost              *ConfigDetail   `json:"ollama_host,omitempty"`
	KaitoAPIToken           *ConfigDetail   `json:"kaito_api_token,omitempty"`
	TwitterConfig           *AITwitter      `json:"twitter,omitempty"`
}

type workerConfig struct {
	ID              ConfigDetail    `json:"id"`
	Network         ConfigDetail    `json:"network,omitempty"`
	Worker          ConfigDetail    `json:"worker"`
	EndpointID      *ConfigDetail   `json:"endpoint,omitempty"`
	IPFSGateways    *ConfigDetail   `json:"ipfs_gateways,omitempty"`
	Parameters      *Parameters     `json:"parameters,omitempty"`
	MinimumResource MinimumResource `json:"minimum_resource"`
}

const (
	relayURLArrayTypeDescription = "List of relay URLs to follow and receive messages from"

	domainPortDescription = "The port number that the mastodon endpoint domain will listen on. This should be an available network port"

	mastodonInstanceEndpointDescription = "Your Mastodon instance must be accessible via a public URL, which exposes your local instance on the port number you select. You can use services like ngrok (https://ngrok.com)."
)

var defaultNetworkParameters = map[network.Protocol]*Parameters{
	network.ActivityPubProtocol: {
		RelayURLList: &ConfigDetail{
			IsRequired:  true,
			Type:        URLArrayType,
			Value:       []string{"https://relay.fedi.buzz/instance/mastodon.social"},
			Description: relayURLArrayTypeDescription,
			Title:       "Relay URL List",
			Key:         "parameters.relay_url_list",
		},
	},
	network.ArweaveProtocol: {
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
			Title:       "Concurrent Block Requests",
			Key:         "parameters.concurrent_block_requests",
		},
	},
	network.EthereumProtocol: {
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
			Title:       "Concurrent Block Requests",
			Key:         "parameters.concurrent_block_requests",
		},
		BlockBatchSize: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(8),
			Description: "The number of blocks to fetch in a single RPC request. Default: 8",
			Title:       "Block Batch Size",
			Key:         "parameters.block_batch_size",
		},
		ReceiptsBatchSize: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(200),
			Description: "The number of receipts to fetch in a single RPC request. Default: 200",
			Title:       "Receipts Batch Size",
			Key:         "parameters.receipts_batch_size",
		},
		BlockReceiptBatchSize: &ConfigDetail{
			IsRequired:  false,
			Type:        UintType,
			Value:       uint(8),
			Description: "The number of blocks to fetch receipts in a single RPC request. Default: 8",
			Title:       "Block Receipt Batch Size",
			Key:         "parameters.block_receipts_batch_size",
		},
	},
	network.NearProtocol: {
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
			Value:       uint(8),
			Description: "The number of concurrent RPC requests to the Near RPC. Default: 8",
			Title:       "Concurrent Block Requests",
			Key:         "parameters.concurrent_block_requests",
		},
	},
}

// getDefaultParametersByNetwork returns the default parameters based on the network.
func getDefaultParametersByNetwork(protocol network.Protocol) *Parameters {
	return defaultNetworkParameters[protocol]
}

// defaultWorkerConfig returns the default worker config based on the worker and network.
// If parameters are supplied, use them instead of the default parameters.
func defaultWorkerConfig(worker worker.Worker, protocol network.Protocol, parameters *Parameters) workerConfig {
	// generate default parameters only if parameters are not provided
	if parameters == nil {
		parameters = getDefaultParametersByNetwork(protocol)
	}

	config := workerConfig{
		ID: ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Description: "Worker's id, must be unique, for example '[network]-[worker]'",
			Title:       "ID",
			Key:         "id",
		},
		Network: ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Description: "The network where the worker operates on",
			Title:       "Network",
			Key:         "network",
		},
		Worker: ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Value:       worker.Name(),
			Description: "Name of the worker",
			Title:       "Worker",
			Key:         "worker",
		},
		EndpointID: &ConfigDetail{
			IsRequired:  true,
			Type:        URLType,
			Description: "An external endpoint to fetch data from, for example, a blockchain RPC endpoint or a Farcaster api",
			Title:       "Endpoint",
			Key:         "endpoint",
		},
		Parameters: parameters,
	}

	return config
}

// customWorkerConfigWithoutEndpoint generates a config with custom fields and no endpoint.
func customWorkerConfigWithoutEndpoint(worker worker.Worker, protocol network.Protocol, parameters *Parameters, requireIPFS bool) workerConfig {
	config := defaultWorkerConfig(worker, protocol, parameters)

	config.EndpointID = nil

	if requireIPFS {
		setIPFSGateways(&config)
	}

	return config
}

// customWorkerConfigWithIPFS generates a config with IPFS and custom fields.
func customWorkerConfigWithIPFS(worker worker.Worker, protocol network.Protocol, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, protocol, nil)

	setIPFSGateways(&config)

	if len(endpointDescription) > 0 {
		config.EndpointID.Description = endpointDescription
	}

	return config
}

// customWorkerConfig generates a config with custom fields.
func customWorkerConfig(worker worker.Worker, protocol network.Protocol, parameters *Parameters, endpointDescription string) workerConfig {
	config := defaultWorkerConfig(worker, protocol, parameters)

	// Update the EndpointID description based on the provided custom description
	config.EndpointID.Description = endpointDescription

	return config
}

// getEndpointConfig returns the endpoint config for a given network.
func getEndpointConfig(n network.Network) Endpoint {
	endpointConfig := Endpoint{
		URL: &ConfigDetail{
			IsRequired:  true,
			Type:        URLType,
			Description: "The URL of the endpoint.",
			Title:       "URL",
			Key:         "url",
		},
		HTTPHeaders: &ConfigDetail{
			IsRequired:  false,
			Type:        StringMapType,
			Description: "HTTP headers to be sent with requests to this endpoint.",
			Title:       "HTTP Headers",
			Key:         "http_headers",
		},
		HTTP2Disabled: &ConfigDetail{
			IsRequired:  false,
			Type:        BooleanType,
			Description: "Some endpoints may not support HTTP2, set this to true to disable HTTP2.",
			Title:       "HTTP2 Disabled",
			Key:         "http2_disabled",
		},
	}

	switch n.Protocol() {
	case network.EthereumProtocol:
		endpointConfig.URL.Value = endpoint.MustGet(n)
	case network.NearProtocol:
		endpointConfig.URL.Value = "https://archival-rpc.mainnet.near.org"
	case network.FarcasterProtocol:
		endpointConfig.URL.Value = "https://your-farcaster-api-endpoint"
	case network.ArweaveProtocol:
		endpointConfig.URL.Value = "https://arweave.net"
	case network.ActivityPubProtocol:
		endpointConfig.URL.Type = StringType
		endpointConfig.URL.Description = mastodonInstanceEndpointDescription
		endpointConfig.URL.Value = "https://domain.ngrok.app"
	default:
		endpointConfig.URL.Value = "https://your-network-endpoint"
	}

	return endpointConfig
}

func setIPFSGateways(config *workerConfig) {
	config.IPFSGateways = &ConfigDetail{
		IsRequired:  true,
		Type:        URLArrayType,
		Description: "A list of IPFS gateways to fetch data from, multiple gateways may improve the availability of IPFS data",
		Title:       "IPFS Gateways",
		Key:         "ipfs_gateways",
	}
}

// NetworkToWorkersMap is a map of Network-has-Workers.
var NetworkToWorkersMap = map[network.Network][]worker.Worker{
	network.Arbitrum: {
		decentralized.Aave,
		decentralized.Arbitrum,
		decentralized.Core,
		decentralized.Cow,
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.Oneinch,
		decentralized.Paraswap,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.Arweave: {
		decentralized.Mirror,
		decentralized.Momoka,
		decentralized.Paragraph,
	},
	network.Avalanche: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Oneinch,
		decentralized.Paraswap,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.Base: {
		decentralized.Aave,
		decentralized.Base,
		decentralized.Core,
		decentralized.Cow,
		decentralized.Curve,
		decentralized.Oneinch,
		decentralized.Paraswap,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.BinanceSmartChain: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Oneinch,
		decentralized.Paraswap,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.Crossbell: {
		decentralized.Core,
		decentralized.Crossbell,
	},
	network.Ethereum: {
		decentralized.Aave,
		decentralized.Arbitrum,
		decentralized.Base,
		decentralized.Core,
		decentralized.Cow,
		decentralized.Curve,
		decentralized.ENS,
		decentralized.Highlight,
		decentralized.Lido,
		decentralized.Linea,
		decentralized.Looksrare,
		decentralized.Nouns,
		decentralized.Oneinch,
		decentralized.OpenSea,
		decentralized.Optimism,
		decentralized.Paraswap,
		decentralized.Rainbow,
		decentralized.RSS3,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.VSL,
	},
	network.Farcaster: {
		decentralized.Core,
	},
	network.Gnosis: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Cow,
		decentralized.Curve,
		decentralized.Oneinch,
		decentralized.Zerion,
	},
	network.Linea: {
		decentralized.Core,
		decentralized.Linea,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.Mastodon: {
		federated.Mastodon,
	},
	network.Near: {
		decentralized.Core,
		decentralized.LiNEAR,
		decentralized.NearSocial,
	},
	network.Optimism: {
		decentralized.Aave,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.KiwiStand,
		decentralized.Matters,
		decentralized.Oneinch,
		decentralized.Optimism,
		decentralized.Paraswap,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.Polygon: {
		decentralized.Aave,
		decentralized.Aavegotchi,
		decentralized.Core,
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.IQWiki,
		decentralized.Lens,
		decentralized.Oneinch,
		decentralized.Paraswap,
		decentralized.Polymarket,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Uniswap,
		decentralized.Zerion,
	},
	network.RSSHub: {
		rss.Core,
	},
	network.SatoshiVM: {
		decentralized.Core,
		decentralized.SAVM,
		decentralized.Uniswap,
	},
	network.VSL: {
		decentralized.Core,
	},
	network.XLayer: {
		decentralized.Core,
		decentralized.Curve,
		decentralized.Zerion,
	},
}

// WorkerToConfigMap is a map of worker to config.
var WorkerToConfigMap = map[network.Protocol]map[worker.Worker]workerConfig{
	network.ActivityPubProtocol: {
		federated.Mastodon: customWorkerConfig(federated.Mastodon, network.ActivityPubProtocol, &Parameters{
			RelayURLList: &ConfigDetail{
				IsRequired:  false,
				Type:        URLArrayType,
				Description: relayURLArrayTypeDescription,
				Title:       "Relay URL List",
				Key:         "parameters.relay_url_list",
			},
			Port: &ConfigDetail{
				IsRequired:  false,
				Type:        UintType,
				Description: domainPortDescription,
				Title:       "Port Number",
				Key:         "parameters.port",
			},
		}, mastodonInstanceEndpointDescription),
		federated.Bluesky: customWorkerConfigWithoutEndpoint(federated.Bluesky, network.ActivityPubProtocol, nil, false),
	},
	network.ArweaveProtocol: {
		decentralized.Mirror:    customWorkerConfigWithoutEndpoint(decentralized.Mirror, network.ArweaveProtocol, nil, true),
		decentralized.Momoka:    customWorkerConfigWithIPFS(decentralized.Momoka, network.ArweaveProtocol, "A Polygon RPC is required for Momoka"),
		decentralized.Paragraph: customWorkerConfigWithoutEndpoint(decentralized.Paragraph, network.ArweaveProtocol, nil, false),
	},
	network.EthereumProtocol: {
		decentralized.Aave:       defaultWorkerConfig(decentralized.Aave, network.EthereumProtocol, nil),
		decentralized.Aavegotchi: defaultWorkerConfig(decentralized.Aavegotchi, network.EthereumProtocol, nil),
		decentralized.Arbitrum:   defaultWorkerConfig(decentralized.Arbitrum, network.EthereumProtocol, nil),
		decentralized.BendDAO:    defaultWorkerConfig(decentralized.BendDAO, network.EthereumProtocol, nil),
		decentralized.Base:       defaultWorkerConfig(decentralized.Base, network.EthereumProtocol, nil),
		decentralized.Core:       defaultWorkerConfig(decentralized.Core, network.EthereumProtocol, nil),
		decentralized.Cow:        defaultWorkerConfig(decentralized.Cow, network.EthereumProtocol, nil),
		decentralized.Crossbell:  customWorkerConfigWithIPFS(decentralized.Crossbell, network.EthereumProtocol, ""),
		decentralized.Curve:      defaultWorkerConfig(decentralized.Curve, network.EthereumProtocol, nil),
		decentralized.ENS:        defaultWorkerConfig(decentralized.ENS, network.EthereumProtocol, nil),
		decentralized.Highlight:  defaultWorkerConfig(decentralized.Highlight, network.EthereumProtocol, nil),
		decentralized.IQWiki:     customWorkerConfigWithIPFS(decentralized.IQWiki, network.EthereumProtocol, ""),
		decentralized.KiwiStand:  defaultWorkerConfig(decentralized.KiwiStand, network.EthereumProtocol, nil),
		decentralized.Lens:       customWorkerConfigWithIPFS(decentralized.Lens, network.EthereumProtocol, ""),
		decentralized.Lido:       defaultWorkerConfig(decentralized.Lido, network.EthereumProtocol, nil),
		decentralized.Linea:      defaultWorkerConfig(decentralized.Linea, network.EthereumProtocol, nil),
		decentralized.Looksrare:  defaultWorkerConfig(decentralized.Looksrare, network.EthereumProtocol, nil),
		decentralized.Matters:    customWorkerConfigWithIPFS(decentralized.Matters, network.EthereumProtocol, ""),
		decentralized.Nouns:      defaultWorkerConfig(decentralized.Nouns, network.EthereumProtocol, nil),
		decentralized.Oneinch:    defaultWorkerConfig(decentralized.Oneinch, network.EthereumProtocol, nil),
		decentralized.OpenSea:    defaultWorkerConfig(decentralized.OpenSea, network.EthereumProtocol, nil),
		decentralized.Optimism:   defaultWorkerConfig(decentralized.Optimism, network.EthereumProtocol, nil),
		decentralized.Paraswap:   defaultWorkerConfig(decentralized.Paraswap, network.EthereumProtocol, nil),
		decentralized.Polymarket: defaultWorkerConfig(decentralized.Polymarket, network.EthereumProtocol, nil),
		decentralized.Rainbow:    defaultWorkerConfig(decentralized.Rainbow, network.EthereumProtocol, nil),
		decentralized.RSS3:       defaultWorkerConfig(decentralized.RSS3, network.EthereumProtocol, nil),
		decentralized.SAVM:       defaultWorkerConfig(decentralized.SAVM, network.EthereumProtocol, nil),
		decentralized.Stargate:   defaultWorkerConfig(decentralized.Stargate, network.EthereumProtocol, nil),
		decentralized.Uniswap:    defaultWorkerConfig(decentralized.Uniswap, network.EthereumProtocol, nil),
		decentralized.VSL:        defaultWorkerConfig(decentralized.VSL, network.EthereumProtocol, nil),
		decentralized.Zerion:     defaultWorkerConfig(decentralized.Zerion, network.EthereumProtocol, nil),
	},
	network.FarcasterProtocol: {
		decentralized.Core: customWorkerConfig(decentralized.Core, network.FarcasterProtocol, &Parameters{
			APIKey: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Description: "API key to access your Farcaster Hubble on Neynar",
				Title:       "API Key",
				Key:         "parameters.api_key",
			},
		}, "A Farcaster Hubble is required"),
	},
	network.NearProtocol: {
		decentralized.Core:       defaultWorkerConfig(decentralized.Core, network.NearProtocol, nil),
		decentralized.LiNEAR:     defaultWorkerConfig(decentralized.LiNEAR, network.NearProtocol, nil),
		decentralized.NearSocial: defaultWorkerConfig(decentralized.NearSocial, network.NearProtocol, nil),
	},
	network.RSSProtocol: {
		rss.Core: customWorkerConfig(rss.Core, network.RSSProtocol, &Parameters{
			Authentication: &Authentication{
				AccessKey: &ConfigDetail{
					IsRequired:  false,
					Type:        StringType,
					Description: "Access key to access your RSSHub",
					Title:       "Access Key",
					Key:         "parameters.authentication.access_key",
				},
			},
		}, "Your RSSHub instance URL"),
	},
}

// AITwitter defines Twitter-specific parameters for AI workers
type AITwitter struct {
	BearerToken       *ConfigDetail `json:"bearer_token,omitempty"`
	APIKey            *ConfigDetail `json:"api_key,omitempty"`
	APISecret         *ConfigDetail `json:"api_secret,omitempty"`
	AccessToken       *ConfigDetail `json:"access_token,omitempty"`
	AccessTokenSecret *ConfigDetail `json:"access_token_secret,omitempty"`
}

// AIWorkerConfig defines the default worker configuration for AI workers
var AIWorkerConfig = workerConfig{
	ID: ConfigDetail{
		IsRequired:  true,
		Type:        StringType,
		Value:       "agentdata-core",
		Description: "Worker's id, must be unique, for example '[network]-[worker]'",
		Title:       "ID",
		Key:         "id",
	},
	Worker: ConfigDetail{
		IsRequired:  true,
		Type:        StringType,
		Value:       "core",
		Description: "Name of the worker",
		Title:       "Worker",
		Key:         "worker",
	},
	EndpointID: &ConfigDetail{
		IsRequired:  false,
		Type:        URLType,
		Description: "Your AgentData service URL",
		Title:       "Endpoint",
		Key:         "endpoint",
	},
	Parameters: &Parameters{},
	MinimumResource: MinimumResource{
		CPUCore:       0.25,
		MemoryInGb:    0.25,
		DiskSpaceInGb: 0,
		Title:         "Minimum Resource",
		Key:           "minimum_resource",
	},
}

// getNetworkConfigDetailForAI generates the AI configuration details
func getNetworkConfigDetailForAI() NetworkConfigDetailForSingleWorker {
	networkDetail := NetworkConfigDetailForSingleWorker{
		ID: "agentdata",
	}

	workerConfig := deepCopyWorkerConfig(AIWorkerConfig)

	workerConfig.Parameters = &Parameters{
		AgentdataDBURL: &ConfigDetail{
			IsRequired:  false,
			Type:        StringType,
			Value:       nil,
			Description: "AgentData database URL",
			Title:       "AgentData DB URL",
			Key:         "agentdata_db_url",
		},
		OpenAIAPIKey: &ConfigDetail{
			IsRequired:  false,
			Type:        StringType,
			Value:       nil,
			Description: "OpenAI API Key for LLM integration",
			Title:       "OpenAI API Key",
			Key:         "openai_api_key",
		},
		OllamaHost: &ConfigDetail{
			IsRequired:  false,
			Type:        StringType,
			Value:       nil,
			Description: "Ollama host URL for local model hosting",
			Title:       "Ollama Host",
			Key:         "ollama_host",
		},
		KaitoAPIToken: &ConfigDetail{
			IsRequired:  false,
			Type:        StringType,
			Value:       nil,
			Description: "Kaito API Token for Kaito integration",
			Title:       "Kaito API Token",
			Key:         "kaito_api_token",
		},
		TwitterConfig: &AITwitter{
			BearerToken: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Value:       nil,
				Description: "Twitter Bearer Token for API access",
				Title:       "Twitter Bearer Token",
				Key:         "twitter.bearer_token",
			},
			APIKey: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Value:       nil,
				Description: "Twitter API Key",
				Title:       "Twitter API Key",
				Key:         "twitter.api_key",
			},
			APISecret: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Value:       nil,
				Description: "Twitter API Secret",
				Title:       "Twitter API Secret",
				Key:         "twitter.api_secret",
			},
			AccessToken: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Value:       nil,
				Description: "Twitter Access Token",
				Title:       "Twitter Access Token",
				Key:         "twitter.access_token",
			},
			AccessTokenSecret: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Value:       nil,
				Description: "Twitter Access Token Secret",
				Title:       "Twitter Access Token Secret",
				Key:         "twitter.access_token_secret",
			},
		},
	}

	networkDetail.WorkerConfig = workerConfig

	zap.L().Debug("successfully retrieved AI network configuration details",
		zap.String("networkID", "agentdata"))

	return networkDetail
}
