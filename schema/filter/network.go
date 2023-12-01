package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Network --linecomment --output network_string.go --json --sql
type Network uint64

const (
	NetworkEthereum Network = iota // ethereum
	NetworkRSSHub                  // rsshub
	NetworkArweave                 // arweave
)
