package near

type Chunk struct {
	Author       string               `json:"author"`
	Header       ChunkHeader          `json:"header"`
	Transactions []TransactionDetails `json:"transactions"`
	Receipts     []ChunkReceipt       `json:"receipts"`
}

type ChunkHeader struct {
	ChunkHash            string        `json:"chunk_hash"`
	PrevBlockHash        string        `json:"prev_block_hash"`
	OutcomeRoot          string        `json:"outcome_root"`
	PrevStateRoot        string        `json:"prev_state_root"`
	EncodedMerkleRoot    string        `json:"encoded_merkle_root"`
	EncodedLength        int           `json:"encoded_length"`
	HeightCreated        int           `json:"height_created"`
	HeightIncluded       int           `json:"height_included"`
	ShardID              int           `json:"shard_id"`
	GasUsed              int64         `json:"gas_used"`
	GasLimit             int64         `json:"gas_limit"`
	RentPaid             string        `json:"rent_paid"`
	ValidatorReward      string        `json:"validator_reward"`
	BalanceBurnt         string        `json:"balance_burnt"`
	OutgoingReceiptsRoot string        `json:"outgoing_receipts_root"`
	TxRoot               string        `json:"tx_root"`
	ValidatorProposals   []interface{} `json:"validator_proposals"`
	Signature            string        `json:"signature"`
	CongestionInfo       struct {
		DelayedReceiptsGas  string `json:"delayed_receipts_gas"`
		BufferedReceiptsGas string `json:"buffered_receipts_gas"`
		ReceiptBytes        int    `json:"receipt_bytes"`
		AllowedShard        int    `json:"allowed_shard"`
	} `json:"congestion_info"`
}

type ChunkReceipt struct {
	PredecessorID string        `json:"predecessor_id"`
	ReceiverID    string        `json:"receiver_id"`
	ReceiptID     string        `json:"receipt_id"`
	Receipt       ReceiptAction `json:"receipt"`
	Priority      int           `json:"priority"`
}

type ReceiptAction struct {
	Action ReceiptActionDetails `json:"Action"`
}

type ReceiptActionDetails struct {
	SignerID            string        `json:"signer_id"`
	SignerPublicKey     string        `json:"signer_public_key"`
	GasPrice            string        `json:"gas_price"`
	OutputDataReceivers []interface{} `json:"output_data_receivers"`
	InputDataIDs        []interface{} `json:"input_data_ids"`
	Actions             []interface{} `json:"actions"` // Change to interface{} to accept any action type
	IsPromiseYield      bool          `json:"is_promise_yield"`
}
