package ethereum_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Ethereum(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *ethereum.Task
		config *engine.Config
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Ethereum native transfer",
			arguments: arguments{
				task: &ethereum.Task{
					Chain: filter.ChainEthereumMainnet,
					Header: &types.Header{
						// TODO Provide all fields.
						Time: 1647774927,
					},
					Transaction: types.NewTx(&types.DynamicFeeTx{
						ChainID:   lo.Must(hexutil.DecodeBig("0x1")),
						Nonce:     0xa,
						GasTipCap: lo.Must(hexutil.DecodeBig("0x59682f00")),
						GasFeeCap: lo.Must(hexutil.DecodeBig("0x3b8c8f46b")),
						Gas:       0x5208,
						To:        lo.ToPtr(common.HexToAddress("0xa1b2dcac834117f38fb0356b5176b5693e165c90")),
						Value:     lo.Must(hexutil.DecodeBig("0xd48ed9972b634")),
						V:         lo.Must(hexutil.DecodeBig("0x0")),
						R:         lo.Must(hexutil.DecodeBig("0x66f2c1c5fbde05362a8d944ee884ba05777150c5dbe5bd414834b567a73f7765")),
						S:         lo.Must(hexutil.DecodeBig("0x4b91fdf20d7d85572836d5c4849949d4d96b6cff36e6a53f2b1dce6a3144ef4")),
					}),
					TransactionIndex: 0xf4,
					Receipt: &types.Receipt{
						Type:              0x2,
						PostState:         common.Hex2Bytes("0x1"),
						Status:            types.ReceiptStatusSuccessful,
						CumulativeGasUsed: 0x18b21de,
						TxHash:            common.HexToHash("0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99"),
						GasUsed:           0x5208,
						EffectiveGasPrice: lo.Must(hexutil.DecodeBig("0x3b8c8f46b")),
						BlockHash:         common.HexToHash("0xea9d0ecd7a085aa998789e8e9c017a7d45f199873380ecb568218525171165b0"),
						BlockNumber:       lo.Must(hexutil.DecodeBig("0xdc1390")),
						TransactionIndex:  0xf4,
					},
				},
			},
			want: &schema.Feed{
				ID:    "0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99",
				Chain: filter.ChainEthereumMainnet,
				Index: 244,
				From:  "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
				To:    "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
				Type:  filter.TypeTransactionTransfer,
				Fee: schema.Fee{
					Amount:  decimal.NewFromInt(0),
					Decimal: 18,
				},
				Actions: []schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
						To:   "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
						Metadata: metadata.TransactionTransfer{
							Value: lo.ToPtr(lo.Must(decimal.NewFromString("3739360016119348"))),
						},
					},
				},
				Status:    true,
				Timestamp: 1647774927,
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config)
			require.NoError(t, err)

			matched, err := instance.Match(ctx, testcase.arguments.task)
			testcase.wantError(t, err)
			require.True(t, matched)

			feed, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)
			require.Equal(t, testcase.want, feed)

			t.Log(feed)
		})
	}
}
