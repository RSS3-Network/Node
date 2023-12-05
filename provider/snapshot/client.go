package snapshot

import (
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

// Command to generate the client code in operation.go.
//go:generate go run --mod=mod github.com/Khan/genqlient@v0.6.0

const (
	EndpointMainnet = "https://hub.snapshot.org/graphql"
	EndpointTestnet = "https://testnet.snapshot.org/graphql"

	Limit = 1000

	ProposalStateClosed = "closed"
)

// Use _ to make sure we implement the interface.
var _ graphql.Doer = (*client)(nil)

// A graphql.Doer that adds the X-API-KEY header to requests.
type client struct {
	httpClient *http.Client
	apiKey     string
}

// Do implements [graphql.Doer].
func (c *client) Do(request *http.Request) (*http.Response, error) {
	// add the X-API-KEY header if it's set.
	if c.apiKey != "" {
		request.Header.Set("X-API-KEY", c.apiKey)
	}

	// do the request
	return c.httpClient.Do(request)
}

// NewClient a function that can be used to configure a [client] with options.
// Options are applied in the order they are passed to the function
func NewClient(endpoint string, options ...Option) (graphql.Client, error) {
	// build a default client
	instance := client{
		httpClient: http.DefaultClient,
	}

	// Apply options.
	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	// Return the client
	return graphql.NewClient(endpoint, &instance), nil
}
