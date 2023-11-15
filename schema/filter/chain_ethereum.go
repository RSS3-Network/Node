package filter

var _ Chain = (*ChainEthereum)(nil)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=ChainEthereum --linecomment --output chain_ethereum_string.go --json --sql
type ChainEthereum uint64

//goland:noinspection GoMixedReceiverTypes
func (i ChainEthereum) Network() Network {
	return NetworkEthereum
}

//goland:noinspection GoMixedReceiverTypes
func (i ChainEthereum) ID() uint64 {
	return uint64(i)
}

const (
	ChainEthereumMainnet ChainEthereum = 1 // mainnet
)
