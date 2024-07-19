package activitypub

import (
	"fmt"
	"time"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/activitypub"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network network.Network
	Message activitypub.Object
}

func (t Task) ID() string {
	return t.Message.ID
}

func (t Task) GetNetwork() network.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	publishedTimeStamp := t.Message.Published

	// if the Task does not have an empty 'Published' field
	if publishedTimeStamp == "" {
		return 0
	}

	parsedTime, err := time.Parse(time.RFC3339, publishedTimeStamp)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}

	// Convert the time.Time object to a Unix timestamp and cast to uint64
	timestamp := uint64(parsedTime.Unix())

	return timestamp
}

func (t Task) Validate() error {
	return nil
}

func (t Task) BuildActivity(options ...activityx.Option) (*activityx.Activity, error) {
	activity := activityx.Activity{
		ID:        t.ID(),
		Network:   t.Network,
		Type:      typex.Unknown,
		Status:    true,
		Actions:   make([]*activityx.Action, 0),
		Timestamp: t.GetTimestamp(),
	}

	// Apply activity options.
	for _, option := range options {
		if err := option(&activity); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &activity, nil
}
