package bundle_test

import (
	"context"
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/arweave/bundle"
	"github.com/samber/lo"
	"github.com/spf13/viper"
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

func TestDecoder(t *testing.T) {
	t.Parallel()
	setup(t)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	type arguments struct {
		transactionID string
	}

	testcases := []struct {
		name      string
		arguments arguments
	}{
		{
			name: "Bundlr transaction",
			arguments: arguments{
				transactionID: "hjyAbCImQgLS7T7Ae8SkgkT-ACmyYnXQk5Gl6LcS66w",
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			response, err := arweaveClient.GetTransactionData(context.Background(), testcase.arguments.transactionID)
			require.NoError(t, err)

			defer lo.Try(response.Close)

			decoder := bundle.NewDecoder(response)

			header, err := decoder.DecodeHeader()
			require.NoError(t, err)

			for _, dataItemInfo := range header.DataItemInfos {
				t.Log(dataItemInfo.ID, dataItemInfo.Size)
			}

			for decoder.Next() {
				dataItem, err := decoder.DecodeDataItem()
				require.NoError(t, err)

				_, err = io.Copy(io.Discard, dataItem)
				require.NoError(t, err)
			}
		})
	}
}
