package filter

import "github.com/labstack/echo/v4"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --yaml --sql
type Network uint64

const (
	NetworkUnknown   Network = iota // unknown
	NetworkEthereum                 // ethereum
	NetworkOptimism                 // optimism
	NetworkBase                     // base
	NetworkPolygon                  // polygon
	NetworkCrossbell                // crossbell
	NetworkArbitrum                 // arbitrum
	NetworkFantom                   // fantom
	NetworkRSS                      // rss
	NetworkArweave                  // arweave
	NetworkFarcaster                // farcaster
)

var _ echo.BindUnmarshaler = (*Network)(nil)

func (n *Network) UnmarshalParam(param string) error {
	network, err := NetworkString(param)
	if err != nil {
		return err
	}

	*n = network

	return nil
}

type NetworkSource string

const (
	NetworkEthereumSource  NetworkSource = "ethereum"
	NetworkArweaveSource   NetworkSource = "arweave"
	NetworkFarcasterSource NetworkSource = "farcaster"
)

func (n Network) Source() NetworkSource {
	switch n {
	case NetworkEthereum, NetworkPolygon, NetworkOptimism, NetworkArbitrum, NetworkFantom, NetworkBase, NetworkCrossbell:
		return NetworkEthereumSource
	case NetworkArweave:
		return NetworkArweaveSource
	case NetworkFarcaster:
		return NetworkFarcasterSource
	default:
		return ""
	}
}

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql
type EthereumChainID uint64

const (
	EthereumChainIDMainnet   EthereumChainID = 1     // ethereum
	EthereumChainIDOptimism  EthereumChainID = 10    // optimism
	EthereumChainIDPolygon   EthereumChainID = 137   // polygon
	EthereumChainIDArbitrum  EthereumChainID = 42161 // arbitrum
	EthereumChainIDFantom    EthereumChainID = 250   // fantom
	EthereumChainIDBase      EthereumChainID = 8453  // base
	EthereumChainIDCrossbell EthereumChainID = 3737  // crossbell
)

func IsOptimismSuperchain(chainID uint64) bool {
	switch chainID {
	case uint64(EthereumChainIDOptimism),
		uint64(EthereumChainIDBase):
		return true
	default:
		return false
	}
}
