package lens_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/lens"
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
		config *engine.Config
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Post By polebug.lens",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x4453f44d391358825a6c00922ee0126c033aafe85258ec913cff61f26da48c92"),
						ParentHash:   common.HexToHash("0xb8c42dd921c150bb2283167d4e980062b019985e332b0695ca0d0e479a9d0adb"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xb9EDE6f94D192073D8eaF85f8db677133d483249"),
						Number:       lo.Must(new(big.Int).SetString("39312003", 0)),
						GasLimit:     30000000,
						GasUsed:      7653945,
						Timestamp:    1676445288,
						BaseFee:      lo.Must(new(big.Int).SetString("185783717525", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x1d4dc90c35d75df96b04e53f3949bad991f432a78dfa2196df39ad60d657e525"),
						From:      common.HexToAddress("0xD1FecCF6881970105dfb2b654054174007f0e07E"),
						Gas:       1500000,
						GasPrice:  lo.Must(new(big.Int).SetString("268151093434", 10)),
						Hash:      common.HexToHash("0x1d4dc90c35d75df96b04e53f3949bad991f432a78dfa2196df39ad60d657e525"),
						Input:     hexutil.MustDecode("0x963ff14100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000015fc300000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000a31ff85e840ed117e172bc9ad89e55128a999205000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000003f68747470733a2f2f617277656176652e6e65742f466534302d2d4f726646746f57464e775165336361524f7735577255334e416646796b52795343583339590000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xDb46d1Dc155634FbC732f92E853b10B288AD5a1d")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x4453f44d391358825a6c00922ee0126c033aafe85258ec913cff61f26da48c92"),
						BlockNumber:       lo.Must(new(big.Int).SetString("39312003", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 2409090,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x3e6f0cf8ba"),
						GasUsed:           152136,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xDb46d1Dc155634FbC732f92E853b10B288AD5a1d"),
							Topics: []common.Hash{
								common.HexToHash("0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000015fc3"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000004"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000a31ff85e840ed117e172bc9ad89e55128a9992050000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000063ec8668000000000000000000000000000000000000000000000000000000000000003f68747470733a2f2f617277656176652e6e65742f466534302d2d4f726646746f57464e775165336361524f7735577255334e416646796b52795343583339590000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("39312003", 0)),
							TransactionHash: common.HexToHash("0x1d4dc90c35d75df96b04e53f3949bad991f432a78dfa2196df39ad60d657e525"),
							Index:           85,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000d1feccf6881970105dfb2b654054174007f0e07e"),
								common.HexToHash("0x000000000000000000000000b9ede6f94d192073d8eaf85f8db677133d483249"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000002c84ead4b7b4680000000000000000000000000000000000000000000000b6b628a132c217cfc70000000000000000000000000000000000000000000005bbf31840652c144eb30000000000000000000000000000000000000000000000b6b5fc1c47ed601b5f0000000000000000000000000000000000000000000005bbf344c55000cc031b"),
							BlockNumber:     lo.Must(new(big.Int).SetString("39312003", 0)),
							TransactionHash: common.HexToHash("0x1d4dc90c35d75df96b04e53f3949bad991f432a78dfa2196df39ad60d657e525"),
							Index:           86,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x1d4dc90c35d75df96b04e53f3949bad991f432a78dfa2196df39ad60d657e525"),
						TransactionIndex: 17,
					},
				},
				config: &engine.Config{
					Network:  filter.NetworkPolygon,
					Endpoint: endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "0x1d4dc90c35d75df96b04e53f3949bad991f432a78dfa2196df39ad60d657e525",
				Network:  filter.NetworkPolygon,
				Index:    17,
				From:     "0xD1FecCF6881970105dfb2b654054174007f0e07E",
				To:       "0xDb46d1Dc155634FbC732f92E853b10B288AD5a1d",
				Type:     filter.TypeSocialPost,
				Platform: lo.ToPtr(filter.PlatformLens),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("40795434750675024")),
					Decimal: 18,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
						Platform: "Lenster",
						From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						To:       "0xDb46d1Dc155634FbC732f92E853b10B288AD5a1d",
						Metadata: metadata.SocialPost{
							Body: "big head",
							Media: []metadata.Media{
								{
									MimeType: "image/jpeg",
									Address:  "ipfs://bafybeihjz7wh2uexzuwlnumtcr7ptjyglicg4ura43z4cnn3pganwiqifa",
								},
							},
							Handle:        "polebug.lens",
							ProfileID:     "0x015fc3",
							PublicationID: "0x04",
							ContentURI:    "https://arweave.net/Fe40--OrfFtoWFNwQe3caROw5WrU3NAfFykRySCX39Y",
							Tags:          nil,
							Timestamp:     1676445288,
						},
					},
				},
				Status:    true,
				Timestamp: 1676445288,
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
			require.Equal(t, testcase.want, feed)

			t.Log(feed)
		})
	}
}
