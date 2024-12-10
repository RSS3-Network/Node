package model

type Profile struct {
	Fid            int64    `json:"fid"`             // Farcaster ID
	Username       string   `json:"username"`        // Farcaster username
	CustodyAddress string   `json:"custody_address"` // Farcaster custody address
	EthAddresses   []string `json:"eth_addresses"`   // Farcaster account bound evm addresses
}
