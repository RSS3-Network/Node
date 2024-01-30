package engine

import (
	"encoding/json"
	"time"

	"github.com/rss3-network/protocol-go/schema/filter"
)

type CheckpointTransformer interface {
	Import(checkpoint *Checkpoint) error
	Export() (*Checkpoint, error)
}

type Checkpoint struct {
	ID         string          `json:"id"`
	Network    filter.Network  `json:"network"`
	Worker     string          `json:"worker"`
	State      json.RawMessage `json:"state"`
	IndexCount int64           `json:"index_count"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
