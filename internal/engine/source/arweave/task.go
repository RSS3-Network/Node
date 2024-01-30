package arweave

import (
	"fmt"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 12

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network     filter.Network
	Block       arweave.Block
	Transaction arweave.Transaction
}

func (t Task) ID() string {
	return fmt.Sprintf("%s-%s", t.Network, t.Transaction.ID)
}

func (t Task) GetNetwork() filter.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return uint64(t.Block.Timestamp)
}

func (t Task) Validate() error {
	return nil
}

// BuildFeed builds a feed from the task.
func (t Task) BuildFeed(options ...schema.FeedOption) (*schema.Feed, error) {
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

	feed := schema.Feed{
		ID:      t.Transaction.ID,
		Network: t.Network,
		From:    from,
		To:      t.Transaction.Target,
		Type:    filter.TypeUnknown,
		Status:  true,
		Fee: &schema.Fee{
			Amount:  feeValue,
			Decimal: defaultFeeDecimal,
		},
		Actions:   make([]*schema.Action, 0),
		Timestamp: uint64(t.Block.Timestamp),
	}

	for _, option := range options {
		if err := option(&feed); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &feed, nil
}
