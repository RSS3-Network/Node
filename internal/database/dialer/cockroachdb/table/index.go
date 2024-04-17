package table

import (
	"fmt"
	"math"
	"time"

	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type Index struct {
	ID        string             `gorm:"column:id"`
	Owner     string             `gorm:"column:owner"`
	Network   network.Network    `gorm:"column:network"`
	Platform  string             `gorm:"column:platform"`
	Index     uint               `gorm:"column:index"`
	Tag       tag.Tag            `gorm:"column:tag"`
	Type      string             `gorm:"column:type"`
	Status    bool               `gorm:"column:status"`
	Direction activity.Direction `gorm:"column:direction"`
	Timestamp time.Time          `gorm:"column:timestamp"`
	CreatedAt time.Time          `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time          `gorm:"column:updated_at;autoUpdateTime"`
}

func (i *Index) TableName() string {
	return "indexes"
}

func (i *Index) PartitionName() string {
	return fmt.Sprintf("%s_%d_q%d", i.TableName(), i.Timestamp.Year(), int(math.Ceil(float64(i.Timestamp.Month())/3)))
}

func (i *Index) Import(_activity *activity.Activity) error {
	i.ID = _activity.ID
	i.Network = _activityNetwork
	i.Platform = _activityPlatform
	i.Index = _activityIndex
	i.Tag = _activity.Type.Tag()
	i.Type = _activity.Type.Name()
	i.Status = _activityStatus
	i.Timestamp = time.Unix(int64(_activityTimestamp), 0)

	return nil
}

type Indexes []*Index

func (i *Indexes) Import(feed []*activity.Activity) error {
	*i = make([]*Index, 0, len(feed))

	for _, feed := range feed {
		var index Index

		if err := index.Import(feed); err != nil {
			return err
		}

		addresses := make(map[string]activity.Direction)

		// Judge the direction of the action.
		for _, action := range _activity.Actions {
			if action.From == action.To {
				addresses[action.To] = activity.DirectionSelf

				continue
			}

			addresses[action.To] = activity.DirectionIn
			addresses[action.From] = activity.DirectionOut
		}

		// The from/to address of the transaction has a higher priority.
		if _activityFrom == _activityTo {
			addresses[_activityTo] = activity.DirectionSelf
		} else {
			addresses[_activityTo] = activity.DirectionIn
			addresses[_activityFrom] = activity.DirectionOut
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
