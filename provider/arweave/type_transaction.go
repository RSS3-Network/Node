package arweave

import (
	"io"
)

// Transaction represents a transaction on the Arweave network.
type Transaction struct {
	Format     int       `json:"format"`
	ID         string    `json:"id"`
	Owner      string    `json:"owner"`
	Tags       []Tag     `json:"tags"`
	Target     string    `json:"target"`
	Quantity   string    `json:"quantity"`
	Data       string    `json:"data"`
	DataReader io.Reader `json:"-"`
	DataSize   string    `json:"data_size"`
	DataRoot   string    `json:"data_root"`
	Reward     string    `json:"reward"`
	Signature  string    `json:"signature"`
}

// Tag represents a tag on the Arweave network.
type Tag struct {
	Name  string `json:"name" avro:"name"`
	Value string `json:"value" avro:"value"`
}
