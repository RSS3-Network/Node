package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Tag --transform=snake --trimprefix=Tag --output tag_string.go --json --sql
type Tag uint64

const (
	TagUnknown Tag = iota
	TagTransaction
	TagCollectible
	TagExchange
)
