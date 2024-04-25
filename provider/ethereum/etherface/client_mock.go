package etherface

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

type MockClient interface {
	Lookup(ctx context.Context, hash string) (string, error)
}

var _ Client = (*mockEtherfaceClient)(nil)

// mockEtherfaceClient is a client for querying the etherface database
type mockEtherfaceClient struct {
	db *leveldb.DB
}

// Lookup queries the etherface database for the function name of a given hash
func (h *mockEtherfaceClient) Lookup(ctx context.Context, hash string) (string, error) {
	functionName, err := h.query(ctx, hash)
	if err != nil {
		return "", fmt.Errorf("failed to lookup hash: %w", err)
	}

	return functionName, nil
}

// query queries the etherface database for the function name of a given hash
func (h *mockEtherfaceClient) query(_ context.Context, hash string) (string, error) {
	data, err := h.db.Get([]byte(hash), nil)
	if errors.Is(err, leveldb.ErrNotFound) {
		return "", ErrorNoResults
	} else if err != nil {
		return "", fmt.Errorf("failed to get data from leveldb: %w", err)
	}

	return string(data), nil
}

// NewMockEtherfaceClient creates a new Etherface client based on leveldb
func NewMockEtherfaceClient() (MockClient, error) {
	tempDir, err := os.MkdirTemp("", "leveldb_mock")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory for leveldb: %w", err)
	}

	database, err := initLevelDB(tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize leveldb: %w", err)
	}

	instance := &mockEtherfaceClient{
		db: database,
	}

	return instance, nil
}

// initLevelDB initializes a new leveldb database
func initLevelDB(tempDir string) (*leveldb.DB, error) {
	// open the database
	database, err := leveldb.OpenFile(tempDir, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open leveldb: %w", err)
	}

	// mock data
	err = database.Put([]byte("8f283970"), []byte("changeAdmin"), nil)
	if err != nil {
		return nil, err
	}

	return database, nil
}
