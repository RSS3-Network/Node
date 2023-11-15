package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/shopspring/decimal"
)

var _ schema.Metadata = (*TransactionTransfer)(nil)

type TransactionTransfer struct {
	Address *string          `json:"address,omitempty"`
	ID      *decimal.Decimal `json:"id,omitempty"`
	Value   *decimal.Decimal `json:"value,omitempty"`
}

func (t TransactionTransfer) Type() filter.Type {
	return filter.TypeTransactionTransfer
}
