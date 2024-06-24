package curve_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/curve"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/curve"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	redisx "github.com/rss3-network/node/provider/redis"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
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
		want      *activityx.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Swap ETH for WBTC on Registry Exchange",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x08efe625d4ee81b07b9159e28138cf2f198a8e0e5cbb7f548a11bd585ff204c0"),
						ParentHash:   common.HexToHash("0x6e691b2e2168e4555ef3da8ce521463e82e20e08fe75f1af841d876ec8a4014b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17925088", 0)),
						GasLimit:     29999943,
						GasUsed:      13343377,
						Timestamp:    1692161063,
						BaseFee:      lo.Must(new(big.Int).SetString("19536863916", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
						From:      common.HexToAddress("0x78d72460De403CB1650Bbd0476256539D61B25b2"),
						Gas:       273044,
						GasPrice:  lo.Must(new(big.Int).SetString("20536863916", 10)),
						Hash:      common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
						Input:     hexutil.MustDecode("0x9db4f7aa000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee0000000000000000000000007f86bf177dd4f3494b841a37e810a34dd56c829b0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c5990000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001bc16d674ec800000000000000000000000000000000000000000000000000000000000000bf0b480000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x99a58482BD75cbab83b27EC03CA68fF489b5788f")),
						Value:     lo.Must(new(big.Int).SetString("2000000000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x08efe625d4ee81b07b9159e28138cf2f198a8e0e5cbb7f548a11bd585ff204c0"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17925088", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 6885925,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x4c817acac"),
						GasUsed:           187334,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000007f86bf177dd4f3494b841a37e810a34dd56c829b"),
								common.HexToHash("0x00000000000000000000000099a58482bd75cbab83b27ec03ca68ff489b5788f"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000bf0dba"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17925088", 0)),
							TransactionHash: common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
							Index:           229,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7F86Bf177Dd4F3494b841a37e810A34dD56c829B"),
							Topics: []common.Hash{
								common.HexToHash("0x143f1f8e861fbdeddd5b46e844b7d3ac7b86a122f36e8c463859ee6811b1f29c"),
								common.HexToHash("0x00000000000000000000000099a58482bd75cbab83b27ec03ca68ff489b5788f"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000001bc16d674ec8000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000bf0dba0000000000000000000000000000000000000000000000000000000000001a36000000000000006398424f9ddf5b03190000000000000634e9f0721944954dcd"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17925088", 0)),
							TransactionHash: common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
							Index:           230,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000099a58482bd75cbab83b27ec03ca68ff489b5788f"),
								common.HexToHash("0x00000000000000000000000078d72460de403cb1650bbd0476256539d61b25b2"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000bf0dba"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17925088", 0)),
							TransactionHash: common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
							Index:           231,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x99a58482BD75cbab83b27EC03CA68fF489b5788f"),
							Topics: []common.Hash{
								common.HexToHash("0x14b561178ae0f368f40fafd0485c4f7129ea71cdc00b4ce1e5940f9bc659c8b2"),
								common.HexToHash("0x00000000000000000000000078d72460de403cb1650bbd0476256539d61b25b2"),
								common.HexToHash("0x00000000000000000000000078d72460de403cb1650bbd0476256539d61b25b2"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee0000000000000000000000007f86bf177dd4f3494b841a37e810a34dd56c829b0000000000000000000000002260fac5e5542a773aa44fbcfedf7c193bc2c59900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001bc16d674ec800000000000000000000000000000000000000000000000000000000000000bf0dba"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17925088", 0)),
							TransactionHash: common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
							Index:           232,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32"),
						TransactionIndex: 68,
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
				ID:      "0xe019155e4cc243b605397d56bb88bd83f2c6f002eccb8e3056ded3a2e2240d32",
				Network: network.Ethereum,
				Index:   68,
				From:    "0x78d72460De403CB1650Bbd0476256539D61B25b2",
				To:      curve.AddressRegistryExchangeEthereum.String(),
				Type:    typex.ExchangeSwap,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x9db4f7aa",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("3847252864839944")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x78d72460De403CB1650Bbd0476256539D61B25b2",
						To:       "0x78d72460De403CB1650Bbd0476256539D61B25b2",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2000000000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("12520890"))),
								Name:     "Wrapped BTC",
								Symbol:   "WBTC",
								Decimals: 8,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692161063,
			},
			wantError: require.NoError,
		},
		{
			name: "Swap DAI for USDT on StableSwap 3Pool",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x3a531f8ba1aefde11738cff16b41012e6e355e47c0c0bc90fcf2defc085e55c6"),
						ParentHash:   common.HexToHash("0x4c2838085629bb606f2a10a08eb37e736ffca4263397d621164c253d05e80f7c"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("17926977", 0)),
						GasLimit:     30000000,
						GasUsed:      20909099,
						Timestamp:    1692183911,
						BaseFee:      lo.Must(new(big.Int).SetString("22343321601", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0"),
						From:      common.HexToAddress("0xb0A79185B7640D6a3DbbDd1072A70D8f82f9095E"),
						Gas:       193055,
						GasPrice:  lo.Must(new(big.Int).SetString("23343321601", 10)),
						Hash:      common.HexToHash("0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0"),
						Input:     hexutil.MustDecode("0x3df0212400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000006a0497ce94801c179d9900000000000000000000000000000000000000000000000000000074ab0eab6c"),
						To:        lo.ToPtr(common.HexToAddress("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x3a531f8ba1aefde11738cff16b41012e6e355e47c0c0bc90fcf2defc085e55c6"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17926977", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8477663,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x56f5ed201"),
						GasUsed:           121603,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b0a79185b7640d6a3dbbdd1072a70d8f82f9095e"),
								common.HexToHash("0x000000000000000000000000bebc44782c7db0a1a60cb6fe97d0b483032ff1c7"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000006a0497ce94801c179d99"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17926977", 0)),
							TransactionHash: common.HexToHash("0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0"),
							Index:           197,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000bebc44782c7db0a1a60cb6fe97d0b483032ff1c7"),
								common.HexToHash("0x000000000000000000000000b0a79185b7640d6a3dbbdd1072a70d8f82f9095e"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000074b40525e1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17926977", 0)),
							TransactionHash: common.HexToHash("0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0"),
							Index:           198,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7"),
							Topics: []common.Hash{
								common.HexToHash("0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140"),
								common.HexToHash("0x000000000000000000000000b0a79185b7640d6a3dbbdd1072a70d8f82f9095e"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006a0497ce94801c179d99000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000074b40525e1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17926977", 0)),
							TransactionHash: common.HexToHash("0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0"),
							Index:           199,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0"),
						TransactionIndex: 41,
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
				ID:      "0x31e337fb6eba7d49b9cf7107497e27c2cfa25c572f098380a530fff4538e80c0",
				Network: network.Ethereum,
				Index:   41,
				From:    "0xb0A79185B7640D6a3DbbDd1072A70D8f82f9095E",
				To:      "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7",
				Type:    typex.ExchangeSwap,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x3df02124",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2838617936646403")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformCurve.String(),
						From:     "0xb0A79185B7640D6a3DbbDd1072A70D8f82f9095E",
						To:       "0xb0A79185B7640D6a3DbbDd1072A70D8f82f9095E",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("500655573004300108406169"))),
								Name:     "Dai Stablecoin",
								Symbol:   "DAI",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("501236442593"))),
								Name:     "Tether USD",
								Symbol:   "USDT",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692183911,
			},
			wantError: require.NoError,
		},
		{
			name: "Add USDC liquidity to StableSwap 3Pool",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xee4ec575a9c6176586d731d8c265bb9bbd0ea3af9519c1e8e3d71a1b423d9dd1"),
						ParentHash:   common.HexToHash("0x0b10a07478d399c833386eb0b7f914346e01a1946e3ba70e55eb47cdcf3eef49"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990"),
						Number:       lo.Must(new(big.Int).SetString("17919359", 0)),
						GasLimit:     30000000,
						GasUsed:      29513471,
						Timestamp:    1692091811,
						BaseFee:      lo.Must(new(big.Int).SetString("13879476397", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641"),
						From:      common.HexToAddress("0x19BA1abE64245f0a9eB12B9efCC94167f9a05064"),
						Gas:       268739,
						GasPrice:  lo.Must(new(big.Int).SetString("14879476397", 10)),
						Hash:      common.HexToHash("0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641"),
						Input:     hexutil.MustDecode("0x4515cef300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000035a4e900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002f86e6a4377cf1c1f7"),
						To:        lo.ToPtr(common.HexToAddress("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xee4ec575a9c6176586d731d8c265bb9bbd0ea3af9519c1e8e3d71a1b423d9dd1"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17919359", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 21173065,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x376e2caad"),
						GasUsed:           184622,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000019ba1abe64245f0a9eb12b9efcc94167f9a05064"),
								common.HexToHash("0x000000000000000000000000bebc44782c7db0a1a60cb6fe97d0b483032ff1c7"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000035a4e900"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919359", 0)),
							TransactionHash: common.HexToHash("0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641"),
							Index:           474,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000019ba1abe64245f0a9eb12b9efcc94167f9a05064"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000002f8a8d5701f584689a"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919359", 0)),
							TransactionHash: common.HexToHash("0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641"),
							Index:           475,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7"),
							Topics: []common.Hash{
								common.HexToHash("0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d"),
								common.HexToHash("0x00000000000000000000000019ba1abe64245f0a9eb12b9efcc94167f9a05064"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000035a4e90000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013fcbed484241c0000000000000000000000000000000000000000000000000000000000006d4e0000000000000000000000000000000000000000000000000000000000005773000000000000000000000000000000000000000000b7783e887a66a577c051bd000000000000000000000000000000000000000000b2a6696981fde5cd311c4e"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919359", 0)),
							TransactionHash: common.HexToHash("0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641"),
							Index:           476,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641"),
						TransactionIndex: 514,
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
				ID:      "0xd68b76ccb5692b43fdac871806b1c989602ee8d04b884a73c745966f30106641",
				Network: network.Ethereum,
				Index:   514,
				From:    "0x19BA1abE64245f0a9eB12B9efCC94167f9a05064",
				To:      "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7",
				Type:    typex.ExchangeLiquidity,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x4515cef3",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("2747078691366934")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeLiquidity,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x19BA1abE64245f0a9eB12B9efCC94167f9a05064",
						To:       "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityAdd,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("900000000"))),
									Name:     "USD Coin",
									Symbol:   "USDC",
									Decimals: 6,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
					{
						Type:     typex.TransactionMint,
						Platform: workerx.PlatformCurve.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0x19BA1abE64245f0a9eB12B9efCC94167f9a05064",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("876980703079224862874"))),
							Name:     "Curve.fi DAI/USDC/USDT",
							Symbol:   "3Crv",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1692091811,
			},
			wantError: require.NoError,
		},
		{
			name: "Add USDC liquidity to StableSwap 3Pool Deposit",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xf83c944ea55e41e4a8f786b5bdb0c8c2d1e5735b18041d9313b70830ff0b5bd3"),
						ParentHash:   common.HexToHash("0x7ad1d26468f2866be51f7b865eb05bb4ed111a65130ed187e5208d9b009c52d6"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"),
						Number:       lo.Must(new(big.Int).SetString("19879519", 0)),
						GasLimit:     30000000,
						GasUsed:      18295452,
						Timestamp:    1715826035,
						BaseFee:      lo.Must(new(big.Int).SetString("3427682187", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
						From:      common.HexToAddress("0x7664f4A2d8c663d713170d0AB25B7C15d06B84a1"),
						Gas:       433342,
						GasPrice:  lo.Must(new(big.Int).SetString("3587627803", 10)),
						Hash:      common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
						Input:     hexutil.MustDecode("0x4515cef3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008c640b185e89a000000000000000000000000000000000000000000000000000e71830a80edcc"),
						To:        lo.ToPtr(common.HexToAddress("0x3993d34e7e99Abf6B6f367309975d1360222D446")),
						Value:     lo.Must(new(big.Int).SetString("2469780972234906", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xf83c944ea55e41e4a8f786b5bdb0c8c2d1e5735b18041d9313b70830ff0b5bd3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19879519", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 18166461,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xd5d6db1b"),
						GasUsed:           272430,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x0000000000000000000000003993d34e7e99abf6b6f367309975d1360222d446"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000008c640b185e89a"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19879519", 0)),
							TransactionHash: common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
							Index:           553,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000003993d34e7e99abf6b6f367309975d1360222d446"),
								common.HexToHash("0x000000000000000000000000d51a44d3fae010294c616388b506acda1bfaae46"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000008c640b185e89a"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19879519", 0)),
							TransactionHash: common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
							Index:           554,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xc4AD29ba4B3c580e6D59105FFf484999997675Ff"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000003993d34e7e99abf6b6f367309975d1360222d446"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000e75368d346ff5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19879519", 0)),
							TransactionHash: common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
							Index:           555,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD51a44d3FaE010294C616388b506AcdA1bfAAE46"),
							Topics: []common.Hash{
								common.HexToHash("0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee"),
								common.HexToHash("0x0000000000000000000000003993d34e7e99abf6b6f367309975d1360222d446"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008c640b185e89a0000000000000000000000000000000000000000000000000000021f76d530420000000000000000000000000000000000000000000003aa4efee6ea4bb03774"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19879519", 0)),
							TransactionHash: common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
							Index:           556,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xc4AD29ba4B3c580e6D59105FFf484999997675Ff"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000003993d34e7e99abf6b6f367309975d1360222d446"),
								common.HexToHash("0x0000000000000000000000007664f4a2d8c663d713170d0ab25b7c15d06b84a1"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000e75368d346ff5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19879519", 0)),
							TransactionHash: common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
							Index:           557,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42"),
						TransactionIndex: 156,
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
				ID:      "0xfe13c9ab435f72a3a7373036f2165f8a8c84b07f3f093ff5e35dcdc937f59c42",
				Network: network.Ethereum,
				Index:   156,
				From:    "0x7664f4A2d8c663d713170d0AB25B7C15d06B84a1",
				To:      "0x3993d34e7e99Abf6B6f367309975d1360222D446",
				Type:    typex.ExchangeLiquidity,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x4515cef3",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("977377442371290")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeLiquidity,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x7664f4A2d8c663d713170d0AB25B7C15d06B84a1",
						To:       "0xD51a44d3FaE010294C616388b506AcdA1bfAAE46",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityAdd,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2469780972234906"))),
									Name:     "Wrapped Ether",
									Symbol:   "WETH",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
					{
						Type:     typex.TransactionMint,
						Platform: workerx.PlatformCurve.String(),
						From:     ethereum.AddressGenesis.String(),
						To:       "0x7664f4A2d8c663d713170d0AB25B7C15d06B84a1",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0xc4AD29ba4B3c580e6D59105FFf484999997675Ff"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("4069526831656949"))),
							Name:     "Curve.fi USD-BTC-ETH",
							Symbol:   "crv3crypto",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
				},
				Status:    true,
				Timestamp: 1715826035,
			},
			wantError: require.NoError,
		},
		{
			name: "Remove CRV liquidity from StableSwap cvxCrv/Crv",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x253edfe9b30e5691aa3d83ae78a59ced30fde709e8b081a7057925a7033e10a7"),
						ParentHash:   common.HexToHash("0x1f88f5071bac567512baa7a3802736c34f3c246e4988f61094e6bee2cf85a0d8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("17919348", 0)),
						GasLimit:     30000000,
						GasUsed:      13291823,
						Timestamp:    1692091679,
						BaseFee:      lo.Must(new(big.Int).SetString("13994808854", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0"),
						From:      common.HexToAddress("0xb9A8f6204F08Df3609AfC4A0a352A1e62C8B8c5D"),
						Gas:       150150,
						GasPrice:  lo.Must(new(big.Int).SetString("14994808854", 10)),
						Hash:      common.HexToHash("0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0"),
						Input:     hexutil.MustDecode("0x1a4d01d200000000000000000000000000000000000000000000065b52b23e4869c9db45000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000063f7fd438df83c21f0f"),
						To:        lo.ToPtr(common.HexToAddress("0x971add32Ea87f10bD192671630be3BE8A11b8623")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x253edfe9b30e5691aa3d83ae78a59ced30fde709e8b081a7057925a7033e10a7"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17919348", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 5249497,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x37dc2a016"),
						GasUsed:           104311,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x971add32Ea87f10bD192671630be3BE8A11b8623"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000b9a8f6204f08df3609afc4a0a352a1e62c8b8c5d"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000065b52b23e4869c9db45"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919348", 0)),
							TransactionHash: common.HexToHash("0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0"),
							Index:           150,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD533a949740bb3306d119CC777fa900bA034cd52"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000971add32ea87f10bd192671630be3be8a11b8623"),
								common.HexToHash("0x000000000000000000000000b9a8f6204f08df3609afc4a0a352a1e62c8b8c5d"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000063ffab5187f178ddcf3"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919348", 0)),
							TransactionHash: common.HexToHash("0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0"),
							Index:           151,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x971add32Ea87f10bD192671630be3BE8A11b8623"),
							Topics: []common.Hash{
								common.HexToHash("0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0"),
								common.HexToHash("0x000000000000000000000000b9a8f6204f08df3609afc4a0a352a1e62c8b8c5d"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000065b52b23e4869c9db4500000000000000000000000000000000000000000000063ffab5187f178ddcf30000000000000000000000000000000000000000003997b12fa9ba7b788fed1b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17919348", 0)),
							TransactionHash: common.HexToHash("0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0"),
							Index:           152,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0"),
						TransactionIndex: 72,
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
				ID:      "0x3dc1676ca76d0d9afd9dded53a0823d32f43454be439df1e56d34891eccb86c0",
				Network: network.Ethereum,
				Index:   72,
				From:    "0xb9A8f6204F08Df3609AfC4A0a352A1e62C8B8c5D",
				To:      "0x971add32Ea87f10bD192671630be3BE8A11b8623",
				Type:    typex.ExchangeLiquidity,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x1a4d01d2",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("1564123506369594")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBurn,
						Platform: workerx.PlatformCurve.String(),
						From:     "0xb9A8f6204F08Df3609AfC4A0a352A1e62C8B8c5D",
						To:       ethereum.AddressGenesis.String(),
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x971add32Ea87f10bD192671630be3BE8A11b8623"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("30018811501663138470725"))),
							Name:     "Curve.fi Factory Plain Pool: cvxCrv/Crv",
							Symbol:   "cvxcrv-crv-f",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     typex.ExchangeLiquidity,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x971add32Ea87f10bD192671630be3BE8A11b8623",
						To:       "0xb9A8f6204F08Df3609AfC4A0a352A1e62C8B8c5D",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRemove,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xD533a949740bb3306d119CC777fa900bA034cd52"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("29514409146275974733043"))),
									Name:     "Curve DAO Token",
									Symbol:   "CRV",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692091679,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit 3Crv to DAI/USDC/USDT Gauge",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xa127faa2fcd741ea97dd0b4f7282872b15f4c92cb3d5bff493188f06e209b1c7"),
						ParentHash:   common.HexToHash("0x9b9ec13949a81836a01f70edd7f0d95b4a57c71141c2af88a415d3135861e3c7"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("17918045", 0)),
						GasLimit:     30000000,
						GasUsed:      29984188,
						Timestamp:    1692075935,
						BaseFee:      lo.Must(new(big.Int).SetString("10828959989", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda"),
						From:      common.HexToAddress("0x2AF2b2e485E1854fd71590C7Ffd104dB0f66F8A6"),
						Gas:       919848,
						GasPrice:  lo.Must(new(big.Int).SetString("11000000000", 10)),
						Hash:      common.HexToHash("0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda"),
						Input:     hexutil.MustDecode("0xb6b55f2500000000000000000000000000000000000000000000000f9748bb0f90ed0682"),
						To:        lo.ToPtr(common.HexToAddress("0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xa127faa2fcd741ea97dd0b4f7282872b15f4c92cb3d5bff493188f06e209b1c7"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17918045", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 9277913,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x28fa6ae00"),
						GasUsed:           433023,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A"),
							Topics: []common.Hash{
								common.HexToHash("0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000002af2b2e485e1854fd71590c7ffd104db0f66f8a600000000000000000000000000000000000000000000000f9748bb0f90ed0682000000000000000000000000000000000000000000157bac038ca6ef34a7e3e700000000000000000000000000000000000000000000000dcea9fffe4d7a39900000000000000000000000000000000000000000000d2f781df34370f6e63205"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17918045", 0)),
							TransactionHash: common.HexToHash("0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda"),
							Index:           182,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000002af2b2e485e1854fd71590c7ffd104db0f66f8a6"),
								common.HexToHash("0x000000000000000000000000bfcf63294ad7105dea65aa58f8ae5be2d9d0952a"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000f9748bb0f90ed0682"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17918045", 0)),
							TransactionHash: common.HexToHash("0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda"),
							Index:           183,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x0000000000000000000000002af2b2e485e1854fd71590c7ffd104db0f66f8a6"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000f9748bb0f90ed0682"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17918045", 0)),
							TransactionHash: common.HexToHash("0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda"),
							Index:           184,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda"),
						TransactionIndex: 87,
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
				ID:      "0xefd90c9bc0560e3eaac3667cc6b671fb73076ed51b1de3d9e17d2128beddcdda",
				Network: network.Ethereum,
				Index:   87,
				From:    "0x2AF2b2e485E1854fd71590C7Ffd104dB0f66F8A6",
				To:      "0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A",
				Type:    typex.ExchangeStaking,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xb6b55f25",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("4763253000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeStaking,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x2AF2b2e485E1854fd71590C7Ffd104dB0f66F8A6",
						To:       "0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingStake,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("287602329679223916162"))),
								Name:     "Curve.fi DAI/USDC/USDT",
								Symbol:   "3Crv",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692075935,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw 3Crv from DAI/USDC/USDT Gauge",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x74329bad4a85c071ada4c4de7ff8e4b0bf2cde350fd9f3f9d4786ca33dea4832"),
						ParentHash:   common.HexToHash("0x6826f5dd71c33eff1e7221e59ef774446b9b2536a2c4273b486f9374aadb8e99"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5"),
						Number:       lo.Must(new(big.Int).SetString("17918050", 0)),
						GasLimit:     30000000,
						GasUsed:      15233071,
						Timestamp:    1692075995,
						BaseFee:      lo.Must(new(big.Int).SetString("10918186082", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f"),
						From:      common.HexToAddress("0x66A2709d2c682Db7f623E1B6B14C918196cC40C8"),
						Gas:       820642,
						GasPrice:  lo.Must(new(big.Int).SetString("10943933754", 10)),
						Hash:      common.HexToHash("0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f"),
						Input:     hexutil.MustDecode("0x2e1a7d4d00000000000000000000000000000000000000000000001ff4a2bedf5c5c93a5"),
						To:        lo.ToPtr(common.HexToAddress("0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x74329bad4a85c071ada4c4de7ff8e4b0bf2cde350fd9f3f9d4786ca33dea4832"),
						BlockNumber:       lo.Must(new(big.Int).SetString("17918050", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 14691938,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x28c4f2d3a"),
						GasUsed:           378620,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A"),
							Topics: []common.Hash{
								common.HexToHash("0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000066a2709d2c682db7f623e1b6b14c918196cc40c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000157b8c0ee9e80fd84b504200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d2f6b55b22a4ad1f45d5d"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17918050", 0)),
							TransactionHash: common.HexToHash("0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f"),
							Index:           359,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000bfcf63294ad7105dea65aa58f8ae5be2d9d0952a"),
								common.HexToHash("0x00000000000000000000000066a2709d2c682db7f623e1b6b14c918196cc40c8"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000001ff4a2bedf5c5c93a5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17918050", 0)),
							TransactionHash: common.HexToHash("0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f"),
							Index:           360,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A"),
							Topics: []common.Hash{
								common.HexToHash("0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364"),
								common.HexToHash("0x00000000000000000000000066a2709d2c682db7f623e1b6b14c918196cc40c8"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000001ff4a2bedf5c5c93a5"),
							BlockNumber:     lo.Must(new(big.Int).SetString("17918050", 0)),
							TransactionHash: common.HexToHash("0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f"),
							Index:           361,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f"),
						TransactionIndex: 191,
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
				ID:      "0x6c5460d21f5b9ae59af3f954146ff21f1acaff9030ceac5537738259438e9e3f",
				Network: network.Ethereum,
				Index:   191,
				From:    "0x66A2709d2c682Db7f623E1B6B14C918196cC40C8",
				To:      "0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A",
				Type:    typex.ExchangeStaking,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x2e1a7d4d",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("4143592197939480")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeStaking,
						Platform: workerx.PlatformCurve.String(),
						From:     "0xbFcF63294aD7105dEa65aA58F8AE5BE2D9d0952A",
						To:       "0x66A2709d2c682Db7f623E1B6B14C918196cC40C8",
						Metadata: metadata.ExchangeStaking{
							Action: metadata.ActionExchangeStakingUnstake,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("589476928043014198181"))),
								Name:     "Curve.fi DAI/USDC/USDT",
								Symbol:   "3Crv",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692075995,
			},
			wantError: require.NoError,
		},
		{
			name: "Remove CRV liquidity from StableSwap cbETH/ETH",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xf8ccea9d3ccbb2f975ada4134a476ccfda38c381ca0838cd897121ca450b4aa5"),
						ParentHash:   common.HexToHash("0x4ea56270ab03b52d6cd539059fa823bfe2208435b958e41a0e189616140f9cce"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326"),
						Number:       lo.Must(new(big.Int).SetString("20117906", 0)),
						GasLimit:     30000000,
						GasUsed:      11473198,
						Timestamp:    1718704727,
						BaseFee:      lo.Must(new(big.Int).SetString("4928434420", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
						From:      common.HexToAddress("0x75347799e8824d9aD096Fa6bB5bb1Ab56d15E89e"),
						Gas:       408000,
						GasPrice:  lo.Must(new(big.Int).SetString("5015535854", 10)),
						Hash:      common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
						Input:     hexutil.MustDecode("0x5b36389c0000000000000000000000000000000000000000000000000079c4071ad2ca9f000000000000000000000000000000000000000000000000007277a1710106e700000000000000000000000000000000000000000000000000825fa0852eee70"),
						To:        lo.ToPtr(common.HexToAddress("0x5FAE7E604FC3e24fd43A72867ceBaC94c65b404A")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xf8ccea9d3ccbb2f975ada4134a476ccfda38c381ca0838cd897121ca450b4aa5"),
						BlockNumber:       lo.Must(new(big.Int).SetString("20117906", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8834540,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x12af300ee"),
						GasUsed:           131513,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x5b6C539b224014A09B3388e51CaAA8e354c959C8"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000075347799e8824d9ad096fa6bb5bb1ab56d15e89e"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000079c4071ad2ca9f"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20117906", 0)),
							TransactionHash: common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
							Index:           200,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x0000000000000000000000005fae7e604fc3e24fd43a72867cebac94c65b404a"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000730ae29304e433"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20117906", 0)),
							TransactionHash: common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
							Index:           201,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005fae7e604fc3e24fd43a72867cebac94c65b404a"),
								common.HexToHash("0x00000000000000000000000075347799e8824d9ad096fa6bb5bb1ab56d15e89e"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000730ae29304e433"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20117906", 0)),
							TransactionHash: common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
							Index:           202,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xBe9895146f7AF43049ca1c1AE358B0541Ea49704"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000005fae7e604fc3e24fd43a72867cebac94c65b404a"),
								common.HexToHash("0x00000000000000000000000075347799e8824d9ad096fa6bb5bb1ab56d15e89e"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000083075800104af8"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20117906", 0)),
							TransactionHash: common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
							Index:           203,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x5FAE7E604FC3e24fd43A72867ceBaC94c65b404A"),
							Topics: []common.Hash{
								common.HexToHash("0xdd3c0336a16f1b64f172b7bb0dad5b2b3c7c76f91e8c4aafd6aae60dce800153"),
								common.HexToHash("0x00000000000000000000000075347799e8824d9ad096fa6bb5bb1ab56d15e89e"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000730ae29304e4330000000000000000000000000000000000000000000000000083075800104af8000000000000000000000000000000000000000000000014035989e086077480"),
							BlockNumber:     lo.Must(new(big.Int).SetString("20117906", 0)),
							TransactionHash: common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
							Index:           204,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5"),
						TransactionIndex: 104,
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
				ID:      "0x2775104b193d825c59376b3eebae10831c76b7204bdfc8c2a9b1b069fc3eaac5",
				Network: network.Ethereum,
				Index:   104,
				From:    "0x75347799e8824d9aD096Fa6bB5bb1Ab56d15E89e",
				To:      "0x5FAE7E604FC3e24fd43A72867ceBaC94c65b404A",
				Type:    typex.ExchangeLiquidity,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x5b36389c",
				},
				Platform: workerx.PlatformCurve.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("659608166767102")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBurn,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x75347799e8824d9aD096Fa6bB5bb1Ab56d15E89e",
						To:       ethereum.AddressGenesis.String(),
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x5b6C539b224014A09B3388e51CaAA8e354c959C8"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("34274006975826591"))),
							Name:     "Curve.fi Factory Crypto Pool: cbETH/ETH",
							Symbol:   "cbETH/ETH-f",
							Decimals: 18,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     typex.ExchangeLiquidity,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x5FAE7E604FC3e24fd43A72867ceBaC94c65b404A",
						To:       "0x75347799e8824d9aD096Fa6bB5bb1Ab56d15E89e",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRemove,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("32381590567183411"))),
									Name:     "Wrapped Ether",
									Symbol:   "WETH",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
					{
						Type:     typex.ExchangeLiquidity,
						Platform: workerx.PlatformCurve.String(),
						From:     "0x5FAE7E604FC3e24fd43A72867ceBaC94c65b404A",
						To:       "0x75347799e8824d9aD096Fa6bB5bb1Ab56d15E89e",
						Metadata: metadata.ExchangeLiquidity{
							Action: metadata.ActionExchangeLiquidityRemove,
							Tokens: []metadata.Token{
								{
									Address:  lo.ToPtr("0xBe9895146f7AF43049ca1c1AE358B0541Ea49704"),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("36881296488680184"))),
									Name:     "Coinbase Wrapped Staked ETH",
									Symbol:   "cbETH",
									Decimals: 18,
									Standard: metadata.StandardERC20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1718704727,
			},
			wantError: require.NoError,
		},
	}

	// Start Redis container
	preset := redis.Preset(
		redis.WithVersion("6.0.9"),
	)

	container, err := gnomock.Start(preset)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, gnomock.Stop(container))
	})

	// Connect to Redis
	redisClient, err := redisx.NewClient(config.Redis{
		Endpoint: container.DefaultAddress(),
	})
	require.NoError(t, err)

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker(testcase.arguments.config, redisClient)
			require.NoError(t, err)

			activity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			// t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
