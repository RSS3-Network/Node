package near

import (
	"encoding/json"
)

type Action struct {
	FunctionCall    *FunctionCallAction    `json:"FunctionCall,omitempty"`
	Transfer        *TransferAction        `json:"Transfer,omitempty"`
	DelegateActions *DelegateActionsAction `json:"DelegateActions,omitempty"`
	Stake           *StakeAction           `json:"Stake,omitempty"`
}

func (a *Action) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		// If it's not a map, try unmarshaling as a string
		var strAction string
		if err := json.Unmarshal(data, &strAction); err != nil {
			return err
		}
		// If it's a string, create a map with the string as the key
		raw = map[string]json.RawMessage{
			strAction: json.RawMessage("{}"),
		}
	}

	for k, v := range raw {
		switch k {
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
		case "Delegate":
			a.DelegateActions = &DelegateActionsAction{}
			if err := json.Unmarshal(v, a.DelegateActions); err != nil {
				return err
			}
		case "Stake":
			a.Stake = &StakeAction{}
			if err := json.Unmarshal(v, a.Stake); err != nil {
				return err
			}
		default:
		}
	}

	return nil
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
