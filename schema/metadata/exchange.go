package metadata

import (
	"time"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
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
