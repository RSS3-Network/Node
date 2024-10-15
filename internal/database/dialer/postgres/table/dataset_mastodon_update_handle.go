package table

import (
	"time"

	"github.com/rss3-network/node/internal/database/model"
)

var _ model.MastodonHandleTransformer = (*DatasetMastodonUpdateHandle)(nil)

type DatasetMastodonUpdateHandle struct {
	Handle    string    `gorm:"column:handle;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (d *DatasetMastodonUpdateHandle) Import(handle *model.MastodonHandle) error {
	d.Handle = handle.Handle
	d.UpdatedAt = handle.UpdatedAt
	d.CreatedAt = handle.CreatedAt

	return nil
}

func (d *DatasetMastodonUpdateHandle) Export() (*model.MastodonHandle, error) {
	handle := model.MastodonHandle{
		Handle:    d.Handle,
		UpdatedAt: d.UpdatedAt,
		CreatedAt: d.CreatedAt,
	}

	return &handle, nil
}

func (DatasetMastodonUpdateHandle) TableName() string {
	return "dataset_mastodon_update_handles"
}
