package schema

import "github.com/naturalselectionlabs/rss3-node/schema/filter"

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
