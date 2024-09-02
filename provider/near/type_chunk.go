package near

type Chunk struct {
	Author         string              `json:"author"`
	Header         ChunkHeader         `json:"header"`
	CongestionInfo ChunkCongestionInfo `json:"congestion_info"`
	Transactions   []ChunkTransaction  `json:"transactions"`
	Receipts       []ChunkReceipt      `json:"receipts"`
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
}

type ChunkCongestionInfo struct {
	DelayedReceiptsGas  string `json:"delayed_receipts_gas"`
	BufferedReceiptsGas string `json:"buffered_receipts_gas"`
	ReceiptBytes        int    `json:"receipt_bytes"`
	AllowedShard        int    `json:"allowed_shard"`
}

type ChunkTransaction struct {
	SignerID    string        `json:"signer_id"`
	PublicKey   string        `json:"public_key"`
	Nonce       int64         `json:"nonce"`
	ReceiverID  string        `json:"receiver_id"`
	Actions     []ChunkAction `json:"actions"`
	PriorityFee int           `json:"priority_fee"`
	Signature   string        `json:"signature"`
	Hash        string        `json:"hash"`
}

type ChunkAction struct {
	Delegate     *DelegateAction     `json:"Delegate,omitempty"`
	FunctionCall *FunctionCallAction `json:"FunctionCall,omitempty"`
}

type DelegateAction struct {
	DelegateAction DelegateActionDetails `json:"delegate_action"`
	Signature      string                `json:"signature"`
}

type DelegateActionDetails struct {
	SenderID       string        `json:"sender_id"`
	ReceiverID     string        `json:"receiver_id"`
	Actions        []ChunkAction `json:"actions"`
	Nonce          int64         `json:"nonce"`
	MaxBlockHeight int           `json:"max_block_height"`
	PublicKey      string        `json:"public_key"`
}

type FunctionCallAction struct {
	MethodName string `json:"method_name"`
	Args       string `json:"args"`
	Gas        int64  `json:"gas"`
	Deposit    string `json:"deposit"`
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
	Actions             []ChunkAction `json:"actions"`
	IsPromiseYield      bool          `json:"is_promise_yield"`
}
