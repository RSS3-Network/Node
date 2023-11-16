package ipfs_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/naturalselectionlabs/rss3-node/internal/provider/ipfs"
	"github.com/stretchr/testify/require"
)

func TestCheckIPFSEndpoints(t *testing.T) {
	t.Parallel()

	gateways, err := ipfs.GetIPFSPublicGateways(context.Background())
	require.NoError(t, err)
	require.Greater(t, len(gateways), 0)
	t.Log(gateways)

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	endpoints := ipfs.CheckIPFSEndpoints(context.Background(), client, gateways)
	require.Equal(t, len(endpoints), len(gateways))
	t.Log(endpoints)
}
