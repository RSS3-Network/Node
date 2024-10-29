package activitypub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"strings"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

var _ engine.DataSource = (*dataSource)(nil)
var defaultStartTime int64

// dataSource implements the engine.DataSource interface.
type dataSource struct {
	config         *config.Module
	databaseClient database.Client
	mastodonClient mastodon.Client
	option         *Option
	state          State
}

// Network returns the network configuration from the config
func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

// Start initializes the data protocol and starts consuming relay messages
func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	if err := s.initialize(ctx, errorChan); err != nil {
		zap.L().Error("Failed to initialize", zap.Error(err))
		errorChan <- fmt.Errorf("failed to initialize dataSource: %w", err)

		return
	}

	if s.mastodonClient == nil {
		zap.L().Error("Mastodon client is nil after initialization")
		return
	}

	// Get the channel to get the notification of the http server being ready
	readyChan := s.mastodonClient.GetReadyChan()

	go func() {
		select {
		case <-ctx.Done():
			zap.L().Info("context cancelled before server was ready",
				zap.Error(ctx.Err()))
			return
		case <-readyChan:
			zap.L().Info("server is ready, starting to process messages")

			if err := s.processMessages(ctx, tasksChan); err != nil {
				zap.L().Error("failed to process messages", zap.Error(err))
				errorChan <- fmt.Errorf("failed to process messages: %w", err)
			}
		}
	}()
}

// processMessages handles incoming ActivityPub messages and generates tasks.
func (s *dataSource) processMessages(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	messageChan, err := s.mastodonClient.GetMessageChan()
	if err != nil {
		return fmt.Errorf("failed to get client message channel: %w", err)
	}

	zap.L().Info("starting message processing",
		zap.Int("channel_length", len(messageChan)))

	// Create a pool to process messages concurrently
	resultPool := pool.NewWithResults[*engine.Tasks]().WithMaxGoroutines(runtime.NumCPU())

	for {
		select {
		case <-ctx.Done():
			zap.L().Info("context canceled, stopping message processing")
			resultPool.Wait()

			return ctx.Err()
		case msg := <-messageChan:
			if msg == "" {
				zap.L().Info("empty message received, skipping")
				continue
			}

			zap.L().Info("received message for processing",
				zap.String("message", msg))

			// Process each message concurrently
			resultPool.Go(func() *engine.Tasks {
				tasks := s.handleMessage(ctx, msg)

				if tasks != nil {
					tasksChan <- tasks

					zap.L().Info("sent message to tasks channel", zap.String("message", msg))
				}

				return tasks
			})
		}
	}
}

// handleMessage processes a single ActivityPub message.
func (s *dataSource) handleMessage(ctx context.Context, msg string) *engine.Tasks {
	zap.L().Info("processing message", zap.Any("message", msg))

	relayObjectID := gjson.Get(msg, "id").String()
	// Check if the received message is a relay message or an ActivityPub message.
	if isRelayMessage(relayObjectID) {
		if gjson.Get(msg, "type").String() == "Accept" {
			zap.L().Info("Relay Subscription request has been approved.", zap.String("relayObjectID", relayObjectID))
			return nil
		}

		objectID := gjson.Get(msg, "object").String()
		zap.L().Info("Found object ID", zap.String("objectID", objectID))

		// Attempt to fetch the detailed ActivityPub object within the relay message.
		fetchedObject, err := s.mastodonClient.FetchAnnouncedObject(ctx, objectID)
		if err != nil {
			zap.L().Error("Failed to fetch announced object",
				zap.String("objectID", objectID), zap.Error(err))
			return nil
		}

		return s.processObjectTask(ctx, fetchedObject)
	}

	// If the message is not a relay message, attempt to unmarshal it directly into an ActivityPub object.
	var fetchedObject activitypub.Object
	if err := json.Unmarshal([]byte(msg), &fetchedObject); err != nil {
		zap.L().Error("Error unmarshalling message", zap.Error(err))
		return nil
	}

	// Remove activity suffix if present
	fetchedObject.ID = strings.TrimSuffix(fetchedObject.ID, mastodon.ActivitySuffix)

	return s.processObjectTask(ctx, &fetchedObject)
}

// processFetchedObject builds and sends tasks based on the fetched ActivityPub object.
func (s *dataSource) processObjectTask(ctx context.Context, obj *activitypub.Object) *engine.Tasks {
	zap.L().Info("processing received ActivityPub object",
		zap.String("objectID", obj.ID),
		zap.String("objectType", obj.Type),
		zap.Any("objectAttributes", obj),
	)

	tasks := s.buildMastodonMessageTasks(ctx, *obj)

	return tasks
}

// isRelayMessage determines if a URL is a relay message based on its host.
func isRelayMessage(id string) bool {
	parsedURL, err := url.Parse(id)
	if err != nil {
		zap.L().Error("parse URL failed",

			zap.String("id", id),
			zap.Error(err))

		return false
	}

	if parsedURL.Host == "" {
		zap.L().Error("Parsed URL has an empty host", zap.String("id", id))
		return false
	}

	if strings.HasPrefix(parsedURL.Host, "relay.") || strings.HasPrefix(parsedURL.Host, "rel.") {
		zap.L().Info("URL identified as a relay message", zap.String("id", id))

		return true
	}

	zap.L().Info("URL is not a relay message", zap.String("id", id))

	return false
}

// initialize creates and configures the Mastodon client.
// It returns an error if the client creation fails.
func (s *dataSource) initialize(ctx context.Context, errorChan chan<- error) (err error) {
	client, err := mastodon.NewClient(ctx, s.config.Endpoint.URL, s.option.RelayURLList, errorChan)
	if err != nil {
		return fmt.Errorf("failed to create activitypub client: %w", err)
	}

	s.mastodonClient = client
	zap.L().Info("initialized mastodon client",
		zap.String("endpoint", s.config.Endpoint.URL),
		zap.String("client_addr", fmt.Sprintf("%p", client)))

	return nil
}

// buildMastodonMessageTasks processes the relay message and creates tasks for the engine
func (s *dataSource) buildMastodonMessageTasks(_ context.Context, object activitypub.Object) *engine.Tasks {
	var tasks engine.Tasks

	// If the object is empty, return an empty task
	if object.Type == "" {
		zap.L().Info("skipping empty object")
		return nil
	}

	task := &Task{
		Network: s.Network(),
		Message: object,
	}

	tasks.Tasks = append(tasks.Tasks, task)

	zap.L().Info("task created for ActivityPub object",
		zap.String("objectID", object.ID),
		zap.Int("totalTasks", len(tasks.Tasks)),
		zap.Any("taskDetails", task),
	)

	return &tasks
}

// NewSource creates a new data protocol instance
func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
	if config == nil {
		return nil, errors.New("config is required")
	}

	var state State
	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, fmt.Errorf("failed to unmarshal checkpoint state: %w", err)
		}
	}

	option, err := NewOption(config.Network, config.Parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to create option: %w", err)
	}

	// Create a new datasource instance
	instance := &dataSource{
		databaseClient: databaseClient,
		config:         config,
		state:          state,
		option:         option,
	}

	defaultStartTime = instance.option.TimestampStart

	zap.L().Info("initialized data source", zap.Any("option", option))

	return instance, nil
}

func GetTimestampStart() int64 {
	return defaultStartTime
}
