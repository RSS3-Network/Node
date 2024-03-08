package highlight_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/highlight"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Ethereum(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Highlight Mint",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x44c76c805bf21b13b29a7acbd0061339a143957cb3a34bfc3c47d95872d07c2a"),
						ParentHash:   common.HexToHash("0xa55e8d62a5a5a11830fb39807272f9b2ac2ac5bb474383cd145fcdc767639a63"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("18416219", 0)),
						GasLimit:     30000000,
						GasUsed:      18311702,
						Timestamp:    1698103355,
						BaseFee:      lo.Must(new(big.Int).SetString("52226673329", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
						From:      common.HexToAddress("0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF"),
						Gas:       189366,
						GasPrice:  lo.Must(new(big.Int).SetString("53226673329", 10)),
						Hash:      common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
						Input:     hexutil.MustDecode("0x705dcf3400000000000000000000000000000000000000000000000000000000000000180000000000000000000000000000000000000000000000000000000000000003000000000000000000000000ea310c966d3ff5e09c65487f1763b21361eb71ef"),
						To:        lo.ToPtr(common.HexToAddress("0x1bf979282181f2b7a640d17aB5D2e25125F2de5e")),
						Value:     lo.Must(new(big.Int).SetString("26400000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x44c76c805bf21b13b29a7acbd0061339a143957cb3a34bfc3c47d95872d07c2a"),
						BlockNumber:       lo.Must(new(big.Int).SetString("18416219", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8928155,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xc648e94b1"),
						GasUsed:           162637,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x1bf979282181f2b7a640d17aB5D2e25125F2de5e"),
							Topics: []common.Hash{
								common.HexToHash("0x9363885e28e7ba67b096932f9f00dff44742731d6cb4fa26ccd4424e78e41e13"),
								common.HexToHash("0x000000000000000000000000a532df3bedc7a30747c45f57603080f37a60df48"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000018"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000005543df729c00000000000000000000000000000000000000000000000000000000000000002710"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18416219", 0)),
							TransactionHash: common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
							Index:           184,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1bf979282181f2b7a640d17aB5D2e25125F2de5e"),
							Topics: []common.Hash{
								common.HexToHash("0x981414aed4973b05aa301314dc13a5a4077f24490497b98bc270852581c1c578"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000018"),
								common.HexToHash("0x000000000000000000000000fc19b4536a4b9b9eec1c362894f7dc505457538d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000003"),
							BlockNumber:     lo.Must(new(big.Int).SetString("18416219", 0)),
							TransactionHash: common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
							Index:           185,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xFC19B4536a4b9B9EEc1C362894f7dC505457538D"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000ea310c966d3ff5e09c65487f1763b21361eb71ef"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000002b8"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18416219", 0)),
							TransactionHash: common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
							Index:           186,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xFC19B4536a4b9B9EEc1C362894f7dC505457538D"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000ea310c966d3ff5e09c65487f1763b21361eb71ef"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000002b9"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18416219", 0)),
							TransactionHash: common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
							Index:           187,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xFC19B4536a4b9B9EEc1C362894f7dC505457538D"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000ea310c966d3ff5e09c65487f1763b21361eb71ef"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000002ba"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("18416219", 0)),
							TransactionHash: common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
							Index:           188,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42"),
						TransactionIndex: 79,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xe278e133786ec1714550361994ad314e6219196d8500eab67cd2df5000898c42",
				Network: filter.NetworkEthereum,
				Index:   79,
				From:    "0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF",
				To:      "0x1bf979282181f2b7a640d17aB5D2e25125F2de5e",
				Type:    filter.TypeCollectibleMint,
				Calldata: &schema.Calldata{
					FunctionHash: "0x705dcf34",
				},
				Platform: lo.ToPtr(filter.PlatformHighlight),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("8656626470208573")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformHighlight.String(),
						From:     common.HexToAddress("0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF").String(),
						To:       common.HexToAddress("0xa532Df3Bedc7A30747c45F57603080F37A60df48").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("24000000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformHighlight.String(),
						From:     common.HexToAddress("0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF").String(),
						To:       common.HexToAddress("0x1bf979282181f2b7a640d17aB5D2e25125F2de5e").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2400000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformHighlight.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       common.HexToAddress("0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF").String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0xFC19B4536a4b9B9EEc1C362894f7dC505457538D"),
							ID:       lo.ToPtr(decimal.NewFromBigInt(big.NewInt(696), 0)),
							Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
							Name:     "R3FUS3",
							Symbol:   "R3FUS3",
							Standard: metadata.StandardERC721,
							URI:      "https://highlight-creator-assets.highlight.xyz/main/base-dir/b0b61e2c-f741-48c4-94ca-f090c9f5817e/onChainDir/696",
						},
					},
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformHighlight.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       common.HexToAddress("0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF").String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0xFC19B4536a4b9B9EEc1C362894f7dC505457538D"),
							ID:       lo.ToPtr(decimal.NewFromBigInt(big.NewInt(697), 0)),
							Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
							Name:     "R3FUS3",
							Symbol:   "R3FUS3",
							Standard: metadata.StandardERC721,
							URI:      "https://highlight-creator-assets.highlight.xyz/main/base-dir/b0b61e2c-f741-48c4-94ca-f090c9f5817e/onChainDir/697",
						},
					},
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformHighlight.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       common.HexToAddress("0xeA310C966D3Ff5E09c65487f1763B21361Eb71eF").String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0xFC19B4536a4b9B9EEc1C362894f7dC505457538D"),
							ID:       lo.ToPtr(decimal.NewFromBigInt(big.NewInt(698), 0)),
							Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
							Name:     "R3FUS3",
							Symbol:   "R3FUS3",
							Standard: metadata.StandardERC721,
							URI:      "https://highlight-creator-assets.highlight.xyz/main/base-dir/b0b61e2c-f741-48c4-94ca-f090c9f5817e/onChainDir/698",
						},
					},
				},
				Status:    true,
				Timestamp: 1698103355,
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

			//t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}
