package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-playground/form/v4"
)

const (
	// EndpointMainnet https://www.thehubble.xyz/intro/hubble.html
	EndpointMainnet = "https://nemes.farcaster.xyz:2281"

	DefaultTimeout  = 5 * time.Second
	DefaultAttempts = 3
)

var _ Client = (*client)(nil)

type Client interface {
	// GetCastsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/casts.md
	GetCastsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken string) (*MessageResponse, error)
	GetCastByFidAndHash(ctx context.Context, fid *int64, hash string) (*Message, error)
	// GetVerificationsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/verification.md
	GetVerificationsByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error)
	// GetUserNameProofsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/usernameproof.md
	GetUserNameProofsByFid(ctx context.Context, fid *int64) (*ProofResponse, error)
	// GetUserDataByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/userdata.md
	GetUserDataByFid(ctx context.Context, fid *int64, pageToken string) (*MessageResponse, error)
	GetUserDataByFidAndType(ctx context.Context, fid *int64, userDataType string) (*Message, error)
	// GetEvents https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/events.md
	GetEvents(ctx context.Context, fromEventID *int64) (*EventResponse, error)
	// GetFids https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/fids.md
	GetFids(ctx context.Context, reverse bool, pageSize *int) (*FidResponse, error)
	// GetReactionsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/reactions.md
	GetReactionsByFid(ctx context.Context, fid *int64, reverse bool, pageSize *int, pageToken, reactionType string) (*MessageResponse, error)
	GetReaction(ctx context.Context, fid, targetFid *int64, targetHash, reactionType string) (*Message, error)
}

type client struct {
	endpointURL *url.URL
	httpClient  *http.Client
	attempts    uint
}

type farcasterQuery struct {
	Fid          *int64 `form:"fid,omitempty"`
	TargetFid    *int64 `form:"target_fid,omitempty"`
	Hash         string `form:"hash,omitempty"`
	TargetHash   string `form:"target_hash,omitempty"`
	FromEventID  string `form:"from_event_id,omitempty"`
	Reverse      bool   `form:"reverse,omitempty"`
	PageSize     *int   `form:"pageSize,omitempty"`
	PageToken    string `form:"pageToken,omitempty"`
	UserDataType string `form:"user_data_type,omitempty"`
	ReactionType string `form:"reaction_type,omitempty"`
}

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

func (c *client) GetUserNameProofsByFid(ctx context.Context, fid *int64) (*ProofResponse, error) {
	var response ProofResponse

	if err := c.call(ctx, "/v1/userNameProofsByFid", farcasterQuery{
		Fid: fid,
	}, &response); err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	return &response, nil
}

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
	encoder := form.NewEncoder()

	values, err := encoder.Encode(query)

	if err != nil {
		return fmt.Errorf("build params %w", err)
	}

	if err := c.fetch(ctx, fmt.Sprintf("%s?%s", path, values.Encode()), &response); err != nil {
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

func ExtractEventIDToTimestamp(eventID int64) int64 {
	binaryEventID := fmt.Sprintf("%b", eventID)
	binaryTimestamp := binaryEventID[:len(binaryEventID)-SequenceBits]

	decimalTimestamp := 0
	for _, digit := range binaryTimestamp {
		decimalTimestamp = decimalTimestamp*2 + int(digit-'0')
	}

	timestampWithEpoch := int64(decimalTimestamp) + FarcasterEpoch

	return timestampWithEpoch
}

func ConvertTimestampToEventID(timestamp int64) uint64 {
	timestampWithoutEpoch := timestamp - FarcasterEpoch
	binaryTimestamp := fmt.Sprintf("%b", timestampWithoutEpoch)

	var eventID int64

	// Calculate the eventID by combining the binary timestamp and a sequence (in this case, 0).
	for i := 0; i < len(binaryTimestamp); i++ {
		eventID = eventID*2 + int64(binaryTimestamp[i]-'0')
	}

	// Add a sequence (in this case, 0) to the eventID.
	eventID <<= SequenceBits

	return uint64(eventID)
}

func CovertFarcasterTimeToTimestamp(timestamp int64) int64 {
	return timestamp*1000 + FarcasterEpoch
}

type ClientOption func(client *client) error

func NewClient(endpoint string, options ...ClientOption) (Client, error) {
	var (
		instance = client{
			httpClient: &http.Client{
				Timeout: DefaultTimeout,
			},
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
