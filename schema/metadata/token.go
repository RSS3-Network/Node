package metadata

import (
	"github.com/rss3-network/serving-node/provider/ethereum/contract"
	"github.com/shopspring/decimal"
)

type Token struct {
	Address  *string           `json:"address,omitempty"`
	ID       *decimal.Decimal  `json:"id,omitempty"`
	Value    *decimal.Decimal  `json:"value,omitempty"`
	Name     string            `json:"name,omitempty"`
	Symbol   string            `json:"symbol,omitempty"`
	URI      string            `json:"uri,omitempty"`
	Decimals uint8             `json:"decimals,omitempty"`
	Standard contract.Standard `json:"standard,omitempty"`
}
