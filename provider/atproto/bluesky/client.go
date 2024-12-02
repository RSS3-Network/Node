package bluesky

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/araddon/dateparse"
	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
	"github.com/bluesky-social/indigo/repo"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/go-playground/form/v4"
	"github.com/ipfs/go-cid"
	at "github.com/rss3-network/node/provider/atproto"
	"github.com/samber/lo"
	cbg "github.com/whyrusleeping/cbor-gen"
	"go.uber.org/zap"
)

const (
	BskyEndpoint     = "https://bsky.social"
	BskySubscribeURI = "wss://bsky.network/xrpc/com.atproto.sync.subscribeRepos"

	SyncListReposLimit = 50
)

type Client struct {
	username       string
	password       string
	filter         []string
	timestampStart int64
	mutex          sync.RWMutex
	defaultClient  *XrpcClient
	cacheClient    map[string]*XrpcClient
	encoder        *form.Encoder
	httpClient     *http.Client
}

type XrpcClient struct {
	*xrpc.Client
	createdAt time.Time
}

// SyncListRepos fetches a paginated list of repository updates.
// Parameters:
// - cursor: Pagination token, empty for first page
// - limit: Maximum number of repos to fetch
// Returns list of messages and the next cursor for pagination.
func (c *Client) SyncListRepos(ctx context.Context, cursor string, limit int64) ([]*at.Message, *string, error) {
	// Get default client for bsky.social endpoint
	client, err := c.GetXrpcClient(ctx, BskyEndpoint)
	if err != nil {
		zap.L().Error("get xrpc client failed", zap.Error(err))

		return nil, nil, fmt.Errorf("get xrpc client failed: %w", err)
	}

	// Fetch the list of repos
	results, err := atproto.SyncListRepos(ctx, client.Client, cursor, limit)
	if err != nil {
		zap.L().Error("sync list repos failed", zap.Error(err))

		return nil, nil, fmt.Errorf("sync list repos failed: %w", err)
	}

	data := make([]*at.Message, 0)

	// Process each repository to extract its records
	for _, rep := range results.Repos {
		feed, err := c.SyncGetRepo(ctx, rep)
		if err != nil {
			zap.L().Error("sync get repo failed", zap.Error(err))

			return nil, nil, fmt.Errorf("sync get repo failed: %w", err)
		}

		if len(feed) == 0 {
			continue
		}

		data = append(data, feed...)
	}

	return data, results.Cursor, nil
}

// SyncGetRepo retrieves all records from a repository snapshot.
// Parameters:
// - repoData: Repository metadata including DID and status
// Returns list of messages containing the repository records.
func (c *Client) SyncGetRepo(ctx context.Context, repoData *atproto.SyncListRepos_Repo) ([]*at.Message, error) {
	if repoData.Active != nil && !lo.FromPtr(repoData.Active) {
		return nil, nil
	}

	// Parse the repository DID
	did, err := syntax.ParseDID(repoData.Did)
	if err != nil {
		zap.L().Error("parse DID failed", zap.Error(err))

		return nil, fmt.Errorf("parse DID: %w", err)
	}

	// Get authenticated client for the endpoint
	client, err := c.GetXrpcClient(ctx, c.LookupDIDEdnpoint(ctx, did))
	if err != nil {
		zap.L().Error("create xrpc client failed", zap.Error(err))

		return nil, fmt.Errorf("create xrpc client: %w", err)
	}

	// Get user handle from profile
	handle, err := c.GetHandle(ctx, did)
	if err != nil {
		zap.L().Error("get profile failed", zap.Error(err))

		return nil, fmt.Errorf("get profile: %w", err)
	}

	// Fetch repo data
	repoBytes, err := atproto.SyncGetRepo(ctx, client.Client, did.String(), "")
	if err != nil {
		zap.L().Error("get repo failed", zap.Error(err))

		return nil, fmt.Errorf("get repo: %w", err)
	}

	// Parse repo data in CAR format
	feedList, err := c.ParseCARList(ctx, did, handle, repoBytes)
	if err != nil {
		zap.L().Error("parse car list failed", zap.Error(err))

		return nil, fmt.Errorf("parse car list: %w", err)
	}

	return feedList, nil
}

