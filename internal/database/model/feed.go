package model

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
)

type FeedQuery struct {
	ID          *string
	Network     *filter.Network
	Owner       *string
	ActionLimit int
	ActionPage  int
}

type FeedsQuery struct {
	Owner          *string
	Cursor         *schema.Feed
	Status         *bool
	Direction      *filter.Direction
	StartTimestamp *uint64
	EndTimestamp   *uint64
	Platform       *filter.Platform
	Owners         []string
	Network        []filter.Network
	Tags           []filter.Tag
	Types          []filter.Type
	Platforms      []filter.Platform
	Distinct       *bool
	RelatedActions *bool
	Limit          int
	ActionLimit    int
}
