package filter

import "fmt"

type Type interface {
	Name() string
	Tag() Tag
}

func TypeString(tag Tag, typeValue string) (Type, error) {
	switch tag {
	case TagTransaction:
		return TransactionTypeString(typeValue)
	case TagCollectible:
		return CollectibleTypeString(typeValue)
	case TagUnknown:
		return UnknownTypeString(typeValue)
	default:
		return nil, fmt.Errorf("unknown tag: %s", tag.String())
	}
}
