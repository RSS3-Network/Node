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
	ProcessIncomingMessages(ctx context.Context) (<-chan *map[string]interface{}, error)
	GetActor() (*activitypub.Actor, error)
	SendMessage(*map[string]interface{})
	FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error)
	// GetKafkaConsumer() *kgo.Client
}

type client struct {
	httpClient *http.Client
	encoder    *form.Encoder
	attempts   uint
	domain     string
	privateKey *rsa.PrivateKey
	signer     httpsig.Signer
	msgChan    chan *map[string]interface{}
	actor      *activitypub.Actor
	relayURLs  []string
}

func NewClient(endpoint string) (Client, error) {
	privateKey, publicKeyPem, err := generateKeyPair()
	if err != nil {
		return nil, err
	}

	// Create actor template
	actor := &activitypub.Actor{
		Context: []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		Type:   "Application",
		ID:     fmt.Sprintf("%s/actor", endpoint),
		Inbox:  fmt.Sprintf("%s/actor/inbox", endpoint),
		Outbox: fmt.Sprintf("%s/actor/outbox", endpoint),
		PublicKey: struct {
			ID           string `json:"id"`
			Owner        string `json:"owner"`
			PublicKeyPem string `json:"publicKeyPem"`
		}{
			ID:           fmt.Sprintf("%s/actor#main-key", endpoint),
			Owner:        fmt.Sprintf("%s/actor", endpoint),
			PublicKeyPem: publicKeyPem,
		},
		Endpoints: struct {
			SharedInbox string `json:"sharedInbox"`
		}{
			SharedInbox: fmt.Sprintf("%s/inbox", endpoint),
		},
	}

	// Create signer
	signer, _, err := httpsig.NewSigner(
		[]httpsig.Algorithm{httpsig.RSA_SHA256},
		httpsig.DigestSha256,
		[]string{httpsig.RequestTarget, "Host", "Date", "Digest", "Content-Type"},
		httpsig.Signature,
		60*60,
	)
	if err != nil {
		zap.L().Error("Error creating NewSigner: %v", zap.Error(err))
	}

	// Set the Relay URL list
	relayURLs := []string{
		"https://relay.fedi.buzz/instance/mas.to",
		"https://relay.fedi.buzz/instance/mastodon.social",
	}

	return &client{
		domain:     endpoint,
		privateKey: privateKey,
		signer:     signer,
		msgChan:    make(chan *map[string]interface{}, MessageChannelMaxLength),
		actor:      actor,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		encoder:   form.NewEncoder(),
		attempts:  DefaultAttempts,
		relayURLs: relayURLs,
	}, nil
}
func (c *client) FetchAnnouncedObject(ctx context.Context, objectURL string) (*activitypub.Object, error) {
	httpClient, err := httpx.NewHTTPClient()
	if err != nil {
		zap.L().Error("failed to create HTTP client", zap.Error(err))
		return nil, fmt.Errorf("create HTTP client: %w", err)
	}

	// Append /activity to the object URL
	objectURLWithActivity := fmt.Sprintf("%s/activity", objectURL)
	zap.L().Info("fetching announced object", zap.String("object_url_with_activity", objectURLWithActivity))

	// Use the embedded request function to fetch the object data
	req, err := httpClient.Fetch(ctx, objectURLWithActivity)
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

	// Modify the 'id' to remove '/activity' at the end if it exists
	if fetchedObject.ID != "" && fetchedObject.ID[len(fetchedObject.ID)-9:] == "/activity" {
		zap.L().Info("Trimming /activity from ID", zap.String("original_id", fetchedObject.ID))
		trimmedID := fetchedObject.ID[:len(fetchedObject.ID)-9]
		fetchedObject.ID = trimmedID
	}

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
	// ToDo: configure with relay URLs rather than domain names
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
	httpClient, err := httpx.NewHTTPClient()
	if err != nil {
		return fmt.Errorf("create http client: %w", err)
	}

	uuidStr := uuid.New().String()
	contentStr := fmt.Sprintf(
		`{"@context":"https://www.w3.org/ns/activitystreams","id":"%s/%s","type":"Follow","actor":"%s/actor","object":"https://www.w3.org/ns/activitystreams#Public"}`,
		c.domain, uuidStr, c.domain,
	)
	content := []byte(contentStr)

	// Parse the instance URL to extract the host dynamically
	instanceURL, err := url.Parse(instance)
	if err != nil {
		return fmt.Errorf("parse instance URL: %w", err)
	}

	reqURL := instanceURL.String()
	zap.L().Info("following relay", zap.String("url", reqURL))
	zap.L().Info("instanceURL.HOST", zap.String("host", instanceURL.Host))

	// Make a POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(content))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
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
		return fmt.Errorf("sign request: %w", err)
	}

	// Use Fetch() to handle the request and retry logic
	resp, err := httpClient.Fetch(ctx, reqURL)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Close()

	// Read the response body
	body, err := io.ReadAll(resp)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check if the body contains an error or unexpected content
	if len(body) == 0 {
		return fmt.Errorf("empty response body from relay server: %s", instance)
	}

	// You can now process the body content, log it, or handle errors based on the content
	zap.L().Info("Successfully followed relay", zap.String("instance", instance), zap.String("body", string(body)))

	return nil
}

func (c *client) GetActor() (*activitypub.Actor, error) {
	return c.actor, nil
}

func (c *client) ProcessIncomingMessages(_ context.Context) (<-chan *map[string]interface{}, error) {
	return c.msgChan, nil
}

// Add the implementation to the client struct
func (c *client) SendMessage(msg *map[string]interface{}) {
	msgType, typeOk := (*msg)["type"].(string)
	msgID, idOk := (*msg)["id"].(string)

	select {
	case c.msgChan <- msg:
		// Log the message details if they exist
		if typeOk && idOk {
			zap.L().Debug("message sent to channel",
				zap.String("type", msgType),
				zap.String("id", msgID))
		} else {
			zap.L().Warn("message sent but type or id missing",
				zap.Any("message", msg))
		}
	default:
		// Log if the channel is full
		if typeOk && idOk {
			zap.L().Warn("message channel full, dropping message",
				zap.String("type", msgType),
				zap.String("id", msgID))
		} else {
			zap.L().Warn("message channel full, dropping message with missing type or id",
				zap.Any("message", msg))
		}
	}
}
