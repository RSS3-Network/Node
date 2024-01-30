package farcaster_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/internal/database/model"
	source "github.com/rss3-network/node/internal/engine/source/farcaster"
	worker "github.com/rss3-network/node/internal/engine/worker/farcaster"
	message "github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestWorker(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task *source.Task
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *schema.Feed
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Post A Cast",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkFarcaster,
					Message: message.Message{
						Data: message.MessageData{
							Type: message.MessageTypeCastAdd.String(),
							Fid:  14142,
							Profile: &model.Profile{
								Fid:            14142,
								Username:       "brucexc.eth",
								CustodyAddress: "0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55",
								EthAddresses:   []string{"0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0"},
							},
							Timestamp: 88297925,
							Network:   "FARCASTER_NETWORK_MAINNET",
							CastAddBody: &message.CastAddBody{
								EmbedsDeprecated:  []string{},
								Mentions:          []uint64{},
								MentionsUsernames: []string{},
								ParentCastID:      nil,
								ParentCast:        nil,
								ParentURL:         "",
								Text:              "Now I’m just posting this to test how fast we can index this post and delivery into various ecosystem projects built on the RSS3 Network (rss3.io). Hopefully it doesn’t take long…",
								MentionsPositions: []int32{},
								Embeds: []message.Embed{
									{
										URL: "http://rss3.io",
									},
								},
							},
						},
						Hash:            "0x9d72f8030aafa43f4c208b013964a51017a2747c",
						HashScheme:      "HASH_SCHEME_BLAKE3",
						Signature:       "jzQzwRXdT5JDmvDdcTpZkqo4ZLGtoNvBD2AVhHcgvlxsSSlHOKu4jdJVN8p+SX+0Mcaq27AwKDnE4wEvPqlmAg==",
						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
					},
				},
			},
			want: &schema.Feed{
				ID:       "0x0000000000000000000000009d72f8030aafa43f4c208b013964a51017a2747c",
				Network:  filter.NetworkFarcaster,
				From:     common.HexToAddress("0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55").String(),
				To:       common.HexToAddress("0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55").String(),
				Type:     filter.TypeSocialPost,
				Status:   true,
				Platform: lo.ToPtr(filter.PlatformFarcaster),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
						Platform: filter.PlatformFarcaster.String(),
						From:     common.HexToAddress("0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0").String(),
						To:       common.HexToAddress("0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0").String(),
						Metadata: metadata.SocialPost{
							Handle:    "brucexc.eth",
							Body:      "Now I’m just posting this to test how fast we can index this post and delivery into various ecosystem projects built on the RSS3 Network (rss3.io). Hopefully it doesn’t take long…",
							ProfileID: "14142",
							Media: []metadata.Media{
								{Address: "http://rss3.io", MimeType: "text/html; charset=utf-8"},
							},
							PublicationID: common.HexToAddress("0x0000000000000000000000009d72f8030aafa43f4c208b013964a51017a2747c").String(),
						},
					},
				},
				Timestamp: 1697757125,
			},
			wantError: require.NoError,
		},
		{
			name: "Comment A Cast",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkFarcaster,
					Message: message.Message{
						Data: message.MessageData{
							Type: message.MessageTypeCastAdd.String(),
							Fid:  14142,
							Profile: &model.Profile{
								Fid:            14142,
								Username:       "brucexc.eth",
								CustodyAddress: "0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55",
								EthAddresses:   []string{"0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0"},
							},
							Timestamp: 88298491,
							Network:   "FARCASTER_NETWORK_MAINNET",
							CastAddBody: &message.CastAddBody{
								EmbedsDeprecated:  []string{},
								Mentions:          []uint64{},
								MentionsUsernames: []string{},
								ParentCastID: &message.CastID{
									Fid:  23901,
									Hash: "0xe3fbd4ca0ac4f89f0524f9f4ceee49d600d21fdd",
								},
								ParentCast: &message.Message{
									Data: message.MessageData{
										Type: message.MessageTypeCastAdd.String(),
										Fid:  23901,
										Profile: &model.Profile{
											Fid:            23901,
											Username:       "henryqw",
											CustodyAddress: "0xe25228a6525A2090be824d66Bdf6DB8836eCc90C",
											EthAddresses:   []string{"0x827431510a5D249cE4fdB7F00C83a3353F471848"},
										},
										Timestamp: 88288706,
										Network:   "FARCASTER_NETWORK_MAINNET",
										CastAddBody: &message.CastAddBody{
											EmbedsDeprecated:  []string{},
											Mentions:          []uint64{},
											MentionsUsernames: []string{},
											ParentCastID:      nil,
											ParentCast:        nil,
											ParentURL:         "",
											Text:              "Now I’m just posting this to test how fast we can index this post and delivery into various ecosystem projects built on the RSS3 Network (rss3.io). Hopefully it doesn’t take long…",
											MentionsPositions: []int32{},
											Embeds: []message.Embed{
												{
													URL: "http://rss3.io",
												},
											},
										},
									},
									Hash:            "0xe3fbd4ca0ac4f89f0524f9f4ceee49d600d21fdd",
									HashScheme:      "HASH_SCHEME_BLAKE3",
									Signature:       "EPjUbN5JUOdSEbsNXX1qXEdC4yfXuwpn97OYU27IlbQt783xMcjFfOCcWvIbSPoONRApBLo5zR3Tcq6yv09wDA==",
									SignatureScheme: "SIGNATURE_SCHEME_ED25519",
									Signer:          "0xc698a2da58d7a25b01ada34ababd6c52b353a1d22e7786a3563e14410bb19834",
								},
								ParentURL:         "",
								Text:              "about 2 minutes",
								MentionsPositions: []int32{},
								Embeds:            []message.Embed{},
							},
						},
						Hash:            "0x10ae8f78cbbad692c3b330b8970770406dc785ef",
						HashScheme:      "HASH_SCHEME_BLAKE3",
						Signature:       "aqY/kQNP5WQ/zQFRHRiKTY8zIFyCbMhDhVSNPhFd1kgn4y1hz72eBwXeIBS9sR126wkhqSKmb/W0pdsncxsDw==",
						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
					},
				},
			},
			want: &schema.Feed{
				ID:       "0x00000000000000000000000010ae8f78cbbad692c3b330b8970770406dc785ef",
				Network:  filter.NetworkFarcaster,
				From:     common.HexToAddress("0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55").String(),
				To:       common.HexToAddress("0xe25228a6525A2090be824d66Bdf6DB8836eCc90C").String(),
				Type:     filter.TypeSocialComment,
				Platform: lo.ToPtr(filter.PlatformFarcaster),
				Status:   true,
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialComment,
						Platform: filter.PlatformFarcaster.String(),
						From:     common.HexToAddress("0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0").String(),
						To:       common.HexToAddress("0x827431510a5D249cE4fdB7F00C83a3353F471848").String(),
						Metadata: metadata.SocialPost{
							Handle:        "brucexc.eth",
							Body:          "about 2 minutes",
							ProfileID:     "14142",
							PublicationID: common.HexToAddress("0x00000000000000000000000010AE8F78cbBad692C3B330B8970770406DC785eF").String(),
							Target: &metadata.SocialPost{
								Handle:    "henryqw",
								Body:      "Now I’m just posting this to test how fast we can index this post and delivery into various ecosystem projects built on the RSS3 Network (rss3.io). Hopefully it doesn’t take long…",
								ProfileID: "23901",
								Media: []metadata.Media{
									{Address: "http://rss3.io", MimeType: "text/html; charset=utf-8"},
								},
								PublicationID: common.HexToAddress("0x000000000000000000000000e3Fbd4Ca0aC4F89f0524F9f4cEEe49D600D21Fdd").String(),
							},
						},
					},
				},
				Timestamp: 1697757691,
			},
			wantError: require.NoError,
		},
		{
			name: "Share A Cast",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkFarcaster,
					Message: message.Message{
						Data: message.MessageData{
							Type: message.MessageTypeReactionAdd.String(),
							Fid:  14142,
							Profile: &model.Profile{
								Fid:            14142,
								Username:       "brucexc.eth",
								CustodyAddress: "0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55",
								EthAddresses:   []string{"0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0"},
							},
							Timestamp: 90547939,
							Network:   "FARCASTER_NETWORK_MAINNET",
							ReactionBody: &message.ReactionBody{
								Type: message.ReactionTypeRecast.String(),
								TargetCastID: message.CastID{
									Fid:  23901,
									Hash: "0x05536f622dbedaa75ae1bd8ec9ff98bb0df10894",
								},
								TargetCast: &message.Message{
									Data: message.MessageData{
										Type: message.MessageTypeCastAdd.String(),
										Fid:  23901,
										Profile: &model.Profile{
											Fid:            23901,
											Username:       "henryqw",
											CustodyAddress: "0xe25228a6525A2090be824d66Bdf6DB8836eCc90C",
											EthAddresses:   []string{"0x827431510a5D249cE4fdB7F00C83a3353F471848"},
										},
										Timestamp: 90537875,
										Network:   "FARCASTER_NETWORK_MAINNET",
										CastAddBody: &message.CastAddBody{
											EmbedsDeprecated:  []string{},
											Mentions:          []uint64{},
											MentionsUsernames: []string{},
											ParentCastID:      nil,
											ParentCast:        nil,
											ParentURL:         "",
											Text:              "If you see a guy wearing this same shirt everyday in Istanbul, he is got shirt ton of them.",
											MentionsPositions: []int32{},
											Embeds: []message.Embed{
												{
													URL: "https://i.imgur.com/fEEn6S1.jpg",
												},
											},
										},
									},
									Hash:            "0x05536f622dbedaa75ae1bd8ec9ff98bb0df10894",
									HashScheme:      "HASH_SCHEME_BLAKE3",
									Signature:       "EPjUbN5JUOdSEbsNXX1qXEdC4yfXuwpn97OYU27IlbQt783xMcjFfOCcWvIbSPoONRApBLo5zR3Tcq6yv09wDA==",
									SignatureScheme: "SIGNATURE_SCHEME_ED25519",
									Signer:          "0xc698a2da58d7a25b01ada34ababd6c52b353a1d22e7786a3563e14410bb19834",
								},
							},
						},
						Hash:            "0x2931794842a5f3a152bff66cc010120cfe9c6102",
						HashScheme:      "HASH_SCHEME_BLAKE3",
						Signature:       "L5DbR5RqsowyJxxfhgIfdh711e3GyxT+pLCCOTYRNLQ1+UyKikiExlXGXVcWlFHxX2grFNIfFmKeTvsFx9lfAg==",
						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
					},
				},
			},
			want: &schema.Feed{
				ID:       "0x0000000000000000000000002931794842a5f3a152bff66cc010120cfe9c6102",
				Network:  filter.NetworkFarcaster,
				From:     common.HexToAddress("0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55").String(),
				To:       common.HexToAddress("0xe25228a6525A2090be824d66Bdf6DB8836eCc90C").String(),
				Type:     filter.TypeSocialShare,
				Status:   true,
				Platform: lo.ToPtr(filter.PlatformFarcaster),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialShare,
						Platform: filter.PlatformFarcaster.String(),
						From:     common.HexToAddress("0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0").String(),
						To:       common.HexToAddress("0x827431510a5D249cE4fdB7F00C83a3353F471848").String(),
						Metadata: metadata.SocialPost{
							Handle:        "brucexc.eth",
							ProfileID:     "14142",
							PublicationID: common.HexToAddress("0x0000000000000000000000002931794842A5f3A152BFF66cC010120CFE9C6102").String(),
							Target: &metadata.SocialPost{
								Handle:    "henryqw",
								Body:      "If you see a guy wearing this same shirt everyday in Istanbul, he is got shirt ton of them.",
								ProfileID: "23901",
								Media: []metadata.Media{
									{Address: "https://i.imgur.com/fEEn6S1.jpg", MimeType: "image/jpeg"},
								},
								PublicationID: common.HexToAddress("0x00000000000000000000000005536f622DbEdAa75ae1Bd8eC9Ff98Bb0dF10894").String(),
							},
						},
					},
				},
				Timestamp: 1700007139,
			},
			wantError: require.NoError,
		},
		{
			name: "Share Own Cast",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkFarcaster,
					Message: message.Message{
						Data: message.MessageData{
							Type: message.MessageTypeReactionAdd.String(),
							Fid:  14142,
							Profile: &model.Profile{
								Fid:            14142,
								Username:       "brucexc.eth",
								CustodyAddress: "0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55",
								EthAddresses:   []string{"0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0"},
							},
							Timestamp: 90548475,
							Network:   "FARCASTER_NETWORK_MAINNET",
							ReactionBody: &message.ReactionBody{
								Type: message.ReactionTypeRecast.String(),
								TargetCastID: message.CastID{
									Fid:  14142,
									Hash: "0x9d72F8030AAfa43F4c208B013964a51017A2747c",
								},
								TargetCast: &message.Message{
									Data: message.MessageData{
										Type: message.MessageTypeCastAdd.String(),
										Fid:  14142,
										Profile: &model.Profile{
											Fid:            14142,
											Username:       "brucexc.eth",
											CustodyAddress: "0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55",
											EthAddresses:   []string{"0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0"},
										},
										Timestamp: 88297925,
										Network:   "FARCASTER_NETWORK_MAINNET",
										CastAddBody: &message.CastAddBody{
											EmbedsDeprecated:  []string{},
											Mentions:          []uint64{},
											MentionsUsernames: []string{},
											ParentCastID:      nil,
											ParentCast:        nil,
											ParentURL:         "",
											Text:              "Now I’m just posting this to test how fast we can index this post and delivery into various ecosystem projects built on the RSS3 Network (rss3.io). Hopefully it doesn’t take long…",
											MentionsPositions: []int32{},
											Embeds: []message.Embed{
												{
													URL: "http://rss3.io",
												},
											},
										},
									},
									Hash:            "0x9d72f8030aafa43f4c208b013964a51017a2747c",
									HashScheme:      "HASH_SCHEME_BLAKE3",
									Signature:       "jzQzwRXdT5JDmvDdcTpZkqo4ZLGtoNvBD2AVhHcgvlxsSSlHOKu4jdJVN8p+SX+0Mcaq27AwKDnE4wEvPqlmAg==",
									SignatureScheme: "SIGNATURE_SCHEME_ED25519",
									Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
								},
							},
						},
						Hash:            "0x0d62b1610e7dd177363cd3571bfe9a112d6185de",
						HashScheme:      "HASH_SCHEME_BLAKE3",
						Signature:       "YKu7eDm4MfkOZ1FQw6hG8rB6/TawlYKN9Ftp1GsY0T1KIlHzQKQt005Y96f6VnN5ygPbyS68atT782bY59l0DA==",
						SignatureScheme: "SIGNATURE_SCHEME_ED25519",
						Signer:          "0x4157bb4fe638e7143b2bf8f3b636d5322242f6586eac191fc78aa54675360195",
					},
				},
			},
			want: &schema.Feed{
				ID:       "0x0000000000000000000000000d62b1610e7dd177363cd3571bfe9a112d6185de",
				Network:  filter.NetworkFarcaster,
				From:     common.HexToAddress("0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55").String(),
				To:       common.HexToAddress("0xe5d6216F0085a7F6B9b692e06cf5856e6fA41B55").String(),
				Type:     filter.TypeSocialShare,
				Status:   true,
				Platform: lo.ToPtr(filter.PlatformFarcaster),
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialShare,
						Platform: filter.PlatformFarcaster.String(),
						From:     common.HexToAddress("0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0").String(),
						To:       common.HexToAddress("0x8888888198FbdC8c017870cC5d3c96D0cf15C4F0").String(),
						Metadata: metadata.SocialPost{
							Handle:        "brucexc.eth",
							ProfileID:     "14142",
							PublicationID: common.HexToAddress("0x0000000000000000000000000d62b1610E7Dd177363cD3571bFE9A112D6185DE").String(),
							Target: &metadata.SocialPost{
								Handle:    "brucexc.eth",
								Body:      "Now I’m just posting this to test how fast we can index this post and delivery into various ecosystem projects built on the RSS3 Network (rss3.io). Hopefully it doesn’t take long…",
								ProfileID: "14142",
								Media: []metadata.Media{
									{Address: "http://rss3.io", MimeType: "text/html; charset=utf-8"},
								},
								PublicationID: common.HexToAddress("0x00000000000000000000000009d72F8030AAfa43F4c208B013964a51017A2747c").String(),
							},
						},
					},
				},
				Timestamp: 1700007675,
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

			data, err := json.MarshalIndent(feed, "", "\x20\x20")
			require.NoError(t, err)

			t.Log(string(data))

			require.Equal(t, testcase.want, feed)
		})
	}
}
