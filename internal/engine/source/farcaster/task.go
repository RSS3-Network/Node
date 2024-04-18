package farcaster

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network network.Network
	Message farcaster.Message
}

func (t Task) ID() string {
	return fmt.Sprintf("%s.%s", t.Network, t.Message.Hash)
}

func (t Task) GetNetwork() network.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return uint64(farcaster.CovertFarcasterTimeToTimestamp(int64(t.Message.Data.Timestamp)))
}

func (t Task) Validate() error {
	return nil
}

func (t Task) BuildActivity(options ...activity.Option) (*activity.Activity, error) {
	Tag:      tag. := activity.Activity{
		ID:        common.HexToHash(t.Message.Hash).String(),
		Network:   t.Network,
		Type:      typex.Unknown,
		Status:    true,
		Actions:   make([]*activity.Action, 0),
		Timestamp: t.GetTimestamp(),
	}

	// Apply _activity options.
	for _, option := range options {
		if err := option(&_activity); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &_activity, nil
}
