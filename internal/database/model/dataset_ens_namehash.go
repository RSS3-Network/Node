package model

import "github.com/ethereum/go-ethereum/common"

type ENSNamehashTransformer interface {
	Import(namehash *ENSNamehash) error
	Export() (*ENSNamehash, error)
}

type ENSNamehash struct {
	Hash common.Hash `json:"hash"` // Hexadecimal hash bytes
	Name string      `json:"name"` // ENS name, *.eth
}
