package table

import (
	"fmt"
	"math"
	"time"

	"github.com/rss3-network/node/provider/ethereum"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type Index struct {
	ID        string              `gorm:"column:id"`
	Owner     string              `gorm:"column:owner"`
	Network   network.Network     `gorm:"column:network"`
	Platform  string              `gorm:"column:platform"`
	Index     uint                `gorm:"column:index"`
	Tag       tag.Tag             `gorm:"column:tag"`
	Type      string              `gorm:"column:type"`
	Status    bool                `gorm:"column:status"`
	Direction activityx.Direction `gorm:"column:direction"`
	Timestamp time.Time           `gorm:"column:timestamp"`
	CreatedAt time.Time           `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time           `gorm:"column:updated_at;autoUpdateTime"`
}

func (i *Index) TableName() string {
	return "indexes"
}

func (i *Index) PartitionName() string {
	return fmt.Sprintf("%s_%d_q%d", i.TableName(), i.Timestamp.Year(), int(math.Ceil(float64(i.Timestamp.Month())/3)))
}

func (i *Index) Import(activity *activityx.Activity) error {
	i.ID = activity.ID
	i.Network = activity.Network
	i.Platform = activity.Platform
	i.Index = activity.Index
	i.Tag = activity.Type.Tag()
	i.Type = activity.Type.Name()
	i.Status = activity.Status
	i.Timestamp = time.Unix(int64(activity.Timestamp), 0)

	return nil
}

type Indexes []*Index

func (i *Indexes) Import(activities []*activityx.Activity) error {
	*i = make([]*Index, 0, len(activities))

	for _, activity := range activities {
		var index Index

		if err := index.Import(activity); err != nil {
			return err
		}

		addresses := make(map[string]activityx.Direction)

		// Judge the direction of the action.
		for _, action := range activity.Actions {
			if action.From == action.To {
				addresses[action.To] = activityx.DirectionSelf

				continue
			}

			addresses[action.To] = activityx.DirectionIn
			addresses[action.From] = activityx.DirectionOut
		}

		// The from/to address of the transaction has a higher priority.
		if activity.From == activity.To {
			addresses[activity.To] = activityx.DirectionSelf
		} else {
			addresses[activity.To] = activityx.DirectionIn
			addresses[activity.From] = activityx.DirectionOut
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
