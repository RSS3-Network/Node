package filter

import "fmt"

var _ Chain = (*ChainRSS)(nil)

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=ChainRSS --linecomment --output chain_rss_string.go --json --sql
type ChainRSS uint64

//goland:noinspection GoMixedReceiverTypes
func (i ChainRSS) Network() Network {
	return NetworkRSS
}

//goland:noinspection GoMixedReceiverTypes
func (i ChainRSS) ID() uint64 {
	return uint64(i)
}

//goland:noinspection GoMixedReceiverTypes
func (i ChainRSS) FullName() string {
	return fmt.Sprintf("%s.%s", i.Network().String(), i.String())
}

const (
	ChainRSSRSSHub ChainRSS = 1 // rsshub
)