// ParseCARList processes repository data in CAR format and extracts messages.
// Parameters:
// - did: Repository owner's DID
// - handle: User's handle
// - repoBytes: Raw CAR format repository data
// Returns list of processed messages from the repository.
func (c *Client) ParseCARList(ctx context.Context, did syntax.DID, handle string, repoBytes []byte) ([]*at.Message, error) {
	r, err := repo.ReadRepoFromCar(ctx, bytes.NewReader(repoBytes))
	if err != nil {
		zap.L().Error("read repo from car failed", zap.Error(err))

		return nil, fmt.Errorf("read repo from car: %w", err)
	}

	recList := make([]*at.Message, 0)

	// Iterate over all records
	err = r.ForEach(ctx, "", func(path string, _ cid.Cid) error {
		// Parse path to get collection and rkey
		collection, rkey := c.ParsePath(path)

		if !lo.Contains(c.filter, collection) {
			return nil
		}

		// Get record from repo
		_, rec, err := r.GetRecord(ctx, path)
		if err != nil {
			zap.L().Error("get record failed", zap.Error(err))

			return fmt.Errorf("get record: %w", err)
		}

		message := &at.Message{
			URI:        c.BuildURI(did, collection, rkey),
			Did:        did,
			Handle:     handle,
			Collection: collection,
			Rkey:       rkey,
		}

		// Parse record and update message
		if err = c.ParseRecord(ctx, rec, message); err != nil {
			zap.L().Error("parse record failed", zap.Error(err))

			return nil
		}

		if message.CreatedAt.Unix() >= c.timestampStart {
			recList = append(recList, message)
		}

		return nil
	})

	if err != nil {
		zap.L().Error("iterate over all records failed", zap.Error(err))

		return nil, err
	}

	return recList, nil
}

// GetRepoRecord retrieves a specific record from a repository.
// Parameters:
// - repo: The DID of the repository
// - path: The path in format "collection/rkey"
// Returns the message containing the record data or an error.
func (c *Client) GetRepoRecord(ctx context.Context, repo string, path string) (*at.Message, error) {
	message, rec, err := c.GetRecord(ctx, repo, path)
	if err != nil {
		zap.L().Error("get record failed", zap.Error(err))

		return nil, fmt.Errorf("get record: %w", err)
	}

	err = c.ParseRecord(ctx, rec, message)
	if err != nil {
		zap.L().Error("parse record failed", zap.Error(err))

		return nil, nil
	}

	if message.CreatedAt.Unix() >= c.timestampStart {
		return message, nil
	}

	return nil, nil
}

func (c *Client) GetRecord(ctx context.Context, repo string, path string) (*at.Message, cbg.CBORMarshaler, error) {
	// Parse the path into collection and rkey components
	collection, rkey := c.ParsePath(path)

	if !lo.Contains(c.filter, collection) {
		return nil, nil, nil
	}

	// Parse and validate the DID
	did, err := syntax.ParseDID(repo)
	if err != nil {
		zap.L().Error("parse DID failed", zap.Error(err), zap.String("repo", repo))

		return nil, nil, fmt.Errorf("parse DID: %w", err)
	}

	// Get authenticated client
	client, err := c.GetXrpcClient(ctx, c.LookupDIDEdnpoint(ctx, did))
	if err != nil {
		zap.L().Error("create xrpc client failed", zap.Error(err))

		return nil, nil, fmt.Errorf("create xrpc client: %w", err)
	}

	// Fetch the handle associated with the DID
	handle, err := c.GetHandle(ctx, did)
	if err != nil {
		zap.L().Error("get profile failed", zap.Error(err))

		return nil, nil, fmt.Errorf("get profile: %w", err)
	}

	message := &at.Message{
		URI:        c.BuildURI(did, collection, rkey),
		Did:        did,
		Handle:     handle,
		Collection: collection,
		Rkey:       rkey,
	}

	// Retrieve the record
	resp, err := atproto.RepoGetRecord(ctx, client.Client, "", collection, repo, rkey)
	if err != nil {
		// Handle CID too short error gracefully
		if strings.Contains(err.Error(), cid.ErrCidTooShort.Error()) || strings.Contains(err.Error(), "XRPC ERROR 400") {
			return message, nil, nil
		}

		return nil, nil, fmt.Errorf("repo get record: %w", err)
	}

	if resp != nil && resp.Value != nil {
		return message, resp.Value.Val, nil
	}

	return message, nil, nil
}

