package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=UnknownType --transform=snake --trimprefix=Type --output type_unknown_string.go --json --sql
type UnknownType uint64

//goland:noinspection GoMixedReceiverTypes
func (i UnknownType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i UnknownType) Tag() Tag {
	return TagUnknown
}

const (
	TypeUnknown UnknownType = iota
)
