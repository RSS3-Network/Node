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
	Outcome   struct {
		Logs        []interface{}              `json:"logs"`
		ReceiptIDs  []string                   `json:"receipt_ids"`
		GasBurnt    int64                      `json:"gas_burnt"`
		TokensBurnt string                     `json:"tokens_burnt"`
		ExecutorID  string                     `json:"executor_id"`
		Status      TransactionOutcomeStatus   `json:"status"`
		Metadata    TransactionOutcomeMetadata `json:"metadata"`
	} `json:"outcome"`
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
	Outcome   struct {
		Logs        []string                   `json:"logs"`
		ReceiptIDs  []string                   `json:"receipt_ids"`
		GasBurnt    int64                      `json:"gas_burnt"`
		TokensBurnt string                     `json:"tokens_burnt"`
		ExecutorID  string                     `json:"executor_id"`
		Status      TransactionStatus          `json:"status"`
		Metadata    TransactionOutcomeMetadata `json:"metadata"`
	} `json:"outcome"`
}

type Action struct {
	FunctionCall *FunctionCallAction `json:"FunctionCall,omitempty"`
	Transfer     *TransferAction     `json:"Transfer,omitempty"`
	Delegate     *DelegateAction     `json:"Delegate,omitempty"`
	// Add other action types as needed
}

type DelegateAction struct {
	DelegateAction struct {
		Actions        []Action `json:"actions"`
		MaxBlockHeight int      `json:"max_block_height"`
		Nonce          int64    `json:"nonce"`
		PublicKey      string   `json:"public_key"`
		ReceiverID     string   `json:"receiver_id"`
		SenderID       string   `json:"sender_id"`
	} `json:"delegate_action"`
	Signature string `json:"signature"`
}

type FunctionCallAction struct {
	Args       string `json:"args"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
	MethodName string `json:"method_name"`
}

type TransferAction struct {
	Deposit string `json:"deposit"`
}
