package arweave

import (
	"math/big"
)

type Config struct {
	BlockNumberStart  *big.Int
	BlockNumberTarget *big.Int
}
