package activitypub

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/IBM/sarama"
	"github.com/avast/retry-go/v4"
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

const (
	kafkaTopic = "activitypub_events" // kafka topic name
)

var _ engine.DataSource = (*dataSource)(nil)

// dataSource struct defines the fields for the data source
type dataSource struct {
	config         *config.Module
	databaseClient database.Client
	mastodonClient mastodon.Client
	state          State
}

// Network returns the network configuration from the config
func (s *dataSource) Network() network.Network {
	return s.config.Network
}

func (s *dataSource) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

// Start initializes the data source and starts consuming Kafka messages
func (s *dataSource) Start(ctx context.Context, tasksChan chan<- *engine.Tasks, errorChan chan<- error) {
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize dataSource: %w", err)
		return
	}

	// Start consuming Kafka messages in a goroutine
	go func() {
		if err := retryOperation(ctx, func(ctx context.Context) error {
			return s.consumeKafkaMessages(ctx, tasksChan)
		}); err != nil {
			errorChan <- err
		}
	}()
}

// consumeKafkaMessages consumes messages from a Kafka topic and processes them
func (s *dataSource) consumeKafkaMessages(ctx context.Context, tasksChan chan<- *engine.Tasks) error {
	consumer := s.mastodonClient.GetKafkaConsumer()
	// Start consuming messages from the specified Kafka partition
	partitionConsumer, err := consumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		return fmt.Errorf("consume partition: %w", err)
	}

	defer partitionConsumer.Close()

	// Create a pool to process messages concurrently
	resultPool := pool.NewWithResults[*engine.Tasks]().WithMaxGoroutines(runtime.NumCPU())

	// Loop to continuously process incoming messages
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// Send the message to the result pool for concurrent processing
			resultPool.Go(func() *engine.Tasks {
				// process each message
				tasks := s.processMessage(ctx, msg)
				if tasks != nil {
					tasksChan <- tasks
				}

				return tasks
			})
		case <-ctx.Done():
			resultPool.Wait()
			return ctx.Err()
		}
	}
}

// processMessage processes a single Kafka message
func (s *dataSource) processMessage(ctx context.Context, msg *sarama.ConsumerMessage) *engine.Tasks {
	// Update the state with the last processed message count number (offset)
	s.state.LastOffset = msg.Offset

	// fmt.Printf("[activitypub/data_source.go] Consumed message offset: %d, value: %s\n", msg.Offset, string(msg.Value))

	// Unmarshal the message and store it as an ActivityPub object
	messageValue := string(msg.Value)

	var object activitypub.Object

	if err := json.Unmarshal([]byte(messageValue), &object); err != nil {
		zap.L().Error("failed to unmarshal message value", zap.Error(err))
		return nil
	}

	// fmt.Printf("[activitypub/data_source.go] Unmarshal object is %+v\n", object)

	// Build the corresponding message task for transformation
	tasks := s.buildMastodonMessageTasks(ctx, object)

	// Print the tasks for debugging
	fmt.Println("[activitypub/data_source.go] Generated tasks:")

	for _, task := range tasks.Tasks {
		fmt.Printf("Task: %+v\n", task)
	}

	return tasks
}

// initialize sets up the Kafka consumer and Mastodon client
func (s *dataSource) initialize() (err error) {
	// Create a new Client instance
	client, err := mastodon.NewClient(s.config.Endpoint.URL)
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
			zap.L().Warn("retry farcaster dataSource", zap.Uint("retry", n), zap.Error(err))
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

	// Filter the message based on type
	// switch object.Type {
	// case "Create":
	//	if note, ok := object.Object.(map[string]interface{}); ok {
	//		if noteType, exists := note["type"].(string); exists && noteType == "Note" {
	//			// Example: Create task for a Note object
	//			fmt.Println("[buildMastodonMessageTasks] Object in Create type (for a Note Object)")
	//		}
	//	}
	// case "Follow":
	//	// Process Follow type messages
	//	fmt.Println("[buildMastodonMessageTasks] Object in Follow type")
	// case "Like":
	//	// Process Like type messages
	//	fmt.Println("[buildMastodonMessageTasks] Object in Like type")
	// default:
	//	zap.L().Debug("unsupported message type", zap.String("type", object.Type))
	//}

	tasks.Tasks = append(tasks.Tasks, &Task{
		Network: s.Network(),
		Message: object,
	})

	return &tasks
}

// NewSource creates a new data source instance
func NewSource(config *config.Module, checkpoint *engine.Checkpoint, databaseClient database.Client) (engine.DataSource, error) {
	var (
		state State
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

	return &instance, nil
}
