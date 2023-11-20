package bundle_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/provider/arweave/bundle"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestDecoder(t *testing.T) {
	t.Parallel()

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

	// Mock Test
	// arweaveMockEndpoint := viper.GetString("")
	// require.NotEmpty(t, arweaveMockEndpoint)

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			transactionDataURL := fmt.Sprintf("%s%s", "https://arweave.net/", testcase.arguments.transactionID)

			request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, transactionDataURL, nil)
			require.NoError(t, err)

			response, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			defer response.Body.Close()

			decoder := bundle.NewDecoder(response.Body)

			header, err := decoder.DecodeHeader()
			require.NoError(t, err)

			for _, dataItemInfo := range header.DataItemInfos {
				t.Log(dataItemInfo.ID, dataItemInfo.Size)
			}

			for decoder.Next() {
				dataItem, err := decoder.DecodeDataItem()
				require.NoError(t, err)

				t.Log(dataItem)

				_, err = io.Copy(io.Discard, dataItem)
				require.NoError(t, err)
			}
		})
	}
}
