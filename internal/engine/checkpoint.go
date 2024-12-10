package engine

import (
	"encoding/json"
	"time"

	"github.com/rss3-network/protocol-go/schema/network"
)

type Checkpoint struct {
	ID         string          `json:"id"`
	Network    network.Network `json:"network"`
	Worker     string          `json:"worker"`
	State      json.RawMessage `json:"state"`
	IndexCount int64           `json:"index_count"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
