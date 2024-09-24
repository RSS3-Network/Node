package table

import (
	"time"

	"github.com/rss3-network/node/internal/database/model"
)

var _ model.MastodonHandleTransformer = (*DatasetMastodonHandle)(nil)

type DatasetMastodonHandle struct {
	Handle      string    `gorm:"column:handle;primaryKey"`
	LastUpdated time.Time `gorm:"column:last_updated;not null"`
}

func (d *DatasetMastodonHandle) Import(handle *model.MastodonHandle) error {
	d.Handle = handle.Handle
	d.LastUpdated = handle.LastUpdated

	return nil
}

func (d *DatasetMastodonHandle) Export() (*model.MastodonHandle, error) {
	handle := model.MastodonHandle{
		Handle:      d.Handle,
		LastUpdated: d.LastUpdated,
	}

	return &handle, nil
}
