package utils

import "crypto/sha256"

// OwnerToAddress returns the arweave address of the owner.
func OwnerToAddress(owner string) (string, error) {
	by, err := Base64Decode(owner)
	if err != nil {
		return "", err
	}

	addr := sha256.Sum256(by)

	return Base64Encode(addr[:]), nil
}
