package model

import (
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
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
	Platforms      []string
	Distinct       *bool
	RelatedActions *bool
	Limit          int
	ActionLimit    int
}

type ActivitiesMetadataQuery struct {
	Network        *network.Network
	Platform       *decentralized.Platform
	Tag            *tag.Tag
	Type           *schema.Type
	Accounts       []string
	Cursor         *activityx.Activity
	Status         *bool
	StartTimestamp *uint64
	EndTimestamp   *uint64
	Limit          int
	ActionLimit    int
	Metadata       *metadata.Metadata
}
