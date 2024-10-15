package info

// This file includes the relations between networks and their associated workers.
// Results are exclusively provided to external users and do not influence internal operations within the Node.
// Each worker has a default configuration, which can be customized based on various factors.

import (
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/node/schema/worker/federated"
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
	KafkaTopic              *ConfigDetail   `json:"kafka_topic,omitempty"`
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

const (
	activityPubKafkaTopicDescription = "The Kafka topic to which the ActivityPub data will be published. By default, the data will be sent to this topic on the Kafka broker running on the Mastodon instance."

	mastodonInstanceDescription = "A Mastodon instance is required. Please follow <a href=\"https://github.com/RSS3-Network/Mastodon-Instance-Kit\" target=\"_blank\">the guide</a> to either deploy a new Mastodon instance or modify an existing Mastodon instance. After completing either option, enter your Mastodon endpoint (format: your_instance_ip:9092) here."
)

var defaultNetworkParameters = map[network.Source]*Parameters{
	network.ActivityPubSource: {
		KafkaTopic: &ConfigDetail{
			IsRequired:  true,
			Type:        StringType,
			Value:       "activitypub_events",
			Description: activityPubKafkaTopicDescription,
			Title:       "Kafka Topic",
			Key:         "parameters.kafka_topic",
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
			Title:       "Concurrent Block Requests",
			Key:         "parameters.concurrent_block_requests",
		},
	},
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
	network.NearSource: {
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

	switch n.Source() {
	case network.EthereumSource:
		endpointConfig.URL.Value = endpoint.MustGet(n)
	case network.NearSource:
		endpointConfig.URL.Value = "https://archival-rpc.mainnet.near.org"
	case network.FarcasterSource:
		endpointConfig.URL.Value = "https://your-farcaster-api-endpoint"
	case network.ArweaveSource:
		endpointConfig.URL.Value = "https://arweave.net"
	case network.ActivityPubSource:
		endpointConfig.URL.Type = StringType
		endpointConfig.URL.Description = mastodonInstanceDescription
		endpointConfig.URL.Value = "127.0.0.1:9092"
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
		decentralized.Curve,
		decentralized.Highlight,
		decentralized.Rainbow,
		decentralized.Stargate,
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
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Zerion,
	},
	network.Base: {
		decentralized.Aave,
		decentralized.Base,
		decentralized.Core,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Zerion,
	},
	network.BinanceSmartChain: {
		decentralized.Core,
		decentralized.Rainbow,
		decentralized.Stargate,
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
		decentralized.Core,
		decentralized.Curve,
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
		decentralized.Optimism,
		decentralized.Rainbow,
		decentralized.Stargate,
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
		decentralized.Polymarket,
		decentralized.Rainbow,
		decentralized.Stargate,
		decentralized.Zerion,
	},
	network.RSS: {
		rss.RSSHub,
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
		decentralized.Zerion,
	},
}

// WorkerToConfigMap is a map of worker to config.
var WorkerToConfigMap = map[network.Source]map[worker.Worker]workerConfig{
	network.ActivityPubSource: {
		federated.Mastodon: customWorkerConfig(federated.Mastodon, network.ActivityPubSource, &Parameters{
			KafkaTopic: &ConfigDetail{
				IsRequired:  true,
				Type:        StringType,
				Value:       "activitypub_events",
				Description: activityPubKafkaTopicDescription,
				Title:       "Kafka Topic",
				Key:         "parameters.kafka_topic",
			},
		}, mastodonInstanceDescription),
	},
	network.ArweaveSource: {
		decentralized.Mirror:    customWorkerConfigWithoutEndpoint(decentralized.Mirror, network.ArweaveSource, nil, true),
		decentralized.Momoka:    customWorkerConfigWithIPFS(decentralized.Momoka, network.ArweaveSource, "A Polygon RPC is required for Momoka"),
		decentralized.Paragraph: customWorkerConfigWithoutEndpoint(decentralized.Paragraph, network.ArweaveSource, nil, false),
	},
	network.EthereumSource: {
		decentralized.Aave:       defaultWorkerConfig(decentralized.Aave, network.EthereumSource, nil),
		decentralized.Aavegotchi: defaultWorkerConfig(decentralized.Aavegotchi, network.EthereumSource, nil),
		decentralized.Arbitrum:   defaultWorkerConfig(decentralized.Arbitrum, network.EthereumSource, nil),
		decentralized.BendDAO:    defaultWorkerConfig(decentralized.BendDAO, network.EthereumSource, nil),
		decentralized.Base:       defaultWorkerConfig(decentralized.Base, network.EthereumSource, nil),
		decentralized.Core:       defaultWorkerConfig(decentralized.Core, network.EthereumSource, nil),
		decentralized.Cow:        defaultWorkerConfig(decentralized.Cow, network.EthereumSource, nil),
		decentralized.Crossbell:  customWorkerConfigWithIPFS(decentralized.Crossbell, network.EthereumSource, ""),
		decentralized.Curve:      defaultWorkerConfig(decentralized.Curve, network.EthereumSource, nil),
		decentralized.ENS:        defaultWorkerConfig(decentralized.ENS, network.EthereumSource, nil),
		decentralized.Highlight:  defaultWorkerConfig(decentralized.Highlight, network.EthereumSource, nil),
		decentralized.IQWiki:     customWorkerConfigWithIPFS(decentralized.IQWiki, network.EthereumSource, ""),
		decentralized.KiwiStand:  defaultWorkerConfig(decentralized.KiwiStand, network.EthereumSource, nil),
		decentralized.Lens:       customWorkerConfigWithIPFS(decentralized.Lens, network.EthereumSource, ""),
		decentralized.Lido:       defaultWorkerConfig(decentralized.Lido, network.EthereumSource, nil),
		decentralized.Linea:      defaultWorkerConfig(decentralized.Linea, network.EthereumSource, nil),
		decentralized.Looksrare:  defaultWorkerConfig(decentralized.Looksrare, network.EthereumSource, nil),
		decentralized.Matters:    customWorkerConfigWithIPFS(decentralized.Matters, network.EthereumSource, ""),
		decentralized.Nouns:      defaultWorkerConfig(decentralized.Nouns, network.EthereumSource, nil),
		decentralized.Oneinch:    defaultWorkerConfig(decentralized.Oneinch, network.EthereumSource, nil),
		decentralized.OpenSea:    defaultWorkerConfig(decentralized.OpenSea, network.EthereumSource, nil),
		decentralized.Optimism:   defaultWorkerConfig(decentralized.Optimism, network.EthereumSource, nil),
		decentralized.Paraswap:   defaultWorkerConfig(decentralized.Paraswap, network.EthereumSource, nil),
		decentralized.Polymarket: defaultWorkerConfig(decentralized.Polymarket, network.EthereumSource, nil),
		decentralized.Rainbow:    defaultWorkerConfig(decentralized.Rainbow, network.EthereumSource, nil),
		decentralized.RSS3:       defaultWorkerConfig(decentralized.RSS3, network.EthereumSource, nil),
		decentralized.SAVM:       defaultWorkerConfig(decentralized.SAVM, network.EthereumSource, nil),
		decentralized.Stargate:   defaultWorkerConfig(decentralized.Stargate, network.EthereumSource, nil),
		decentralized.Uniswap:    defaultWorkerConfig(decentralized.Uniswap, network.EthereumSource, nil),
		decentralized.VSL:        defaultWorkerConfig(decentralized.VSL, network.EthereumSource, nil),
		decentralized.Zerion:     defaultWorkerConfig(decentralized.Zerion, network.EthereumSource, nil),
	},
	network.FarcasterSource: {
		decentralized.Core: customWorkerConfig(decentralized.Core, network.FarcasterSource, &Parameters{
			APIKey: &ConfigDetail{
				IsRequired:  false,
				Type:        StringType,
				Description: "API key to access your Farcaster Hubble on Neynar",
				Title:       "API Key",
				Key:         "parameters.api_key",
			},
		}, "A Farcaster Hubble is required"),
	},
	network.NearSource: {
		decentralized.Core:       defaultWorkerConfig(decentralized.Core, network.NearSource, nil),
		decentralized.LiNEAR:     defaultWorkerConfig(decentralized.LiNEAR, network.NearSource, nil),
		decentralized.NearSocial: defaultWorkerConfig(decentralized.NearSocial, network.NearSource, nil),
	},
	network.RSSSource: {
		rss.RSSHub: customWorkerConfig(rss.RSSHub, network.RSSSource, &Parameters{
			Authentication: &Authentication{
				AccessKey: &ConfigDetail{
					IsRequired:  false,
					Type:        StringType,
					Description: "A key to access the RSS Feed",
					Title:       "Access Key",
					Key:         "parameters.authentication.access_key",
				},
			},
		}, "Your RSSHub instance URL"),
	},
}
