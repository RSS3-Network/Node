package arweave_test

import (
	"context"
	"io"
	"sync"
	"testing"

	"github.com/rss3-network/node/provider/arweave"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce     sync.Once
	arweaveClient arweave.Client
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		arweaveClient, err = arweave.NewClient()
		require.NoError(t, err)
	})
}

func TestClient_GetTransactionData(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		id string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "C3JdvFDS7lx-OXcauYitjYvJNt8kJdQJo8GX8kC5RbM",
			arguments: arguments{
				id: "C3JdvFDS7lx-OXcauYitjYvJNt8kJdQJo8GX8kC5RbM",
			},
		},
		{
			name: "cNlfQ2J_DBvHHDIV09k0oPQYru91afRDMnSN8bqb-to",
			arguments: arguments{
				id: "cNlfQ2J_DBvHHDIV09k0oPQYru91afRDMnSN8bqb-to",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			response, err := arweaveClient.GetTransactionData(context.Background(), testcase.arguments.id)
			require.NoError(t, err)

			defer lo.Try(response.Close)

			data, err := io.ReadAll(response)
			require.NoError(t, err)

			t.Log(string(data))
		})
	}
}

func TestClient_GetBlockByHeight(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		height int64
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "1312246",
			arguments: arguments{
				height: 1312246,
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			block, err := arweaveClient.GetBlockByHeight(context.Background(), testcase.arguments.height)
			require.NoError(t, err)

			t.Log(block.Height)
			t.Log(block.Hash)
			t.Log(block.Timestamp)
		})
	}
}

func TestClient_GetTransactionByID(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		id string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "PcJuwNWK_NR8JUeoaxid81UHk2oazls4OtE6320TEl0",
			arguments: arguments{
				id: "PcJuwNWK_NR8JUeoaxid81UHk2oazls4OtE6320TEl0",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			transaction, err := arweaveClient.GetTransactionByID(context.Background(), testcase.arguments.id)
			require.NoError(t, err)

			t.Log(transaction.ID)
		})
	}
}
