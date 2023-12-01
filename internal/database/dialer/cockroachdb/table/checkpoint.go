package table

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

var _ engine.CheckpointTransformer = (*Checkpoint)(nil)

type Checkpoint struct {
	ID        string          `gorm:"column:id"`
	Chain     string          `gorm:"column:chain"`
	Worker    string          `gorm:"column:worker"`
	State     json.RawMessage `gorm:"column:state;type:jsonb"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time       `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Checkpoint) Import(checkpoint *engine.Checkpoint) (err error) {
	c.ID = checkpoint.ID
	c.Chain = checkpoint.Chain.FullName()
	c.Worker = checkpoint.Worker
	c.State = checkpoint.State

	return nil
}

func (c *Checkpoint) Export() (*engine.Checkpoint, error) {
	// TODO Refine it in filter package.
	splits := strings.Split(c.Chain, ".")
	if len(splits) != 2 {
		return nil, fmt.Errorf("invalid chain %s", c.Chain)
	}

	checkpoint := engine.Checkpoint{
		ID:     c.ID,
		Worker: c.Worker,
		State:  c.State,
	}

	var err error

	if checkpoint.Network, err = filter.NetworkString(splits[0]); err != nil {
		return nil, err
	}

	if checkpoint.Chain, err = filter.ChainEthereumString(splits[1]); err != nil {
		return nil, err
	}

	return &checkpoint, nil
}
