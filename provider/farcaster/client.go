package farcaster

import (
	"context"
	"encoding/json"
	"errors"
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
	// Public Instances https://www.thehubble.xyz/intro/hubble.html
	EndpointMainnet = "https://nemes.farcaster.xyz:2281"

	DefaultTimeout  = 5 * time.Second
	DefaultAttempts = 3
	// https://github.com/farcasterxyz/hub-monorepo/blob/77ff79ed804104956eb153247c22c00099c7b122/packages/core/src/time.ts#L4
	FarcasterEpoch = 1609459200 // January 1, 2021 UTC
	SequenceBits   = 12
)

var _ Client = (*client)(nil)

type Client interface {
	GetCastsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken string) (*MessageResponse, error)
	GetCastByFidAndHash(ctx context.Context, fid *int64, hash string) (*Message, error)
	GetVerificationsByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error)
	GetUserNameProofsByFid(ctx context.Context, fid *int64) (*ProofResponse, error)
	GetUserNameProofByName(ctx context.Context, name string) (*UserNameProof, error)
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

// GetUserNameProofByName Get a proof by a user name.
func (c *client) GetUserNameProofByName(ctx context.Context, name string) (*UserNameProof, error) {
	var response UserNameProof

	if err := c.call(ctx, "/v1/userNameProofByName", farcasterQuery{
		Name: name,
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

// GetUserDataByFidAndType Get UserData by an fid and user data typex.
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
		zap.L().Error("fetch farcaster request, retrying", zap.Error(err), zap.Uint("attempts", n))
	})

	retryableFunc := func() error {
		err = c.fetch(ctx, fmt.Sprintf("%s?%s", path, values.Encode()), &response)
		if err != nil {
			var httpErr *HTTPError

			// If the error is an HTTP error and the status code is 4xx, we will not retry.
			if errors.As(err, &httpErr) && httpErr.StatusCode >= http.StatusBadRequest && httpErr.StatusCode < http.StatusInternalServerError {
				zap.L().Warn("failed to fetch farcaster request, will not retry", zap.Error(err), zap.Int("status.code", httpErr.StatusCode))

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

// CovertFarcasterTimeToTimestamp Converts a Farcaster seconds timestamp to a Unix timestamp.
func CovertFarcasterTimeToTimestamp(timestamp int64) int64 {
	return timestamp + FarcasterEpoch
}

// CovertTimestampToFarcasterTime Converts a Unix timestamp to a Farcaster seconds timestamp.
func CovertTimestampToFarcasterTime(timestamp int64) uint32 {
	return uint32(timestamp) - FarcasterEpoch
}

// ConvertEventIDToTimestampMilli Convert the timestampMilli from a farcaster event ID.
func ConvertEventIDToTimestampMilli(eventID uint64) uint64 {
	timestampMilliMask := ^uint64((1 << SequenceBits) - 1)
	timestampMilli := (eventID & timestampMilliMask) >> SequenceBits

	return timestampMilli + FarcasterEpoch*1000
}

// ConvertTimestampMilliToEventID Convert a timestampMilli to a farcaster event ID.
func ConvertTimestampMilliToEventID(timestamp int64) uint64 {
	timestampWithoutEpoch := timestamp - FarcasterEpoch*1000
	eventID := timestampWithoutEpoch << SequenceBits

	return uint64(eventID)
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
