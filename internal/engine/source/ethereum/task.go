package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 18

var _ engine.Task = (*Task)(nil)

type Task struct {
	Chain       filter.ChainEthereum
	Header      *types.Header
	Transaction *types.Transaction
	Receipt     *types.Receipt
}

func (t Task) ID() string {
	return fmt.Sprintf("%s.%s.%s", t.Chain.Network(), t.Chain, t.Transaction.Hash())
}

func (t Task) Network() filter.Network {
	return t.Chain.Network()
}

func (t Task) Timestamp() uint64 {
	return t.Header.Time
}

func (t Task) Validate() error {
	// TODO Implement it.
	return nil
}

func (t Task) BuildFeed( /* TODO Implementing options. */ ) (*schema.Feed, error) {
	var to common.Address

	if t.Transaction.To() == nil {
		// The Nethermind node may be missing this field,
		// Reference https://github.com/NethermindEth/nethermind/issues/5823.
		if t.Receipt.ContractAddress == ethereum.AddressGenesis {
			return nil, fmt.Errorf("invalid receipt %s", t.Transaction.Hash().String())
		}

		to = t.Receipt.ContractAddress
	} else {
		to = *t.Transaction.To()
	}

	from, err := types.LatestSignerForChainID(t.Transaction.ChainId()).Sender(t.Transaction)
	if err != nil {
		return nil, fmt.Errorf("recovery transaction sender: %w", err)
	}

	feeAmount, err := t.buildFee()
	if err != nil {
		return nil, fmt.Errorf("build fee: %w", err)
	}

	feed := schema.Feed{
		ID:    t.Transaction.Hash().String(),
		Chain: t.Chain,
		Index: t.Receipt.TransactionIndex,
		From:  from.String(),
		To:    to.String(),
		Type:  filter.TypeUnknown,
		Fee: schema.Fee{
			Amount:  decimal.NewFromBigInt(feeAmount, 0),
			Decimal: defaultFeeDecimal,
		},
		Actions:   make([]*schema.Action, 0),
		Status:    t.Receipt.Status == types.ReceiptStatusSuccessful,
		Timestamp: t.Header.Time,
	}

	return &feed, nil
}

func (t Task) buildFee() (*big.Int, error) {
	switch t.Receipt.Type {
	case types.LegacyTxType, types.AccessListTxType:
		return new(big.Int).Mul(t.Transaction.GasPrice(), new(big.Int).SetUint64(t.Receipt.GasUsed)), nil
	case types.DynamicFeeTxType:
		// TODO Add support for EIP-1559 transaction.
		return big.NewInt(0), nil
	default:
		return nil, fmt.Errorf("unsupported transaction type %d", t.Receipt.Type)
	}
}
