package mastodon

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/IBM/sarama"
	"github.com/go-playground/form/v4"
)

const (
	DefaultTimeout         = 5 * time.Second
	DefaultAttempts        = 3
	mastodonSourceEndpoint = "34.125.35.124" // External IP of the Mastodon VM instance
)

var _ Client = (*client)(nil)

type Client interface {
	GetKafkaConsumer() sarama.Consumer
}

type client struct {
	endpointURL   *url.URL
	httpClient    *http.Client
	encoder       *form.Encoder
	attempts      uint
	kafkaConsumer sarama.Consumer
}

func (c *client) GetKafkaConsumer() sarama.Consumer {
	return c.kafkaConsumer
}

func NewClient(endpoint string) (Client, error) {
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
	var kafkaBrokers = []string{mastodonSourceEndpoint + ":9092"}

	// Create a new Kafka consumer using the broker endpoint
	consumer, err := sarama.NewConsumer(kafkaBrokers, nil)
	if err != nil {
		return nil, fmt.Errorf("create kafka consumer: %w", err)
	}

	instance.kafkaConsumer = consumer

	if instance.endpointURL, err = url.Parse(endpoint); err != nil {
		return nil, fmt.Errorf("parse endpoint: %w", err)
	}

	return &instance, nil
}
