package utils

import (
	"fmt"
	"math/big"

	"github.com/rss3-network/protocol-go/schema"
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
