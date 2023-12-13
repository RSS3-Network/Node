package table

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/mirror/model"
	"github.com/shopspring/decimal"
)

var _ model.PostTransformer = (*Post)(nil)

type Post struct {
	TransactionID        string          `gorm:"column:transaction_id"`
	Height               decimal.Decimal `gorm:"column:height"`
	Contributor          common.Address  `gorm:"column:contributor"`
	Title                string          `gorm:"column:title"`
	Content              string          `gorm:"column:content"`
	Version              string          `gorm:"column:version"`
	Language             string          `gorm:"column:language"`
	ContentDigital       string          `gorm:"column:content_digital"`
	OriginContentDigital string          `gorm:"column:origin_content_digital"`
	Timestamp            time.Time       `gorm:"column:timestamp"`
}

func (p *Post) Import(post *model.Post) (err error) {
	p.TransactionID = post.TransactionID
	p.Height = post.Height
	p.Contributor = post.Contributor
	p.Title = post.Title
	p.Content = post.Content
	p.Version = post.Version
	p.Language = post.Language
	p.ContentDigital = post.ContentDigital
	p.OriginContentDigital = post.OriginContentDigital
	p.Timestamp = post.Timestamp

	return nil
}

func (p *Post) Export() (*model.Post, error) {
	post := model.Post{
		TransactionID:        p.TransactionID,
		Height:               p.Height,
		Contributor:          p.Contributor,
		Title:                p.Title,
		Content:              p.Content,
		Version:              p.Version,
		Language:             p.Language,
		ContentDigital:       p.ContentDigital,
		OriginContentDigital: p.OriginContentDigital,
		Timestamp:            p.Timestamp,
	}

	return &post, nil
}
