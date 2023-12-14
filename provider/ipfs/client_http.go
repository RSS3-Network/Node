package ipfs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/gabriel-vasile/mimetype"
	syncx "github.com/naturalselectionlabs/rss3-node/common/sync"
	"github.com/samber/lo"
)

var (
	ErrorUnsupportedMode = errors.New("unsupported mode")
	ErrorNoResults       = errors.New("no results")
)

// FetchMode is the mode of fetching data from IPFS gateways.
type FetchMode string

const (
	FetchModeQuick  FetchMode = "quick"
	FetchModeStable FetchMode = "stable"

	DefaultTimeout  = 3 * time.Second
	DefaultAttempts = 3
)

type HTTPClient interface {
	Fetch(ctx context.Context, path string, fetchMode FetchMode) (io.ReadCloser, error)
	GetContentType(ctx context.Context, path string) (string, error)
}

var _ HTTPClient = (*httpClient)(nil)

type httpClient struct {
	httpClient *http.Client
	gateways   []string
	attempts   uint
	locker     sync.RWMutex
}

func (h *httpClient) Fetch(ctx context.Context, path string, fetchMode FetchMode) (readCloser io.ReadCloser, err error) {
	retryableFunc := func() error {
		switch fetchMode {
		case FetchModeStable:
			readCloser, err = h.fetchStable(ctx, path)
		case FetchModeQuick:
			readCloser, err = h.fetchQuick(ctx, path)
		default:
			return fmt.Errorf("%w %s", ErrorUnsupportedMode, fetchMode)
		}

		return err
	}

	retryIfFunc := func(err error) bool {
		return !errors.Is(err, ErrorUnsupportedMode)
	}

	if err := retry.Do(retryableFunc, retry.Attempts(h.attempts), retry.RetryIf(retryIfFunc)); err != nil {
		return nil, fmt.Errorf("retry attempts: %w", err)
	}

	return readCloser, nil
}

func (h *httpClient) fetch(ctx context.Context, gateway string, path string) (io.ReadCloser, error) {
	fileURL, err := url.JoinPath(gateway, path)
	if err != nil {
		return nil, fmt.Errorf("invalid gateway and path: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		defer lo.Try(response.Body.Close)

		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	return response.Body, nil
}

func (h *httpClient) fetchStable(ctx context.Context, path string) (readCloser io.ReadCloser, err error) {
	h.locker.RLock()
	defer h.locker.RUnlock()

	for _, gateway := range h.gateways {
		if readCloser, err = h.fetch(ctx, gateway, path); err == nil {
			return readCloser, nil
		}
	}

	return nil, ErrorNoResults
}

func (h *httpClient) fetchQuick(ctx context.Context, path string) (io.ReadCloser, error) {
	h.locker.RLock()
	defer h.locker.RUnlock()

	quickGroup := syncx.NewQuickGroup[io.ReadCloser](ctx)

	for _, gateway := range h.gateways {
		gateway := gateway

		quickGroup.Go(func(ctx context.Context) (io.ReadCloser, error) {
			return h.fetch(ctx, gateway, path)
		})
	}

	result, err := quickGroup.Wait()
	if err != nil {
		return nil, ErrorNoResults
	}

	return result, nil
}

func (h *httpClient) GetContentType(ctx context.Context, uri string) (string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	request.Header.Set("User-Agent", "curl/7.86.0")

	response, err := h.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	mimeType, err := mimetype.DetectReader(response.Body)

	if err != nil || mimeType == nil {
		return "", fmt.Errorf("fail to detect mime type %s", uri)
	}

	return mimeType.String(), nil
}

func NewHTTPClient(options ...HTTPClientOption) (HTTPClient, error) {
	instance := httpClient{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		gateways: DefaultGateways,
		attempts: DefaultAttempts,
	}

	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply options: %w", err)
		}
	}

	return &instance, nil
}

type HTTPClientOption func(*httpClient) error

func WithAttempts(attempts uint) HTTPClientOption {
	return func(h *httpClient) error {
		h.attempts = attempts

		return nil
	}
}

func WithTimeout(timeout time.Duration) HTTPClientOption {
	return func(h *httpClient) error {
		h.httpClient.Timeout = timeout

		return nil
	}
}

func WithGateways(gateways []string) HTTPClientOption {
	return func(h *httpClient) error {
		h.gateways = gateways // Overwrite gateways.

		return nil
	}
}
