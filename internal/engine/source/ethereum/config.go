package ethereum

import (
	"math/big"
)

type Config struct { // TODO Implement support for customizable configurations.
	BlockNumberStart  *big.Int
	BlockNumberTarget *big.Int
}
