package filter

import "fmt"

type Chain interface {
	Network() Network
	String() string
	ID() uint64
	FullName() string
}

func ChainString(network Network, chainValue string) (Chain, error) {
	switch network {
	case NetworkEthereum:
		return ChainEthereumString(chainValue)
	case NetworkRSS:
		return ChainRSSString(chainValue)
	default:
		return nil, fmt.Errorf("unknown network: %s", network)
	}
}
