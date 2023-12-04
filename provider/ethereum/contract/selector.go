package contract

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// MethodID calculates the method ID of the function selector.
func MethodID(method string) [4]byte {
	var (
		methodID [4]byte
		hash     = crypto.Keccak256Hash([]byte(method))
	)

	copy(methodID[:], hash[:4])

	return methodID
}

// EventHash calculates the event hash of the event signature.
func EventHash(event string) common.Hash {
	return crypto.Keccak256Hash([]byte(event))
}

// MatchMethodIDs checks if the input matches any of the method IDs.
func MatchMethodIDs(input []byte, methodIDs ...[4]byte) bool {
	for _, methodID := range methodIDs {
		if bytes.HasPrefix(input, methodID[:]) {
			return true
		}
	}

	return false
}

// MatchEventHashes checks if the hash matches any of the event hashes.
func MatchEventHashes(hash common.Hash, eventHashes ...common.Hash) bool {
	for _, eventHash := range eventHashes {
		if hash == eventHash {
			return true
		}
	}

	return false
}

// MatchAddresses checks if the address matches any of the addresses.
func MatchAddresses(address common.Address, addresses ...common.Address) bool {
	for _, addr := range addresses {
		if address == addr {
			return true
		}
	}

	return false
}

// ContainsMethodIDs checks if the code contains all the method IDs.
func ContainsMethodIDs(code []byte, methodIDs ...[4]byte) bool {
	for _, methodID := range methodIDs {
		if exists := bytes.Contains(code, methodID[:]); !exists {
			return exists
		}
	}

	return true
}
