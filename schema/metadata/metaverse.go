package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=MetaverseTradeAction --transform=snake --trimprefix=ActionMetaverseTrade --output metaverse_trade.go --json --sql
type MetaverseTradeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t MetaverseTradeAction) Type() filter.Type {
	return filter.TypeMetaverseTrade
}

const (
	ActionMetaverseTradeSell MetaverseTradeAction = iota + 1
	ActionMetaverseTradeBuy
	ActionMetaverseTradeList
)

var _ Metadata = (*MetaverseTransfer)(nil)

type MetaverseTransfer Token

func (m MetaverseTransfer) Tag() filter.Tag {
	return filter.TagMetaverse
}

func (m MetaverseTransfer) Type() filter.Type {
	return filter.TypeMetaverseTransfer
}

var _ Metadata = (*MetaverseTrade)(nil)

type MetaverseTrade struct {
	Action MetaverseTradeAction `json:"action,omitempty"`

	Token
	Cost Token `json:"cost,omitempty"`
}

func (m MetaverseTrade) Tag() filter.Tag {
	return filter.TagMetaverse
}

func (m MetaverseTrade) Type() filter.Type {
	return filter.TypeMetaverseTrade
}
