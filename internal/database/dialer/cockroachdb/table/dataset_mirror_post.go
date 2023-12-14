package table

import (
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror/model"
)

var _ model.DatasetMirrorPostTransformer = (*DatasetMirrorPost)(nil)

type DatasetMirrorPost struct {
	ID                   string `gorm:"column:id"`
	OriginContentDigital string `gorm:"column:origin_content_digital"`
}

func (p *DatasetMirrorPost) Import(post *model.DatasetMirrorPost) (err error) {
	p.ID = post.TransactionID
	p.OriginContentDigital = post.OriginContentDigital

	return nil
}

func (p *DatasetMirrorPost) Export() (*model.DatasetMirrorPost, error) {
	post := model.DatasetMirrorPost{
		TransactionID:        p.ID,
		OriginContentDigital: p.OriginContentDigital,
	}

	return &post, nil
}
