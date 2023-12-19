package arweave

// Block represents a block on the Arweave network.
type Block struct {
	Nonce     string   `json:"nonce"`
	Timestamp int64    `json:"timestamp"`
	Height    int64    `json:"height"`
	Hash      string   `json:"hash"`
	Txs       []string `json:"txs"`
}
