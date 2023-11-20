package rara_test

import (
	"context"
	"testing"

	"github.com/naturalselectionlabs/rss3-node/provider/rara"
	"github.com/stretchr/testify/require"
)

func Test_Client(t *testing.T) {
	t.Parallel()

	client, err := rara.NewClient(rara.Endpoint)

	require.NoError(t, err, "build RARA client")

	ipfsResponse, err := rara.GetIpfs(context.Background(), client, "QmVNYxrJtCFUZLq7MvXrUNvt6ctH93C7tKyAFYRmX3BJMa")

	require.NoError(t, err)

	require.Equal(t, ipfsResponse.GetUserSpends()[0].IpfsHash, "QmVNYxrJtCFUZLq7MvXrUNvt6ctH93C7tKyAFYRmX3BJMa")
}
