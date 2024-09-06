package endpoint

import (
	"os"
	"strings"

	networkx "github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

var endpointsMap = map[networkx.Network][]string{
	networkx.Ethereum: {
		"https://rpc.ankr.com/eth",
		"https://ethereum.blockpi.network/v1/rpc/public",
		"https://eth.llamarpc.com",
	},
	networkx.Optimism: {
		"https://rpc.ankr.com/optimism",
		"https://optimism.blockpi.network/v1/rpc/public",
	},
	networkx.Polygon: {
		"https://rpc.ankr.com/polygon",
		"https://polygon.blockpi.network/v1/rpc/public",
		"https://polygon.llamarpc.com",
	},
	networkx.Base: {
		"https://mainnet.base.org",
		"https://rpc.ankr.com/base",
		"https://base.blockpi.network/v1/rpc/public",
		"https://base.llamarpc.com",
	},
	networkx.Arbitrum: {
		"https://arb1.arbitrum.io/rpc",
		"https://rpc.ankr.com/arbitrum",
		"https://arbitrum.blockpi.network/v1/rpc/public",
		"https://arbitrum.llamarpc.com",
	},
	networkx.Avalanche: {
		"https://api.avax.network/ext/bc/C/rpc",
		"https://rpc.ankr.com/avalanche",
		"https://avalanche.blockpi.network/v1/rpc/public",
	},
	networkx.Crossbell: {
		"https://rpc.crossbell.io",
	},
	networkx.VSL: {
		"https://rpc.rss3.io",
	},
	networkx.SatoshiVM: {
		"https://alpha-rpc-node-http.svmscan.io",
	},
	networkx.BinanceSmartChain: {
		"https://rpc.ankr.com/bsc",
	},
	networkx.Gnosis: {
		"https://rpc.ankr.com/gnosis",
	},
	networkx.Linea: {
		"https://rpc.linea.build",
	},
	networkx.XLayer: {
		"https://rpc.xlayer.tech",
		"https://xlayerrpc.okx.com",
	},
}

func Get(n networkx.Network) ([]string, bool) {
	endpoints, exists := endpointsMap[n]
	if exists && len(endpoints) == 0 {
		return nil, false
	}

	if !exists {
		return nil, exists
	}

	return endpoints, exists
}

func MustGet(n networkx.Network) string {
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
				network, err := networkx.NetworkString(splits[1])
				if err != nil {
					continue
				}

				if _, exists := endpointsMap[network]; !exists {
					endpointsMap[network] = make([]string, 0)
				}

				endpointsMap[network] = append([]string{value}, endpointsMap[network]...)
			}
		}
	}
}
