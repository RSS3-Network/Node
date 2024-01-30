package arweave_test

import (
	"context"
	"testing"

	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/arweave"
	worker "github.com/rss3-network/node/internal/engine/worker/fallback/arweave"
	"github.com/rss3-network/node/provider/arweave"
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
		{
			name: "Arweave native transfer",
			arguments: arguments{
				task: &source.Task{
					Network: filter.NetworkArweave,
					Transaction: arweave.Transaction{
						ID:       "mQFakpEtbvv8eAjxmWYLcIO8QJJP2ZFYOhP1imDcnuY",
						Reward:   "3847185",
						Quantity: "14243667552815",
						Owner:    "t46TzeVlqOpgyJnqOin9fSIluSy_s2SnKI2-WtJ-BOD1WXZIg13DcmBFWgDZ7zDai4Jf46-UpjuMmxTorLuSAxNRZJMzPnhqqWXONhMaiESjwM2tWc7d6bSxntkLcZuywGw9tpK4rh8hfXEtWQe65t_rxT_THGj1Il4P5sSbfLg1s5Xam9to_rxB5hoHGPdNmOTIDi-RA_DmV6KJ30spZDu8FZk7HHFOrUzrhy5TgDg5G4KNxllc3Mo8W1U5MTWlJ1A4nC1DGuRFJ7m0YZ-e0_Yf7H6BfYlRruW485t6f8cLdEFwALK3LeRCaN3VgZ5U6g2vWloorEBruiCSACFeN2lZf8kEdkaJWHymCU9FNhrdqyPtzToQaaSpykWLqwBcjSvYaGpAyr5IrKxHZ9Qj8oD7nyQ2tpGtasQAHY3sn7rokc2dA_srivn25RffuyrFfs5SPBMJTU0MqoDoPT5HjyZaMiq06IcH8wMCqx4bde02kLwPEgVMJtQ12SmsHNmGQmUjF3F1JfcfKD5yYhCUfPeAkxo9MGYgxlIanKHk1MNi1aiayCCGhd_OX9jCbsAyxC0cnaxDIF_4RHB_0rNXTrWBwGiV53GC3ncAE-SnNeejwOxXV_wc1RoaTYh7EkAxQV6whzYMLXeQ0yovzrU9VNe2ryoC5Eul4iqukNgA00k",
						Target:   "4u5gMvlfVhkn_atzuagjO92H_xJLtVNjucSfEYBrL0E",
					},
					Block: arweave.Block{
						Timestamp: 1689914323,
					},
				},
			},
			want: &schema.Feed{
				ID:      "mQFakpEtbvv8eAjxmWYLcIO8QJJP2ZFYOhP1imDcnuY",
				Network: filter.NetworkArweave,
				Index:   0,
				From:    "JaUubKRNhJP9i1iDFt-n_s0zzqV97x8d_7ex3ZZv3CE",
				To:      "4u5gMvlfVhkn_atzuagjO92H_xJLtVNjucSfEYBrL0E",
				Type:    filter.TypeTransactionTransfer,
				Fee: &schema.Fee{
					Amount:  decimal.NewFromInt(3847185),
					Decimal: 12,
				},
				Actions: []*schema.Action{
					{
						Type: filter.TypeTransactionTransfer,
						From: "JaUubKRNhJP9i1iDFt-n_s0zzqV97x8d_7ex3ZZv3CE",
						To:   "4u5gMvlfVhkn_atzuagjO92H_xJLtVNjucSfEYBrL0E",
						Metadata: metadata.TransactionTransfer{
							Value: lo.ToPtr(lo.Must(decimal.NewFromString("14243667552815"))),
						},
					},
				},
				Status:    true,
				Timestamp: 1689914323,
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
