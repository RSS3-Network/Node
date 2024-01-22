package crossbell_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/rss3-node/config"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/crossbell"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/endpoint"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
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
			name: "Crossbell Create Profile",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkCrossbell,
					ChainID: 3737,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x208cd7fa40c76bd071edc0db8f6d107b70e194cd27e69027e7c642523f15640d"),
						ParentHash:   common.HexToHash("0xe9de8de4fc50554480ee7018abf94a0876af9b4db6f50d0dd57ec50325635df8"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("5277407", 0)),
						GasLimit:     8000000,
						GasUsed:      336353,
						Timestamp:    1655636500,
						BaseFee:      nil,
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
						From:      common.HexToAddress("0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101"),
						Gas:       409471,
						GasPrice:  lo.Must(new(big.Int).SetString("1000000000", 10)),
						Hash:      common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
						Input:     hexutil.MustDecode("0xbd5f69cb00000000000000000000000000000000000000000000000000000000000000200000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd657810100000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000362656800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042697066733a2f2f6261666b726569646a346a64707778693661697a343375626f6b76356333746e716c66346e66706b6f737332727973666a7a726465746769656b6d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   nil,
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x208cd7fa40c76bd071edc0db8f6d107b70e194cd27e69027e7c642523f15640d"),
						BlockNumber:       lo.Must(new(big.Int).SetString("5277407", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 336353,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3b9aca00"),
						GasUsed:           336353,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
								common.HexToHash("0x0000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd6578101"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000df1"),
							},
							Data:            nil,
							BlockNumber:     lo.Must(new(big.Int).SetString("5277407", 0)),
							TransactionHash: common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
							Index:           0,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8"),
							Topics: []common.Hash{
								common.HexToHash("0xa5802a04162552328d75eaac538a033704a7c3beab65d0a83e52da1c8c9b7cdf"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000df1"),
								common.HexToHash("0x0000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd6578101"),
								common.HexToHash("0x0000000000000000000000003f54953e56a0fe272f3e1e61615bda1fd6578101"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000062af021400000000000000000000000000000000000000000000000000000000000000036265680000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("5277407", 0)),
							TransactionHash: common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
							Index:           1,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3"),
						TransactionIndex: 0,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0x70fa8c5b0fc9ce76fd27a1b99c1a3083bc5e4d0d72a7630cefc27e7badeb6ad3",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialProfile,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("336353000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformCrossbell.String(),
						From:     "0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101",
						To:       "0x3F54953E56A0Fe272F3E1E61615Bda1fd6578101",
						Metadata: metadata.SocialProfile{
							ProfileID: "3569",
							Action:    metadata.ActionSocialProfileCreate,
							Address:   common.HexToAddress("0x3f54953e56a0fe272f3e1e61615bda1fd6578101"),
							Handle:    "beh.csb",
							ImageURI:  "ipfs://bafkreidj4jdpwxi6aiz43ubokv5c3tnqlf4nfpkoss2rysfjzrdetgiekm",
						},
					},
				},
				Status:    true,
				Timestamp: 1655636500,
			},
			wantError: require.NoError,
		},
		{
			name: "Crossbell Set Profile URI",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0x24bc9af800a8d543e2a8cafd0015c5cb9166203e5114f36fd28dbc244a65a659",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x08d66b34054a174841e2361bd4746Ff9F4905cC2",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialProfile,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("51306000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformCrossbell.String(),
						From:     "0x08d66b34054a174841e2361bd4746Ff9F4905cC2",
						To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
						Metadata: metadata.SocialProfile{
							ProfileID: "30",
							Action:    metadata.ActionSocialProfileUpdate,
							Address:   common.HexToAddress("0x08d66b34054a174841e2361bd4746ff9f4905cc2"),
							Handle:    "song.csb",
							ImageURI:  "csb://asset:0x5452c7fb99d99fab3cc1875e9da9829cb50f7a13-1008@ethereum",
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0x8ba289a3cd8ea21a52834961fa5dde07b2f52d51bd571fc40b18af09f841a3ef",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialProfile,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("327967000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformCrossbell.String(),
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0x96a778fa21f708c24de5f995256155573815d96498de490155366f92aae512c7",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialProfile,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("56047000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformCrossbell.String(),
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0xc6268ec5e794165322787e56aa483f6e1eb61586c8762c45593d0ee65a8740f3",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x0fefeD77Bb715E96f1c35c1a4E0D349563d6f6c0",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialProfile,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("47651000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialProfile,
						Platform: filter.PlatformCrossbell.String(),
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0xf81f06f80b929e9f8c7bd9bbac6759a79c89d10fb3e7ce3f34d05b3f50253c41",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x0F588318A494e4508A121a32B6670b5494Ca3357",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialPost,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("236812000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0x4add740166bde1a1fc64045c3098a28f7335972db3bfb73455df3020d1365e04",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x0F588318A494e4508A121a32B6670b5494Ca3357",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialComment,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("175319000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialComment,
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0x2b24d08600be7ce57f664d31c1ea0ff0a6ded55d5be8266cfee669305297cc93",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialRevise,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("53229000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialRevise,
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0xd0b0341118411ec564737458824544d00b75074989c775bac3db47a47dfd0b2a",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0x8a6dDC78E3aA24f4F31980623f489a274b305762",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialDelete,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("49945000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialDelete,
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
					Network: filter.NetworkCrossbell,
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
					Network:  filter.NetworkCrossbell,
					Endpoint: endpoint.MustGet(filter.NetworkCrossbell),
				},
			},
			want: &schema.Feed{
				ID:       "0xf4e9d4397b364c1ae0202e836fbae505b77382d25a238f0d98f73c72c81fa693",
				Network:  filter.NetworkCrossbell,
				Index:    0,
				From:     "0xc560eb6fd0c2eb80Df50E5e06715295AE1205049",
				To:       "0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8",
				Type:     filter.TypeSocialMint,
				Platform: lo.ToPtr(filter.PlatformCrossbell),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("416564000000000")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialMint,
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
