package endpoint

import (
	"os"
	"strings"

	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

var endpointsMap = map[network.Network][]string{
	network.Ethereum: {
		"https://rpc.ankr.com/eth",
		"https://ethereum.blockpi.network/v1/rpc/public",
		"https://eth.llamarpc.com",
	},
	network.Optimism: {
		"https://rpc.ankr.com/optimism",
		"https://optimism.blockpi.network/v1/rpc/public",
	},
	network.Polygon: {
		"https://rpc.ankr.com/polygon",
		"https://polygon.blockpi.network/v1/rpc/public",
		"https://polygon.llamarpc.com",
	},
	network.Base: {
		"https://mainnet.base.org",
		"https://rpc.ankr.com/base",
		"https://base.blockpi.network/v1/rpc/public",
		"https://base.llamarpc.com",
	},
	network.Arbitrum: {
		"https://arb1.arbitrum.io/rpc",
		"https://rpc.ankr.com/arbitrum",
		"https://arbitrum.blockpi.network/v1/rpc/public",
		"https://arbitrum.llamarpc.com",
	},
	network.Avalanche: {
		"https://api.avax.network/ext/bc/C/rpc",
		"https://rpc.ankr.com/avalanche",
		"https://avalanche.blockpi.network/v1/rpc/public",
	},
	network.Crossbell: {
		"https://rpc.crossbell.io",
	},
	network.VSL: {
		"https://rpc.rss3.io",
	},
	network.SatoshiVM: {
		"https://alpha-rpc-node-http.svmscan.io",
	},
	network.BinanceSmartChain: {
		"https://rpc.ankr.com/bsc",
	},
	network.Gnosis: {
		"https://rpc.ankr.com/gnosis",
	},
	network.Linea: {
		"https://rpc.linea.build",
	},
}

func Get(n network.Network) ([]string, bool) {
	endpoints, exists := endpointsMap[n]
	if exists && len(endpoints) == 0 {
		return nil, false
	}

	if !exists {
		return nil, exists
	}

	return endpoints, exists
}

func MustGet(n network.Network) string {
	return lo.Must(Get(n))[0]
}

func init() {
	// Load endpoints from the environment variables.
	for _, environment := range os.Environ() {
		var (
			splits     = strings.Split(environment, "=")
			key, value = splits[0], splits[1]
		)

		// Example of the environment variable `ENDPOINT_ETHEREUM=https://rpc.ankr.com/eth`.
		if strings.HasPrefix(key, "ENDPOINT_") {
			splits := strings.Split(key, "_")

			if len(splits) == 2 {
				_network, err := network.NetworkString(splits[1])
				if err != nil {
					continue
				}

				if _, exists := endpointsMap[_network]; !exists {
					endpointsMap[_network] = make([]string, 0)
				}

				endpointsMap[_network] = append([]string{value}, endpointsMap[_network]...)
			}
		}
	}
}
