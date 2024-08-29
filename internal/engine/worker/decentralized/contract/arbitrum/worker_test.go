package arbitrum_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/arbitrum"
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

func TestWorker_Arbitrum(t *testing.T) {
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
			name: "Deposit ETH from Ethereum to Arbitrum",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:       common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
						ParentHash: common.HexToHash("0x4ad4c3a5e7c9e9d1424498d297b03b1aa766a870de2a000bc746f62e03d7ba14"),
						UncleHash:  common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:   common.HexToAddress("0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97"),
						Number:     big.NewInt(19773020),
						GasLimit:   30000000,
						GasUsed:    20381092,
						Timestamp:  1714651005,
						BaseFee:    big.NewInt(7389489056),
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
						From:      common.HexToAddress("0xE43D08D35e55Ab46E89352Bd2526aFa4447D2C10"),
						Gas:       97810,
						GasPrice:  lo.Must(new(big.Int).SetString("7389489056", 10)),
						Hash:      common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
						Input:     hexutil.MustDecode("0x439370b1"),
						To:        lo.ToPtr(common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f")),
						Value:     big.NewInt(300000000000000000),
						Type:      2,
						ChainID:   big.NewInt(1),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
						BlockNumber:       big.NewInt(19773020),
						ContractAddress:   nil,
						CumulativeGasUsed: 12537198,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3d786f7ac"),
						GasUsed:           91261,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x8315177ab297ba92a06054ce80a67ed4dbd7ed3a"),
								Topics: []common.Hash{
									common.HexToHash("0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e5d5"),
									common.HexToHash("0x60b59f12ab85a33feec31d7306659e923622ba0212e56c3dff5267519748a8c5"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000004dbd4fc535ac27206064b68ffcf827b0a60bab3f000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000f54e08d35e55ab46e89352bd2526afa4447d3d21ad4b2c5c1911e7ae28dcde05543eeb9d86862d51e4223228f11ce526b54ee8a900000000000000000000000000000000000000000000000000000001b872a7a00000000000000000000000000000000000000000000000000000000066322b77"),
								BlockNumber:     big.NewInt(19773020),
								TransactionHash: common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
								//TxIndex:     123,
								BlockHash: common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
								Index:     304,
								Removed:   false,
							},
							{
								Address: common.HexToAddress("0x4dbd4fc535ac27206064b68ffcf827b0a60bab3f"),
								Topics: []common.Hash{
									common.HexToHash("0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e5d5"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000034e43d08d35e55ab46e89352bd2526afa4447d2c100000000000000000000000000000000000000000000000000429d069189e0000000000000000000000000000"),
								BlockNumber:     big.NewInt(19773020),
								TransactionHash: common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
								//TxIndex:     123,
								BlockHash: common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
								Index:     305,
								Removed:   false,
							},
						},
						Status:           1,
						TransactionHash:  common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
						TransactionIndex: 123,
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180",
				Network:  network.Ethereum,
				Index:    123,
				Platform: "Arbitrum",
				From:     "0xE43D08D35e55Ab46E89352Bd2526aFa4447D2C10",
				To:       "0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f",
				Type:     typex.TransactionBridge,
				Status:   true,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x439370b1",
				},
				Timestamp: 1714651005,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1505883940069116")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformArbitrum.String(),
						From:     "0xf54E08D35E55AB46e89352bD2526AfA4447D3D21",
						To:       "0xf54E08D35E55AB46e89352bD2526AfA4447D3D21",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Arbitrum,
							Token: metadata.Token{
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("300000000000000000"))),
							},
						},
					},
				},
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit USDT from Ethereum to Arbitrum",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:       common.HexToHash("0xb25b2d57c631bebe4592117d9dd8fa451bc56e80da90f072816ed3432fd1547e"),
						ParentHash: common.HexToHash("0x76d301cb6f5f718ba388d93270c1744970ac501d1162faac758bf103f0745114"),
						UncleHash:  common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:   common.HexToAddress("0x1f9090aae28b8a3dceadf281b0f12828e676c326"),
						Number:     big.NewInt(19773103),
						GasLimit:   30000000,
						GasUsed:    18366601,
						Timestamp:  1714659247,
						BaseFee:    big.NewInt(19402891507),
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xb25b2d57c631bebe4592117d9dd8fa451bc56e80da90f072816ed3432fd1547e"),
						From:      common.HexToAddress("0xA3DE4Bc4Cd34E591b08bf9d0378581c5416fA66A"),
						Gas:       227759,
						GasPrice:  lo.Must(new(big.Int).SetString("19489399987", 10)),
						Hash:      common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
						Input:     hexutil.MustDecode("0xd2ce7d65000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000001312d000000000000000000000000000000000000000000000000000000000000043238000000000000000000000000000000000000000000000000000000000b92bcf000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000e6f4e6851b8000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef")),
						Value:     big.NewInt(1200000000000000),
						Type:      2,
						ChainID:   big.NewInt(1),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xb25b2d57c631bebe4592117d9dd8fa451bc56e80da90f072816ed3432fd1547e"),
						BlockNumber:       big.NewInt(19773103),
						ContractAddress:   nil,
						CumulativeGasUsed: 11529807,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x488b27eb3"),
						GasUsed:           199879,
						Status:            1,
						TransactionHash:   common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
						TransactionIndex:  157,
						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0x72ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef"),
								Topics: []common.Hash{
									common.HexToHash("0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5"),
									common.HexToHash("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000cee284f754e854890e311e3280b767f80797180d"),
								BlockNumber:     big.NewInt(19773103),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           304,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000cee284f754e854890e311e3280b767f80797180d"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000001312d00"),
								BlockNumber:     big.NewInt(19773103),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           305,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x8315177ab297ba92a06054ce80a67ed4dbd7ed3a"),
								Topics: []common.Hash{
									common.HexToHash("0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
									common.HexToHash("0x277ab7dc33f766c4c919e05811bbdb3f6526d202c78a8b7ba0793661fc4b54b0"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000004dbd4fc535ac27206064b68ffcf827b0a60bab3f0000000000000000000000000000000000000000000000000000000000000009000000000000000000000000dff384f754e854890e311e3280b767f80797291ea0968ecc6b568bf83b8edcb18fcae655f8f18eadc7e29600d6ca42fe27a40a2100000000000000000000000000000000000000000000000000000004882b68f30000000000000000000000000000000000000000000000000000000066328faf"),
								BlockNumber:     big.NewInt(19773103),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           306,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0x4dbd4fc535ac27206064b68ffcf827b0a60bab3f"),
								Topics: []common.Hash{
									common.HexToHash("0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000264000000000000000000000000096760f208390250649e3e8763348e783aef55620000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000117850b4550000000000000000000000000000000000000000000000000000000e6f4e6851b80000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000000043238000000000000000000000000000000000000000000000000000000000b92bcf000000000000000000000000000000000000000000000000000000000000001442e567b36000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000001312d0000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     big.NewInt(19773103),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           307,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xcee284f754e854890e311e3280b767f80797180d"),
								Topics: []common.Hash{
									common.HexToHash("0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000096760f208390250649e3e8763348e783aef5562"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001442e567b36000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000001312d0000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     big.NewInt(19773103),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           308,
								Removed:         false,
							},
							{
								Address: common.HexToAddress("0xcee284f754e854890e311e3280b767f80797180d"),
								Topics: []common.Hash{
									common.HexToHash("0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000001312d00"),
								BlockNumber:     big.NewInt(19773103),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           309,
								Removed:         false,
							},
						},
					},
				},
				config: &config.Module{
					Network: network.Ethereum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Ethereum),
					},
				},
			},
			want: &activityx.Activity{
				ID:       "0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8",
				Network:  network.Ethereum,
				Index:    157,
				Platform: "Arbitrum",
				From:     "0xA3DE4Bc4Cd34E591b08bf9d0378581c5416fA66A",
				To:       "0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef",
				Type:     typex.TransactionBridge,
				Status:   true,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xd2ce7d65",
				},
				Timestamp: 1714659247,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3892297415123237")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformArbitrum.String(),
						From:     "0xA3DE4Bc4Cd34E591b08bf9d0378581c5416fA66A",
						To:       "0xA3DE4Bc4Cd34E591b08bf9d0378581c5416fA66A",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Arbitrum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1200000000000000"))),
								Name:     "Tether USD",
								Symbol:   "USDT",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
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

			require.Equal(t, testcase.want, activity)
		})
	}
}
