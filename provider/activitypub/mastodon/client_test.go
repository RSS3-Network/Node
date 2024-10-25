package mastodon_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
	"github.com/stretchr/testify/require"
)

// testConfig holds test configuration parameters
var testConfig = struct {
	domain     string
	relayURLs  []string
	serverPort string
}{
	domain: "https://newdomain6.ngrok.app",
	relayURLs: []string{
		"https://relay.fedi.buzz/instance/mastodon.social",
	},
	serverPort: ":8181",
}

// TestClientMonitorRelayMessages tests the relay message monitoring functionality.
func TestClientMonitorRelayMessages(t *testing.T) {
	t.Parallel()

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	var (
		client mastodon.Client
		err    error
	)

	// Start test server
	go func() {
		if err := e.Start(testConfig.serverPort); err != nil {
			if err.Error() != "http: Server closed" {
				t.Logf("server error: %v", err)
			}
		}
	}()

	defer e.Close()

	// Wait for server initialization
	time.Sleep(time.Second)

	// Initialize client
	client, err = mastodon.NewClient(testConfig.domain, testConfig.relayURLs)
	require.NoError(t, err)

	// Configure server endpoints
	setupTestServer(e, client)

	tests := []struct {
		name        string
		duration    time.Duration
		minMessages int
	}{
		{
			name:        "receive multiple messages",
			duration:    50 * time.Second,
			minMessages: 3,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Allow server to stabilize
			time.Sleep(10 * time.Second)

			// Follow relay service
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			err := client.FollowRelayServices(ctx)

			cancel()
			require.NoError(t, err, "follow relay service")

			// Monitor messages
			messages, err := monitorMessages(t, client, tt.duration, tt.minMessages)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(messages), tt.minMessages,
				"insufficient messages received: got %d, want %d",
				len(messages), tt.minMessages)

			// Log received messages
			for i, msg := range messages {
				t.Logf("message %d: %s", i+1, msg)
			}
		})
	}
}

// TestClientFetchAnnouncedObject tests fetching announced objects.
func TestClientFetchAnnouncedObject(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		objectURL string
		wantErr   bool
	}{
		{
			name:      "fetch valid status",
			objectURL: "https://mas.to/users/yantor3d/statuses/113371504232175109",
			wantErr:   false,
		},
		{
			name:      "fetch invalid status",
			objectURL: "https://invalid.url/status/123",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := mastodon.NewClient(testConfig.domain, testConfig.relayURLs)
			require.NoError(t, err)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			obj, err := client.FetchAnnouncedObject(ctx, tt.objectURL)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, obj)
			require.NotEmpty(t, obj.ID)
			t.Logf("fetched object: id=%s, type=%s", obj.ID, obj.Type)
		})
	}
}

// setupTestServer configures the test server endpoints.
func setupTestServer(e *echo.Echo, client mastodon.Client) {
	e.GET("/actor", func(c echo.Context) error {
		actor, err := client.GetActor()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, actor)
	})

	// Handle both inbox endpoints
	inboxHandler := func(c echo.Context) error {
		body := c.Request().Body
		defer body.Close()

		data, err := io.ReadAll(body)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		client.SendMessage(string(data))

		return c.NoContent(http.StatusAccepted)
	}

	e.POST("/inbox", inboxHandler)
	e.POST("/actor/inbox", inboxHandler)
}

// monitorMessages monitors incoming messages for the specified duration.
func monitorMessages(t *testing.T, client mastodon.Client, duration time.Duration, minCount int) ([]string, error) {
	msgChan, err := client.GetMessageChan()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	messages := make([]string, 0, minCount)
	done := make(chan bool, 1)

	go func() {
		for len(messages) < minCount {
			select {
			case msg := <-msgChan:
				if msg == "" {
					continue
				}

				messages = append(messages, msg)
				t.Logf("received message #%d: %s", len(messages), msg)
			case <-ctx.Done():
				done <- false
				return
			}
		}
		done <- true
	}()

	if success := <-done; !success {
		return messages, fmt.Errorf("timeout waiting for messages: got %d, want %d",
			len(messages), minCount)
	}

	return messages, nil
}
