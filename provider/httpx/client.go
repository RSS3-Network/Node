package httpx

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/gabriel-vasile/mimetype"
	"github.com/samber/lo"
)

var (
	ErrorNoResults = errors.New("no results")
)

const (
	DefaultTimeout  = 3 * time.Second
	DefaultAttempts = 3
)

type Client interface {
	Fetch(ctx context.Context, path string) (io.ReadCloser, error)
	GetContentType(ctx context.Context, path string) (string, error)
}

var _ Client = (*httpClient)(nil)

type httpClient struct {
	httpClient *http.Client
	attempts   uint
}

func (h *httpClient) Fetch(ctx context.Context, path string) (readCloser io.ReadCloser, err error) {
	retryableFunc := func() error {
		readCloser, err = h.fetch(ctx, path)
		return err
	}

	retryIfFunc := func(err error) bool {
		return !errors.Is(err, ErrorNoResults)
	}

	if err := retry.Do(retryableFunc, retry.Attempts(h.attempts), retry.RetryIf(retryIfFunc)); err != nil {
		return nil, fmt.Errorf("retry attempts: %w", err)
	}

	return readCloser, nil
}

func (h *httpClient) fetch(ctx context.Context, path string) (io.ReadCloser, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
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

func (h *httpClient) GetContentType(ctx context.Context, uri string) (string, error) {
	// TODO: Support chain:// protocol.
	if strings.HasPrefix(uri, "chain://") {
		return "text/html", nil
	}

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

func NewHTTPClient(options ...ClientOption) (Client, error) {
	instance := httpClient{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		attempts: DefaultAttempts,
	}

	for _, option := range options {
		if err := option(&instance); err != nil {
			return nil, fmt.Errorf("apply options: %w", err)
		}
	}

	return &instance, nil
}

type ClientOption func(*httpClient) error

func WithAttempts(attempts uint) ClientOption {
	return func(h *httpClient) error {
		h.attempts = attempts

		return nil
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(h *httpClient) error {
		h.httpClient.Timeout = timeout

		return nil
	}
}
