package arweave

import (
	"fmt"

	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/node/v2/provider/arweave"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 12

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network     network.Network
	Block       arweave.Block
	Transaction arweave.Transaction
}

func (t Task) ID() string {
	return fmt.Sprintf("%s-%s", t.Network, t.Transaction.ID)
}

func (t Task) GetNetwork() network.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return uint64(t.Block.Timestamp)
}

func (t Task) Validate() error {
	return nil
}

// BuildActivity builds an activity  from the task.
func (t Task) BuildActivity(options ...activityx.Option) (*activityx.Activity, error) {
	var feeValue decimal.Decimal

	// Set fee value if the reward is not empty.
	if t.Transaction.Reward != "" {
		var err error

		feeValue, err = decimal.NewFromString(t.Transaction.Reward)
		if err != nil {
			return nil, fmt.Errorf("parse transaction reward: %w", err)
		}
	} else {
		feeValue = decimal.Zero
	}

	// From address is the owner of the transaction.
	from, err := arweave.PublicKeyToAddress(t.Transaction.Owner)
	if err != nil {
		return nil, fmt.Errorf("parse transaction owner: %w", err)
	}

	activity := activityx.Activity{
		ID:      t.Transaction.ID,
		Network: t.Network,
		From:    from,
		To:      t.Transaction.Target,
		Type:    typex.Unknown,
		Status:  true,
		Fee: &activityx.Fee{
			Amount:  feeValue,
			Decimal: defaultFeeDecimal,
		},
		Actions:   make([]*activityx.Action, 0),
		Timestamp: uint64(t.Block.Timestamp),
	}

	for _, option := range options {
		if err := option(&activity); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &activity, nil
}
