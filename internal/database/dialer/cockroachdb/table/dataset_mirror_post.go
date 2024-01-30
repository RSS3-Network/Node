package table

import (
	"github.com/rss3-network/node/internal/engine/worker/contract/mirror/model"
)

var _ model.DatasetMirrorPostTransformer = (*DatasetMirrorPost)(nil)

// DatasetMirrorPost represents a mirror post for revise logic check.
type DatasetMirrorPost struct {
	ID                  string `gorm:"column:id;primary_key"`
	OriginContentDigest string `gorm:"column:origin_content_digest;index:idx_origin_content_digest"`
}

func (p *DatasetMirrorPost) Import(post *model.DatasetMirrorPost) (err error) {
	p.ID = post.TransactionID
	p.OriginContentDigest = post.OriginContentDigest

	return nil
}

func (p *DatasetMirrorPost) Export() (*model.DatasetMirrorPost, error) {
	post := model.DatasetMirrorPost{
		TransactionID:       p.ID,
		OriginContentDigest: p.OriginContentDigest,
	}

	return &post, nil
}
