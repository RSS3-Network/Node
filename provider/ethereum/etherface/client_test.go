package etherface_test

import (
	"context"
	"sync"
	"testing"

	"github.com/rss3-network/node/provider/ethereum/etherface"
	"github.com/stretchr/testify/require"
)

var (
	setupOnce       sync.Once
	etherfaceClient etherface.Client
)

func setup(t *testing.T) {
	setupOnce.Do(func() {
		var err error

		etherfaceClient, err = etherface.NewEtherfaceClient()
		require.NoError(t, err)
	})
}

func TestEtherfaceClient_Lookup(t *testing.T) {
	t.Parallel()

	setup(t)

	type arguments struct {
		hash string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "Lookup Function Signature",
			arguments: arguments{
				hash: "0xb4e4b296",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			functionName, err := etherfaceClient.Lookup(context.TODO(), testcase.arguments.hash)
			require.NoError(t, err)

			require.Equal(t, "matchAskWithTakerBidUsingETHAndWETH", functionName)
		})
	}
}
