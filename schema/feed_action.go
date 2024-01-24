package schema

import (
	"encoding/json"
	"fmt"

	"github.com/rss3-network/serving-node/schema/filter"
	"github.com/rss3-network/serving-node/schema/metadata"
)

type Action struct {
	Tag         filter.Tag        `json:"tag"`
	Type        filter.Type       `json:"type"`
	Platform    string            `json:"platform,omitempty"`
	From        string            `json:"from"`
	To          string            `json:"to"`
	Metadata    metadata.Metadata `json:"metadata"`
	RelatedURLs []string          `json:"related_urls,omitempty"`
}

type Actions []*Action

var _ json.Unmarshaler = (*Action)(nil)

func (a *Action) UnmarshalJSON(bytes []byte) error {
	type ActionAlias Action

	type action struct {
		ActionAlias

		TypeX     string          `json:"type"`
		MetadataX json.RawMessage `json:"metadata"`
	}

	var temp action

	err := json.Unmarshal(bytes, &temp)
	if err != nil {
		return fmt.Errorf("unmarshal action: %w", err)
	}

	temp.Type, err = filter.TypeString(temp.Tag, temp.TypeX)
	if err != nil {
		return fmt.Errorf("invalid action type: %w", err)
	}

	temp.Metadata, err = metadata.Unmarshal(temp.Type, temp.MetadataX)
	if err != nil {
		return fmt.Errorf("invalid action metadata: %w", err)
	}

	*a = Action(temp.ActionAlias)

	return nil
}
