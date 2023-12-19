package metadata

import (
	"encoding/json"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Metadata interface {
	Type() filter.Type
}

func Unmarshal(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	switch metadataType.Tag() {
	case filter.TagCollectible:
		return unmarshalCollectibleMetadata(metadataType, data)
	case filter.TagTransaction:
		return unmarshalTransactionMetadata(metadataType, data)
	case filter.TagSocial:
		return unmarshalSocialMetadata(metadataType, data)
	case filter.TagRSS:
		return unmarshalRSSMetadata(metadataType, data)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}
}

func unmarshalCollectibleMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeCollectibleTransfer, filter.TypeCollectibleMint, filter.TypeCollectibleBurn:
		result = new(CollectibleTransfer)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalTransactionMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeTransactionApproval:
		result = new(TransactionApproval)
	case filter.TypeTransactionTransfer:
		result = new(TransactionTransfer)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalSocialMetadata(metadataType filter.Type, data json.RawMessage) (Metadata, error) {
	var result Metadata

	switch metadataType {
	case filter.TypeSocialPost, filter.TypeSocialComment, filter.TypeSocialShare:
		result = new(SocialPost)
	default:
		return nil, fmt.Errorf("invalid metadata type: %s.%s", metadataType.Tag(), metadataType.Name())
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}

func unmarshalRSSMetadata(_ filter.Type, data json.RawMessage) (Metadata, error) {
	result := new(RSS)

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid metadata: %w", err)
	}

	return result, nil
}
