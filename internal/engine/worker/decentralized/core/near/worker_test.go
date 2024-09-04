package near_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rss3-network/node/config"
	source "github.com/rss3-network/node/internal/engine/source/near"
	worker "github.com/rss3-network/node/internal/engine/worker/decentralized/core/near"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/samber/lo"
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
		// {
		// 	name: "Near Transaction",
		// 	arguments: arguments{
		// 		task: &source.Task{
		// 			Network: network.Near,
		// 			Block: near.Block{
		// 				Header: near.BlockHeader{
		// 					Timestamp: 1689914323,
		// 				},
		// 			},
		// 			Transaction: near.TransactionDetails{
		// 				Hash:        "7yY4AXKyMpNcC9qzghiJVSbRgJhEdmkP8qZyVH3u6Erx",
		// 				SignerID:    "relayer.pagodaplatform.near",
		// 				ReceiverID:  "djku30.near",
		// 				PriorityFee: 0,
		// 			},
		// 		},
		// 	},
		// 	want: &activityx.Activity{
		// 		ID:      "7yY4AXKyMpNcC9qzghiJVSbRgJhEdmkP8qZyVH3u6Erx",
		// 		Network: network.Near,
		// 		Index:   0,
		// 		From:    "relayer.pagodaplatform.near",
		// 		To:      "djku30.near",
		// 		Type:    typex.TransactionTransfer,
		// 		Actions: []*activityx.Action{
		// 			{
		// 				Type:     typex.TransactionTransfer,
		// 				From:     "relayer.pagodaplatform.near",
		// 				To:       "djku30.near",
		// 				Metadata: metadata.TransactionTransfer{},
		// 			},
		// 		},
		// 		Status:    true,
		// 		Timestamp: 1689914323,
		// 	},
		// 	wantError: require.NoError,
		// },
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
