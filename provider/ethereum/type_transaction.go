// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethereum

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*transactionMarshal)(nil)

// MarshalJSON marshals as JSON.
func (t Transaction) MarshalJSON() ([]byte, error) {
	type Transaction struct {
		BlockHash   common.Hash     `json:"blockHash"`
		BlockNumber *hexutil.Big    `json:"blockNumber"`
		From        common.Address  `json:"from"`
		Gas         hexutil.Uint64  `json:"gas"`
		GasPrice    *hexutil.Big    `json:"gasPrice"`
		GasTipCap   *hexutil.Big    `json:"maxPriorityFeePerGas"`
		Hash        common.Hash     `json:"hash"`
		Input       hexutil.Bytes   `json:"input"`
		To          *common.Address `json:"to"`
		Index       hexutil.Uint    `json:"index"`
		Value       *hexutil.Big    `json:"value"`
		Type        hexutil.Uint64  `json:"type"`
		ChainID     *hexutil.Big    `json:"chainId"`
	}
	var enc Transaction
	enc.BlockHash = t.BlockHash
	enc.BlockNumber = (*hexutil.Big)(t.BlockNumber)
	enc.From = t.From
	enc.Gas = hexutil.Uint64(t.Gas)
	enc.GasPrice = (*hexutil.Big)(t.GasPrice)
	enc.GasTipCap = (*hexutil.Big)(t.GasTipCap)
	enc.Hash = t.Hash
	enc.Input = t.Input
	enc.To = t.To
	enc.Index = hexutil.Uint(t.Index)
	enc.Value = (*hexutil.Big)(t.Value)
	enc.Type = hexutil.Uint64(t.Type)
	enc.ChainID = (*hexutil.Big)(t.ChainID)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	type Transaction struct {
		BlockHash   *common.Hash    `json:"blockHash"`
		BlockNumber *hexutil.Big    `json:"blockNumber"`
		From        *common.Address `json:"from"`
		Gas         *hexutil.Uint64 `json:"gas"`
		GasPrice    *hexutil.Big    `json:"gasPrice"`
		GasTipCap   *hexutil.Big    `json:"maxPriorityFeePerGas"`
		Hash        *common.Hash    `json:"hash"`
		Input       *hexutil.Bytes  `json:"input"`
		To          *common.Address `json:"to"`
		Index       *hexutil.Uint   `json:"index"`
		Value       *hexutil.Big    `json:"value"`
		Type        *hexutil.Uint64 `json:"type"`
		ChainID     *hexutil.Big    `json:"chainId"`
	}
	var dec Transaction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BlockHash != nil {
		t.BlockHash = *dec.BlockHash
	}
	if dec.BlockNumber != nil {
		t.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.From != nil {
		t.From = *dec.From
	}
	if dec.Gas != nil {
		t.Gas = uint64(*dec.Gas)
	}
	if dec.GasPrice != nil {
		t.GasPrice = (*big.Int)(dec.GasPrice)
	}
	if dec.GasTipCap != nil {
		t.GasTipCap = (*big.Int)(dec.GasTipCap)
	}
	if dec.Hash != nil {
		t.Hash = *dec.Hash
	}
	if dec.Input != nil {
		t.Input = *dec.Input
	}
	if dec.To != nil {
		t.To = dec.To
	}
	if dec.Index != nil {
		t.Index = uint(*dec.Index)
	}
	if dec.Value != nil {
		t.Value = (*big.Int)(dec.Value)
	}
	if dec.Type != nil {
		t.Type = uint64(*dec.Type)
	}
	if dec.ChainID != nil {
		t.ChainID = (*big.Int)(dec.ChainID)
	}
	return nil
}
