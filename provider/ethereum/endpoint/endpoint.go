package endpoint

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
)

var endpointsMap = map[filter.ChainEthereum][]string{
	filter.ChainEthereumMainnet: {
		"https://rpc.ankr.com/eth",
		"https://ethereum.blockpi.network/v1/rpc/public",
		"https://eth.llamarpc.com",
	},
}

func Get(chain filter.ChainEthereum) ([]string, bool) {
	endpoints, exists := endpointsMap[chain]
	if exists && len(endpoints) == 0 {
		return nil, !exists
	}

	if !exists {
		return nil, exists
	}

	return endpoints, exists
}

func MustGet(chain filter.ChainEthereum) string {
	return lo.Must(Get(chain))[0]
}
