package crossbell_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/v2/config"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	worker "github.com/rss3-network/node/v2/internal/engine/worker/decentralized/contract/crossbell"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/v2/schema/worker/decentralized"
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
		// {
		// 	name: "Crossbell Create Profile",
		// 	arguments: arguments{
		// 		task: &protocol.Task{
		// 			Network: network.Crossbell,
		// 			ChainID: 3737,
		// 			Header: &ethereum.Header{
		// 				Hash:         common.HexToHash("0x208cd7fa40c76bd071edc0db8f6d107b70e194cd27e69027e7c642523f15640d"),
		// 				ParentHash:   common.HexToHash("0xe9de8de4fc50554480ee7018abf94a0876af9b4db6f50d0dd57ec50325635df8"),
		// 				UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
		// 				Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
		// 				Number:       lo.Must(new(big.Int).SetString("5277407", 0)),
		// 				GasLimit:     8000000,
		// 				GasUsed:      336353,
		// 				Timestamp:    1655636500,
		// 				BaseFee:      nil,
		// 				Transactions: nil,
		// 			},
		// 			Transaction: &ethereum.Transaction{
		// 				BlockHash: common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
		// 				From:      common.HexToAddress("0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101"),
		// 				Gas:       409471,
		// 				GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
		// 				Hash:      common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
		// 				Input:     hexutil.MustDecode("0xbd5f69cb00000000000000000000000000000000000000000000000000000000000000200000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd657810100000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000362656800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b726569646a346a64707778693661697a343375626f6b76356333746e716c66346e66706b6f737332727973666a7a726465746769656b6d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000"),
		// 				To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
		// 				Value:     lo.Must(new(big.Int).SetString("0", 0)),
		// 				Type:      0,
		// 				ChainID:   nil,
		// 			},
		// 			Receipt: &ethereum.Receipt{
		// 				BlockHash:         common.HexToHash("0x208cd7fa40c76bd071edc0db8f6d107b70e194cd27e69027e7c642523f15640d"),
		// 				BlockNumber:       lo.Must(new(big.Int).SetString("5277407", 0)),
		// 				ContractAddress:   nil,
		// 				CumulativeGasUsed: 336353,
		// 				EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
		// 				GasUsed:           336353,
		// 				Logs: []*ethereum.Log{{
		// 					Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
		// 					Topics: []common.Hash{
		// 						common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
		// 						common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		// 						common.HexToHash("0x0000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd6578101"),
		// 						common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000df1"),
		// 					},
		// 					Data:            nil,
		// 					BlockNumber:     lo.Must(new(big.Int).SetString("5277407", 0)),
		// 					TransactionHash: common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
		// 					Index:           0,
		// 					Removed:         false,
		// 				}, {
		// 					Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
		// 					Topics: []common.Hash{
		// 						common.HexToHash("0xa5802a04162552328d75eaac538a033704a7c3beab65d0a83e52da1c8c9b7cdf"),
		// 						common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000df1"),
		// 						common.HexToHash("0x0000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd6578101"),
		// 						common.HexToHash("0x0000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd6578101"),
		// 					},
		// 					Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000062af021400000000000000000000000000000000000000000000000000000000000000036265680000000000000000000000000000000000000000000000000000000000"),
		// 					BlockNumber:     lo.Must(new(big.Int).SetString("5277407", 0)),
		// 					TransactionHash: common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
		// 					Index:           1,
		// 					Removed:         false,
		// 				}},
		// 				Status:           1,
		// 				TransactionHash:  common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
		// 				TransactionIndex: 0,
		// 			},
		// 		},
		// 		config: &config.Module{
		// 			Network: network.Crossbell,
		// 			Endpoint: config.Endpoint{
		// 				URL: endpoint.MustGet(network.Crossbell),
		// 			},
		// 		},
		// 	},
		// 	want: &activityx.Activity{
		// 		ID:      "0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3",
		// 		Network: network.Crossbell,
		// 		Index:   0,
		// 		From:    "0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101",
		// 		To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
		// 		Type:    typex.SocialProfile,
		// 		Calldata: &activityx.Calldata{
		// 			FunctionHash: "0xbd5f69cb",
		// 		},
		// 		Platform: workerx.PlatformCrossbell.String(),
		// 		Fee: &activityx.Fee{
		// 			Amount:  lo.Must(decimal.NewFromString("336353000000000")),
		// 			Decimal: 18,
		// 		},
		// 		Actions: []*activityx.Action{
		// 			{
		// 				Type:     typex.SocialProfile,
		// 				Platform: workerx.PlatformCrossbell.String(),
		// 				From:     "0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101",
		// 				To:       "0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101",
		// 				Metadata: metadata.SocialProfile{
		// 					ProfileID: "3569",
		// 					Action:    metadata.ActionSocialProfileCreate,
		// 					Address:   common.HexToAddress("0x3f54953e56a0fe272f3e1e61615bda1fd6578101"),
		// 					Handle:    "beh.csb",
		// 					ImageURI:  "ipfs://bafkreidj4jdpwxi6aiz43ubokv5c3tnqlf4nfpkoss2rysfjzrdetgiekm",
		// 				},
		// 			},
		// 		},
		// 		Status:    true,
		// 		Timestamp: 1655636500,
		// 	},
		// 	wantError: require.NoError,
		// },
		{
			name: "Crossbell Set Profile URI",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x395944fc954a4f6854f47b2bcbadc0f68ceb2980c9efdd8e32f435ec576610e3"),
						ParentHash:   common.HexToHash("0xd7c2901af3c02cb9af8675c4453af349ee25a37bd7e6ec51018303ebe5a28f71"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("4390809", 0)),
						GasLimit:     8000000,
						GasUsed:      51306,
						Timestamp:    1654749902,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x24bc9af800a8d543e2a8cafd0015c5cb9166203e5114f36fd28dbc244a65a659"),
						From:      common.HexToAddress("0x08d66b34054a174841e2361bd4746Ff9F4905cC2"),
						Gas:       51677,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x24bc9af800a8d543e2a8cafd0015c5cb9166203e5114f36fd28dbc244a65a659"),
						Input:     hexutil.MustDecode("0x7c392b51000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b72656967697063623467366a7777746d373764637263796a6565636473736874716573666134326a35346b6c6478617a74776537737a79000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x395944fc954a4f6854f47b2bcbadc0f68ceb2980c9efdd8e32f435ec576610e3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("4390809", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 51306,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           51306,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0xc6b6b2c87fb0784186dd03398c7203c866d0ae59539fa3158aecbc86cb587a95"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000001e"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b72656967697063623467366a7777746d373764637263796a6565636473736874716573666134326a35346b6c6478617a74776537737a79000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("4390809", 0)),
							TransactionHash: common.HexToHash("0x24bc9af800a8d543e2a8cafd0015c5cb9166203e5114f36fd28dbc244a65a659"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x24bc9af800a8d543e2a8cafd0015c5cb9166203e5114f36fd28dbc244a65a659"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x24bc9af800a8d543e2a8cafd0015c5cb9166203e5114f36fd28dbc244a65a659",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x08d66b34054a174841e2361bd4746Ff9F4905cC2",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialProfile,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x7c392b51",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("51306000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialProfile,
						Platform: workerx.PlatformCrossbell.String(),
						From:     "0x08d66b34054a174841e2361bd4746Ff9F4905cC2",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialProfile{
							ProfileID: "30",
							Action:    metadata.ActionSocialProfileUpdate,
							Address:   common.HexToAddress("0x08d66b34054a174841e2361bd4746ff9f4905cc2"),
							Handle:    "song.csb",
							ImageURI:  "ipfs://QmSX9QiwjTGBk5m22UscTg3vrbMwUfFsmxVzMH57hkPD5U/1008.png",
						},
					},
				},
				Status:    true,
				Timestamp: 1654749902,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Create Character",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x73c391ff1d8cb24ba51393f006ae5b2c7f36cbd6ef08a0d29a5ca7a4f0075929"),
						ParentHash:   common.HexToHash("0x3820581f4704a63fae9776208c3cf28306e491e33355cbabec84a4be56367d1b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("13449044", 0)),
						GasLimit:     60000000,
						GasUsed:      327967,
						Timestamp:    1663808139,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef"),
						From:      common.HexToAddress("0xe9c57C291340Ef34DB3646A10af99FE2A0E03827"),
						Gas:       332706,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef"),
						Input:     hexutil.MustDecode("0xcd69fe610000000000000000000000000000000000000000000000000000000000000020000000000000000000000000e9c57c291340ef34db3646a10af99fe2a0e0382700000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000007706f6c65627567000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b726569617477763363706267357174346374336e67666a656c6c65377667356267776f36326261776574776c3373616c6d7479626b786100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x73c391ff1d8cb24ba51393f006ae5b2c7f36cbd6ef08a0d29a5ca7a4f0075929"),
						BlockNumber:       lo.Must(new(big.Int).SetString("13449044", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 327967,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           327967,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x000000000000000000000000e9c57c291340ef34db3646a10af99fe2a0e03827"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000081a0"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("13449044", 0)),
							TransactionHash: common.HexToHash("0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000081a0"),
								common.HexToHash("0x000000000000000000000000e9c57c291340ef34db3646a10af99fe2a0e03827"),
								common.HexToHash("0x000000000000000000000000e9c57c291340ef34db3646a10af99fe2a0e03827"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000632bb28b0000000000000000000000000000000000000000000000000000000000000007706f6c6562756700000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("13449044", 0)),
							TransactionHash: common.HexToHash("0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef"),
							Index:           1,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef",
				Network: network.Crossbell,
				Index:   0,
				From:    "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialProfile,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xcd69fe61",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("327967000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialProfile,
						Platform: workerx.PlatformCrossbell.String(),
						From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						To:       "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						Metadata: metadata.SocialProfile{
							ProfileID: "33184",
							Name:      "polebug",
							Action:    metadata.ActionSocialProfileCreate,
							Address:   common.HexToAddress("0xe9c57c291340ef34db3646a10af99fe2a0e03827"),
							Handle:    "polebug.csb",
						},
					},
				},
				Status:    true,
				Timestamp: 1663808139,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Character Set Handle",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x76db549a236c3979eee7ef994b0bd0c2a99e3e111c7f91307e1ce62341aec173"),
						ParentHash:   common.HexToHash("0x1a6a2e50dde83aaa834b20142f17396f12eb57245004b2f20a00205f29c76a5b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("35330381", 0)),
						GasLimit:     60000000,
						GasUsed:      56047,
						Timestamp:    1685689476,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x96a778fa21f708c24de5f995256155573815d96498de490155366f92aae512c7"),
						From:      common.HexToAddress("0xe9c57C291340Ef34DB3646A10af99FE2A0E03827"),
						Gas:       72305,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x96a778fa21f708c24de5f995256155573815d96498de490155366f92aae512c7"),
						Input:     hexutil.MustDecode("0xa6e6178d00000000000000000000000000000000000000000000000000000000000081a00000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000a706f6c656275675f763100000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x76db549a236c3979eee7ef994b0bd0c2a99e3e111c7f91307e1ce62341aec173"),
						BlockNumber:       lo.Must(new(big.Int).SetString("35330381", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 56047,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           56047,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e"),
								common.HexToHash("0x000000000000000000000000e9c57c291340ef34db3646a10af99fe2a0e03827"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000081a0"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000a706f6c656275675f763100000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("35330381", 0)),
							TransactionHash: common.HexToHash("0x96a778fa21f708c24de5f995256155573815d96498de490155366f92aae512c7"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x96a778fa21f708c24de5f995256155573815d96498de490155366f92aae512c7"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x96a778fa21f708c24de5f995256155573815d96498de490155366f92aae512c7",
				Network: network.Crossbell,
				Index:   0,
				From:    "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialProfile,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xa6e6178d",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("56047000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialProfile,
						Platform: workerx.PlatformCrossbell.String(),
						From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialProfile{
							ProfileID: "33184",
							Action:    metadata.ActionSocialProfileUpdate,
							Address:   common.HexToAddress("0xe9c57c291340ef34db3646a10af99fe2a0e03827"),
							Handle:    "polebug_v1.csb",
							Name:      "polebug",
						},
					},
				},
				Status:    true,
				Timestamp: 1685689476,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Character Set Uri",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x5037a92b9e783285779cf7ba1e180146d82424559d4b150efb3abb66feaff1d8"),
						ParentHash:   common.HexToHash("0x5650c621df6dc696e6ca9ff3ba4dea300c7d7cb1e2a218c2d1970c45d91c1485"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("15240402", 0)),
						GasLimit:     60000000,
						GasUsed:      47651,
						Timestamp:    1665599497,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xc6268ec5e794165322787e56aa483f6e1eb61586c8762c45593d0ee65a8740f3"),
						From:      common.HexToAddress("0x0fefeD77Bb715E96f1c35c1a4E0D349563d6f6c0"),
						Gas:       47964,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0xc6268ec5e794165322787e56aa483f6e1eb61586c8762c45593d0ee65a8740f3"),
						Input:     hexutil.MustDecode("0x47f94de7000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b726569617472616874786d633467376a7a737536676271366b62657036717970707176746472617064357167616170696d697633703265000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x5037a92b9e783285779cf7ba1e180146d82424559d4b150efb3abb66feaff1d8"),
						BlockNumber:       lo.Must(new(big.Int).SetString("15240402", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 47651,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           47651,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000000c"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b726569617472616874786d633467376a7a737536676271366b62657036717970707176746472617064357167616170696d697633703265000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("15240402", 0)),
							TransactionHash: common.HexToHash("0xc6268ec5e794165322787e56aa483f6e1eb61586c8762c45593d0ee65a8740f3"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xc6268ec5e794165322787e56aa483f6e1eb61586c8762c45593d0ee65a8740f3"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xc6268ec5e794165322787e56aa483f6e1eb61586c8762c45593d0ee65a8740f3",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x0fefeD77Bb715E96f1c35c1a4E0D349563d6f6c0",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialProfile,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x47f94de7",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("47651000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialProfile,
						Platform: workerx.PlatformCrossbell.String(),
						From:     "0x0fefeD77Bb715E96f1c35c1a4E0D349563d6f6c0",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialProfile{
							ProfileID: "12",
							Action:    metadata.ActionSocialProfileUpdate,
							Address:   common.HexToAddress("0x0fefed77bb715e96f1c35c1a4e0d349563d6f6c0"),
							ImageURI:  "ipfs://bafkreie6eneu4ihz7qcqu3i6jlhbz4wb4sw252vkpep6a4dgviz3iabz6e",
							Handle:    "joshua.csb",
							Name:      "Joshua",
							Bio:       "Founder @RSS3",
						},
					},
				},
				Status:    true,
				Timestamp: 1665599497,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Post Note",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x34905b0848cb58ee88d692d9c4cc095743d85f3a18bc595d206c234d3f43350b"),
						ParentHash:   common.HexToHash("0x7001856f2b5b188435c72baa04a710d664100269d7140178e195a8da01c6cbbf"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("35256662", 0)),
						GasLimit:     60000000,
						GasUsed:      236812,
						Timestamp:    1685615757,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf81f06f80b929e9f8c7bd9bbac6759a79c89d10fb3e7ce3f34d05b3f50253c41"),
						From:      common.HexToAddress("0x0F588318A494e4508A121a32B6670b5494Ca3357"),
						Gas:       300000,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0xf81f06f80b929e9f8c7bd9bbac6759a79c89d10fb3e7ce3f34d05b3f50253c41"),
						Input:     hexutil.MustDecode("0xf316bacd000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000000000000000000000000000000000000000081a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b7265696275686c76356a6f74766a6e77616b6265753771686a616770793772747079747a716a677766326576787a693269676e7a6d7a6500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003468747470733a2f2f747769747465722e636f6d2f30783830692f7374617475732f31363634313938363439393032313337333434000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x34905b0848cb58ee88d692d9c4cc095743d85f3a18bc595d206c234d3f43350b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("35256662", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 236812,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           236812,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000081a0"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000080"),
								common.HexToHash("0xd84ffca7c96c2d4c29116b91d9da8c6d033bfbbc4994dc7a17d55787991390e4"),
							},
							Data:            hexutil.MustDecode("0x416e7955726900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000003468747470733a2f2f747769747465722e636f6d2f30783830692f7374617475732f31363634313938363439393032313337333434000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("35256662", 0)),
							TransactionHash: common.HexToHash("0xf81f06f80b929e9f8c7bd9bbac6759a79c89d10fb3e7ce3f34d05b3f50253c41"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf81f06f80b929e9f8c7bd9bbac6759a79c89d10fb3e7ce3f34d05b3f50253c41"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xf81f06f80b929e9f8c7bd9bbac6759a79c89d10fb3e7ce3f34d05b3f50253c41",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x0F588318A494e4508A121a32B6670b5494Ca3357",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialPost,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xf316bacd",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("236812000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialPost,
						Platform: "xSync",
						From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialPost{
							ProfileID:     "33184",
							Handle:        "polebug.csb",
							PublicationID: "128",
							Body:          "@0x80i @scnace 不行！牙齿要酸掉了",
							ContentURI:    "ipfs://bafkreibuhlv5jotvjnwakbeu7qhjagpy7rtpytzqjgwf2evxzi2ignzmze",
							Media:         []metadata.Media{},
						},
					},
				},
				Status:    true,
				Timestamp: 1685615757,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Comment Note",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xfecf46318eb6be1b87b10c0c10edf6d2520c4f558bd14e780ef9012d41ad375b"),
						ParentHash:   common.HexToHash("0x57beaf65c5c10a7a298db5f468b7ebec8549a3f70acd1e681e34ac0834e2784d"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("34836623", 0)),
						GasLimit:     60000000,
						GasUsed:      175319,
						Timestamp:    1685195718,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x4add740166bde1a1fc64045c3098a28f7335972db3bfb73455df3020d1365e04"),
						From:      common.HexToAddress("0x0F588318A494e4508A121a32B6670b5494Ca3357"),
						Gas:       300000,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x4add740166bde1a1fc64045c3098a28f7335972db3bfb73455df3020d1365e04"),
						Input:     hexutil.MustDecode("0x9a4dec180000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000e200000000000000000000000000000000000000000000000000000000000081a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b72656964343266666a786b776a3569696c70696b756e367135776c796d647a76656e6f793370787976697375667563647033757a6b687900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xfecf46318eb6be1b87b10c0c10edf6d2520c4f558bd14e780ef9012d41ad375b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("34836623", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 175319,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           175319,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000081a0"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000007d"),
								common.HexToHash("0x30d9277151eb220f8df4e7723ab24739e9ff8357aefc69ff287acff4c14e5b13"),
							},
							Data:            hexutil.MustDecode("0x4e6f74650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000e2"),
							BlockNumber:     lo.Must(new(big.Int).SetString("34836623", 0)),
							TransactionHash: common.HexToHash("0x4add740166bde1a1fc64045c3098a28f7335972db3bfb73455df3020d1365e04"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x4add740166bde1a1fc64045c3098a28f7335972db3bfb73455df3020d1365e04"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x4add740166bde1a1fc64045c3098a28f7335972db3bfb73455df3020d1365e04",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x0F588318A494e4508A121a32B6670b5494Ca3357",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialComment,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x9a4dec18",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("175319000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialComment,
						Platform: "xSync",
						From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialPost{
							Body:          "@songkeys 天哪 祝好",
							Media:         []metadata.Media{},
							Handle:        "polebug.csb",
							ProfileID:     "33184",
							PublicationID: "125",
							ContentURI:    "ipfs://bafkreid42ffjxkwj5iilpikun6q5wlymdzvenoy3pxyvisufucdp3uzkhy",
							Tags:          nil,
							Target: &metadata.SocialPost{
								Body:          "咳嗽咳到眼前出现黑影…到医院发现是视网膜脱落。现在我在等待手术😭",
								Media:         []metadata.Media{},
								Handle:        "song.csb",
								ProfileID:     "30",
								PublicationID: "226",
								ContentURI:    "ipfs://bafkreicuzhibec2otyq2s6oqtjmslczzglt6txfbjlcfokr2okgqcyadpu",
								Tags:          nil,
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1685195718,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Set Note Uri",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x2eea61e14b05ac37016435339bc38509eb3eef7f1fbfa7c1bf29aaeddd28db1e"),
						ParentHash:   common.HexToHash("0x4ee8f18dde050b2c356cafbc2748b42472283a7440f688017a08d2f92db9a0d8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("35313427", 0)),
						GasLimit:     60000000,
						GasUsed:      53229,
						Timestamp:    1685672522,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x2b24d08600be7ce57f664d31c1ea0ff0a6ded55d5be8266cfee669305297cc93"),
						From:      common.HexToAddress("0x8a6dDC78E3aA24f4F31980623f489a274b305762"),
						Gas:       80877,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x2b24d08600be7ce57f664d31c1ea0ff0a6ded55d5be8266cfee669305297cc93"),
						Input:     hexutil.MustDecode("0x628b644a000000000000000000000000000000000000000000000000000000000000d8f8000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b72656963377433747767696d61336375353472376674797575743668373477706477746c716f78636436366f783571746b70746f6b6d79000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x2eea61e14b05ac37016435339bc38509eb3eef7f1fbfa7c1bf29aaeddd28db1e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("35313427", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 53229,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           53229,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000d8f8"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b72656963377433747767696d61336375353472376674797575743668373477706477746c716f78636436366f783571746b70746f6b6d79000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("35313427", 0)),
							TransactionHash: common.HexToHash("0x2b24d08600be7ce57f664d31c1ea0ff0a6ded55d5be8266cfee669305297cc93"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x2b24d08600be7ce57f664d31c1ea0ff0a6ded55d5be8266cfee669305297cc93"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x2b24d08600be7ce57f664d31c1ea0ff0a6ded55d5be8266cfee669305297cc93",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialRevise,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x628b644a",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("53229000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialRevise,
						Platform: "xlog",
						From:     "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialPost{
							Title:         "YouTube Top100",
							Body:          "https://cn.noxinfluencer.com/youtube-channel-rank/top-100-all-all-youtuber-sorted-by-subs-weekly\n\n\n1 T-Series https://www.youtube.com/@tseries\n",
							Media:         []metadata.Media{},
							Handle:        "ordi.csb",
							ProfileID:     "55544",
							PublicationID: "7",
							ContentURI:    "ipfs://bafkreic7t3twgima3cu54r7ftyuut6h74wpdwtlqoxcd66ox5qtkptokmy",
							Tags:          nil,
							Target:        nil,
						},
					},
				},
				Status:    true,
				Timestamp: 1685672522,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Delete Note",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x35972b9b40bd7165ab2af893bb16b710796e12454f019eadcb8e9e717d41923e"),
						ParentHash:   common.HexToHash("0xfb5449512b4ae4635a393712acfe18954ab7453e47e2ac633551cf2b65cda364"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("35313474", 0)),
						GasLimit:     60000000,
						GasUsed:      49945,
						Timestamp:    1685672569,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd0b0341118411ec564737458824544d00b75074989c775bac3db47a47dfd0b2a"),
						From:      common.HexToAddress("0x8a6dDC78E3aA24f4F31980623f489a274b305762"),
						Gas:       75474,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0xd0b0341118411ec564737458824544d00b75074989c775bac3db47a47dfd0b2a"),
						Input:     hexutil.MustDecode("0xc2a6fe3b000000000000000000000000000000000000000000000000000000000000d8f80000000000000000000000000000000000000000000000000000000000000007"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x35972b9b40bd7165ab2af893bb16b710796e12454f019eadcb8e9e717d41923e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("35313474", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 49945,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           49945,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000d8f8"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000007"),
							BlockNumber:     lo.Must(new(big.Int).SetString("35313474", 0)),
							TransactionHash: common.HexToHash("0xd0b0341118411ec564737458824544d00b75074989c775bac3db47a47dfd0b2a"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd0b0341118411ec564737458824544d00b75074989c775bac3db47a47dfd0b2a"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xd0b0341118411ec564737458824544d00b75074989c775bac3db47a47dfd0b2a",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialDelete,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xc2a6fe3b",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("49945000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialDelete,
						Platform: "xlog",
						From:     "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialPost{
							Title:         "YouTube Top100",
							Body:          "https://cn.noxinfluencer.com/youtube-channel-rank/top-100-all-all-youtuber-sorted-by-subs-weekly\n\n\n1 T-Series https://www.youtube.com/@tseries\n",
							Media:         []metadata.Media{},
							Handle:        "ordi.csb",
							ProfileID:     "55544",
							PublicationID: "7",
							ContentURI:    "ipfs://bafkreic7t3twgima3cu54r7ftyuut6h74wpdwtlqoxcd66ox5qtkptokmy",
							Tags:          nil,
							Target:        nil,
						},
					},
				},
				Status:    true,
				Timestamp: 1685672569,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Mint Note",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x6c3fb9b4c1153572a35361a2e7771048a3fd4350fb42a0b9e54613d78c764e4e"),
						ParentHash:   common.HexToHash("0xdcd4d21e9779e753c37e1c99dd1341930b25aa68852290986d078952d919cb4b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("29835277", 0)),
						GasLimit:     60000000,
						GasUsed:      416564,
						Timestamp:    1680194372,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
						From:      common.HexToAddress("0xc560eb6fd0c2eb80Df50E5e06715295AE1205049"),
						Gas:       428734,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
						Input:     hexutil.MustDecode("0xa7ccb4bf0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002d00000000000000000000000000000000000000000000000000000000000000bc00000000000000000000000008d66b34054a174841e2361bd4746ff9f4905cc2000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x6c3fb9b4c1153572a35361a2e7771048a3fd4350fb42a0b9e54613d78c764e4e"),
						BlockNumber:       lo.Must(new(big.Int).SetString("29835277", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 416564,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           416564,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x59531Cf802170C646597A9e1715827C9ccB76D92"),
							Topics: []common.Hash{
								common.HexToHash("0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000006425bb44000000000000000000000000000000000000000000000000000000000000000b4e6f74652d34352d313838000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b4e6f74652d34352d313838000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("29835277", 0)),
							TransactionHash: common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59531Cf802170C646597A9e1715827C9ccB76D92"),
							Topics: []common.Hash{
								common.HexToHash("0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000002d00000000000000000000000000000000000000000000000000000000000000bc000000000000000000000000000000000000000000000000000000006425bb44"),
							BlockNumber:     lo.Must(new(big.Int).SetString("29835277", 0)),
							TransactionHash: common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59531Cf802170C646597A9e1715827C9ccB76D92"),
							Topics: []common.Hash{
								common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("29835277", 0)),
							TransactionHash: common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x59531Cf802170C646597A9e1715827C9ccB76D92"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x00000000000000000000000008d66b34054a174841e2361bd4746ff9f4905cc2"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("29835277", 0)),
							TransactionHash: common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
							Index:           3,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212"),
								common.HexToHash("0x00000000000000000000000008d66b34054a174841e2361bd4746ff9f4905cc2"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000002d"),
								common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000bc"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000059531cf802170c646597a9e1715827c9ccb76d920000000000000000000000000000000000000000000000000000000000000001"),
							BlockNumber:     lo.Must(new(big.Int).SetString("29835277", 0)),
							TransactionHash: common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
							Index:           4,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693",
				Network: network.Crossbell,
				Index:   0,
				From:    "0xc560eb6fd0c2eb80Df50E5e06715295AE1205049",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialMint,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xa7ccb4bf",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("416564000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialMint,
						Platform: "Crossbell",
						From:     "0x08d66b34054a174841e2361bd4746Ff9F4905cC2",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialPost{
							Body: "This is to certify that @Songkeys is 文案大师.",
							Media: []metadata.Media{
								{
									MimeType: "image/jpg",
									Address:  "ipfs://bafkreicaj6ehv5wzmzr5uptythkxgrdf5emvgi4merdao3f37dkquq363a",
								},
							},
							Title:         "文案大师",
							Handle:        "atlas.csb",
							ProfileID:     "45",
							PublicationID: "188",
							ContentURI:    "ipfs://bafkreiflhf3mygphlq32f5hud4232h452wh7gmgc42cfwiip2mfgdxpwwm",
						},
					},
				},
				Status:    true,
				Timestamp: 1680194372,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Tip Character for Note",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x45a117f01e5951b3fa2d1adf7263902ca0ceb2ad82ee4fc810cf0fb4894717b7"),
						ParentHash:   common.HexToHash("0x892ef53108bda6d1f08ae14ce569036e22fec6a394be238a6a1b33cf57df1c05"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("42147627", 0)),
						GasLimit:     60000000,
						GasUsed:      91854,
						Timestamp:    1692506722,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
						From:      common.HexToAddress("0x39F9e912C1F696F533e7A2267Ea233AeC9742b35"),
						Gas:       168426,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
						Input:     hexutil.MustDecode("0x9bd9bbc60000000000000000000000000058be0845952d887d1668b5545de995e12e87830000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000002a4d000000000000000000000000000000000000000000000000000000000000b34b0000000000000000000000000000000000000000000000000000000000000500"),
						To:        lo.ToPtr(common.HexToAddress("0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x45a117f01e5951b3fa2d1adf7263902ca0ceb2ad82ee4fc810cf0fb4894717b7"),
						BlockNumber:       lo.Must(new(big.Int).SetString("42147627", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 91854,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           91854,

						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64"),
							Topics: []common.Hash{
								common.HexToHash("0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987"),
								common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
								common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
								common.HexToHash("0x0000000000000000000000000058be0845952d887d1668b5545de995e12e8783"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000002a4d000000000000000000000000000000000000000000000000000000000000b34b00000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("42147627", 0)),
							TransactionHash: common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000039f9e912c1f696f533e7a2267ea233aec9742b35"),
								common.HexToHash("0x0000000000000000000000000058be0845952d887d1668b5545de995e12e8783"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000de0b6b3a7640000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("42147627", 0)),
							TransactionHash: common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
							Index:           1,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64"),
							Topics: []common.Hash{
								common.HexToHash("0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987"),
								common.HexToHash("0x0000000000000000000000000058be0845952d887d1668b5545de995e12e8783"),
								common.HexToHash("0x0000000000000000000000000058be0845952d887d1668b5545de995e12e8783"),
								common.HexToHash("0x000000000000000000000000b461c1521ee9d96a9c950337f0851b79bd66cae1"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000002a4d000000000000000000000000000000000000000000000000000000000000b34b0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("42147627", 0)),
							TransactionHash: common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
							Index:           2,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000058be0845952d887d1668b5545de995e12e8783"),
								common.HexToHash("0x000000000000000000000000b461c1521ee9d96a9c950337f0851b79bd66cae1"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000de0b6b3a7640000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("42147627", 0)),
							TransactionHash: common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
							Index:           3,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0058be0845952D887D1668B5545de995E12e8783"),
							Topics: []common.Hash{
								common.HexToHash("0x3cfc312a05bba2ef0c4efdf6a7a4ef883e9e35f403da1b4cbc810b3ba738dab5"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000002a4d"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000b34b"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000500"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000afb95cc0bd320648b3e8df6223d9cdd05ebedc640000000000000000000000000000000000000000000000000de0b6b3a7640000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("42147627", 0)),
							TransactionHash: common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
							Index:           4,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xd8072e5fd4d0de8a620b665fc99eabe252ac4e1b18347829ec9e2f4e6641b865",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
				To:      "0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64",
				Type:    typex.SocialReward,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x9bd9bbc6",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("91854000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialReward,
						Platform: "xlog",
						From:     "0x39F9e912C1F696F533e7A2267Ea233AeC9742b35",
						To:       "0xb461c1521eE9D96a9C950337f0851b79bD66CAe1",
						Metadata: &metadata.SocialPost{
							Target: &metadata.SocialPost{
								Handle:        "pseudoyu.csb",
								Title:         "周报 #44 - 静寂工人、xLog 与日常",
								Body:          "---\ntitle: \"周报 #44 - 静寂工人、xLog 与日常\"\ndate: 2023-08-15T10:01:24+08:00\ndraft: false\ntags: [\"review\", \"life\", \"romance\", \"love\", \"run\", \"time\", \"boyi\"]\ncategories: [\"Ideas\"]\nauthors:\n- \"pseudoyu\"\n---\n\n## 前言\n\n![weekly_review_20230816](https://image.pseudoyu.com/images/weekly_review_20230816.png)\n\n本篇是对 `2023-07-11` 到 `2023-08-15` 这几周生活的记录与思考。\n\n停滞了将近一个月，再打开周报文档的时候也有了些许陌生感。其实积攒了很多想写的主题，也有很多想说的话，尽管输出变少了，分享欲并未因为生活状态的改变消减，却蔓延在静水流深的日常中。在渐渐习惯了新的生活节奏后，也逐渐开始调整（填补）年初制定的一些计划，周报和博客文章也开始继续更新。\n\n## 静寂工人\n\n![yixi_ticket](https://image.pseudoyu.com/images/yixi_ticket.jpg)\n\n一个周末去听了一席的演讲，有点像是 Ted 的本地化版本，有不少有趣的人，有为了拯救濒危猫头鹰在俄罗斯森林里生活了数年的可爱学者，有探寻山河褶皱残留荒庙的艺术家，自己做书的建筑师...\n\n![yixi_speech](https://image.pseudoyu.com/images/yixi_speech.jpg)\n\n印象最深也最喜欢的是一位专注于“寻短”的心理咨询师魏明毅老师，现场买了书也找作者签了名，这周挺集中地看完了。很喜欢她“人类学学徒”的自称，深入被台湾基隆码头辉煌年代所驱逐与禁锢的一群工人，与他们同行，理解他们的内心世界与生活的视角。全书其实没有多少去渲染什么苦难，他们辉煌过也经历了被时代抛弃的没落，比起很多真正的难以温饱的群体来说还相距甚远。\n\n![book_with_sig](https://image.pseudoyu.com/images/book_with_sig.jpg)\n\n但震撼的是这一群人似乎仅仅在几年间就因为政治/经济方向的一些变动而迅速被遗忘，似乎那段“不夜城”的岁月从未属于过他们，而他们却要花上大半辈子的时间去让社会重新接受自己，同时，也要重新认识自己，对外界而言，他们似乎只是“不够努力”而被抛弃的一群“没用”的人，没有人认真倾听过他们的声音，他们也渐渐学会了沉默。\n\n历史总是在以自己（或是某些政治用意）的标准来决定哪些应该被记录下来，他们就是被略过的笔墨，然而真正的历史感往往在史料之外，那些被掩埋的记忆。我们现今的时代也如是，有了太多不可说不可议的东西，那些反而才是真正的时代的声音。\n\n笔触让我想到之前看过的一部台湾电影《阳光普照》，许多对话、独白和长镜头，就像是像把很多生活的痕迹原样地展现出来，大概也是台湾特有的风格，娓娓道来。对台湾的一些生活方式和时代特色有了很多兴趣，接下来想去看一部之前很早就想看的剧《做工的人》。\n\n会对这一点尤其感兴趣也是因为在我初高中的时光里因为我爸的工作原因，有过跟他的几个学徒住过员工宿舍的经历。那时的室友哥哥们大多没念/没念完高中就到异乡拼搏，想学一门技术在异乡谋求一席之地。\n\n在智能手机和短视频还并不流行的那个时代，也许是因为晚上的时间太过漫长，下班后的他们总是需要找一个地方消磨时光，我也曾有几次被带入他们的生活圈子，有时会去室内轮滑场玩上几个小时，有时会在某个热闹广场的一角聊聊天，也有时候仅仅是在房间里用着老式 DVD 一起看一晚上奥特曼。那时候的我还很小，跟他们的相处和交流自然算不上深入，后续随着各种变动也不再有联系。\n\n十几年后的现在想想，他们当时也只是十八九岁，却似乎有着超过他们年龄的成熟感，他们在很小的年纪就需要在这个热闹城市里找到自己的位置，却又不得不面临着奔波的生活、孤独感以及和这个社会的疏离感，也许只有夜晚的自由时光才能让他们真正感受到一些属于自己的东西。后来有稍微打听了下他们后来各自的生活，有的回老家了，做一些小小的事业；有的继续在杭州打拼，换了一个又一个地方，似乎也还是没能实现自己想要的生活。就像书中描述的那样，他们并不会去指责这个时代，而是更多将生活中不满的部分归咎于自己，仅此而已。\n\n## 个人生活剪影\n\n![boyi_universal](https://image.pseudoyu.com/images/boyi_universal.png)\n\n妹妹放暑假了，于是来北京玩了，和学姐一起去了环球影城，尽管是暑假的周末，但也没想象得那么拥挤。\n\n几个热门的项目都有玩到，在「哈利·波特禁忌之旅」项目时眼镜飞了，至今还没找回来，于是去眼镜城配了一副，大概来到魔法世界总得留下些什么吧。\n\n其他的就是静水流深的日常，缓慢却也美好。有一起见了赴京考察的亮亮和其他老师，聊到了挺多之前的事；会偶尔去探索一些街头巷尾的馆子，也会在工作日的晚上去雍和宫听民谣，热闹中找到片刻歇息；还会在远程办公的间隙接送学姐上下班，等候的间隙跟门口的保安小哥聊聊天...\n\n大约是心境有了些变化，生活状态和看待生活的心情也有了许多改变，很难确切形容，大体是充满了更多美好和善意了吧。\n\n也好久没晒过捏捏的日常了，现在毛越来越厚重，可能下周会带去洗个澡梳理一下毛。\n\n![nienie_2023_0816_01](https://image.pseudoyu.com/images/nienie_2023_0816_01.jpg)\n\n![nienie_2023_0816_02](https://image.pseudoyu.com/images/nienie_2023_0816_02.jpg)\n\n## 有趣的事与物\n\n### xLog\n\n![yu_xlog_profile](https://image.pseudoyu.com/images/yu_xlog_profile.png)\n\nxLog 目前是我自己博客的主要同步和备份平台，由于舍不得之前的数据统计和 Cusdis 积累的评论，一直还没办法整个迁移过去，最新推出了「作品集」功能后，配合着 xSync, xFeed 和 xChar 倒是很好地成为了我的个人 Profile 页面，也是越来越依赖了。\n\n这是我的作品集页面：[xlog.pseudoyu.com/portfolios](https://xlog.pseudoyu.com/portfolios)\n\n![boyi_xlog](https://image.pseudoyu.com/images/boyi_xlog.png)\n\n之前学姐的博客一直使用的是 Hashnode 平台，也算是简易好用的平台，但随着 xLog 功能的更迭和给予创作者的自由度，我还是把它迁移了过来，除了领取 CSB Token 以外，onboarding 体验已经很好了，绑定域名和 umami 解析这些也很方便，建议想拥有自己博客的朋友可以尝试一下。\n\n这是学姐的博客：[boyilu.com](https://www.boyilu.com)\n\n迁移和编辑过程中也都再重新读了一遍，依然觉得文字是我所能想到的了解和慢慢靠近一个人最舒适也最温暖的方式了。\n\n### 个人信息流同步系统\n\n由于 Railway 在 8.1 正式关闭了 Free Plan，n8n 同步系统停了两天，调研了一番后来还是在自己的 vps 上用 docker 运行了（一台搬瓦工的 ThePlan，2c2g，配合 NginxProxyManager 进行反向代理，跑了我大部分需要数据持久化的服务），顺便更新了一下版本，继续进行同步啦。\n\n截至写周报的时候 Telegram 频道已经有了 536 subscribers，渐渐地也形成了自己输入输出的习惯模式，Twitter（现在可能得叫 X 了）虽然渐渐用得少了，也到了 2000+ fo，也算是个小里程碑了。\n\n### 开源预算\n\n在之前的一篇周报「[周报 #30 - 开源预算、写作初心与对技术的谦卑](https://www.pseudoyu.com/zh/2023/02/12/weekly_review_20230212/)」中提到过这样一段：\n\n\u003e 在 Randy 的一篇文章『[我给自己设立了每月 $20 的开源捐赠预算](https://lutaonan.com/blog/my-oss-donation-budget/)』中看到他对于开源项目的理念与态度，觉得很有意思，也引发了我想为自己也设立一个同样的开源预算的想法。目前的设定是每月至少 $20（约 130 元人民币）或等价值的预算，根据自己的日常使用与技术栈灵活选择，我会选择以下项目进行捐赠：\n\u003e\n\u003e - 对我有启发的独立博客作者与开发者\n\u003e - 我在做 side projects 时常用且解决了很切实问题的项目\n\u003e - 我高频使用的一些有趣的开源工具与服务\n\n之前固定给 [reorx](https://github.com/reorx) 在 GitHub Sponsor 赞助，这个月在 Randy 发布了 [Cusdis Pro](https://blog.cusdis.com/announcing-cusdis-pro/) 和 [Notepal](https://notepal.randynamic.org/) 之后开始给 [Randy](https://lutaonan.com/) 赞助。\n\n![sponsor_randy](https://image.pseudoyu.com/images/sponsor_randy.png)\n\n\u003e 💖 I'm sponsoring @randyloop for his passion and dedication to open-source. His products, blogs and podcasts really bring inspiration to my growth as a programmer and life learner.\n\n### Nuphy\n\n![nuphy_keyboard](https://image.pseudoyu.com/images/nuphy_keyboard.png)\n\n自己偶尔会写一些效率工具类的文章，没想到因此被 Nuphy 官方联系送测了 Nuphy Air60 键盘，轻便且颜值高，打算外出都带它了。\n\n### 其他\n\n感觉还发生了很多好玩的事，却因为堆积到月的维度而一时难以都想起来，开始集中冲刺学日语，在出一套教程以及很多想写的文章主题，希望后续能够慢慢补坑。\n\n### 输入\n\n虽然大部分有意思的输入会在 「[Yu's Life](https://t.me/pseudoyulife)」 Telegram 频道里自动同步，不过还是挑选一部分在这里列举一下，感觉更像一个 newsletter 了。\n\n#### 书籍\n\n- **世界尽头与冷酷仙境**，描述得很美，主人公的视角有一种我很喜欢的与世界的疏离感，结局也冷静而不失意外。但是很可惜的是这本书分了好几个月且几乎都是在飞机或是高铁上读的，加上平行的两条世界线的设定，在剧情上断续感很强，也许有机会可以再好好读一遍。\n- **静寂工人**，有挺多感想的，写在第一部分了。\n\n#### 文章\n\n- [Local-first software: You own your data, in spite of the cloud](https://www.inkandswitch.com/local-first/)\n- [一个独立创造者的五年](https://mp.weixin.qq.com/s/x6PLSIMn_1qcKnXWPT-J-Q)\n- [Borg, Omega, and Kubernetes - ACM Queue](https://queue.acm.org/detail.cfm?id=2898444)\n- [The Modular World](https://rainandcoffee.substack.com/p/the-modular-world)\n- [对 Newsletter 说不 - DIYgod](https://diygod.cc/say-no-to-newsletter)\n- [如何复刻本网站，零氪快速建博客 | Cali Castle](https://cali.so/blog/guide-for-cloning-my-site)\n- [《平台工人》丨青山资本 2023 年中消费报告](https://mp.weixin.qq.com/s/j7zKFgxLIsDz12SFxmmaBA)\n- [AI 正在杀死网络, 相对失败与成长, 纯文本 -#39](https://geekplux.com/newsletters/39)\n- [让玄学可靠：构建复杂 LLM 应用](https://mp.weixin.qq.com/s/4ALipJhxCLmRZGguDROyEw)\n- [穿透 Web3](https://mp.weixin.qq.com/s/pacNtxhtAByaGgLUr-BNGQ)\n- [Go 工程化(二) 项目目录结构](https://lailin.xyz/post/go-training-week4-project-layout.html)\n- [The past is not true | Derek Sivers](https://sive.rs/pnt)\n- [发布一年多，终于入手了 Steam Deck](https://houjoe.notion.site/Steam-Deck-69bd595650e141da94146521e5a3c10a)\n- [Hashing](https://samwho.dev/hashing/)\n- [向量数据库](https://guangzhengli.com/blog/zh/vector-database/)\n- [到底是谁的成就感 - Fulfillment Of Platforms](https://someonegao.com/fulfillment-of-platforms)\n- [Thinking something nice about someone? Tell them. | Derek Sivers](https://sive.rs/nice)\n- [数据库碎碎念 - 知乎](https://zhuanlan.zhihu.com/p/645811161)\n- [A Beginner’s Guide to Docker Networking – Collabnix](https://collabnix.com/a-beginners-guide-to-docker-networking/)\n- [Wireshark 常用功能笔记 | NoPanic](https://www.ilikejobs.com/posts/wireshark/)\n- [开发一个浏览器插件在第三天卖出 1000 元 | Randy's Blog](https://lutaonan.com/blog/my-extension-sold-1k-yuan/)\n- [ChromeOS is splitting the browser from the OS, getting more Linux-y | Ars Technica](https://arstechnica.com/gadgets/2023/08/google-is-finally-separating-chrome-from-chromeos-for-easier-updates/)\n- [The Reluctant Sysadmin's Guide to Securing a Linux Server](https://pboyd.io/posts/securing-a-linux-vm/)\n- [掘力计划第 20 期：Tw93 Pake —— 利用 Rust 轻松构建跨端轻量级应用 - 掘金](https://juejin.cn/post/7261897250648932389)\n- [账户抽象：链接 Web3 和 Web2](https://mp.weixin.qq.com/s/lrFk0UWZg0LFq6HTke3YBA)\n- [Pake - 利用 Rust 轻松构建轻量级应用 - Tw93](https://tw93.fun/2023-08-03/pake.html)\n- [使用 Zeabur 构建我的 n8n 赛博空间 | 今是昨非](https://zuofei.net/5091.html)\n- [创造了不起丨 TiKV Committer 的最快养成方法 - 知乎](https://zhuanlan.zhihu.com/p/367623333)\n- [在博客融入一个跨平台作品集 - DIYgod](https://diygod.cc/xlog-portfolios)\n- [全链游戏-重新定义游戏服务端 — dashuo](https://mirror.xyz/shuo.eth/4bFNscq2EBxPTQRhxbyMhDmL6BwuE3eEksY7hkNvWm0)\n- [运动 | 有哪些事是学拳击以后才知道的？](https://www.boyilu.com/boxing-beginner)\n- [运动 | 攀岩手记](https://www.boyilu.com/rock-climbing-record-01)\n- [观影｜Tercett – Zsigmond Móricz’s Women](https://www.boyilu.com/tercett-review)\n- [十斤糖炒栗子那么爱](https://www.boyilu.com/love-grandfather)\n- [快乐的甲板](https://www.boyilu.com/happy-deck)\n- [YC's essential startup advice : YC Startup Library | Y Combinator](https://www.ycombinator.com/library/4D-yc-s-essential-startup-advice)\n\n#### 视频\n\n记录一下看过的有意思的视频：\n\n- [ClippyGPT - How I Built Supabase’s OpenAI Doc Search (Embeddings)](https://www.youtube.com/watch?v=Yhtjd7yGGGA)\n- [How to create an OpenAI Q\u0026A bot with ChatGPT API + embeddings](https://www.youtube.com/watch?v=RM-v7zoYQo0)\n- [入门到毕业，一台全搞定！拍照视频两不误｜a6700 全新评测](https://www.bilibili.com/video/BV1uM4y1x7Sv)\n- [月融资 1.6 亿美元！大模型带火的新风口，竟然是它？| 向量数据库爆火背后的逻辑](https://www.bilibili.com/video/BV1W94y1B7Vd)\n- [【硬核】苹果海外千亿税款如何避税 | 欧美政府税收反击战](https://www.youtube.com/watch?v=dQ2bjo07aNs)\n- [对抗孤独需要肤浅朋友，不要什么精神之交](https://www.bilibili.com/video/BV1vV4y1b76v)\n- [一个小时调研全新领域，能了解到什么程度？](https://www.bilibili.com/video/BV16x4y1R7MH)\n- [北京 City walk：属于老北京的童年记忆，在城南找到了](https://www.bilibili.com/video/BV1L94y1q7ft)\n- [带没去过漫展的 UP 逛 BW，是怎样一种体验？](https://www.bilibili.com/video/BV1hV411L7M8)\n- [有被卷到！剪映——更适合中国人体质的剪辑软件](https://www.youtube.com/watch?v=ydFvA29KtzE)\n- [实录对比丨谁是影响听歌音质的最大因素？](https://www.bilibili.com/video/BV1hk4y1V7oV)\n- [致敬新海诚｜打开往门，奔赴三年之约【关西旅拍】](https://www.bilibili.com/video/BV1814y1B77t)\n- [离谱！Apple TV 也能当软路由？！](https://www.youtube.com/watch?v=1WzhoxdwpeQ)\n- [面试必备！每个人都应该掌握的快速调研法](https://www.bilibili.com/video/BV1pF411f7yD)\n- [90%的摄影师从来都没思考过的问题！每天发作品发了个寂寞！](https://www.bilibili.com/video/BV1Cr4y1C7Bz)\n- [都是演的？！旅拍 UP 主吐槽大会](https://www.bilibili.com/video/BV1X94y1r7e3)\n- [b 站千万播放的“三年动画”什么水平？大佬：有点东西](https://www.bilibili.com/video/BV1f841197c2)\n- [博得之門 Baldur's Gate》為何被稱為傳奇？它復活的意義？【就知道玩遊戲 65】](https://www.youtube.com/watch?v=n5VE49MXWFs)\n\n#### 播客\n\n- [ActivityPub is the next big thing in social](https://www.listennotes.com/e/180f6b6ed2ee4b0e89f23db185f50462)\n- [Vol. 94 是光诗歌: 用诗歌发现大山孩子手里的光](https://www.listennotes.com/e/a08521d0a1024ed9a8d5b3f1cba5f444)\n- [编码人声：面向 AI 的新编程范式](https://www.listennotes.com/e/676c51f85479449ea2503b446937883e)\n- [第 4 集 |「捕蛇者说」的起源、做公益、如何保存数字遗产、遇到过的名人，清华和二本](https://www.listennotes.com/e/40d6d93c94ac4816ad1195f5479c702f)\n\n#### 音乐\n\n- [Zenzenzense - movie ver. by RADWIMPS](https://open.spotify.com/track/2DLrgv7HhJanCuD8L9uJLR)\n- [ありがとう by Takuya Ohashi](https://open.spotify.com/track/7vgvfXICUyLXzgtGBP5Y6Z)\n- [小さな恋のうた by 高木さん](https://open.spotify.com/track/3rysOpH29UDtPpv5W2pzWZ)\n- [21 Guns by Green Day](https://open.spotify.com/track/64yrDBpcdwEdNY9loyEGbX)\n\n#### 剧集\n\n- **东京爱情故事**，重温。\n- **伪装者**，因为还是挺喜欢琅琊榜的这一波人，mark 了很久，找了个周末集中看了下\n",
								Media:         []metadata.Media{},
								ProfileID:     "45899",
								PublicationID: "1280",
								ContentURI:    "ipfs://bafkreicrsytu3gc2kjyruyojejf2eavggc3mm3egnz5nc3wdza56gtq4iq",
								Reward: &metadata.Token{
									Address:  lo.ToPtr(common.HexToAddress("0xAfB95CC0BD320648B3E8Df6223d9CDD05EbeDC64").String()),
									Value:    lo.ToPtr(lo.Must(decimal.NewFromString("1000000000000000000"))),
									Name:     "Mira Token",
									Symbol:   "MIRA",
									Decimals: 18,
									Standard: 20,
								},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692506722,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Set Operator Permission Appoint",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x13017260840c42c075e2b74985964bf75ef7dcdd2daf429ebf605472006727a2"),
						ParentHash:   common.HexToHash("0x2e2bfea2e3ad05d83a5dba93888d8d97e1a73691dbee124ee422df647b8cf563"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("33502728", 0)),
						GasLimit:     60000000,
						GasUsed:      113102,
						Timestamp:    1683861823,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0xd1f402c52687b954bba45d1c1012df743ab0facec8d9212f3b2a2a2d54fcc6d5"),
						From:      common.HexToAddress("0x8a6dDC78E3aA24f4F31980623f489a274b305762"),
						Gas:       115796,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0xd1f402c52687b954bba45d1c1012df743ab0facec8d9212f3b2a2a2d54fcc6d5"),
						Input:     hexutil.MustDecode("0x206657f2000000000000000000000000000000000000000000000000000000000000d8f8000000000000000000000000bbc2918c9003d264c25ecae45b44a846702c0e7c0000100000000fffffff00000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x13017260840c42c075e2b74985964bf75ef7dcdd2daf429ebf605472006727a2"),
						BlockNumber:       lo.Must(new(big.Int).SetString("33502728", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 113102,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           113102,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x4b0dcce4c30a5691ff14b7d4a8a27e2738b66289dc90120ebbc092812941bd27"),
								common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000d8f8"),
								common.HexToHash("0x000000000000000000000000bbc2918c9003d264c25ecae45b44a846702c0e7c"),
							},
							Data:            hexutil.MustDecode("0x0000100000000fffffff00000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("33502728", 0)),
							TransactionHash: common.HexToHash("0xd1f402c52687b954bba45d1c1012df743ab0facec8d9212f3b2a2a2d54fcc6d5"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0xd1f402c52687b954bba45d1c1012df743ab0facec8d9212f3b2a2a2d54fcc6d5"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0xd1f402c52687b954bba45d1c1012df743ab0facec8d9212f3b2a2a2d54fcc6d5",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialProxy,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x206657f2",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("113102000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialProxy,
						Platform: "Crossbell",
						From:     "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
						To:       "0xBBC2918C9003D264c25EcAE45B44a846702C0E7c",
						Metadata: metadata.SocialProxy{
							ProxyAddress: common.HexToAddress("0xbbc2918c9003d264c25ecae45b44a846702c0e7c"),
							Action:       metadata.ActionSocialProxyAppoint,
							Profile: metadata.SocialProfile{
								Handle:    "ordi.csb",
								ProfileID: "55544",
								Address:   common.HexToAddress("0x8a6dDC78E3aA24f4F31980623f489a274b305762"),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1683861823,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Set Operator Permission Remove",
			arguments: arguments{
				task: &source.Task{
					Network: network.Crossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x8134df3d86c7c98ee590abc6cc9e593c6e37cc82e74e15129e45c164ba922ff3"),
						ParentHash:   common.HexToHash("0x9d0dfd3accb4e4485579081ee61abea11e55c2b4058dc2031a9f9f44f4b25ae5"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("35336576", 0)),
						GasLimit:     60000000,
						GasUsed:      30915,
						Timestamp:    1685695671,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x84257dbc6d86831d12735836779db5eb0a9d1848a7d91a6715f9a11ab10bf5af"),
						From:      common.HexToAddress("0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A"),
						Gas:       62869,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x84257dbc6d86831d12735836779db5eb0a9d1848a7d91a6715f9a11ab10bf5af"),
						Input:     hexutil.MustDecode("0x206657f20000000000000000000000000000000000000000000000000000000000007d0a0000000000000000000000000f588318a494e4508a121a32b6670b5494ca33570000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x8134df3d86c7c98ee590abc6cc9e593c6e37cc82e74e15129e45c164ba922ff3"),
						BlockNumber:       lo.Must(new(big.Int).SetString("35336576", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 30915,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           30915,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0x4b0dcce4c30a5691ff14b7d4a8a27e2738b66289dc90120ebbc092812941bd27"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000007d0a"),
								common.HexToHash("0x0000000000000000000000000f588318a494e4508a121a32b6670b5494ca3357"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("35336576", 0)),
							TransactionHash: common.HexToHash("0x84257dbc6d86831d12735836779db5eb0a9d1848a7d91a6715f9a11ab10bf5af"),
							Index:           0,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x84257dbc6d86831d12735836779db5eb0a9d1848a7d91a6715f9a11ab10bf5af"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network: network.Crossbell,
					Endpoint: config.Endpoint{
						URL: endpoint.MustGet(network.Crossbell),
					},
				},
			},
			want: &activityx.Activity{
				ID:      "0x84257dbc6d86831d12735836779db5eb0a9d1848a7d91a6715f9a11ab10bf5af",
				Network: network.Crossbell,
				Index:   0,
				From:    "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A",
				To:      "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:    typex.SocialProxy,
				Calldata: &activityx.Calldata{
					FunctionHash: "0x206657f2",
				},
				Platform: workerx.PlatformCrossbell.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("30915000000000")),
					Decimal: 18,
				},
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialProxy,
						Platform: "Crossbell",
						From:     "0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A",
						To:       "0x0F588318A494e4508A121a32B6670b5494Ca3357",
						Metadata: metadata.SocialProxy{
							ProxyAddress: common.HexToAddress("0x0f588318a494e4508a121a32b6670b5494ca3357"),
							Action:       metadata.ActionSocialProxyRemove,
							Profile: metadata.SocialProfile{
								Handle:    "brucexx.csb",
								ProfileID: "32010",
								Address:   common.HexToAddress("0x23c46e912b34C09c4bCC97F4eD7cDd762cee408A"),
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1685695671,
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
