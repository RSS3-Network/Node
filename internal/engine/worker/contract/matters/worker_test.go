package matters_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/matters"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/matters"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
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
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Curation",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xd287e85c28567ff5e0ecde319f9e9d3cd39c399c3cd97015af10adfc0b971427"),
						ParentHash:   common.HexToHash("0x623c004b7150f523123181934c5d7d71afa229f7f98bad19448348eab0fbce24"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xbDBd4347b082D9d6BdF2Da4555a37Ce52a2e2120"),
						Number:       lo.Must(new(big.Int).SetString("46169613", 0)),
						GasLimit:     30000000,
						GasUsed:      8071254,
						Timestamp:    1691743370,
						BaseFee:      lo.Must(new(big.Int).SetString("46050144999", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
						From:      common.HexToAddress("0x18Fb694EbAE03a78f038F54362592Dd89c0e300C"),
						Gas:       118339,
						GasPrice:  lo.Must(new(big.Int).SetString("76050144999", 10)),
						Hash:      common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
						Input:     hexutil.MustDecode("0xdbcdaf5b000000000000000000000000099d854cd6f996cd1cb1ffc1df0e197a774426e2000000000000000000000000c2132d05d31c914a87c6611c10748aeb04b58e8f00000000000000000000000000000000000000000000000000000000000a875000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000035697066733a2f2f516d52546e4a4e316b77746b766b576d326669635455433771486244714e316361483763574b79397158595336310000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0x5edebbdae7B5C79a69AaCF7873796bb1Ec664DB8")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      2,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xd287e85c28567ff5e0ecde319f9e9d3cd39c399c3cd97015af10adfc0b971427"),
						BlockNumber:       lo.Must(new(big.Int).SetString("46169613", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 7216947,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x11b4f11ee7"),
						GasUsed:           74710,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
							Topics: []common.Hash{
								common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
								common.HexToHash("0x00000000000000000000000018fb694ebae03a78f038f54362592dd89c0e300c"),
								common.HexToHash("0x000000000000000000000000099d854cd6f996cd1cb1ffc1df0e197a774426e2"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000a8750"),
							BlockNumber:     lo.Must(new(big.Int).SetString("46169613", 0)),
							TransactionHash: common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
							Index:           203,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
							Topics: []common.Hash{
								common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
								common.HexToHash("0x00000000000000000000000018fb694ebae03a78f038f54362592dd89c0e300c"),
								common.HexToHash("0x0000000000000000000000005edebbdae7b5c79a69aacf7873796bb1ec664db8"),
							},
							Data:            hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1ee296"),
							BlockNumber:     lo.Must(new(big.Int).SetString("46169613", 0)),
							TransactionHash: common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
							Index:           204,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x5edebbdae7B5C79a69AaCF7873796bb1Ec664DB8"),
							Topics: []common.Hash{
								common.HexToHash("0xc2e41b3d49bbccbac6ceb142bad6119608adf4f1ee1ca5cc6fc332e0ca2fc602"),
								common.HexToHash("0x00000000000000000000000018fb694ebae03a78f038f54362592dd89c0e300c"),
								common.HexToHash("0x000000000000000000000000099d854cd6f996cd1cb1ffc1df0e197a774426e2"),
								common.HexToHash("0x000000000000000000000000c2132d05d31c914a87c6611c10748aeb04b58e8f"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000a87500000000000000000000000000000000000000000000000000000000000000035697066733a2f2f516d52546e4a4e316b77746b766b576d326669635455433771486244714e316361483763574b79397158595336310000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("46169613", 0)),
							TransactionHash: common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
							Index:           205,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x00000000000000000000000018fb694ebae03a78f038f54362592dd89c0e300c"),
								common.HexToHash("0x000000000000000000000000bdbd4347b082d9d6bdf2da4555a37ce52a2e2120"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000007f6735255c8000000000000000000000000000000000000000000000000009ac8d1aa4def2b4200000000000000000000000000000000000000000000306673636b994b170d050000000000000000000000000000000000000000000000009ac0db36fb996342000000000000000000000000000000000000000000003066736b620c9d6cd505"),
							BlockNumber:     lo.Must(new(big.Int).SetString("46169613", 0)),
							TransactionHash: common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
							Index:           206,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34"),
						TransactionIndex: 41,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkPolygon,
					Endpoint: endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:      "0x406420c77151dd2779f647967aaee7b2ae25248bbb51b53de4d8912c904dbd34",
				Network: filter.NetworkPolygon,
				Index:   41,
				From:    "0x18Fb694EbAE03a78f038F54362592Dd89c0e300C",
				To:      matters.AddressCuration.String(),
				Type:    filter.TypeSocialReward,
				Calldata: &schema.Calldata{
					FunctionHash: "0xdbcdaf5b",
				},
				Tag:      filter.TagSocial,
				Platform: lo.ToPtr(filter.PlatformMatters),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("5681706332875290")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialReward,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformMatters.String(),
						From:     "0x18Fb694EbAE03a78f038F54362592Dd89c0e300C",
						To:       "0x099d854cD6F996cD1CB1ffc1Df0E197A774426e2",
						Metadata: &metadata.SocialPost{
							Reward: &metadata.Token{
								Address:  lo.ToPtr(common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F").String()),
								Value:    lo.ToPtr(decimal.NewFromBigInt(big.NewInt(690000), 0)),
								Name:     "(PoS) Tether USD",
								Symbol:   "USDT",
								Standard: metadata.StandardERC20,
								Decimals: 6,
							},
							Target: &metadata.SocialPost{
								Handle:  "chorsee",
								Title:   "與他瑪同哭 —— 終止暴力對待婦女音樂會的共振 - 楚思 (chorsee)",
								Summary: "近日在台灣揭發的連串Me too事件，還有印度女摔跤手奮力就性侵事件示威。在臉書上見到非常多關於台灣事件的評論和分析，獲益良多，也望塵莫及，自己未有新的觀點和見解可分享，謹以一篇舊文章為發聲的女性聊表聲援支持，也希望當中分享的歌曲能帶來一些安慰。（抱）",
								Body:    "【註：2018年寫的舊文章，於香港基督教媒體《時代論壇》刊登，今次刊登時作了少許修改。】\n\n性侵犯受害人在網上勇敢自白，隨之而來一群判官，嘲諷攻擊。我最看不過眼的不是當中分析，是目中無人的態度，「佢擺明就係攞光環」（他一定是為了搏取光環）這些話明擺著是討厭，但故作持平大剌剌地指點如何是更好的做法，都可能是一種冷漠的傷害。\n\n參加過就同性戀作信仰反省的講座，立場兩極的講員搬出不同理論，腔調平穩，辯論時彬彬有禮，但聽來很遙遠。我提問：「如果有個同性戀者就在這裡，你會跟他說什麼？」講員沒有直接回應，後補的一番分析依舊稱呼「同性戀者」、「他們」。\n\n當受害者在我們跟前，我們懂得怎樣直接向「您」說話嗎？\n\n有天收到電話，朋友跟我凌亂地說了些酒吧、陌生人、醒來等字眼，繞過關鍵而太痛苦的一詞，問該怎辦。她還在強裝鎮定，我更覺一片空白，失語。\n\n* * *\n\n早前參與「他瑪的呼聲——終止暴力對待婦女音樂會」，由基督教協進會的性別公義事工籌辦。音樂會名字的「他瑪」所指的是聖經的撒母耳記下所記，大衛的女兒，她被同父異母兄長性侵犯。\n\n參與音樂會時沒特別期待，怎料好像將內心赤裸裸地送進了SPA，是一場意想不到的觸動震撼，無法旁觀， 只能認真地感受內心一切起伏。\n\n第一首歌就是反覆的一聲聲「容許我淒泣」，開闢一個讓情緒自由流動的空間，托住了讀經的姊妹以及在座每人，我們看到姊妹讀畢他瑪故事的淚水，我們也絕對容許淒泣。之後，我們聽到更多其他的「他瑪」的故事，甚至有姊妹勇敢講親身遭遇，不是罐頭式「我改變了」的樂觀，但也不是糾纏於痛苦的細節，是切實的掙扎，也是向上帝切切的祈求。\n\n找不到中文歌詞版，推介我喜歡的Celtic Woman 版，歌詞就是《容許我淒泣》的意思\n\n向受害者說話不容易，何況跟他們同哭——不只是為他們而哭，而是不得不發現自己內心同樣脆弱，胸腔同樣流轉一股沸騰酸苦。不只是聽眾，籌備音樂會的人，也進入這段旅程，內心驟然要敞開，舒坦前必先疼痛，像拉開久未伸展的筋骨。\n\n敞開的赤裸是如此危險，只能求上主以愛填補，這時同讀「婦女信經」，「我信聖靈，上主的陰性之靈，她就像母雞，創造孕育我們，以她的翅膀覆庇我們。」還齊誦「醫治連禱」：「求你的醫治大能，觸摸一切混亂、自我懷疑的身心，並光照他們。」份外體會到，我們真的需要向上主呼求。\n\n最終尋求的希望，是《For everyone born》，作詞的Shirley Erena Murray將「世界人權宣言」放置於福音的處境，”God will delight when we are creators of justice and joy, compassion and peace.” 關懷不只限於某性別，是由此出發，去為祈願一個更公義平等的世界而歌。\n\n性別公義事工有辦工作坊作教育，也有安排專業輔導和牧師等人設立性侵受害人支援小組，還辦這音樂會，滋養安慰、省察、呼求，塑造著各人的內心，難能可貴。議題的事理固然要釐清，但其情感面向，更不能忽視的是活生生的人的故事、呼求、脆弱，願我們能同哀哭，才有可能同喜樂。\n\n音樂會的精華片段\n\n願公義得以伸張。願每一個受害人得到更多安慰、滋潤，大大力給你們一個擁抱，辛苦了。\n\n* * *\n\nP.S. 我將當天音樂會的歌曲串成了一個 [歌單](https://www.youtube.com/playlist?list=PL25ePuQiYMITe9DmoKdSKfx8ryToHHDCN) 作分享。脫離了整個音樂會的鋪排後，未必能傳遞到歌曲在音樂會中的含義，可是或能分享到一些安慰和激勵，特別也推介《Rise Up》和《Nobody knows the trouble I've seen》這兩首歌。",
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1691743370,
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
			//t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))
			require.Equal(t, testcase.want, feed)

			t.Log(feed)
		})
	}
}
