package metadata

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ schema.Metadata = (*CollectibleTransfer)(nil)

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

var _ schema.Metadata = (*CollectibleApproval)(nil)

type CollectibleApproval struct {
	Action CollectibleApprovalAction `json:"action"`

	Token
}

func (c CollectibleApproval) Type() filter.Type {
	return filter.TypeCollectibleApproval
}
