package farcaster

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/provider/farcaster"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ engine.Task = (*Task)(nil)

type Task struct {
	Chain   filter.ChainFarcaster
	Message farcaster.Message
}

func (t Task) ID() string {
	return fmt.Sprintf("%s.%s.%s", t.Chain.Network(), t.Chain, t.Message.Hash)
}

func (t Task) Network() filter.Network {
	return t.Chain.Network()
}

func (t Task) Timestamp() uint64 {
	return uint64(farcaster.CovertFarcasterTimeToTimestamp(int64(t.Message.Data.Timestamp)))
}

func (t Task) Validate() error {
	return nil
}

func (t Task) BuildFeed(options ...schema.FeedOption) (*schema.Feed, error) {
	feed := schema.Feed{
		ID:        common.HexToHash(t.Message.Hash).String(),
		Chain:     t.Chain,
		Type:      filter.TypeUnknown,
		Status:    true,
		Actions:   make([]*schema.Action, 0),
		Timestamp: t.Timestamp(),
	}

	// Apply feed options.
	for _, option := range options {
		if err := option(&feed); err != nil {
			return nil, fmt.Errorf("apply option: %w", err)
		}
	}

	return &feed, nil
}
