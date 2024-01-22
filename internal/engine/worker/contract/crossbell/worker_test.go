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
