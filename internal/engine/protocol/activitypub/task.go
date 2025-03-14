package activitypub

import (
	"fmt"
	"time"

	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/node/v2/provider/activitypub"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"go.uber.org/zap"
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
	// Get the defaultStartTime that was pulled from VSL NetworkParams
	defaultStartTime := GetTimestampStart()

	// Use default time if Published is empty
	timeStr := t.Message.Published

	if timeStr == "" {
		return uint64(defaultStartTime)
	}

	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		zap.L().Error("failed to parse time", zap.Error(err))

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
