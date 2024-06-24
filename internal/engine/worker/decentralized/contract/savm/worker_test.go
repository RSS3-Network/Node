package savm_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/savm"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/savm"
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
			name: "Withdraw WBTC",
			arguments: arguments{
				task: &source.Task{
					Network: network.SatoshiVM,
					ChainID: 3109,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x06ade7dc3b83d30455a99a0a28da08e7cac601ffdeb207c5b8270d7327109f60"),
						ParentHash:   common.HexToHash("0x3c9b96ff9eab5d06c3cd5492b05cd0f49aeb2d651007c0051fe649dc8471ddb1"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x10Eed66cD72390e6b2CCf3150aBbB9836337a544"),
						Number:       lo.Must(new(big.Int).SetString("65961", 0)),
						GasLimit:     50000000,
						GasUsed:      177859,
						Timestamp:    1710430332,
						BaseFee:      lo.Must(new(big.Int).SetString("75000000", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
						From:      common.HexToAddress("0x9e5635611981425B7AcbF827516123143F1d2237"),
						Gas:       261270,
						GasPrice:  lo.Must(new(big.Int).SetString("75000000", 10)),
						Hash:      common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
						Input:     hexutil.MustDecode("0x030ba25d00000000000000000000000000000000000000000000000000002f4b3187400000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000040010062633170646730636d6b713664743430686333746165363335676c3767667a657a30756a6d6e7575643070767a3374643673396134663971377434777836"),
						To:        lo.ToPtr(common.HexToAddress("0xF70Af817B07118CBF7ACCC38767899598e045408")),
						Value:     lo.Must(new(big.Int).SetString("200000000000000", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("3109", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x06ade7dc3b83d30455a99a0a28da08e7cac601ffdeb207c5b8270d7327109f60"),
						BlockNumber:       lo.Must(new(big.Int).SetString("65961", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 177859,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x47868c0"),
						GasUsed:           177859,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xca4098FcEcC3Bf7ad04691E1186f714A72E70E6d"),
							Topics: []common.Hash{
								common.HexToHash("0x274ed44bc09f9ff15c7aab597c835942a05dbf38024e688e775b3dd5c30772d9"),
							},
							Data:            hexutil.MustDecode("0x00000c25a089881600000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000246139ca8000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("65961", 0)),
							TransactionHash: common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x10E82B34E757Cf4fdEFC0E9Fc8CCA4e5b9749bcE"),
							Topics: []common.Hash{
								common.HexToHash("0x838807fe70997cb429689dd35bc6ccb21720aa7d7daac8736b398060104fb0fa"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000c25a089881600000000000000000000000000000000000000000000000be549a52afcf896c43357b20825d7faaf1576cdecf7c3a307ab7339e70483305d0000000000000000000000002c304412842606fe810515e7b36c1544d769b57a00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000ece9983eebb068f20a4e5547796f7539bdd4f4ea00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000001450000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000040010062633170646730636d6b713664743430686333746165363335676c3767667a657a30756a6d6e7575643070767a33746436733961346639713774347778360000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("65961", 0)),
							TransactionHash: common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4B98599D7EC7D8484A0b6b561F65293a526bf4Dd"),
							Topics: []common.Hash{
								common.HexToHash("0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364"),
								common.HexToHash("0x000000000000000000000000f70af817b07118cbf7accc38767899598e045408"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000002f4b31874000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("65961", 0)),
							TransactionHash: common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xF70Af817B07118CBF7ACCC38767899598e045408"),
							Topics: []common.Hash{
								common.HexToHash("0x181177e6e9bbf7d7a45604cf4fdb0444641e20905d8d4daf2a7863036e56dc15"),
							},
							Data:            hexutil.MustDecode("0x00000c25a089881600000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000000004e200000000000000000000000009e5635611981425b7acbf827516123143f1d2237"),
							BlockNumber:     lo.Must(new(big.Int).SetString("65961", 0)),
							TransactionHash: common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
							Index:           3,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.SatoshiVM,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.SatoshiVM),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xf9d9724861f5e5558cbf70b9347e561462252990d2f4ed284a20514c78c77363",
				Network: network.SatoshiVM,
				Index:   0,
				From:    "0x9e5635611981425B7AcbF827516123143F1d2237",
				To:      savm.AddressBTCBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x030ba25d",
				},
				Platform: workerx.PlatformSAVM.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("13339425000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformSAVM.String(),
						From:     "0x9e5635611981425B7AcbF827516123143F1d2237",
						To:       "0x9e5635611981425B7AcbF827516123143F1d2237",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.SatoshiVM,
							TargetNetwork: network.Bitcoin,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("20000"))),
								Name:     "Bitcoin",
								Symbol:   "BTC",
								Decimals: 8,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1710430332,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdraw SAVM",
			arguments: arguments{
				task: &source.Task{
					Network: network.SatoshiVM,
					ChainID: 3109,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x147c37f2d1e5d743fb09aab9ffefc67b6e884a3e1627db133d52bc3636829921"),
						ParentHash:   common.HexToHash("0xe711a216afa2e72c1c5c0d497560ec5e6fc0d6fe97dea6128709fc5134c6c171"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x10Eed66cD72390e6b2CCf3150aBbB9836337a544"),
						Number:       lo.Must(new(big.Int).SetString("64773", 0)),
						GasLimit:     50000000,
						GasUsed:      145040,
						Timestamp:    1710426768,
						BaseFee:      lo.Must(new(big.Int).SetString("75000000", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
						From:      common.HexToAddress("0x9e5635611981425B7AcbF827516123143F1d2237"),
						Gas:       163155,
						GasPrice:  lo.Must(new(big.Int).SetString("75000000", 10)),
						Hash:      common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
						Input:     hexutil.MustDecode("0xe43772fb0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000009e5635611981425b7acbf827516123143f1d223700000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x1109F6221F684BCb9B2529b8803a7b8c3411d45f")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("3109", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x147c37f2d1e5d743fb09aab9ffefc67b6e884a3e1627db133d52bc3636829921"),
						BlockNumber:       lo.Must(new(big.Int).SetString("64773", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 145040,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x47868c0"),
						GasUsed:           145040,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x0E02765992f946397E6d2e65642eABb9cc674928"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000009e5635611981425b7acbf827516123143f1d2237"),
								common.HexToHash("0x0000000000000000000000001109f6221f684bcb9b2529b8803a7b8c3411d45f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000016345785d8a0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("64773", 0)),
							TransactionHash: common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0E02765992f946397E6d2e65642eABb9cc674928"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000001109f6221f684bcb9b2529b8803a7b8c3411d45f"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000016345785d8a0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("64773", 0)),
							TransactionHash: common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xca4098FcEcC3Bf7ad04691E1186f714A72E70E6d"),
							Topics: []common.Hash{
								common.HexToHash("0x274ed44bc09f9ff15c7aab597c835942a05dbf38024e688e775b3dd5c30772d9"),
							},
							Data:            hexutil.MustDecode("0x00000c25000000010000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("64773", 0)),
							TransactionHash: common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x10E82B34E757Cf4fdEFC0E9Fc8CCA4e5b9749bcE"),
							Topics: []common.Hash{
								common.HexToHash("0x838807fe70997cb429689dd35bc6ccb21720aa7d7daac8736b398060104fb0fa"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000c2500000001000000000000000000000000000000000000000000000004966c63d14939ec9ace2dc744f5ea970e1cc6f20f12afefdcdff58ed5d321637e00000000000000000000000017f5b02b99ff5f74add7babe57dc08295b4f3de500000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000104959e3ce9135e60e0300daafaea6b262f0e51a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000009e5635611981425b7acbf827516123143f1d223700000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("64773", 0)),
							TransactionHash: common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
							Index:           3,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x1109F6221F684BCb9B2529b8803a7b8c3411d45f"),
							Topics: []common.Hash{
								common.HexToHash("0x799e69b309353cce2426b9f6a5c9e5aaae30669a28e9aa4b90c21503fe82bbf4"),
								common.HexToHash("0x0000000000000000000000009e5635611981425b7acbf827516123143f1d2237"),
							},
							Data:            hexutil.MustDecode("0x00000c2500000001000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000016345785d8a0000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("64773", 0)),
							TransactionHash: common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
							Index:           4,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.SatoshiVM,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.SatoshiVM),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x21eafa992790b01bbd5b0445d1ba4fd0dfcc6ba2c7f46821588feabe42012ffa",
				Network: network.SatoshiVM,
				Index:   0,
				From:    "0x9e5635611981425B7AcbF827516123143F1d2237",
				To:      savm.AddressSAVMBridge.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xe43772fb",
				},
				Platform: workerx.PlatformSAVM.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("10878000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformSAVM.String(),
						From:     "0x9e5635611981425B7AcbF827516123143F1d2237",
						To:       "0x9e5635611981425B7AcbF827516123143F1d2237",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.SatoshiVM,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr(savm.AddressSAVMToken.String()),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("100000000000000000"))),
								Name:     "SatoshiVM",
								Symbol:   "SAVM",
								Decimals: 18,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1710426768,
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

			// t.Log(string(lo.Must(json.MarshalIndent(activity, "", "\x20\x20"))))

			require.Equal(t, testcase.want, activity)
		})
	}
}
