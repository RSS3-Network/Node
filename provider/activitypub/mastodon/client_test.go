package mastodon

//
//import (
//	"context"
//	"fmt"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/require"
//	"go.uber.org/zap"
//)
//
//// testConfig holds test configuration parameters
//var testConfig = struct {
//	domain    string
//	relayURLs []string
//}{
//	// Set your ngrok or tunnel domain here
//	domain: "https://newdomain7.ngrok.app",
//	// Add your relay URLs here
//	relayURLs: []string{
//		"https://relay.fedi.buzz/instance/mastodon.social",
//	},
//}
//
//// TestClientMonitorRelayMessages tests the relay message monitoring functionality
//func TestClientMonitorRelayMessages(t *testing.T) {
//	t.Parallel()
//
//	// Initialize logger for tests
//	logger, err := zap.NewDevelopment()
//	require.NoError(t, err)
//	zap.ReplaceGlobals(logger)
//
//	// Initialize client with error channel
//	errChan := make(chan error, 1)
//	ctx := context.Background()
//	client, err := NewClient(ctx, testConfig.domain, testConfig.relayURLs, errChan)
//	require.NoError(t, err)
//
//	// Wait for server ready signal
//	select {
//	case <-client.GetReadyChan():
//		t.Log("Server is ready")
//	case err := <-errChan:
//		t.Fatalf("Server initialization error: %v", err)
//	case <-time.After(5 * time.Second):
//		t.Fatal("Timeout waiting for server to be ready")
//	}
//
//	tests := []struct {
//		name        string
//		duration    time.Duration
//		minMessages int
//	}{
//		{
//			name:        "receive multiple messages",
//			duration:    50 * time.Second,
//			minMessages: 3,
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//			// Monitor messages
//			messages, err := monitorMessages(t, client, tt.duration, tt.minMessages)
//			require.NoError(t, err)
//			require.GreaterOrEqual(t, len(messages), tt.minMessages,
//				"insufficient messages received: got %d, want %d",
//				len(messages), tt.minMessages)
//
//			// Log received messages
//			for i, msg := range messages {
//				t.Logf("message %d: %s", i+1, msg)
//			}
//		})
//	}
//}
//
//// TestClientFetchAnnouncedObject tests fetching announced objects
//func TestClientFetchAnnouncedObject(t *testing.T) {
//	t.Parallel()
//
//	// Initialize logger for tests
//	logger, err := zap.NewDevelopment()
//	require.NoError(t, err)
//	zap.ReplaceGlobals(logger)
//
//	tests := []struct {
//		name      string
//		objectURL string
//		wantErr   bool
//	}{
//		{
//			name:      "fetch valid status",
//			objectURL: "https://mastodon.social/users/Gargron/statuses/103270115826048975",
//			wantErr:   false,
//		},
//		{
//			name:      "fetch invalid status",
//			objectURL: "https://invalid.url/status/123",
//			wantErr:   true,
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//		t.Run(tt.name, func(t *testing.T) {
//			t.Parallel()
//			// Initialize client with error channel
//			errChan := make(chan error, 1)
//			ctx := context.Background()
//			client, err := NewClient(ctx, testConfig.domain, testConfig.relayURLs, errChan)
//			require.NoError(t, err)
//
//			// Wait for server ready signal
//			select {
//			case <-client.GetReadyChan():
//				t.Log("Server is ready")
//			case err := <-errChan:
//				t.Fatalf("Server initialization error: %v", err)
//			case <-time.After(5 * time.Second):
//				t.Fatal("Timeout waiting for server to be ready")
//			}
//
//			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//			defer cancel()
//
//			obj, err := client.FetchAnnouncedObject(ctx, tt.objectURL)
//			if tt.wantErr {
//				require.Error(t, err)
//				return
//			}
//
//			require.NoError(t, err)
//			require.NotNil(t, obj)
//			require.NotEmpty(t, obj.ID)
//			t.Logf("fetched object: id=%s, type=%s", obj.ID, obj.Type)
//		})
//	}
//}
//
//// monitorMessages monitors incoming messages for the specified duration
//func monitorMessages(t *testing.T, client Client, duration time.Duration, minCount int) ([]string, error) {
//	msgChan, err := client.GetMessageChan()
//	if err != nil {
//		return nil, err
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), duration)
//	defer cancel()
//
//	messages := make([]string, 0, minCount)
//	done := make(chan bool, 1)
//
//	go func() {
//		for len(messages) < minCount {
//			select {
//			case msg := <-msgChan:
//				if msg == "" {
//					continue
//				}
//
//				messages = append(messages, msg)
//				t.Logf("received message #%d: %s", len(messages), msg)
//			case <-ctx.Done():
//				done <- false
//				return
//			}
//		}
//		done <- true
//	}()
//
//	if success := <-done; !success {
//		return messages, fmt.Errorf("timeout waiting for messages: got %d, want %d",
//			len(messages), minCount)
//	}
//
//	return messages, nil
//}
