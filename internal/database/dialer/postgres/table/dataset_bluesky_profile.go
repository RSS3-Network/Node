package table

import (
	"time"

	"github.com/rss3-network/node/v2/internal/database/model"
)

type DatasetBlueskyProfile struct {
	DID       string    `gorm:"column:did;primaryKey"`
	Handle    string    `gorm:"column:handle"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (DatasetBlueskyProfile) TableName() string {
	return "dataset_bluesky_profiles"
}

func (d *DatasetBlueskyProfile) Import(profile *model.BlueskyProfile) error {
	d.DID = profile.DID
	d.Handle = profile.Handle
	d.CreatedAt = profile.CreatedAt
	d.UpdatedAt = profile.UpdatedAt

	return nil
}

func (d *DatasetBlueskyProfile) Export() (*model.BlueskyProfile, error) {
	profile := model.BlueskyProfile{
		DID:       d.DID,
		Handle:    d.Handle,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	return &profile, nil
}
