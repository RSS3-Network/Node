package schema

import (
	"encoding/json"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
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

func (f *Feed) MarshalJSON() ([]byte, error) {
	type Filler Feed

	filler := Filler(*f)

	filler.Network = f.Chain.Network()
	filler.Tag = f.Type.Tag()

	return json.Marshal(filler)
}
