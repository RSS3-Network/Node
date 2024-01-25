package iqwiki

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

//go:generate go run --mod=mod github.com/Khan/genqlient

const Endpoint = "https://graph.everipedia.org/graphql"

var _ graphql.Doer = (*client)(nil)

type client struct {
	httpClient *http.Client
}

func (c *client) Do(request *http.Request) (*http.Response, error) {
	return c.httpClient.Do(request)
}

func NewClient(endpoint string) (graphql.Client, error) {
	instance := client{
		httpClient: http.DefaultClient,
	}

	return graphql.NewClient(endpoint, &instance), nil
}
