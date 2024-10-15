package table

import (
	"github.com/rss3-network/node/internal/database/model"
)

var _ model.MastodonHandleTransformer = (*DatasetMastodonUpdateHandle)(nil)

type DatasetMastodonUpdateHandle struct {
	Handle      string `gorm:"column:handle;primaryKey"`
	LastUpdated uint64 `gorm:"column:last_updated;not null"`
}

func (d *DatasetMastodonUpdateHandle) Import(handle *model.MastodonHandle) error {
	d.Handle = handle.Handle
	d.LastUpdated = handle.LastUpdated

	return nil
}

func (d *DatasetMastodonUpdateHandle) Export() (*model.MastodonHandle, error) {
	handle := model.MastodonHandle{
		Handle:      d.Handle,
		LastUpdated: d.LastUpdated,
	}

	return &handle, nil
}
