package table

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/farcaster/model"
	"github.com/samber/lo"
)

var _ model.ProfileTransformer = (*Profile)(nil)

type Profile struct {
	Fid            int64            `gorm:"column:fid"`
	Username       string           `gorm:"column:username"`
	CustodyAddress common.Address   `gorm:"column:custody_address"`
	EthAddresses   []common.Address `gorm:"column:eth_addresses"`
}

func (p *Profile) TableName() string {
	return "farcaster_profiles"
}

func (p *Profile) Import(profile *model.Profile) (err error) {
	p.Fid = profile.Fid
	p.Username = profile.Username
	p.CustodyAddress = common.HexToAddress(profile.CustodyAddress)
	p.EthAddresses = lo.Map(profile.EthAddresses, func(address string, index int) common.Address {
		return common.HexToAddress(address)
	})

	return nil
}

func (p *Profile) Export() (*model.Profile, error) {
	profile := model.Profile{
		Fid:            p.Fid,
		Username:       p.Username,
		CustodyAddress: p.CustodyAddress.String(),
		EthAddresses: lo.Map(p.EthAddresses, func(address common.Address, index int) string {
			return address.String()
		}),
	}

	return &profile, nil
}
