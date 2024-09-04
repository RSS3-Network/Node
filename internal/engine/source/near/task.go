package near

import (
	"fmt"
	"time"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/near"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 12

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network     network.Network
	Block       near.Block
	Transaction near.TransactionDetails
}

func (t Task) ID() string {
	return fmt.Sprintf("%s-%s", t.Network, t.Transaction.Hash)
}

func (t Task) GetNetwork() network.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	// Convert nanoseconds to seconds using time package
	return uint64(time.Duration(t.Block.Header.Timestamp).Seconds())
}

func (t Task) Validate() error {
	return nil
}

// BuildActivity builds an activity from the task.
func (t Task) BuildActivity(options ...activityx.Option) (*activityx.Activity, error) {
	var feeValue decimal.Decimal

	// Set fee value if the reward is not empty.
	if t.Transaction.PriorityFee != 0 {
		var err error

		feeValue, err = decimal.NewFromString(fmt.Sprintf("%d", t.Transaction.PriorityFee))
		if err != nil {
			return nil, fmt.Errorf("parse transaction priority fee: %w", err)
		}
	} else {
		feeValue = decimal.Zero
	}

	// From address is the owner of the transaction.
	from := t.Transaction.SignerID

	activity := activityx.Activity{
		ID:      t.Transaction.Hash,
		Network: t.Network,
		From:    from,
		To:      t.Transaction.ReceiverID,
		Type:    typex.Unknown,
		Status:  true,
		Fee: &activityx.Fee{
			Amount:  feeValue,
			Decimal: defaultFeeDecimal,
		},
		Actions:   make([]*activityx.Action, 0),
		Timestamp: uint64(time.Duration(t.Block.Header.Timestamp).Seconds()),
	}

	for _, option := range options {
		if err := option(&activity); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &activity, nil
}
