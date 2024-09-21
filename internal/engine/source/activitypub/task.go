package activitypub

import (
	"fmt"
	"time"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/activitypub"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"go.uber.org/zap"
)

var _ engine.Task = (*Task)(nil)

// TODO: should be pulled from VSL (NetworkParams contract)
var defaultStartTime = "2024-07-22T00:00:00Z"

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
	// Use default time if Published is empty
	timeStr := t.Message.Published
	if timeStr == "" {
		timeStr = defaultStartTime
	}

	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		zap.L().Info("Error parsinig time")
		return 0
	}
	// Convert the time.Time object to a Unix timestamp and cast to uint64
	return uint64(parsedTime.Unix())
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
