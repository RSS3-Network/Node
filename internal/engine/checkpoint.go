package engine

import (
	"encoding/json"

	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type CheckpointTransformer interface {
	Import(checkpoint *Checkpoint) error
	Export() (*Checkpoint, error)
}

type Checkpoint struct {
	ID      string          `json:"id"`
	Network filter.Network  `json:"network"`
	Chain   filter.Chain    `json:"chain"`
	Worker  string          `json:"worker"`
	State   json.RawMessage `json:"state"`
}
