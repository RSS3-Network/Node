package mastodon

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-fed/httpsig"
	"github.com/go-playground/form/v4"
	"github.com/google/uuid"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/httpx"
	"go.uber.org/zap"
)

const (
	defaultAttempts         = 2
	messageChannelMaxLength = 1000
	keySize                 = 2048
	sigExpiry               = 60 * 60
)

var _ Client = (*client)(nil)

type Client interface {
	FollowRelayServices(ctx context.Context) error
	GetMessageChan() (<-chan string, error)
	GetActor() (*activitypub.Actor, error)
	SendMessage(object string)
	FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error)
}

type client struct {
	httpClient httpx.Client
	encoder    *form.Encoder
	attempts   uint
	domain     string
	privateKey *rsa.PrivateKey
	msgChan    chan string
	actor      *activitypub.Actor
	relayURLs  []string
}

// NewClient creates a new Mastodon client with the specified public endpoint.
// It initializes cryptographic keys and ActivityPub actor information.
func NewClient(endpoint string, relayURLList []string) (Client, error) {
	// Input Validation
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	if len(relayURLList) == 0 {
		return nil, fmt.Errorf("relayURLList cannot be empty")
	}

	// Generate the Key Pair
	privateKey, publicKeyPem, err := generateKeyPair()
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %w", err)
	}

	// Create actor
	actor, err := createActor(endpoint, publicKeyPem)
	if err != nil {
		return nil, fmt.Errorf("failed to create actor: %w", err)
	}

	httpClient, err := httpx.NewHTTPClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	c := &client{
		domain:     endpoint,
		privateKey: privateKey,
		msgChan:    make(chan string, messageChannelMaxLength),
		actor:      actor,
		httpClient: httpClient,
		encoder:    form.NewEncoder(),
		attempts:   defaultAttempts,
		relayURLs:  relayURLList,
	}

	return c, nil
}

// createActor creates a new ActivityPub actor with the public endpoint and public key.
func createActor(endpoint, publicKeyPem string) (*activitypub.Actor, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	if publicKeyPem == "" {
		return nil, fmt.Errorf("publicKeyPem cannot be empty")
	}

	actorPath := endpoint + actorPath

	return &activitypub.Actor{
		Context: []string{
			ActivityStreamsContext,
			SecurityV1Context,
		},
		Type:   "Application",
		ID:     actorPath,
		Inbox:  actorPath + inBoxPath,
		Outbox: actorPath + outBoxPath,
		PublicKey: activitypub.PublicKey{
			ID:           actorPath + "#main-key",
			Owner:        actorPath,
			PublicKeyPem: publicKeyPem,
		},
		Endpoints: activitypub.EndPoint{
			SharedInbox: endpoint + inBoxPath,
		},
	}, nil
}

// newSigner creates a new HTTP signature signer.
func newSigner(instanceURL *url.URL) (httpsig.Signer, error) {
	if instanceURL == nil {
		return nil, fmt.Errorf("instanceURL cannot be nil")
	}

	headers := []string{httpsig.RequestTarget, headerDate, headerContentType}
	if !strings.HasSuffix(instanceURL.Path, inBoxPath) {
		headers = append(headers, headerHost, headerDigest)
	}

	zap.L().Info("Creating signer", zap.String("instanceURL", instanceURL.String()), zap.Strings("headers", headers))

	signer, _, err := httpsig.NewSigner(
		[]httpsig.Algorithm{httpsig.RSA_SHA256},
		httpsig.DigestSha256,
		headers,
		httpsig.Signature,
		sigExpiry,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to make a NewSigner call: %w", err)
	}

	return signer, nil
}

// generateKeyPair creates a new RSA key pair for signing ActivityPub messages.
func generateKeyPair() (*rsa.PrivateKey, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate key using rsa: %w", err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get make marshal public key: %w", err)
	}

	// Encode in PEM format with PUBLIC KEY header
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	if publicKeyPem == nil {
		return nil, "", fmt.Errorf("failed to encode public key")
	}

	return privateKey, string(publicKeyPem), nil
}

