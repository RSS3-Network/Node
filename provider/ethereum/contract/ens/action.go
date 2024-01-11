package ens

import "fmt"

var _ fmt.Stringer = (*UpdateAction)(nil)

type UpdateAction string

func (n UpdateAction) String() string {
	return string(n)
}

const (
	FusesSet           UpdateAction = "Fuses"
	AddressChanged     UpdateAction = "Address"
	ContentHashChanged UpdateAction = "ContentHash"
	NameChanged        UpdateAction = "Name"
	PubkeyChanged      UpdateAction = "Pubkey"
)
