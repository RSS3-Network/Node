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
