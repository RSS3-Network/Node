package lens

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func EncodeID(id *big.Int) string {
	value := strconv.FormatInt(id.Int64(), 16)

	return fmt.Sprintf("0x%s%s", strings.Repeat("0", len(value)%2), value)
}
