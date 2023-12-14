package model

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type DatasetMirrorPostTransformer interface {
	Import(post *DatasetMirrorPost) error
	Export() (*DatasetMirrorPost, error)
}

type DatasetMirrorPost struct {
	TransactionID        string
	Height               decimal.Decimal
	Contributor          common.Address
	Title                string
	Content              string
	Version              string
	Language             string
	ContentDigital       string
	OriginContentDigital string
	Timestamp            time.Time
}
