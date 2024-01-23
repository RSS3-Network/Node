package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=SocialType --transform=snake --trimprefix=TypeSocial --output type_social_string.go --json --sql
type SocialType uint64

//goland:noinspection GoMixedReceiverTypes
func (i SocialType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i SocialType) Tag() Tag {
	return TagSocial
}

const (
	TypeSocialPost SocialType = iota + 1
	TypeSocialRevise
	TypeSocialComment
	TypeSocialShare
	TypeSocialProfile
	TypeSocialMint
	TypeSocialReward
)
