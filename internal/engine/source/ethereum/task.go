package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 18

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network     filter.Network
	ChainID     uint64
	Header      *ethereum.Header
	Transaction *ethereum.Transaction
	Receipt     *ethereum.Receipt
}

func (t Task) ID() string {
	return fmt.Sprintf("%s.%s", t.Network, t.Transaction.Hash)
}

func (t Task) GetNetwork() filter.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return t.Header.Timestamp
}

func (t Task) Validate() error {
	// TODO Implement it.
	return nil
}

func (t Task) BuildFeed(options ...schema.FeedOption) (*schema.Feed, error) {
	var to, from common.Address

	if t.Transaction.To == nil {
		// The Nethermind node may be missing this field,
		// Reference https://github.com/NethermindEth/nethermind/issues/5823.
		if t.Receipt.ContractAddress == nil {
			return nil, fmt.Errorf("invalid receipt %s", t.Transaction.Hash)
		}

		to = *t.Receipt.ContractAddress
	} else {
		to = *t.Transaction.To
	}

	// TODO Verify transaction signature.
	// There is no signer implementation that supports all of layer 2 network at the same time.
	from = t.Transaction.From

	feeAmount, err := t.buildFee()
	if err != nil {
		return nil, fmt.Errorf("build fee: %w", err)
	}

	var functionHash string

	if t.Transaction.Input != nil && len(t.Transaction.Input) >= 4 {
		functionHashBytes := t.Transaction.Input[:4]
		functionHash = hexutil.Encode(functionHashBytes)
	}

	feed := schema.Feed{
		ID:      t.Transaction.Hash.String(),
		Network: t.Network,
		Index:   t.Receipt.TransactionIndex,
		From:    from.String(),
		To:      to.String(),
		Type:    filter.TypeUnknown,
		Fee: &schema.Fee{
			Amount:  decimal.NewFromBigInt(feeAmount, 0),
			Decimal: defaultFeeDecimal,
		},
		Calldata: &schema.Calldata{
			FunctionHash: functionHash,
		},
		Actions:   make([]*schema.Action, 0),
		Status:    t.Receipt.Status == types.ReceiptStatusSuccessful,
		Timestamp: t.Header.Timestamp,
	}

	// Apply feed options.
	for _, option := range options {
		if err := option(&feed); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &feed, nil
}

func (t Task) buildFee() (*big.Int, error) {
	switch {
	case filter.IsOptimismSuperchain(t.ChainID):
		return t.buildFeeOptimismSuperchain()
	default:
		return t.buildFeeDefault()
	}
}

func (t Task) buildFeeDefault() (*big.Int, error) {
	switch t.Transaction.Type {
	case types.LegacyTxType, types.AccessListTxType:
		return new(big.Int).Mul(t.Transaction.GasPrice, new(big.Int).SetUint64(t.Receipt.GasUsed)), nil
	case types.DynamicFeeTxType: // https://eips.ethereum.org/EIPS/eip-1559
		var (
			gasPrice decimal.Decimal
			gasUsed  = decimal.NewFromBigInt(new(big.Int).SetUint64(t.Receipt.GasUsed), 0)
		)

		if t.Receipt.EffectiveGasPrice != nil {
			gasPrice = decimal.NewFromBigInt(t.Receipt.EffectiveGasPrice, 0)
		} else {
			var (
				baseFee   = decimal.NewFromBigInt(t.Header.BaseFee, 0)
				gasTipCap = decimal.NewFromBigInt(t.Transaction.GasTipCap, 0)
			)

			gasPrice = baseFee.Add(gasTipCap)
		}

		return gasPrice.Mul(gasUsed).BigInt(), nil
	default:
		return nil, fmt.Errorf("unsupported transaction type %d", t.Transaction.Type)
	}
}

func (t Task) buildFeeOptimismSuperchain() (*big.Int, error) {
	switch t.Transaction.Type {
	case types.LegacyTxType, types.AccessListTxType, types.DynamicFeeTxType:
		fee, err := t.buildFeeDefault()
		if err != nil {
			return nil, err
		}

		return new(big.Int).Add(fee, t.Receipt.L1Fee), nil
	case ethereum.TransactionTypeOptimismDeposit:
		return big.NewInt(0), nil
	default:
		return nil, fmt.Errorf("unsupported transaction type %d", t.Transaction.Type)
	}
}
