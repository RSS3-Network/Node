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
		{
			name: "Momoka V1 Mirror",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1684142674,
					},
					Transaction: arweave.Transaction{
						ID:       "Hx828fn_nXZM1xCST_Vejz5JSCbz0XfjhCAOiOuugiM",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweDcwYzc0NTEwMmU2NzFlNGU1ZGIzNWMxN2JlZDg4YjljNGY1Y2NlZjYzNGQ4YjE1ZGE4NTBjZjM1NzFjOTk5MTg0YjEzYzQ2YTgwYTk2MDg2NTQ4YmY4NTVhOWY3OTNiODA1MDNiNTA1ODMzMGZiZTllYjMyYTZmMzBiZTkxYmU1MWIiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiJkMDNkNTMwOS04YWFjLTRjMGEtODE3NS0xZjhkNTlmNzM1MjQiLCJ0eXBlIjoiTUlSUk9SX0NSRUFURUQiLCJ0aW1lc3RhbXBQcm9vZnMiOnsidHlwZSI6IkJVTkRMUiIsImhhc2hQcmVmaXgiOiIxIiwicmVzcG9uc2UiOnsiaWQiOiJJMTVqbjdfRC1qVHZ3U2hOVDNYUlBqN3l4akR0dm1aOUk2UHJSbTM2WUNRIiwidGltZXN0YW1wIjoxNjg0MTQyNTE3MDU2LCJ2ZXJzaW9uIjoiMS4wLjAiLCJwdWJsaWMiOiJzcTlKYnBwS0xsQUt0UXdhbGZYNURhZ25HTWxUaXJkaXRYazd5NGpnb2VBN0RFTTBaNmNWUEU1eE1ROWt6X1Q5VnBwUDZCRkh0SHlaQ1pPRGVyY0VWV2lwemtyMzZ0ZlFrUjVFREdVUXlMaXZkeFV6YldnVmt6dzdEMjdQSkVhNGNkMVV5NnIxOHJZTHFFUmdiUnZBWnBoNVlKWm1wU0prN3IzTXduUXF1dWt0anZTcGZDTEZ3U3hQMXc4Nzktc3NfSmFsTTlJQ3pSaTM4aGVuT05pbzhnbGw2R1Y5LW9tcld3Uk1aZXJfMTVic3BDSzV0eEN3cFkxMzduZkt3S0Q1WUJBdXp4eGNqNDI0TTd6bFNIbHNhZkJ3YVJ3RmJmOGdIdFcwM2lKRVI0bFI0R3hlWTBXdm5ZYUIzS0RJU0hRcDUzYTlubGJtaVdPNVdjSEhZc1I4M09UMmVKMFBsM1JXQS1faW1rX1NOd0dRVENqbUE2dGZfVVZ3TDhIellTMml5dXU4NWI3aVlLOVpRb2g4bnFiTkM2cWliSUNFNGg5RmUzYk43QWdpdEllOVh6Q1RPWERmTXI0YWhqQzhra3FKMXo0ek5BSTYtTGVlaV9NZ2Q4SnRaaDJ2cUZOWmhYSzBsU2FkRmxfOU9oM0FFVDd0VWRzMkU3cy02enBSUGQ5b0JadTYta051SERSSjZUUWhaU3dKOVpPNUhZc2NjYl9HXzFzbzcyYVhKeW1SOWdnSmdXcjRKM2Jhd0FZWW5xbXZtekdrbFlPbEVfNUhWbk14Zi1VeHBUN3p0ZHNIYmM5UUVINlcyYnp3eGJwalRjekVaczNKQ0NCM2MtTmV3TkhzajlQWU0zYjV0VGxUTlA5a05Bd1BaSFdwdDExdDc5THVOa05HdDlMZk9layIsInNpZ25hdHVyZSI6ImlVMGwtNWlGM25PaFZfRzRkeXZaRi0wREpUbTNhNTBkMzZDSzNBY0ZuUXdRR2FYeEhKcHNPQW9vdU02RVZtZ2dxTHl0Zk9zdkZldkY5VUw4RzhFQmZYTXg4TC04NVhIXzdpS1VwVVNPQ0JSM0Vwc2tsa0pHVWg2TVVuaDJoeVA3emppMHBucHIyRU1kbUJueWQzVXpHTkhtTzVPeHJsd2UyRmxaSkJHWUJxaUZZRDcteW1iYmhWSTFQbWpVSWU3UDJ5NFFuVUpoSVNvYkUwUi02S2R5b0phUWJxS1hKcjZHZHluM2l3OS1JeDZKLTlBWXRMa2g3TExBVGJ1dzE4MXRhbnVtM1l4M3J1bHRFY1UxUlFyM1NBQzJwdGJGSkdYcUJ3TG1YTGt1VGpNSFdmbEVZdWhyeVJaWUh6cHFGQzh3OXlrb2FoNVVOWjlRd2hFMWtRaVJ0cXF3NF9JU2tEc0tFb3hMM0xXQ2s0dEJBbEpaeE9RNzViNkNwaTZiU1Rzc2VfVWcwdldwVlVzZlJGbFpBenZtdGhCM1BXaDh3M0djeWJsR211NUZPa08wMXJ4ZGw0MHR4UUZmTEQ1dlVPbDlxaklYbGY5aGprLTcyUVpjaTFuZzQ1aS1aT2w5aUk2d3VpbHpxaXU3am0ySG1BOHVzcDFxUHV0WVFzYzFMYV9YUzFDUVM1NjVlOVI3UG5XbkNiZnMyeVFzeU0wQWdlLWI3aHNtMDJMOXdfa0hIY0RZMklreVowZDJQS19XRnpzNTc4YzQzOENTNDNZbGRkS2FQdzNieTZEaGp2cE1LWFZCSmxaRnk3dlR1NUxGczh3RFNFWVMwT2xlMmJPZEdodkVsSTFDTklxOHFKdlJkMHkyelVKNjFGTzZjVEF3SjZFVHFTMXRCYTdCSXN3IiwiZGVhZGxpbmVIZWlnaHQiOjExODMyMzQsImJsb2NrIjoxMTgzMjM0LCJ2YWxpZGF0b3JTaWduYXR1cmVzIjpbXX19LCJjaGFpblByb29mcyI6eyJ0aGlzUHVibGljYXRpb24iOnsic2lnbmF0dXJlIjoiMHgwMWViMjBiODg5MmUyNzg2NDBjOWEwY2E2ZDIzNzVmNGQwYzViYWE5MGE2OTYxNmY2NjYwNjU2MzY1MzI0M2Y2M2FkZTU2M2E0YmE2ZDg2YmQyYmZlMzI3MzYwY2Q4OGY3ODlmNDM1NjhkNTAxYjYwODYwYzNlMmI5MWNiYjZlMjFjIiwic2lnbmVkQnlEZWxlZ2F0ZSI6dHJ1ZSwic2lnbmF0dXJlRGVhZGxpbmUiOjE2ODQxNDI1MTYsInR5cGVkRGF0YSI6eyJ0eXBlcyI6eyJNaXJyb3JXaXRoU2lnIjpbeyJuYW1lIjoicHJvZmlsZUlkIiwidHlwZSI6InVpbnQyNTYifSx7Im5hbWUiOiJwcm9maWxlSWRQb2ludGVkIiwidHlwZSI6InVpbnQyNTYifSx7Im5hbWUiOiJwdWJJZFBvaW50ZWQiLCJ0eXBlIjoidWludDI1NiJ9LHsibmFtZSI6InJlZmVyZW5jZU1vZHVsZURhdGEiLCJ0eXBlIjoiYnl0ZXMifSx7Im5hbWUiOiJyZWZlcmVuY2VNb2R1bGUiLCJ0eXBlIjoiYWRkcmVzcyJ9LHsibmFtZSI6InJlZmVyZW5jZU1vZHVsZUluaXREYXRhIiwidHlwZSI6ImJ5dGVzIn0seyJuYW1lIjoibm9uY2UiLCJ0eXBlIjoidWludDI1NiJ9LHsibmFtZSI6ImRlYWRsaW5lIiwidHlwZSI6InVpbnQyNTYifV19LCJkb21haW4iOnsibmFtZSI6IkxlbnMgUHJvdG9jb2wgUHJvZmlsZXMiLCJ2ZXJzaW9uIjoiMSIsImNoYWluSWQiOjEzNywidmVyaWZ5aW5nQ29udHJhY3QiOiIweERiNDZkMURjMTU1NjM0RmJDNzMyZjkyRTg1M2IxMEIyODhBRDVhMWQifSwidmFsdWUiOnsicHJvZmlsZUlkIjoiMHgwMTVmYzMiLCJwcm9maWxlSWRQb2ludGVkIjoiMHgwNSIsInB1YklkUG9pbnRlZCI6IjB4MWQ4NSIsInJlZmVyZW5jZU1vZHVsZURhdGEiOiIweCIsInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZUluaXREYXRhIjoiMHgiLCJub25jZSI6MCwiZGVhZGxpbmUiOjE2ODQxNDI1MTZ9fSwiYmxvY2tIYXNoIjoiMHg3NDdmYTQ5NjkxYWYzNTM1OWI0YzBkNDk2YzIzNTY2YmVkNmMxYTQ2YTdiNjUyZDYyMjRjNjY2ZWJlYzI4OTRmIiwiYmxvY2tOdW1iZXIiOjQyNzI5MjU5LCJibG9ja1RpbWVzdGFtcCI6MTY4NDE0MjUxNn0sInBvaW50ZXIiOnsibG9jYXRpb24iOiJhcjovL3Fpb0ExWkVDSXgtZ2t5cWRUajk2RWtSWWdLeXRwUmhJcnhQLU5nb3ItUFkiLCJ0eXBlIjoiT05fREEifX0sInB1YmxpY2F0aW9uSWQiOiIweDAxNWZjMy0weDBjLURBLWQwM2Q1MzA5IiwiZXZlbnQiOnsicHJvZmlsZUlkIjoiMHgwMTVmYzMiLCJwdWJJZCI6IjB4MGMiLCJwcm9maWxlSWRQb2ludGVkIjoiMHgwNSIsInB1YklkUG9pbnRlZCI6IjB4MWQ4NSIsInJlZmVyZW5jZU1vZHVsZURhdGEiOiIweCIsInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZVJldHVybkRhdGEiOiIweCIsInRpbWVzdGFtcCI6MTY4NDE0MjUxNn19",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "Hx828fn_nXZM1xCST_Vejz5JSCbz0XfjhCAOiOuugiM",
				Network:  filter.NetworkArweave,
				Index:    0,
				From:     "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
				To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
				Type:     filter.TypeSocialShare,
				Platform: lo.ToPtr(filter.PlatformLens),
				Fee: &schema.Fee{
					Address: nil,
					Amount:  decimal.Zero,
					Decimal: 12,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialShare,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformLens.String(),
						From:     "0xe9c57C291340Ef34DB3646A10af99FE2A0E03827",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle:        "polebug.lens",
							ProfileID:     "0x015fc3",
							PublicationID: "0x015fc3-0x0c-DA-d03d5309",
							Timestamp:     1684142516,
							Target: &metadata.SocialPost{
								Handle:        "stani.lens",
								Title:         "How much gas prices affect your decision to mint an NFT?",
								Body:          "How much gas prices affect your decision to mint an NFT?",
								ProfileID:     "0x05",
								PublicationID: "0x05-0x1d85-DA-0ae20b9a",
								ContentURI:    "https://arweave.net/XM5x0FSJS-V6mkl1Bh6UPnKy9F2pRmM5SCqvbvWtLHE",
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1684142674,
			},
			wantError: require.NoError,
		},
		{
			name: "Momoka V1 Comment",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1692691555,
					},
					Transaction: arweave.Transaction{
						ID:       "0ZDiUg-TX9ne6Q0Bef1Pm1-q2giPUZIdDFDNa4YjEbg",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweGM2ODk0MTY4OGE2ZjgyNGFjZGI3MjExNjI5OTUzMzQxZWI3ZDYzYTcyYjQ4YmZmMWMxMmJhY2ViNWMzOWQxYmM2ZGNlNThmMjA2MzhkMTg5NjVkYzgxYTNjOGNkZDE2NjM0N2M2MmM0OTg4NWNiMjU2MWY0ZTE0ZjQwY2YwOGY3MWMiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiIzZjI1NjM3Ny02OGQyLTQ5OGItOWY3NC03MjU4N2I1NmNkMDYiLCJ0eXBlIjoiQ09NTUVOVF9DUkVBVEVEIiwidGltZXN0YW1wUHJvb2ZzIjp7InR5cGUiOiJCVU5ETFIiLCJoYXNoUHJlZml4IjoiMSIsInJlc3BvbnNlIjp7ImlkIjoiN0toNFYwZTdtdGtaUW9uU0szcXRQY3NhMVkxSHBsTTZNREJINW1rQ3E0VSIsInRpbWVzdGFtcCI6MTY5MjY5MDMxMDk5OSwidmVyc2lvbiI6IjEuMC4wIiwicHVibGljIjoic3E5SmJwcEtMbEFLdFF3YWxmWDVEYWduR01sVGlyZGl0WGs3eTRqZ29lQTdERU0wWjZjVlBFNXhNUTlrel9UOVZwcFA2QkZIdEh5WkNaT0RlcmNFVldpcHprcjM2dGZRa1I1RURHVVF5TGl2ZHhVemJXZ1Zrenc3RDI3UEpFYTRjZDFVeTZyMThyWUxxRVJnYlJ2QVpwaDVZSlptcFNKazdyM013blFxdXVrdGp2U3BmQ0xGd1N4UDF3ODc5LXNzX0phbE05SUN6UmkzOGhlbk9OaW84Z2xsNkdWOS1vbXJXd1JNWmVyXzE1YnNwQ0s1dHhDd3BZMTM3bmZLd0tENVlCQXV6eHhjajQyNE03emxTSGxzYWZCd2FSd0ZiZjhnSHRXMDNpSkVSNGxSNEd4ZVkwV3ZuWWFCM0tESVNIUXA1M2E5bmxibWlXTzVXY0hIWXNSODNPVDJlSjBQbDNSV0EtX2lta19TTndHUVRDam1BNnRmX1VWd0w4SHpZUzJpeXV1ODViN2lZSzlaUW9oOG5xYk5DNnFpYklDRTRoOUZlM2JON0FnaXRJZTlYekNUT1hEZk1yNGFoakM4a2txSjF6NHpOQUk2LUxlZWlfTWdkOEp0WmgydnFGTlpoWEswbFNhZEZsXzlPaDNBRVQ3dFVkczJFN3MtNnpwUlBkOW9CWnU2LWtOdUhEUko2VFFoWlN3SjlaTzVIWXNjY2JfR18xc283MmFYSnltUjlnZ0pnV3I0SjNiYXdBWVlucW12bXpHa2xZT2xFXzVIVm5NeGYtVXhwVDd6dGRzSGJjOVFFSDZXMmJ6d3hicGpUY3pFWnMzSkNDQjNjLU5ld05Ic2o5UFlNM2I1dFRsVE5QOWtOQXdQWkhXcHQxMXQ3OUx1TmtOR3Q5TGZPZWsiLCJzaWduYXR1cmUiOiJSVXZTcGlEOTE4aXhwa2dKNkNfM083Wkd6eTNzM1BIWEJLLTBxTXYtRzQtclpPeTRPdzBaS0Y4SjJHZmUyRnBUQTR6Vk5tUHN5b3VaZzNxOXBtZEEyQ0RWV3hvclcxRl9WUXF1RmVfczV2V0xzRmR2NlphcVJxanNOT2tLejVWZ00xcnBSQWlTQXdUUDVjTnVOTW5RVXBTVWc3eXVmSWtYX1JtQUd6VVVaTWtSOEs0cW9MRzI4RndwSWFZb0JURTFuOGJaVUZRVGptaFJmOGZOTFhEYk9za0lfZVdfSzZqci1hZmVDaVkxYm1kd2pUNWxLVnJ1VVpmaXlTWnYzQ0p6U2NLeDgtbDQwekxpeV9hVTdwSjVtVGxKMXROR0ljc0poVl9zLW95T2JPUzVZOHVHTXZhMEpiWGNSakNiSDJCQ1o4WnFFUDJOSVFOS3lXWmdxNHNrdy0tTk1ROU9OVHd5ejdIU195OWZHZDB3U1lNWVB3Qy0wcEZwSWNSSHN5c184Qm9YUzZ2TEFlTTdzR3YtVHd2ajc5bkZDZ2x2dlBqamV0SzhUQTRvQWNDSnA2ZmsxaGQwdndXUzE1aDllc29WY3hyN0xMekdyNDVEMUpVQ0NOd1lOeWdhM0RBZ1pBSzlQU2p2clB5QWl5dHJTSzBvOXdvMHFKa1p4OVZuclNid2ZESTBXblZzT2JoTi1tZTk0ckJpdU9UalEtZ0g2X2J6Q3g0YzBiUUdkZzE3MlFWdFc4VUNMLW1MWVdCQW9MY0l1TXlzSEJYX0w0V0F2amRrSFVPcjdaYmlYNWxUcnhKYnBHaDA1dENHZGRVWVZVR0w2eTQ1T1R3LXdjMENpRVkxa3VFQUZncFNmVE1DWmdNcmdrWU1XN194LWd0b1VXYVFNdmgwZTlETEE4TSIsImRlYWRsaW5lSGVpZ2h0IjoxMjQ5NzMzLCJibG9jayI6MTI0OTczMywidmFsaWRhdG9yU2lnbmF0dXJlcyI6W119fSwiY2hhaW5Qcm9vZnMiOnsidGhpc1B1YmxpY2F0aW9uIjp7InNpZ25hdHVyZSI6IjB4ZDViYWY4ODEyMWViNzBmMDUxMDlhNDlhYzI2Mzg0NTdjZDdiZmU5NDk5MTQ1ODUwY2M3MTIxNWI2OThkNGM4OTQ2MmQ5ODgwODhhYzg0NTY3NmRjNDZiOGM3NmU1MzhjZTY5NTViOWQ2ZmJjN2M1MWI5M2UwMjViMGYxMzYzMTcxYyIsInNpZ25lZEJ5RGVsZWdhdGUiOnRydWUsInNpZ25hdHVyZURlYWRsaW5lIjoxNjkyNjkwMzA5LCJ0eXBlZERhdGEiOnsidHlwZXMiOnsiQ29tbWVudFdpdGhTaWciOlt7Im5hbWUiOiJwcm9maWxlSWQiLCJ0eXBlIjoidWludDI1NiJ9LHsibmFtZSI6ImNvbnRlbnRVUkkiLCJ0eXBlIjoic3RyaW5nIn0seyJuYW1lIjoicHJvZmlsZUlkUG9pbnRlZCIsInR5cGUiOiJ1aW50MjU2In0seyJuYW1lIjoicHViSWRQb2ludGVkIiwidHlwZSI6InVpbnQyNTYifSx7Im5hbWUiOiJyZWZlcmVuY2VNb2R1bGVEYXRhIiwidHlwZSI6ImJ5dGVzIn0seyJuYW1lIjoiY29sbGVjdE1vZHVsZSIsInR5cGUiOiJhZGRyZXNzIn0seyJuYW1lIjoiY29sbGVjdE1vZHVsZUluaXREYXRhIiwidHlwZSI6ImJ5dGVzIn0seyJuYW1lIjoicmVmZXJlbmNlTW9kdWxlIiwidHlwZSI6ImFkZHJlc3MifSx7Im5hbWUiOiJyZWZlcmVuY2VNb2R1bGVJbml0RGF0YSIsInR5cGUiOiJieXRlcyJ9LHsibmFtZSI6Im5vbmNlIiwidHlwZSI6InVpbnQyNTYifSx7Im5hbWUiOiJkZWFkbGluZSIsInR5cGUiOiJ1aW50MjU2In1dfSwiZG9tYWluIjp7Im5hbWUiOiJMZW5zIFByb3RvY29sIFByb2ZpbGVzIiwidmVyc2lvbiI6IjEiLCJjaGFpbklkIjoxMzcsInZlcmlmeWluZ0NvbnRyYWN0IjoiMHhEYjQ2ZDFEYzE1NTYzNEZiQzczMmY5MkU4NTNiMTBCMjg4QUQ1YTFkIn0sInZhbHVlIjp7InByb2ZpbGVJZCI6IjB4MDFkMWFjIiwicHJvZmlsZUlkUG9pbnRlZCI6IjB4MDFkMTliIiwicHViSWRQb2ludGVkIjoiMHgyMSIsImNvbnRlbnRVUkkiOiJhcjovLzc4RlhmVHh3NjllaDZ3T0t2UkNaRkpRUjVBYWRadEJjYklwZ3ZHZUtZMEUiLCJyZWZlcmVuY2VNb2R1bGUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAiLCJjb2xsZWN0TW9kdWxlIjoiMHhhMzFGRjg1RTg0MEVEMTE3RTE3MkJDOUFkODlFNTUxMjhBOTk5MjA1IiwiY29sbGVjdE1vZHVsZUluaXREYXRhIjoiMHgiLCJyZWZlcmVuY2VNb2R1bGVJbml0RGF0YSI6IjB4IiwicmVmZXJlbmNlTW9kdWxlRGF0YSI6IjB4Iiwibm9uY2UiOjAsImRlYWRsaW5lIjoxNjkyNjkwMzA5fX0sImJsb2NrSGFzaCI6IjB4NjUzOWY1NmZiOTdiNmQxZmM1YWNmOGMxY2YxMTBiMzNjOTllM2FjNDExOGQ5ODgwNDI0NWY0NmIzYWM0MWE2ZiIsImJsb2NrTnVtYmVyIjo0NjYwOTMwNCwiYmxvY2tUaW1lc3RhbXAiOjE2OTI2OTAzMDl9LCJwb2ludGVyIjp7ImxvY2F0aW9uIjoiYXI6Ly9fdU1DNUV1X3YtMjRuZ1A2c3RuOGxKVENaQXktbEJNMmxxdHpQYTBtVk93IiwidHlwZSI6Ik9OX0RBIn19LCJwdWJsaWNhdGlvbklkIjoiMHgwMWQxYWMtMHgwMi1EQS0zZjI1NjM3NyIsImV2ZW50Ijp7InByb2ZpbGVJZCI6IjB4MDFkMWFjIiwicHViSWQiOiIweDAyIiwiY29udGVudFVSSSI6ImFyOi8vNzhGWGZUeHc2OWVoNndPS3ZSQ1pGSlFSNUFhZFp0QmNiSXBndkdlS1kwRSIsInByb2ZpbGVJZFBvaW50ZWQiOiIweDAxZDE5YiIsInB1YklkUG9pbnRlZCI6IjB4MjEiLCJyZWZlcmVuY2VNb2R1bGVEYXRhIjoiMHgiLCJjb2xsZWN0TW9kdWxlIjoiMHhhMzFGRjg1RTg0MEVEMTE3RTE3MkJDOUFkODlFNTUxMjhBOTk5MjA1IiwiY29sbGVjdE1vZHVsZVJldHVybkRhdGEiOiIweCIsInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZVJldHVybkRhdGEiOiIweCIsInRpbWVzdGFtcCI6MTY5MjY5MDMwOX19",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "0ZDiUg-TX9ne6Q0Bef1Pm1-q2giPUZIdDFDNa4YjEbg",
				Network:  filter.NetworkArweave,
				Index:    0,
				From:     "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
				To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
				Type:     filter.TypeSocialComment,
				Platform: lo.ToPtr(filter.PlatformLens),
				Fee: &schema.Fee{
					Address: nil,
					Amount:  decimal.Zero,
					Decimal: 12,
				},
				Actions: []*schema.Action{
					{
						Type:     filter.TypeSocialComment,
						Tag:      filter.TagSocial,
						Platform: filter.PlatformLens.String(),
						From:     "0xE6b532E63F228087e26a5897131f2e1D043e27f2",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle:        "hypnogaba.lens",
							Title:         "Comment by @hypnogaba.lens",
							Body:          "im so hyped and waiting for them) ",
							ProfileID:     "0x01d1ac",
							PublicationID: "0x01d1ac-0x02-DA-3f256377",
							ContentURI:    "ar://78FXfTxw69eh6wOKvRCZFJQR5AadZtBcbIpgvGeKY0E",
							Tags:          []string{"music"},
							Timestamp:     1692690309,
							Target: &metadata.SocialPost{
								Handle: "agoria.lens",
								Title:  "Video by @agoria.lens",
								Body:   "Let's welcome (y)our 2000 Dynamic Avatars!\nAllow List Opens today.\n\n{Agorians} are the first ever web3 avatar collection that implements an automatic appearance transformation to reflect the real-life time of day... Are they virtual?!",
								Media: []metadata.Media{
									{
										Address:  "ipfs://bafybeihkxh7ndcdnn6kqnmn5l3a42vjlp7vfec7ybrpvcymatof2elrf4i",
										MimeType: "video/mp4",
									},
								},
								ProfileID:     "0x01d19b",
								PublicationID: "0x01d19b-0x21-DA-af626ace",
								ContentURI:    "ar://uT1Ph_o95nleQ-Yu_mHBbFHt-l3sSdJC9iB0zLO71Xg",
								Tags:          []string{"science_&_technology"},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1692691555,
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
