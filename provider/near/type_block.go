package near

type Block struct {
	Author string       `json:"author"`
	Chunks []BlockChunk `json:"chunks"`
	Header BlockHeader  `json:"header"`
}

type BlockChunk struct {
	ChunkHash            string         `json:"chunk_hash"`
	PrevBlockHash        string         `json:"prev_block_hash"`
	OutcomeRoot          string         `json:"outcome_root"`
	PrevStateRoot        string         `json:"prev_state_root"`
	EncodedMerkleRoot    string         `json:"encoded_merkle_root"`
	EncodedLength        int            `json:"encoded_length"`
	HeightCreated        int            `json:"height_created"`
	HeightIncluded       int            `json:"height_included"`
	ShardID              int            `json:"shard_id"`
	GasUsed              int64          `json:"gas_used"`
	GasLimit             int64          `json:"gas_limit"`
	RentPaid             string         `json:"rent_paid"`
	ValidatorReward      string         `json:"validator_reward"`
	BalanceBurnt         string         `json:"balance_burnt"`
	OutgoingReceiptsRoot string         `json:"outgoing_receipts_root"`
	TxRoot               string         `json:"tx_root"`
	ValidatorProposals   []interface{}  `json:"validator_proposals"`
	Signature            string         `json:"signature"`
	CongestionInfo       CongestionInfo `json:"congestion_info"`
}

type CongestionInfo struct {
	DelayedReceiptsGas  string `json:"delayed_receipts_gas"`
	BufferedReceiptsGas string `json:"buffered_receipts_gas"`
	ReceiptBytes        int    `json:"receipt_bytes"`
	AllowedShard        int    `json:"allowed_shard"`
}

type BlockHeader struct {
	Approvals             []string      `json:"approvals"`
	BlockBodyHash         string        `json:"block_body_hash"`
	BlockMerkleRoot       string        `json:"block_merkle_root"`
	BlockOrdinal          int           `json:"block_ordinal"`
	ChallengesResult      []interface{} `json:"challenges_result"`
	ChallengesRoot        string        `json:"challenges_root"`
	ChunkHeadersRoot      string        `json:"chunk_headers_root"`
	ChunkMask             []bool        `json:"chunk_mask"`
	ChunkReceiptsRoot     string        `json:"chunk_receipts_root"`
	ChunkTxRoot           string        `json:"chunk_tx_root"`
	ChunksIncluded        int           `json:"chunks_included"`
	EpochID               string        `json:"epoch_id"`
	EpochSyncDataHash     interface{}   `json:"epoch_sync_data_hash"`
	GasPrice              string        `json:"gas_price"`
	Hash                  string        `json:"hash"`
	Height                int           `json:"height"`
	LastDsFinalBlock      string        `json:"last_ds_final_block"`
	LastFinalBlock        string        `json:"last_final_block"`
	LatestProtocolVersion int           `json:"latest_protocol_version"`
	NextBpHash            string        `json:"next_bp_hash"`
	NextEpochID           string        `json:"next_epoch_id"`
	OutcomeRoot           string        `json:"outcome_root"`
	PrevHash              string        `json:"prev_hash"`
	PrevHeight            int           `json:"prev_height"`
	PrevStateRoot         string        `json:"prev_state_root"`
	RandomValue           string        `json:"random_value"`
	RentPaid              string        `json:"rent_paid"`
	Signature             string        `json:"signature"`
	Timestamp             int64         `json:"timestamp"`
	TimestampNanosec      string        `json:"timestamp_nanosec"`
	TotalSupply           string        `json:"total_supply"`
	ValidatorProposals    []interface{} `json:"validator_proposals"`
	ValidatorReward       string        `json:"validator_reward"`
}

type BlockNumberResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Block  `json:"result"`
	ID      string `json:"id"`
}
