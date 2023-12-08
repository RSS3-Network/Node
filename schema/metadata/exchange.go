package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ schema.Metadata = (*ExchangeSwap)(nil)

type ExchangeSwap struct {
	From Token `json:"from"`
	To   Token `json:"to"`
}

func (e ExchangeSwap) Type() filter.Type {
	return filter.TypeExchangeSwap
}

var _ schema.Metadata = (*ExchangeLiquidity)(nil)

type ExchangeLiquidity struct {
	Action ExchangeLiquidityAction `json:"action"`
	Tokens []Token                 `json:"tokens"`
}

func (e ExchangeLiquidity) Type() filter.Type {
	return filter.TypeExchangeSwap
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeLiquidityAction --transform=snake --trimprefix=ActionExchangeLiquidity --output exchange_liquidity.go --json --sql
type ExchangeLiquidityAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeLiquidityAction) Type() filter.Type {
	return filter.TypeExchangeLiquidity
}

const (
	ActionExchangeLiquidityAdd ExchangeLiquidityAction = iota + 1
	ActionExchangeLiquidityRemove
	ActionExchangeLiquidityCollect
)
