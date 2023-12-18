package endpoint

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
)

var endpointsMap = map[filter.Network][]string{
	filter.NetworkEthereum: {
		"https://rpc.ankr.com/eth",
		"https://ethereum.blockpi.network/v1/rpc/public",
		"https://eth.llamarpc.com",
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
