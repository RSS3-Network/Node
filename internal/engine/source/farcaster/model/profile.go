package model

type ProfileTransformer interface {
	Import(profile *Profile) error
	Export() (*Profile, error)
}

type Profile struct {
	Fid            int64    `json:"fid"`
	Username       string   `json:"username"`
	CustodyAddress string   `json:"custody_address"`
	EthAddresses   []string `json:"eth_addresses"`
}
