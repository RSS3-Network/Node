package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rss3-network/node/internal/stream"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

type Client struct {
	kafkaClient *kgo.Client
	topic       string
}

func New(ctx context.Context, uri, topic string) (stream.Client, error) {
	brokers := strings.Split(uri, ",")

	if len(brokers) == 0 {
		return nil, fmt.Errorf("invalid uri: %s", uri)
	}

	kafkaClient, err := kgo.NewClient([]kgo.Opt{kgo.SeedBrokers(brokers...)}...)

	if err != nil {
		return nil, fmt.Errorf("new kafka client: %w", err)
	}

	kafkaAdminClient := kadm.NewClient(kafkaClient)

	topicDetails, err := kafkaAdminClient.ListTopics(ctx)

	if err != nil {
		return nil, fmt.Errorf("list topics: %w", err)
	}

	if _, exists := topicDetails[topic]; !exists {
		if _, err = kafkaAdminClient.CreateTopic(ctx, 1, 1, nil, topic); err != nil {
			return nil, fmt.Errorf("create %s topic: %w", topic, err)
		}
	}

	return &Client{
		kafkaClient: kafkaClient,
		topic:       topic,
	}, nil
}

func (c *Client) PushActivities(ctx context.Context, activities []*activityx.Activity) error {
	records := make([]*kgo.Record, 0, len(activities))

	for _, activity := range activities {
		record, err := c.encodeActivity(activity)
		if err != nil {
			return fmt.Errorf("encode activity %s: %w", activity.ID, err)
		}

		records = append(records, record)
	}

	produceResults := c.kafkaClient.ProduceSync(ctx, records...)

	if err := produceResults.FirstErr(); err != nil {
		return fmt.Errorf("push activities: %w", err)
	}

	zap.L().Info("pushed activities", zap.Int("records", len(records)))

	return nil
}

func (c *Client) encodeActivity(activity *activityx.Activity) (*kgo.Record, error) {
	value, err := json.Marshal(activity)
	if err != nil {
		return nil, err
	}

	record := kgo.Record{
		Topic: c.topic,
		Key:   []byte(activity.ID),
		Value: value,
	}

	return &record, nil
}
