package schema

import "github.com/naturalselectionlabs/rss3-node/schema/filter"

type Feed struct {
	ID           string           `json:"id"`
	Owner        string           `json:"owner,omitempty"`
	Network      filter.Network   `json:"network"`
	Index        uint             `json:"index"`
	From         string           `json:"from"`
	To           string           `json:"to"`
	Tag          filter.Tag       `json:"tag"`
	Type         filter.Type      `json:"type"`
	Platform     *filter.Platform `json:"platform,omitempty"`
	Fee          *Fee             `json:"fee,omitempty"`
	TotalActions uint             `json:"total_actions"`
	Actions      []*Action        `json:"actions"`
	Direction    filter.Direction `json:"direction,omitempty"`
	Status       bool             `json:"status"`
	Timestamp    uint64           `json:"timestamp"`
}

// FeedOption is a function that can be used to modify a feed,
// it is used in the feed builder.
type FeedOption func(feed *Feed) error

// WithFeedPlatform is a feed option that sets the platform of the feed.
func WithFeedPlatform(platform filter.Platform) FeedOption {
	return func(feed *Feed) error {
		feed.Platform = &platform

		return nil
	}
}
