package mastodon

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/twmb/franz-go/pkg/kgo"
)

const (
	DefaultTimeout  = 3 * time.Second
	DefaultAttempts = 2
)

var _ Client = (*client)(nil)

type Client interface {
	GetKafkaConsumer() *kgo.Client
}

type client struct {
	httpClient    *http.Client
	encoder       *form.Encoder
	attempts      uint
	kafkaConsumer *kgo.Client
}

func (c *client) GetKafkaConsumer() *kgo.Client {
	return c.kafkaConsumer
}

func NewClient(endpoint string, kafkaTopic string) (Client, error) {
	var (
		instance = client{
			httpClient: &http.Client{
				Timeout: DefaultTimeout,
			},
			encoder:  form.NewEncoder(),
			attempts: DefaultAttempts,
		}
	)
	// form the kafka broker endpoint
	var kafkaBrokers = []string{endpoint}

	// Create a new Kafka client
	options := []kgo.Opt{
		kgo.SeedBrokers(kafkaBrokers...),
		kgo.ConsumeTopics(kafkaTopic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtEnd()),
	}
	consumer, err := kgo.NewClient(options...)

	if err != nil {
		return nil, fmt.Errorf("create kafka consumer: %w", err)
	}

	instance.kafkaConsumer = consumer

	return &instance, nil
}
