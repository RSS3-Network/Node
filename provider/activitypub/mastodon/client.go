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
	httpClient httpx.Client // ToDo: check if needed to pass by reference?
	encoder    *form.Encoder
	attempts   uint
	domain     string
	privateKey *rsa.PrivateKey
	signer     httpsig.Signer
	msgChan    chan *activitypub.RelayObject
	actor      *activitypub.Actor
	relayURLs  []string
}

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
		zap.L().Error("Error creating NewSigner: %v", zap.Error(err))
	}

	if err != nil {
		zap.L().Error("Error creating NewSigner: %v", zap.Error(err))
	}

	// Set the Relay URL list
	relayURLs := []string{ // ToDo: This should not be hardcoded...
		// "https://relay.fedi.buzz/tag/hello",
		"https://relay.fedi.buzz/instance/mas.to",
		// "https://relay.fedi.buzz/instance/mastodon.social",
	}

	httpClient, err := httpx.NewHTTPClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	return &client{
		domain:     endpoint,
		privateKey: privateKey,
		signer:     signer,
		msgChan:    make(chan *activitypub.RelayObject, MessageChannelMaxLength),
		actor:      actor,
		httpClient: httpClient,
		encoder:    form.NewEncoder(),
		attempts:   DefaultAttempts,
		relayURLs:  relayURLs,
	}, nil
}

func createActor(endpoint string, publicKeyPem string) *activitypub.Actor {
	actorPath := endpoint + "/actor"

	return &activitypub.Actor{
		Context: []string{
			ActivityStreamsContext,
			SecurityV1Context,
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

func (c *client) FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error) {
	// Append /activity to the object URL
	objectURLWithActivity := fmt.Sprintf("%s/activity", objectURL)
	zap.L().Info("fetching announced object", zap.String("object_url_with_activity", objectURLWithActivity))

	// Use the embedded request function to fetch the object data
	req, err := c.httpClient.Fetch(ctx, objectURLWithActivity)
	if err != nil {
		zap.L().Error("failed to fetch announced object", zap.String("url", objectURLWithActivity), zap.Error(err))
		return nil, fmt.Errorf("create request: %w", err)
	}

	zap.L().Info("request successful", zap.String("url", objectURLWithActivity))

	// Read the response body
	body, err := io.ReadAll(req)
	if err != nil {
		zap.L().Error("failed to read response body", zap.String("url", objectURLWithActivity), zap.Error(err))
		return nil, fmt.Errorf("read response body: %w", err)
	}

	zap.L().Debug("response body received", zap.ByteString("body", body))

	// Unmarshal the JSON response into an activitypub.Object
	var fetchedObject activitypub.Object
	if err := json.Unmarshal(body, &fetchedObject); err != nil {
		zap.L().Error("failed to decode response", zap.String("url", objectURLWithActivity), zap.Error(err))
		return nil, fmt.Errorf("decode response: %w", err)
	}

	//ToDo: test if this works or not
	fetchedObject.ID = strings.TrimSuffix(fetchedObject.ID, "/activity")

	zap.L().Info("fetched object successfully", zap.Any("fetchedObject", fetchedObject))

	return &fetchedObject, nil
}

// Function to generate RSA private and public key pair
func generateKeyPair() (*rsa.PrivateKey, string, error) {
	// Generate private key (2048-bit RSA key)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, "", err
	}

	// Marshal the public key to PEM format
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return privateKey, string(publicKeyPem), nil
}

func (c *client) FollowRelayServices(ctx context.Context) error {
	for _, instance := range c.relayURLs {
		if err := c.followRelay(ctx, instance); err != nil {
			zap.L().Error("failed to follow relay",
				zap.String("instance", instance),
				zap.Error(err))
		}
	}

	return nil
}
func (c *client) followRelay(ctx context.Context, instance string) error {
	if c.domain == "" {
		return fmt.Errorf("domain is empty")
	}

	content := c.generateFollowContent()

	// Parse the instance URL to extract the host dynamically
	instanceURL, err := url.Parse(instance)
	if err != nil {
		return fmt.Errorf("parse instance URL: %w", err)
	}

	// Create and sign the follow request
	req, err := c.createAndSignFollowRequest(ctx, instanceURL, content)
	if err != nil {
		return fmt.Errorf("create and sign follow request: %w", err)
	}

	// Send the request and handle the response
	if err := c.sendRequest(req); err != nil {
		return err
	}

	return nil
}

func (c *client) generateFollowContent() []byte {
	uuidStr := uuid.New().String()
	contentStr := fmt.Sprintf(
		`{"@context":"%s","id":"%s/%s","type":"Follow","actor":"%s/actor","object":"%s"}`,
		ActivityStreamsContext,
		c.domain, uuidStr, c.domain,
		ActivityStreamsPublicContext,
	)
	content := []byte(contentStr)

	return content
}

func (c *client) createAndSignFollowRequest(ctx context.Context, instanceURL *url.URL, content []byte) (*http.Request, error) {
	reqURL := instanceURL.String()
	zap.L().Info("following relay", zap.String("url", reqURL))
	zap.L().Info("instanceURL.Host", zap.String("host", instanceURL.Host))

	// Create a POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
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
		return nil, fmt.Errorf("sign request: %w", err)
	}

	return req, nil
}

func (c *client) sendRequest(req *http.Request) error {
	httpClient := http.Client{}

	resp, err := httpClient.Do(req) // ToDo: Use Fetch() to handle the request (now using it will cause hanging issue)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status: %s, body: %s", resp.Status, string(body))
	}

	return nil
}

func (c *client) GetActor() (*activitypub.Actor, error) {
	return c.actor, nil
}

func (c *client) GetMessageChan() (<-chan *activitypub.RelayObject, error) {
	return c.msgChan, nil
}

// Add the implementation to the client struct
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
