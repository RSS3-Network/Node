package activitypub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/avast/retry-go/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

const (
	defaultBlockTime         = 3 * time.Second
	kafkaTopic               = "activitypub_events" // kafka topic name
	MASTODON_SOURCE_ENDPOINT = "34.125.79.79"       // External IP of the Mastodon VM instance
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

			// Uncomment and implement task processing
			//tasks := s.buildMastodonMessageTasks(ctx, msg.Value)
			//tasksChan <- tasks

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// initialize sets up the Kafka consumer and Mastodon client
func (s *dataSource) initialize() (err error) {

	var kafkaBrokers = []string{MASTODON_SOURCE_ENDPOINT + ":9092"}

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
