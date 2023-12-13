package model

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type PostTransformer interface {
	Import(post *Post) error
	Export() (*Post, error)
}

type Post struct {
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
