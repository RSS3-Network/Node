package iqwiki_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/iqwiki"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
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
			name: "Create a Wiki",
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
				config: &config.Module{
					Network:  filter.NetworkPolygon,
					Endpoint: endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:      "0x43dc470bbec2f3c585ac8f7a8340870b774a5be52ab9cf0836a8d534761be85e",
				Network: filter.NetworkPolygon,
				Index:   0x10,
				From:    "0x191a41c307373211D08613B68df4031977589069",
				To:      "0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D",
				Type:    filter.TypeSocialPost,
				Calldata: &schema.Calldata{
					FunctionHash: "0xed53ddb9",
				},
				Tag:      filter.TagSocial,
				Platform: lo.ToPtr(filter.PlatformIQWiki),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("7792791087169620")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*schema.Action{
					{
						Tag:      filter.TagSocial,
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
		{
			name: "Revise a Wiki",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkPolygon,
					ChainID: 137,
					Header: &ethereum.Header{
						Hash:         common.HexToHash("0x6450d21272936b5b825b2177c21acc7b7432652019082ad1fe3a1657b986429b"),
						ParentHash:   common.HexToHash("0x687a50f591df244770ed9a6d264be8971bffe78a2b09a1ed24af6d718077d1b7"),
						UncleHash:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Number:       lo.Must(new(big.Int).SetString("46942958", 0)),
						GasLimit:     29505428,
						GasUsed:      15544333,
						Timestamp:    1693414028,
						BaseFee:      lo.Must(new(big.Int).SetString("81776886634", 0)),
						Transactions: nil,
					},
					Transaction: &ethereum.Transaction{
						BlockHash: common.HexToHash("0x395a0ea73962d7f6e22cecc7d74c8f489a6707cc65f7cebdb39355bf01e8694a"),
						From:      common.HexToAddress("0x191a41c307373211D08613B68df4031977589069"),
						Gas:       50000,
						GasPrice:  lo.Must(new(big.Int).SetString("307123462857", 10)),
						Hash:      common.HexToHash("0x395a0ea73962d7f6e22cecc7d74c8f489a6707cc65f7cebdb39355bf01e8694a"),
						Input:     hexutil.MustDecode("0xed53ddb900000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000212cb3f4ae6611054637f9f78f18fb628ad258bb0000000000000000000000000000000000000000000000000000000064f0c3ff000000000000000000000000000000000000000000000000000000000000001c2e12df626fb630f095de4639a27ddb18216b5cdee6d45dfa2588e7ff3ad1e365385c6e61213f5307bb55b7a838f251a6e357b681b9652d465753d97986bbbcc7000000000000000000000000000000000000000000000000000000000000002e516d5444314738587478767063717672596d794662544d4146384d44486f67505242326d536271583548356d4432000000000000000000000000000000000000"),
						To:        lo.ToPtr(common.HexToAddress("0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D")),
						Value:     lo.Must(new(big.Int).SetString("0", 0)),
						Type:      0,
						ChainID:   lo.Must(new(big.Int).SetString("137", 0)),
					},
					Receipt: &ethereum.Receipt{
						BlockHash:         common.HexToHash("0x6450d21272936b5b825b2177c21acc7b7432652019082ad1fe3a1657b986429b"),
						BlockNumber:       lo.Must(new(big.Int).SetString("46942958", 0)),
						ContractAddress:   nil,
						CumulativeGasUsed: 297429,
						EffectiveGasPrice: hexutil.MustDecodeBig("0x4781fc22c9"),
						GasUsed:           39151,
						Logs: []*ethereum.Log{{
							Address: common.HexToAddress("0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D"),
							Topics: []common.Hash{
								common.HexToHash("0x0fe877471cb763db81561ce83d01be57b6699361a3febbc7a0cdfb18df66246b"),
								common.HexToHash("0x000000000000000000000000212cb3f4ae6611054637f9f78f18fb628ad258bb"),
							},
							Data:            hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002e516d5444314738587478767063717672596d794662544d4146384d44486f67505242326d536271583548356d4432000000000000000000000000000000000000"),
							BlockNumber:     lo.Must(new(big.Int).SetString("46942958", 0)),
							TransactionHash: common.HexToHash("0x395a0ea73962d7f6e22cecc7d74c8f489a6707cc65f7cebdb39355bf01e8694a"),
							Index:           8,
							Removed:         false,
						}, {
							Address: common.HexToAddress("0x0000000000000000000000000000000000001010"),
							Topics: []common.Hash{
								common.HexToHash("0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63"),
								common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000001010"),
								common.HexToHash("0x000000000000000000000000191a41c307373211d08613b68df4031977589069"),
								common.HexToHash("0x00000000000000000000000067b94473d81d0cd00849d563c94d0432ac988b49"),
							},
							Data:            hexutil.MustDecode("0x000000000000000000000000000000000000000000000000001f580e8d8df1b10000000000000000000000000000000000000000000000011be05885fa596b00000000000000000000000000000000000000000000000481d735ae55732e6e3e0000000000000000000000000000000000000000000000011bc100776ccb794f000000000000000000000000000000000000000000000481d755066400bc5fef"),
							BlockNumber:     lo.Must(new(big.Int).SetString("46942958", 0)),
							TransactionHash: common.HexToHash("0x395a0ea73962d7f6e22cecc7d74c8f489a6707cc65f7cebdb39355bf01e8694a"),
							Index:           9,
							Removed:         false,
						}},
						Status:           1,
						TransactionHash:  common.HexToHash("0x395a0ea73962d7f6e22cecc7d74c8f489a6707cc65f7cebdb39355bf01e8694a"),
						TransactionIndex: 3,
					},
				},
				config: &config.Module{
					Network:  filter.NetworkPolygon,
					Endpoint: endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:      "0x395a0ea73962d7f6e22cecc7d74c8f489a6707cc65f7cebdb39355bf01e8694a",
				Network: filter.NetworkPolygon,
				Index:   0x3,
				From:    "0x191a41c307373211D08613B68df4031977589069",
				To:      "0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D",
				Type:    filter.TypeSocialRevise,
				Calldata: &schema.Calldata{
					FunctionHash: "0xed53ddb9",
				},
				Tag:      filter.TagSocial,
				Platform: lo.ToPtr(filter.PlatformIQWiki),
				Fee: &schema.Fee{
					Amount:  lo.Must(decimal.NewFromString("12024190694314407")),
					Decimal: 18,
				},
				TotalActions: 1,
				Actions: []*schema.Action{
					{
						Tag:      filter.TagSocial,
						Type:     filter.TypeSocialRevise,
						Platform: filter.PlatformIQWiki.String(),
						From:     "0x212Cb3F4aE6611054637f9f78F18fB628AD258bb",
						To:       "0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D",
						Metadata: metadata.SocialPost{
							Handle:  "zainab",
							Title:   "Dex Finance",
							Summary: "Dex Finance is a DAO founded in 2020, providing financial products with market yields such as Dex Money Market, dexIRA & dexETF.",
							Body:    "**Dex Finance,** founded in 2020, is a [decentralized autonomous organization](https://iq.wiki/wiki/dao) (DAO) offering a range of financial products with market yields. The core products consist of the Dex Money Market, dexIRA & dexETF. [\\[1\\]](#cite-id-1gtk87hhdur)[\\[2\\]](#cite-id-4h80aojny4f)  \n$$widget0 [YOUTUBE@VID](BL3kH_mRm4k)$$  \n  \n# Team  \n  \nDex Finance was co-founded in July 2020 by Jake[\\[9\\]](#cite-id-4zriaxlx9pk), the Chief Executive Officer (CEO), and Matt[\\[10\\]](#cite-id-3zdig6t82gh) the Chief Technical Officer (CTO) while Mark Allen[\\[11\\]](#cite-id-by610czmwkj) serves as the Chief Operations Officer (COO) of the company. The team consists of software engineers, marketing professionals, traders, and legal personnel. The core members own and operate a full-stack media company that provides services including web development, media production, and digital marketing. [\\[1\\]](#cite-id-1gtk87hhdur)  \n  \n# Products  \n  \n## Dex Money Market  \n  \nDex Money Market is an algorithmic reserve protocol with multi-step auto-compounded strategies, Exchange-Traded Fund (ETF) rewards, and bonding mechanisms used to allow investors to purchase Money Market tokens at a discount by buying with a preselected list of currencies, including dexIRA & dexETF. [\\[3\\]](#cite-id-0chyy2c4bebf)  \n  \nThe protocol offers a single regulatory mechanism and protocol-owned liquidity. The Money Market consists of two corresponding tokens – USDEX, and dexSHARES. [\\[3\\]](#cite-id-0chyy2c4bebf)  \n  \n### USDEX  \n  \nUSDEX (part of a multi-token token system) is an algorithmic [stablecoin](https://iq.wiki/wiki/stablecoin) with a built-in stability mechanism that aims to maintain its price to 1 USDC token and is used as a medium of exchange. [\\[3\\]](#cite-id-0chyy2c4bebf)  \n  \n### dexSHARES  \n  \ndexSHARE is a measure of value in the Money Market Protocol. dexSHARE holders have voting rights (governance) on proposals to improve the protocol and future use cases within the Money Market protocol. [\\[3\\]](#cite-id-0chyy2c4bebf)  \n  \n## dexIRA  \n  \ndexIRA, launched in July 2021, is an individual retirement account designed for [decentralized exchanges](https://iq.wiki/wiki/decentralized-exchange). It consists of an automated [BNB](https://iq.wiki/wiki/binance-coin) redistribution, multi-currency dividend rewards, buyback protocols, and tokenomics that promote long-term holding and wealth generation. [\\[4\\]](#cite-id-atsywgse3xr)  \n  \ndexIRA carries a 30% sales tax for all “non-qualified” distributions. The tax is split between rewards pool, liquidity pool, and proprietary buyback protocol. Investors have the option to have fees waived for “qualified” distributions—and thus tax-free and penalty-free—if they haven’t sold any dexIRA tokens within one year. [\\[4\\]](#cite-id-atsywgse3xr)  \n  \n## dexETF  \n  \ndexETF allows users to build a portfolio with a non-custodial [crypto](https://iq.wiki/wiki/cryptocurrency) management protocol. dexETF is a crypto investment method whereby the first ETF combines top large-cap choices with select DEX brand tokens. There are also plans for additional ETF funds to be created based on community governance through the [DAO](https://iq.wiki/wiki/dao). This approach allows investors to collectively shape the inclusion of tokens within future funds. [\\[5\\]](#cite-id-5omrh7mdkzk)  \n  \n### deETF Tokens  \n  \nAn ETF token is a token whose value is formed from the price of the assets within it. This token serves to diversify funds and is already divided into a list of coins. The price of this token changes depending on the price of the assets within it. Fund tokens in an ETF are visible by expanding the fund and looking at the list. [\\[6\\]](#cite-id-zrndo24zdea)  \n  \n## dexPAD  \n  \ndexPAD is a project that aims to provide new project creators with multiple services and technologies needed to succeed, including an Initial Decentralized Offering (IDO) platform, community support, and other platform solutions. [\\[7\\]](#cite-id-kdn44vguwne)  \n  \n> \"Projects launched on dexPAD are 100% fairly allocated. Allocation is based on the amount of dexIRA tokens held prior to a given launch. The numbers of tokens held per tier will be used for governance.\" - the website stated[\\[7\\]](#cite-id-kdn44vguwne)  \n  \ndexPAD uses a tier-based system to determine the amount of dexIRA tokens needed to secure a seat for the allocation. Tier 1 includes a dex holding of $100K, Tier 2 — $500K, Tier 3 — $1 million, and Tier 4 — $5 million. [\\[7\\]](#cite-id-kdn44vguwne)  \n  \n## dexVAULTS  \n  \ndexVAULTS simplifies the advanced trading strategies employed by experienced investors. The task of Dex Finance Vaults is to collect [liquidity pool](https://iq.wiki/wiki/liquidity-pool) pairs of tokens on different Dexes, combine them into one [smart contract](https://iq.wiki/wiki/smart-contract) (vault), and then negotiate them (profit collection and reinvestment). Once the vault with the specified LP pairs has been created and the user credits his funds to them, they are allocated according to the specified shares. Immediately after the investment, rewards are credited, which the user can automatically (if the Auto-harvest threshold is met) or manually collect, part of which is reinvested and part of which is credited as Profit to the account. [\\[8\\]](#cite-id-alxjma035o)  \n  \nEach vault consists of 1 to 3 farms, which differ in both the level of risk and the percentage of rewards accrued. Using Dex Finance users can either invest in stableLP and earn rewards or combine different LPs to create a balanced level of risk and APY of the vaults. [\\[8\\]](#cite-id-alxjma035o)  \n  \n# GDEX Token  \n  \nThe DexFi Governance (GDEX) token is the native token of the DexFi platform. GDEX token holders can participate in the governance of the platform and benefit from the fees generated by the index funds. [\\[12\\]](#cite-id-j7ao907utu)",
							Media: []metadata.Media{{
								Address:  "ipfs://QmamLgsY8XSyYQFZw4rhKWWbVTR8kBrFFegBdLvPjzuLrY",
								MimeType: "image/png",
							}},
							ProfileID:     "",
							PublicationID: "dex-finance",
							ContentURI:    "https://iq.wiki/wiki/dex-finance",
							Tags: []string{
								"Protocols",
							},
							AuthorURL: "https://iq.wiki/account/0x212Cb3F4aE6611054637f9f78F18fB628AD258bb",
							Reward:    nil,
							Timestamp: 0,
							Target:    nil,
							TargetURL: "",
						},
					},
				},
				Status:    true,
				Timestamp: 1693414028,
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
