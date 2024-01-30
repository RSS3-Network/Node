package ipfs_test

import (
	"context"
	"testing"
	"time"

	"github.com/rss3-network/node/provider/ipfs"
	"github.com/stretchr/testify/require"
)

func TestGetGatewayURLs(t *testing.T) {
	t.Parallel()

	type arguments struct {
		ctx         context.Context
		timeout     time.Duration
		gatewayList string
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      require.ValueAssertionFunc
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "Default gateway list",
			arguments: arguments{
				ctx:         context.Background(),
				timeout:     ipfs.DefaultTimeout,
				gatewayList: ipfs.DefaultGatewayList,
			},
			want: func(t require.TestingT, value interface{}, msgAndArgs ...interface{}) {
				gatewayURLs, ok := value.([]string)
				require.True(t, ok)

				require.Greater(t, len(gatewayURLs), 0)
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(testcase.arguments.ctx, testcase.arguments.timeout)
			defer cancel()

			gatewayURLs, err := ipfs.FetchGateways(ctx, testcase.arguments.gatewayList)
			testcase.wantError(t, err)
			testcase.want(t, gatewayURLs)
		})
	}
}
