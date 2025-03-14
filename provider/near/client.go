package near

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/rss3-network/node/v2/provider/httpx"
	"github.com/samber/lo"
)

// Client provides basic RPC methods for NEAR Protocol.
type Client interface {
	BlockByHeight(ctx context.Context, blockHeight *big.Int) (*Block, error)
	ChunkByHash(ctx context.Context, hash string) (*Chunk, error)
	ChunkByHeight(ctx context.Context, blockHeight *big.Int, shardID int) (*Chunk, error)
	TransactionByHash(ctx context.Context, txHash string, senderAccountID string) (*Transaction, error)
	GetBlockHeight(ctx context.Context) (int64, error)
}

var _ Client = (*client)(nil)

type client struct {
	httpClient httpx.Client
	endpoint   string
}

// Dial creates a new client for a given URL.
func Dial(_ context.Context, endpoint string) (Client, error) {
	httpClient, err := httpx.NewHTTPClient(httpx.WithTimeout(15 * time.Second))
	if err != nil {
		return nil, err
	}

	return &client{
		httpClient: httpClient,
		endpoint:   endpoint,
	}, nil
}

func (c *client) post(ctx context.Context, method string, params map[string]interface{}, result interface{}) error {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "dontcare",
		"method":  method,
		"params":  params,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body, err := c.httpClient.Post(ctx, c.endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer lo.Try(body.Close)

	var jsonResponse struct {
		JSONRPC string          `json:"jsonrpc"`
		Result  json.RawMessage `json:"result"`
		ID      string          `json:"id"`
	}

	if err := json.NewDecoder(body).Decode(&jsonResponse); err != nil {
		return err
	}

	return json.Unmarshal(jsonResponse.Result, result)
}

func (c *client) rpcCall(ctx context.Context, method string, params map[string]interface{}, result interface{}) error {
	return c.post(ctx, method, params, result)
}

// ChunkByHash returns the chunk for the given chunk hash.
func (c *client) ChunkByHash(ctx context.Context, hash string) (*Chunk, error) {
	var result Chunk

	err := c.rpcCall(ctx, "chunk", map[string]interface{}{"chunk_id": hash}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ChunkByHeight returns the chunk for the given block height and shard ID.
func (c *client) ChunkByHeight(ctx context.Context, blockHeight *big.Int, shardID int) (*Chunk, error) {
	var result Chunk

	err := c.rpcCall(ctx, "chunk", map[string]interface{}{
		"block_id": blockHeight.Int64(),
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

	err := c.rpcCall(ctx, "tx", map[string]interface{}{"tx_hash": txHash, "sender_account_id": senderAccountID}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// BlockByHeight returns the block for the given block height.
func (c *client) BlockByHeight(ctx context.Context, blockHeight *big.Int) (*Block, error) {
	var result Block

	err := c.rpcCall(ctx, "block", map[string]interface{}{"block_id": blockHeight.Int64()}, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBlockHeight returns the current block height.
func (c *client) GetBlockHeight(ctx context.Context) (int64, error) {
	var result Block

	err := c.rpcCall(ctx, "block", map[string]interface{}{"finality": "final"}, &result)
	if err != nil {
		return 0, err
	}

	return int64(result.Header.Height), nil
}
