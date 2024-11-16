package near

import (
	"fmt"
	"math/big"
	"time"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/utils"
	"github.com/rss3-network/node/provider/near"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/shopspring/decimal"
)

const defaultFeeDecimal = 24

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network     network.Network
	Block       near.Block
	Transaction near.Transaction
}

func (t Task) ID() string {
	return fmt.Sprintf("%s-%s", t.Network, t.Transaction.Transaction.Hash)
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
	// From address is the owner of the transaction.
	from := t.Transaction.Transaction.SignerID

	feeAmount, err := t.buildFee()
	if err != nil {
		return nil, fmt.Errorf("build fee: %w", err)
	}

	// Build Calldata
	var calldata *activityx.Calldata

	if len(t.Transaction.Transaction.Actions) > 0 {
		action := t.Transaction.Transaction.Actions[0]

		var parsedFunction string

		switch {
		case action.FunctionCall != nil:
			parsedFunction = action.FunctionCall.MethodName
		case action.Transfer != nil:
			parsedFunction = "Transfer"
		case action.CreateAccount != nil:
			parsedFunction = "CreateAccount"
		case action.DeployContract != nil:
			parsedFunction = "DeployContract"
		case action.Stake != nil:
			parsedFunction = "Stake"
		case action.AddKey != nil:
			parsedFunction = "AddKey"
		case action.DeleteKey != nil:
			parsedFunction = "DeleteKey"
		case action.DeleteAccount != nil:
			parsedFunction = "DeleteAccount"
		case action.DelegateActions != nil:
			parsedFunction = "DelegateActions"
		default:
			parsedFunction = "Unknown"
		}

		calldata = &activityx.Calldata{
			ParsedFunction: parsedFunction,
		}
	}

	activity := activityx.Activity{
		ID:      t.Transaction.Transaction.Hash,
		Network: t.Network,
		From:    from,
		To:      t.Transaction.Transaction.ReceiverID,
		Type:    typex.Unknown,
		Status:  true,
		Fee: &activityx.Fee{
			Amount:  decimal.NewFromBigInt(utils.GetBigInt(feeAmount), 0),
			Decimal: defaultFeeDecimal,
		},
		Actions:   make([]*activityx.Action, 0),
		Calldata:  calldata,
		Timestamp: uint64(time.Duration(t.Block.Header.Timestamp).Seconds()),
	}

	for _, option := range options {
		if err := option(&activity); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &activity, nil
}

func (t Task) buildFee() (*big.Int, error) {
	gasPrice, ok := new(big.Int).SetString(t.Block.Header.GasPrice, 10)
	if !ok {
		return nil, fmt.Errorf("invalid gas price: %s", t.Block.Header.GasPrice)
	}

	gasBurnt := new(big.Int).SetUint64(uint64(t.Transaction.TransactionOutcome.Outcome.GasBurnt))

	return new(big.Int).Mul(gasPrice, gasBurnt), nil
}
