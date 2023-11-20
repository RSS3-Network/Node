package snapshot

import (
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

//go:generate go run --mod=mod github.com/Khan/genqlient@v0.6.0

const (
	EndpointMainnet = "https://hub.snapshot.org/graphql"
	EndpointTestnet = "https://testnet.snapshot.org/graphql"

	Limit = 1000

	ProposalStateClosed = "closed"
)

var _ graphql.Doer = (*client)(nil)

type client struct {
	httpClient *http.Client
	apiKey     string
}

func (c *client) Do(request *http.Request) (*http.Response, error) {
	if c.apiKey != "" {
		request.Header.Set("X-API-KEY", c.apiKey)
	}

	return c.httpClient.Do(request)
}

func NewClient(endpoint string, options ...Option) (graphql.Client, error) {
	instance := client{
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return graphql.NewClient(endpoint, &instance), nil
}
