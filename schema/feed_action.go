package schema

import (
	"encoding/json"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type ActionTransformer interface {
	Import(action *Action) error
}

type Action struct {
	Tag      filter.Tag  `json:"tag"`
	Type     filter.Type `json:"type"`
	Platform string      `json:"platform,omitempty"`
	From     string      `json:"from"`
	To       string      `json:"to"`
	Metadata Metadata    `json:"metadata"`
}

func (a *Action) MarshalJSON() ([]byte, error) {
	type Filler Action

	filler := Filler(*a)

	filler.Tag = a.Type.Tag()

	return json.Marshal(filler)
}
