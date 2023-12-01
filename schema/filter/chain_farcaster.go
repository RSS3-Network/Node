package filter

import "fmt"

var _ Chain = (*ChainFarcaster)(nil)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=ChainFarcaster --linecomment --output chain_farcaster_string.go --json --sql
type ChainFarcaster uint64

//goland:noinspection GoMixedReceiverTypes
func (i ChainFarcaster) Network() Network {
	return NetworkFarcaster
}

//goland:noinspection GoMixedReceiverTypes
func (i ChainFarcaster) ID() uint64 {
	return uint64(i)
}

//goland:noinspection GoMixedReceiverTypes
func (i ChainFarcaster) FullName() string {
	return fmt.Sprintf("%s.%s", i.Network().String(), i.String())
}

const (
	ChainFarcasterMainnet ChainFarcaster = 1 // mainnet
)
