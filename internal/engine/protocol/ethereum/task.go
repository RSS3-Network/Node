package ethereum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/provider/ethereum"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 18

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network     network.Network
	ChainID     uint64
	Header      *ethereum.Header
	Transaction *ethereum.Transaction
	Receipt     *ethereum.Receipt
}

func (t Task) ID() string {
	return fmt.Sprintf("%s.%s", t.Network, t.Transaction.Hash)
}

func (t Task) GetNetwork() network.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return t.Header.Timestamp
}

func (t Task) Validate() error {
	// TODO Implement it.
	return nil
}

func (t Task) BuildActivity(options ...activityx.Option) (*activityx.Activity, error) {
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

	activity := activityx.Activity{
		ID:      t.Transaction.Hash.String(),
		Network: t.Network,
		Index:   t.Receipt.TransactionIndex,
		From:    from.String(),
		To:      to.String(),
		Type:    typex.Unknown,
		Fee: &activityx.Fee{
			Amount:  decimal.NewFromBigInt(utils.GetBigInt(feeAmount), 0),
			Decimal: defaultFeeDecimal,
		},
		Calldata: &activityx.Calldata{
			FunctionHash: functionHash,
		},
		Actions:   make([]*activityx.Action, 0),
		Status:    t.Receipt.Status == types.ReceiptStatusSuccessful,
		Timestamp: t.Header.Timestamp,
	}

	// Apply activity options.
	for _, option := range options {
		if err := option(&activity); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &activity, nil
}

func (t Task) buildFee() (*big.Int, error) {
	switch {
	case network.IsOptimismSuperchain(t.ChainID):
		return t.buildFeeOptimismSuperchain()
	case t.Network == network.Arbitrum:
		return t.buildFeeArbitrumNitro()
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
			gasPrice = decimal.NewFromBigInt(utils.GetBigInt(t.Receipt.EffectiveGasPrice), 0)
		} else {
			var (
				baseFee   = decimal.NewFromBigInt(utils.GetBigInt(t.Header.BaseFee), 0)
				gasTipCap = decimal.NewFromBigInt(utils.GetBigInt(t.Transaction.GasTipCap), 0)
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
	case ethereum.TransactionTypeOptimismDeposit: // The `effectiveGasPrice` of the transaction is always 0.
		return big.NewInt(0), nil
	default:
		return nil, fmt.Errorf("unsupported transaction type %d", t.Transaction.Type)
	}
}

// buildFeeArbitrumNitro calculates the fee of the Arbitrum Nitro transactions.
func (t Task) buildFeeArbitrumNitro() (*big.Int, error) {
	switch t.Transaction.Type {
	case
		types.LegacyTxType,
		types.AccessListTxType,
		types.DynamicFeeTxType:
		return t.buildFeeDefault()
	case // The transaction fee is `effectiveGasPrice` * `gasUsed`.
		ethereum.TransactionTypeArbitrumContract,
		ethereum.TransactionTypeArbitrumUnsigned,
		ethereum.TransactionTypeArbitrumRetry,
		ethereum.TransactionTypeArbitrumLegacy:
		return new(big.Int).Mul(t.Receipt.EffectiveGasPrice, new(big.Int).SetUint64(t.Receipt.GasUsed)), nil
	case // The `gasUsed` of the transaction is always 0.
		ethereum.TransactionTypeArbitrumSubmitRetryable,
		ethereum.TransactionTypeArbitrumDeposit,
		ethereum.TransactionTypeArbitrumInternal:
		return big.NewInt(0), nil
	default:
		return nil, fmt.Errorf("unsupported transaction type %d", t.Transaction.Type)
	}
}
