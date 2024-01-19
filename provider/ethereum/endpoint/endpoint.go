package endpoint

import (
	"os"
	"strings"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
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
	filter.NetworkCrossbell: {
		"https://rpc.crossbell.io",
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