// GetHandle retrieves the handle (username) for a given DID.
// Parameters:
// - did: User's decentralized identifier
// Returns the user's handle or an error.
func (c *Client) GetHandle(ctx context.Context, did syntax.DID) (string, error) {
	client, err := c.GetXrpcClient(ctx, c.LookupDIDEdnpoint(ctx, did))
	if err != nil {
		return "", fmt.Errorf("get xrpc client: %w", err)
	}

	resp, err := bsky.ActorGetProfile(ctx, client.Client, did.String())
	if err != nil {
		if strings.Contains(err.Error(), "XRPC ERROR 400") {
			return "", nil
		}

		return "", fmt.Errorf("profile get: %w", err)
	}

	if resp == nil {
		return "", nil
	}

	return resp.Handle, nil
}

// ParseRecord processes different types of records and updates the message.
// Parameters:
// - rec: CBOR marshaled record to parse
// - message: Message struct to be updated with record data
// Returns error if record parsing fails.
func (c *Client) ParseRecord(ctx context.Context, rec cbg.CBORMarshaler, message *at.Message) error {
	var (
		err       error
		createdAt string
	)

	switch rec := rec.(type) {
	case *bsky.FeedPost:
		message.Feed = rec
		createdAt = rec.CreatedAt

		if rec.Reply != nil && rec.Reply.Parent != nil {
			message.RefMessage, err = c.ParseRepoStrongRef(ctx, rec.Reply.Parent)
			if err != nil {
				zap.L().Error("parse repo strong ref failed", zap.Error(err))

				return fmt.Errorf("parse repo strong ref: %w", err)
			}
		}
	case *bsky.ActorProfile:
		message.Profile = rec

		if rec.CreatedAt != nil {
			createdAt = *rec.CreatedAt
		}
	case *bsky.FeedRepost:
		message.RefMessage, err = c.ParseRepoStrongRef(ctx, rec.Subject)
		if err != nil {
			zap.L().Error("parse repo strong ref failed", zap.Error(err))

			return fmt.Errorf("parse repo strong ref: %w", err)
		}

		createdAt = rec.CreatedAt
	case *bsky.FeedLike:
		message.RefMessage, err = c.ParseRepoStrongRef(ctx, rec.Subject)
		if err != nil {
			zap.L().Error("parse repo strong ref failed", zap.Error(err))

			return fmt.Errorf("parse repo strong ref: %w", err)
		}

		createdAt = rec.CreatedAt
	}

	timestamp, err := dateparse.ParseAny(createdAt)
	if err != nil {
		return fmt.Errorf("parse timestamp: %w", err)
	}

	message.CreatedAt = timestamp

	return nil
}

func (c *Client) ParseRepoStrongRef(ctx context.Context, ref *atproto.RepoStrongRef) (*at.Message, error) {
	if ref == nil {
		return nil, fmt.Errorf("repo strong ref is nil")
	}

	repo, path := c.ParseURI(ref.Uri)

	target, rec, err := c.GetRecord(ctx, repo, path)
	if err != nil {
		zap.L().Error("get record failed", zap.Error(err))

		return nil, fmt.Errorf("get record: %w", err)
	}

	if rec == nil {
		return nil, nil
	}

	switch rec := rec.(type) {
	case *bsky.FeedPost:
		target.Feed = rec
	default:
		return nil, fmt.Errorf("unsupported record type: %T", rec)
	}

	return target, nil
}

