package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ schema.Metadata = (*TransactionTransfer)(nil)

type TransactionTransfer Token

func (t TransactionTransfer) Type() filter.Type {
	return filter.TypeTransactionTransfer
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=TransactionApprovalAction --transform=snake --trimprefix=ActionTransaction --output transaction_approval.go --json --sql
type TransactionApprovalAction uint64

const (
	ActionTransactionApprove TransactionApprovalAction = iota + 1
	ActionTransactionRevoke
)

var _ schema.Metadata = (*TransactionApproval)(nil)

type TransactionApproval struct {
	Action TransactionApprovalAction `json:"action"`

	Token
}

func (t TransactionApproval) Type() filter.Type {
	return filter.TypeTransactionApproval
}
