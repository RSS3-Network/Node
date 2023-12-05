package arweave

import (
	"io"
)

type Transaction struct {
	Format     int        `json:"format"`
	ID         string     `json:"id"`
	Owner      string     `json:"owner"` // Need to use utils.Base64Encode(wallet.PubKey.N.Bytes()).
	Tags       []Tag      `json:"tags"`
	Target     string     `json:"target"`
	Quantity   string     `json:"quantity"`
	Data       string     `json:"data"` // Base64 encoded.
	DataReader *io.Reader `json:"-"`    // When data size too big, use data reader, set data to an empty string.
	DataSize   string     `json:"data_size"`
	DataRoot   string     `json:"data_root"`
	Reward     string     `json:"reward"`
	Signature  string     `json:"signature"`
}

type Tag struct {
	Name  string `json:"name" avro:"name"`
	Value string `json:"value" avro:"value"`
}
