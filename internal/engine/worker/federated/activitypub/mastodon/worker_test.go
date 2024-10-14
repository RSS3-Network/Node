package mastodon

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/adrianbrad/psqldocker"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/engine/source/activitypub"
	message "github.com/rss3-network/node/provider/activitypub"
	redisx "github.com/rss3-network/node/provider/redis"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce      sync.Once
	redisClient    rueidis.Client
	databaseClient database.Client
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		// Setup database client
		databaseClient = setupDatabaseClient(t)

		// Start Redis container with TLS
		preset := redis.Preset(
			redis.WithVersion("6.0.9"),
		)

		var container *gnomock.Container
		for attempts := 0; attempts < 3; attempts++ {
			container, err = gnomock.Start(preset)
			if err == nil {
				break
			}
			// If starting fails, wait a bit and try again
			time.Sleep(time.Second * 2)
		}
		require.NoError(t, err, "Failed to start Redis container after multiple attempts")

		t.Cleanup(func() {
			require.NoError(t, gnomock.Stop(container))
		})

		// Connect to Redis without TLS
		redisClient, err = redisx.NewClient(config.Redis{
			Endpoint: container.DefaultAddress(),
			TLS: config.RedisTLS{
				Enabled:            false,
				CAFile:             "/path/to/ca.crt",
				CertFile:           "/path/to/client.crt",
				KeyFile:            "/path/to/client.key",
				InsecureSkipVerify: false,
			},
		})
		require.NoError(t, err)
	})
}

func setupDatabaseClient(t *testing.T) database.Client {
	container, dataSourceName, err := createContainer(context.Background(), database.DriverPostgreSQL, true)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, container.Close())
	})

	// Dial the database
	client, err := dialer.Dial(context.Background(), &config.Database{
		Driver:    database.DriverPostgreSQL,
		URI:       dataSourceName,
		Partition: lo.ToPtr(true),
	})
	require.NoError(t, err)
	require.NotNil(t, client)

	// Migrate the database
	require.NoError(t, client.Migrate(context.Background()))

	return client
}

func createContainer(_ context.Context, driver database.Driver, _ bool) (container *psqldocker.Container, dataSourceName string, err error) {
	switch driver {
	case database.DriverPostgreSQL:
		nBig, err := rand.Int(rand.Reader, big.NewInt(100000))
		if err != nil {
			return nil, "", fmt.Errorf("failed to generate secure random number: %w", err)
		}

		containerName := fmt.Sprintf("psql-test-container-%d", nBig.Int64())

		c, err := psqldocker.NewContainer(
			"user",
			"password",
			"test",
			psqldocker.WithContainerName(containerName),
		)
		if err != nil {
			return nil, "", fmt.Errorf("create psql container: %w", err)
		}

		return c, formatContainerURI(c), nil
	default:
		return nil, "", fmt.Errorf("unsupported driver: %s", driver)
	}
}

func formatContainerURI(container *psqldocker.Container) string {
	return fmt.Sprintf(
		"postgres://user:password@%s:%s/%s?sslmode=disable",
		"127.0.0.1",
		container.Port(),
		"test",
	)
}

