package near

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
)

// Client provides basic RPC methods for NEAR Protocol.
type Client interface {
	BlockByNumber(ctx context.Context, blockNumber *big.Int) (*Block, error)
	ChunkByHash(ctx context.Context, hash string) (*Chunk, error)
	ChunkByBlockNumber(ctx context.Context, blockNumber *big.Int, shardID int) (*Chunk, error)
	TransactionByHash(ctx context.Context, txHash string, senderAccountID string) (*Transaction, error)
}

var _ Client = (*client)(nil)

type client struct {
	endpoint string
}

// Dial creates a new client for a given URL.
func Dial(_ context.Context, endpoint string) (Client, error) {
	return &client{
		endpoint: endpoint,
	}, nil
}

func (c *client) post(ctx context.Context, method string, params map[string]interface{}, result interface{}) error {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  method,
		"params":  params,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var jsonResponse struct {
		JSONRPC string          `json:"jsonrpc"`
		Result  json.RawMessage `json:"result"`
		ID      string          `json:"id"`
	}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return err
	}

	if err := json.Unmarshal(jsonResponse.Result, result); err != nil {
		return err
	}

	return nil
}

// ChunkByHash returns the chunk for the given chunk hash.
func (c *client) ChunkByHash(ctx context.Context, hash string) (*Chunk, error) {
	var result Chunk

	err := c.post(ctx, "chunk", map[string]interface{}{
		"chunk_id": hash,
	}, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ChunkByBlockNumber returns the chunk for the given block number and shard ID.
func (c *client) ChunkByBlockNumber(ctx context.Context, blockNumber *big.Int, shardID int) (*Chunk, error) {
	var result Chunk

	err := c.post(ctx, "chunk", map[string]interface{}{
		"block_id": blockNumber.Int64(),
		"shard_id": shardID,
	}, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// TransactionByHash returns the transaction for the given transaction hash and sender account ID.
func (c *client) TransactionByHash(ctx context.Context, txHash string, senderAccountID string) (*Transaction, error) {
	var result Transaction

	err := c.post(ctx, "tx", map[string]interface{}{
		"tx_hash":           txHash,
		"sender_account_id": senderAccountID,
		"wait_until":        "EXECUTED",
	}, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// BlockByNumber returns the block for the given block number.
func (c *client) BlockByNumber(ctx context.Context, blockNumber *big.Int) (*Block, error) {
	var result Block

	err := c.post(ctx, "block", map[string]interface{}{
		"block_id": blockNumber.Int64(),
	}, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
