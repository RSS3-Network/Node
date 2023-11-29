package farcaster

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	GetCastsByFid(ctx context.Context, fid int64, reverse *bool, pageSize *int, pageToken *string) (*MessageResponse, error)
	GetCastByFidAndHash(ctx context.Context, fid int64, hash string) (*Message, error)
	// GetVerificationsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/verification.md
	GetVerificationsByFid(ctx context.Context, fid int64, pageToken *string) (*MessageResponse, error)
	// GetUserNameProofsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/usernameproof.md
	GetUserNameProofsByFid(ctx context.Context, fid int64) (*ProofResponse, error)
	// GetUserDataByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/userdata.md
	GetUserDataByFid(ctx context.Context, fid int64, pageToken *string) (*MessageResponse, error)
	GetUserDataByFidAndType(ctx context.Context, fid int64, userDataType *string) (*Message, error)
	// GetEvents https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/events.md
	GetEvents(ctx context.Context, fromEventID int64) (*EventResponse, error)
	// GetFids https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/fids.md
	GetFids(ctx context.Context, reverse *bool, pageSize *int) (*FidResponse, error)
	// GetReactionsByFid https://github.com/farcasterxyz/hub-monorepo/blob/main/apps/hubble/www/docs/docs/httpapi/reactions.md
	GetReactionsByFid(ctx context.Context, fid int64, reverse *bool, pageSize *int, pageToken *string, reactionType *string) (*MessageResponse, error)
	GetReaction(ctx context.Context, fid, targetFid int64, targetHash, reactionType string) (*Message, error)
}

type client struct {
	endpointURL *url.URL
	httpClient  *http.Client
	attempts    uint
}

func (c *client) GetCastsByFid(ctx context.Context, fid int64, reverse *bool, pageSize *int, pageToken *string) (*MessageResponse, error) {
	var response MessageResponse

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))

	if pageSize != nil {
		params.Add("pageSize", strconv.Itoa(*pageSize))
	}

	if pageToken != nil {
		params.Add("pageToken", *pageToken)
	}

	if reverse != nil {
		params.Add("reverse", strconv.FormatBool(*reverse))
	}

	str := fmt.Sprintf("/v1/castsByFid?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetCastByFidAndHash(ctx context.Context, fid int64, hash string) (*Message, error) {
	var response Message

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))
	params.Add("hash", hash)

	str := fmt.Sprintf("/v1/castById?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetVerificationsByFid(ctx context.Context, fid int64, pageToken *string) (*MessageResponse, error) {
	var response MessageResponse

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))

	if pageToken != nil {
		params.Add("nextPageToken", *pageToken)
	}

	str := fmt.Sprintf("/v1/verificationsByFid?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetUserNameProofsByFid(ctx context.Context, fid int64) (*ProofResponse, error) {
	var response ProofResponse

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))

	str := fmt.Sprintf("/v1/userNameProofsByFid?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetUserDataByFid(ctx context.Context, fid int64, pageToken *string) (*MessageResponse, error) {
	var response MessageResponse

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))

	if pageToken != nil {
		params.Add("nextPageToken", *pageToken)
	}

	str := fmt.Sprintf("/v1/userDataByFid?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetUserDataByFidAndType(ctx context.Context, fid int64, userDataType *string) (*Message, error) {
	var response Message

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))

	params.Add("user_data_type", *userDataType)

	str := fmt.Sprintf("/v1/userDataByFid?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetEvents(ctx context.Context, fromEventID int64) (*EventResponse, error) {
	var response EventResponse

	params := url.Values{}

	params.Add("from_event_id", strconv.FormatInt(fromEventID, 10))

	str := fmt.Sprintf("/v1/events?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetFids(ctx context.Context, reverse *bool, pageSize *int) (*FidResponse, error) {
	var response FidResponse

	params := url.Values{}

	if reverse != nil {
		params.Add("reverse", strconv.FormatBool(*reverse))
	}

	if pageSize != nil {
		params.Add("pageSize", strconv.Itoa(*pageSize))
	}

	str := fmt.Sprintf("/v1/fids?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetReactionsByFid(ctx context.Context, fid int64, reverse *bool, pageSize *int, pageToken *string, reactionType *string) (*MessageResponse, error) {
	var response MessageResponse

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))

	if pageSize != nil {
		params.Add("pageSize", strconv.Itoa(*pageSize))
	}

	if pageToken != nil {
		params.Add("pageToken", *pageToken)
	}

	if reverse != nil {
		params.Add("reverse", strconv.FormatBool(*reverse))
	}

	if reactionType != nil {
		params.Add("reaction_type", *reactionType)
	}

	str := fmt.Sprintf("/v1/reactionsByFid?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
}

func (c *client) GetReaction(ctx context.Context, fid, targetFid int64, targetHash, reactionType string) (*Message, error) {
	var response Message

	params := url.Values{}

	params.Add("fid", strconv.FormatInt(fid, 10))
	params.Add("target_fid", strconv.FormatInt(targetFid, 10))
	params.Add("target_hash", targetHash)
	params.Add("reaction_type", reactionType)

	str := fmt.Sprintf("/v1/reactionById?%s", params.Encode())

	if err := c.fetch(ctx, str, &response); err != nil {
		return nil, fmt.Errorf("call: %w", err)
	}

	return &response, nil
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
