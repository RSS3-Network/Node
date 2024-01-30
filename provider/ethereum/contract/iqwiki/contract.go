package iqwiki

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// Wiki
// https://polygonscan.com/address/0xb8aa8cabfba7ee3ccb218a9969aef86dff3b9d2d
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Wiki.abi --pkg iqwiki --type IqWiki --out contract_wiki.go

var (
	AddressWiki = common.HexToAddress("0xb8aA8CabfBa7eE3ccb218a9969AEF86DFf3b9d2D")
	AddressSig  = common.HexToAddress("0x191a41c307373211d08613b68df4031977589069")

	EventPosted = contract.EventHash("Posted(address,string)")
)
