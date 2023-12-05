package schema

import (
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
)

type FeedTransformer interface {
	Import(feed *Feed) error
}

type FeedsTransformer interface {
	Import(feed []*Feed) error
}

type Feed struct {
	ID        string           `json:"id"`
	Network   filter.Network   `json:"network"`
	Chain     filter.Chain     `json:"chain"`
	Index     uint             `json:"index"`
	From      string           `json:"from"`
	To        string           `json:"to"`
	Tag       filter.Tag       `json:"tag"`
	Type      filter.Type      `json:"type"`
	Platform  *filter.Platform `json:"platform,omitempty"`
	Fee       Fee              `json:"fee"`
	Actions   []*Action        `json:"actions"`
	Status    bool             `json:"status"`
	Timestamp uint64           `json:"timestamp"`
}

type FeedOption func(feed *Feed) error

func WithFeedPlatform(platform filter.Platform) FeedOption {
	return func(feed *Feed) error {
		feed.Platform = lo.ToPtr(platform)

		return nil
	}
}
