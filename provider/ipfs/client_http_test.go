package ipfs_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/rss3-network/node/provider/ipfs"
	"github.com/stretchr/testify/require"
)

func TestHttpClient_Fetch(t *testing.T) {
	t.Parallel()

	type arguments struct {
		ctx       context.Context
		options   []ipfs.HTTPClientOption
		path      string
		fetchMode ipfs.FetchMode
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      require.ValueAssertionFunc
		wantError require.ErrorAssertionFunc
	}{
		//{
		//	name: "Crossbell profile of kallydev.csb",
		//	arguments: arguments{
		//		ctx: context.Background(),
		//		options: []ipfs.HTTPClientOption{
		//			ipfs.WithTimeout(10 * time.Second),
		//		},
		//		path:      "/ipfs/QmRohM66fF9WLqoLTCi6qQEtLiav4JMEsah21nNMeXxEfx",
		//		fetchMode: ipfs.FetchModeStable,
		//	},
		//	want: func(t require.TestingT, v interface{}, msgAndArgs ...interface{}) {
		//		length, ok := v.(int64)
		//		require.True(t, ok)
		//		require.Greater(t, length, int64(0))
		//	},
		//	wantError: require.NoError,
		//},
		{
			name: "Crossbell avatar of kallydev.csb",
			arguments: arguments{
				ctx: context.Background(),
				options: []ipfs.HTTPClientOption{
					ipfs.WithTimeout(10 * time.Second),
				},
				path:      "/ipfs/QmPkTNGYSUDx5n9hzEDgM19xd2aRTZMfwCuvhcPk3Qazhh",
				fetchMode: ipfs.FetchModeStable,
			},
			want: func(t require.TestingT, v interface{}, msgAndArgs ...interface{}) {
				length, ok := v.(int64)
				require.True(t, ok)
				require.Greater(t, length, int64(0))
			},
			wantError: require.NoError,
		},
		{
			name: "Token lists of Uniswap",
			arguments: arguments{
				ctx: context.Background(),
				options: []ipfs.HTTPClientOption{
					ipfs.WithTimeout(10 * time.Second),
				},
				path:      "/ipns/tokens.uniswap.org",
				fetchMode: ipfs.FetchModeStable,
			},
			want: func(t require.TestingT, v interface{}, msgAndArgs ...interface{}) {
				length, ok := v.(int64)
				require.True(t, ok)
				require.Greater(t, length, int64(0))
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			httpClient, err := ipfs.NewHTTPClient(testcase.arguments.options...)
			testcase.wantError(t, err)

			result, err := httpClient.Fetch(testcase.arguments.ctx, testcase.arguments.path, testcase.arguments.fetchMode)
			testcase.wantError(t, err)

			length, err := io.Copy(io.Discard, result)
			testcase.wantError(t, err)
			testcase.want(t, length)
		})
	}
}
