package arweave

import (
	"crypto/sha256"
	"encoding/base64"
)

// PublicKeyToAddress returns the arweave address of the owner.
func PublicKeyToAddress(publicKey string) (string, error) {
	by, err := Base64Decode(publicKey)
	if err != nil {
		return "", err
	}

	addr := sha256.Sum256(by)

	return Base64Encode(addr[:]), nil
}

// Base64Encode returns the base64 encoding of data.
func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// Base64Decode returns the bytes represented by the base64 string.
func Base64Decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
