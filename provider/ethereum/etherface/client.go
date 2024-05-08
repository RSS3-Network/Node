package etherface

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/avast/retry-go/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ErrorNoResults = errors.New("no results")
)

const (
	EtherfaceDBPath = "/root/node/etherface"
	DefaultAttempts = 3
)

type Client interface {
	Lookup(ctx context.Context, hash string) (string, error)
}

var _ Client = (*etherfaceClient)(nil)

// etherfaceClient is a client for querying the etherface database
type etherfaceClient struct {
	db       *leveldb.DB
	attempts uint
}

// Lookup queries the etherface database for the function name of a given hash
func (h *etherfaceClient) Lookup(ctx context.Context, hash string) (functionName string, err error) {
	retryableFunc := func() error {
		functionName, err = h.query(ctx, hash)
		return err
	}

	retryIfFunc := func(err error) bool {
		return !errors.Is(err, ErrorNoResults)
	}

	if err := retry.Do(retryableFunc, retry.Attempts(h.attempts), retry.RetryIf(retryIfFunc)); err != nil {
		return "", fmt.Errorf("retry attempts: %w", err)
	}

	return functionName, nil
}

// query queries the etherface database for the function name of a given hash
func (h *etherfaceClient) query(_ context.Context, hash string) (string, error) {
	normalizedHash := hash

	if strings.HasPrefix(hash, "0x") {
		normalizedHash = hash[2:]
	}

	data, err := h.db.Get([]byte(normalizedHash), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return "", ErrorNoResults
		}

		return "", fmt.Errorf("failed to get data from leveldb: %w", err)
	}

	return extractFunctionName(string(data)), nil
}

// extractFunctionName extracts the function name from a string before the first bracket
func extractFunctionName(str string) string {
	index := strings.Index(str, "(")
	if index != -1 {
		return str[:index]
	}

	return ""
}

// NewEtherfaceClient creates a new Etherface client based on leveldb
func NewEtherfaceClient() (Client, error) {
	database, err := leveldb.OpenFile(EtherfaceDBPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open leveldb: %w", err)
	}

	instance := etherfaceClient{
		db:       database,
		attempts: DefaultAttempts,
	}

	return &instance, nil
}

type ClientOption func(client *etherfaceClient) error

// WithAttempts sets the number of attempts for querying the etherface database
func WithAttempts(attempts uint) ClientOption {
	return func(h *etherfaceClient) error {
		h.attempts = attempts

		return nil
	}
}
