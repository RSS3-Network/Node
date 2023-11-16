package ipfs_test

import (
	"context"
	"io"
	"sync"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/internal/provider/ipfs"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce      sync.Once
	ipfsHTTPClient ipfs.Client
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		ipfsHTTPClient, err = ipfs.NewHTTPClient()
		require.NoError(t, err)
	})
}

func TestQuickFetch(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		path string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "character metadata",
			arguments: arguments{
				path: "/ipfs/QmRohM66fF9WLqoLTCi6qQEtLiav4JMEsah21nNMeXxEfx",
			},
		},
		{
			name: "profile avatar",
			arguments: arguments{
				path: "/ipfs/QmPkTNGYSUDx5n9hzEDgM19xd2aRTZMfwCuvhcPk3Qazhh",
			},
		},
		{
			name: "uniswap tokenlists",
			arguments: arguments{
				path: "/ipns/tokens.uniswap.org",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			body, err := ipfsHTTPClient.Fetch(context.TODO(), testcase.arguments.path, ipfs.QuickMode)
			require.NoError(t, err)

			defer lo.Try(body.Close)

			data, err := io.ReadAll(body)
			require.NoError(t, err)
			require.NotEmpty(t, data)
		})
	}
}

func TestStableFetch(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		path string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "character metadata",
			arguments: arguments{
				path: "/ipfs/QmRohM66fF9WLqoLTCi6qQEtLiav4JMEsah21nNMeXxEfx",
			},
		},
		{
			name: "profile avatar",
			arguments: arguments{
				path: "/ipfs/QmPkTNGYSUDx5n9hzEDgM19xd2aRTZMfwCuvhcPk3Qazhh",
			},
		},
		{
			name: "uniswap tokenlists",
			arguments: arguments{
				path: "/ipns/tokens.uniswap.org",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			body, err := ipfsHTTPClient.Fetch(context.TODO(), testcase.arguments.path, ipfs.StableMode)
			defer lo.Try(body.Close)

			data, err := io.ReadAll(body)
			require.NoError(t, err)
			require.NotEmpty(t, data)
		})
	}
}
