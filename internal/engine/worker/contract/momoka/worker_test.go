package momoka_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/config"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/arweave"
	worker "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/momoka"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/endpoint"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Arweave(t *testing.T) {
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
			name: "Momoka V1 Post",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1692672514,
					},
					Transaction: arweave.Transaction{
						ID:       "7Nv-EG4QBa0yqwvOfL8vH_mKlj6-VkoobZ9CySDHPqA",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweGNlMWYyY2E4ZmU2NzFlZTQ1ZjlmMzI4MjRlZmM1ODRiM2Y2MGUxNGI5NjExZDYwYjYyOGE2MWFhN2ZlODYxZWQ3ZDcwZDExNjYwNWFjZjEwNGE1NGZlNDNhNjc4M2ZmMTY4MjU1NDczNzQ3MjY4ZjNlYTRmNWY4ZDMzOGMxMzk2MWMiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiI1ZDc1ZjMxZC0wMzc0LTQzNjktYjIwYy0xOGYxYmExYzA2OTkiLCJ0eXBlIjoiUE9TVF9DUkVBVEVEIiwidGltZXN0YW1wUHJvb2ZzIjp7InR5cGUiOiJCVU5ETFIiLCJoYXNoUHJlZml4IjoiMSIsInJlc3BvbnNlIjp7ImlkIjoiRzdJbnVFX0FJRE1UMzliMnl1YzZvdjRNZW16V212SGZsVVVfVmFFY2ltRSIsInRpbWVzdGFtcCI6MTY5MjY3MDQ3NDM1OCwidmVyc2lvbiI6IjEuMC4wIiwicHVibGljIjoic3E5SmJwcEtMbEFLdFF3YWxmWDVEYWduR01sVGlyZGl0WGs3eTRqZ29lQTdERU0wWjZjVlBFNXhNUTlrel9UOVZwcFA2QkZIdEh5WkNaT0RlcmNFVldpcHprcjM2dGZRa1I1RURHVVF5TGl2ZHhVemJXZ1Zrenc3RDI3UEpFYTRjZDFVeTZyMThyWUxxRVJnYlJ2QVpwaDVZSlptcFNKazdyM013blFxdXVrdGp2U3BmQ0xGd1N4UDF3ODc5LXNzX0phbE05SUN6UmkzOGhlbk9OaW84Z2xsNkdWOS1vbXJXd1JNWmVyXzE1YnNwQ0s1dHhDd3BZMTM3bmZLd0tENVlCQXV6eHhjajQyNE03emxTSGxzYWZCd2FSd0ZiZjhnSHRXMDNpSkVSNGxSNEd4ZVkwV3ZuWWFCM0tESVNIUXA1M2E5bmxibWlXTzVXY0hIWXNSODNPVDJlSjBQbDNSV0EtX2lta19TTndHUVRDam1BNnRmX1VWd0w4SHpZUzJpeXV1ODViN2lZSzlaUW9oOG5xYk5DNnFpYklDRTRoOUZlM2JON0FnaXRJZTlYekNUT1hEZk1yNGFoakM4a2txSjF6NHpOQUk2LUxlZWlfTWdkOEp0WmgydnFGTlpoWEswbFNhZEZsXzlPaDNBRVQ3dFVkczJFN3MtNnpwUlBkOW9CWnU2LWtOdUhEUko2VFFoWlN3SjlaTzVIWXNjY2JfR18xc283MmFYSnltUjlnZ0pnV3I0SjNiYXdBWVlucW12bXpHa2xZT2xFXzVIVm5NeGYtVXhwVDd6dGRzSGJjOVFFSDZXMmJ6d3hicGpUY3pFWnMzSkNDQjNjLU5ld05Ic2o5UFlNM2I1dFRsVE5QOWtOQXdQWkhXcHQxMXQ3OUx1TmtOR3Q5TGZPZWsiLCJzaWduYXR1cmUiOiJpUkF4UlFYNjZsLUozOFBkS2hwMi1oekxCOUFuSTUwR2ktMmlIOHo3SUZVZkVvTFpDN1BLZlVSMEtpREZpRUloU2JZWldob3N1NjZSNlNkeDJWSElSTnhNNmhNMklNb0ZUcWxSeVZCdDVnUEhHbFhFSC02N081OUdqZ0VEb180aWpXZ1lVUG9hRGlqYVVqT002T0Uzd2RVeFdfcW9wTG12RjUtX0JpM1Bqc3VOY2F1TGZPVGQyUVhrNk5EcG5NVXRvYTlFY1lBdThlZlFaZzc1VjNHUS1nY09JYXlTYW9YQ3ZLbjlHMGp3YmZWNGxid0ZXVFA5UTJrRENyYVkxX203WS1xeXlzY0diWHhDMWF6b3JVZENZaExnTXFZTUhUbVg3aFU4Y2tTQWdtQmFiZE1NUXhXMUVsV1lJUTN1aHhpUklHc0dzVUtnaldEbE5lRFYzX2dPekFSNjVOZVIwZ1dwczNjcmFGRDN1cEhLN01SMUlxa052YmlDMDVqemRsZXNJWmZXamxwQmZCSkpuWnNCNEhLTGxnWWNmd0xvT1MySDdkS1BQaDhUeUttaW9zNW5NcGpTS1hTMy1FU2VMU2ltZmVncWE2MU1UOUFoQ0N5akNFcExVc1FScGlnQl9HeXFUNlNHaW80V2VSWWlkMENjRGJoU2dwQTlMYnNDNnR3a0xGYjh4QS1PUS1hcXJpTGYzZFhiTnRwVldHenl3YmJZckMwRkxKTzVONk5uTXZQM3BCVGtBaUhCVmRUWjlHcjVUNUdxTjBJQ19vZFdMdDhzUW92WUpyWFRDME85UEVRaldIVldmMXI1ZVlCNG5taDd0ZzJOcE1OclJUVEwzb0RMMVUzbHNfdU55YjdvWmx2TFI1S2hmbWdmcnhieVBoYlo0eWdsNzZnNkNBdyIsImRlYWRsaW5lSGVpZ2h0IjoxMjQ5NTkyLCJibG9jayI6MTI0OTU5MiwidmFsaWRhdG9yU2lnbmF0dXJlcyI6W119fSwiY2hhaW5Qcm9vZnMiOnsidGhpc1B1YmxpY2F0aW9uIjp7InNpZ25hdHVyZSI6IjB4YTVhZGYxM2NhOTU2ODczYTZiZDA3Yzg4ZWMzYjNhMGMxYzA0YjdhYTM2OWM0ZWQxNzk5MzcwY2I3NGViYzNlMzQyOTNjZTAzY2YwMjRmZWUyNmIwYmUyYTVmNDFlNjk4NmU3NDE0ZDVjOWExZTEyNzczM2VjYWUxYmQ4YjJlNzAxYiIsInNpZ25lZEJ5RGVsZWdhdGUiOnRydWUsInNpZ25hdHVyZURlYWRsaW5lIjoxNjkyNjcwNDczLCJ0eXBlZERhdGEiOnsidHlwZXMiOnsiUG9zdFdpdGhTaWciOlt7Im5hbWUiOiJwcm9maWxlSWQiLCJ0eXBlIjoidWludDI1NiJ9LHsibmFtZSI6ImNvbnRlbnRVUkkiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoiY29sbGVjdE1vZHVsZSIsInR5cGUiOiJhZGRyZXNzIn0seyJuYW1lIjoiY29sbGVjdE1vZHVsZUluaXREYXRhIiwidHlwZSI6ImJ5dGVzIn0seyJuYW1lIjoicmVmZXJlbmNlTW9kdWxlIiwidHlwZSI6ImFkZHJlc3MifSx7Im5hbWUiOiJyZWZlcmVuY2VNb2R1bGVJbml0RGF0YSIsInR5cGUiOiJieXRlcyJ9LHsibmFtZSI6Im5vbmNlIiwidHlwZSI6InVpbnQyNTYifSx7Im5hbWUiOiJkZWFkbGluZSIsInR5cGUiOiJ1aW50MjU2In1dfSwiZG9tYWluIjp7Im5hbWUiOiJMZW5zIFByb3RvY29sIFByb2ZpbGVzIiwidmVyc2lvbiI6IjEiLCJjaGFpbklkIjoxMzcsInZlcmlmeWluZ0NvbnRyYWN0IjoiMHhEYjQ2ZDFEYzE1NTYzNEZiQzczMmY5MkU4NTNiMTBCMjg4QUQ1YTFkIn0sInZhbHVlIjp7InByb2ZpbGVJZCI6IjB4NTNhZiIsImNvbnRlbnRVUkkiOiJodHRwczovL2RhdGEubGVucy5waGF2ZXIuY29tL2FwaS9sZW5zL3Bvc3RzLzFmZGNjN2NlLTkxYTctNGFmNy04MDIyLTEzMTMyODQyYTVlYyIsImNvbGxlY3RNb2R1bGUiOiIweGEzMUZGODVFODQwRUQxMTdFMTcyQkM5QWQ4OUU1NTEyOEE5OTkyMDUiLCJjb2xsZWN0TW9kdWxlSW5pdERhdGEiOiIweCIsInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZUluaXREYXRhIjoiMHgiLCJub25jZSI6MCwiZGVhZGxpbmUiOjE2OTI2NzA0NzN9fSwiYmxvY2tIYXNoIjoiMHg0YTU3MWM0YTcwYWEzN2QxYWE1NWJlYmQ3ZGY1MTJmOTkyOTMyNWM3MDJjNjEwZWM1ODQzMTJlZDcxNGI4OTc1IiwiYmxvY2tOdW1iZXIiOjQ2NjAwMTM2LCJibG9ja1RpbWVzdGFtcCI6MTY5MjY3MDQ3M30sInBvaW50ZXIiOm51bGx9LCJwdWJsaWNhdGlvbklkIjoiMHg1M2FmLTB4MDUtREEtNWQ3NWYzMWQiLCJldmVudCI6eyJwcm9maWxlSWQiOiIweDUzYWYiLCJwdWJJZCI6IjB4MDUiLCJjb250ZW50VVJJIjoiaHR0cHM6Ly9kYXRhLmxlbnMucGhhdmVyLmNvbS9hcGkvbGVucy9wb3N0cy8xZmRjYzdjZS05MWE3LTRhZjctODAyMi0xMzEzMjg0MmE1ZWMiLCJjb2xsZWN0TW9kdWxlIjoiMHhhMzFGRjg1RTg0MEVEMTE3RTE3MkJDOUFkODlFNTUxMjhBOTk5MjA1IiwiY29sbGVjdE1vZHVsZVJldHVybkRhdGEiOiIweCIsInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZVJldHVybkRhdGEiOiIweCIsInRpbWVzdGFtcCI6MTY5MjY3MDQ3M319",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "7Nv-EG4QBa0yqwvOfL8vH_mKlj6-VkoobZ9CySDHPqA",
				Network:  filter.NetworkArweave,
				Index:    0,
				From:     "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
				To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
				Type:     filter.TypeSocialPost,
				Platform: lo.ToPtr(filter.PlatformLens),
				Fee: &schema.Fee{
					Address: nil,
					Amount:  decimal.Zero,
					Decimal: 12,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialPost,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformLens.String(),
						From:     "0xa5c2A34384e38f74181756860DaA48e8C80e991A",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle: "pooly66.lens",
							Body:   "😕🙁☹️😞😔😟",
							Media: []metadata.Media{
								{
									Address:  "https://images.lens.phaver.com/insecure/raw:1/plain/a1aebf7d47ef88929b25f962138eed10",
									MimeType: "image/jpeg",
								},
							},
							ProfileID:     "0x53af",
							PublicationID: "0x53af-0x05-DA-5d75f31d",
							ContentURI:    "https://data.lens.phaver.com/api/lens/posts/1fdcc7ce-91a7-4af7-8022-13132842a5ec",
							Timestamp:     1692670473,
						},
					},
				},
				Status:    true,
				Timestamp: 1692672514,
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