// FollowRelayServices attempts to follow all configured relay services.
func (c *client) FollowRelayServices(ctx context.Context) error {
	var errs []error

	for _, instance := range c.relayURLs {
		if err := c.followRelay(ctx, instance); err != nil {
			zap.L().Error("failed to follow relay", zap.String("instance", instance), zap.Error(err))
			errs = append(errs, fmt.Errorf("failed to follow relay instance %s: %w", instance, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("multiple errors occurred while following relay services: %v", errs)
	}

	return nil
}

// followRelay attempts to follow a single relay instance.
func (c *client) followRelay(ctx context.Context, instance string) error {
	if instance == "" {
		return fmt.Errorf("instance cannot be empty")
	}

	// Generate ActivityPub Message to subscribe relay services
	content := c.generateFollowContent()

	instanceURL, err := url.Parse(instance)
	if err != nil {
		return fmt.Errorf("failed to parse instance URL: %w", err)
	}

	// Create signer
	signer, err := newSigner(instanceURL)
	if err != nil {
		return fmt.Errorf("failed to create a new signer: %w", err)
	}

	// Create and sign the follow request
	req, err := c.createAndSignFollowRequest(ctx, instanceURL, content, signer)
	if err != nil {
		return fmt.Errorf("failed to create and sign follow request: %w", err)
	}

	// Send the request and handle the response
	if err := c.sendRequest(req); err != nil {
		return fmt.Errorf("failed to send request for relay subscription: %w", err)
	}

	zap.L().Info("successfully followed relay", zap.String("instance", instance))

	return nil
}

// generateFollowContent creates the JSON content for a follow request.
func (c *client) generateFollowContent() []byte {
	followContent := fmt.Sprintf(
		`{"@context":"%s","id":"%s/%s","type":"Follow","actor":"%s/actor","object":"%s"}`,
		ActivityStreamsContext,
		c.domain, uuid.New().String(), c.domain,
		ActivityStreamsPublicContext,
	)

	return []byte(followContent)
}

// createAndSignFollowRequest creates and signs an HTTP request for following a relay.
func (c *client) createAndSignFollowRequest(ctx context.Context, instanceURL *url.URL, content []byte, signer httpsig.Signer) (*http.Request, error) {
	switch {
	case instanceURL == nil:
		return nil, fmt.Errorf("instance URL cannot be nil")
	case len(content) == 0:
		return nil, fmt.Errorf("content cannot be empty")
	case signer == nil:
		return nil, fmt.Errorf("signer cannot be nil")
	}

	zap.L().Info("Creating follow request", zap.String("instanceURL", instanceURL.String()), zap.ByteString("content", content))
	reqURL := instanceURL.String()

	// Create a POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to create context of the request: %w", err)
	}

	req.Header.Set(headerContentType, activityJSONType)
	req.Header.Set(headerDate, time.Now().Format(time.RFC1123))

	// Set Host header only if instance URL doesn't end in /inbox
	if !strings.HasSuffix(instanceURL.Path, inBoxPath) {
		req.Header.Set(headerHost, instanceURL.Host)
		zap.L().Info("Set Host header", zap.String(headerHost, instanceURL.Host))
	}

	// Sign the request
	keyID := fmt.Sprintf("%s/actor#main-key", c.domain)
	if err := signer.SignRequest(
		c.privateKey,
		keyID,
		req,
		content,
	); err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	return req, nil
}

// sendRequest sends an HTTP request and validates the response.
func (c *client) sendRequest(req *http.Request) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}

	// Close request body
	defer func() {
		if err := resp.Body.Close(); err != nil {
			zap.L().Error("closing response body", zap.Error(err))
		}
	}()

	// Log basic response details (status code and headers)
	zap.L().Info("received response",
		zap.Int("status_code", resp.StatusCode),
		zap.Any("headers", resp.Header),
		zap.Any("resp", resp),
	)

	if resp.StatusCode != http.StatusAccepted {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unexpected status %s: failed to read error response", resp.Status)
		}

		return fmt.Errorf("unexpected status %s: %s", resp.Status, body)
	}

	return nil
}

// FetchAnnouncedObject retrieves an ActivityPub object from the specified URL.
func (c *client) FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error) {
	if objectURL == "" {
		return nil, fmt.Errorf("objectURL cannot be empty")
	}

	// Append "/activity" in case if the objectURL doesn't already end with it
	activityURL := objectURL
	if !strings.HasSuffix(objectURL, ActivitySuffix) {
		activityURL += ActivitySuffix
	}

	zap.L().Info("Fetching announced object with activity URL",
		zap.String("activityURL", activityURL),
	)

	body, err := c.fetchObject(ctx, activityURL)
	if err != nil {
		return nil, fmt.Errorf("fetching object: %w", err)
	}

	// Unmarshal the JSON response into an ActivityPub object
	var obj activitypub.Object
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, fmt.Errorf("failed to decode object: %w", err)
	}

	// Handle ID manipulation by trimming /activity
	obj.ID = strings.TrimSuffix(obj.ID, ActivitySuffix)

	zap.L().Info("fetched ActivityPub object", zap.Any("fetchedObject", obj))

	return &obj, nil
}

// fetchObject fetches an object from a URL and return its response body
func (c *client) fetchObject(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.httpClient.Fetch(ctx, url)

	if err != nil {
		return nil, fmt.Errorf("fetch object: %w", err)
	}

	zap.L().Info("Request was successful")

	body, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	zap.L().Info("response body received", zap.ByteString("body", body))

	return body, nil
}

// GetActor returns the ActivityPub actor associated with this client.
func (c *client) GetActor() (*activitypub.Actor, error) {
	if c.actor == nil {
		return nil, fmt.Errorf("actor not initialized")
	}

	return c.actor, nil
}

// GetMessageChan returns the channel for receiving relay messages.
func (c *client) GetMessageChan() (<-chan string, error) {
	if c.msgChan == nil {
		return nil, fmt.Errorf("message channel not initialized")
	}

	return c.msgChan, nil
}

// SendMessage queues a message string to the relay message channel.
// If the channel is full, the message is dropped and logged.
func (c *client) SendMessage(msg string) {
	if msg == "" {
		zap.L().Warn("empty message dropped")
		return
	}

	select {
	case c.msgChan <- msg:
		zap.L().Info("message queued", zap.String("message", msg))
	default:
		zap.L().Info("channel full, dropping message", zap.String("message", msg))
	}
}
