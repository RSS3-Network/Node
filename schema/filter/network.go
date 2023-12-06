package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --yaml --sql
type Network uint64

const (
	NetworkEthereum  Network = iota // ethereum
	NetworkRSS                      // rss
	NetworkFarcaster                // farcaster
)
