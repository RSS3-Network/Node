package schema

import (
	"encoding/json"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
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
	Status       bool             `json:"status"`
	Timestamp    uint64           `json:"timestamp"`
}

func (f *Feed) MarshalJSON() ([]byte, error) {
	type Filler Feed

	filler := Filler(*f)

	filler.Network = f.Network
	filler.Tag = f.Type.Tag()

	return json.Marshal(filler)
}

type FeedOption func(feed *Feed) error

func WithFeedPlatform(platform filter.Platform) FeedOption {
	return func(feed *Feed) error {
		feed.Platform = lo.ToPtr(platform)

		return nil
	}
}
