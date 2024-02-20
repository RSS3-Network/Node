package endpoint

import (
	"os"
	"strings"

	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
)

var endpointsMap = map[filter.Network][]string{
	filter.NetworkEthereum: {
		"https://rpc.ankr.com/eth",
		"https://ethereum.blockpi.network/v1/rpc/public",
		"https://eth.llamarpc.com",
	},
	filter.NetworkOptimism: {
		"https://rpc.ankr.com/optimism",
		"https://optimism.blockpi.network/v1/rpc/public",
	},
	filter.NetworkPolygon: {
		"https://rpc.ankr.com/polygon",
		"https://polygon.blockpi.network/v1/rpc/public",
		"https://polygon.llamarpc.com",
	},
	filter.NetworkBase: {
		"https://mainnet.base.org",
		"https://rpc.ankr.com/base",
		"https://base.blockpi.network/v1/rpc/public",
		"https://base.llamarpc.com",
	},
	filter.NetworkArbitrum: {
		"https://arb1.arbitrum.io/rpc",
		"https://rpc.ankr.com/arbitrum",
		"https://arbitrum.blockpi.network/v1/rpc/public",
		"https://arbitrum.llamarpc.com",
	},
	filter.NetworkFantom: {
		"https://rpcapi.fantom.network",
		"https://rpc.ankr.com/fantom",
		"https://fantom.blockpi.network/v1/rpc/public",
	},
	filter.NetworkAvalanche: {
		"https://api.avax.network/ext/bc/C/rpc",
		"https://rpc.ankr.com/avalanche",
		"https://avalanche.blockpi.network/v1/rpc/public",
	},
	filter.NetworkCrossbell: {
		"https://rpc.crossbell.io",
	},
	filter.NetworkRSS3Testnet: {
		"https://rpc.testnet.rss3.io",
	},
}

func Get(network filter.Network) ([]string, bool) {
	endpoints, exists := endpointsMap[network]
	if exists && len(endpoints) == 0 {
		return nil, !exists
	}

	if !exists {
		return nil, exists
	}

	return endpoints, exists
}

func MustGet(network filter.Network) string {
	return lo.Must(Get(network))[0]
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
				network, err := filter.NetworkString(splits[1])
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
