package filter

import "fmt"

type Type interface {
	Name() string
	Tag() Tag
}

func TypeString(tag Tag, typeValue string) (Type, error) {
	switch tag {
	case TagCollectible:
		return CollectibleTypeString(typeValue)
	case TagExchange:
		return ExchangeTypeString(typeValue)
	case TagMetaverse:
		return MetaverseTypeString(typeValue)
	case TagRSS:
		return RSSTypeString(typeValue)
	case TagSocial:
		return SocialTypeString(typeValue)
	case TagTransaction:
		return TransactionTypeString(typeValue)
	case TagUnknown:
		return UnknownTypeString(typeValue)
	default:
		return nil, fmt.Errorf("unknown tag: %s", tag.String())
	}
}
