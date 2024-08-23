package table

import (
	"github.com/lib/pq"
	"github.com/rss3-network/node/internal/database/model"
)

var _ model.ProfileTransformer = (*DatasetFarcasterProfile)(nil)

type DatasetFarcasterProfile struct {
	Fid            int64          `gorm:"column:fid"`
	Username       string         `gorm:"column:username"`
	CustodyAddress string         `gorm:"column:custody_address"`
	EthAddresses   pq.StringArray `gorm:"column:eth_addresses;type:text[]"`
}

func (d *DatasetFarcasterProfile) Import(profile *model.Profile) (err error) {
	d.Fid = profile.Fid
	d.Username = profile.Username
	d.CustodyAddress = profile.CustodyAddress
	d.EthAddresses = profile.EthAddresses

	return nil
}

func (d *DatasetFarcasterProfile) Export() (*model.Profile, error) {
	profile := model.Profile{
		Fid:            d.Fid,
		Username:       d.Username,
		CustodyAddress: d.CustodyAddress,
		EthAddresses:   d.EthAddresses,
	}

	return &profile, nil
}
