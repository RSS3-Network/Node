package irys_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/rss3-network/node/provider/arweave/bundle/irys"
	"github.com/rss3-network/node/provider/arweave/contract/momoka"
	"github.com/stretchr/testify/require"
)

func TestQueryTransactions(t *testing.T) {
	t.Parallel()

	graphqlClient := graphql.NewClient(irys.EndpointMainnet, http.DefaultClient)

	transactions, err := irys.Transactions(context.Background(), graphqlClient, momoka.AddressesSubmitter, "", irys.DefaultLimit)
	require.NoError(t, err)

	require.Len(t, transactions.Transactions.Edges, irys.DefaultLimit)
}
