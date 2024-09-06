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
	FunctionCall    *FunctionCallAction    `json:"FunctionCall,omitempty"`
	Transfer        *TransferAction        `json:"Transfer,omitempty"`
	DeployContract  *DeployContractAction  `json:"DeployContract,omitempty"`
	CreateAccount   *CreateAccountAction   `json:"CreateAccount,omitempty"`
	DeleteAccount   *DeleteAccountAction   `json:"DeleteAccount,omitempty"`
	AddKey          *AddKeyAction          `json:"AddKey,omitempty"`
	DeleteKey       *DeleteKeyAction       `json:"DeleteKey,omitempty"`
	DelegateActions *DelegateActionsAction `json:"DelegateActions,omitempty"`
	Stake           *StakeAction           `json:"Stake,omitempty"`
}

type FunctionCallAction struct {
	ReceiverID string `json:"receiver_id"`
	MethodName string `json:"method_name"`
	Args       string `json:"args"`
	Allowance  string `json:"allowance"`
	Deposit    string `json:"deposit"`
	Gas        int64  `json:"gas"`
}

type TransferAction struct {
	Deposit string `json:"deposit"`
}

type DeployContractAction struct {
	Code string `json:"code"`
}

type CreateAccountAction struct{}

type DeleteAccountAction struct {
	BeneficiaryID string `json:"beneficiary_id"`
}

type AddKeyAction struct {
	PublicKey string      `json:"public_key"`
	AccessKey interface{} `json:"access_key"` // Can be FullAccess or FunctionCall
}

type DeleteKeyAction struct {
	PublicKey string `json:"public_key"`
}

type DelegateActionsAction struct {
	Actions        []Action `json:"actions"`
	MaxBlockHeight int      `json:"max_block_height"`
	Nonce          int64    `json:"nonce"`
	PublicKey      string   `json:"public_key"`
	ReceiverID     string   `json:"receiver_id"`
	SenderID       string   `json:"sender_id"`
	Signature      string   `json:"signature"`
}

type StakeAction struct {
	Stake     string `json:"stake"`
	PublicKey string `json:"public_key"`
}

type TransactionDetails struct {
	SignerID    string   `json:"signer_id"`
	PublicKey   string   `json:"public_key"`
	Nonce       int64    `json:"nonce"`
	ReceiverID  string   `json:"receiver_id"`
	Actions     []Action `json:"actions"`
	Signature   string   `json:"signature"`
	Hash        string   `json:"hash"`
	PriorityFee int      `json:"priority_fee"`
}
