package matters_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/matters"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/matters"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Matters(t *testing.T) {
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
			name: "Curation",
			arguments: arguments{
				task: &source.Task{
					Network: network.Optimism,
					ChainID: 10,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x7ba6c85ff9d6eba5cb5f0fb6974ff4efbdef4dc4f093a2a00f50a851c8a49bcf"),
						ParentHash:   common.HexToHash("0x347a738508aec78e101d6790b564201959b7841473dd6f4e556d13443161cbd2"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x4200000000000000000000000000000000000011"),
						Number:       lo.Must(new(big.Int).SetString("117061980", 0)),
						GasLimit:     30000000,
						GasUsed:      23316448,
						Timestamp:    1709722737,
						BaseFee:      lo.Must(new(big.Int).SetString("227967", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0"),
						From:      common.HexToAddress("0x869871ba7eC8BCB59530a04e1e28A42A208a697b"),
						Gas:       70081,
						GasPrice:  lo.Must(new(big.Int).SetString("1011502", 10)),
						Hash:      common.HexToHash("0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0"),
						Input:     hexutil.MustDecode("0xdbcdaf5b00000000000000000000000073d5f95b8f6fc4178aa934c44d7acc7dc2a6daf800000000000000000000000094b008aa00579c1307b0ef2c499ad98a8ce58e5800000000000000000000000000000000000000000000000000000000000f424000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000035697066733a2f2f516d58325245656f724b786d347170667479467a514850683243654a654143364368796e51537a62506a675866470000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x5edebbdae7B5C79a69AaCF7873796bb1Ec664DB8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("10", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x7ba6c85ff9d6eba5cb5f0fb6974ff4efbdef4dc4f093a2a00f50a851c8a49bcf"),
						BlockNumber:       lo.Must(new(big.Int).SetString("117061980", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 22764157,
						EffectiveGasPrice: hexutil.MustDecodeBig("0xf6f2e"),
						GasUsed:           69223,
						L1GasPrice:        lo.Must(new(big.Int).SetString("67629803535", 0)),
						L1GasUsed:         lo.Must(new(big.Int).SetString("4072", 0)),
						L1Fee:             lo.Must(new(big.Int).SetString("188365775036251", 0)),
						FeeScalar:         lo.Must(new(big.Float).SetString("0.684")),
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0x94b008aA00579c1307B0EF2c499aD98a8ce58e58"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x000000000000000000000000869871ba7ec8bcb59530a04e1e28a42a208a697b"),
								common.HexToHash("0x00000000000000000000000073d5f95b8f6fc4178aa934c44d7acc7dc2a6daf8"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000f4240"),
							BlockNumber:     lo.Must(new(big.Int).SetString("117061980", 0)),
							TransactionHash: common.HexToHash("0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0"),
							Index:           147,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x94b008aA00579c1307B0EF2c499aD98a8ce58e58"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x000000000000000000000000869871ba7ec8bcb59530a04e1e28a42a208a697b"),
								common.HexToHash("0x0000000000000000000000005edebbdae7b5c79a69aacf7873796bb1ec664db8"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000006acfc0"),
							BlockNumber:     lo.Must(new(big.Int).SetString("117061980", 0)),
							TransactionHash: common.HexToHash("0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0"),
							Index:           148,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x5edebbdae7B5C79a69AaCF7873796bb1Ec664DB8"),
							Topics: []common.Hash{
								common.HexToHash("0xc2e41b3d49bbccbac6ceb142bad6119608adf4f1ee1ca5cc6fc332e0ca2fc602"),
								common.HexToHash("0x000000000000000000000000869871ba7ec8bcb59530a04e1e28a42a208a697b"),
								common.HexToHash("0x00000000000000000000000073d5f95b8f6fc4178aa934c44d7acc7dc2a6daf8"),
								common.HexToHash("0x00000000000000000000000094b008aa00579c1307b0ef2c499ad98a8ce58e58"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000f42400000000000000000000000000000000000000000000000000000000000000035697066733a2f2f516d58325245656f724b786d347170667479467a514850683243654a654143364368796e51537a62506a675866470000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("117061980", 0)),
							TransactionHash: common.HexToHash("0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0"),
							Index:           149,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0"),
						TransactionIndex: 8,
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
				ID:      "0x3000bbbfe0fd1d7f13eb8ff9e324db0c9ae75466e6666881216377f8d6fce2a0",
				Network: network.Optimism,
				Index:   8,
				From:    "0x869871ba7eC8BCB59530a04e1e28A42A208a697b",
				To:      matters.AddressCuration.String(),
				Type:    typex.SocialReward,
				Calldata: &activityx.Calldata{
					FunctionHash: "0xdbcdaf5b",
				},
				Tag:      tag.Social,
				Platform: workerx.PlatformMatters.String(),
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("188435794239197")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*activityx.Action{
					{
						Type:     typex.SocialReward,
						Tag:      tag.Social,
						Platform: workerx.PlatformMatters.String(),
						From:     "0x869871ba7eC8BCB59530a04e1e28A42A208a697b",
						To:       "0x73d5f95B8f6FC4178aA934c44D7Acc7dC2a6DAf8",
						Metadata: &metadata.SocialPost{
							Reward: &metadata.Token{
								Address:  lo.ToPtr(common.HexToAddress("0x94b008aA00579c1307B0EF2c499aD98a8ce58e58").String()),
								Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1000000), 0)),
								Name:     "Tether USD",
								Symbol:   "USDT",
								Standard: metadata.StandardERC20,
								Decimals: 6,
							},
							Target: &metadata.SocialPost{
								Handle:  "guo",
								Title:   "ZuSocial: notes and reflections - 刘果 | Guo Liu (guo)",
								Summary: "In November Istanbul, alongside the vibrant communities of ZuConnect and Devconnect, we held a 2-week hacker house focusing on decentralized social.",
								Body:    "![](./06e20887-c9da-4efd-b02f-47367cc5277f.png)Sunset in Istanbul, by Eunhye\n\nZuSocial was more than just an event; it was a chance to connect with a diverse group of individuals, who share the belief of decentralization and participate in an emerging vision of a better social network.\n\nIn the areas where blockchain technology could fundamentally change, social network might just rival its impact on finance, if not surpass it. More and more people get their information from some kind of social network, while at the same time, the flaws in our centralized social media platforms are becoming more apparent.\n\nThese platforms often prioritize seeking attention over fostering genuine human connections. They centralize control over the flow of information, becoming tools of manipulation and even weapons during conflicts. Our thoughts, shared in these virtual public spaces, are gradually privatized without us even realizing it.\n\nDecentralization offers the potential to address these issues. It can potentially make censorship and propaganda more difficult, replacing the dominance of a few platforms with a landscape of fairer protocols and more efficient markets.\n\nHowever, we know very little about how to build these decentralized alternatives. The world of blockchain and related technologies looks vastly different from our current social media landscape, and might not be directly meaningful to an end user. We also lack a deep understanding of how to construct digital public spaces since they are a relatively new concept.\n\nTo navigate this uncharted territory, we must seek guidance from fields that have amassed more experience. At ZuSocial, we intentionally drew inspiration from the physical world and borrowed tools from disciplines like architecture and urban planning. It's a fascinating journey, especially during events like ZuConnect, ZuSocial, and DevConnect, where we can experiment and learn by observing our own interactions.\n\nBelow are some of my learnings and reflections during ZuSocial. The most profound conversations were often free-flowing and unstructured, but I've organized my thoughts into four distinct threads.\n\n- [ZuConnect decentralized social day: lessons from the past](https://liuguo.eth.limo/505333-zu-connect-decentralized-day-lessons-from-the-past/)\n\n- [What does blockchain bring to social protocol?](https://liuguo.eth.limo/505337-what-does-blockchain-bring-to-social-protocol/)\n\n- [The economics of Matters Town](https://liuguo.eth.limo/505387-the-economics-of-matters-town/)\n\n- [Groups and autonomous spaces: a mental model for better social networks](https://liuguo.eth.limo/505396-groups-and-autonomous-spaces-a-mental-model-for-better-social-networks/)",
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1709722737,
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

			t.Log(activity)
		})
	}
}
