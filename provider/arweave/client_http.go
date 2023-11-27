package arweave

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	syncx "github.com/naturalselectionlabs/rss3-node/common/sync"
	"github.com/samber/lo"
)

// DefaultClient is the default HTTP client.
var DefaultClient = lo.Must(NewClient())

// Client is the interface that wraps the Fetch method.
type HTTPClient interface {
	Fetch(ctx context.Context, id string) (io.ReadCloser, error)
}

// Ensure that client implements Client.
var _ HTTPClient = (*httpClient)(nil)

// client is the default implementation of Client with Endpoints
type httpClient struct {
	httpClient   *http.Client
	endpointURLs []*url.URL
	locker       sync.RWMutex
}

// Fetch fetches the data of the given hash from arweave network.
func (h *httpClient) Fetch(ctx context.Context, id string) (io.ReadCloser, error) {
	h.locker.RLock()
	defer h.locker.RUnlock()

	quickGroup := syncx.NewQuickGroup[io.ReadCloser](ctx)

	// try to fetch from all endpoints until we get a response (arweave gateways).
	for _, endpoint := range h.endpointURLs {
		endpoint := endpoint

		quickGroup.Go(func(ctx context.Context) (io.ReadCloser, error) {
			return h.fetch(ctx, endpoint, id)
		})
	}

	result, err := quickGroup.Wait()
	if err != nil {
		return nil, fmt.Errorf("fetch from all endpoints: %w", err)
	}

	return result, nil
}

// fetch fetches the data of the given hash from the given endpoint.
func (h *httpClient) fetch(ctx context.Context, endpointURL *url.URL, id string) (io.ReadCloser, error) {
	gatewaysURL := *endpointURL
	gatewaysURL.Path = id

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, gatewaysURL.String(), nil)
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
type HTTPClientOption func(*httpClient) error

// WithEndpointURLs WithHTTPClient sets the HTTP client with arweave endpoints.
func WithEndpointURLs(endpoints []string) HTTPClientOption {
	return func(h *httpClient) error {
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
func NewClient(options ...HTTPClientOption) (HTTPClient, error) {
	instance := httpClient{
		httpClient:   http.DefaultClient,
		endpointURLs: make([]*url.URL, 0),
	}

	if err := WithEndpointURLs(EndpointStrings())(&instance); err != nil {
		return nil, fmt.Errorf("apply default option: %w", err)
	}

	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &instance, nil
}
