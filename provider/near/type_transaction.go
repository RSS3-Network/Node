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

type TransactionDetails struct {
	SignerID    string               `json:"signer_id"`
	PublicKey   string               `json:"public_key"`
	Nonce       int64                `json:"nonce"`
	ReceiverID  string               `json:"receiver_id"`
	Actions     []FunctionCallAction `json:"actions"`
	PriorityFee int                  `json:"priority_fee"`
	Signature   string               `json:"signature"`
	Hash        string               `json:"hash"`
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
	Delegate     *DelegateAction     `json:"Delegate,omitempty"`
	FunctionCall *FunctionCallAction `json:"FunctionCall,omitempty"`
}

type DelegateAction struct {
	DelegateAction DelegateActionDetails `json:"delegate_action"`
	Signature      string                `json:"signature"`
}

type DelegateActionDetails struct {
	SenderID       string   `json:"sender_id"`
	ReceiverID     string   `json:"receiver_id"`
	Actions        []Action `json:"actions"`
	Nonce          int64    `json:"nonce"`
	MaxBlockHeight int      `json:"max_block_height"`
	PublicKey      string   `json:"public_key"`
}

type FunctionCallAction struct {
	MethodName string `json:"method_name"`
	Args       string `json:"args"`
	Gas        int64  `json:"gas"`
	Deposit    string `json:"deposit"`
}
