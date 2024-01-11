package iqwiki_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/iqwiki"
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
			name: "Parse IQWiki feed",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0xddc555e3da6de36d21f1266af0591b8434d509209273430bbc643870df5c8554"),
						ParentHash:   common.HexToHash("0x8c515d803c87a46ab01ec0c0bba323f005d5743fe3f4669443f4d69abedb4e8b"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0xa8B52F02108AA5F4B675bDcC973760022D7C6020"),
						Number:       lo.Must(new(big.Int).SetString("52070001", 0)),
						GasLimit:     30000000,
						GasUsed:      16300598,
						Timestamp:    1704686506,
						BaseFee:      lo.Must(new(big.Int).SetString("61227334651", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e"),
						From:      common.HexToAddress("0x191a41c307373211D08613B68df4031977589069"),
						Gas:       50000,
						GasPrice:  lo.Must(new(big.Int).SetString("199044496620", 10)),
						Hash:      common.HexToHash("0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e"),
						Input:     hexutil.MustDecode("0xed53ddb900000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000caf1cd3c9c76382e38fd813aecf103325ebd0dbe00000000000000000000000000000000000000000000000000000000659cc51a000000000000000000000000000000000000000000000000000000000000001bc2c3956065ff0e3b606f51cd599589a4a9d8e5ecbbc33cb3ac86d51b3ea105970e6944a86d3bc4ab126f29e9fcd18baeec8aae07f774cbad32655419a1dc298d000000000000000000000000000000000000000000000000000000000000002e516d6267637a72376f53636f71725964706950574a336676597173585255557a42675937426e7a41503250393846000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0xddc555e3da6de36d21f1266af0591b8434d509209273430bbc643870df5c8554"),
						BlockNumber:       lo.Must(new(big.Int).SetString("52070001", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 3808332,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x2e57f9fcec"),
						GasUsed:           39151,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D"),
							Topics: []common.Hash{
								common.HexToHash("0x0fe877471cb763db81561ce83d01be57b6699361a3febbc7a0cdfb18df66246b"),
								common.HexToHash("0x000000000000000000000000caf1cd3c9c76382e38fd813aecf103325ebd0dbe"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002e516d6267637a72376f53636f71725964706950574a336676597173585255557a42675937426e7a41503250393846000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("52070001", 0)),
							TransactionHash: common.HexToHash("0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e"),
							Index:           95,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000191a41c307373211d08613b68df4031977589069"),
								common.HexToHash("0x000000000000000000000000a8b52f02108aa5f4b675bdcc973760022d7c6020"),
							},
							Data:            hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000132b579454f8ff00000000000000000000000000000000000000000000000119106de9cb56145b000000000000000000000000000000000000000000000f63ce919d96bc73d91200000000000000000000000000000000000000000000000118fd429237011b5c000000000000000000000000000000000000000000000f63cea4c8ee50c8d211"),
							BlockNumber:     lo.Must(new(big.Int).SetString("52070001", 0)),
							TransactionHash: common.HexToHash("0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e"),
							Index:           96,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e"),
						TransactionIndex: 16,
					},
				},
				config: &engine.Config{
					Network:  filter.NetworkPolygon,
					Endpoint: endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e",
				Network:  filter.NetworkPolygon,
				Index:    0x10,
				From:     "0x191a41c307373211D08613B68df4031977589069",
				To:       "0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D",
				Type:     filter.TypeSocialPost,
				Tag:      filter.TagSocial,
				Platform: lo.ToPtr(filter.PlatformIQWiki),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("1799863423694410")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
						Platform: filter.PlatformIQWiki.String(),
						From:     "0xcAF1CD3c9C76382E38fd813aecf103325ebD0dbE",
						To:       "0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D",
						Metadata: metadata.SocialPost{
							Handle:  "",
							Title:   "Aleksander Larsen",
							Summary: "Aleksander Larsen is a game developer and entrepreneur, known for co-founding Axie Infinity, Ronin Network, Sky Mavis, and the gaming community Dota.no.",
							Body:    "**Aleksander Larsen** is the co-founder, chairman, and COO of Axie Infinity, Ronin Network, and Sky Mavis. He also is an occasional investor and venture partner at South Korean VC firm [#Hashed](https://iq.wiki/wiki/hashed). Having always been interested in video game and game development, he helped grow gaming communities before moving on to co-found his own game studios.  \n  \n# Career  \n  \nHe began his entrepreneurial career early on in 2004, serving as the administrator and co-founder of Dota.no, a Nordic E-Sport community. This continued for over 13 years until December 2017. In 2018, he was responsible for Community Management and Business Development at Parsec Frontiers, a multiplayer online game where all the assets are secured on blockchain. He helped grow the community on Discord, Telegram, and Twitter.  \n  \nLarsen soon took that experience and co-founded [Sky Mavis](https://iq.wiki/wiki/sky-mavis) in April 2018, a game studio, with Larsen at the helm. One of the flagship products is [Axie Infinity](https://iq.wiki/wiki/axie-infinity), a complex digital nation involving blockchain technology and incentive design. Larsen handles the business aspect of things, including fundraising, corporate development, and partnerships.  \n  \n# Sky Mavis  \n  \nSky Mavis focuses on creating virtual worlds with player-owned economies. It is a marketplace for trustless trading of unique assets. They build infrastructure that enables the products to reach millions of players worldwide. Larsen leads business, legal, finance, venture, and operations here. It was mentioned[\\[1\\]](#cite-id-o1zrellbc4) as a Fortune Top 40 Crypto Company in 2023.  \n  \nAs the COO, Larsen helped raise $150 million in 2022 led by [Binance](https://iq.wiki/wiki/binance) to reimburse victims of a $625 million exploit of Axie Infinity's Ronin network. Ronin Network was a critical bridge chain that powers Axie Infinity. The attack resulted in a loss of 173,600 Ethereum and 25.5M [USDC](https://iq.wiki/wiki/usdc), equating to over $600 million. The attack happened on March 23, 2022. Larsen addressed the matter on Twitter, claiming to have worked with the Sky Mavis board and key cybersecurity personnel to get an overview of the situation.  \n  \n> Been working with the Sky Mavis board and key cybersecurity personnel to get a complete overview of the situation Our internal network is currently going through a deep forensics review to ensure there is no lingering threat. This was a social engineering attack combined with a human error from December 2021. @SkyMavisHQ tech is solid and we will be adding several new validators to @Ronin\\_Network shortly to further decentralize the network. We are committed to ensuring that all of the drained funds are recovered or reimbursed, and we are continuing conversations with our stakeholders to determine the best course of action.[\\[3\\]](#cite-id-2e5a7rg5s64)  \n  \n# Axie Infinity  \n  \nAxie Infinity is Sky Mavis' flagship product and is the #1 game on [Ethereum](https://iq.wiki/wiki/ethereum) by daily, weekly, and monthly active players. A Pokémon-inspired game, players can battle, raise and collect fantasy creatures called Axies. It started when Larsen and the other founders met playing [CryptoKitties](https://iq.wiki/wiki/cryptokitties), and they shared the common goal of making a better game.[\\[9\\]](#cite-id-mlqkxog80m) They began to develop Axie, which allowed gamers to farm the time spent on the game, turning it into [Ether](https://iq.wiki/wiki/ethereum), and then money. This was a new concept to games and crypto at that point in time.  \n  \nThe game is growing at an incredible speed, reporting $100,000 of revenue in January 2022, over $190 million in July 2022, and $360 million in August.[\\[2\\]](#cite-id-tgxcjy2fw2)  \n  \nLarsen firmly believes that the game's play-to-earn characteristics help it stand out from other blockchain-based titles and boost its sustainability. He believes that the internet and mobile games will move towards being decentralized through Web3.  \n  \n> \"And right now the internet is the metaverse to an extent. The only difference here is that it's not truly decentralized and you didn't have ownership on the internet before now you have ownership through the blockchain. So that means that the internet is now becoming more metaverse like, because people can own their own assets. They can interact with the blockchain through their web explorer for the front ends to where the metaverse are right now.\"[\\[7\\]](#cite-id-ezl99wsbdlb)  \n  \n# Blockchain Game Alliance (BGA)  \n  \nWith his expertise, Larsen soon rose to become the Secretary of the Board of Directors for the Blockchain Game Alliance (BGA) in 2020. The BGA is an organization committed to the adoption of blockchain technologies in gaming, by promoting the creation of common standards, facilitating networking and knowledge sharing, as well as encouraging cooperation and the implementation of best practices.  \n  \nLarsen's job as the secretary includes creating a long-term strategy for the BGA in cooperation with the rest of the board, making the value proposition clear, onboarding additional companies to the BGA, administrative management of projects, and last but not least, organizing meetups and connecting the members to potential investors.  \n  \n# Education & Personal Life  \n  \nLarsen was born and raised in Oslo, Norway. He graduated from BI Norwegian Business School in 2011 with a degree in economics and administration[\\[5\\]](#cite-id-blds5ln5pzb).  \n  \nHe speaks Norwegian fluently and English professionally[\\[6\\]](#cite-id-rjbzszrc23b).  \n  \nHe describes himself \"more of a gaming person than a crypto person,\" originally wanting to work in the gaming sector.[\\[8\\]](#cite-id-0y0xetm49yzp)  \n  \n# Trivia  \n  \nLarsen spent his free time playing games and building gamin communities. He has represented Norway in Warcraft3 and Dota, even once ranking as one of the Top 200 players in Europe in Dota 2.[\\[4\\]](#cite-id-lcds4vg073c)  \n  \n<br>  \n<br>  \n<br>  \n<br>  \n",
							Media: []metadata.Media{{
								Address:  "ipfs://QmXG6bi5vb4wNtTLaY2Sgo5WAVZ4kpNp7BKpgyzc6H3CsP",
								MimeType: "image/png",
							}},
							ProfileID:     "",
							PublicationID: "aleksander-larsen",
							ContentURI:    "https://iq.wiki/wiki/aleksander-larsen",
							Tags: []string{
								"Founders",
							},
							AuthorURL: "https://iq.wiki/account/0xcAF1CD3c9C76382E38fd813aecf103325ebD0dbE",
							Reward:    nil,
							Timestamp: 0,
							Target:    nil,
							TargetURL: "",
						},
					},
				},
				Status:    true,
				Timestamp: 1704686506,
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := worker.NewWorker()
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
