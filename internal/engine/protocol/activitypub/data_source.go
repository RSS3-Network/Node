package activitypub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"sync"
	"time"

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
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

var _ engine.DataSource = (*dataSource)(nil)
var defaultStartTime int64

// serverPort is the default port for the ActivityPub server.
const (
	defaultServerPort  = ":8181"
	maxRequestBodySize = 1 << 20 // 1 MB

	// Version information
	nodeInfoSchemaURL          = "http://nodeinfo.diaspora.software/ns/schema/2.0"
	softwareName               = "custom-activitypub-relay"
	softwareVersion            = "1.0.0"
	protocolName               = "activitypub"
	httpServerTimeOut          = 20
	httpServerReadWriteTimeOut = 30
)

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
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("failed to initialize dataSource: %w", err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(3)

	// Start HTTP server
	go func() {
		defer wg.Done()

		if err := s.startHTTPServer(ctx); err != nil {
			errorChan <- fmt.Errorf("failed to start HTTP server: %w", err)
		}
	}()

	// Start relay service following
	go func() {
		defer wg.Done()

		if err := s.mastodonClient.FollowRelayServices(ctx); err != nil {
			errorChan <- fmt.Errorf("failed to follow relay services: %w", err)
		}
	}()

	// Start message processing
	go func() {
		defer wg.Done()

		if err := s.processMessages(ctx, tasksChan); err != nil {
			errorChan <- fmt.Errorf("failed to process messages: %w", err)
		}
	}()

	wg.Wait()
}

// startHTTPServer initializes and runs the HTTP server for ActivityPub endpoints.
func (s *dataSource) startHTTPServer(ctx context.Context) error {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Configure timeouts and middleware
	server.Server.ReadTimeout = httpServerReadWriteTimeOut * time.Second
	server.Server.WriteTimeout = httpServerReadWriteTimeOut * time.Second
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Setup routes
	server.GET("/actor", s.handleGetActor)
	server.POST("/actor/inbox", s.handleActorInbox)
	server.POST("/inbox", s.handleActorInbox)
	server.GET("/.well-known/nodeinfo", s.handleNodeInfo)
	server.GET("/nodeinfo/2.0", s.handleNodeInfoDetails)
	server.GET("/api/v1/instance", s.handleInstanceInfo)

	// Start server with graceful shutdown
	errCh := make(chan error, 1)

	go func() {
		if err := server.Start(defaultServerPort); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				zap.L().Error("server error", zap.Error(err))
				errCh <- err
			}
		}
	}()

	// Wait for context cancellation or server error
	select {
	case err := <-errCh:
		return fmt.Errorf("server startup failed: %w", err)

	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), httpServerTimeOut*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown failed: %w", err)
		}

		return nil
	}
}

// handleGetActor retrieves and returns the ActivityPub actor information.
func (s *dataSource) handleGetActor(c echo.Context) error {
	actor, err := s.mastodonClient.GetActor()
	if err != nil {
		zap.L().Error("failed to retrieve actor", zap.Error(err))

		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve actor")
	}

	return c.JSON(http.StatusOK, actor)
}

// handleActorInbox processes incoming ActivityPub messages sent to the actor's inbox.
func (s *dataSource) handleActorInbox(c echo.Context) error {
	requestBody := http.MaxBytesReader(c.Response(), c.Request().Body, maxRequestBodySize)

	// Close request body
	defer func(requestBody io.ReadCloser) {
		err := requestBody.Close()
		if err != nil {
			zap.L().Error("failed to close request body", zap.Error(err))
		}
	}(requestBody)

	data, err := io.ReadAll(requestBody)
	if err != nil {
		zap.L().Error("failed to read request body", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "failed to read request body")
	}

	message := string(data)
	zap.L().Info("received relay object", zap.String("message", message))
	s.mastodonClient.SendMessage(message)

	return c.NoContent(http.StatusAccepted)
}

// handleNodeInfo returns basic NodeInfo for service discovery.
func (s *dataSource) handleNodeInfo(c echo.Context) error {
	domain := s.config.Endpoint.URL

	info := activitypub.NodeInfo{
		Links: []activitypub.NodeInfoLink{
			{
				Rel:  nodeInfoSchemaURL,
				Href: fmt.Sprintf("%s/nodeinfo/2.0", domain),
			},
		},
	}

	return c.JSON(http.StatusOK, info)
}

// handleNodeInfoDetails returns detailed node information.
func (s *dataSource) handleNodeInfoDetails(c echo.Context) error {
	details := activitypub.NodeInfoDetails{
		Version: "2.0",
		Software: activitypub.SoftwareInfo{
			Name:    softwareName,
			Version: softwareVersion,
		},
		Protocols: []string{protocolName},
		Services: activitypub.ServicesInfo{
			Inbound:  []string{protocolName},
			Outbound: []string{protocolName},
		},
		OpenRegistrations: false,
		Usage: activitypub.NodeUsageInfo{
			Users: activitypub.UsersInfo{
				Total: 1,
			},
		},
	}

	return c.JSON(http.StatusOK, details)
}

// handleInstanceInfo provides instance-specific information.
func (s *dataSource) handleInstanceInfo(c echo.Context) error {
	info := activitypub.InstanceInfo{
		URI:         s.config.Endpoint.URL,
		Title:       "Custom Relay Instance",
		Description: "",
	}

	return c.JSON(http.StatusOK, info)
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
			if msg == "" {
				continue
			}
			// Process each message concurrently
			resultPool.Go(func() *engine.Tasks {
				return s.handleMessage(ctx, msg, tasksChan)
			})
		}
	}
}

// handleMessage processes a single ActivityPub message.
func (s *dataSource) handleMessage(ctx context.Context, msg string, tasksChan chan<- *engine.Tasks) *engine.Tasks {
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

		return s.processObjectTask(ctx, fetchedObject, tasksChan)
	}

	// If the message is not a relay message, attempt to unmarshal it directly into an ActivityPub object.
	var fetchedObject activitypub.Object
	if err := json.Unmarshal([]byte(msg), &fetchedObject); err != nil {
		zap.L().Error("Error unmarshalling message", zap.Error(err))
		return nil
	}

	// Remove activity suffix if present
	fetchedObject.ID = strings.TrimSuffix(fetchedObject.ID, mastodon.ActivitySuffix)

	return s.processObjectTask(ctx, &fetchedObject, tasksChan)
}

// processFetchedObject builds and sends tasks based on the fetched ActivityPub object.
func (s *dataSource) processObjectTask(ctx context.Context, obj *activitypub.Object, tasksChan chan<- *engine.Tasks) *engine.Tasks {
	zap.L().Info("processing received ActivityPub object",
		zap.String("objectID", obj.ID),
		zap.String("objectType", obj.Type),
		zap.Any("objectAttributes", obj),
	)

	tasks := s.buildMastodonMessageTasks(ctx, *obj)

	if tasks != nil {
		zap.L().Info("sending tasks", zap.Any("tasks", tasks))
		tasksChan <- tasks
	}

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
func (s *dataSource) initialize() (err error) {
	client, err := mastodon.NewClient(s.config.Endpoint.URL, s.option.RelayURLList)
	if err != nil {
		return fmt.Errorf("failed to create activitypub client: %w", err)
	}

	s.mastodonClient = client
	zap.L().Info("initialized mastodon client",
		zap.String("endpoint", s.config.Endpoint.URL))

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
