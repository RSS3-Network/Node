package utils

import (
	"encoding/base64"
)

// Base64Encode returns the base64 encoding of data.
func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// Base64Decode returns the bytes represented by the base64 string.
func Base64Decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
