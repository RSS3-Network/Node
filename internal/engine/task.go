package engine

import (
	"github.com/rss3-network/serving-node/schema"
	"github.com/rss3-network/serving-node/schema/filter"
)

type Task interface {
	ID() string
	GetNetwork() filter.Network
	GetTimestamp() uint64
	Validate() error
	BuildFeed(options ...schema.FeedOption) (*schema.Feed, error)
}
