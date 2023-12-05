package types

import "os"

type Transaction struct {
	Format     int      `json:"format"`
	ID         string   `json:"id"`
	LastTx     string   `json:"last_tx"`
	Owner      string   `json:"owner"` // Need to use utils.Base64Encode(wallet.PubKey.N.Bytes()).
	Tags       []Tag    `json:"tags"`
	Target     string   `json:"target"`
	Quantity   string   `json:"quantity"`
	Data       string   `json:"data"` // Base64 encoded.
	DataReader *os.File `json:"-"`    // When data size too big, use data reader, set data to an empty string.
	DataSize   string   `json:"data_size"`
	DataRoot   string   `json:"data_root"`
	Reward     string   `json:"reward"`
	Signature  string   `json:"signature"`

	// Computed when needed.
	Chunks *Chunks `json:"-"`
}

type Chunks struct {
	DataRoot []byte   `json:"data_root"`
	Chunks   []Chunk  `json:"chunks"`
	Proofs   []*Proof `json:"proofs"`
}

type Chunk struct {
	DataHash     []byte
	MinByteRange int
	MaxByteRange int
}

type Proof struct {
	Offest int
	Proof  []byte
}

type Tag struct {
	Name  string `json:"name" avro:"name"`
	Value string `json:"value" avro:"value"`
}
