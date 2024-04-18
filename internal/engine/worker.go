package engine

import (
	"context"

	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type Worker interface {
	// Name is the name of the worker.
	Name() string
	// Platform returns the display name of the worker as the `platform` in the final Activity response.
	Platform() string
	// Tag the `tag` in the final Activity response.
	Tag() tag.Tag
	// Types the possible `type` value in the final Activity response.
	Types() []*schema.Type
	Filter() SourceFilter
	Match(ctx context.Context, task Task) (bool, error)
	Transform(ctx context.Context, task Task) (*activityx.Activity, error)
}
