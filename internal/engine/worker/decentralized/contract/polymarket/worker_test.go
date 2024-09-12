package polymarket_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/polymarket"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Polymarket(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *activityx.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Test Prediction Offer Match Finalization (buy, sell actions)",
			arguments: struct {
				task   *source.Task
				config *config.Module
			}{
				task: &source.Task{
					Network: network.Polygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x32422585cf4fdb3977cbb80f53e02211bc88b8bed1f4528907937d7b8db2bb9d"),
						ParentHash:   common.HexToHash("0x2bcbe0471d721245fdf728d54f2c9148ea1997ba0bb7d16a06e918b11c20babb"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4F080de0d0CC72e091a77cD03036750Adc47060e"),
						Number:       lo.Must(new(big.Int).SetString("61724368", 0)),
						GasLimit:     30000000,
						GasUsed:      11394017,
						Timestamp:    1726122849,
						BaseFee:      lo.Must(new(big.Int).SetString("23", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xee69b38fad2d28c86d6fd838ba0ab328844c753b67ceec37e3433aec58d71c1b"),
						From:      common.HexToAddress("0x429B8474BD7308b7787d364985bB4b8eA7De1D47"),
						Gas:       1200000,
						GasPrice:  lo.Must(new(big.Int).SetString("105000000023", 10)),
						Hash:      common.HexToHash("0xee69b38fad2d28c86d6fd838ba0ab328844c753b67ceec37e3433aec58d71c1b"),
						Input:     hexutil.MustDecode("0xd2539b3700000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000002c0000000000000000000000000000000000000000000000000000000000dee08b000000000000000000000000000000000000000000000000000000000000007600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f496bb175d00000000000000000000000043ed7d1bf7c703136971ae5e64f6e7feea435535000000000000000000000000d13f5d567d425204dc79db61839408cd58ecc59500000000000000000000000000000000000000000000000000000000000000008e269104c7193a7d6c37b7c7862a74def81b03e5fba830e80b2400b1c41a4edd000000000000000000000000000000000000000000000000000000000dee08b000000000000000000000000000000000000000000000000000000000139e95400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000041888fa35d117707a58af15849f001756e8ec66b06dcaca82f8a991d095712ddbe7c522243dfe1f0bf0630ed2993ae05706d317c1647e67e23d6d5610e62ad63f81c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000260000000000000000000000000000000000000000000000000000000a1e952a096000000000000000000000000abf500a9b3481ee143fde5078df1a202eb88a4a6000000000000000000000000f5e89105b63e8e14d2fd7cdfd191252bb0521169000000000000000000000000000000000000000000000000000000000000000012b2de357d4a2d0c4160fcda37eb93483a7d40180de99454026156695463f2b50000000000000000000000000000000000000000000000000000000011490c80000000000000000000000000000000000000000000000000000000003b9aca000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000041097b2cd11d6c98fb95df8686586be5d271a0550457c278a96e424fda2d24cbca770f0dad6010796a65f58dbba23504b818d536194704ef7f2818fce495d617c41b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d163d7e4f20000000000000000000000006b7ec4ba079c0a258435ea36025f6190cd4245620000000000000000000000006e5fa2f2b3268c32966d0683a0a6fbd87f080855000000000000000000000000000000000000000000000000000000000000000012b2de357d4a2d0c4160fcda37eb93483a7d40180de99454026156695463f2b50000000000000000000000000000000000000000000000000000000000f360b00000000000000000000000000000000000000000000000000000000003473bc00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000001a0000000000000000000000000000000000000000000000000000000000000004140e6ed17550e508e9c24f9e53180bee01c24cb08b73fe53a0b6532655c57d3d6324faea0ca8ae989b14778a7b5c69a8b6c4a77d890c8519a4e00a7db6e8e55d01b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000004bd2be00000000000000000000000000000000000000000000000000000000000f360b0"),
						To:        lo.ToPtr(common.HexToAddress("0x56C79347e95530c01A2FC76E732f9566dA16E113")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x32422585cf4fdb3977cbb80f53e02211bc88b8bed1f4528907937d7b8db2bb9d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("61724368", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 1033301,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x18727cda17"),
						GasUsed:           478405,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
							Topics: []common.Hash{
								common.HexToHash("0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6"),
								common.HexToHash("0x629828c8f231e1ee97e793768b807a5f940bb30dc5e28fe1e7fa30137e76eac1"),
								common.HexToHash("0x000000000000000000000000abf500a9b3481ee143fde5078df1a202eb88a4a6"),
								common.HexToHash("0x00000000000000000000000043ed7d1bf7c703136971ae5e64f6e7feea435535"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000012b2de357d4a2d0c4160fcda37eb93483a7d40180de99454026156695463f2b50000000000000000000000000000000000000000000000000000000004bd2be000000000000000000000000000000000000000000000000000000000105759800000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("61724368", 0)),
							TransactionHash: common.HexToHash("0xee69b38fad2d28c86d6fd838ba0ab328844c753b67ceec37e3433aec58d71c1b"),
							Index:           33,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xee69b38fad2d28c86d6fd838ba0ab328844c753b67ceec37e3433aec58d71c1b"),
						TransactionIndex: 4,
					},
				},
				config: &config.Module{
					Network: network.Polygon,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Polygon),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0xee69b38fad2d28c86d6fd838ba0ab328844c753b67ceec37e3433aec58d71c1b",
				Network:  network.Polygon,
				Index:    4,
				From:     "0x429B8474BD7308b7787d364985bB4b8eA7De1D47",
				To:       "0x56C79347e95530c01A2FC76E732f9566dA16E113",
				Type:     typex.CollectibleTrade,
				Platform: workerx.PlatformPolymarket.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("50232525011003315")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0xd2539b37",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.CollectibleTrade,
						Platform: workerx.PlatformPolymarket.String(),
						From:     "0xABf500a9b3481Ee143FDE5078df1A202Eb88A4A6",
						To:       "0x43eD7d1Bf7c703136971Ae5E64F6E7fEeA435535",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Name:     "Polygon",
								Symbol:   "MATIC",
								Decimals: 18,
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("79506400"))),
							},
							Cost: &metadata.Token{
								Address:  lo.ToPtr("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("8457663681790058947550769538996974245890751367516560139630487435494614626997"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("274160000"))),
								Standard: metadata.StandardERC1155,
							},
						},
					},
					{
						Type:     typex.CollectibleTrade,
						Platform: workerx.PlatformPolymarket.String(),
						From:     "0xABf500a9b3481Ee143FDE5078df1A202Eb88A4A6",
						To:       "0x43eD7d1Bf7c703136971Ae5E64F6E7fEeA435535",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeSell,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("8457663681790058947550769538996974245890751367516560139630487435494614626997"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("274160000"))),
								Standard: metadata.StandardERC1155,
							},
							Cost: &metadata.Token{
								Name:     "Polygon",
								Symbol:   "MATIC",
								Decimals: 18,
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("79506400"))),
							},
						},
					}},
				Status:    true,
				Timestamp: 1726122849,
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

			activity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
