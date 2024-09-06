package near

type Transaction struct {
	Status               TransactionStatus  `json:"status"`
	Transaction          TransactionDetails `json:"transaction"`
	TransactionOutcome   TransactionOutcome `json:"transaction_outcome"`
	ReceiptsOutcome      []ReceiptOutcome   `json:"receipts_outcome"`
	FinalExecutionStatus string             `json:"final_execution_status"`
}

type TransactionStatus struct {
	SuccessValue string `json:"SuccessValue"`
}

type TransactionOutcome struct {
	Proof     []TransactionOutcomeProof `json:"proof"`
	BlockHash string                    `json:"block_hash"`
	ID        string                    `json:"id"`
	Outcome   Outcome                   `json:"outcome"`
}

type Outcome struct {
	Logs        []interface{}              `json:"logs"`
	ReceiptIDs  []string                   `json:"receipt_ids"`
	GasBurnt    int64                      `json:"gas_burnt"`
	TokensBurnt string                     `json:"tokens_burnt"`
	ExecutorID  string                     `json:"executor_id"`
	Status      TransactionOutcomeStatus   `json:"status"`
	Metadata    TransactionOutcomeMetadata `json:"metadata"`
}

type TransactionOutcomeProof struct {
	Hash      string `json:"hash"`
	Direction string `json:"direction"`
}

type TransactionOutcomeStatus struct {
	SuccessReceiptID string `json:"SuccessReceiptId"`
}

type TransactionOutcomeMetadata struct {
	Version    int         `json:"version"`
	GasProfile interface{} `json:"gas_profile"`
}

type ReceiptOutcome struct {
	Proof     []TransactionOutcomeProof `json:"proof"`
	BlockHash string                    `json:"block_hash"`
	ID        string                    `json:"id"`
	Outcome   Outcome                   `json:"outcome"`
}

type TransactionDetails struct {
	SignerID    string   `json:"signer_id"`
	PublicKey   string   `json:"public_key"`
	Nonce       uint64   `json:"nonce"`
	ReceiverID  string   `json:"receiver_id"`
	Actions     []Action `json:"actions"`
	Signature   string   `json:"signature"`
	Hash        string   `json:"hash"`
	PriorityFee uint64   `json:"priority_fee"`
}
