package farcaster

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
)

var _ engine.Task = (*Task)(nil)

type Task struct {
	Network filter.Network
	Message farcaster.Message
}

func (t Task) ID() string {
	return fmt.Sprintf("%s.%s", t.Network, t.Message.Hash)
}

func (t Task) GetNetwork() filter.Network {
	return t.Network
}

func (t Task) GetTimestamp() uint64 {
	return uint64(farcaster.CovertFarcasterTimeToTimestamp(int64(t.Message.Data.Timestamp)))
}

func (t Task) Validate() error {
	return nil
}

func (t Task) BuildFeed(options ...schema.FeedOption) (*schema.Feed, error) {
	feed := schema.Feed{
		ID:        common.HexToHash(t.Message.Hash).String(),
		Network:   t.Network,
		Type:      filter.TypeUnknown,
		Status:    true,
		Actions:   make([]*schema.Action, 0),
		Timestamp: t.GetTimestamp(),
	}

	// Apply feed options.
	for _, option := range options {
		if err := option(&feed); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &feed, nil
}
