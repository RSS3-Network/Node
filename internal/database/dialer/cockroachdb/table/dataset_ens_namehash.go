package table

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/internal/database/model"
)

var _ model.ENSNamehashTransformer = (*DatasetENSNamehash)(nil)

type DatasetENSNamehash struct {
	Hash common.Hash `gorm:"column:hash;primary_key"`
	Name string      `gorm:"column:name;index:idx_ensnamehash_name"`
}

func (d *DatasetENSNamehash) Import(ensNamehash *model.ENSNamehash) (err error) {
	d.Hash = ensNamehash.Hash
	d.Name = ensNamehash.Name

	return nil
}

func (d *DatasetENSNamehash) Export() (*model.ENSNamehash, error) {
	profile := model.ENSNamehash{
		Hash: d.Hash,
		Name: d.Name,
	}

	return &profile, nil
}
