package arweave

import (
	"math/big"
)

type Option struct {
	BlockHeightStart  *big.Int
	BlockHeightTarget *big.Int
}
