package schema

import (
	"encoding/json"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

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
	Status       bool             `json:"success"`
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

type Feeds []*Feed

var _ json.Unmarshaler = (*Feeds)(nil)

func (f *Feeds) UnmarshalJSON(bytes []byte) error {
	type FeedAlias Feed

	type feed struct {
		*FeedAlias
		TypeX string `json:"type"`
	}

	var temp []*feed

	err := json.Unmarshal(bytes, &temp)
	if err != nil {
		return fmt.Errorf("unmarshal feeds: %w", err)
	}

	for _, v := range temp {
		v.TotalActions = uint(len(v.Actions))

		v.Type, err = filter.TypeString(v.Tag, v.TypeX)
		if err != nil {
			return fmt.Errorf("unmarshal type: %w", err)
		}

		*f = append(*f, (*Feed)(v.FeedAlias))
	}

	return nil
}
