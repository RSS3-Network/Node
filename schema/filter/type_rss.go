package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=RSSType --transform=snake --trimprefix=TypeRSS --output type_rss_string.go --json --sql
type RSSType uint64

//goland:noinspection GoMixedReceiverTypes
func (r RSSType) Name() string {
	return r.String()
}

//goland:noinspection GoMixedReceiverTypes
func (r RSSType) Tag() Tag {
	return TagRSS
}

const (
	TypeRSSFeed RSSType = iota + 1
)
