package utils

import "math/big"

// GetBigInt returns the value if not nil, otherwise returns big.NewInt(0)
func GetBigInt(value *big.Int) *big.Int {
	if value == nil {
		return big.NewInt(0)
	}

	return value
}
