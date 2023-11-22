package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleType --transform=snake --trimprefix=TypeCollectible --output type_collectible_string.go --json --sql
type CollectibleType uint64

//goland:noinspection GoMixedReceiverTypes
func (i CollectibleType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i CollectibleType) Tag() Tag {
	return TagCollectible
}

const (
	TypeCollectibleTransfer CollectibleType = iota + 1
	TypeCollectibleApproval
	TypeCollectibleMint
	TypeCollectibleBurn
)
