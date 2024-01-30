package table

import (
	"fmt"
	"math"
	"time"

	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
)

type Index struct {
	ID        string           `gorm:"column:id"`
	Owner     string           `gorm:"column:owner"`
	Network   filter.Network   `gorm:"column:network"`
	Platform  *filter.Platform `gorm:"column:platform"`
	Index     uint             `gorm:"column:index"`
	Tag       filter.Tag       `gorm:"column:tag"`
	Type      string           `gorm:"column:type"`
	Status    bool             `gorm:"column:status"`
	Direction filter.Direction `gorm:"column:direction"`
	Timestamp time.Time        `gorm:"column:timestamp"`
	CreatedAt time.Time        `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time        `gorm:"column:updated_at;autoUpdateTime"`
}

func (i *Index) TableName() string {
	return "indexes"
}

func (i *Index) PartitionName() string {
	return fmt.Sprintf("%s_%d_q%d", i.TableName(), i.Timestamp.Year(), int(math.Ceil(float64(i.Timestamp.Month())/3)))
}

func (i *Index) Import(feed *schema.Feed) error {
	i.ID = feed.ID
	i.Network = feed.Network
	i.Platform = feed.Platform
	i.Index = feed.Index
	i.Tag = feed.Type.Tag()
	i.Type = feed.Type.Name()
	i.Status = feed.Status
	i.Timestamp = time.Unix(int64(feed.Timestamp), 0)

	return nil
}

type Indexes []*Index

func (i *Indexes) Import(feed []*schema.Feed) error {
	*i = make([]*Index, 0, len(feed))

	for _, feed := range feed {
		var index Index

		if err := index.Import(feed); err != nil {
			return err
		}

		addresses := make(map[string]filter.Direction)

		// Judge the direction of the action.
		for _, action := range feed.Actions {
			if action.From == action.To {
				addresses[action.To] = filter.DirectionSelf

				continue
			}

			addresses[action.To] = filter.DirectionIn
			addresses[action.From] = filter.DirectionOut
		}

		// The from/to address of the transaction has a higher priority.
		if feed.From == feed.To {
			addresses[feed.To] = filter.DirectionSelf
		} else {
			addresses[feed.To] = filter.DirectionIn
			addresses[feed.From] = filter.DirectionOut
		}

		for address, direction := range addresses {
			if address == ethereum.AddressGenesis.String() || address == "" {
				continue
			}

			data := index
			data.Owner = address
			data.Direction = direction

			*i = append(*i, &data)
		}
	}

	return nil
}
