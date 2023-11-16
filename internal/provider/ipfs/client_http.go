package ipfs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/naturalselectionlabs/rss3-node/common/syncx"
)

var _ Client = (*httpClient)(nil)

// Mode is the mode of fetching data from IPFS gateways.
type Mode string

const (
	QuickMode  Mode = "quick"
	StableMode Mode = "stable"
)

// DefaultTimeout is the default timeout for http client.
// DefaultRetryCount is the default retry count for http client.
const (
	DefaultTimeout    = 3 * time.Second
	DefaultRetryCount = 3
)

// httpClient is an implementation of Client interface
// that uses HTTP protocol to fetch data from IPFS gateways.
// retryCount is the number of retries for each gateway.
// httpClient is a wrapper around http.Client.
// endpoints is a list of IPFS endpoints.
type httpClient struct {
	retryCount uint
	httpClient http.Client
	endpoints  []string
}

// Fetch fetches data from IPFS gateways.
// path is the path of the ipfs data, such as /ipfs/Qm... or /ipns/Qm...
// mode is the mode of fetching data from IPFS gateways, there are two modes: QuickMode and StableMode.
func (h *httpClient) Fetch(ctx context.Context, path string, mode Mode) (io.ReadCloser, error) {
	var data io.ReadCloser

	retryableFunc := func() error {
		var err error

		switch mode {
		case QuickMode:
			data, err = h.quickFetch(ctx, path)
		case StableMode:
			data, err = h.stableFetch(ctx, path)
		}

		return err
	}

	if err := retry.Do(retryableFunc, retry.Attempts(h.retryCount)); err != nil {
		return nil, fmt.Errorf("retry error: %w", err)
	}

	return data, nil
}

// quickFetch fetches ipfs data through concurrent query gateways,
// and exit once there is a return.
// This mode returns quickly, but consumes redundant requests and causes instability.
func (h *httpClient) quickFetch(ctx context.Context, path string) (io.ReadCloser, error) {
	quickGroup := syncx.NewQuickGroup[io.ReadCloser](ctx)

	gateways := make([]string, len(h.endpoints))
	copy(gateways, h.endpoints)

	for _, endpoint := range gateways {
		endpoint := endpoint

		quickGroup.Go(func(ctx context.Context) (io.ReadCloser, error) {
			return h.fetch(ctx, endpoint, path)
		})
	}

	return quickGroup.Wait()
}

// stableFetch fetches ipfs data through serial polling gateways.
// If the first gateway fails, it will continue to query the next gateway.
// This mode is stable, but the response time can be long.
func (h *httpClient) stableFetch(ctx context.Context, path string) (io.ReadCloser, error) {
	gateways := make([]string, len(h.endpoints))
	copy(gateways, h.endpoints)

	for _, endpoint := range gateways {
		data, err := h.fetch(ctx, endpoint, path)
		if err == nil {
			return data, nil
		}
	}

	return nil, fmt.Errorf("all gateways failed")
}

// fetch fetches ipfs data using http protocol.
func (h *httpClient) fetch(ctx context.Context, endpoint string, path string) (io.ReadCloser, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint: %w", err)
	}

	endpointURL.Path = path

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpointURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("new request error: %w", err)
	}

	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return response.Body, nil
}

// monitorIPFSEndpoints monitors the response time of IPFS endpoints.
func (h *httpClient) monitorIPFSEndpoints(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for range ticker.C {
		h.endpoints = CheckIPFSEndpoints(ctx, h.httpClient, h.endpoints)
	}
}

type HTTPClientOption func(*httpClient) error

// WithRetryCount sets the retry count for http client.
func WithRetryCount(retryCount uint) HTTPClientOption {
	return func(h *httpClient) error {
		h.retryCount = retryCount

		return nil
	}
}

// WithTimeout sets the timeout for http client.
func WithTimeout(timeout time.Duration) HTTPClientOption {
	return func(h *httpClient) error {
		h.httpClient.Timeout = timeout

		return nil
	}
}

// WithEndpoints sets the endpoints for ipfs client.
func WithEndpoints(endpoints []string) HTTPClientOption {
	return func(h *httpClient) error {
		h.endpoints = append(h.endpoints, endpoints...)

		return nil
	}
}

// NewHTTPClient creates a new ipfs client.
func NewHTTPClient(options ...HTTPClientOption) (Client, error) {
	ctx := context.Background()

	client := &httpClient{
		retryCount: DefaultRetryCount,
		httpClient: http.Client{
			Timeout: DefaultTimeout,
		},
	}

	client.endpoints, _ = GetIPFSPublicGateways(ctx)

	for _, option := range options {
		if err := option(client); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	if len(client.endpoints) == 0 {
		return nil, fmt.Errorf("no available endpoints")
	}

	go client.monitorIPFSEndpoints(ctx)

	return client, nil
}
