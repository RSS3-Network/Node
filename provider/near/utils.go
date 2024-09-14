package near

import (
	"encoding/base64"
)

// DecodeBase64 decodes a base64 string.
func DecodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
