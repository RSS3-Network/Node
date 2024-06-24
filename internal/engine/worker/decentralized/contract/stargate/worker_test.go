package stargate_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/stargate"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/layerzero"
	"github.com/rss3-network/node/provider/ethereum/contract/stargate"
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
			name: "Deposit ETH from Ethereum Arbitrum One to Ethereum Base on Ethereum Arbitrum",
			arguments: arguments{
				task: &source.Task{
					Network: network.Arbitrum,
					ChainID: 42161,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x5958fe64312e0751d5036731682b1ae518f7236efa7608cb73bae6e48cfc88cf"),
						ParentHash:   common.HexToHash("0xead23c4148b5d80e02149a5f1e1cd69e16293a0e287265d04b797fd3b475154e"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xA4b000000000000000000073657175656e636572"),
						Number:       lo.Must(new(big.Int).SetString("136635694", 0)),
						GasLimit:     1125899906842624,
						GasUsed:      1332027,
						Timestamp:    1696214752,
						BaseFee:      lo.Must(new(big.Int).SetString("100000000", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
						From:      common.HexToAddress("0x32C8e4D2Cb2642A29bC06115336AC96Bf0160485"),
						Gas:       833042,
						GasPrice:  lo.Must(new(big.Int).SetString("100000000", 10)),
						Hash:      common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
						Input:     hexutil.MustDecode("0x1114cd2a00000000000000000000000000000000000000000000000000000000000000b800000000000000000000000032c8e4d2cb2642a29bc06115336ac96bf016048500000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000d18df9ff2c660000000000000000000000000000000000000000000000000000d081bf2593db4000000000000000000000000000000000000000000000000000000000000000001432c8e4d2cb2642a29bc06115336ac96bf0160485000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xbf22f0f184bCcbeA268dF387a49fF5238dD23E40")),
						Value:     lo.Must(new(big.Int).SetString("15100154593365273712", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("42161", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x5958fe64312e0751d5036731682b1ae518f7236efa7608cb73bae6e48cfc88cf"),
						BlockNumber:       lo.Must(new(big.Int).SetString("136635694", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 1116552,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x5f5e100"),
						GasUsed:           702946,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x82CbeCF39bEe528B5476FE6d1550af59a9dB6Fc0"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x000000000000000000000000bf22f0f184bccbea268df387a49ff5238dd23e40"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000d18df9ff2c660000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x82CbeCF39bEe528B5476FE6d1550af59a9dB6Fc0"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000bf22f0f184bccbea268df387a49ff5238dd23e40"),
								common.HexToHash("0x00000000000000000000000053bf833a5d6c4dda888f69c22c88c9f356a41614"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000d18df9ff2c660000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x915A55e36A01285A14f05dE6e81ED9cE89772f8e"),
							Topics: []common.Hash{
								common.HexToHash("0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000b8000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000bf22f0f184bccbea268df387a49ff5238dd23e40000000000000000000000000000000000000000000000000d16dc9f98abbc000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000019c0048155000000000000000000000000000000000000000000000000000001e94055994f0000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           3,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x82CbeCF39bEe528B5476FE6d1550af59a9dB6Fc0"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000bf22f0f184bccbea268df387a49ff5238dd23e40"),
								common.HexToHash("0x000000000000000000000000915a55e36a01285a14f05de6e81ed9ce89772f8e"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000d18df9ff2c660000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           4,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x915A55e36A01285A14f05dE6e81ED9cE89772f8e"),
							Topics: []common.Hash{
								common.HexToHash("0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000b8000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004e31a39abb5373e8fc"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           5,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x177d36dBE2271A4DdB2Ad8304d82628eb921d790"),
							Topics: []common.Hash{
								common.HexToHash("0xdf21c415b78ed2552cc9971249e32a053abce6087a0ae0fbf3f78db5174a3493"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000004f8ddc18c688"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           6,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
							Topics: []common.Hash{
								common.HexToHash("0xb0c632f55f1e1b3b2c3d82f41ee4716bb4c00f0f5d84cdafc141581bb8757a4f"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000220001000000000000000000000000000000000000000000000000000000000002ab98000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           7,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD56e4eAb23cb81f43168F9F45211Eb027b9aC7cc"),
							Topics: []common.Hash{
								common.HexToHash("0x87e46b0a6199bc734632187269a103c05714ee0adae5b28f30723955724f37ef"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000003d0c368665e8"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           8,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
							Topics: []common.Hash{
								common.HexToHash("0xe9bded5f24a4168e4f3bf44e00298c993b22376aad8c58c7dda9718a54cbea82"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000002540000000000047ce7006e352d8275aae3e0c2404d9f68f6cee084b5beb3dd00b8af54be5b6eec24d6bfacf1cce4eaf680a82393980000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004e31a39abb5373e8fc000000000000000000000000000000000000000000000000d16dc9f98abbc00000000000000000000000000000000000000000000000000000019c004815500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001e94055994f000000000000000000000000000000000000000000000000000d18df9ff2c66000000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001432c8e4d2cb2642a29bc06115336ac96bf01604850000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           9,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x352d8275AAE3e0c2404d9f68f6cEE084B5bEB3DD"),
							Topics: []common.Hash{
								common.HexToHash("0x8d3ee0df6a4b7e82a7f20a763f1c6826e6176323e655af64f32318827d2112d4"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000047ce7"),
							BlockNumber:     lo.Must(new(big.Int).SetString("136635694", 0)),
							TransactionHash: common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
							Index:           10,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d"),
						TransactionIndex: 2,
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
				ID:      "0x9c14874078e6c92eefe807d1b8595ba05295a6a1d5abaa2634ca286e8409b61d",
				Network: network.Arbitrum,
				Index:   2,
				From:    "0x32C8e4D2Cb2642A29bC06115336AC96Bf0160485",
				To:      stargate.AddressRouterETHArbitrum.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x1114cd2a",
				},
				Platform: workerx.PlatformStargate.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("70294600000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
						From: "0x32C8e4D2Cb2642A29bC06115336AC96Bf0160485",
						To:   layerzero.AddressUltraLightNodeArbitrumOne.String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("154593365273712"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type: typex.TransactionTransfer,
						From: "0x32C8e4D2Cb2642A29bC06115336AC96Bf0160485",
						To:   "0x915A55e36A01285A14f05dE6e81ED9cE89772f8e",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("9060000000000000"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformStargate.String(),
						From:     "0x32C8e4D2Cb2642A29bC06115336AC96Bf0160485",
						To:       "0x32C8e4D2Cb2642A29bC06115336AC96Bf0160485",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Arbitrum,
							TargetNetwork: network.Base,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("15090940000000000000"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696214752,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdrawal ETH from Ethereum Optimism to Ethereum Avalanche on Ethereum Optimism",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
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
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x1ee6eca8ad26a8240e2b83352cb5e642636359aa0bd82a88ce72aba7a988305c",
				Network: network.Optimism,
				Index:   1,
				From:    "0xe93685f3bBA03016F02bD1828BaDD6195988D950",
				To:      layerzero.AddressRelayerOptimism.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x252f7b01",
				},
				Platform: workerx.PlatformStargate.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("65275689948467")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformStargate.String(),
						From:     "0xd22363e3762cA7339569F3d33EADe20127D5F98C",
						To:       "0x66d4f4A4A743A6136D32d4c93eA03e839A62Dad7",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Arbitrum,
							TargetNetwork: network.Optimism,
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
		{
			name: "Deposit USDC from Ethereum Optimism to Ethereum Avalanche on Ethereum Optimism",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x166fff874bf97a5f1fb8add7b06eb78affcb528e9772965101eb2aa36ef3cad3"),
						ParentHash:   common.HexToHash("0x13d38b5e0e4c758ba3bf05131fdd3e40f429b96bfceb996eb52973fadb239bc2"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("110308341", 0)),
						GasLimit:     30000000,
						GasUsed:      2567492,
						Timestamp:    1696215459,
						BaseFee:      lo.Must(new(big.Int).SetString("25004152", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
						From:      common.HexToAddress("0xA9fdDD748b21898C7E26E05ad895760FDd2723Bd"),
						Gas:       1489020,
						GasPrice:  lo.Must(new(big.Int).SetString("26671573", 10)),
						Hash:      common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
						Input:     hexutil.MustDecode("0x9fbf10fc000000000000000000000000000000000000000000000000000000000000006a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a9fddd748b21898c7e26e05ad895760fdd2723bd00000000000000000000000000000000000000000000000000000000a8b9e84900000000000000000000000000000000000000000000000000000000a709f7e1000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000026c975e37b7e29a00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000014a9fddd748b21898c7e26e05ad895760fdd2723bd0000000000000000000000000000000000000000000000000000000000000000000000000000000000000014a9fddd748b21898c7e26e05ad895760fdd2723bd0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xB0D502E938ed5f4df2E681fE6E419ff29631d62b")),
						Value:     lo.Must(new(big.Int).SetString("1482831510889074", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x166fff874bf97a5f1fb8add7b06eb78affcb528e9772965101eb2aa36ef3cad3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("110308341", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2315849,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x196f9d5"),
						GasUsed:           371488,
						L1GasPrice:        lo.Must(new(big.Int).SetString("6414149718", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("5264", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("23094633535037", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xDecC0c09c3B5f6e92EF4184125D5648a66E35298"),
							Topics: []common.Hash{
								common.HexToHash("0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000006a0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a9fddd748b21898c7e26e05ad895760fdd2723bd00000000000000000000000000000000000000000000000000000000a8926f31000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d8e84000000000000000000000000000000000000000000000000000000000019ea940000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           33,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000a9fddd748b21898c7e26e05ad895760fdd2723bd"),
								common.HexToHash("0x000000000000000000000000decc0c09c3b5f6e92ef4184125d5648a66e35298"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000a8b9e849"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           34,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000a9fddd748b21898c7e26e05ad895760fdd2723bd"),
								common.HexToHash("0x000000000000000000000000b0d502e938ed5f4df2e681fe6e419ff29631d62b"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000016ce6d68c81790465a36"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           35,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xDecC0c09c3B5f6e92EF4184125D5648a66E35298"),
							Topics: []common.Hash{
								common.HexToHash("0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000006a000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013a49a66372"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           36,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x81E792e5a9003CC1C8BF5569A00f34b65d75b017"),
							Topics: []common.Hash{
								common.HexToHash("0xdf21c415b78ed2552cc9971249e32a053abce6087a0ae0fbf3f78db5174a3493"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000050146b351bbfd"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           37,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
							Topics: []common.Hash{
								common.HexToHash("0xb0c632f55f1e1b3b2c3d82f41ee4716bb4c00f0f5d84cdafc141581bb8757a4f"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000560002000000000000000000000000000000000000000000000000000000000002ab98000000000000000000000000000000000000000000000000026c975e37b7e29aa9fddd748b21898c7e26e05ad895760fdd2723bd00000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           38,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD56e4eAb23cb81f43168F9F45211Eb027b9aC7cc"),
							Topics: []common.Hash{
								common.HexToHash("0x87e46b0a6199bc734632187269a103c05714ee0adae5b28f30723955724f37ef"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000d1d85874bb1"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           39,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
							Topics: []common.Hash{
								common.HexToHash("0xe9bded5f24a4168e4f3bf44e00298c993b22376aad8c58c7dda9718a54cbea82"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000002540000000000062c13006f701a95707a0290ac8b90b3719e8ee5b210360883006a9d1b1669c73b033dfe47ae5a0164ab96df25b944000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013a49a6637200000000000000000000000000000000000000000000000000000000a8926f3100000000000000000000000000000000000000000000000000000000000d8e8400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000019ea9400000000000000000000000000000000000000000000000000000000a8b9e84900000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000014a9fddd748b21898c7e26e05ad895760fdd2723bd0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           40,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x701a95707A0290AC8B90b3719e8EE5b210360883"),
							Topics: []common.Hash{
								common.HexToHash("0x8d3ee0df6a4b7e82a7f20a763f1c6826e6176323e655af64f32318827d2112d4"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000062c13"),
							BlockNumber:     lo.Must(new(big.Int).SetString("110308341", 0)),
							TransactionHash: common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
							Index:           41,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40"),
						TransactionIndex: 6,
					},
				},
				config: &config.Module{
					Network: network.Optimism,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Optimism),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xa445b44b0d5da6487ff45c951fd3ec1a23c8802f41c53d95203cd8b6658e6e40",
				Network: network.Optimism,
				Index:   6,
				From:    "0xA9fdDD748b21898C7E26E05ad895760FDd2723Bd",
				To:      stargate.AddressRouterOptimism.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x9fbf10fc",
				},
				Platform: workerx.PlatformStargate.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("33002802845661")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
						From: "0xA9fdDD748b21898C7E26E05ad895760FDd2723Bd",
						To:   layerzero.AddressUltraLightNodeArbitrumOne.String(),
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1482831510889074"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type: typex.TransactionTransfer,
						From: "0xA9fdDD748b21898C7E26E05ad895760FDd2723Bd",
						To:   "0xDecC0c09c3B5f6e92EF4184125D5648a66E35298",
						Metadata: metadata.TransactionTransfer{
							Address:  lo.ToPtr("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2586904"))),
							Name:     "USD Coin",
							Symbol:   "USDC",
							Decimals: 6,
							Standard: metadata.StandardERC20,
						},
					},
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformStargate.String(),
						From:     "0xA9fdDD748b21898C7E26E05ad895760FDd2723Bd",
						To:       "0xA9fdDD748b21898C7E26E05ad895760FDd2723Bd",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Optimism,
							TargetNetwork: network.Avalanche,
							Token: metadata.Token{
								Address:  lo.ToPtr("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2828169009"))),
								Name:     "USD Coin",
								Symbol:   "USDC",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696215459,
			},
			wantError: require.NoError,
		},
		{
			name: "Withdrawal USDT from Ethereum Arbitrum One to Ethereum Polygon on Ethereum Polygon",
			arguments: arguments{
				task: &source.Task{
					Network: network.Polygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x1105e484b0308c1ffb3ed86adf3a3895cf36189aadb177a34bcd74c800d61915"),
						ParentHash:   common.HexToHash("0x16df08e2f66f12d94b2c63b58be6f92bc44cd025e27fad5aa2b70368892ed9e7"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x90b11143a0CB64e067402307Bc7f2276dcEC8250"),
						Number:       lo.Must(new(big.Int).SetString("48222949", 0)),
						GasLimit:     29824656,
						GasUsed:      22667970,
						Timestamp:    1696216798,
						BaseFee:      lo.Must(new(big.Int).SetString("11915200554", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
						From:      common.HexToAddress("0xe93685f3bBA03016F02bD1828BaDD6195988D950"),
						Gas:       1308608,
						GasPrice:  lo.Must(new(big.Int).SetString("52495269297", 10)),
						Hash:      common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
						Input:     hexutil.MustDecode("0x252f7b01000000000000000000000000000000000000000000000000000000000000006e0000000000000000000000009d1b1669c73b033dfe47ae5a0164ab96df25b9440000000000000000000000000000000000000000000000000000000000043238959a267b005b9e91e5a8735dd8be0dc6b838069ce4c51dd07d10d1eca590aeed959a267b005b9e91e5a8735dd8be0dc6b838069ce4c51dd07d10d1eca590aeed00000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000002740000000000000000000000004d73adb72bc3dd368966edd0f0b2148401a178e2000000000017a39d006e352d8275aae3e0c2404d9f68f6cee084b5beb3dd006d9d1b1669c73b033dfe47ae5a0164ab96df25b94400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f39d4fa1270000000000000000000000000000000000000000000000000000000006cc0e0300000000000000000000000000000000000000000000000000000000000061ea00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001bdfe0000000000000000000000000000000000000000000000000000000006ce2deb00000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000014536ef5ebb4682981babbafa29b93e43ca5b285110000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x75dC8e5F50C8221a82CA6aF64aF811caA983B65f")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x1105e484b0308c1ffb3ed86adf3a3895cf36189aadb177a34bcd74c800d61915"),
						BlockNumber:       lo.Must(new(big.Int).SetString("48222949", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 4482327,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xc38f63db1"),
						GasUsed:           214289,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2"),
							Topics: []common.Hash{
								common.HexToHash("0x2bd2d8a84b748439fd50d79a49502b4eb5faa25b864da6a9ab5c150704be9a4d"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000006e"),
								common.HexToHash("0x0000000000000000000000009d1b1669c73b033dfe47ae5a0164ab96df25b944"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000017a39d807b9b7da968928b8cafd63b4f2244e17c19d2237d816568bf7fbaa158be90470000000000000000000000000000000000000000000000000000000000000014352d8275aae3e0c2404d9f68f6cee084b5beb3dd000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("48222949", 0)),
							TransactionHash: common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
							Index:           227,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x29e38769f23701A2e4A8Ef0492e19dA4604Be62c"),
							Topics: []common.Hash{
								common.HexToHash("0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000006e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f39d4fa127"),
							BlockNumber:     lo.Must(new(big.Int).SetString("48222949", 0)),
							TransactionHash: common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
							Index:           228,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000029e38769f23701a2e4a8ef0492e19da4604be62c"),
								common.HexToHash("0x000000000000000000000000536ef5ebb4682981babbafa29b93e43ca5b28511"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000006cc0e03"),
							BlockNumber:     lo.Must(new(big.Int).SetString("48222949", 0)),
							TransactionHash: common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
							Index:           229,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x29e38769f23701A2e4A8Ef0492e19dA4604Be62c"),
							Topics: []common.Hash{
								common.HexToHash("0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000536ef5ebb4682981babbafa29b93e43ca5b285110000000000000000000000000000000000000000000000000000000006cc0e03000000000000000000000000000000000000000000000000000000000001bdfe00000000000000000000000000000000000000000000000000000000000061ea"),
							BlockNumber:     lo.Must(new(big.Int).SetString("48222949", 0)),
							TransactionHash: common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
							Index:           230,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000e93685f3bba03016f02bd1828badd6195988d950"),
								common.HexToHash("0x00000000000000000000000090b11143a0cb64e067402307bc7f2276dcec8250"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000001ee4d73a7070f7000000000000000000000000000000000000000000006a2b093661585a3e73a0000000000000000000000000000000000000000000000013c55d1ba06b2ee3cd000000000000000000000000000000000000000000006a2b09177c811fce02a9000000000000000000000000000000000000000000000013c57c0077a59f54c4"),
							BlockNumber:     lo.Must(new(big.Int).SetString("48222949", 0)),
							TransactionHash: common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
							Index:           231,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b"),
						TransactionIndex: 18,
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
				ID:      "0xd0793bf3462b2764afdefb71716543bcded109439457e65fa848e81d04e4079b",
				Network: network.Polygon,
				Index:   18,
				From:    "0xe93685f3bBA03016F02bD1828BaDD6195988D950",
				To:      layerzero.AddressRelayerPolygon.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x252f7b01",
				},
				Platform: workerx.PlatformStargate.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("11249158762384833")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformStargate.String(),
						From:     "0x29e38769f23701A2e4A8Ef0492e19dA4604Be62c",
						To:       "0x536ef5EBB4682981babBAfA29B93E43CA5b28511",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.Arbitrum,
							TargetNetwork: network.Polygon,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("114036227"))),
								Name:     "(PoS) Tether USD",
								Symbol:   "USDT",
								Decimals: 6,
								Standard: metadata.StandardERC20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1696216798,
			},
			wantError: require.NoError,
		},
		{
			name: "Deposit ETH from Ethereum Linea to Ethereum Base on Ethereum Linea",
			arguments: arguments{
				task: &source.Task{
					Network: network.Linea,
					ChainID: 59144,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x61463d9c64f9f0cf57bb58e9e4f09627070ce61a15a55cc1edce5df69c370539"),
						ParentHash:   common.HexToHash("0x0451306882437c58b2e9194549df844776f782093d6af2c763883f9a34e882b0"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("3129166", 0)),
						GasLimit:     61000000,
						GasUsed:      16343539,
						Timestamp:    1711374637,
						BaseFee:      lo.Must(new(big.Int).SetString("7", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
						From:      common.HexToAddress("0xD7f9e3e2A28b7D0f0a534511f2a4cAFA7e517a6c"),
						Gas:       452854,
						GasPrice:  lo.Must(new(big.Int).SetString("746888043", 10)),
						Hash:      common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
						Input:     hexutil.MustDecode("0x1114cd2a00000000000000000000000000000000000000000000000000000000000000b8000000000000000000000000d7f9e3e2a28b7d0f0a534511f2a4cafa7e517a6c00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000006379da05b600000000000000000000000000000000000000000000000000000062fa85bb7140000000000000000000000000000000000000000000000000000000000000000014d7f9e3e2a28b7d0f0a534511f2a4cafa7e517a6c000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x8731d54E9D02c286767d56ac03e8037C07e01e98")),
						Value:     lo.Must(new(big.Int).SetString("28255130018775908", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("59144", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x61463d9c64f9f0cf57bb58e9e4f09627070ce61a15a55cc1edce5df69c370539"),
						BlockNumber:       lo.Must(new(big.Int).SetString("3129166", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 13005303,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x2c849b6b"),
						GasUsed:           387665,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x224D8Fd7aB6AD4c6eb4611Ce56EF35Dec2277F03"),
							Topics: []common.Hash{
								common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
								common.HexToHash("0x0000000000000000000000008731d54e9d02c286767d56ac03e8037c07e01e98"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000006379da05b60000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           134,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x224D8Fd7aB6AD4c6eb4611Ce56EF35Dec2277F03"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x0000000000000000000000008731d54e9d02c286767d56ac03e8037c07e01e98"),
								common.HexToHash("0x0000000000000000000000002f6f07cdcf3588944bf4c42ac74ff24bf56e7590"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000006379da05b60000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           135,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAad094F6A75A14417d39f04E690fC216f080A41a"),
							Topics: []common.Hash{
								common.HexToHash("0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000b8000000000000000000000000000000000000000000000000000000000000000d0000000000000000000000008731d54e9d02c286767d56ac03e8037c07e01e98000000000000000000000000000000000000000000000000006363a4452c82940000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006ee32813d6c00000000000000000000000000000000000000000000000000000cbba106e0000000000000000000000000000000000000000000000000000000028bed016000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           136,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x224D8Fd7aB6AD4c6eb4611Ce56EF35Dec2277F03"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000008731d54e9d02c286767d56ac03e8037c07e01e98"),
								common.HexToHash("0x000000000000000000000000aad094f6a75a14417d39f04e690fc216f080a41a"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000006379da05b60000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           137,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAad094F6A75A14417d39f04E690fC216f080A41a"),
							Topics: []common.Hash{
								common.HexToHash("0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000b8000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001aa971cea4bb288b0e"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           138,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xA658742d33ebd2ce2F0bdFf73515Aa797Fd161D9"),
							Topics: []common.Hash{
								common.HexToHash("0xdf21c415b78ed2552cc9971249e32a053abce6087a0ae0fbf3f78db5174a3493"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000a25ccfdbdb19"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           139,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x38dE71124f7a447a01D67945a51eDcE9FF491251"),
							Topics: []common.Hash{
								common.HexToHash("0xb0c632f55f1e1b3b2c3d82f41ee4716bb4c00f0f5d84cdafc141581bb8757a4f"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000220001000000000000000000000000000000000000000000000000000000000002ab98000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           140,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xD56e4eAb23cb81f43168F9F45211Eb027b9aC7cc"),
							Topics: []common.Hash{
								common.HexToHash("0x87e46b0a6199bc734632187269a103c05714ee0adae5b28f30723955724f37ef"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000045ad4648284b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           141,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x38dE71124f7a447a01D67945a51eDcE9FF491251"),
							Topics: []common.Hash{
								common.HexToHash("0xe9bded5f24a4168e4f3bf44e00298c993b22376aad8c58c7dda9718a54cbea82"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000254000000000003ff0400b745f1a95a4d3f3836523f5c83673c797f4d4d263b00b8af54be5b6eec24d6bfacf1cce4eaf680a82393980000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000000d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001aa971cea4bb288b0e000000000000000000000000000000000000000000000000006363a4452c8294000000000000000000000000000000000000000000000000000006ee32813d6c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000028bed01600000000000000000000000000000000000000000000000000000000cbba106e0000000000000000000000000000000000000000000000000000063774e18b4a00000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000014d7f9e3e2a28b7d0f0a534511f2a4cafa7e517a6c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           142,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x45f1A95A4D3f3836523F5c83673c797f4d4d263B"),
							Topics: []common.Hash{
								common.HexToHash("0x8d3ee0df6a4b7e82a7f20a763f1c6826e6176323e655af64f32318827d2112d4"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000003ff04"),
							BlockNumber:     lo.Must(new(big.Int).SetString("3129166", 0)),
							TransactionHash: common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
							Index:           143,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d"),
						TransactionIndex: 74,
					},
				},
				config: &config.Module{
					Network: network.Linea,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Linea),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x717dd473abb35804f76ddb3b74f395d7e6fe8d09959eb346b6420001735f5e9d",
				Network: network.Linea,
				Index:   74,
				From:    "0xD7f9e3e2A28b7D0f0a534511f2a4cAFA7e517a6c",
				To:      "0x8731d54E9D02c286767d56ac03e8037C07e01e98",
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x1114cd2a",
				},
				Platform: workerx.PlatformStargate.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("289542353189595")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
						From: "0xD7f9e3e2A28b7D0f0a534511f2a4cAFA7e517a6c",
						To:   "0x38dE71124f7a447a01D67945a51eDcE9FF491251",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("255130018775908"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type: typex.TransactionTransfer,
						From: "0xD7f9e3e2A28b7D0f0a534511f2a4cAFA7e517a6c",
						To:   "0xAad094F6A75A14417d39f04E690fC216f080A41a",
						Metadata: metadata.TransactionTransfer{
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("24420119313772"))),
							Name:     "Ethereum",
							Symbol:   "ETH",
							Decimals: 18,
						},
					},
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformStargate.String(),
						From:     "0xD7f9e3e2A28b7D0f0a534511f2a4cAFA7e517a6c",
						To:       "0xD7f9e3e2A28b7D0f0a534511f2a4cAFA7e517a6c",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeDeposit,
							SourceNetwork: network.Linea,
							TargetNetwork: network.Base,
							Token: metadata.Token{
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("27975579880686228"))),
								Name:     "Ethereum",
								Symbol:   "ETH",
								Decimals: 18,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1711374637,
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
