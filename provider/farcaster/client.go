package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/go-playground/form/v4"
	"go.uber.org/zap"
)

const (
	EndpointMainnet = "https://nemes.farcaster.xyz:2281" // Public Instances https://www.thehubble.xyz/intro/hubble.html

	DefaultTimeout  = 5 * time.Second
	DefaultAttempts = 5
)

var _ Client = (*client)(nil)

type Client interface {
	GetCastsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken string) (*MessageResponse, error)
	GetCastByFidAndHash(ctx context.Context, fid *int64, hash string) (*Message, error)
	GetVerificationsByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error)
	GetUserNameProofsByFid(ctx context.Context, fid *int64) (*ProofResponse, error)
	GetUserDataByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error)
	GetUserDataByFidAndType(ctx context.Context, fid *int64, userDataType string) (*Message, error)
	GetEvents(ctx context.Context, fromEventID *int64) (*EventResponse, error)
	GetFids(ctx context.Context, reverse bool, pageSize *int) (*FidResponse, error)
	GetReactionsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken, reactionType string) (*MessageResponse, error)
	GetReaction(ctx context.Context, fid, targetFid *int64, targetHash, reactionType string) (*Message, error)
}

type client struct {
	endpointURL *url.URL
	httpClient  *http.Client
	encoder     *form.Encoder
	attempts    uint
}

// GetCastsByFid Fetch all casts for authored by an fid.
func (c *client) GetCastsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken string) (*MessageResponse, error) {
	var response MessageResponse

	if err := c.call(ctx, "/v1/castsByFid", farcasterQuery{
		Fid:       fid,
		Reverse:   reverse,
		PageSize:  pageSize,
		PageToken: pageToken,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetCastByFidAndHash Get a cast by its fid and hash.
func (c *client) GetCastByFidAndHash(ctx context.Context, fid *int64, hash string) (*Message, error) {
	var response Message

	if err := c.call(ctx, "/v1/castById", farcasterQuery{
		Fid:  fid,
		Hash: hash,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetVerificationsByFid Get a list of verifications provided by an fid.
func (c *client) GetVerificationsByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error) {
	var response MessageResponse

	if err := c.call(ctx, "/v1/verificationsByFid", farcasterQuery{
		Fid:       fid,
		PageToken: pageToken,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetUserNameProofsByFid Get a list of proofs provided by an fid.
func (c *client) GetUserNameProofsByFid(ctx context.Context, fid *int64) (*ProofResponse, error) {
	var response ProofResponse

	if err := c.call(ctx, "/v1/userNameProofsByFid", farcasterQuery{
		Fid: fid,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetUserDataByFid Get UserData for an fid.
func (c *client) GetUserDataByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error) {
	var response MessageResponse

	if err := c.call(ctx, "/v1/userDataByFid", farcasterQuery{
		Fid:       fid,
		PageToken: pageToken,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetUserDataByFidAndType Get UserData by an fid and user data type.
func (c *client) GetUserDataByFidAndType(ctx context.Context, fid *int64, userDataType string) (*Message, error) {
	var response Message

	if err := c.call(ctx, "/v1/userDataByFid", farcasterQuery{
		Fid:          fid,
		UserDataType: userDataType,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetEvents Get a page of Hub events.
func (c *client) GetEvents(ctx context.Context, fromEventID *int64) (*EventResponse, error) {
	var (
		response EventResponse
		query    farcasterQuery
	)

	if fromEventID != nil {
		query.FromEventID = strconv.FormatInt(*fromEventID, 10)
	}

	if err := c.call(ctx, "/v1/events", query, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetFids Get a list of all the fids.
func (c *client) GetFids(ctx context.Context, reverse bool, pageSize *int) (*FidResponse, error) {
	var response FidResponse

	if err := c.call(ctx, "/v1/fids", farcasterQuery{
		Reverse:  reverse,
		PageSize: pageSize,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetReactionsByFid Get all reactions by an fid.
func (c *client) GetReactionsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken, reactionType string) (*MessageResponse, error) {
	var response MessageResponse

	if err := c.call(ctx, "/v1/reactionsByFid", farcasterQuery{
		Fid:          fid,
		Reverse:      reverse,
		PageSize:     pageSize,
		PageToken:    pageToken,
		ReactionType: reactionType,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

// GetReaction Get a reaction by its created fid and target Cast.
func (c *client) GetReaction(ctx context.Context, fid, targetFid *int64, targetHash, reactionType string) (*Message, error) {
	var response Message

	if err := c.call(ctx, "/v1/reactionById", farcasterQuery{
		Fid:          fid,
		TargetFid:    targetFid,
		TargetHash:   targetHash,
		ReactionType: reactionType,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

func (c *client) call(ctx context.Context, path string, query farcasterQuery, response any) error {
	values, err := c.encoder.Encode(query)

	if err != nil {
		return fmt.Errorf("build params %w", err)
	}

	onRetry := retry.OnRetry(func(n uint, err error) {
		zap.L().Error("fetch farcaster request", zap.Error(err), zap.Uint("attempts", n))
	})

	retryableFunc := func() error {
		return c.fetch(ctx, fmt.Sprintf("%s?%s", path, values.Encode()), &response)
	}

	if err = retry.Do(retryableFunc, retry.Delay(time.Second), retry.Attempts(c.attempts), onRetry); err != nil {
		return fmt.Errorf("call: %w", err)
	}

	return nil
}

func (c *client) fetch(ctx context.Context, path string, result any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", c.endpointURL, path), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", response.Status)
	}

	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}

// CovertFarcasterTimeToTimestamp Converts a Farcaster seconds timestamp to a Unix milliseconds timestamp.
func CovertFarcasterTimeToTimestamp(timestamp int64) int64 {
	return timestamp + FarcasterEpoch
}

type ClientOption func(client *client) error

func WithAPIKey(apiKey *string) ClientOption {
	return func(h *client) error {
		if apiKey != nil {
			h.httpClient.Transport = NewAuthenticationTransport(*apiKey)
		}

		return nil
	}
}

type AuthenticationTransport struct {
	APIKey string

	roundTripper http.RoundTripper
}

func (a *AuthenticationTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("api_key", a.APIKey)

	return a.roundTripper.RoundTrip(request)
}

func NewAuthenticationTransport(apiKey string) http.RoundTripper {
	return &AuthenticationTransport{
		APIKey:       apiKey,
		roundTripper: http.DefaultTransport,
	}
}

func NewClient(endpoint string, options ...ClientOption) (Client, error) {
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

	for _, option := range options {
		if err = option(&instance); err != nil {
			return nil, fmt.Errorf("apply options: %w", err)
		}
	}

	return &instance, nil
}