// ParseURI splits an AT Protocol URI into components.
// Parameters:
// - uri: Complete AT Protocol URI (format: at://did/collection/rkey)
// Returns the repository DID and the path.
func (c *Client) ParseURI(uri string) (string, string) {
	// Remove "at://" prefix
	uri = strings.TrimPrefix(uri, "at://")

	// Split into DID and collection/rkey parts
	parts := strings.SplitN(uri, "/", 2)
	if len(parts) != 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

// ParsePath splits a path string into collection and rkey components.
// Parameters:
// - path: String in format "collection/rkey"
// Returns collection and rkey as separate strings.
func (c *Client) ParsePath(path string) (string, string) {
	parts := strings.Split(path, "/")

	if len(parts) < 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

// BuildURI constructs an AT Protocol URI from its components.
// Parameters:
// - did: Repository owner's DID
// - collection: Record collection identifier
// - rkey: Record key
// Returns complete URI in format "at://did/collection/rkey".
func (c *Client) BuildURI(did syntax.DID, collection string, rkey string) string {
	return fmt.Sprintf("at://%s/%s/%s", did, collection, rkey)
}

// GetXrpcClient returns a cached or new authenticated XRPC client.
// Parameters:
// - endpoint: API endpoint URL to connect to
// Returns authenticated client or error if authentication fails.
func (c *Client) GetXrpcClient(ctx context.Context, endpoint string) (*XrpcClient, error) {
	if endpoint == BskyEndpoint {
		if c.defaultClient != nil && time.Since(c.defaultClient.createdAt) < time.Hour*2 {
			return c.defaultClient, nil
		}

		client, err := c.createAndAuthenticateClient(ctx, BskyEndpoint)
		if err != nil {
			zap.L().Error("server create session failed", zap.Error(err))
			return nil, err
		}

		c.defaultClient = client

		return c.defaultClient, nil
	}

	c.mutex.RLock()
	cacheClient, ok := c.cacheClient[endpoint]
	c.mutex.RUnlock()

	if ok && time.Since(cacheClient.createdAt) < time.Hour*2 {
		return cacheClient, nil
	}

	client, err := c.createAndAuthenticateClient(ctx, endpoint)
	if err != nil {
		return c.GetXrpcClient(ctx, BskyEndpoint)
	}

	c.mutex.Lock()
	c.cacheClient[endpoint] = client
	c.mutex.Unlock()

	return client, nil
}

func (c *Client) createAndAuthenticateClient(ctx context.Context, endpoint string) (*XrpcClient, error) {
	client := &xrpc.Client{
		Host:   endpoint,
		Client: http.DefaultClient,
	}

	// Create context with 5 second timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Authenticate the client
	auth, err := atproto.ServerCreateSession(timeoutCtx, client, &atproto.ServerCreateSession_Input{
		Identifier: c.username,
		Password:   c.password,
	})
	if err != nil {
		return nil, fmt.Errorf("server create session: %w", err)
	}

	// Update client with authentication info
	client.Auth = &xrpc.AuthInfo{
		AccessJwt:  auth.AccessJwt,
		RefreshJwt: auth.RefreshJwt,
		Handle:     auth.Handle,
		Did:        auth.Did,
	}

	return &XrpcClient{
		Client:    client,
		createdAt: time.Now(),
	}, nil
}

func (c *Client) LookupDIDEdnpoint(ctx context.Context, did syntax.DID) string {
	// Creates a new BaseDirectory with default settings.
	baseDirectory := &identity.BaseDirectory{
		PLCURL: identity.DefaultPLCURL,
		HTTPClient: http.Client{
			Timeout: time.Second * 15,
		},
		Resolver: net.Resolver{
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{Timeout: time.Second * 5}
				return d.DialContext(ctx, network, address)
			},
		},
		TryAuthoritativeDNS:   true,
		SkipDNSDomainSuffixes: []string{".bsky.social"},
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	identity, err := baseDirectory.LookupDID(timeoutCtx, did)
	if err == nil {
		return identity.PDSEndpoint()
	}

	return BskyEndpoint
}

func NewClient(_ context.Context, filter []string, username string, password string, timestamp int64) (*Client, error) {
	client := &Client{
		username:       username,
		password:       password,
		filter:         filter,
		timestampStart: timestamp,
		cacheClient:    make(map[string]*XrpcClient),
		encoder:        form.NewEncoder(),
		httpClient:     http.DefaultClient,
	}

	defaultXrpcClient, err := client.createAndAuthenticateClient(context.Background(), BskyEndpoint)
	if err != nil {
		return nil, fmt.Errorf("create and authenticate client: %w", err)
	}

	client.defaultClient = defaultXrpcClient

	return client, nil
}
