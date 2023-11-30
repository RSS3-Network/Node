package arweave

import (
	"fmt"

	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 18

var _ engine.Task = (*Task)(nil)

type Task struct {
	Chain       filter.ChainArweave
	Block       types.Block
	Transaction types.Transaction
}

func (t Task) ID() string {
	return fmt.Sprintf("%s-%s-%s", t.Chain.Network(), t.Chain.Name(), t.Transaction.ID)
}

func (t Task) Network() filter.Network {
	return t.Chain.Network()
}

func (t Task) Timestamp() uint64 {
	return uint64(t.Block.Timestamp)
}

func (t Task) Validate() error {
	return nil
}

func (t Task) BuildFeed( /* TODO Implementing options. */ ) (*schema.Feed, error) {
	var feeValue decimal.Decimal

	if t.Transaction.Reward != "" {
		var err error

		feeValue, err = decimal.NewFromString(t.Transaction.Reward)
		if err != nil {
			return nil, fmt.Errorf("parse transaction reward: %w", err)
		}

		feeValue = feeValue.Shift(-12)
	} else {
		feeValue = decimal.Zero
	}

	from, err := utils.OwnerToAddress(t.Transaction.Owner)
	if err != nil {
		return nil, fmt.Errorf("parse transaction owner: %w", err)
	}

	feed := schema.Feed{
		ID:     t.Transaction.ID,
		Chain:  t.Chain,
		From:   from,
		To:     t.Transaction.Target,
		Type:   filter.TypeUnknown,
		Status: true,
		Fee: schema.Fee{
			Amount:  feeValue,
			Decimal: defaultFeeDecimal,
		},
		Actions:   make([]*schema.Action, 0),
		Timestamp: uint64(t.Block.Timestamp),
	}

	return &feed, nil
}
