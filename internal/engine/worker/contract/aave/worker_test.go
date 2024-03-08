package aave_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/aave"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/aave"
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
			name: "Deposit stETH to AAVE V2 LendingPool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x2944e0b4d7a8f8530da13880211eb675bb18de7340bda8c3f703dd371d45e3e9"),
						ParentHash:   common.HexToHash("0x7448f11755f1cba0d1e167eca137c333fb24a805ee99dfb531cdc8c3e5b9376c"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17883118", 0)),
						GasLimit:     30000000,
						GasUsed:      10664885,
						Timestamp:    1691653955,
						BaseFee:      lo.Must(new(big.Int).SetString("16224967415", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
						From:      common.HexToAddress("0xF8E4517dC4fd4bfeF9903336ADB1Ede20803430d"),
						Gas:       300000,
						GasPrice:  lo.Must(new(big.Int).SetString("16324967415", 10)),
						Hash:      common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
						Input:     hexutil.MustDecode("0xe8eda9df000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe840000000000000000000000000000000000000000000000068155a43676dfffff000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d0000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x2944e0b4d7a8f8530da13880211eb675bb18de7340bda8c3f703dd371d45e3e9"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17883118", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6523882,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3cd0b3bf7"),
						GasUsed:           244391,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000033b3fdd66b253e7cda7a3f60000000000000000000000000000000000000000033b2e3c9fd0803ce8000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           179,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x0000000000000000000000007d2768de32b0b80b7a3454c06bdac94a69ddc7a9"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           180,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x0000000000000000000000001982b2f5814301d4e9a8b0201555376e62f82428"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000068155a43676dfffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           181,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
							Topics: []common.Hash{
								common.HexToHash("0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x0000000000000000000000001982b2f5814301d4e9a8b0201555376e62f82428"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000005bae746d84eb59c85"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           182,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1982b2F5814301d4e9a8b0201555376e62F82428"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000068155a43676dfffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           183,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1982b2F5814301d4e9a8b0201555376e62F82428"),
							Topics: []common.Hash{
								common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000068155a43676dfffff0000000000000000000000000000000000000000033b3fdd66b253e7cda7a3f6"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           184,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2"),
								common.HexToHash("0x000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           185,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951"),
								common.HexToHash("0x000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d0000000000000000000000000000000000000000000000068155a43676dfffff"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883118", 0)),
							TransactionHash: common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
							Index:           186,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c"),
						TransactionIndex: 75,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x488591689796c234eb6a876be238c67a9092d755070d6992dddccd2f8beee57c",
				Network: filter.NetworkEthereum,
				Index:   75,
				From:    "0xF8E4517dC4fd4bfeF9903336ADB1Ede20803430d",
				To:      aave.AddressV2LendingPoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0xe8eda9df",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3989675111519265")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0xF8E4517dC4fd4bfeF9903336ADB1Ede20803430d",
						To:       aave.AddressV2LendingPoolMainnet.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquiditySupply,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("119999999999999999999"))),
									Name:     "Liquid staked Ether 2.0",
									Symbol:   "stETH",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691653955,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw USDC from AAVE V2 LendingPool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x74d8c5cbf1a3df7c45f431553216a7c304a80f3ef4fecb7b5611bb6f282a421a"),
						ParentHash:   common.HexToHash("0xa10d8a57fee5963862fb019d45e6325b1513aff11e40a475f8ba7f580543f49e"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17883249", 0)),
						GasLimit:     30000000,
						GasUsed:      14329964,
						Timestamp:    1691655527,
						BaseFee:      lo.Must(new(big.Int).SetString("13648369325", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
						From:      common.HexToAddress("0x0f1DfeF1a40557d279d0de6E49aB306891A638b8"),
						Gas:       335046,
						GasPrice:  lo.Must(new(big.Int).SetString("13748369325", 10)),
						Hash:      common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
						Input:     hexutil.MustDecode("0x69328dec000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000048c273950000000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
						To:        lo.ToPtr(common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x74d8c5cbf1a3df7c45f431553216a7c304a80f3ef4fecb7b5611bb6f282a421a"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17883249", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 7867074,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3337777ad"),
						GasUsed:           254165,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000464c71f6c2f760dda6093dcb91c24c39e5d6e18c"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000345025f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           226,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"),
							Topics: []common.Hash{
								common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"),
								common.HexToHash("0x000000000000000000000000464c71f6c2f760dda6093dcb91c24c39e5d6e18c"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000345025f0000000000000000000000000000000000000000038c7bb08d181615b947c0d1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           227,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000003daf6590bbe1c61b114d3400000000000000000000000000000000000000000089b1af5c699e0d9c0d72320000000000000000000000000000000000000000004fca9bc1dd9513960d72320000000000000000000000000000000000000000038c7bb08d181615b947c0d1000000000000000000000000000000000000000003b2ab37c009fadb0cdca929"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           228,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000bcca60bb61934080951369a648fb03df4f96263c"),
								common.HexToHash("0x0000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000048c27395000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           229,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000048c27395000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           230,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"),
							Topics: []common.Hash{
								common.HexToHash("0x5d624aa9c148153ab3446c1b154f660ee7701e549fe9b62dab7171b1c80e6fa2"),
								common.HexToHash("0x0000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
								common.HexToHash("0x0000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000048c273950000000000000000000000000000000000000000000038c7bb08d181615b947c0d1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           231,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
								common.HexToHash("0x0000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
								common.HexToHash("0x0000000000000000000000000f1dfef1a40557d279d0de6e49ab306891a638b8"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000048c27395000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883249", 0)),
							TransactionHash: common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
							Index:           232,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6"),
						TransactionIndex: 100,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x759532f09bd943d6da2964d58d7fc9a6b3f3914ef0643c37cb5f0eb1f3fe03a6",
				Network: filter.NetworkEthereum,
				Index:   100,
				From:    "0x0f1DfeF1a40557d279d0de6E49aB306891A638b8",
				To:      aave.AddressV2LendingPoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0x69328dec",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3494354289488625")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0x0f1DfeF1a40557d279d0de6E49aB306891A638b8",
						To:       aave.AddressV2LendingPoolMainnet.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityWithdraw,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("5000000000000"))),
									Name:     "USD Coin",
									Symbol:   "USDC",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691655527,
			},
			wantError: require.NoError,
		},
		{
			name: "Borrow USDT from AAVE V2 LendingPool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x644b27dbc2a39e9abacdec00d7347f254531a2a87b997523a50a8a2496fa7393"),
						ParentHash:   common.HexToHash("0x8de61afb4ed0d1f7d51b3c979797ae65bef849fe295ebe51bb056cbbefb6c316"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17883124", 0)),
						GasLimit:     30000000,
						GasUsed:      12645053,
						Timestamp:    1691654027,
						BaseFee:      lo.Must(new(big.Int).SetString("16215464117", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
						From:      common.HexToAddress("0xF8E4517dC4fd4bfeF9903336ADB1Ede20803430d"),
						Gas:       470281,
						GasPrice:  lo.Must(new(big.Int).SetString("16315464117", 10)),
						Hash:      common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
						Input:     hexutil.MustDecode("0xa415bcad000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000ba43b740000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
						To:        lo.ToPtr(common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x644b27dbc2a39e9abacdec00d7347f254531a2a87b997523a50a8a2496fa7393"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17883124", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10763994,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3cc7a39b5"),
						GasUsed:           403016,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x3Ed3B47Dd13EC9a98b44e6204A523E766B225811"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000464c71f6c2f760dda6093dcb91c24c39e5d6e18c"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000005682adf"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           240,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x3Ed3B47Dd13EC9a98b44e6204A523E766B225811"),
							Topics: []common.Hash{
								common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"),
								common.HexToHash("0x000000000000000000000000464c71f6c2f760dda6093dcb91c24c39e5d6e18c"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000005682adf000000000000000000000000000000000000000003a15347b75a78a68af0a303"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           241,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd784927Ff2f95ba542BfC824c8a8a98F3495f6b5"),
							Topics: []common.Hash{
								common.HexToHash("0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x000000000000000000000000531842cebbdd378f8ee36d171d6cc9c4fcf475ec"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000051a008816f5c9f3e8feb6a"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           242,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x531842cEbbdD378f8ee36D171d6cC9C4fcf475Ec"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000ba43b7400"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           243,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x531842cEbbdD378f8ee36D171d6cC9C4fcf475Ec"),
							Topics: []common.Hash{
								common.HexToHash("0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000ba43b7400000000000000000000000000000000000000000003d4ec669e20dccbb17bddd4"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           244,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000028b1737b093a937c0e76510000000000000000000000000000000000000000007d52ed06d0446e7e9384900000000000000000000000000000000000000000003b264456303a2c2e938490000000000000000000000000000000000000000003a15347b75a78a68af0a303000000000000000000000000000000000000000003d4ec669e20dccbb17bddd4"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           245,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000003ed3b47dd13ec9a98b44e6204a523e766b225811"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000ba43b7400"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           246,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b"),
								common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
								common.HexToHash("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000f8e4517dc4fd4bfef9903336adb1ede20803430d0000000000000000000000000000000000000000000000000000000ba43b740000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000003b264456303a2c2e938490"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883124", 0)),
							TransactionHash: common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
							Index:           247,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f"),
						TransactionIndex: 108,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xd01b5eab0d6abceb5f21da3dc6c329b35b5f854fd4c0e7d05b2cd81e9a410c9f",
				Network: filter.NetworkEthereum,
				Index:   108,
				From:    "0xF8E4517dC4fd4bfeF9903336ADB1Ede20803430d",
				To:      aave.AddressV2LendingPoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0xa415bcad",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("6575393086576872")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0xF8E4517dC4fd4bfeF9903336ADB1Ede20803430d",
						To:       aave.AddressV2LendingPoolMainnet.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityBorrow,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("50000000000"))),
									Name:     "Tether USD",
									Symbol:   "USDT",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691654027,
			},
			wantError: require.NoError,
		},
		{
			name: "Repay USDC to AAVE V2 LendingPool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xd05278242892db7a04b060514abd7f5616b3f135acd87a05037edfe42aac0393"),
						ParentHash:   common.HexToHash("0x8b9be781a2b9b3a2d208efdb3831735f5e00cfdb8f2f8ef9cff9373f6070d7ff"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x30583a30D1698358E1101650C4E7f534f03e7A73"),
						Number:       lo.Must(new(big.Int).SetString("17883185", 0)),
						GasLimit:     30000000,
						GasUsed:      8580084,
						Timestamp:    1691654759,
						BaseFee:      lo.Must(new(big.Int).SetString("18005942185", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
						From:      common.HexToAddress("0xc11E5B31008F3750b43f006c3026beD653888E07"),
						Gas:       326396,
						GasPrice:  lo.Must(new(big.Int).SetString("18055942185", 10)),
						Hash:      common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
						Input:     hexutil.MustDecode("0x573ade81000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4800000000000000000000000000000000000000000000000000000006e000854d0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c11e5b31008f3750b43f006c3026bed653888e07"),
						To:        lo.ToPtr(common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xd05278242892db7a04b060514abd7f5616b3f135acd87a05037edfe42aac0393"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17883185", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8321640,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x43437d029"),
						GasUsed:           242817,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000464c71f6c2f760dda6093dcb91c24c39e5d6e18c"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000959b078"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           152,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xBcca60bB61934080951369a648Fb03DF4F96263C"),
							Topics: []common.Hash{
								common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"),
								common.HexToHash("0x000000000000000000000000464c71f6c2f760dda6093dcb91c24c39e5d6e18c"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000959b0780000000000000000000000000000000000000000038c7b82eb637b940e79a406"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           153,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x619beb58998eD2278e08620f97007e1116D5D25b"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000c11e5b31008f3750b43f006c3026bed653888e07"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000006e000854d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           154,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x619beb58998eD2278e08620f97007e1116D5D25b"),
							Topics: []common.Hash{
								common.HexToHash("0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a"),
								common.HexToHash("0x000000000000000000000000c11e5b31008f3750b43f006c3026bed653888e07"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000006e000854d000000000000000000000000000000000000000003b2aafb6f81496b1fe87270"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           155,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000001a0850992baab07bbd0ebf0000000000000000000000000000000000000000005af4ee1413c0d48c42b7b80000000000000000000000000000000000000000002105609abf6a93e4856f700000000000000000000000000000000000000000038c7b82eb637b940e79a406000000000000000000000000000000000000000003b2aafb6f81496b1fe87270"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           156,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000c11e5b31008f3750b43f006c3026bed653888e07"),
								common.HexToHash("0x000000000000000000000000bcca60bb61934080951369a648fb03df4f96263c"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000006e000854d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           157,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9"),
							Topics: []common.Hash{
								common.HexToHash("0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa"),
								common.HexToHash("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
								common.HexToHash("0x000000000000000000000000c11e5b31008f3750b43f006c3026bed653888e07"),
								common.HexToHash("0x000000000000000000000000c11e5b31008f3750b43f006c3026bed653888e07"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000006e000854d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17883185", 0)),
							TransactionHash: common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
							Index:           158,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec"),
						TransactionIndex: 86,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x3c309fa16fb7e995fba8f236d2eb3a20a44f3ce116a13899683f4580154e85ec",
				Network: filter.NetworkEthereum,
				Index:   86,
				From:    "0xc11E5B31008F3750b43f006c3026beD653888E07",
				To:      aave.AddressV2LendingPoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0x573ade81",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("4384289713535145")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0xc11E5B31008F3750b43f006c3026beD653888E07",
						To:       aave.AddressV2LendingPoolMainnet.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRepay,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("29527934285"))),
									Name:     "USD Coin",
									Symbol:   "USDC",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691654759,
			},
			wantError: require.NoError,
		},
		{
			name: "Supply rETH to AAVE V3 Pool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x06745b2b0c2f4ab248de64f1f744217ccbc68003012e0e50e85970a1b3b36d28"),
						ParentHash:   common.HexToHash("0x7390411df78651942ea6503b7b64c97eb4f22b84be463817daaaef2cec7fab8a"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17882365", 0)),
						GasLimit:     30000000,
						GasUsed:      13132938,
						Timestamp:    1691644787,
						BaseFee:      lo.Must(new(big.Int).SetString("14398093469", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
						From:      common.HexToAddress("0xf01bb28137121c063D73e7B61DBABE352467292b"),
						Gas:       300000,
						GasPrice:  lo.Must(new(big.Int).SetString("14498093469", 10)),
						Hash:      common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
						Input:     hexutil.MustDecode("0x617ba037000000000000000000000000ae78736cd615f374d3085123a210448e74fc6393000000000000000000000000000000000000000000000000022be058804b909f000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b0000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x06745b2b0c2f4ab248de64f1f744217ccbc68003012e0e50e85970a1b3b36d28"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17882365", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10542495,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x36027599d"),
						GasUsed:           175191,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000ae78736cd615f374d3085123a210448e74fc6393"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000b190bb4c83c3950117b600000000000000000000000000000000000000000054b172ca5915f7ae45f64f0000000000000000000000000000000000000000000a3f3503a50a6d1445f64f0000000000000000000000000000000000000000033bc3c00adbe19c35d763920000000000000000000000000000000000000000033ea8b27c7d3a86ac0f5563"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882365", 0)),
							TransactionHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
							Index:           288,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae78736Cd615f374D3085123A210448E74Fc6393"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b"),
								common.HexToHash("0x000000000000000000000000cc9ee9483f662091a1de4795249e24ac0ac2630f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000022be058804b909f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882365", 0)),
							TransactionHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
							Index:           289,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xae78736Cd615f374D3085123A210448E74Fc6393"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b"),
								common.HexToHash("0x00000000000000000000000087870bca3f3fd6335c3f4ce8392d69350b4fa4e2"),
							},
							Data:            hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffffffff3a26faa89e0d0ab"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882365", 0)),
							TransactionHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
							Index:           290,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xCc9EE9483f662091a1de4795249E24aC0aC2630f"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000022be06252f1a87c"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882365", 0)),
							TransactionHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
							Index:           291,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xCc9EE9483f662091a1de4795249E24aC0aC2630f"),
							Topics: []common.Hash{
								common.HexToHash("0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196"),
								common.HexToHash("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b"),
								common.HexToHash("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000022be06252f1a87c00000000000000000000000000000000000000000000000000000009d2a617dd0000000000000000000000000000000000000000033bc3c00adbe19c35d76392"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882365", 0)),
							TransactionHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
							Index:           292,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61"),
								common.HexToHash("0x000000000000000000000000ae78736cd615f374d3085123a210448e74fc6393"),
								common.HexToHash("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000f01bb28137121c063d73e7b61dbabe352467292b000000000000000000000000000000000000000000000000022be058804b909f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882365", 0)),
							TransactionHash: common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
							Index:           293,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce"),
						TransactionIndex: 85,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x2626b06195b4e7699495b1ea542a954a7b164195f3fa2dbd1c063882cf697bce",
				Network: filter.NetworkEthereum,
				Index:   85,
				From:    "0xf01bb28137121c063D73e7B61DBABE352467292b",
				To:      aave.AddressV3PoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0x617ba037",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2539935492927579")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0xf01bb28137121c063D73e7B61DBABE352467292b",
						To:       aave.AddressV3PoolMainnet.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquiditySupply,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xae78736Cd615f374D3085123A210448E74Fc6393"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("156465282788593823"))),
									Name:     "Rocket Pool ETH",
									Symbol:   "rETH",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691644787,
			},
			wantError: require.NoError,
		},
		{
			name: "Supply USDCbC to AAVE V3 Pool on Ethereum Base",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkBase,
					ChainID: 8453,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x8b5d7e878bf84a83fc537a05154461d33906c7e470b6f58ea233d8a3b3de288c"),
						ParentHash:   common.HexToHash("0x1b67512207131dcce344741ba74b8343e94364002010f0fabb23fdace1d0a23d"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("2963358", 0)),
						GasLimit:     30000000,
						GasUsed:      1685052,
						Timestamp:    1692716063,
						BaseFee:      lo.Must(new(big.Int).SetString("59", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
						From:      common.HexToAddress("0x1aEB9F637B88d9c5422E0cb094F3bCDa20fCE37B"),
						Gas:       300000,
						GasPrice:  lo.Must(new(big.Int).SetString("16008", 10)),
						Hash:      common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
						Input:     hexutil.MustDecode("0x617ba037000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca00000000000000000000000000000000000000000000000000000000000f42400000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b0000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("8453", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x8b5d7e878bf84a83fc537a05154461d33906c7e470b6f58ea233d8a3b3de288c"),
						BlockNumber:       lo.Must(new(big.Int).SetString("2963358", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 1685052,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3e88"),
						GasUsed:           233332,
						L1GasPrice:        lo.Must(new(big.Int).SetString("92801097015", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("3024", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("191951273883378", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000295be96e6406697200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000033b2e3c9fd0803ce80000000000000000000000000000000000000000000000033b2e3c9fd0803ce8000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("2963358", 0)),
							TransactionHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
							Index:           25,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b"),
								common.HexToHash("0x0000000000000000000000000a1d576f3efef75b330424287a95a366e8281d54"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000f4240"),
							BlockNumber:     lo.Must(new(big.Int).SetString("2963358", 0)),
							TransactionHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
							Index:           26,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0a1d576f3eFeF75b330424287a95A366e8281D54"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000f4240"),
							BlockNumber:     lo.Must(new(big.Int).SetString("2963358", 0)),
							TransactionHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
							Index:           27,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0a1d576f3eFeF75b330424287a95A366e8281D54"),
							Topics: []common.Hash{
								common.HexToHash("0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196"),
								common.HexToHash("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b"),
								common.HexToHash("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000033b2e3c9fd0803ce8000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("2963358", 0)),
							TransactionHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
							Index:           28,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
							Topics: []common.Hash{
								common.HexToHash("0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2"),
								common.HexToHash("0x000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca"),
								common.HexToHash("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("2963358", 0)),
							TransactionHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
							Index:           29,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
							Topics: []common.Hash{
								common.HexToHash("0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61"),
								common.HexToHash("0x000000000000000000000000d9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca"),
								common.HexToHash("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000001aeb9f637b88d9c5422e0cb094f3bcda20fce37b00000000000000000000000000000000000000000000000000000000000f4240"),
							BlockNumber:     lo.Must(new(big.Int).SetString("2963358", 0)),
							TransactionHash: common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
							Index:           30,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a"),
						TransactionIndex: 22,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkBase,
					Endpoint: endpoint.MustGet(filter.NetworkBase),
				},
			},
			want: &schema.Feed{
				ID:      "0x3ee309226dfd6a4d26be87082e5e1ce6c8769f130c9495b2e44684ad5cc2682a",
				Network: filter.NetworkBase,
				Index:   22,
				From:    "0x1aEB9F637B88d9c5422E0cb094F3bCDa20fCE37B",
				To:      aave.AddressV3PoolBase.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0x617ba037",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("191955009062034")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0x1aEB9F637B88d9c5422E0cb094F3bCDa20fCE37B",
						To:       aave.AddressV3PoolBase.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquiditySupply,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1000000"))),
									Name:     "USD Base Coin",
									Symbol:   "USDbC",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692716063,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw WBTC from AAVE V3 Pool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x2ac43f77fc58845ceec77c3f91dc7e582c6d0fbe1b414527ebaaba0de35c7114"),
						ParentHash:   common.HexToHash("0xa0b9bfed818813a608528148d4c1907f914691e4a63a346741f9ded3043f12b4"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17882399", 0)),
						GasLimit:     29999972,
						GasUsed:      11748770,
						Timestamp:    1691645195,
						BaseFee:      lo.Must(new(big.Int).SetString("13119335234", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
						From:      common.HexToAddress("0x5A0D6d0DE7c74899F09d3509A429bEb7D3b4b1d0"),
						Gas:       401310,
						GasPrice:  lo.Must(new(big.Int).SetString("13219335234", 10)),
						Hash:      common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
						Input:     hexutil.MustDecode("0x69328dec0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
						To:        lo.ToPtr(common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x2ac43f77fc58845ceec77c3f91dc7e582c6d0fbe1b414527ebaaba0de35c7114"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17882399", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 10150148,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x313ef0c42"),
						GasUsed:           300457,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000a1877cc1c704b1d67d020000000000000000000000000000000000000000003ef5210644a23a495e95a2000000000000000000000000000000000000000000079d816ebe0f29755a9ea70000000000000000000000000000000000000000033c34573ef9de0b49e7cadf00000000000000000000000000000000000000000341c60b2d2f5317d82475d1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882399", 0)),
							TransactionHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
							Index:           266,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd"),
								common.HexToHash("0x0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("17882399", 0)),
							TransactionHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
							Index:           267,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x5Ee5bf7ae06D1Be5997A1A72006FE6C607eC6DE8"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000be5873f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882399", 0)),
							TransactionHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
							Index:           268,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x5Ee5bf7ae06D1Be5997A1A72006FE6C607eC6DE8"),
							Topics: []common.Hash{
								common.HexToHash("0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000be5873f000000000000000000000000000000000000000000000000000000000000192c0000000000000000000000000000000000000000033c34573ef9de0b49e7cadf"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882399", 0)),
							TransactionHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
							Index:           269,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005ee5bf7ae06d1be5997a1a72006fe6c607ec6de8"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000be5a06b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882399", 0)),
							TransactionHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
							Index:           270,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7"),
								common.HexToHash("0x0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c599"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
								common.HexToHash("0x0000000000000000000000005a0d6d0de7c74899f09d3509a429beb7d3b4b1d0"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000be5a06b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882399", 0)),
							TransactionHash: common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
							Index:           271,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11"),
						TransactionIndex: 108,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xef6e2adee0defc810147847b2c3c5255e5e5e75cbc7a4e6ff009444f0aa53f11",
				Network: filter.NetworkEthereum,
				Index:   108,
				From:    "0x5A0D6d0DE7c74899F09d3509A429bEb7D3b4b1d0",
				To:      aave.AddressV3PoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0x69328dec",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3971841806401938")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     aave.AddressV3PoolMainnet.String(),
						To:       "0x5A0D6d0DE7c74899F09d3509A429bEb7D3b4b1d0",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityWithdraw,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("199598187"))),
									Name:     "Wrapped BTC",
									Symbol:   "WBTC",
									Decimals: 8,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691645195,
			},
			wantError: require.NoError,
		},
		{
			name: "Borrow USDT from AAVE V3 Pool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x931bea6a782b8f18027759b250a73d17fc32f3a6c3d1fd115419335bcb177e6d"),
						ParentHash:   common.HexToHash("0x4d0f0501b5e9e083e91949c4d9fcefafd1ca9cf4acf51454c8d6e314b53e8830"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17882281", 0)),
						GasLimit:     30000000,
						GasUsed:      16659485,
						Timestamp:    1691643779,
						BaseFee:      lo.Must(new(big.Int).SetString("12551475790", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
						From:      common.HexToAddress("0xC4dBC7D5957dceF7eD2B0778C597c16Ce2769E7d"),
						Gas:       400000,
						GasPrice:  lo.Must(new(big.Int).SetString("12651475790", 10)),
						Hash:      common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
						Input:     hexutil.MustDecode("0xa415bcad000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec700000000000000000000000000000000000000000000000000000001a13b860000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d"),
						To:        lo.ToPtr(common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x931bea6a782b8f18027759b250a73d17fc32f3a6c3d1fd115419335bcb177e6d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17882281", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 13143019,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x2f216334e"),
						GasUsed:           301906,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x6df1C1E379bC5a00a7b4C6e67A203333772f45A8"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000001a14c826b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882281", 0)),
							TransactionHash: common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
							Index:           291,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6df1C1E379bC5a00a7b4C6e67A203333772f45A8"),
							Topics: []common.Hash{
								common.HexToHash("0x458f5fa412d0f69b08dd84872b0215675cc67bc1d5b6fd93300a1c3878b86196"),
								common.HexToHash("0x000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d"),
								common.HexToHash("0x000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000001a14c826b000000000000000000000000000000000000000000000000000000000010fc6b0000000000000000000000000000000000000000034e0ca4cb8b8f1ad1cf7357"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882281", 0)),
							TransactionHash: common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
							Index:           292,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000172f53c8ab86fb33ecf2ba0000000000000000000000000000000000000000004a000fe3d6c9f030da6d7a00000000000000000000000000000000000000000020a4267572c386beda6d7a00000000000000000000000000000000000000000348c27fbf92067ad5d921490000000000000000000000000000000000000000034e0ca4cb8b8f1ad1cf7357"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882281", 0)),
							TransactionHash: common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
							Index:           293,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000023878914efe38d27c4d67ab83ed1b93a74d4086a"),
								common.HexToHash("0x000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000001a13b8600"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882281", 0)),
							TransactionHash: common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
							Index:           294,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0"),
								common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
								common.HexToHash("0x000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000c4dbc7d5957dcef7ed2b0778c597c16ce2769e7d00000000000000000000000000000000000000000000000000000001a13b8600000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000020a4267572c386beda6d7a"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882281", 0)),
							TransactionHash: common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
							Index:           295,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295"),
						TransactionIndex: 96,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0x060a2a7e1db5be295f14171e98a39dec1a7282101c148833333409b99dad4295",
				Network: filter.NetworkEthereum,
				Index:   96,
				From:    "0xC4dBC7D5957dceF7eD2B0778C597c16Ce2769E7d",
				To:      aave.AddressV3PoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0xa415bcad",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("3819556449855740")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     aave.AddressV3PoolMainnet.String(),
						To:       "0xC4dBC7D5957dceF7eD2B0778C597c16Ce2769E7d",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityBorrow,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("7000000000"))),
									Name:     "Tether USD",
									Symbol:   "USDT",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691643779,
			},
			wantError: require.NoError,
		},
		{
			name: "Repay USDT to AAVE V3 Pool on Ethereum Mainnet",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkEthereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x744f4914a79e72159e11169076fc6cf1db483d6f2c318f2f98cb95e94f535fec"),
						ParentHash:   common.HexToHash("0x19e8fccf2d190cdd7676657d3adfa1d6965e37c191a28388de165bc592b51a41"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17882851", 0)),
						GasLimit:     30000000,
						GasUsed:      28174472,
						Timestamp:    1691650679,
						BaseFee:      lo.Must(new(big.Int).SetString("13282618298", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
						From:      common.HexToAddress("0x790c9422839FD93a3A4E31e531f96cC87F397c00"),
						Gas:       300000,
						GasPrice:  lo.Must(new(big.Int).SetString("13382618298", 10)),
						Hash:      common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
						Input:     hexutil.MustDecode("0x573ade81000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000000000000000000000000000000000b09949e1800000000000000000000000000000000000000000000000000000000000000002000000000000000000000000790c9422839fd93a3a4e31e531f96cc87f397c00"),
						To:        lo.ToPtr(common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x744f4914a79e72159e11169076fc6cf1db483d6f2c318f2f98cb95e94f535fec"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17882851", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 25418104,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x31daa8cba"),
						GasUsed:           192149,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x6df1C1E379bC5a00a7b4C6e67A203333772f45A8"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000790c9422839fd93a3a4e31e531f96cc87f397c00"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000b098c80af1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882851", 0)),
							TransactionHash: common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
							Index:           513,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6df1C1E379bC5a00a7b4C6e67A203333772f45A8"),
							Topics: []common.Hash{
								common.HexToHash("0x4cf25bc1d991c17529c25213d3cc0cda295eeaad5f13f361969b12ea48015f90"),
								common.HexToHash("0x000000000000000000000000790c9422839fd93a3a4e31e531f96cc87f397c00"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000b098c80af1000000000000000000000000000000000000000000000000000000000081d68f0000000000000000000000000000000000000000034e0e835960574ff2229eb1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882851", 0)),
							TransactionHash: common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
							Index:           514,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a"),
								common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000016c88037b7801f792a42cb00000000000000000000000000000000000000000049b75ce4d7163368f49d11000000000000000000000000000000000000000000205b7376730fc9f6f49d1100000000000000000000000000000000000000000348c3d16aaceb6aec29a33b0000000000000000000000000000000000000000034e0e835960574ff2229eb1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882851", 0)),
							TransactionHash: common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
							Index:           515,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000790c9422839fd93a3a4e31e531f96cc87f397c00"),
								common.HexToHash("0x00000000000000000000000023878914efe38d27c4d67ab83ed1b93a74d4086a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000b09949e180"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882851", 0)),
							TransactionHash: common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
							Index:           516,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
							Topics: []common.Hash{
								common.HexToHash("0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051"),
								common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
								common.HexToHash("0x000000000000000000000000790c9422839fd93a3a4e31e531f96cc87f397c00"),
								common.HexToHash("0x000000000000000000000000790c9422839fd93a3a4e31e531f96cc87f397c00"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000b09949e1800000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17882851", 0)),
							TransactionHash: common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
							Index:           517,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0"),
						TransactionIndex: 158,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkEthereum,
					Endpoint: endpoint.MustGet(filter.NetworkEthereum),
				},
			},
			want: &schema.Feed{
				ID:      "0xdc2cd838127a068db79c70508ac87b73d083a2ea22d230ae6f1d8b27001e58d0",
				Network: filter.NetworkEthereum,
				Index:   158,
				From:    "0x790c9422839FD93a3A4E31e531f96cC87F397c00",
				To:      aave.AddressV3PoolMainnet.String(),
				Type:    filter.TypeExchangeLiquidity,
				Calldata: &schema.Calldata{
					FunctionHash: "0x573ade81",
				},
				Platform: lo.ToPtr(filter.PlatformAAVE),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("2571456723342402")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeExchangeLiquidity,
						Platform: filter.PlatformAAVE.String(),
						From:     "0x790c9422839FD93a3A4E31e531f96cC87F397c00",
						To:       aave.AddressV3PoolMainnet.String(),
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRepay,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("758486000000"))),
									Name:     "Tether USD",
									Symbol:   "USDT",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691650679,
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
