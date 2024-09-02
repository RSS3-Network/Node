package arbitrum_test

import (
	"context"
	"encoding/json"
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
						Hash:         common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
						ParentHash:   common.HexToHash("0x4ad4c3a5e7c9e9d1424498d297b03b1aa766a870de2a000bc746f62e03d7ba14"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("19775068", 0)),
						GasLimit:     30000000,
						GasUsed:      20306340,
						Timestamp:    1714563959,
						BaseFee:      lo.Must(new(big.Int).SetString("7389489056", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
						From:      common.HexToAddress("0xE43D08D35e55Ab46E89352Bd2526aFa4447D2C10"),
						Gas:       97810,
						GasPrice:  lo.Must(new(big.Int).SetString("7401129056", 10)),
						Hash:      common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
						Input:     hexutil.MustDecode("0x439370b1"),
						To:        lo.ToPtr(common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f")),
						Value:     lo.Must(new(big.Int).SetString("300000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa87ffa74eacfbfec0063f0eebce583e563e8a212c50ebcad71ddabad90539529"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19775068", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 12531054,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x1b9244460"),
						GasUsed:           91261,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x8315177aB297bA92A06054cE80a67Ed4DBd7ed3a"),
							Topics: []common.Hash{
								common.HexToHash("0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e5d5"),
								common.HexToHash("0x60b59f12ab85a33feec31d7306659e923622ba0212e56c3dff5267519748a8c5"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000004dbd4fc535ac27206064b68ffcf827b0a60bab3f000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000f54e08d35e55ab46e89352bd2526afa4447d3d21ad4b2c5c1911e7ae28dcde05543eeb9d86862d51e4223228f11ce526b54ee8a900000000000000000000000000000000000000000000000000000001b872a7a00000000000000000000000000000000000000000000000000000000066322b77"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19775068", 0)),
							TransactionHash: common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
							Index:           304,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f"),
							Topics: []common.Hash{
								common.HexToHash("0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e5d5"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000034e43d08d35e55ab46e89352bd2526afa4447d2c100000000000000000000000000000000000000000000000000429d069189e0000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19775068", 0)),
							TransactionHash: common.HexToHash("0x0b590dbd9b4e8887f91c6d3c4553e832e1ac8f47b18e9a881005cf6d1849e180"),
							Index:           305,
							Removed:         false,
						}},
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
				From:     "0xE43D08D35e55Ab46E89352Bd2526aFa4447D2C10",
				To:       "0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f",
				Type:     typex.TransactionBridge,
				Platform: workerx.PlatformArbitrum.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("675434438779616")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x439370b1",
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
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("300000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1714563959,
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
						Hash:         common.HexToHash("0xb25b2d57c631bebe4592117d9dd8fa451bc56e80da90f072816ed3432fd1547e"),
						ParentHash:   common.HexToHash("0x76d301cb6f5f718ba388d93270c1744970ac501d1162faac758bf103f0745114"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("19777199", 0)),
						GasLimit:     30000000,
						GasUsed:      18444425,
						Timestamp:    1714589615,
						BaseFee:      lo.Must(new(big.Int).SetString("19464415475", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
						From:      common.HexToAddress("0xA3DE4Bc4Cd34E591b08bf9d0378581c5416fA66A"),
						Gas:       227247,
						GasPrice:  lo.Must(new(big.Int).SetString("19473268403", 10)),
						Hash:      common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
						Input:     hexutil.MustDecode("0xd2ce7d65000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000001312d000000000000000000000000000000000000000000000000000000000000043238000000000000000000000000000000000000000000000000000000000b92bcf000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000e6f4e6851b8000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef")),
						Value:     lo.Must(new(big.Int).SetString("307335163891712", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xb25b2d57c631bebe4592117d9dd8fa451bc56e80da90f072816ed3432fd1547e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19777199", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 11513423,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x488b27eb3"),
						GasUsed:           199879,

						Logs: []*ethereum.Log{
							{
								Address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
								Topics: []common.Hash{
									common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000cee284f754e854890e311e3280b767f80797180d"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000001312d00"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19777199", 0)),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           289,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0x8315177aB297bA92A06054cE80a67Ed4DBd7ed3a"),
								Topics: []common.Hash{
									common.HexToHash("0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
									common.HexToHash("0x277ab7dc33f766c4c919e05811bbdb3f6526d202c78a8b7ba0793661fc4b54b0"),
								},
								Data:            hexutil.MustDecode("0x0000000000000000000000004dbd4fc535ac27206064b68ffcf827b0a60bab3f0000000000000000000000000000000000000000000000000000000000000009000000000000000000000000dff384f754e854890e311e3280b767f80797291ea0968ecc6b568bf83b8edcb18fcae655f8f18eadc7e29600d6ca42fe27a40a2100000000000000000000000000000000000000000000000000000004882b68f30000000000000000000000000000000000000000000000000000000066328faf"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19777199", 0)),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           290,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f"),
								Topics: []common.Hash{
									common.HexToHash("0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
								},
								Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000264000000000000000000000000096760f208390250649e3e8763348e783aef55620000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000117850b4550000000000000000000000000000000000000000000000000000000e6f4e6851b80000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000000043238000000000000000000000000000000000000000000000000000000000b92bcf000000000000000000000000000000000000000000000000000000000000001442e567b36000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000001312d0000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19777199", 0)),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           291,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0xcEe284F754E854890e311e3280b767F80797180d"),
								Topics: []common.Hash{
									common.HexToHash("0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000096760f208390250649e3e8763348e783aef5562"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001442e567b36000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a0000000000000000000000000000000000000000000000000000000001312d0000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19777199", 0)),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           292,
								Removed:         false,
							}, {
								Address: common.HexToAddress("0xcEe284F754E854890e311e3280b767F80797180d"),
								Topics: []common.Hash{
									common.HexToHash("0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000a3de4bc4cd34e591b08bf9d0378581c5416fa66a"),
									common.HexToHash("0x000000000000000000000000000000000000000000000000000000000016e738"),
								},
								Data:            hexutil.MustDecode("0x000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000001312d00"),
								BlockNumber:     lo.Must(new(big.Int).SetString("19777199", 0)),
								TransactionHash: common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
								Index:           293,
								Removed:         false,
							}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x39aa325d33222ab596525fb382e8a775b6b661c4fb9471e11f6b34a9478568e8"),
						TransactionIndex: 157,
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
				From:     "0xA3DE4Bc4Cd34E591b08bf9d0378581c5416fA66A",
				To:       "0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef",
				Type:     typex.TransactionBridge,
				Platform: workerx.PlatformArbitrum.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3892297415123237")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0xd2ce7d65",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformArbitrum.String(),
						From:     "0xdFf384F754E854890E311e3280B767F80797291e",
						To:       "0xdFf384F754E854890E311e3280B767F80797291e",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Ethereum,
							TargetNetwork: network.Arbitrum,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("307335163891712"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
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
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("307335163891712"))),
								Name:     "Tether USD",
								Symbol:   "USDT",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1714589615,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw USDC from Arbitrum to Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xad8e86c014493b90cb2882ab41a8ebf202b866b7895bfb9241160bc59e05ac4c"),
						ParentHash:   common.HexToHash("0xff8c81f1c81c6b32b1d99d9458b400becb47437825a99c983da672a5dabcf7d2"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x388C818CA8B9251b393131C08a736A67ccB19297"),
						Number:       lo.Must(new(big.Int).SetString("20648434", 0)),
						GasLimit:     30000000,
						GasUsed:      9933802,
						Timestamp:    1725105851,
						BaseFee:      lo.Must(new(big.Int).SetString("667809446", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
						From:      common.HexToAddress("0x9136D8E0F5A09729b8A9dEDF4B5c7C0b1EF6365d"),
						Gas:       500000,
						GasPrice:  lo.Must(new(big.Int).SetString("2667809446", 10)),
						Hash:      common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
						Input:     hexutil.MustDecode("0x08635a950000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000001f0a5000000000000000000000000096760f208390250649e3e8763348e783aef5562000000000000000000000000cee284f754e854890e311e3280b767f80797180d000000000000000000000000000000000000000000000000000000000eacd7a000000000000000000000000000000000000000000000000000000000013a4e1f0000000000000000000000000000000000000000000000000000000066c9cab60000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000036000000000000000000000000000000000000000000000000000000000000000110eb071d07bc82af7b71eaede3ecca2d3dd8d7c502cd850db37a597879a63ef7c1c1307571071115ddcd456e9fa1d2682764da0b6c65244f6197a27c203745b3d2e26fed1dbe5e5f29f47110a4b1c2b750730bd1ea57bcdce67b1822061d7001e4ecab25c8a4a4a3ed26d5218f69df84922805b19f3d07c7d16a1ee057e051c0478cc1adc63f247777522642f89c03ed01ba26141dfc56e6b9c6ac5b293a217b01fe9cf1bb89f4a8d9c6563f41bcdc7f72dde80763362735f1dfe48b376a336cebbf06c742146f2cb138d57fb90a5e306b1d1e3e7aa84c8e1055099e30a1404af8d302602268bebec402c41dfebb77c138b4112d47603adde3864b7d5eb1fc728000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060b7b72ebdcd016b750724760612db47d8d37b066cdbc42f5458029d71df2badf2ba2a36f24697a4a158278feea7c61d80a7095422e10a4ef78b94c9ed096508e293e3dc5befe34c72d718dc7b2e5f5cfdeb11a8c3387f0811a76c1d8825903cc0425084107ea9f7a4118f5ed1e3566cda4e90b550363fc804df1e52ed5f2386b43a6b28077d49f37d58c87aec0b51f7bce13b648143f3295385f3b3d5ac3b9b00000000000000000000000000000000000000000000000000000000000001242e567b36000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d0000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d0000000000000000000000000000000000000000000000000000002266180d4e00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000044230000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x0B9857ae2D4A3DBe74ffE1d7DF045bb7F96E4840")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xad8e86c014493b90cb2882ab41a8ebf202b866b7895bfb9241160bc59e05ac4c"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20648434", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 919649,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x9f038aa6"),
						GasUsed:           178357,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x0B9857ae2D4A3DBe74ffE1d7DF045bb7F96E4840"),
							Topics: []common.Hash{
								common.HexToHash("0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964"),
								common.HexToHash("0x000000000000000000000000cee284f754e854890e311e3280b767f80797180d"),
								common.HexToHash("0x000000000000000000000000096760f208390250649e3e8763348e783aef5562"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000001f0a5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20648434", 0)),
							TransactionHash: common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
							Index:           12,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000cee284f754e854890e311e3280b767f80797180d"),
								common.HexToHash("0x0000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000002266180d4e"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20648434", 0)),
							TransactionHash: common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
							Index:           13,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xcEe284F754E854890e311e3280b767F80797180d"),
							Topics: []common.Hash{
								common.HexToHash("0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3"),
								common.HexToHash("0x0000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d"),
								common.HexToHash("0x0000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000004423"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000002266180d4e"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20648434", 0)),
							TransactionHash: common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
							Index:           14,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x8315177aB297bA92A06054cE80a67Ed4DBd7ed3a"),
							Topics: []common.Hash{
								common.HexToHash("0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466"),
								common.HexToHash("0x0000000000000000000000000b9857ae2d4a3dbe74ffe1d7df045bb7f96e4840"),
								common.HexToHash("0x000000000000000000000000cee284f754e854890e311e3280b767f80797180d"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001242e567b36000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d0000000000000000000000009136d8e0f5a09729b8a9dedf4b5c7c0b1ef6365d0000000000000000000000000000000000000000000000000000002266180d4e00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000044230000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20648434", 0)),
							TransactionHash: common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
							Index:           15,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624"),
						TransactionIndex: 18,
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
				ID:       "0x902724b37aa1ec66c21ab2f5a7769ad332d2061b0a748ea1b51b39d5f97c5624",
				Network:  network.Ethereum,
				Index:    18,
				Platform: workerx.PlatformArbitrum.String(),
				From:     "0x9136D8E0F5A09729b8A9dEDF4B5c7C0b1EF6365d",
				To:       "0x0B9857ae2D4A3DBe74ffE1d7DF045bb7F96E4840",
				Type:     typex.TransactionBridge,
				Status:   true,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x08635a95",
				},
				Timestamp: 1725105851,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("475822489360222")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformArbitrum.String(),
						From:     "0x9136D8E0F5A09729b8A9dEDF4B5c7C0b1EF6365d",
						To:       "0x9136D8E0F5A09729b8A9dEDF4B5c7C0b1EF6365d",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Arbitrum,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("147741740366"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw ARB from Arbitrum to Ethereum",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 42161,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x7e22efbf9570e9bf21695248c751de8479f58c1162485a8fb5fc9fbae78b6175"),
						ParentHash:   common.HexToHash("0x3e8af39ae25740a3d1b45ad53909ee84c2466c79963998f90162a8e02d31e56e"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xA4b000000000000000000073657175656e636572"),
						Number:       lo.Must(new(big.Int).SetString("100260539", 0)),
						GasLimit:     1125899906842624,
						GasUsed:      614919,
						Timestamp:    1686538396,
						BaseFee:      lo.Must(new(big.Int).SetString("100000000", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x974e267c3601b535d2f7965c24473ba500a495450bbf3822569ad18d929b752f"),
						From:      common.HexToAddress("0x2789f0554679F1f58491236196ECf8D8cc94c91C"),
						Gas:       724081,
						GasPrice:  lo.Must(new(big.Int).SetString("100000000", 10)),
						Hash:      common.HexToHash("0x974e267c3601b535d2f7965c24473ba500a495450bbf3822569ad18d929b752f"),
						Input:     hexutil.MustDecode("0x7b3a3c8b000000000000000000000000b50721bcf8d664c30412cfbc6cf7a15145234ad10000000000000000000000002789f0554679f1f58491236196ecf8d8cc94c91c000000000000000000000000000000000000000000000004eee0f1dc808c000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x5288c571Fd7aD117beA99bF60FE0846C4E84F933")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("42161", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x7e22efbf9570e9bf21695248c751de8479f58c1162485a8fb5fc9fbae78b6175"),
						BlockNumber:       lo.Must(new(big.Int).SetString("100260539", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 614919,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x5f5e100"),
						GasUsed:           614919,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x0000000000000000000000000000000000000064"),
							Topics: []common.Hash{
								common.HexToHash("0x3e7aafa77dbf186b7fd488006beff893744caa3c4f6f299e8a709fa2087374fc"),
								common.HexToHash("0x000000000000000000000000bbce8aa77782f13d4202a230d978f361b011db27"),
								common.HexToHash("0xea45bd538103218f8a1a43aae971ec8530336ebc16cccf9ea60f8107d50561cf"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013892"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000cad7828a19b363a2b44717afb1786b5196974d8e0000000000000000000000000000000000000000000000000000000005f9dabb00000000000000000000000000000000000000000000000000000000010a6f92000000000000000000000000000000000000000000000000000000006486889c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001242e567b36000000000000000000000000b50721bcf8d664c30412cfbc6cf7a15145234ad10000000000000000000000002789f0554679f1f58491236196ecf8d8cc94c91c0000000000000000000000002789f0554679f1f58491236196ecf8d8cc94c91c000000000000000000000000000000000000000000000004eee0f1dc808c000000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000002d30000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("100260539", 0)),
							TransactionHash: common.HexToHash("0x974e267c3601b535d2f7965c24473ba500a495450bbf3822569ad18d929b752f"),
							Index:           3,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xCaD7828a19b363A2B44717AFB1786B5196974D8E"),
							Topics: []common.Hash{
								common.HexToHash("0x3073a74ecb728d10be779fe19a74a1428e20468f5b4d167bf9c73d9067847d73"),
								common.HexToHash("0x0000000000000000000000002789f0554679f1f58491236196ecf8d8cc94c91c"),
								common.HexToHash("0x0000000000000000000000002789f0554679f1f58491236196ecf8d8cc94c91c"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000013892"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000b50721bcf8d664c30412cfbc6cf7a15145234ad100000000000000000000000000000000000000000000000000000000000002d3000000000000000000000000000000000000000000000004eee0f1dc808c0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("100260539", 0)),
							TransactionHash: common.HexToHash("0x974e267c3601b535d2f7965c24473ba500a495450bbf3822569ad18d929b752f"),
							Index:           5,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x974e267c3601b535d2f7965c24473ba500a495450bbf3822569ad18d929b752f"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network: network.Arbitrum,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Arbitrum),
					},
				},
			},

			want: &activityx.Activity{
				ID:       "0x974e267c3601b535d2f7965c24473ba500a495450bbf3822569ad18d929b752f",
				Network:  network.Ethereum,
				Index:    1,
				From:     "0x2789f0554679F1f58491236196ECf8D8cc94c91C",
				To:       "0x5288c571Fd7aD117beA99bF60FE0846C4E84F933",
				Type:     typex.TransactionBridge,
				Platform: workerx.PlatformArbitrum.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("61491900000000")),
					Decimal: 18,
				},
				Calldata: &activityx.Calldata{
					FunctionHash: "0x7b3a3c8b",
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformArbitrum.String(),
						From:     "0xCaD7828a19b363A2B44717AFB1786B5196974D8E",
						To:       "0xbbcE8aA77782F13D4202a230d978F361B011dB27",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Arbitrum,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("0"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformArbitrum.String(),
						From:     "0x2789f0554679F1f58491236196ECf8D8cc94c91C",
						To:       "0x2789f0554679F1f58491236196ECf8D8cc94c91C",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Arbitrum,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x912CE59144191C1204E64559FE8253a0e49E6548"),
								Name:     "Arbitrum",
								Symbol:   "ARB",
								Decimals: 18,
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("91000000000000000000"))),
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1686538396,
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
