package mastodon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/go-playground/form/v4"
	"github.com/rss3-network/node/provider/activitypub"
	"go.uber.org/zap"
)

const (
	DefaultTimeout  = 5 * time.Second
	DefaultAttempts = 3
)

var _ Client = (*client)(nil)

type Client interface {
	requestActivityPubData(ctx context.Context) (*activitypub.MessageResponse, error)
}

type client struct {
	endpointURL *url.URL
	httpClient  *http.Client
	encoder     *form.Encoder
	attempts    uint
}

// GetCastsByFid
func (c *client) requestActivityPubData(ctx context.Context) (*activitypub.MessageResponse, error) {
	if err := c.call(ctx, "/activitypub/data", activitypub.MastodonQuery{}, nil); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	// return &response, nil
	return nil, nil
}

func (c *client) call(ctx context.Context, path string, query activitypub.MastodonQuery, response any) error {
	values, err := c.encoder.Encode(query)

	if err != nil {
		return fmt.Errorf("build params %w", err)
	}

	onRetry := retry.OnRetry(func(n uint, err error) {
		zap.L().Error("fetch activitypub request, retrying", zap.Error(err), zap.Uint("attempts", n))
	})

	retryableFunc := func() error {
		err = c.fetch(ctx, fmt.Sprintf("%s?%s", path, values.Encode()), &response)
		if err != nil {
			var httpErr *HTTPError

			// If the error is an HTTP error and the status code is 4xx, we will not retry.
			if errors.As(err, &httpErr) && httpErr.StatusCode >= http.StatusBadRequest && httpErr.StatusCode < http.StatusInternalServerError {
				zap.L().Warn("failed to fetch activitypub request, will not retry", zap.Error(err), zap.Int("status.code", httpErr.StatusCode))

				return retry.Unrecoverable(err)
			}
		}

		return err
	}

	if err = retry.Do(retryableFunc, retry.Delay(time.Second), retry.Attempts(c.attempts), onRetry); err != nil {
		return fmt.Errorf("call: %w", err)
	}

	return nil
}

type HTTPError struct {
	StatusCode int
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func (c *client) fetch(ctx context.Context, path string, result any) error {
	httpErr := &HTTPError{}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.endpointURL, path), nil)
	if err != nil {
		httpErr.Message = fmt.Sprintf("create request: %s", err.Error())

		return httpErr
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		httpErr.Message = fmt.Sprintf("do request: %s", err.Error())

		return httpErr
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		httpErr.StatusCode = response.StatusCode
		httpErr.Message = fmt.Sprintf("unexpected status: %s", response.Status)

		return httpErr
	}

	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		httpErr.Message = fmt.Sprintf("failed to decode response: %s", err.Error())

		return httpErr
	}

	return nil
}

type ClientOption func(client *client) error

func NewClient(endpoint string) (Client, error) {
	var (
		instance = client{
			httpClient: &http.Client{
				Timeout: DefaultTimeout,
			},
			encoder:  form.NewEncoder(),
			attempts: DefaultAttempts,
		}
		err error
	)

	if instance.endpointURL, err = url.Parse(endpoint); err != nil {
		return nil, fmt.Errorf("parse endpoint: %w", err)
	}

	// for _, option := range options {
	//	if err = option(&instance); err != nil {
	//		return nil, fmt.Errorf("apply options: %w", err)
	//	}
	//}

	return &instance, nil
}
