package arweave

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/samber/lo"
)

// DefaultClient is the default HTTP client.
var DefaultClient = lo.Must(NewClient())

// Client is the interface that wraps the Fetch method.
type Client interface {
	Fetch(ctx context.Context, id string) (io.ReadCloser, error)
}

// Ensure that client implements Client.
var _ Client = (*client)(nil)

// client is the default implementation of Client with Endpoints
type client struct {
	httpClient   *http.Client
	endpointURLs []*url.URL
}

// Fetch fetches the data of the given hash from arweave network.
func (c *client) Fetch(ctx context.Context, id string) (io.ReadCloser, error) {
	// try to fetch from all endpoints until we get a response (arweave gateways).
	for _, endpoint := range c.endpointURLs {
		response, err := c.fetch(ctx, endpoint, id)
		if err == nil {
			return response, nil
		}
	}

	return nil, fmt.Errorf("failed to fetch %s", id)
}

// fetch fetches the data of the given hash from the given endpoint.
func (c *client) fetch(ctx context.Context, endpointURL *url.URL, id string) (io.ReadCloser, error) {
	requestURL := endpointURL
	requestURL.Path = id

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	return response.Body, nil
}

// ClientOption is the type of the options passed to NewClient.
type ClientOption func(*client) error

// WithEndpointURLs WithHTTPClient sets the HTTP client with arweave endpoints.
func WithEndpointURLs(endpoints []string) ClientOption {
	return func(h *client) error {
		for _, endpoint := range endpoints {
			endpointURL, err := url.Parse(endpoint)
			if err != nil {
				return fmt.Errorf("invalid endpoint %s: %w", endpoint, err)
			}

			h.endpointURLs = append(h.endpointURLs, endpointURL)
		}

		return nil
	}
}

// NewClient creates a new Client with the given options.
func NewClient(options ...ClientOption) (Client, error) {
	client := client{
		httpClient:   http.DefaultClient,
		endpointURLs: make([]*url.URL, 0),
	}

	if err := WithEndpointURLs(EndpointStrings())(&client); err != nil {
		return nil, fmt.Errorf("apply default option: %w", err)
	}

	for _, option := range options {
		if err := option(&client); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &client, nil
}
