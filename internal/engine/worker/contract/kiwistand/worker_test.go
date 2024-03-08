package kiwistand_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/kiwistand"
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
			name: "Kiwistand Mint",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkOptimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x34e94a884c38722a326185213fc4b7788c51a0b5bb8fe2b9f1c9127b7f0ad935"),
						ParentHash:   common.HexToHash("0x17319c4cc9818b4acdadfd062bdf161a3ae3c570f413072913e12dc49ed6eb8a"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("115584324", 0)),
						GasLimit:     30000000,
						GasUsed:      4412970,
						Timestamp:    1706767425,
						BaseFee:      lo.Must(new(big.Int).SetString("4425002", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
						From:      common.HexToAddress("0xf70da97812CB96acDF810712Aa562db8dfA3dbEF"),
						Gas:       164252,
						GasPrice:  lo.Must(new(big.Int).SetString("4530002", 10)),
						Hash:      common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
						Input:     hexutil.MustDecode("0x4536818100000000000000000000000068b885f956e3f0457f39b1d9efded041a99af91500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034c464700000000000000000000000000000000000000000000000000000000001d4da48b274ce2200364ae"),
						To:        lo.ToPtr(common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45")),
						Value:     lo.Must(new(big.Int).SetString("3337000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x34e94a884c38722a326185213fc4b7788c51a0b5bb8fe2b9f1c9127b7f0ad935"),
						BlockNumber:       lo.Must(new(big.Int).SetString("115584324", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2680241,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x451f52"),
						GasUsed:           130304,
						L1GasPrice:        lo.Must(new(big.Int).SetString("18224688442", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("3400", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("42383335440715", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x7777777F279eba3d3Ad8F4E708545291A6fDBA8B"),
							Topics: []common.Hash{
								common.HexToHash("0x90e8cce6b15b450d1e56e9ef986d1cd376838a90944336c02886ca12b9e6ebd7"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000007a810dcd0f8d83b20212326813db6ef7e9fd030c"),
								common.HexToHash("0x0000000000000000000000007a810dcd0f8d83b20212326813db6ef7e9fd030c"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000ee324c588cef1bf1c1360883e4318834af66366d0000000000000000000000007a810dcd0f8d83b20212326813db6ef7e9fd030c00000000000000000000000066747bdc903d17c586fa09ee5d6b54cc85bbea4500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c9e86723e0000000000000000000000000000000000000000000000000000000c9e86723e000000000000000000000000000000000000000000000000000000064f43391f0000000000000000000000000000000000000000000000000000000c9e86723e000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("115584324", 0)),
							TransactionHash: common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
							Index:           79,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000068b885f956e3f0457f39b1d9efded041a99af915"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000003d3"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("115584324", 0)),
							TransactionHash: common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
							Index:           80,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45"),
							Topics: []common.Hash{
								common.HexToHash("0x4e26b0356a15833a75d497ecc40ebbb716b99466ed0dba9454f1fff451e25a90"),
								common.HexToHash("0x00000000000000000000000068b885f956e3f0457f39b1d9efded041a99af915"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000009184e72a00000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000003d2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("115584324", 0)),
							TransactionHash: common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
							Index:           81,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45"),
							Topics: []common.Hash{
								common.HexToHash("0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a"),
								common.HexToHash("0x000000000000000000000000f70da97812cb96acdf810712aa562db8dfa3dbef"),
								common.HexToHash("0x00000000000000000000000066747bdc903d17c586fa09ee5d6b54cc85bbea45"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000003d2"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000034c46470000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("115584324", 0)),
							TransactionHash: common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
							Index:           82,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557"),
						TransactionIndex: 44,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkOptimism,
					Endpoint: endpoint.MustGet(filter.NetworkOptimism),
				},
			},
			want: &schema.Feed{
				ID:       "0xa4ded3ac75f8aef67c01e836306097b588548fae57d51da376f0acbaa8387557",
				Network:  filter.NetworkOptimism,
				Index:    44,
				From:     "0xf70da97812CB96acDF810712Aa562db8dfA3dbEF",
				To:       "0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45",
				Type:     filter.TypeCollectibleMint,
				Platform: lo.ToPtr(filter.PlatformKiwiStand),
				Calldata: &schema.Calldata{
					FunctionHash: "0x45368181",
				},
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("42973612821323")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformKiwiStand.String(),
						From:     common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45").String(),
						To:       common.HexToAddress("0x7777777F279eba3d3Ad8F4E708545291A6fDBA8B").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("222000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformKiwiStand.String(),
						From:     common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45").String(),
						To:       common.HexToAddress("0x7777777F279eba3d3Ad8F4E708545291A6fDBA8B").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("222000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformKiwiStand.String(),
						From:     common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45").String(),
						To:       common.HexToAddress("0x7777777F279eba3d3Ad8F4E708545291A6fDBA8B").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("111000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformKiwiStand.String(),
						From:     common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45").String(),
						To:       common.HexToAddress("0x7777777F279eba3d3Ad8F4E708545291A6fDBA8B").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("222000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeCollectibleMint,
						Platform: filter.PlatformKiwiStand.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       common.HexToAddress("0x68b885F956e3F0457f39B1d9eFDeD041A99AF915").String(),
						Metadata: metadata.CollectibleTransfer{
							Address:  lo.ToPtr("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45"),
							ID:       lo.ToPtr(decimal.NewFromBigInt(big.NewInt(979), 0)),
							Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0)),
							Name:     "Kiwi Pass",
							Symbol:   "$KIWI",
							Standard: metadata.StandardERC721,
							URI:      "data:application/json;base64,eyJuYW1lIjogIktpd2kgUGFzcyA5NzkiLCAiZGVzY3JpcHRpb24iOiAiS2l3aSBQYXNzIG9wZW5zIHRoZSBkb29yIHRvIHRoZSBjb21tdW5pdHkgb2YgY3VyYXRvcnMgd2hvIHNoYXBlIHRoZSBLaXdpIGZlZWQuXG5cbkdvIHRvIGtpd2luZXdzLnh5eiBmb3IgeW91ciBkYWlseSBkb3NlIG9mIGtpd2kg8J+lnVxuXG5EbyBOT1QgYnV5IEtpd2kgUGFzc2VzIG9uIHRoZSBzZWNvbmRhcnkgbWFya2V0LiBUaGV5IHdpbGwgTk9UIHdvcmsuIiwgImltYWdlIjogImlwZnM6Ly9iYWZ5YmVpYzJhcmFqNWtvanVhZ2tqZXJ5aHF5Y2Y3dnZhZmczMmJjdzNsNXUzZ3Q2YXM0MmZseWEyeSIsICJwcm9wZXJ0aWVzIjogeyJudW1iZXIiOiA5NzksICJuYW1lIjogIktpd2kgUGFzcyJ9fQ==",
						},
					},
					{
						Type:     filter.TypeTransactionTransfer,
						Platform: filter.PlatformKiwiStand.String(),
						From:     common.HexToAddress("0xf70da97812CB96acDF810712Aa562db8dfA3dbEF").String(),
						To:       common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45").String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2560000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     filter.TypeSocialMint,
						Platform: filter.PlatformKiwiStand.String(),
						From:     common.HexToAddress("0xf70da97812CB96acDF810712Aa562db8dfA3dbEF").String(),
						To:       common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45").String(),
						Metadata: &metadata.SocialPost{
							Title: "LFG",
							Body:  "LFG",
						},
					},
				},
				Status:    true,
				Timestamp: 1706767425,
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

			t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}
