package metadata

import (
	"github.com/rss3-network/serving-node/schema/filter"
)

var _ Metadata = (*TransactionTransfer)(nil)

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

var _ Metadata = (*TransactionApproval)(nil)

type TransactionApproval struct {
	Action TransactionApprovalAction `json:"action"`

	Token
}

func (t TransactionApproval) Type() filter.Type {
	return filter.TypeTransactionApproval
}

var _ Metadata = (*TransactionBridge)(nil)

type TransactionBridge struct {
	Action        TransactionBridgeAction `json:"action"`
	SourceNetwork filter.Network          `json:"source_network"`
	TargetNetwork filter.Network          `json:"target_network"`
	Token         Token                   `json:"token"`
}

func (t TransactionBridge) Type() filter.Type {
	return filter.TypeTransactionBridge
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=TransactionBridgeAction --transform=snake --trimprefix=ActionTransactionBridge --output transaction_bridge.go --json --sql
type TransactionBridgeAction uint64

//goland:noinspection GoMixedReceiverTypes
func (t TransactionBridgeAction) Type() filter.Type {
	return filter.TypeTransactionBridge
}

const (
	ActionTransactionBridgeWithdraw TransactionBridgeAction = iota + 1
	ActionTransactionBridgeDeposit
)
