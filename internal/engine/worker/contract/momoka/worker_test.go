package momoka_test

import (
	"context"
	"testing"

	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/arweave"
	worker "github.com/rss3-network/node/internal/engine/worker/contract/momoka"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
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
		// V1
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

		// V2
		{
			name: "Momoka V2 Post",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1699526849,
					},
					Transaction: arweave.Transaction{
						ID:       "XL9qZdAKrBQleCiaVOxy2veWWoUB2mGeTH2D__Ksx6g",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweDhhNTFhNDdhMzgxMmE2MjkxZTNmMTk3MWY5NjRiZThiMzE4MDZiMjgxMTE5ZDNlYzEwNDhlYTU4NjZjNmUwZmUzNWVmZDFjZDhmMGI0YTNhMDlmMzAxNmNhNmNkMTU5OGM0NzI2N2E4ZTE1ZmRlYTIwMzVjYzI1OTI3Mzk3MDlkMWIiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiJhNGIzMGZiZS1hNjNiLTQ4ZWItYWZmZC02ZGE3NTAyYjUyMGMiLCJ0eXBlIjoiUE9TVF9DUkVBVEVEIiwidGltZXN0YW1wUHJvb2ZzIjp7InR5cGUiOiJCVU5ETFIiLCJoYXNoUHJlZml4IjoiMSIsInJlc3BvbnNlIjp7ImlkIjoiZFFtd25FaXI0Z3hPdXJ1dDMwQk9XMkRtMDZIbk5MYTRjZHV3MHZJZjI5dyIsInRpbWVzdGFtcCI6MTY5OTUyNjAyNTM1NiwidmVyc2lvbiI6IjEuMC4wIiwicHVibGljIjoic3E5SmJwcEtMbEFLdFF3YWxmWDVEYWduR01sVGlyZGl0WGs3eTRqZ29lQTdERU0wWjZjVlBFNXhNUTlrel9UOVZwcFA2QkZIdEh5WkNaT0RlcmNFVldpcHprcjM2dGZRa1I1RURHVVF5TGl2ZHhVemJXZ1Zrenc3RDI3UEpFYTRjZDFVeTZyMThyWUxxRVJnYlJ2QVpwaDVZSlptcFNKazdyM013blFxdXVrdGp2U3BmQ0xGd1N4UDF3ODc5LXNzX0phbE05SUN6UmkzOGhlbk9OaW84Z2xsNkdWOS1vbXJXd1JNWmVyXzE1YnNwQ0s1dHhDd3BZMTM3bmZLd0tENVlCQXV6eHhjajQyNE03emxTSGxzYWZCd2FSd0ZiZjhnSHRXMDNpSkVSNGxSNEd4ZVkwV3ZuWWFCM0tESVNIUXA1M2E5bmxibWlXTzVXY0hIWXNSODNPVDJlSjBQbDNSV0EtX2lta19TTndHUVRDam1BNnRmX1VWd0w4SHpZUzJpeXV1ODViN2lZSzlaUW9oOG5xYk5DNnFpYklDRTRoOUZlM2JON0FnaXRJZTlYekNUT1hEZk1yNGFoakM4a2txSjF6NHpOQUk2LUxlZWlfTWdkOEp0WmgydnFGTlpoWEswbFNhZEZsXzlPaDNBRVQ3dFVkczJFN3MtNnpwUlBkOW9CWnU2LWtOdUhEUko2VFFoWlN3SjlaTzVIWXNjY2JfR18xc283MmFYSnltUjlnZ0pnV3I0SjNiYXdBWVlucW12bXpHa2xZT2xFXzVIVm5NeGYtVXhwVDd6dGRzSGJjOVFFSDZXMmJ6d3hicGpUY3pFWnMzSkNDQjNjLU5ld05Ic2o5UFlNM2I1dFRsVE5QOWtOQXdQWkhXcHQxMXQ3OUx1TmtOR3Q5TGZPZWsiLCJzaWduYXR1cmUiOiJsel9DU3BvSUxBeFlmLVZBWW1vS2YwdWptTjJvdGlCcW9MZUJwNVVvVjFydkRPMmJlREVfS3Z1bzRGYVViZ3R5eHNac1FpNnJpSkUySlNNQUo4WHBCUDRtOW9LNE92ZTlIX3o5VzRQZmU0V0JTbzNhcmxYWXY0d2pnMUx0RUx6NUlZcVA4Vk95YUc2akEyZFJjMjdpcEhjcm9hYXctcFd6eHd2ZEJCY1lKTjJoRzZ2cmpKbENfYWJ3Nlo5TDNIQTY0dUZ0TXhnd2h2VXJob1dVQkozWF9QOHViYlI1bUx0ZUo4cl9xTVZicGNDR2Q5ZUtkNllnd3ozX2RseUZOOUFESGhyOU1UbVI2WXdWYVcycHJWdmluN0N2OGQ4ZV9UMXlaX01lSnduVG4wWUMyaExxTXVsdkhCUUFFZjU4M1FWZnJRMWd3amtscXZyRHNtRGxEUnVZbndYSlNIQ0J5ZWtKd19ORnBSdDZlcTNZamFpd3pEb3J3NVBtRjBsMEhERkpJYUtWZ21VVUlRMEVDd05JWjh1Tjh2eUF5RmxBdkdkbEpnNzJxYS1VV2dnUnRvV1NVQ2dCTXpuVTlPSWtUdnZuajRtQURCZExheVowRFBPSG82SDNKVjBIOWRKYVMtYlRBQmNFLXgzM0NvMGw5NmNPbFJGbDhTNnNNajdSQ1J0WUJmVWY5aTZVcXdVbTZadFMyZGlGSlFQWGZ0bmVFOG5ud2M1b0x3bWt2NmVYaVBTS1poMWljQnFHOUxVdlViTkJ3bnVpVVVTMC0zMF9IaXRYWUxlN29NdTkyN2w4VXFTRG1ZaVE3QnYyX1dlNGp4a1E2Ty0yWnVhbklrdXFfWXFoek1STVNwVmtJNnFsZWlmVUdDMzlWOEFlak9teEVqWVVTbEg5UnJxOGw5cyIsImRlYWRsaW5lSGVpZ2h0IjoxMzAyODEwLCJibG9jayI6MTMwMjgxMCwidmFsaWRhdG9yU2lnbmF0dXJlcyI6W119fSwiY2hhaW5Qcm9vZnMiOnsidGhpc1B1YmxpY2F0aW9uIjp7InNpZ25hdHVyZSI6IjB4MWE4YzgxYTgzY2QyOGQwYWM1NTlkMWE1NTYxZTkzODQyMDJiMDIxMGY5ZTQzYjIzMjAxMmYwMWYwYzYyYTc0MzRkNTk4YzcyMjg3YWRjMzcyMzMyNjQyMTkxMDUzZTFiZDc2MjEyOWIyZGQyMDQyZjc1MjdkNzQ5ZTRmODgxNGUxYyIsInNpZ25lZEJ5RGVsZWdhdGUiOnRydWUsInNpZ25hdHVyZURlYWRsaW5lIjoxNjk5NTI2MDI0LCJ0eXBlZERhdGEiOnsidHlwZXMiOnsiUG9zdCI6W3sidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoicHJvZmlsZUlkIn0seyJ0eXBlIjoic3RyaW5nIiwibmFtZSI6ImNvbnRlbnRVUkkifSx7InR5cGUiOiJhZGRyZXNzW10iLCJuYW1lIjoiYWN0aW9uTW9kdWxlcyJ9LHsidHlwZSI6ImJ5dGVzW10iLCJuYW1lIjoiYWN0aW9uTW9kdWxlc0luaXREYXRhcyJ9LHsidHlwZSI6ImFkZHJlc3MiLCJuYW1lIjoicmVmZXJlbmNlTW9kdWxlIn0seyJ0eXBlIjoiYnl0ZXMiLCJuYW1lIjoicmVmZXJlbmNlTW9kdWxlSW5pdERhdGEifSx7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6Im5vbmNlIn0seyJ0eXBlIjoidWludDI1NiIsIm5hbWUiOiJkZWFkbGluZSJ9XX0sImRvbWFpbiI6eyJuYW1lIjoiTGVucyBQcm90b2NvbCBQcm9maWxlcyIsInZlcnNpb24iOiIyIiwiY2hhaW5JZCI6MTM3LCJ2ZXJpZnlpbmdDb250cmFjdCI6IjB4RGI0NmQxRGMxNTU2MzRGYkM3MzJmOTJFODUzYjEwQjI4OEFENWExZCJ9LCJ2YWx1ZSI6eyJwcm9maWxlSWQiOiIweDAxIiwiY29udGVudFVSSSI6ImFyOi8vZ2J0UlR1aDA4UEljMHpSd21kUE1nNElxb3BkQmwwSDVRWW10TUctNVMtMCIsImFjdGlvbk1vZHVsZXMiOltdLCJhY3Rpb25Nb2R1bGVzSW5pdERhdGFzIjpbXSwicmVmZXJlbmNlTW9kdWxlIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwicmVmZXJlbmNlTW9kdWxlSW5pdERhdGEiOiIweCIsIm5vbmNlIjowLCJkZWFkbGluZSI6MTY5OTUyNjAyNH19LCJibG9ja0hhc2giOiIweGRkNGEyOGI4NzRmZjgyMmUxYjYzZDIwNDEyNGU0YWUwMzcwNDJiMjQxYTllNzc0NzI1ZDU0ZjYyMWFlMDBlZjUiLCJibG9ja051bWJlciI6NDk3Mjk4MDgsImJsb2NrVGltZXN0YW1wIjoxNjk5NTI2MDI0fSwicG9pbnRlciI6bnVsbH0sInB1YmxpY2F0aW9uSWQiOiIweDAxLTB4MDJhMi1EQS1hNGIzMGZiZSIsImV2ZW50Ijp7InBvc3RQYXJhbXMiOnsicHJvZmlsZUlkIjoiMHgwMSIsImNvbnRlbnRVUkkiOiJhcjovL2didFJUdWgwOFBJYzB6UndtZFBNZzRJcW9wZEJsMEg1UVltdE1HLTVTLTAiLCJhY3Rpb25Nb2R1bGVzIjpbXSwiYWN0aW9uTW9kdWxlc0luaXREYXRhcyI6W10sInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZUluaXREYXRhIjoiMHgifSwicHViSWQiOiIweDAyYTIiLCJhY3Rpb25Nb2R1bGVzSW5pdFJldHVybkRhdGFzIjpbXSwicmVmZXJlbmNlTW9kdWxlSW5pdFJldHVybkRhdGEiOiIweCIsInRyYW5zYWN0aW9uRXhlY3V0b3IiOiIweDA1MDkyY0Y2OUJERDQzNWY3QmE0QjhlRjk3YzlDQWVjRjJCQTY5QUQiLCJ0aW1lc3RhbXAiOjE2OTk1MjYwMjR9fQ",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "XL9qZdAKrBQleCiaVOxy2veWWoUB2mGeTH2D__Ksx6g",
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
						From:     "0x05092cF69BDD435f7Ba4B8eF97c9CAecF2BA69AD",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle: "lensprotocol.lens",
							Title:  "Post by @lensprotocol",
							Body:   "gm @lens/shefi, we can't wait to see you in İstanbul this Sunday.\u00a0\n\nAll attendees of the SheFi Summit are in for an exclusive perk. Access to @lens/letsraave awaits you 🪩",
							Media: []metadata.Media{
								{
									Address:  "ipfs://bafybeifqioqsxyvvu67luvvmj5hcenl773yi2pn46l2o4wkdgkethn23ua",
									MimeType: "image/png",
								},
							},
							ProfileID:     "0x01",
							PublicationID: "0x01-0x02a2-DA-a4b30fbe",
							ContentURI:    "ar://gbtRTuh08PIc0zRwmdPMg4IqopdBl0H5QYmtMG-5S-0",
							Timestamp:     1699526024,
						},
					},
				},
				Status:    true,
				Timestamp: 1699526849,
			},
			wantError: require.NoError,
		},
		{
			name: "Momoka V2 Mirror",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1699512726,
					},
					Transaction: arweave.Transaction{
						ID:       "TwtfazAvUhHOMDt0nRsuTVB_HywfrtJ3IRUX0l2ucns",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweGI1ZWQ4ZDUzMGI0ZWJlNDY0OGYzMTVmYjc1Njk0YTVmMzcxOGFiOGFlNzBkYjg5MDhkMmZjNzFkYTRlMzc1YjU2MmQ1ZDAzMGVjMjcwYmEwZWUxMjhmMmIwMjFjZTQ0NTFhMzZmODg1NzYwZmFjOGJiMjlhY2I3MDQ0NTk2MTU2MWIiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiI3OTBmOWUyMC1lMzM1LTQxNDctYjgyZi1mNTE0NmFhNDhhNjMiLCJ0eXBlIjoiTUlSUk9SX0NSRUFURUQiLCJ0aW1lc3RhbXBQcm9vZnMiOnsidHlwZSI6IkJVTkRMUiIsImhhc2hQcmVmaXgiOiIxIiwicmVzcG9uc2UiOnsiaWQiOiJpTF9sRU10V0xDUS14a1RRei04cXBIZmp4QUppSnRJWFNFbVpKcFBMMkVjIiwidGltZXN0YW1wIjoxNjk5NTEyNTk3OTMwLCJ2ZXJzaW9uIjoiMS4wLjAiLCJwdWJsaWMiOiJzcTlKYnBwS0xsQUt0UXdhbGZYNURhZ25HTWxUaXJkaXRYazd5NGpnb2VBN0RFTTBaNmNWUEU1eE1ROWt6X1Q5VnBwUDZCRkh0SHlaQ1pPRGVyY0VWV2lwemtyMzZ0ZlFrUjVFREdVUXlMaXZkeFV6YldnVmt6dzdEMjdQSkVhNGNkMVV5NnIxOHJZTHFFUmdiUnZBWnBoNVlKWm1wU0prN3IzTXduUXF1dWt0anZTcGZDTEZ3U3hQMXc4Nzktc3NfSmFsTTlJQ3pSaTM4aGVuT05pbzhnbGw2R1Y5LW9tcld3Uk1aZXJfMTVic3BDSzV0eEN3cFkxMzduZkt3S0Q1WUJBdXp4eGNqNDI0TTd6bFNIbHNhZkJ3YVJ3RmJmOGdIdFcwM2lKRVI0bFI0R3hlWTBXdm5ZYUIzS0RJU0hRcDUzYTlubGJtaVdPNVdjSEhZc1I4M09UMmVKMFBsM1JXQS1faW1rX1NOd0dRVENqbUE2dGZfVVZ3TDhIellTMml5dXU4NWI3aVlLOVpRb2g4bnFiTkM2cWliSUNFNGg5RmUzYk43QWdpdEllOVh6Q1RPWERmTXI0YWhqQzhra3FKMXo0ek5BSTYtTGVlaV9NZ2Q4SnRaaDJ2cUZOWmhYSzBsU2FkRmxfOU9oM0FFVDd0VWRzMkU3cy02enBSUGQ5b0JadTYta051SERSSjZUUWhaU3dKOVpPNUhZc2NjYl9HXzFzbzcyYVhKeW1SOWdnSmdXcjRKM2Jhd0FZWW5xbXZtekdrbFlPbEVfNUhWbk14Zi1VeHBUN3p0ZHNIYmM5UUVINlcyYnp3eGJwalRjekVaczNKQ0NCM2MtTmV3TkhzajlQWU0zYjV0VGxUTlA5a05Bd1BaSFdwdDExdDc5THVOa05HdDlMZk9layIsInNpZ25hdHVyZSI6IkxGaUVKT0MtZy1xWmQ0cTh5UWtpb3FJQllJOUFKaUwxakRpZ0RkLUNCU3FhYXJlMUJxRDctRnF3Y04yS0hVR2I3d05kZXBmLVhINHZLWkpFVml2NVZGRG8yZ3R6SjNibmZEcEU5VUVNVW9acDNPMlBqM0d0VnQxOWoxTEpQM3ExOWlQLWhYMkgxOWI0bEdKSmRnVlNDWi1Oam9udHM1LTZhNU1FNUtWZjIzaUM0TUtzZzMtZ2pNYl9mRElKbzdTZlRiY2RlanpIaDM2Mm5mZU9LMnNqNFZSdEZSbDVIMHVydG1hUWNLOVFYVnZRcDBnbzBfUmQ4bHc3SE5rdEZsaGlScTdBS3Z4eUI4UnpIYVVIdjg3Y2hnMmJTSlR4X2FjTVktcXpaQTBPSXFSSEZXQTFHeUpUN2RuR2p3Z1JzRGFjZXRPNnNZb0ZDS1FJekF6UW13eXFNWUNZZEk2MXNUNHdQNFB0b1pES3QxVTctbkN0WTNWQVRacEZua1JVdW8xbmVWSXFKV01EWWllWjBMMUwwS0hxTDdtSUZmeC1UdDVLWGk0Q3NZY1MyVWN1MU1BS083MjVkeWRDZ3ktbUljRVYtcjBnd0RjZWtQZWZTbmJnRE5TNDY1cUNxOHY0LWRyNkU0OUYtblZSUU1YQnBaYWpEQlg5Rkg3RmtJZHhacFhZSTNfTk9IQnBmZDVmSE16X3Nic3dJd0FoNWoxLUUxUno4MGVWc2tCM0pmTnNydldVR0NPXzlSanZqRE5YM0d1MENJd3pwOURHX1lfUTZsd3k5YUJYZjR3anphU19rcDBfRm5FeDdXSlFrcWM3VDBJTDFBVG1jZUVBNlB1WXRYYnRTN0xqVnBvSVNzQlJYLWRXSDJSamVDeEN2ejA1dHhYbUFhbUhQM1JwRmtVIiwiZGVhZGxpbmVIZWlnaHQiOjEzMDI3MDMsImJsb2NrIjoxMzAyNzAzLCJ2YWxpZGF0b3JTaWduYXR1cmVzIjpbXX19LCJjaGFpblByb29mcyI6eyJ0aGlzUHVibGljYXRpb24iOnsic2lnbmF0dXJlIjoiMHgwMjAyM2RmZTBiNWY5NGI0NjljOGZlYTRkMzQ0N2VhZTU4OGQ0Njk2MWIwYTMyYjQ1MzgxNDg0NWVhNjQ2ZDBjNTVkN2MxOWM3YjVlODhmNjlkNGM4OTExNTg0OTY5MTA0N2I5NmVkMmI1YzMyYzk5NDU4ZTRkZWM0MTI1NzY0MzFjIiwic2lnbmVkQnlEZWxlZ2F0ZSI6dHJ1ZSwic2lnbmF0dXJlRGVhZGxpbmUiOjE2OTk1MTI1OTYsInR5cGVkRGF0YSI6eyJ0eXBlcyI6eyJNaXJyb3IiOlt7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6InByb2ZpbGVJZCJ9LHsidHlwZSI6InN0cmluZyIsIm5hbWUiOiJtZXRhZGF0YVVSSSJ9LHsidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoicG9pbnRlZFByb2ZpbGVJZCJ9LHsidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoicG9pbnRlZFB1YklkIn0seyJ0eXBlIjoidWludDI1NltdIiwibmFtZSI6InJlZmVycmVyUHJvZmlsZUlkcyJ9LHsidHlwZSI6InVpbnQyNTZbXSIsIm5hbWUiOiJyZWZlcnJlclB1YklkcyJ9LHsidHlwZSI6ImJ5dGVzIiwibmFtZSI6InJlZmVyZW5jZU1vZHVsZURhdGEifSx7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6Im5vbmNlIn0seyJ0eXBlIjoidWludDI1NiIsIm5hbWUiOiJkZWFkbGluZSJ9XX0sImRvbWFpbiI6eyJuYW1lIjoiTGVucyBQcm90b2NvbCBQcm9maWxlcyIsInZlcnNpb24iOiIyIiwiY2hhaW5JZCI6MTM3LCJ2ZXJpZnlpbmdDb250cmFjdCI6IjB4RGI0NmQxRGMxNTU2MzRGYkM3MzJmOTJFODUzYjEwQjI4OEFENWExZCJ9LCJ2YWx1ZSI6eyJkZWFkbGluZSI6MTY5OTUxMjU5NiwibWV0YWRhdGFVUkkiOiIiLCJub25jZSI6MCwicG9pbnRlZFByb2ZpbGVJZCI6IjB4ZWNhNyIsInBvaW50ZWRQdWJJZCI6IjB4MDMwYSIsInByb2ZpbGVJZCI6IjB4MjA2MSIsInJlZmVyZW5jZU1vZHVsZURhdGEiOiIweCIsInJlZmVycmVyUHJvZmlsZUlkcyI6W10sInJlZmVycmVyUHViSWRzIjpbXX19LCJibG9ja0hhc2giOiIweDkwZWIyOTdiZTA1MmFjNGE2YThhYTliYjU4MzdlNzliZTY4YWMwYmU2N2I2Nzc0OGNiN2UwN2VlNzcxM2JkN2UiLCJibG9ja051bWJlciI6NDk3MjM3MDMsImJsb2NrVGltZXN0YW1wIjoxNjk5NTEyNTk2fSwicG9pbnRlciI6eyJsb2NhdGlvbiI6ImFyOi8vR3JhWWtTV19fSW9Pa2otck90OWFZcERzSThsdHd0ZFk4ekwwSV9oamlRUSIsInR5cGUiOiJPTl9EQSJ9fSwicHVibGljYXRpb25JZCI6IjB4MjA2MS0weDAxODEtREEtNzkwZjllMjAiLCJldmVudCI6eyJtaXJyb3JQYXJhbXMiOnsicHJvZmlsZUlkIjoiMHgyMDYxIiwibWV0YWRhdGFVUkkiOiIiLCJyZWZlcmVuY2VNb2R1bGVEYXRhIjoiMHgiLCJyZWZlcnJlclByb2ZpbGVJZHMiOltdLCJyZWZlcnJlclB1YklkcyI6W10sInBvaW50ZWRQcm9maWxlSWQiOiIweGVjYTciLCJwb2ludGVkUHViSWQiOiIweDAzMGEifSwicHViSWQiOiIweDAxODEiLCJyZWZlcmVuY2VNb2R1bGVSZXR1cm5EYXRhIjoiMHgiLCJ0cmFuc2FjdGlvbkV4ZWN1dG9yIjoiMHg1QkY2NGY5RUIzQkU3YjNGQkRiNEJCZjYzNTRhRDUyMkI1QTI4Zjk0IiwidGltZXN0YW1wIjoxNjk5NTEyNTk2fX0",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "TwtfazAvUhHOMDt0nRsuTVB_HywfrtJ3IRUX0l2ucns",
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
						From:     "0x5BF64f9EB3BE7b3FBDb4BBf6354aD522B5A28f94",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle:        "prayag.lens",
							ProfileID:     "0x2061",
							PublicationID: "0x2061-0x0181-DA-790f9e20",
							Timestamp:     1699512596,
							Target: &metadata.SocialPost{
								Handle: "ethdaily.lens",
								Title:  "zkWasm L2 Prover On Polygon CDK",
								Body:   "Ethereum News Roundup for Wednesday, November 8, 2023: Polygon Labs announces the development of zkWasm. Optimism launches an art contest. Caldera supports custom gas tokens. And Stackr Labs raises a $5.5 million seed round.\n\nRead more: https://ethdaily.io/zkwasm-l2-prover-on-polygon-cdk/\n\n0:00 Intro\n0:32 zkWasm L2 Prover On Polygon CDK\n1:13 Optimism Launches An Art Contest\n2:02 Caldera Supports Custom Gas Tokens\n2:35 Stackr Labs Raises $5.5 Million Seed\n3:22 Fin\n\n🔗 Website - https://ethdaily.io\n🎙️ Podcast - https://link.chtbl.com/eth\n🪞 Mirror - https://ethdaily.mirror.xyz\n🎥 Video - https://youtube.com/@ethdaily\n🌿 Lens - https://lenster.xyz/u/ethdaily\n🐦 Twitter - https://twitter.com/ethdaily\n🟪 Farcaster - https://warpcast.com/ethdaily.eth\n\n#Polygon #zkWasm #ZK",
								Media: []metadata.Media{
									{
										Address:  "ipfs://bafybeianvdlfrsxlywymsynfkmp7n5r7ozltxfpbugat3hzbaqgfbiu22u",
										MimeType: "video/mp4",
									},
								},
								ProfileID:     "0xeca7",
								PublicationID: "0xeca7-0x030a-DA-2eef483d",
								ContentURI:    "ar://RKaCRNfg-Ni9pN5cYKLdeyF0W82bMgC6g3ukWAX6fC0",
								Tags:          []string{"people"},
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1699512726,
			},
			wantError: require.NoError,
		},
		{
			name: "Momoka V2 Comment",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1699512726,
					},
					Transaction: arweave.Transaction{
						ID:       "bN4Sf7lUO4TZwpPjtP3BUzBSG92io0McNgMXTxXUOcM",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweDZlMWJlOTkyMDU1YzNlYjgzYTM0M2NhOTMwYWFlNDFkN2M4ZGEzYjdhYWUyMTBmNWU1NzMwNzYyZGFlMTRhMDQzNDY0MTY5MGY5Mjk4MzEzYzNjMTI2Nzk0NWIyZDg1YWQ2OWE1MWZkYTQxZjkyM2Q2NDU2OGUwMGQ3ZGQ5NjJhMWMiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiI0ZTFiYzA5MC00ZWVjLTQ3MjAtYmEzNy02NzQwZTcwYWRjNjkiLCJ0eXBlIjoiQ09NTUVOVF9DUkVBVEVEIiwidGltZXN0YW1wUHJvb2ZzIjp7InR5cGUiOiJCVU5ETFIiLCJoYXNoUHJlZml4IjoiMSIsInJlc3BvbnNlIjp7ImlkIjoiMHhOM0tYZEJxa1U4SkVfcEo1QWsxSGR6eWN1T0ZsdDNhamlOaUlzdndSTSIsInRpbWVzdGFtcCI6MTY5OTUxMjI5MzQ4NCwidmVyc2lvbiI6IjEuMC4wIiwicHVibGljIjoic3E5SmJwcEtMbEFLdFF3YWxmWDVEYWduR01sVGlyZGl0WGs3eTRqZ29lQTdERU0wWjZjVlBFNXhNUTlrel9UOVZwcFA2QkZIdEh5WkNaT0RlcmNFVldpcHprcjM2dGZRa1I1RURHVVF5TGl2ZHhVemJXZ1Zrenc3RDI3UEpFYTRjZDFVeTZyMThyWUxxRVJnYlJ2QVpwaDVZSlptcFNKazdyM013blFxdXVrdGp2U3BmQ0xGd1N4UDF3ODc5LXNzX0phbE05SUN6UmkzOGhlbk9OaW84Z2xsNkdWOS1vbXJXd1JNWmVyXzE1YnNwQ0s1dHhDd3BZMTM3bmZLd0tENVlCQXV6eHhjajQyNE03emxTSGxzYWZCd2FSd0ZiZjhnSHRXMDNpSkVSNGxSNEd4ZVkwV3ZuWWFCM0tESVNIUXA1M2E5bmxibWlXTzVXY0hIWXNSODNPVDJlSjBQbDNSV0EtX2lta19TTndHUVRDam1BNnRmX1VWd0w4SHpZUzJpeXV1ODViN2lZSzlaUW9oOG5xYk5DNnFpYklDRTRoOUZlM2JON0FnaXRJZTlYekNUT1hEZk1yNGFoakM4a2txSjF6NHpOQUk2LUxlZWlfTWdkOEp0WmgydnFGTlpoWEswbFNhZEZsXzlPaDNBRVQ3dFVkczJFN3MtNnpwUlBkOW9CWnU2LWtOdUhEUko2VFFoWlN3SjlaTzVIWXNjY2JfR18xc283MmFYSnltUjlnZ0pnV3I0SjNiYXdBWVlucW12bXpHa2xZT2xFXzVIVm5NeGYtVXhwVDd6dGRzSGJjOVFFSDZXMmJ6d3hicGpUY3pFWnMzSkNDQjNjLU5ld05Ic2o5UFlNM2I1dFRsVE5QOWtOQXdQWkhXcHQxMXQ3OUx1TmtOR3Q5TGZPZWsiLCJzaWduYXR1cmUiOiJiamRGamdxLS1LTEpFWFlwV1JDdmthazdPUWlQcERfYnZWTmFrMm1mRWp6TlN3VEw2aUxrRlRfMVFvVmNxaEEyVGFJbC1lcWdWMXpBVm10ZVg1WGJMVHZIQkJGNm1rTGtyX09nSVFQbnFpdGU0TUdjcUUtNEI3SkV4dmRaeVk4NU5ta1hXNzFOa09VeW53WW9iVURYcVdQQkZpeHl6aGFUd2hOYjY1SU9iVHBVei1vSjBCODNaNnFSTDdPbFN6OGlRLXFCOVppSDFmT3JvRTRjWWNRbk9DWGtJLVdDRU5nZzB4VU94Z0h0MUdsTUh1cnJEU1Yyd01aQUJGSFQtOHdzendDQm1qQ3B0SUJHUFVwV1FURlp6aWxOVHdmeE5ZejZUd2lfcEpWaWtCYXRDUE1TeGtpc216ZFBFX3M1bUpod2lfZHB1MF9MSmRIV3lVa0x5cDFqQ1YxM2d1SmphbVh6TDRsR3Z4S3FQT1JNOGhOTTZibGR2MWJ3TGV0a0VYcDV5MW5YQktmUkx0VlhDZmZHSmdlV3Y3cVNydWMwYkRhQUFyUGFVdVJRVjN1UVNwUGJycnNuNGs4anJybldUeHRBelctMnplQWo1LV80R0lMaklvQ1VxbzhJZUlYOVBrNWNNanYxNWxTdGFoclFOOWJScFRwamVlTjZxb2tGT3d2T0NJMnJmN25MVFQ1N3ZSNU1PU2tRNVJrS243OHRlYUxWdHNkVW1jOFRDd1pCNlprTWEtX3lvUmNnR1JfazhSX25HRC0zUE9GdUUzTjU5ZXpRNTJwNk02ZFhfZmlCc3c2enhPaDJPX1NpQmk5SjBSVWl6bWpDNExJbXdZbURhdDhZeWltT2FNbWJ3cDJMR3VvV3RFVlB6MGlNUUR1ZkpOTDVZd05YWDZSSUdVUSIsImRlYWRsaW5lSGVpZ2h0IjoxMzAyNjk5LCJibG9jayI6MTMwMjY5OSwidmFsaWRhdG9yU2lnbmF0dXJlcyI6W119fSwiY2hhaW5Qcm9vZnMiOnsidGhpc1B1YmxpY2F0aW9uIjp7InNpZ25hdHVyZSI6IjB4NjFhNDE1ZGY2ZGMxYmY5MWNmZGZkMzViOTUwZDRiNTA4NzI1YmNhNTM2MTFmZWNhOTIxOTk0M2U2NmIwOGI2NDY5ZjA2M2IwNzVlNGM2Nzk1YjM0Y2EzMDI3NGZkYzBhZDQ4NzM0Zjc0NDFlM2JmZDc4MmJjZmY3MDM4Y2IwNjgxYyIsInNpZ25lZEJ5RGVsZWdhdGUiOnRydWUsInNpZ25hdHVyZURlYWRsaW5lIjoxNjk5NTEyMjkyLCJ0eXBlZERhdGEiOnsidHlwZXMiOnsiQ29tbWVudCI6W3sidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoicHJvZmlsZUlkIn0seyJ0eXBlIjoic3RyaW5nIiwibmFtZSI6ImNvbnRlbnRVUkkifSx7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6InBvaW50ZWRQcm9maWxlSWQifSx7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6InBvaW50ZWRQdWJJZCJ9LHsidHlwZSI6InVpbnQyNTZbXSIsIm5hbWUiOiJyZWZlcnJlclByb2ZpbGVJZHMifSx7InR5cGUiOiJ1aW50MjU2W10iLCJuYW1lIjoicmVmZXJyZXJQdWJJZHMifSx7InR5cGUiOiJieXRlcyIsIm5hbWUiOiJyZWZlcmVuY2VNb2R1bGVEYXRhIn0seyJ0eXBlIjoiYWRkcmVzc1tdIiwibmFtZSI6ImFjdGlvbk1vZHVsZXMifSx7InR5cGUiOiJieXRlc1tdIiwibmFtZSI6ImFjdGlvbk1vZHVsZXNJbml0RGF0YXMifSx7InR5cGUiOiJhZGRyZXNzIiwibmFtZSI6InJlZmVyZW5jZU1vZHVsZSJ9LHsidHlwZSI6ImJ5dGVzIiwibmFtZSI6InJlZmVyZW5jZU1vZHVsZUluaXREYXRhIn0seyJ0eXBlIjoidWludDI1NiIsIm5hbWUiOiJub25jZSJ9LHsidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoiZGVhZGxpbmUifV19LCJkb21haW4iOnsibmFtZSI6IkxlbnMgUHJvdG9jb2wgUHJvZmlsZXMiLCJ2ZXJzaW9uIjoiMiIsImNoYWluSWQiOjEzNywidmVyaWZ5aW5nQ29udHJhY3QiOiIweERiNDZkMURjMTU1NjM0RmJDNzMyZjkyRTg1M2IxMEIyODhBRDVhMWQifSwidmFsdWUiOnsiYWN0aW9uTW9kdWxlcyI6W10sImFjdGlvbk1vZHVsZXNJbml0RGF0YXMiOltdLCJjb250ZW50VVJJIjoiYXI6Ly80ZC1jWUY3amt0bW14aUdnejlraGNLbW84UGJINzQwTzB3YjlXODlqM1RFIiwiZGVhZGxpbmUiOjE2OTk1MTIyOTIsIm5vbmNlIjowLCJwb2ludGVkUHJvZmlsZUlkIjoiMHhiZWUxIiwicG9pbnRlZFB1YklkIjoiMHgwZWNhIiwicHJvZmlsZUlkIjoiMHgwMWEyN2UiLCJyZWZlcmVuY2VNb2R1bGUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAiLCJyZWZlcmVuY2VNb2R1bGVEYXRhIjoiMHgiLCJyZWZlcmVuY2VNb2R1bGVJbml0RGF0YSI6IjB4IiwicmVmZXJyZXJQcm9maWxlSWRzIjpbXSwicmVmZXJyZXJQdWJJZHMiOltdfX0sImJsb2NrSGFzaCI6IjB4NjEyZWZiNjU4MGJhZmFmNmZmMzZiMTQ0ZWY1ZTVlYTVhYTJhZGI2NmQwMTdjNWUyYTcwYjhjMGM2MWNkYTQ1MCIsImJsb2NrTnVtYmVyIjo0OTcyMzU2MCwiYmxvY2tUaW1lc3RhbXAiOjE2OTk1MTIyOTJ9LCJwb2ludGVyIjp7ImxvY2F0aW9uIjoiYXI6Ly9QRUVTRVI0cVhXNkVRNGtYazlfLXQwSXpfSUNBTzMyYU1US2JMNTZIbXBVIiwidHlwZSI6Ik9OX0RBIn19LCJwdWJsaWNhdGlvbklkIjoiMHgwMWEyN2UtMHgwNTNiLURBLTRlMWJjMDkwIiwiZXZlbnQiOnsiY29tbWVudFBhcmFtcyI6eyJwcm9maWxlSWQiOiIweDAxYTI3ZSIsImNvbnRlbnRVUkkiOiJhcjovLzRkLWNZRjdqa3RtbXhpR2d6OWtoY0ttbzhQYkg3NDBPMHdiOVc4OWozVEUiLCJhY3Rpb25Nb2R1bGVzIjpbXSwiYWN0aW9uTW9kdWxlc0luaXREYXRhcyI6W10sInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZUluaXREYXRhIjoiMHgiLCJyZWZlcmVuY2VNb2R1bGVEYXRhIjoiMHgiLCJyZWZlcnJlclByb2ZpbGVJZHMiOltdLCJyZWZlcnJlclB1YklkcyI6W10sInBvaW50ZWRQcm9maWxlSWQiOiIweGJlZTEiLCJwb2ludGVkUHViSWQiOiIweDBlY2EifSwicHViSWQiOiIweDA1M2IiLCJhY3Rpb25Nb2R1bGVzSW5pdFJldHVybkRhdGFzIjpbXSwicmVmZXJlbmNlTW9kdWxlUmV0dXJuRGF0YSI6IjB4IiwicmVmZXJlbmNlTW9kdWxlSW5pdFJldHVybkRhdGEiOiIweCIsInRyYW5zYWN0aW9uRXhlY3V0b3IiOiIweGFkOWNDMWE4NWU3NEJkMTMwMGM1NzA4MmJhODZFZDIxRTZBQUE1ZTgiLCJ0aW1lc3RhbXAiOjE2OTk1MTIyOTJ9fQ",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "bN4Sf7lUO4TZwpPjtP3BUzBSG92io0McNgMXTxXUOcM",
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
						From:     "0xad9cC1a85e74Bd1300c57082ba86Ed21E6AAA5e8",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle:        "sababa.lens",
							Title:         "It’s on Amazon so unsure. I have it in my stan store as a direct link. Soo dunno",
							Body:          "It’s on Amazon so unsure. I have it in my stan store as a direct link. Soo dunno",
							ProfileID:     "0x01a27e",
							PublicationID: "0x01a27e-0x053b-DA-4e1bc090",
							ContentURI:    "ar://4d-cYF7jktmmxiGgz9khcKmo8PbH740O0wb9W89j3TE",
							Timestamp:     1699512292,
							Target: &metadata.SocialPost{
								Handle:        "dmcsvn.lens",
								Body:          "Wow 🤩, Congrats Samm,is it available only in the US?",
								ProfileID:     "0xbee1",
								PublicationID: "0xbee1-0x0eca-DA-0a949fb1",
								ContentURI:    "https://arweave.net/D_1C5K0hoBU6Cbfg3lb1bUIwVhkZV-O6hkqRX_K-yko",
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1699512726,
			},
			wantError: require.NoError,
		},
		{
			name: "Momoka V2 Quote",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Block: arweave.Block{
						Timestamp: 1699537271,
					},
					Transaction: arweave.Transaction{
						ID:       "wJ40YsMEMybIZBXMu5a-lH6RGv7A54B2Lr0vYJCLvAQ",
						Quantity: "0",
						Owner:    "BPEOUTIHi3pEBieHAHp4NK8alylNB9ueEuYdyxhZ43vcX6HtRK_HoyRwVDCd-jjBNKalc8Q-E5mGntgwl0-_NNM",
						Data:     "eyJzaWduYXR1cmUiOiIweDRmYjI5ZTI2M2Q3YzY4YzU3NTUxMmVjNmRkOGY3MGRiMzY0YmIyYzY3YzlhNDEzNTI2YTFmMzQ0NjFmZjAwMTYwMzk1MmY0NTQwYWEzMTExOTY5ODZiYjhkNGE0ZWRkZTA3MTA0YzQxZDdjYjY3MmY2OGEzMDU5NGM5ZTY5MDA3MWIiLCJkYXRhQXZhaWxhYmlsaXR5SWQiOiIxNWVhZjE4Mi1jZmY1LTQ4MDItOTkxYS1iNTU5Yjk0Zjg1YTkiLCJ0eXBlIjoiUVVPVEVfQ1JFQVRFRCIsInRpbWVzdGFtcFByb29mcyI6eyJ0eXBlIjoiQlVORExSIiwiaGFzaFByZWZpeCI6IjEiLCJyZXNwb25zZSI6eyJpZCI6Ijd0dFF5Uk5HVUxJUkowOHlZU09HSzZRcjhuMWgtQUdzT1ZVZjJCQUx2V2ciLCJ0aW1lc3RhbXAiOjE2OTk1MzcwODg1MjEsInZlcnNpb24iOiIxLjAuMCIsInB1YmxpYyI6InNxOUpicHBLTGxBS3RRd2FsZlg1RGFnbkdNbFRpcmRpdFhrN3k0amdvZUE3REVNMFo2Y1ZQRTV4TVE5a3pfVDlWcHBQNkJGSHRIeVpDWk9EZXJjRVZXaXB6a3IzNnRmUWtSNUVER1VReUxpdmR4VXpiV2dWa3p3N0QyN1BKRWE0Y2QxVXk2cjE4cllMcUVSZ2JSdkFacGg1WUpabXBTSms3cjNNd25RcXV1a3RqdlNwZkNMRndTeFAxdzg3OS1zc19KYWxNOUlDelJpMzhoZW5PTmlvOGdsbDZHVjktb21yV3dSTVplcl8xNWJzcENLNXR4Q3dwWTEzN25mS3dLRDVZQkF1enh4Y2o0MjRNN3psU0hsc2FmQndhUndGYmY4Z0h0VzAzaUpFUjRsUjRHeGVZMFd2bllhQjNLRElTSFFwNTNhOW5sYm1pV081V2NISFlzUjgzT1QyZUowUGwzUldBLV9pbWtfU053R1FUQ2ptQTZ0Zl9VVndMOEh6WVMyaXl1dTg1YjdpWUs5WlFvaDhucWJOQzZxaWJJQ0U0aDlGZTNiTjdBZ2l0SWU5WHpDVE9YRGZNcjRhaGpDOGtrcUoxejR6TkFJNi1MZWVpX01nZDhKdFpoMnZxRk5aaFhLMGxTYWRGbF85T2gzQUVUN3RVZHMyRTdzLTZ6cFJQZDlvQlp1Ni1rTnVIRFJKNlRRaFpTd0o5Wk81SFlzY2NiX0dfMXNvNzJhWEp5bVI5Z2dKZ1dyNEozYmF3QVlZbnFtdm16R2tsWU9sRV81SFZuTXhmLVV4cFQ3enRkc0hiYzlRRUg2VzJiend4YnBqVGN6RVpzM0pDQ0IzYy1OZXdOSHNqOVBZTTNiNXRUbFROUDlrTkF3UFpIV3B0MTF0NzlMdU5rTkd0OUxmT2VrIiwic2lnbmF0dXJlIjoiYnRPdm9Sd3hCSWp4aVVoUlJiT1Y3d1BuT3lfZ2hPZDdIekF6SFdnWjlHY1NWTXl1T0pRQVozNnp6SW5FNDd0WVBtUzBoUDNBTVhvRUJubTNIUGJfMTNvYldrSXVhS1pxamh5bDZaRmp4YXBSZlcxVm8ycDBneUk0MjdTa3JERUlsWVJWaWQxSTMzbkFIaVNaQVlUYmJHQzBhYlNYenVhR04zd0pFZ18yU19STXFTcWZzT2lEOExFdlI1QkpJNFRuMFA3QnJFcXFUUHZ3dkpDZm93bFBiZkdzaEJ6RVNlajNDeG5WSDZBYmJ0RFJhZjdkcDVDWWN3VUNUdjNGTTQzZkZGUWFyZ0ZkUjhYYkxDVlJpc29jTmlnUFEwb3A3QVVSSjhENTVKdU1CZmtxUm9SSnRtanppY25Pb1VCMkZqSzZKazg1bjV3eEI0bm1feEhYSWJ2RUhxcFJKVjdIbjNfTVlaWEQ2RDRBbTJyY2FXUTVoRExZVTU0YkNHaXM5UXJqTVZ3MXNPUUdTS3VxRG5EMFNjY19SN2pLbkc3MHRaRmRpM0NoZUxCSUFzYWlOUk5iN25vbFd1ZEZiYTAzWVE5akhiV3ExRUJQYXpYM1cwdERfX1J5S05JakxpVW4xVXRpS2dBclpZWFlRUGhDYWxib0xZRGlDcC1jZDhTREZPNWhOSGZDM0ZrVS1XYWJpeHFyeDIzMk9NWE1JZG1WLWk5ZWdUVUxXMUlFcHNKQjc2S2czS3M4YnFObWpBczBqSloxQVYwNTJock9RX294d1gwYWQyMXFkQS1ycENFc0lwVVEtRzhsUWVDUGdyVTNBejh6M3hIQ0FIMXZhVnFxYUQyMGdCM1lqYVVDMHFlRXB5aVNFOGpkMTAzSHF5MmxzZkhCNG1KeGhJNk1IQmciLCJkZWFkbGluZUhlaWdodCI6MTMwMjg4MywiYmxvY2siOjEzMDI4ODMsInZhbGlkYXRvclNpZ25hdHVyZXMiOltdfX0sImNoYWluUHJvb2ZzIjp7InRoaXNQdWJsaWNhdGlvbiI6eyJzaWduYXR1cmUiOiIweGJlYjEzMTMyNDUzNTc4NGNjNTk1ZjE0ZjNkNTVlMjYyMWIzZGQ3MTA1NGQwZWQ2NjY0NjdhNmQwYjk2ZjMxMjg0OTU4YjMyYmZhZWVlYjQzZGVlNDI1ZDc4OGQzZDNjZTMyNzgxYjc1MTY0OTU4ZDEwYWFhMDA1MzQ1OTM0ZTc5MWIiLCJzaWduZWRCeURlbGVnYXRlIjpmYWxzZSwic2lnbmF0dXJlRGVhZGxpbmUiOjE2OTk1MzcwODYsInR5cGVkRGF0YSI6eyJ0eXBlcyI6eyJRdW90ZSI6W3sidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoicHJvZmlsZUlkIn0seyJ0eXBlIjoic3RyaW5nIiwibmFtZSI6ImNvbnRlbnRVUkkifSx7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6InBvaW50ZWRQcm9maWxlSWQifSx7InR5cGUiOiJ1aW50MjU2IiwibmFtZSI6InBvaW50ZWRQdWJJZCJ9LHsidHlwZSI6InVpbnQyNTZbXSIsIm5hbWUiOiJyZWZlcnJlclByb2ZpbGVJZHMifSx7InR5cGUiOiJ1aW50MjU2W10iLCJuYW1lIjoicmVmZXJyZXJQdWJJZHMifSx7InR5cGUiOiJieXRlcyIsIm5hbWUiOiJyZWZlcmVuY2VNb2R1bGVEYXRhIn0seyJ0eXBlIjoiYWRkcmVzc1tdIiwibmFtZSI6ImFjdGlvbk1vZHVsZXMifSx7InR5cGUiOiJieXRlc1tdIiwibmFtZSI6ImFjdGlvbk1vZHVsZXNJbml0RGF0YXMifSx7InR5cGUiOiJhZGRyZXNzIiwibmFtZSI6InJlZmVyZW5jZU1vZHVsZSJ9LHsidHlwZSI6ImJ5dGVzIiwibmFtZSI6InJlZmVyZW5jZU1vZHVsZUluaXREYXRhIn0seyJ0eXBlIjoidWludDI1NiIsIm5hbWUiOiJub25jZSJ9LHsidHlwZSI6InVpbnQyNTYiLCJuYW1lIjoiZGVhZGxpbmUifV19LCJkb21haW4iOnsibmFtZSI6IkxlbnMgUHJvdG9jb2wgUHJvZmlsZXMiLCJ2ZXJzaW9uIjoiMiIsImNoYWluSWQiOjEzNywidmVyaWZ5aW5nQ29udHJhY3QiOiIweERiNDZkMURjMTU1NjM0RmJDNzMyZjkyRTg1M2IxMEIyODhBRDVhMWQifSwidmFsdWUiOnsiYWN0aW9uTW9kdWxlcyI6W10sImFjdGlvbk1vZHVsZXNJbml0RGF0YXMiOltdLCJjb250ZW50VVJJIjoiYXI6Ly9mRWROSE1OMzdzazFwTkxKNmRYREswc183Q29rbzdBQ0VTRWt2N0xkQ0xBIiwiZGVhZGxpbmUiOjE2OTk1MzcwODYsIm5vbmNlIjo2NiwicG9pbnRlZFByb2ZpbGVJZCI6IjB4NjMyZCIsInBvaW50ZWRQdWJJZCI6IjB4NTgiLCJwcm9maWxlSWQiOiIweDAxNDU3MyIsInJlZmVyZW5jZU1vZHVsZSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInJlZmVyZW5jZU1vZHVsZURhdGEiOiIweCIsInJlZmVyZW5jZU1vZHVsZUluaXREYXRhIjoiMHgiLCJyZWZlcnJlclByb2ZpbGVJZHMiOltdLCJyZWZlcnJlclB1YklkcyI6W119fSwiYmxvY2tIYXNoIjoiMHg3ZGUzZThkMDVhZTIyY2NjNTA1ZWFjYzY5NjEwZGQ4ODYzYWI0YTcxMjA0OGI3NDkxZDk0YzE3YTk1OTYxZDliIiwiYmxvY2tOdW1iZXIiOjQ5NzM0OTkwLCJibG9ja1RpbWVzdGFtcCI6MTY5OTUzNzA4Nn0sInBvaW50ZXIiOnsibG9jYXRpb24iOiJhcjovLzZUUkpOX2NLS3F5WFhsT3pTME1WbnFybmVtdWdBZGlpcXRRQl85RFdjSVUiLCJ0eXBlIjoiT05fREEifX0sInB1YmxpY2F0aW9uSWQiOiIweDAxNDU3My0weGZhLURBLTE1ZWFmMTgyIiwiZXZlbnQiOnsicXVvdGVQYXJhbXMiOnsicHJvZmlsZUlkIjoiMHgwMTQ1NzMiLCJjb250ZW50VVJJIjoiYXI6Ly9mRWROSE1OMzdzazFwTkxKNmRYREswc183Q29rbzdBQ0VTRWt2N0xkQ0xBIiwiYWN0aW9uTW9kdWxlcyI6W10sImFjdGlvbk1vZHVsZXNJbml0RGF0YXMiOltdLCJyZWZlcmVuY2VNb2R1bGUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAiLCJyZWZlcmVuY2VNb2R1bGVJbml0RGF0YSI6IjB4IiwicmVmZXJlbmNlTW9kdWxlRGF0YSI6IjB4IiwicmVmZXJyZXJQcm9maWxlSWRzIjpbXSwicmVmZXJyZXJQdWJJZHMiOltdLCJwb2ludGVkUHJvZmlsZUlkIjoiMHg2MzJkIiwicG9pbnRlZFB1YklkIjoiMHg1OCJ9LCJwdWJJZCI6IjB4ZmEiLCJhY3Rpb25Nb2R1bGVzSW5pdFJldHVybkRhdGFzIjpbXSwicmVmZXJlbmNlTW9kdWxlUmV0dXJuRGF0YSI6IjB4IiwicmVmZXJlbmNlTW9kdWxlSW5pdFJldHVybkRhdGEiOiIweCIsInRyYW5zYWN0aW9uRXhlY3V0b3IiOiIweDJiNjA5MWRhMzEyQTM1OTFERjZhOTRiMjRkRDNjMzRkMkIwMUNhMjIiLCJ0aW1lc3RhbXAiOjE2OTk1MzcwODZ9fQ",
					},
				},
				config: &config.Module{
					IPFSGateways: []string{"https://ipfs.rss3.page/"},
					Endpoint:     endpoint.MustGet(filter.NetworkPolygon),
				},
			},
			want: &schema.Feed{
				ID:       "wJ40YsMEMybIZBXMu5a-lH6RGv7A54B2Lr0vYJCLvAQ",
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
						From:     "0x2b6091da312A3591DF6a94b24dD3c34d2B01Ca22",
						To:       "llq4QhHA7fQWBKd6V8vzAhK-qy_JulpzlYUEgxAgJ4E",
						Metadata: &metadata.SocialPost{
							Handle:        "69_69.lens",
							ProfileID:     "0x014573",
							PublicationID: "0x014573-0xfa-DA-15eaf182",
							Title:         "Quote by @69_69",
							Body:          "highjacking consensus mechanism",
							ContentURI:    "ar://fEdNHMN37sk1pNLJ6dXDK0s_7Coko7ACESEkv7LdCLA",
							Timestamp:     1699537086,
							Target: &metadata.SocialPost{
								Handle:        "olimpio.lens",
								Body:          "If you hold spot ETH, there are a few LSD protocols where you can stake it that don't have a token (🪂):\n\nMany of them have some sort of \"points\" system.\n\n• ether fi: just announced that points\u003c\u003edecentralized governance. \n@ether_fi\n\n• mev io: 50M TVL and backers like 0xMaki. \n@mevdotio\n\n• Swell: most of you already know this one, and they have their pearl points and Voyage, but it's probably too diluted. Alpha: swETH is trading below book value, so it could be a small RFV play if you trust the team. For each swETH there is 1.04144628 ETH in backing- but you can buy one swETH with ~1.012 ETH. \n\n• EigenLayer: this is a re-staking platform where you stake LSD tokens like stETH and rETH. Currently, there's no deposit limit available, but there was a vote to add a basket of more assets so it will be opened soon. This is the biggest opportunity - to monitor. More here: https://twitter.com/OlimpioCrypto/status/1721958916295180491\n\n💡 Important: You can't unstake in any of them for now (you can in Eigen with a 7-day delay). It's one-way staking, so only put whatever you can afford to park for a few months. Also, consider smart contract risk.\n\nIf you hold ETH then you are probably staking it somewhere, the play is to diversify the ETH staking solutions into some that also might or might not do an airdrop in the future, while still earning staking APR.\n\n⚡️",
								ProfileID:     "0x632d",
								PublicationID: "0x632d-0x58-DA-0883e9b6",
								ContentURI:    "https://data.lens.phaver.com/api/lens/posts/8b279a78-119f-430a-a00d-af5551f0097b",
							},
						},
					},
				},
				Status:    true,
				Timestamp: 1699537271,
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
		})
	}
}
