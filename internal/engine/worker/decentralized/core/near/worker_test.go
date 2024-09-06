package near_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/near"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/core/near"
	"github.com/rss3-network/node/provider/near"
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
			name: "Near Transfer Transaction",
			arguments: arguments{
				task: &source.Task{
					Network: network.Near,
					Block: near.Block{
						Header: near.BlockHeader{
							Height:    127337369,
							GasPrice:  "100000000",
							Timestamp: 1725525629383529996,
						},
					},
					Transaction: near.Transaction{
						Transaction: near.TransactionDetails{
							SignerID:   "sweat_welcome.near",
							ReceiverID: "bba0ca9ead96290db66a3f6c28df6e9bfb9598d3dc36860a77364de6c28d8b88",
							Actions: []near.Action{
								{
									Transfer: &near.TransferAction{
										Deposit: "2000000000000000000000",
									},
								},
							},
							Hash: "CaPbVFXTdzH9qUJ7WZZvB2CV8GpkqoxzVhBZa75FZvq5",
						},
						TransactionOutcome: near.TransactionOutcome{
							Outcome: near.Outcome{
								GasBurnt: 4174947687500,
							},
						},
					},
				},
			},
			want: &activityx.Activity{
				ID:        "CaPbVFXTdzH9qUJ7WZZvB2CV8GpkqoxzVhBZa75FZvq5",
				Network:   network.Near,
				Index:     0,
				From:      "sweat_welcome.near",
				To:        "bba0ca9ead96290db66a3f6c28df6e9bfb9598d3dc36860a77364de6c28d8b88",
				Type:      typex.TransactionTransfer,
				Timestamp: 1725525629,
				Fee: &activityx.Fee{
					Amount:  lo.Must(decimal.NewFromString("417494768750000000000")),
					Decimal: 24,
				},
				Status: true,
				Actions: []*activityx.Action{
					{
						Type: typex.TransactionTransfer,
						From: "sweat_welcome.near",
						To:   "bba0ca9ead96290db66a3f6c28df6e9bfb9598d3dc36860a77364de6c28d8b88",
						Metadata: metadata.TransactionTransfer{
							Name:     "NEAR",
							Symbol:   "NEAR",
							Value:    lo.ToPtr(lo.Must(decimal.NewFromString("2000000000000000000000"))),
							Decimals: 24,
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

			t.Log(string(lo.Must(json.MarshalIndent(feed, "", "\x20\x20"))))

			require.Equal(t, testcase.want, feed)
		})
	}
}
