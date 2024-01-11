package erc1577

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

func Parse(contentHash []byte) (path string, err error) {
	contentCode, cursor, err := varint.FromUvarint(contentHash)
	if err != nil {
		return "", fmt.Errorf("invaild content hash: %w", err)
	}

	var protocol string

	switch contentCode := multicodec.Code(contentCode); contentCode {
	case multicodec.Ipfs:
		protocol = "ipfs"
	case multicodec.Ipns:
		protocol = "ipns"
	default:
		return "", fmt.Errorf("unsupported content code: %d", contentCode)
	}

	_, contentCID, err := cid.CidFromBytes(contentHash[cursor:])
	if err != nil {
		return "", fmt.Errorf("invalid content hash: %w", err)
	}

	return fmt.Sprintf("/%s/%s", protocol, contentCID.String()), nil
}
