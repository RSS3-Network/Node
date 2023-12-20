package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ Metadata = (*CollectibleTransfer)(nil)

type CollectibleTransfer Token

func (c CollectibleTransfer) Type() filter.Type {
	return filter.TypeCollectibleTransfer
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleApprovalAction --transform=snake --trimprefix=ActionCollectibleApproval --output collectible_approval.go --json --sql
type CollectibleApprovalAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t CollectibleApprovalAction) Type() filter.Type {
	return filter.TypeCollectibleApproval
}

const (
	ActionCollectibleApprovalApprove CollectibleApprovalAction = iota + 1
	ActionCollectibleApprovalRevoke
)

var _ Metadata = (*CollectibleApproval)(nil)

type CollectibleApproval struct {
	Action CollectibleApprovalAction `json:"action"`

	Token
}

func (c CollectibleApproval) Type() filter.Type {
	return filter.TypeCollectibleApproval
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=CollectibleTradeAction --transform=snake --trimprefix=ActionCollectibleTrade --output collectible_trade.go --json --sql
type CollectibleTradeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (r CollectibleTradeAction) Type() filter.Type {
	return filter.TypeCollectibleTrade
}

const (
	ActionCollectibleTradeBuy CollectibleTradeAction = iota + 1
	ActionCollectibleTradeSell
)

var _ Metadata = (*CollectibleTradeAction)(nil)

type CollectibleTrade struct {
	Action CollectibleTradeAction `json:"action"`
	Token
	Cost *Token `json:"cost,omitempty"`
}

func (r CollectibleTrade) Type() filter.Type {
	return filter.TypeCollectibleTrade
}
