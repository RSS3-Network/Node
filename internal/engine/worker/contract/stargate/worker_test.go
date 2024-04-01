package stargate_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/stargate"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/layerzero"
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
			name: "Withdrawal ETH from Ethereum Optimism to Ethereum Avalanche on Ethereum Optimism",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkOptimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x86968212025a81726ea274c5ce80e7c4ced9d76daa62802888c59df21861e27d"),
						ParentHash:   common.HexToHash("0xc758d41ae8ad1d430f1f5756636d271f982d768cdd85f806f6bf1aebe7e20f34"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("110308470", 0)),
						GasLimit:     30000000,
						GasUsed:      23605998,
						Timestamp:    1696215717,
						BaseFee:      lo.Must(new(big.Int).SetString("26291662", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
						From:      common.HexToAddress("0xe93685f3bBA03016F02bD1828BaDD6195988D950"),
						Gas:       1182874,
						GasPrice:  lo.Must(new(big.Int).SetString("150084666", 10)),
						Hash:      common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
						Input:     hexutil.MustDecode("0x252f7b01000000000000000000000000000000000000000000000000000000000000006e000000000000000000000000701a95707a0290ac8b90b3719e8ee5b21036088300000000000000000000000000000000000000000000000000000000000432387117dbb455d8f038413e7099ccfc2d1639a846166e95ff697c63449d661b2c247117dbb455d8f038413e7099ccfc2d1639a846166e95ff697c63449d661b2c2400000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000002740000000000000000000000004d73adb72bc3dd368966edd0f0b2148401a178e200000000002e2ceb006e352d8275aae3e0c2404d9f68f6cee084b5beb3dd006f701a95707a0290ac8b90b3719e8ee5b2103608830000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008465736214537f3e170000000000000000000000000000000000000000000000000427dea0f167bb99000000000000000000000000000000000000000000000000000044d76b635e67000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a38cc07a76000000000000000000000000000000000000000000000000000428c7051d45900000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001466d4f4a4a743a6136d32d4c93ea03e839a62dad70000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x81E792e5a9003CC1C8BF5569A00f34b65d75b017")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x86968212025a81726ea274c5ce80e7c4ced9d76daa62802888c59df21861e27d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("110308470", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 252392,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x8f21c3a"),
						GasUsed:           205479,
						L1GasPrice:        lo.Must(new(big.Int).SetString("6271260177", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("8028", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("34436442863453", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
							Topics: []common.Hash{
								common.HexToHash("0x2bd2d8a84b748439fd50d79a49502b4eb5faa25b864da6a9ab5c150704be9a4d"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000006e"),
								common.HexToHash("0x000000000000000000000000701a95707a0290ac8b90b3719e8ee5b210360883"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000002e2ceba79700fff941a72bd3fcc82dd239d1b3283c508081ed3e125e70c548cd94545b0000000000000000000000000000000000000000000000000000000000000014352d8275aae3e0c2404d9f68f6cee084b5beb3dd000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308470", 0)),
							TransactionHash: common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd22363e3762cA7339569F3d33EADe20127D5F98C"),
							Topics: []common.Hash{
								common.HexToHash("0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000006e000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008465736214537f3e17"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308470", 0)),
							TransactionHash: common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xb69c8CBCD90A39D8D3d3ccf0a3E968511C3856A0"),
							Topics: []common.Hash{
								common.HexToHash("0xb4a87134099d10c48345145381989042ab07dc53e6e62a6511fca55438562e26"),
								common.HexToHash("0x000000000000000000000000d22363e3762ca7339569f3d33eade20127d5f98c"),
								common.HexToHash("0x00000000000000000000000066d4f4a4a743a6136d32d4c93ea03e839a62dad7"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000427dea0f167bb99"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308470", 0)),
							TransactionHash: common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xd22363e3762cA7339569F3d33EADe20127D5F98C"),
							Topics: []common.Hash{
								common.HexToHash("0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000066d4f4a4a743a6136d32d4c93ea03e839a62dad70000000000000000000000000000000000000000000000000427dea0f167bb990000000000000000000000000000000000000000000000000000a38cc07a7600000000000000000000000000000000000000000000000000000044d76b635e67"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308470", 0)),
							TransactionHash: common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
							Index:           3,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c"),
						TransactionIndex: 1,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkOptimism,
					Endpoint: endpoint.MustGet(filter.NetworkOptimism),
				},
			},
			want: &schema.Feed{
				ID:      "0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c",
				Network: filter.NetworkOptimism,
				Index:   1,
				From:    "0xe93685f3bBA03016F02bD1828BaDD6195988D950",
				To:      layerzero.AddressRelayerOptimism.String(),
				Type:    filter.TypeTransactionBridge,
				Calldata: &schema.Calldata{
					FunctionHash: "0x252f7b01",
				},
				Platform: lo.ToPtr(filter.PlatformStargate),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("65275689948467")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeTransactionBridge,
						Platform: filter.PlatformStargate.String(),
						From:     "0xd22363e3762cA7339569F3d33EADe20127D5F98C",
						To:       "0x66d4f4A4A743A6136D32d4c93eA03e839A62Dad7",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: filter.NetworkArbitrum,
							TargetNetwork: filter.NetworkOptimism,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("299452683069668249"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696215717,
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
