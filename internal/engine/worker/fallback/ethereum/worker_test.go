package ethereum_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/fallback/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/endpoint"
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
		task   *source.Task
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
				task: &source.Task{
					Chain: filter.ChainEthereumMainnet,
					Header: &ethereum.Header{
						// TODO Provide all fields.
						Timestamp: 1647774927,
					},
					Transaction: &ethereum.Transaction{
						BlockHash:   common.HexToHash("0xea9d0ecd7a085aa998789e8e9c017a7d45f199873380ecb568218525171165b0"),
						BlockNumber: lo.Must(hexutil.DecodeBig("0xdc1390")),
						From:        common.HexToAddress("0x000000A52a03835517E9d193B3c27626e1Bc96b1"),
						Gas:         0x5208,
						GasPrice:    nil,
						Hash:        common.HexToHash("0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99"),
						Input:       nil,
						To:          lo.ToPtr(common.HexToAddress("0xa1b2dcac834117f38fb0356b5176b5693e165c90")),
						Index:       0xf4,
						Value:       lo.Must(hexutil.DecodeBig("0xd48ed9972b634")),
						Type:        types.DynamicFeeTxType,
						ChainID:     lo.Must(hexutil.DecodeBig("0x1")),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xea9d0ecd7a085aa998789e8e9c017a7d45f199873380ecb568218525171165b0"),
						BlockNumber:       lo.Must(hexutil.DecodeBig("0xdc1390")),
						ContractAddress:   nil,
						CumulativeGasUsed: 0x18b21de,
						EffectiveGasPrice: lo.Must(hexutil.DecodeBig("0x3b8c8f46b")),
						GasUsed:           0x5208,
						Logs:              nil,
						Status:            types.ReceiptStatusSuccessful,
						TransactionHash:   common.HexToHash("0x0c2f413efbc243f3bb8edac7e70bdc21936e01401a21b0d63e97732aa80f5d99"),
						TransactionIndex:  0xf4,
					},
				},
				config: &engine.Config{
					Chain:    filter.ChainEthereumMainnet,
					Endpoint: endpoint.MustGet(filter.ChainEthereumMainnet),
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
					Amount:  lo.Must(decimal.NewFromString("335686667463000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "0x000000A52a03835517E9d193B3c27626e1Bc96b1",
						To:   "0xA1b2DCAC834117F38FB0356b5176B5693E165c90",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("3739360016119348"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
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
