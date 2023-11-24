package schema

import (
	"github.com/shopspring/decimal"
)

type FeeTransformer interface {
	Import(fee Fee) error
}

type Fee struct {
	Address *string         `json:"address,omitempty"`
	Amount  decimal.Decimal `json:"amount"`
	Decimal uint            `json:"decimal"`
}
