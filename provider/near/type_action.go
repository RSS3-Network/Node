package near

import (
	"encoding/json"
)

type Action struct {
	CreateAccount   *CreateAccountAction  `json:"CreateAccount,omitempty"`
	DeployContract  *DeployContractAction `json:"DeployContract,omitempty"`
	FunctionCall    *FunctionCallAction   `json:"FunctionCall,omitempty"`
	Transfer        *TransferAction       `json:"Transfer,omitempty"`
	Stake           *StakeAction          `json:"Stake,omitempty"`
	AddKey          *AddKeyAction         `json:"AddKey,omitempty"`
	DeleteKey       *DeleteKeyAction      `json:"DeleteKey,omitempty"`
	DeleteAccount   *DeleteAccountAction  `json:"DeleteAccount,omitempty"`
	DelegateActions *SignedDelegateAction `json:"Delegate,omitempty"`
}

// nolint:gocognit
func (a *Action) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		var strAction string

		if err := json.Unmarshal(data, &strAction); err != nil {
			return err
		}

		raw = map[string]json.RawMessage{
			strAction: json.RawMessage("{}"),
		}
	}

	for k, v := range raw {
		switch k {
		case "CreateAccount":
			a.CreateAccount = &CreateAccountAction{}
			if err := json.Unmarshal(v, a.CreateAccount); err != nil {
				return err
			}
		case "DeployContract":
			a.DeployContract = &DeployContractAction{}
			if err := json.Unmarshal(v, a.DeployContract); err != nil {
				return err
			}
		case "FunctionCall":
			a.FunctionCall = &FunctionCallAction{}
			if err := json.Unmarshal(v, a.FunctionCall); err != nil {
				return err
			}
		case "Transfer":
			a.Transfer = &TransferAction{}
			if err := json.Unmarshal(v, a.Transfer); err != nil {
				return err
			}
		case "Stake":
			a.Stake = &StakeAction{}
			if err := json.Unmarshal(v, a.Stake); err != nil {
				return err
			}
		case "AddKey":
			a.AddKey = &AddKeyAction{}
			if err := json.Unmarshal(v, a.AddKey); err != nil {
				return err
			}
		case "DeleteKey":
			a.DeleteKey = &DeleteKeyAction{}
			if err := json.Unmarshal(v, a.DeleteKey); err != nil {
				return err
			}
		case "DeleteAccount":
			a.DeleteAccount = &DeleteAccountAction{}
			if err := json.Unmarshal(v, a.DeleteAccount); err != nil {
				return err
			}
		case "Delegate":
			a.DelegateActions = &SignedDelegateAction{}
			if err := json.Unmarshal(v, a.DelegateActions); err != nil {
				return err
			}
		default:
		}
	}

	return nil
}

type CreateAccountAction struct{}

type DeployContractAction struct {
	Code []byte `json:"code"`
}

type FunctionCallAction struct {
	MethodName string `json:"method_name"`
	Args       string `json:"args"`
	Gas        int64  `json:"gas"`
	Deposit    string `json:"deposit"`
}

type TransferAction struct {
	Deposit string `json:"deposit"`
}

type StakeAction struct {
	Stake     string `json:"stake"`
	PublicKey string `json:"public_key"`
}

type AddKeyAction struct {
	PublicKey string    `json:"public_key"`
	AccessKey AccessKey `json:"access_key"`
}

type DeleteKeyAction struct {
	PublicKey string `json:"public_key"`
}

type DeleteAccountAction struct {
	BeneficiaryID string `json:"beneficiary_id"`
}

type SignedDelegateAction struct {
	DelegateAction DelegateAction `json:"delegate_action"`
	Signature      string         `json:"signature"`
}

type DelegateAction struct {
	SenderID       string              `json:"sender_id"`
	ReceiverID     string              `json:"receiver_id"`
	Actions        []NonDelegateAction `json:"actions"`
	Nonce          int64               `json:"nonce"`
	MaxBlockHeight int64               `json:"max_block_height"`
	PublicKey      string              `json:"public_key"`
}

type NonDelegateAction struct{}

type AccessKey struct{}
