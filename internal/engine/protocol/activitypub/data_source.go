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

// Global constant for the server port
const ServerPort = ":8181" // ToDo: hardcode the server port here

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
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)
		return
	}

	// Start consuming Kafka messages in a goroutine
	go func() {
		if err := s.startHTTPServer(ctx); err != nil {
			errorChan <- fmt.Errorf("start HTTP server: %w", err)
			return
		}
	}()
	// Start following relay services
	go func() {
		if err := s.mastodonClient.FollowRelayServices(ctx); err != nil {
			errorChan <- fmt.Errorf("follow relay services: %w", err)
			return
		}
	}()

	go func() {
		if err := s.processMessages(ctx, tasksChan); err != nil {
			errorChan <- fmt.Errorf("process messages: %w", err)
			return
		}
	}()
}

func (s *dataSource) startHTTPServer(ctx context.Context) error {
	httpServer := echo.New()
	httpServer.HideBanner = true
	httpServer.HidePort = true
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())

	// Setup routes
	httpServer.GET("/actor", s.handleGetActor)
	httpServer.POST("/actor/inbox", s.handleActorInbox)
	httpServer.POST("/inbox", s.handleActorInbox)

	// Start server with graceful shutdown
	go func() { //ToDo: When the HTTP server fails to start, is it expected that the node will not stop running?
		if err := httpServer.Start(ServerPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
		zap.L().Error("Error getting the actor: %v", zap.Error(err))
	}

	// Return the JSON to the client
	return c.JSON(http.StatusOK, actor)
}

func (s *dataSource) handleActorInbox(c echo.Context) error {
	//ToDO: se the max size to avoid overflow
	maxBodySize := int64(1048576) // 1 MB in bytes
	requestBody := http.MaxBytesReader(c.Response(), c.Request().Body, maxBodySize)

	decoder := json.NewDecoder(requestBody)

	var message activitypub.RelayObject
	if err := decoder.Decode(&message); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid message format")
	}

	// Log the relay object
	zap.L().Info("handleActorInbox - received relay object", zap.Any("relayObject", message))

	s.mastodonClient.SendMessage(&message)

	return c.NoContent(http.StatusAccepted)
}

func (s *dataSource) processMessages(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	messageChan, err := s.mastodonClient.GetMessageChan()
	if err != nil {
		return fmt.Errorf("get message channel: %w", err)
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
			// Process each message concurrently using the result pool
			resultPool.Go(func() *engine.Tasks {
				zap.L().Info("processMessages - received message", zap.Any("message", msg))

				if msg.Type != "Announce" {
					return nil
				}

				// Fetch and process the announced object
				var fetchedObject *activitypub.Object

				if err := retryOperation(ctx, func(ctx context.Context) error {
					var err error
					fetchedObject, err = s.mastodonClient.FetchAnnouncedObject(ctx, msg.Object)
					return err
				}); err != nil {
					zap.L().Error("fetch failed",
						zap.String("object_url", msg.Object),
						zap.Error(err))

					return nil
				}

				tasks := s.buildMastodonMessageTasks(ctx, *fetchedObject)
				if tasks != nil {
					tasksChan <- tasks
				}

				return tasks
			})
		}
	}
}

// initialize sets up the Kafka consumer and Mastodon client
func (s *dataSource) initialize() (err error) {
	client, err := mastodon.NewClient(s.config.Endpoint.URL) //ToDo: Resolve LastOffset feature
	if err != nil {
		return fmt.Errorf("create activitypub client: %w", err)
	}

	s.mastodonClient = client

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
		return &tasks
	}

	tasks.Tasks = append(tasks.Tasks, &Task{
		Network: s.Network(),
		Message: object,
	})

	return &tasks
}

// NewSource creates a new data protocol instance
func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
	var (
		state State

		err error
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	// Create a new datasource instance
	instance := dataSource{
		databaseClient: databaseClient,
		config:         config,
		state:          state,
	}

	if instance.option, err = NewOption(config.Network, config.Parameters); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	DefaultStartTime = instance.option.TimestampStart

	zap.L().Info("apply option", zap.Any("option", instance.option))

	return &instance, nil
}

func GetTimestampStart() int64 {
	return DefaultStartTime
}
