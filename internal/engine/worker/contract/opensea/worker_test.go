package opensea_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/opensea"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
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
			name: "Buy ZGC #5 on Seaport V1.1",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x3958c2c95e2048e95344dd4846d387b309edafbf96e0a9ac4c4f4152ff253c59"),
						ParentHash:   common.HexToHash("0x9da3486783daab23246967b3f067ef82238a687e6a6c3b229a7fb58e2ada5452"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1959062D601c40Ef3B093a63f0954C6eE54A0a0F"),
						Number:       lo.Must(new(big.Int).SetString("16216786", 0)),
						GasLimit:     30000000,
						GasUsed:      24326175,
						Timestamp:    1671429323,
						BaseFee:      lo.Must(new(big.Int).SetString("11717147146", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781"),
						From:      common.HexToAddress("0x934B510D4C9103E6a87AEf13b816fb080286D649"),
						Gas:       173902,
						Hash:      common.HexToHash("0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781"),
						Input:     hexutil.MustDecode("0xfb0f3ee100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000339a69195cd0000000000000000000000000006fec7e57a21c3594619ae3ccc72330d5c71e6d11000000000000000000000000004c00500000ad104d7dbd00e3ae0a5c00560c0000000000000000000000000074ee68a33f6c9f113e22b3b77418b75f85d07d2200000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000639ff95e0000000000000000000000000000000000000000000000000000000063c8d7de0000000000000000000000000000000000000000000000000000000000000000360c6ebe0000000000000000000000000000000000000000976dfaaf14cabe1e0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f00000000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f00000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002e0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000017970b794f0000000000000000000000000000000a26b00c1f0df003000390027140000faa7190000000000000000000000000000000000000000000000000005e5c2de53c000000000000000000000000000d0f26bb9ec263b540e7eec7afb25ecfd10b9ceac0000000000000000000000000000000000000000000000000000000000000041ee8bde9e3ab60f91e25867b929e0a13627a516b9905a99dce3e3464c470c6a6403bf661671ffbb260e292b097d3cd61261c9646084d0d0bf7162782d64ac9b6d1b00000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")),
						Value:     lo.Must(new(big.Int).SetString("16600000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x3958c2c95e2048e95344dd4846d387b309edafbf96e0a9ac4c4f4152ff253c59"),
						BlockNumber:       lo.Must(new(big.Int).SetString("16216786", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 24280078,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x2bc9e86d5"),
						GasUsed:           153122,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31"),
								common.HexToHash("0x0000000000000000000000006fec7e57a21c3594619ae3ccc72330d5c71e6d11"),
								common.HexToHash("0x000000000000000000000000004c00500000ad104d7dbd00e3ae0a5c00560c00"),
							},
							Data:            hexutil.MustDecode("0x4d99ffe7fcc5f7e0700f00afbc2db580bd1814ec0782fad330831f33d230051f000000000000000000000000934b510d4c9103e6a87aef13b816fb080286d649000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000074ee68a33f6c9f113e22b3b77418b75f85d07d2200000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000339a69195cd0000000000000000000000000006fec7e57a21c3594619ae3ccc72330d5c71e6d1100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000017970b794f0000000000000000000000000000000a26b00c1f0df003000390027140000faa7190000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005e5c2de53c000000000000000000000000000d0f26bb9ec263b540e7eec7afb25ecfd10b9ceac"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16216786", 0)),
							TransactionHash: common.HexToHash("0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781"),
							Index:           138,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x74EE68a33f6c9f113e22B3B77418B75f85d07D22"),
							Topics: []common.Hash{
								common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
								common.HexToHash("0x0000000000000000000000001e0049783f008a0085193e00003d00cd54003c71"),
								common.HexToHash("0x0000000000000000000000006fec7e57a21c3594619ae3ccc72330d5c71e6d11"),
								common.HexToHash("0x000000000000000000000000934b510d4c9103e6a87aef13b816fb080286d649"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16216786", 0)),
							TransactionHash: common.HexToHash("0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781"),
							Index:           139,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD0f26Bb9Ec263b540e7eEC7aFb25ecFd10B9CeAC"),
							Topics: []common.Hash{
								common.HexToHash("0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d"),
								common.HexToHash("0x00000000000000000000000000000000006c3852cbef3e08e8df289169ede581"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000005e5c2de53c000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("16216786", 0)),
							TransactionHash: common.HexToHash("0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781"),
							Index:           140,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781"),
						TransactionIndex: 96,
					},
				},
				config: &engine.Config{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:       "0xab3b7cc0bdbe6bde86855c82adc59e65fef6b9fe5abc5ccd4966076637a90781",
				Network:  filter.NetworkEthereum,
				Index:    96,
				From:     "0x934B510D4C9103E6a87AEf13b816fb080286D649",
				To:       "0x00000000006c3852cbEf3e08E8dF289169EdE581",
				Type:     filter.TypeCollectibleTrade,
				Platform: lo.ToPtr(filter.PlatformOpenSea),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1799863423694410")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeCollectibleTrade,
						Platform: filter.PlatformOpenSea.String(),
						From:     "0x6FEC7E57A21C3594619ae3CCc72330D5C71e6d11",
						To:       "0x934B510D4C9103E6a87AEf13b816fb080286D649",
						Metadata: metadata.CollectibleTrade{
							Action: metadata.ActionCollectibleTradeBuy,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x74EE68a33f6c9f113e22B3B77418B75f85d07D22"),
								ID:       lo.ToPtr(lo.Must(decimal.NewFromString("5"))),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1"))),
								Name:     "Zerion Genesis Collection",
								Symbol:   "ZGC",
								Standard: contract.StandardERC1155,
								URI:      "ipfs://QmXBUGFTXuAeBfK9oB9G1NAhGq7AwosWjHFRHMdahETeRK",
							},

							Cost: &metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("14525000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1671429323,
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
