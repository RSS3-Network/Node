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
	DefaultTimeout          = 2 * time.Second
	DefaultAttempts         = 2
	MessageChannelMaxLength = 1000
)

var _ Client = (*client)(nil)

type Client interface {
	FollowRelayServices(ctx context.Context) error
	GetMessageChan() (<-chan *activitypub.RelayObject, error)
	GetActor() (*activitypub.Actor, error)
	SendMessage(object *activitypub.RelayObject)
	FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error)
}

type client struct {
	httpClient httpx.Client //ToDo: check here
	encoder    *form.Encoder
	attempts   uint
	domain     string
	privateKey *rsa.PrivateKey
	signer     httpsig.Signer
	msgChan    chan *activitypub.RelayObject
	actor      *activitypub.Actor
	relayURLs  []string
}

// NewClient creates a new Mastodon client with the specified public endpoint.
// It initializes necessary cryptographic keys and ActivityPub actor information.
func NewClient(endpoint string) (Client, error) {
	privateKey, publicKeyPem, err := generateKeyPair()
	if err != nil {
		return nil, err
	}

	// Create actor
	actor := createActor(endpoint, publicKeyPem)

	// Create signer
	signer, err := createSigner()
	if err != nil {
		return nil, fmt.Errorf("failed to create signer: %w", err)
	}

	// RelayURLs is a list of URLs for the Mastodon relay service.
	// ToDo: Pass it as a parameter or read it from the config file.
	relayURLs := []string{
		"https://relay.fedi.buzz/instance/mas.to",
		"https://relay.fedi.buzz/instance/mastodon.social",
	}

	httpClient, err := httpx.NewHTTPClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	c := &client{
		domain:     endpoint,
		privateKey: privateKey,
		signer:     signer,
		msgChan:    make(chan *activitypub.RelayObject, MessageChannelMaxLength),
		actor:      actor,
		httpClient: httpClient,
		encoder:    form.NewEncoder(),
		attempts:   DefaultAttempts,
		relayURLs:  relayURLs,
	}

	return c, nil
}

// createActor creates a new ActivityPub actor with the public endpoint and public key.
func createActor(endpoint, publicKeyPem string) *activitypub.Actor {
	actorPath := endpoint + "/actor"

	return &activitypub.Actor{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		Type:   "Application",
		ID:     actorPath,
		Inbox:  actorPath + "/inbox",
		Outbox: actorPath + "/outbox",
		PublicKey: activitypub.PublicKey{
			ID:           actorPath + "#main-key",
			Owner:        actorPath,
			PublicKeyPem: publicKeyPem,
		},
		Endpoints: activitypub.EndPoint{
			SharedInbox: endpoint + "/inbox",
		},
	}
}

func createSigner() (httpsig.Signer, error) {
	signer, _, err := httpsig.NewSigner(
		[]httpsig.Algorithm{httpsig.RSA_SHA256},
		httpsig.DigestSha256,
		[]string{httpsig.RequestTarget, "Host", "Date", "Digest", "Content-Type"},
		httpsig.Signature,
		60*60,
	)
	if err != nil {
		return nil, err
	}

	return signer, nil
}

// FetchAnnouncedObject retrieves an ActivityPub object from the specified URL.
func (c *client) FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error) {
	objectURLWithActivity := objectURL + "/activity"

	zap.L().Info("Fetching announced object with activity URL",
		zap.String("activityURL", objectURLWithActivity),
	)

	req, err := c.httpClient.Fetch(ctx, objectURLWithActivity)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch object: %w", err)
	}

	zap.L().Info("Request successful")

	// Read the response body
	body, err := io.ReadAll(req)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	zap.L().Info("response body received", zap.ByteString("body", body))

	// Unmarshal the JSON response into an activitypub.Object
	var obj activitypub.Object
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, fmt.Errorf("failed to decode object: %w", err)
	}

	//ToDo: test if this works or not
	obj.ID = strings.TrimSuffix(obj.ID, "/activity")

	zap.L().Info("fetched ActivityPub object", zap.Any("fetchedObject", obj))

	return &obj, nil
}

// generateKeyPair creates a new RSA key pair for signing ActivityPub messages.
func generateKeyPair() (*rsa.PrivateKey, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate RSA key: %w", err)
	}
	// Marshal the public key to PEM format
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

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
	if c.domain == "" {
		return fmt.Errorf("domain not configured")
	}

	// Generate ActivityPub Message to subscribe relay services
	content := c.generateFollowContent()

	instanceURL, err := url.Parse(instance)
	if err != nil {
		return fmt.Errorf("failed to parse instance URL: %w", err)
	}

	// Create and sign the follow request
	req, err := c.createAndSignFollowRequest(ctx, instanceURL, content)
	if err != nil {
		return fmt.Errorf("failed to create and sign follow request: %w", err)
	}

	// Send the request and handle the response
	if err := c.sendRequest(req); err != nil {
		return fmt.Errorf("failed to send request for relay subscription: %w", err)
	}

	return nil
}

// generateFollowContent creates the JSON content for a follow request.
func (c *client) generateFollowContent() []byte {
	return []byte(fmt.Sprintf(
		`{"@context":"%s","id":"%s/%s","type":"Follow","actor":"%s/actor","object":"%s"}`,
		ActivityStreamsContext,
		c.domain, uuid.New().String(), c.domain,
		ActivityStreamsPublicContext,
	))
}

// createAndSignFollowRequest creates and signs an HTTP request for following a relay.
func (c *client) createAndSignFollowRequest(ctx context.Context, instanceURL *url.URL, content []byte) (*http.Request, error) {
	reqURL := instanceURL.String()

	// Create a POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Host", instanceURL.Host)
	req.Header.Set("Content-Type", "application/activity+json")
	req.Header.Set("Date", time.Now().Format(time.RFC1123))

	// Sign the request
	if err := c.signer.SignRequest(
		c.privateKey,
		fmt.Sprintf("%s/actor#main-key", c.domain),
		req,
		content,
	); err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	return req, nil
}

// sendRequest sends an HTTP request and validates the response.
func (c *client) sendRequest(req *http.Request) error {
	httpClient := http.Client{}

	resp, err := httpClient.Do(req) // ToDo: Use Fetch() to handle the request (now using it will cause hanging issue)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unexpected status %s: failed to read error response", resp.Status)
		}

		return fmt.Errorf("unexpected status %s: %s", resp.Status, body)
	}

	return nil
}

// GetActor returns the ActivityPub actor associated with this client.
func (c *client) GetActor() (*activitypub.Actor, error) {
	return c.actor, nil
}

// GetMessageChan returns the channel for receiving relay messages.
func (c *client) GetMessageChan() (<-chan *activitypub.RelayObject, error) {
	return c.msgChan, nil
}

// SendMessage sends a message to the relay message channel.
// If the channel is full, the message is dropped and logged.
func (c *client) SendMessage(msg *activitypub.RelayObject) {
	select {
	case c.msgChan <- msg:
		zap.L().Info("message queued",
			zap.String("type", msg.Type),
			zap.String("id", msg.ID))
	default:
		zap.L().Info("channel full, dropping message",
			zap.String("type", msg.Type),
			zap.String("id", msg.ID))
	}
}
