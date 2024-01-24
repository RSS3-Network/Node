package metadata

import (
	"time"

	"github.com/rss3-network/serving-node/schema/filter"
)

var _ Metadata = (*ExchangeSwap)(nil)

type ExchangeSwap struct {
	From Token `json:"from"`
	To   Token `json:"to"`
}

func (e ExchangeSwap) Type() filter.Type {
	return filter.TypeExchangeSwap
}

var _ Metadata = (*ExchangeLiquidity)(nil)

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
	ActionExchangeLiquiditySupply
	ActionExchangeLiquidityBorrow
	ActionExchangeLiquidityRepay
	ActionExchangeLiquidityWithdraw
)

var _ Metadata = (*ExchangeStaking)(nil)

type ExchangeStaking struct {
	Action ExchangeStakingAction  `json:"action"`
	Token  Token                  `json:"token"`
	Period *ExchangeStakingPeriod `json:"period,omitempty"`
}

func (e ExchangeStaking) Type() filter.Type {
	return filter.TypeExchangeStaking
}

type ExchangeStakingPeriod struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=ExchangeStakingAction --transform=snake --trimprefix=ActionExchangeStaking --output exchange_staking.go --json --sql
type ExchangeStakingAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t ExchangeStakingAction) Type() filter.Type {
	return filter.TypeExchangeStaking
}

const (
	ActionExchangeStakingStake ExchangeStakingAction = iota + 1
	ActionExchangeStakingUnstake
	ActionExchangeStakingClaim
)
