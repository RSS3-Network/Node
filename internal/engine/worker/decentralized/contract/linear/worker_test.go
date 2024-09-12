package linear_test

import (
	"context"
	"testing"

	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/near"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/contract/linear"
	"github.com/rss3-network/node/provider/near"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestWorker_Near(t *testing.T) {
	t.Parallel()

	type arguments struct {
		task   *source.Task
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      *activityx.Activity
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "LiNEAR Swap",
			arguments: arguments{
				task: &source.Task{
					Network: network.Near,
					Block: near.Block{
						Header: near.BlockHeader{
							Height:    127708374,
							GasPrice:  "100000000",
							Timestamp: 1725949386712739154,
						},
					},
					Transaction: near.Transaction{
						Transaction: near.TransactionDetails{
							SignerID:   "strawberry1212.near",
							ReceiverID: "802d89b6e511b335f05024a65161bce7efc3f311.factory.bridge.near",
							Actions: []near.Action{
								{
									FunctionCall: &near.FunctionCallAction{
										Deposit:    "1",
										MethodName: "ft_transfer_call",
										Gas:        180000000000000,
										Args:       "eyJyZWNlaXZlcl9pZCI6InYyLnJlZi1maW5hbmNlLm5lYXIiLCJhbW91bnQiOiI0MjU1NzQwMDAwMDAwMDAwMDAwMDAwIiwibXNnIjoie1wiZm9yY2VcIjowLFwiYWN0aW9uc1wiOlt7XCJwb29sX2lkXCI6NDQxOCxcInRva2VuX2luXCI6XCI4MDJkODliNmU1MTFiMzM1ZjA1MDI0YTY1MTYxYmNlN2VmYzNmMzExLmZhY3RvcnkuYnJpZGdlLm5lYXJcIixcInRva2VuX291dFwiOlwid3JhcC5uZWFyXCIsXCJhbW91bnRfaW5cIjpcIjMxOTE4MDUwMDAwMDAwMDAwMDAwMDBcIixcIm1pbl9hbW91bnRfb3V0XCI6XCIwXCJ9LHtcInBvb2xfaWRcIjozODc5LFwidG9rZW5faW5cIjpcIndyYXAubmVhclwiLFwidG9rZW5fb3V0XCI6XCJ1c2R0LnRldGhlci10b2tlbi5uZWFyXCIsXCJtaW5fYW1vdW50X291dFwiOlwiMFwifSx7XCJwb29sX2lkXCI6NDE3OSxcInRva2VuX2luXCI6XCJ1c2R0LnRldGhlci10b2tlbi5uZWFyXCIsXCJ0b2tlbl9vdXRcIjpcImEwYjg2OTkxYzYyMThiMzZjMWQxOWQ0YTJlOWViMGNlMzYwNmViNDguZmFjdG9yeS5icmlkZ2UubmVhclwiLFwibWluX2Ftb3VudF9vdXRcIjpcIjI1MDIwMjM1XCJ9LHtcInBvb2xfaWRcIjo0NDE4LFwidG9rZW5faW5cIjpcIjgwMmQ4OWI2ZTUxMWIzMzVmMDUwMjRhNjUxNjFiY2U3ZWZjM2YzMTEuZmFjdG9yeS5icmlkZ2UubmVhclwiLFwidG9rZW5fb3V0XCI6XCJ3cmFwLm5lYXJcIixcImFtb3VudF9pblwiOlwiMTA2MzkzNTAwMDAwMDAwMDAwMDAwMFwiLFwibWluX2Ftb3VudF9vdXRcIjpcIjBcIn0se1wicG9vbF9pZFwiOjMsXCJ0b2tlbl9pblwiOlwid3JhcC5uZWFyXCIsXCJ0b2tlbl9vdXRcIjpcImEwYjg2OTkxYzYyMThiMzZjMWQxOWQ0YTJlOWViMGNlMzYwNmViNDguZmFjdG9yeS5icmlkZ2UubmVhclwiLFwibWluX2Ftb3VudF9vdXRcIjpcIjgyODMwMTZcIn1dfSJ9",
									},
								},
							},
							Hash: "2jMSAk3qe1sBUJnQMMvhPaVrjEHLpcrWe8psvH4Jf5Ke",
						},
						TransactionOutcome: near.TransactionOutcome{
							Outcome: near.Outcome{
								GasBurnt: 352882192100,
							},
						},
						ReceiptsOutcome: []near.ReceiptOutcome{
							{
								Outcome: near.Outcome{
									ExecutorID: "802d89b6e511b335f05024a65161bce7efc3f311.factory.bridge.near",
									Logs: []string{
										"EVENT_JSON:{\"standard\":\"nep141\",\"version\":\"1.0.0\",\"event\":\"ft_transfer\",\"data\":[{\"old_owner_id\":\"strawberry1212.near\",\"new_owner_id\":\"v2.ref-finance.near\",\"amount\":\"4255740000000000000000\"}]}",
									},
								},
							},
							{
								Outcome: near.Outcome{
									ExecutorID: "a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48.factory.bridge.near",
									Logs: []string{
										"EVENT_JSON:{\"standard\":\"nep141\",\"version\":\"1.0.0\",\"event\":\"ft_transfer\",\"data\":[{\"old_owner_id\":\"v2.ref-finance.near\",\"new_owner_id\":\"strawberry1212.near\",\"amount\":\"33470457\"}]}",
									},
								},
							},
						},
					},
				},
			},
			want: &activityx.Activity{
				ID:        "2jMSAk3qe1sBUJnQMMvhPaVrjEHLpcrWe8psvH4Jf5Ke",
				Network:   network.Near,
				Index:     0,
				From:      "strawberry1212.near",
				To:        "802d89b6e511b335f05024a65161bce7efc3f311.factory.bridge.near",
				Type:      typex.ExchangeSwap,
				Platform:  workerx.PlatformLiNEAR.String(),
				Timestamp: 1725949386,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("35288219210000000000")),
					Decimal: 24,
				},
				Status: true,
				Actions: []*activityx.Action{
					{
						Type:     typex.ExchangeSwap,
						Platform: workerx.PlatformLiNEAR.String(),
						From:     "strawberry1212.near",
						To:       "v2.ref-finance.near",
						Metadata: metadata.ExchangeSwap{
							From: metadata.Token{
								Address:  lo.ToPtr("802d89b6e511b335f05024a65161bce7efc3f311.factory.bridge.near"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("4255740000000000000000"))),
								Standard: metadata.StandardNEP141,
							},
							To: metadata.Token{
								Address:  lo.ToPtr("a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48.factory.bridge.near"),
								Value:    lo.ToPtr(lo.Must(decimal.NewFromString("33470457"))),
								Standard: metadata.StandardNEP141,
							},
						},
					},
				},
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

			feed, err := instance.Transform(ctx, testcase.arguments.task)
			testcase.wantError(t, err)

			// t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}
