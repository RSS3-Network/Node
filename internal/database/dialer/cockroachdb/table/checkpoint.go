package table

import (
	"encoding/json"
	"time"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/protocol-go/schema/filter"
)

var _ engine.CheckpointTransformer = (*Checkpoint)(nil)

type Checkpoint struct {
	ID         string          `gorm:"column:id"`
	Network    filter.Network  `gorm:"column:network"`
	Worker     string          `gorm:"column:worker"`
	State      json.RawMessage `gorm:"column:state;type:jsonb"`
	IndexCount int64           `gorm:"column:index_count"`
	CreatedAt  time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Checkpoint) Import(checkpoint *engine.Checkpoint) (err error) {
	c.ID = checkpoint.ID
	c.Network = checkpoint.Network
	c.Worker = checkpoint.Worker
	c.State = checkpoint.State
	c.IndexCount = checkpoint.IndexCount
	c.UpdatedAt = checkpoint.UpdatedAt

	return nil
}

func (c *Checkpoint) Export() (*engine.Checkpoint, error) {
	return &engine.Checkpoint{
		ID:         c.ID,
		Network:    c.Network,
		Worker:     c.Worker,
		State:      c.State,
		IndexCount: c.IndexCount,
		UpdatedAt:  c.UpdatedAt,
	}, nil
}
