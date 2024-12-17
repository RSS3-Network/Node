package atproto

import (
	"fmt"
	"strings"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/atproto"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network network.Network
	Message atproto.Message
}

func (t Task) ID() string {
	return strings.TrimPrefix(t.Message.URI, "at://")
}

func (t Task) GetNetwork() network.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return uint64(t.Message.CreatedAt.Unix())
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
