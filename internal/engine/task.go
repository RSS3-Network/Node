package engine

import (
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Task interface {
	ID() string
	GetNetwork() filter.Network
	GetTimestamp() uint64
	Validate() error
	BuildFeed(options ...schema.FeedOption) (*schema.Feed, error)
}