func TestWorker(t *testing.T) {
	t.Parallel()
	setup(t)

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
						ID:        "https://digitalcourage.social/users/Volksverpetzer/statuses/113287832117292671/activity",
						Type:      "Create",
						Actor:     "https://digitalcourage.social/users/Volksverpetzer",
						Published: "2024-10-11T08:25:33Z",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://digitalcourage.social/users/Volksverpetzer/statuses/113287832117292671",
							"content":   "<p>Die Lügen der Rechten kosten uns Millionen: Arbeitende Schutzsuchende schieben wir ab und die steigenden Kosten müssen junge Menschen zahlen. Der Rassismus von AfD &amp; Co. wird junge Menschen teuer zu stehen kommen. <a href=\"https://www.volksverpetzer.de/analyse/wollen-nicht-fuer-rassismus-zahlen/?utm_source=mstdn\" target=\"_blank\" rel=\"nofollow noopener noreferrer\" translate=\"no\"><span class=\"invisible\">https://www.</span><span class=\"ellipsis\">volksverpetzer.de/analyse/woll</span><span class=\"invisible\">en-nicht-fuer-rassismus-zahlen/?utm_source=mstdn</span></a></p></p>",
							"published": "2024-10-11T08:25:33Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:           "https://digitalcourage.social/users/Volksverpetzer/statuses/113287832117292671/activity",
				Network:      network.Mastodon,
				Platform:     federated.PlatformMastodon.String(),
				From:         "@Volksverpetzer@digitalcourage.social",
				To:           "@Volksverpetzer@digitalcourage.social",
				Type:         typex.SocialPost,
				Tag:          tag.Social,
				Status:       true,
				TotalActions: 1,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialPost,
						Tag:      tag.Social,
						Platform: federated.PlatformMastodon.String(),
						From:     "@Volksverpetzer@digitalcourage.social",
						To:       "@Volksverpetzer@digitalcourage.social",
						Metadata: &metadata.SocialPost{
							PublicationID: "113287832117292671",
							Body:          "Die Lügen der Rechten kosten uns Millionen: Arbeitende Schutzsuchende schieben wir ab und die steigenden Kosten müssen junge Menschen zahlen. Der Rassismus von AfD & Co. wird junge Menschen teuer zu stehen kommen. https://www. volksverpetzer.de/analyse/woll en-nicht-fuer-rassismus-zahlen/?utm_source=mstdn ( https://www.volksverpetzer.de/analyse/wollen-nicht-fuer-rassismus-zahlen/?utm_source=mstdn )",
							Handle:        "@Volksverpetzer@digitalcourage.social",
							Timestamp:     1728635133,
						},
						RelatedURLs: []string{"https://digitalcourage.social/users/Volksverpetzer/statuses/113287832117292671"},
					},
				},
				Timestamp: 1728635133,
			},
			wantError: require.NoError,
		},
		{
			name: "Create A Note with Media",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://mastodon.world/users/ctxt/statuses/113287758853959995/activity",
						Type:      "Create",
						Actor:     "https://mastodon.world/users/ctxt",
						Published: "2024-10-11T08:06:55Z",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://mastodon.world/users/ctxt/statuses/113287758853959995",
							"content":   "<p>Y por si se les pasó, les dejamos también la pieza que nos mandó Lola Matamala contándonos en primera persona el desahucio que sufrió su madre el mes pasado después de llevar setenta años viviendo en la misma calle. </p><p>La empresa Dapamali Works compró el edifio y fue echando uno por uno a todos los vecinos: </p><p><a href=\"https://ctxt.es/es/20241001/Firmas/47602/lola-matamala-desahucio-Dapamali-Works-primera-persona-aurora-getafe.htm\" target=\"_blank\" rel=\"nofollow noopener noreferrer\" translate=\"no\"><span class=\"invisible\">https://</span><span class=\"ellipsis\">ctxt.es/es/20241001/Firmas/476</span><span class=\"invisible\">02/lola-matamala-desahucio-Dapamali-Works-primera-persona-aurora-getafe.htm</span></a></p>",
							"published": "2024-10-11T08:06:55Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
							"cc": []string{
								"https://mastodon.world/users/ctxt/followers",
							},
							"sensitive":    false,
							"conversation": "tag:mastodon.world,2024-10-11:objectId=314975275:objectType=Conversation",
							"attachment": []map[string]interface{}{
								{
									"type":      "Document",
									"mediaType": "image/jpeg",
									"url":       "https://s3.eu-central-2.wasabisys.com/mastodonworld/media_attachments/files/113/287/758/697/714/115/original/2c6c341deae150c0.jpg",
									"blurhash":  "UOFiJm4nxv-=_M9Yxu-ptRM{odaes:jEayay",
									"width":     800,
									"height":    493,
								},
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:           "https://mastodon.world/users/ctxt/statuses/113287758853959995/activity",
				Network:      network.Mastodon,
				Platform:     federated.PlatformMastodon.String(),
				From:         "@ctxt@mastodon.world",
				To:           "@ctxt@mastodon.world",
				Type:         typex.SocialPost,
				Tag:          tag.Social,
				Status:       true,
				TotalActions: 1,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialPost,
						Tag:      tag.Social,
						Platform: federated.PlatformMastodon.String(),
						From:     "@ctxt@mastodon.world",
						To:       "@ctxt@mastodon.world",
						Metadata: &metadata.SocialPost{
							PublicationID: "113287758853959995",
							Body:          "Y por si se les pasó, les dejamos también la pieza que nos mandó Lola Matamala contándonos en primera persona el desahucio que sufrió su madre el mes pasado después de llevar setenta años viviendo en la misma calle.\n\nLa empresa Dapamali Works compró el edifio y fue echando uno por uno a todos los vecinos:\n\nhttps:// ctxt.es/es/20241001/Firmas/476 02/lola-matamala-desahucio-Dapamali-Works-primera-persona-aurora-getafe.htm ( https://ctxt.es/es/20241001/Firmas/47602/lola-matamala-desahucio-Dapamali-Works-primera-persona-aurora-getafe.htm )",
							Handle:        "@ctxt@mastodon.world",
							Timestamp:     1728634015,
							Media: []metadata.Media{
								{
									Address:  "https://s3.eu-central-2.wasabisys.com/mastodonworld/media_attachments/files/113/287/758/697/714/115/original/2c6c341deae150c0.jpg",
									MimeType: "image/jpeg",
								},
							},
						},
						RelatedURLs: []string{"https://mastodon.world/users/ctxt/statuses/113287758853959995"},
					},
				},
				Timestamp: 1728634015,
			},
			wantError: require.NoError,
		},
		{
			name: "Create A Note with Tags",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
							map[string]string{
								"ostatus":          "http://ostatus.org#",
								"atomUri":          "ostatus:atomUri",
								"inReplyToAtomUri": "ostatus:inReplyToAtomUri",
								"conversation":     "ostatus:conversation",
								"sensitive":        "as:sensitive",
								"toot":             "http://joinmastodon.org/ns#",
								"votersCount":      "toot:votersCount",
							},
						},
						ID:        "https://mastodon.social/users/DJGummikuh/statuses/113287749405285684/activity",
						Type:      "Create",
						Actor:     "https://mastodon.social/users/DJGummikuh",
						Published: "2024-10-11T08:04:31Z",
						To: []string{
							"https://www.w3.org/ns/activitystreams#Public",
						},
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://mastodon.social/users/DJGummikuh/statuses/113287749405285684",
							"content":   "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://mastodon.energy/@Sustainable2050\" class=\"u-url mention\">@<span>Sustainable2050</span></a></span> who ever could have expected that?</p>",
							"published": "2024-10-11T08:04:31Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
							"cc": []string{
								"https://mastodon.social/users/DJGummikuh/followers",
								"https://mastodon.energy/users/Sustainable2050",
							},
							"sensitive":    false,
							"conversation": "tag:mastodon.energy,2024-10-11:objectId=32875409:objectType=Conversation",
							"tag": []map[string]interface{}{
								{
									"type": "Mention",
									"href": "https://mastodon.energy/users/Sustainable2050",
									"name": "@Sustainable2050@mastodon.energy",
								},
								{
									"type": "Mention",
									"href": "https://mastodon.social/users/DJGummikuh",
									"name": "@DJGummikuh@mastodon.social",
								},
								{
									"type": "Hashtag",
									"href": "https://ard.social/tags/ubernachtungen",
									"name": "#ubernachtungen",
								},
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:           "https://mastodon.social/users/DJGummikuh/statuses/113287749405285684/activity",
				Network:      network.Mastodon,
				Platform:     federated.PlatformMastodon.String(),
				From:         "@DJGummikuh@mastodon.social",
				To:           "@DJGummikuh@mastodon.social",
				Type:         typex.SocialPost,
				Tag:          tag.Social,
				Status:       true,
				TotalActions: 1,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialPost,
						Tag:      tag.Social,
						Platform: federated.PlatformMastodon.String(),
						From:     "@DJGummikuh@mastodon.social",
						To:       "@DJGummikuh@mastodon.social",
						Metadata: &metadata.SocialPost{
							PublicationID: "113287749405285684",
							Body:          "@ Sustainable2050 ( https://mastodon.energy/@Sustainable2050 ) who ever could have expected that?",
							Handle:        "@DJGummikuh@mastodon.social",
							Timestamp:     1728633871,
							Tags:          []string{"@Sustainable2050@mastodon.energy", "@DJGummikuh@mastodon.social", "ubernachtungen"},
						},
						RelatedURLs: []string{"https://mastodon.social/users/DJGummikuh/statuses/113287749405285684"},
					},
				},
				Timestamp: 1728633871,
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
						ID:        "https://social.timespiral.co.jp/users/find575/statuses/113287741759175713/activity",
						Type:      "Create",
						Published: "2024-10-11T08:02:34Z",
						Actor:     "https://social.timespiral.co.jp/users/find575",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://social.timespiral.co.jp/users/find575/statuses/113287741759175713",
							"content":   "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://pawoo.net/@cs133\" class=\"u-url mention\">@<span>cs133</span></a></span> 俳句を発見しました！<br />『各駅に エスペラントで 愛称が』</p>",
							"inReplyTo": "https://pawoo.net/users/cs133/statuses/113287741711853329",
							"published": "2024-10-11T08:02:34Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:           "https://social.timespiral.co.jp/users/find575/statuses/113287741759175713/activity",
				Network:      network.Mastodon,
				Platform:     federated.PlatformMastodon.String(),
				From:         "@find575@social.timespiral.co.jp",
				To:           "@cs133@pawoo.net",
				Type:         typex.SocialComment,
				Tag:          tag.Social,
				Status:       true,
				TotalActions: 1,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialComment,
						Tag:      tag.Social,
						Platform: federated.PlatformMastodon.String(),
						From:     "@find575@social.timespiral.co.jp",
						To:       "@cs133@pawoo.net",
						Metadata: &metadata.SocialPost{
							PublicationID: "113287741759175713",
							Body:          "@ cs133 ( https://pawoo.net/@cs133 ) 俳句を発見しました！\n『各駅に エスペラントで 愛称が』",
							Handle:        "@find575@social.timespiral.co.jp",
							Target: &metadata.SocialPost{
								Handle:        "@cs133@pawoo.net",
								PublicationID: "113287741711853329",
								Timestamp:     1728633754,
								Body:          "賢治に毒されて各駅にエスペラントで愛称が付けられている釜石線",
							},
							TargetURL: "https://pawoo.net/users/cs133/statuses/113287741711853329",
							Timestamp: 1728633754,
						},
						RelatedURLs: []string{"https://social.timespiral.co.jp/users/find575/statuses/113287741759175713"},
					},
				},
				Timestamp: 1728633754,
			},
			wantError: require.NoError,
		},
		{
			name: "Create A Comment whose target Post contains Media",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://social.timespiral.co.jp/users/find575/statuses/113287741759175713/activity",
						Type:      "Create",
						Published: "2024-10-11T08:02:34Z",
						Actor:     "https://social.timespiral.co.jp/users/find575",
						Object: map[string]interface{}{
							"type":      "Note",
							"id":        "https://social.timespiral.co.jp/users/find575/statuses/113287741759175713",
							"content":   "<p><span class=\"h-card\" translate=\"no\"><a href=\"https://pawoo.net/@cs133\" class=\"u-url mention\">@<span>cs133</span></a></span> 俳句を発見しました！<br />『各駅に エスペラントで 愛称が』</p>",
							"inReplyTo": "https://mastodon.world/users/ctxt/statuses/113287758853959995",
							"published": "2024-10-11T08:02:34Z",
							"to": []string{
								"https://www.w3.org/ns/activitystreams#Public",
							},
						},
					},
				},
			},
			want: &activity.Activity{
				ID:           "https://social.timespiral.co.jp/users/find575/statuses/113287741759175713/activity",
				Network:      network.Mastodon,
				Platform:     federated.PlatformMastodon.String(),
				From:         "@find575@social.timespiral.co.jp",
				To:           "@ctxt@mastodon.world",
				Type:         typex.SocialComment,
				Tag:          tag.Social,
				TotalActions: 1,
				Status:       true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialComment,
						Tag:      tag.Social,
						Platform: federated.PlatformMastodon.String(),
						From:     "@find575@social.timespiral.co.jp",
						To:       "@ctxt@mastodon.world",
						Metadata: &metadata.SocialPost{
							PublicationID: "113287741759175713",
							Handle:        "@find575@social.timespiral.co.jp",
							Body:          "@ cs133 ( https://pawoo.net/@cs133 ) 俳句を発見しました！\n『各駅に エスペラントで 愛称が』",
							Timestamp:     1728633754,
							Target: &metadata.SocialPost{
								Handle:        "@ctxt@mastodon.world",
								Body:          "Y por si se les pasó, les dejamos también la pieza que nos mandó Lola Matamala contándonos en primera persona el desahucio que sufrió su madre el mes pasado después de llevar setenta años viviendo en la misma calle.\n\nLa empresa Dapamali Works compró el edifio y fue echando uno por uno a todos los vecinos:\n\nhttps:// ctxt.es/es/20241001/Firmas/476 02/lola-matamala-desahucio-Dapamali-Works-primera-persona-aurora-getafe.htm ( https://ctxt.es/es/20241001/Firmas/47602/lola-matamala-desahucio-Dapamali-Works-primera-persona-aurora-getafe.htm )",
								PublicationID: "113287758853959995",
								Timestamp:     1728634015,
								Media: []metadata.Media{
									{
										Address:  "https://s3.eu-central-2.wasabisys.com/mastodonworld/media_attachments/files/113/287/758/697/714/115/original/2c6c341deae150c0.jpg",
										MimeType: "Document",
									},
								},
							},
							TargetURL: "https://mastodon.world/users/ctxt/statuses/113287758853959995",
						},
						RelatedURLs: []string{"https://social.timespiral.co.jp/users/find575/statuses/113287741759175713"},
					},
				},
				Timestamp: 1728633754,
			},
			wantError: require.NoError,
		},
		{
			name: "Share A Post",
			arguments: arguments{
				task: &activitypub.Task{
					Network: network.Mastodon,
					Message: message.Object{
						Context: []interface{}{
							"https://www.w3.org/ns/activitystreams",
						},
						ID:        "https://fosstodon.org/users/bert_hubert/statuses/113287843427757878/activity",
						Type:      "Announce",
						Actor:     "https://fosstodon.org/users/bert_hubert",
						Published: "2024-10-11T08:28:26Z",
						Object:    "https://infosec.exchange/users/ravirockks/statuses/113286145052908177",
					},
				},
			},
			want: &activity.Activity{
				ID:           "https://fosstodon.org/users/bert_hubert/statuses/113287843427757878/activity",
				Network:      network.Mastodon,
				Platform:     federated.PlatformMastodon.String(),
				From:         "@bert_hubert@fosstodon.org",
				To:           "@ravirockks@infosec.exchange",
				Type:         typex.SocialShare,
				Tag:          tag.Social,
				TotalActions: 1,
				Status:       true,
				Actions: []*activity.Action{
					{
						Type:     typex.SocialShare,
						Tag:      tag.Social,
						Platform: federated.PlatformMastodon.String(),
						From:     "@bert_hubert@fosstodon.org",
						To:       "@ravirockks@infosec.exchange",
						Metadata: &metadata.SocialPost{
							PublicationID: "113287843427757878",
							Handle:        "@bert_hubert@fosstodon.org",
							Target: &metadata.SocialPost{
								PublicationID: "113286145052908177",
								Handle:        "@ravirockks@infosec.exchange",
								Timestamp:     1728609391,
								Body:          "THE CYBER RESILIENCE ACT HAS BEEN ADOPTED BY THE COUNCIL OF THE EUROPEAN UNION!\n\nThe EU is weeks away from becoming the first jurisdiction with a bespoke regulatory framework for the product security _and_ labelling of all software sold commercially in the EU (save stuff covered by other EU rules like cars and healthtech).\n\nYes, the Yanks (via EO14028–>NIST) defined critical software (the CRA has ‘important products with digital elements’ and ‘critical products with digital elements’), but the Yanks, for now at least, have only gone down the procurement route for regulating vendor SDLCs. The EU, on the other hand, is covering everything sold commercially (bar the stated exceptions) to anyone in the EU.\n\nBig day for all us SDLC regulation people!\n\nWhat happens next: Council and EuroParl President sign it —> Publication in the EU OJ —> Entry into force 20 days later —> Application of most provisions 36 months later.\n\nPress release (includes link to final text): https://www. consilium.europa.eu/en/press/p ress-releases/2024/10/10/cyber-resilience-act-council-adopts-new-law-on-security-requirements-for-digital-products/ ( https://www.consilium.europa.eu/en/press/press-releases/2024/10/10/cyber-resilience-act-council-adopts-new-law-on-security-requirements-for-digital-products/ )",
							},
							TargetURL: "https://infosec.exchange/users/ravirockks/statuses/113286145052908177",
							Timestamp: 1728635306,
						},
						RelatedURLs: []string{"https://fosstodon.org/users/bert_hubert/statuses/113287843427757878"},
					},
				},
				Timestamp: 1728635306,
			},
			wantError: require.NoError,
		},
	}
	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			instance, err := NewWorker(databaseClient, redisClient)
			require.NoError(t, err)

			transformedActivity, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			data, err := json.MarshalIndent(transformedActivity, "", "\x20\x20")
			require.NoError(t, err)

			t.Log(string(data))

			require.Equal(t, testcase.want, transformedActivity)
		})
	}
}
