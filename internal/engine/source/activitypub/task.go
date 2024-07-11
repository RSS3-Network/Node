package activitypub

import (
	"fmt"
	"strconv"

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
	timestamp, _ := strconv.ParseUint(publishedTimeStamp, 10, 64)

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
