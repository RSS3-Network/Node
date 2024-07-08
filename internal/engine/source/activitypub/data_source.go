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
	defaultBlockTime       = 3 * time.Second
	kafkaTopic             = "activitypub_events" // kafka topic name
	mastodonSourceEndpoint = "34.125.79.79"       // External IP of the Mastodon VM instance
)

var _ engine.DataSource = (*dataSource)(nil)

// dataSource struct defines the fields needed for the data source
type dataSource struct {
	config         *config.Module
	databaseClient database.Client
	mastodonClient mastodon.Client
	kafkaConsumer  sarama.Consumer
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
	// Start consuming messages from the specified Kafka partition
	partitionConsumer, err := s.kafkaConsumer.ConsumePartition(kafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		return fmt.Errorf("consume partition: %w", err)
	}

	defer partitionConsumer.Close()

	// Loop to continuously process incoming messages
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// Update the state with the last processed message offset
			s.state.LastOffset = msg.Offset

			fmt.Printf("Consumed message offset: %d, value: %s\n", msg.Offset, string(msg.Value))

			// unmarshal the message and store it as a ActivityPub object
			messageValue := string(msg.Value)

			var object activitypub.Object

			if err := json.Unmarshal([]byte(messageValue), &object); err != nil {
				zap.L().Error("failed to unmarshal message value", zap.Error(err))
				return nil
			}

			fmt.Printf("Unmarshal object is %+v\n", object)

			// Build the corresponding message task for transformation
			tasks := s.buildMastodonMessageTasks(ctx, object)

			// Print the tasks for debugging
			fmt.Println("Generated tasks:")

			for _, task := range tasks.Tasks {
				fmt.Printf("Task: %+v\n", task)
			}

			tasksChan <- tasks

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// initialize sets up the Kafka consumer and Mastodon client
func (s *dataSource) initialize() (err error) {
	var kafkaBrokers = []string{mastodonSourceEndpoint + ":9092"}

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer(kafkaBrokers, nil)
	if err != nil {
		return fmt.Errorf("create kafka consumer: %w", err)
	}

	s.kafkaConsumer = consumer

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

	// Create a pool to process tasks concurrently (Note: Currently disabled concurrent processing)
	resultPool := pool.NewWithResults[*Task]().WithMaxGoroutines(lo.Ternary(len(object.To) < 20*runtime.NumCPU(), len(object.To), 20*runtime.NumCPU()))

	// Process the ActivityPub object based on its type
	resultPool.Go(func() *Task {
		// Filter the message based on type
		switch object.Type {
		case "Create":
			if note, ok := object.Object.(map[string]interface{}); ok {
				if noteType, exists := note["type"].(string); exists && noteType == "Note" {
					// Example: Create task for a Note object
					fmt.Println("[buildMastodonMessageTasks] Object in Create type (for a Note Object)")
				}
			}
		case "Follow":
			// Process Follow type messages
			fmt.Println("[buildMastodonMessageTasks] Object in Follow type")
		case "Like":
			// Process Like type messages
			fmt.Println("[buildMastodonMessageTasks] Object in Like type")
		default:
			zap.L().Debug("unsupported message type", zap.String("type", object.Type))
		}

		return &Task{
			Network: s.Network(),
			Message: object,
		}
	})

	// Wait for all tasks to complete and handle them
	for _, task := range resultPool.Wait() {
		if task != nil {
			tasks.Tasks = append(tasks.Tasks, task)
		}
	}

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
