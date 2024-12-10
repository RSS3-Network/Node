package model

import "github.com/ethereum/go-ethereum/common"

type ENSNamehash struct {
	Hash common.Hash `json:"hash"` // Hexadecimal hash bytes
	Name string      `json:"name"` // ENS name, *.eth
}
