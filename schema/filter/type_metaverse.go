package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MetaverseType --transform=snake --trimprefix=TypeMetaverse --output type_metaverse_string.go --json --sql
type MetaverseType uint64

//goland:noinspection GoMixedReceiverTypes
func (i MetaverseType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i MetaverseType) Tag() Tag {
	return TagMetaverse
}

const (
	TypeMetaverseTransfer MetaverseType = iota + 1
	TypeMetaverseMint
	TypeMetaverseBurn
	TypeMetaverseTrade
)
