package client_test

import (
	"context"
	"io"
	"sync"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave/client"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce         sync.Once
	arweaveHTTPClient client.HTTPClient
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		arweaveHTTPClient, err = client.NewHTTPClient()
		require.NoError(t, err)
	})
}

func TestHTTPClient_Fetch(t *testing.T) {
	setup(t)

	t.Parallel()

	type arguments struct {
		id string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "Get arweave transaction",
			arguments: arguments{
				id: "C3JdvFDS7lx-OXcauYitjYvJNt8kJdQJo8GX8kC5RbM",
			},
		},
		{
			name: "Get arweave transaction",
			arguments: arguments{
				id: "cNlfQ2J_DBvHHDIV09k0oPQYru91afRDMnSN8bqb-to",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			response, err := arweaveHTTPClient.Fetch(context.Background(), testcase.arguments.id)
			require.NoError(t, err)

			defer lo.Try(response.Close)

			data, err := io.ReadAll(response)
			require.NoError(t, err)
			log.Debug(string(data))
		})
	}
}
