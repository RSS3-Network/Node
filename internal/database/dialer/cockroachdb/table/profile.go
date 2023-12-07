package table

import (
	"github.com/lib/pq"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/farcaster/model"
)

var _ model.ProfileTransformer = (*Profile)(nil)

type Profile struct {
	Fid            int64          `gorm:"column:fid"`
	Username       string         `gorm:"column:username"`
	CustodyAddress string         `gorm:"column:custody_address"`
	EthAddresses   pq.StringArray `gorm:"column:eth_addresses;type:text[]"`
}

func (p *Profile) Import(profile *model.Profile) (err error) {
	p.Fid = profile.Fid
	p.Username = profile.Username
	p.CustodyAddress = profile.CustodyAddress
	p.EthAddresses = profile.EthAddresses

	return nil
}

func (p *Profile) Export() (*model.Profile, error) {
	profile := model.Profile{
		Fid:            p.Fid,
		Username:       p.Username,
		CustodyAddress: p.CustodyAddress,
		EthAddresses:   p.EthAddresses,
	}

	return &profile, nil
}
