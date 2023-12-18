package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --yaml --sql
type Network uint64

const (
	NetworkUnknown  Network = iota // unknown
	NetworkEthereum                // ethereum
	NetworkRSSHub                  // rsshub
	NetworkArweave                 // arweave
)

type NetworkSource string

const (
	NetworkEthereumSource NetworkSource = "ethereum"
	NetworkArweaveSource  NetworkSource = "arweave"
)

func (n Network) Source() NetworkSource {
	switch n {
	case NetworkEthereum:
		return NetworkEthereumSource
	case NetworkArweave:
		return NetworkArweaveSource
	default:
		return ""
	}
}

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql
type EthereumChainID uint64

const (
	EthereumChainIDMainnet EthereumChainID = 1 // ethereum
)
