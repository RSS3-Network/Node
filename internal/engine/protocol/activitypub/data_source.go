package activitypub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

var _ engine.DataSource = (*dataSource)(nil)
var DefaultStartTime int64

// serverPort is the default port for the ActivityPub server.
const DefaultServerPort = ":8181"
const MaxRequestBodySize = int64(1048576) // 1 MB in bytes

// dataSource struct defines the fields for the data protocol
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

// Start initializes the data protocol and starts consuming Kafka messages
func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("failed to initialize dataSource: %w", err)
		return
	}

	// Start HTTP server
	go func() {
		if err := s.startHTTPServer(ctx); err != nil {
			errorChan <- fmt.Errorf("failed to start HTTP server: %w", err)
		}
	}()

	// Start relay service following
	go func() {
		if err := s.mastodonClient.FollowRelayServices(ctx); err != nil {
			errorChan <- fmt.Errorf("failed to follow relay services: %w", err)
		}
	}()

	// Start message processing
	go func() {
		if err := s.processMessages(ctx, tasksChan); err != nil {
			errorChan <- fmt.Errorf("failed to process messages: %w", err)
		}
	}()
}

// startHTTPServer initializes and runs the HTTP server for ActivityPub endpoints.
func (s *dataSource) startHTTPServer(ctx context.Context) error {
	httpServer := echo.New()
	httpServer.HideBanner = true
	httpServer.HidePort = true

	// Configure middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())

	// Setup routes
	httpServer.GET("/actor", s.handleGetActor)
	httpServer.POST("/actor/inbox", s.handleActorInbox)
	httpServer.POST("/inbox", s.handleActorInbox)

	// Start server with graceful shutdown
	go func() { //ToDo: When the HTTP server fails to start, is it expected that the node will not stop running?
		if err := httpServer.Start(DefaultServerPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Error("server error", zap.Error(err))
		}
	}()

	<-ctx.Done()

	return httpServer.Shutdown(ctx)
}

// HTTP Handlers
func (s *dataSource) handleGetActor(c echo.Context) error {
	// Modify fields dynamically
	actor, err := s.mastodonClient.GetActor()
	if err != nil {
		zap.L().Error("failed to get the actor: %v", zap.Error(err))
	}

	// Return the JSON to the client
	return c.JSON(http.StatusOK, actor)
}

// handleActorInbox processes incoming ActivityPub messages.
func (s *dataSource) handleActorInbox(c echo.Context) error {
	requestBody := http.MaxBytesReader(c.Response(), c.Request().Body, MaxRequestBodySize)

	var message activitypub.RelayObject
	if err := json.NewDecoder(requestBody).Decode(&message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid message format")
	}

	zap.L().Info("received relay object", zap.Any("message", message))
	s.mastodonClient.SendMessage(&message)

	return c.NoContent(http.StatusAccepted)
}

// processMessages handles incoming ActivityPub messages and generates tasks.
func (s *dataSource) processMessages(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	messageChan, err := s.mastodonClient.GetMessageChan()
	if err != nil {
		return fmt.Errorf("failed to get client message channel: %w", err)
	}

	// Create a pool to process messages concurrently
	resultPool := pool.NewWithResults[*engine.Tasks]().WithMaxGoroutines(runtime.NumCPU())

	for {
		select {
		case <-ctx.Done():
			resultPool.Wait()
			return ctx.Err()
		case msg := <-messageChan:
			if msg == nil {
				continue
			}
			// Process each message concurrently
			resultPool.Go(func() *engine.Tasks {
				zap.L().Info("processing message", zap.Any("message", msg))

				if msg.Type != "Announce" {
					zap.L().Info("message ignored, not an Announce type", zap.String("type", msg.Type))
					return nil
				}

				// Fetch and process the announced object
				var fetchedObject *activitypub.Object

				if err := retryOperation(ctx, func(ctx context.Context) error {
					var err error
					fetchedObject, err = s.mastodonClient.FetchAnnouncedObject(ctx, msg.Object)
					return err
				}); err != nil {
					zap.L().Error("failed to fetch",
						zap.String("object_url", msg.Object),
						zap.Error(err))

					return nil
				}

				tasks := s.buildMastodonMessageTasks(ctx, *fetchedObject)
				if tasks != nil {
					zap.L().Info("sending tasks", zap.Any("tasks", tasks))
					tasksChan <- tasks
				}

				return tasks
			})
		}
	}
}

// initialize creates and configures the Mastodon client.
// It returns an error if the client creation fails.
func (s *dataSource) initialize() (err error) {
	client, err := mastodon.NewClient(s.config.Endpoint.URL) //ToDo: Resolve LastOffset feature
	if err != nil {
		return fmt.Errorf("failed to create activitypub client: %w", err)
	}

	s.mastodonClient = client
	zap.L().Info("initialized mastodon client",
		zap.String("endpoint", s.config.Endpoint.URL))

	return nil
}

// retryOperation retries an operation with specified delay and backoff
func retryOperation(ctx context.Context, operation func(ctx context.Context) error) error {
	return retry.Do(
		func() error {
			return operation(ctx)
		},
		retry.Attempts(0),
		retry.Delay(1*time.Second),
		retry.DelayType(retry.BackOffDelay),
		retry.OnRetry(func(n uint, err error) {
			zap.L().Warn("retry activityPub dataSource", zap.Uint("retry", n), zap.Error(err))
		}),
	)
}

// buildMastodonMessageTasks processes the Kafka message and creates tasks for the engine
func (s *dataSource) buildMastodonMessageTasks(_ context.Context, object activitypub.Object) *engine.Tasks {
	var tasks engine.Tasks

	// If the object is empty, return an empty task
	if object.Type == "" {
		zap.L().Info("skipping empty object")
		return nil
	}

	tasks.Tasks = append(tasks.Tasks, &Task{
		Network: s.Network(),
		Message: object,
	})

	return &tasks
}

// NewSource creates a new data protocol instance
func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
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

	DefaultStartTime = instance.option.TimestampStart

	zap.L().Info("initialized data source", zap.Any("option", option))

	return instance, nil
}

func GetTimestampStart() int64 {
	return DefaultStartTime
}
