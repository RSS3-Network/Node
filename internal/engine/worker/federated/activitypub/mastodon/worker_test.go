package mastodon

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rss3-network/node/internal/engine/source/activitypub"
	message "github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/stretchr/testify/require"
)

func TestWorker(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task *activitypub.Task
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *activity.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Create A Note",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://airwaves.social/users/VicRB/statuses/112836523057177095",
						Type:      "Create",
						Actor:     "https://airwaves.social/users/VicRB",
						Published: "2024-07-23T15:31:43Z",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://airwaves.social/users/VicRB/statuses/112836523057177095",
							"content":   "<p>#VesselAlert #Vaixell ... </p>",
							"published": "2024-07-23T15:31:43Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:       "https://airwaves.social/users/VicRB/statuses/112836523057177095",
				Network:  network.Mastodon,
				Platform: federated.PlatformMastodon.String(),
				From:     "@VicRB@airwaves.social",
				To:       "@VicRB@airwaves.social",
				Type:     typex.SocialPost,
				Status:   true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialPost,
						Platform: federated.PlatformMastodon.String(),
						From:     "@VicRB@airwaves.social",
						To:       "@VicRB@airwaves.social",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://airwaves.social/users/VicRB/statuses/112836523057177095",
							Body:          "<p>#VesselAlert #Vaixell ... </p>",
							Handle:        "@VicRB@airwaves.social",
							ProfileID:     "https://airwaves.social/users/VicRB",
							Tags:          nil,
							Timestamp:     1721748703,
						},
					},
				},
				Timestamp: 1721748703,
			},
			wantError: require.NoError,
		},
		{
			name: "Create A Comment",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://beekeeping.ninja/users/Pagan_Animist/statuses/112840117527501203",
						Type:      "Create",
						Published: "2024-07-24T06:45:51Z",
						Actor:     "https://beekeeping.ninja/users/Pagan_Animist",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://beekeeping.ninja/users/Pagan_Animist/statuses/112840117527501203",
							"content":   "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mas.to/@evedazzle\" class=\"u-url mention\">@<span>evedazzle</span></a></span> </p><p>Can communities band together ...</p><p>Would they help?</p><p>I’m just thinking of your power grid and next time.</p>",
							"inReplyTo": "https://mas.to/users/evedazzle/statuses/112802025232873362",
							"published": "2024-07-24T06:45:51Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:       "https://beekeeping.ninja/users/Pagan_Animist/statuses/112840117527501203",
				Network:  network.Mastodon,
				Platform: federated.PlatformMastodon.String(),
				From:     "@Pagan_Animist@beekeeping.ninja",
				To:       "@evedazzle@mas.to",
				Type:     typex.SocialComment,
				Status:   true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialComment,
						Platform: federated.PlatformMastodon.String(),
						From:     "@Pagan_Animist@beekeeping.ninja",
						To:       "@evedazzle@mas.to",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://beekeeping.ninja/users/Pagan_Animist/statuses/112840117527501203",
							Body:          "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mas.to/@evedazzle\" class=\"u-url mention\">@<span>evedazzle</span></a></span> </p><p>Can communities band together ...</p><p>Would they help?</p><p>I’m just thinking of your power grid and next time.</p>",
							Handle:        "@Pagan_Animist@beekeeping.ninja",
							ProfileID:     "https://beekeeping.ninja/users/Pagan_Animist",
							Target: &metadata.SocialPost{
								PublicationID: "https://mas.to/users/evedazzle/statuses/112802025232873362",
							},
							Timestamp: 1721803551,
						},
					},
				},
				Timestamp: 1721803551,
			},
			wantError: require.NoError,
		},
		{
			name: "Create A Comment (Includes 2 Mentions)",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://epicure.social/users/Island_Martha/statuses/112840097961400438",
						Type:      "Create",
						Actor:     "https://epicure.social/users/Island_Martha",
						Published: "2024-07-24T06:40:52Z",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://epicure.social/users/Island_Martha/statuses/112840097961400438",
							"content":   "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@SpookieRobieTheCat\" class=\"u-url mention\">@<span>SpookieRobieTheCat</span></a></span> <span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@eunews\" class=\"u-url mention\">@<span>eunews</span></a></span> <br />Or at very least restrict or remove their voting rights. Orban is not MAGA Mike</p>",
							"inReplyTo": "https://mastodon.social/users/SpookieRobieTheCat/statuses/112840076342641439",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
							"tag": []map[string]string{
								{
									"type": "Mention",
									"href": "https://mastodon.social/@SpookieRobieTheCat",
									"name": "@SpookieRobieTheCat",
								},
								{
									"type": "Mention",
									"href": "https://mastodon.social/@eunews",
									"name": "@eunews",
								},
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:       "https://epicure.social/users/Island_Martha/statuses/112840097961400438",
				Network:  network.Mastodon,
				Platform: federated.PlatformMastodon.String(),
				From:     "@Island_Martha@epicure.social",
				To:       "@SpookieRobieTheCat@mastodon.social",
				Type:     typex.SocialComment,
				Status:   true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialComment,
						Platform: federated.PlatformMastodon.String(),
						From:     "@Island_Martha@epicure.social",
						To:       "@SpookieRobieTheCat@mastodon.social",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://epicure.social/users/Island_Martha/statuses/112840097961400438",
							Body:          "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@SpookieRobieTheCat\" class=\"u-url mention\">@<span>SpookieRobieTheCat</span></a></span> <span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@eunews\" class=\"u-url mention\">@<span>eunews</span></a></span> <br />Or at very least restrict or remove their voting rights. Orban is not MAGA Mike</p>",
							Handle:        "@Island_Martha@epicure.social",
							ProfileID:     "https://epicure.social/users/Island_Martha",
							Target: &metadata.SocialPost{
								PublicationID: "https://mastodon.social/users/SpookieRobieTheCat/statuses/112840076342641439",
							},
							Timestamp: 1721803252,
						},
					},
					{
						Type:     typex.SocialShare,
						Platform: federated.PlatformMastodon.String(),
						From:     "@Island_Martha@epicure.social",
						To:       "@SpookieRobieTheCat@mastodon.social",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://epicure.social/users/Island_Martha/statuses/112840097961400438",
							Body:          "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@SpookieRobieTheCat\" class=\"u-url mention\">@<span>SpookieRobieTheCat</span></a></span> <span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@eunews\" class=\"u-url mention\">@<span>eunews</span></a></span> <br />Or at very least restrict or remove their voting rights. Orban is not MAGA Mike</p>",
							Handle:        "@Island_Martha@epicure.social",
							ProfileID:     "https://epicure.social/users/Island_Martha",
							Target: &metadata.SocialPost{
								PublicationID: "https://mastodon.social/users/SpookieRobieTheCat/statuses/112840076342641439",
							},
							Timestamp: 1721803252,
						},
					},
					{
						Type:     typex.SocialShare,
						Platform: federated.PlatformMastodon.String(),
						From:     "@Island_Martha@epicure.social",
						To:       "@eunews@mastodon.social",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://epicure.social/users/Island_Martha/statuses/112840097961400438",
							Body:          "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@SpookieRobieTheCat\" class=\"u-url mention\">@<span>SpookieRobieTheCat</span></a></span> <span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.social/@eunews\" class=\"u-url mention\">@<span>eunews</span></a></span> <br />Or at very least restrict or remove their voting rights. Orban is not MAGA Mike</p>",
							Handle:        "@Island_Martha@epicure.social",
							ProfileID:     "https://epicure.social/users/Island_Martha",
							Target: &metadata.SocialPost{
								PublicationID: "https://mastodon.social/users/SpookieRobieTheCat/statuses/112840076342641439",
							},
							Timestamp: 1721803252,
						},
					},
				},
				Timestamp: 1721803252,
			},
			wantError: require.NoError,
		},
		{
			name: "Announce(Share)",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://relay.an.exchange/activities/d93bf6f6-832d-49d0-b841-3654d8da0b79",
						Type:      "Announce",
						Actor:     "https://relay.an.exchange/actor",
						Published: "2024-07-22T00:00:00Z",
						Object: map[string]interface{}{
							"type": "Note",
							"id":   "https://cr8r.gg/users/SarraceniaWilds#announces/112250337669855051/undo",
						},
					},
				},
			},
			want: &activity.Activity{
				ID:       "https://relay.an.exchange/activities/d93bf6f6-832d-49d0-b841-3654d8da0b79",
				Network:  network.Mastodon,
				Platform: federated.PlatformMastodon.String(),
				From:     "@relay@relay.an.exchange",
				To:       "@relay@relay.an.exchange",
				Type:     typex.SocialShare,
				Status:   true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialShare,
						Platform: federated.PlatformMastodon.String(),
						From:     "@relay@relay.an.exchange",
						To:       "@SarraceniaWilds@cr8r.gg",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://relay.an.exchange/activities/d93bf6f6-832d-49d0-b841-3654d8da0b79",
							ProfileID:     "https://relay.an.exchange/actor",
							Handle:        "@SarraceniaWilds@cr8r.gg",
							Target: &metadata.SocialPost{
								PublicationID: "https://cr8r.gg/users/SarraceniaWilds#announces/112250337669855051/undo",
							},
							Timestamp: 1721606400,
						},
					},
				},
				Timestamp: 1721606400,
			},
			wantError: require.NoError,
		},
		{
			name: "Like",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://mock.social/activities/like123",
						Type:      "Like",
						Actor:     "https://beekeeping.ninja/users/Pagan_Animist",
						Published: "2024-07-22T00:00:00Z",
						Object: map[string]interface{}{
							"type": "Note",
							"id":   "https://mock.social/notes/note123",
						},
					},
				},
			},
			want: &activity.Activity{
				ID:       "https://mock.social/activities/like123",
				Network:  network.Mastodon,
				Platform: federated.PlatformMastodon.String(),
				From:     "@Pagan_Animist@beekeeping.ninja",
				To:       "",
				Type:     typex.SocialComment,
				Status:   true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialComment,
						Platform: federated.PlatformMastodon.String(),
						From:     "https://beekeeping.ninja/users/Pagan_Animist",
						To:       "",
						Metadata: &metadata.SocialPost{
							PublicationID: "https://mock.social/activities/like123",
							ProfileID:     "https://beekeeping.ninja/users/Pagan_Animist",
							Target: &metadata.SocialPost{
								PublicationID: "https://mock.social/notes/note123",
							},
							Timestamp: 1721606400,
						},
					},
				},
				Timestamp: 1721606400,
			},
			wantError: require.NoError,
		},
	}
	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := NewWorker()
			require.NoError(t, err)

			activity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			data, err := json.MarshalIndent(activity, "", "\x20\x20")
			require.NoError(t, err)

			t.Log(string(data))

			require.Equal(t, testcase.want, activity)
		})
	}
}
