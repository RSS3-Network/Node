package utils

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/tag"
)

// GetBigInt returns the value if not nil, otherwise returns big.NewInt(0)
func GetBigInt(value *big.Int) *big.Int {
	if value == nil {
		return big.NewInt(0)
	}

	return value
}

// ParseTypes parses the types and returns the schema types
func ParseTypes(types []string, tags []tag.Tag) ([]schema.Type, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	schemaTypes := make([]schema.Type, 0)

	for _, typex := range types {
		var (
			value schema.Type
			err   error
		)

		for _, tagx := range tags {
			value, err = schema.ParseTypeFromString(tagx, typex)
			if err == nil {
				schemaTypes = append(schemaTypes, value)

				break
			}
		}

		if err != nil {
			return nil, fmt.Errorf("invalid type: %s", typex)
		}
	}

	return schemaTypes, nil
}

// ParseMetadata parses the metadata and returns the schema metadata
func ParseMetadata(metadataJSON json.RawMessage, typex schema.Type) (metadata.Metadata, error) {
	if len(metadataJSON) == 0 {
		return nil, nil
	}

	value, err := metadata.Unmarshal(typex, metadataJSON)
	if err != nil {
		return nil, err
	}

	return value, nil
}
