package model

import (
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type ActivityQuery struct {
	ID          *string
	Network     *network.Network
	Owner       *string
	ActionLimit int
	ActionPage  int
}

type ActivitiesQuery struct {
	Owner          *string
	Cursor         *activityx.Activity
	Status         *bool
	Direction      *activityx.Direction
	StartTimestamp *uint64
	EndTimestamp   *uint64
	Platform       string
	Owners         []string
	Network        []network.Network
	Tags           []tag.Tag
	Types          []schema.Type
	Platforms      []decentralized.Platform
	Distinct       *bool
	RelatedActions *bool
	Limit          int
	ActionLimit    int
}
