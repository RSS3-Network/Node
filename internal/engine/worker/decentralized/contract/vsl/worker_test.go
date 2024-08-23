package vsl_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/vsl"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
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
			name: "Withdraw RSS3 from L2 to L1",
			arguments: arguments{
				task: &source.Task{
					Network: network.Ethereum,
					ChainID: 1,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xb4dda29fcd3d61d8bad67ee620a9ec9e4c5469109cc21fbe406b72506363ecb1"),
						ParentHash:   common.HexToHash("0x2460da053f0aafe6a673216b2946bb092f90bc0d878e1c90e8d160791974a747"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"),
						Number:       lo.Must(new(big.Int).SetString("19481383", 0)),
						GasLimit:     30000000,
						GasUsed:      11226005,
						Timestamp:    1711004039,
						BaseFee:      lo.Must(new(big.Int).SetString("26223695286", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
						From:      common.HexToAddress("0x30286DD245338292F319809935a1037CcD4573Ea"),
						Gas:       664131,
						GasPrice:  lo.Must(new(big.Int).SetString("26267959927", 10)),
						Hash:      common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
						Input:     hexutil.MustDecode("0x8c3152e9000000000000000000000000000000000000000000000000000000000000002000010000000000000000000000000000000000000000000000000000000000040000000000000000000000004200000000000000000000000000000000000007000000000000000000000000892caa506c86c5101f5ec11c6f09589c9dc8a85c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004678800000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e4d764ad0b000100000000000000000000000000000000000000000000000000000000000400000000000000000000000042000000000000000000000000000000000000100000000000000000000000004cbab69108aa72151eda5a3c164ea86845f184380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000e40166a07a000000000000000000000000c98d64da73a6616c42117b582e832812e7b8d57f000000000000000000000000420000000000000000000000000000000000004200000000000000000000000030286dd245338292f319809935a1037ccd4573ea00000000000000000000000030286dd245338292f319809935a1037ccd4573ea00000000000000000000000000000000000000000000010f0cf064dd5920000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x6A12432491bbbE8d3babf75F759766774C778Db4")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("1", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xb4dda29fcd3d61d8bad67ee620a9ec9e4c5469109cc21fbe406b72506363ecb1"),
						BlockNumber:       lo.Must(new(big.Int).SetString("19481383", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 8567785,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x61db14277"),
						GasUsed:           222625,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000004cbab69108aa72151eda5a3c164ea86845f18438"),
								common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000010f0cf064dd59200000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19481383", 0)),
							TransactionHash: common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
							Index:           192,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4cbab69108Aa72151EDa5A3c164eA86845f18438"),
							Topics: []common.Hash{
								common.HexToHash("0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3"),
								common.HexToHash("0x000000000000000000000000c98d64da73a6616c42117b582e832812e7b8d57f"),
								common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000042"),
								common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea00000000000000000000000000000000000000000000010f0cf064dd5920000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19481383", 0)),
							TransactionHash: common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
							Index:           193,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x4cbab69108Aa72151EDa5A3c164eA86845f18438"),
							Topics: []common.Hash{
								common.HexToHash("0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd"),
								common.HexToHash("0x000000000000000000000000c98d64da73a6616c42117b582e832812e7b8d57f"),
								common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000042"),
								common.HexToHash("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000030286dd245338292f319809935a1037ccd4573ea00000000000000000000000000000000000000000000010f0cf064dd5920000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19481383", 0)),
							TransactionHash: common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
							Index:           194,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x892CAa506c86C5101f5eC11C6f09589c9dC8A85C"),
							Topics: []common.Hash{
								common.HexToHash("0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c"),
								common.HexToHash("0x21e852a360908a2e966ab1ccce08f7b938ed2e52c30d664876620a37e4347d07"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("19481383", 0)),
							TransactionHash: common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
							Index:           195,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x6A12432491bbbE8d3babf75F759766774C778Db4"),
							Topics: []common.Hash{
								common.HexToHash("0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b"),
								common.HexToHash("0xaf9ccf435b0ad249c09835b34e1ba59ab923c97bd6a20b14eefde0a45f514cfc"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("19481383", 0)),
							TransactionHash: common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
							Index:           196,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec"),
						TransactionIndex: 140,
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
				ID:      "0xafd8a40579b52ffb424fca87d17442b915029202fa104c9bee2bd805d12f7eec",
				Network: network.Ethereum,
				Index:   140,
				From:    "0x30286DD245338292F319809935a1037CcD4573Ea",
				To:      vsl.AddressL1OptimismPortal.String(),
				Type:    typex.TransactionBridge,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x8c3152e9",
				},
				Platform: workerx.PlatformVSL.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("5847904578748375")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.TransactionBridge,
						Platform: workerx.PlatformVSL.String(),
						From:     "0x30286DD245338292F319809935a1037CcD4573Ea",
						To:       "0x30286DD245338292F319809935a1037CcD4573Ea",
						Metadata: metadata.TransactionBridge{
							Action:        metadata.ActionTransactionBridgeWithdraw,
							SourceNetwork: network.VSL,
							TargetNetwork: network.Ethereum,
							Token: metadata.Token{
								Address:  lo.ToPtr("0xc98D64DA73a6616c42117b582e832812e7B8D57F"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("5000000000000000000000"))),
								Name:     "RSS3",
								Symbol:   "RSS3",
								Decimals: 18,
								Standard: 20,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1711004039,
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
