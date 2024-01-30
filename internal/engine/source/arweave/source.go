package arweave

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/arweave/bundle"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

const (
	// The block time in Arweave mainnet is designed to be approximately 2 minutes.
	defaultBlockTime = 120 * time.Second
)

// TODO get from command line arguments
var bundlrNodes = []string{
	"OXcT1sVRSA5eGwt2k6Yuz8-3e3g9WJi5uSE99CWqsBs", // Bundlr Node 1
	"ZE0N-8P9gXkhtK-07PQu9d8me5tGDxa_i4Mee5RzVYg", // Bundlr Node 2
}

// Ensure that source implements Source.
var _ engine.Source = (*source)(nil)

type source struct {
	config        *config.Module
	option        *Option
	filter        *Filter
	arweaveClient arweave.Client
	state         State
	pendingState  State
}

func (s *source) Network() filter.Network {
	return s.config.Network
}

func (s *source) State() json.RawMessage {
	return lo.Must(json.Marshal(s.state))
}

func (s *source) Start(ctx context.Context, tasksChan chan<- []engine.Task, errorChan chan<- error) {
	// Initialize source.
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize source: %w", err)

		return
	}

	// Start a goroutine to poll blocks.
	go func() {
		errorChan <- s.pollBlocks(ctx, tasksChan, s.filter)
	}()
}

// initialize initializes the source.
func (s *source) initialize() (err error) {
	// Initialize arweave client.
	if s.arweaveClient, err = arweave.NewClient(); err != nil {
		return fmt.Errorf("create arweave client: %w", err)
	}

	return nil
}

func (s *source) pollBlocks(ctx context.Context, tasksChan chan<- []engine.Task, filter *Filter) error {
	var (
		blockHeightLatestRemote int64
		err                     error
	)

	// Get start block height from config
	// if not set, use default value 0
	if s.option.BlockHeightStart != nil && s.option.BlockHeightStart.Uint64() > s.state.BlockHeight {
		s.pendingState.BlockHeight = s.option.BlockHeightStart.Uint64()
		s.state.BlockHeight = s.option.BlockHeightStart.Uint64()
	}

	// Get target block height from config
	// if not set, use the latest block height from arweave network
	if s.option.BlockHeightTarget != nil {
		zap.L().Info("block height target", zap.Uint64("block.height.target", s.option.BlockHeightTarget.Uint64()))
		blockHeightLatestRemote = int64(s.option.BlockHeightTarget.Uint64())
	} else {
		// Get remote block height from arweave network.
		blockHeightLatestRemote, err = s.arweaveClient.GetBlockHeight(ctx)
		if err != nil {
			return fmt.Errorf("get latest block height: %w", err)
		}

		zap.L().Info("get latest block height", zap.Int64("block.height", blockHeightLatestRemote))
	}

	for {
		if s.option.BlockHeightTarget != nil && s.option.BlockHeightTarget.Uint64() <= s.state.BlockHeight {
			break
		}

		// Check if block height is latest.
		if s.state.BlockHeight >= uint64(blockHeightLatestRemote) {
			// Get the latest block height from arweave network for reconfirming.
			if blockHeightLatestRemote, err = s.arweaveClient.GetBlockHeight(ctx); err != nil {
				return fmt.Errorf("get latest block height: %w", err)
			}

			zap.L().Info("get latest block height", zap.Int64("block.height", blockHeightLatestRemote))

			if s.state.BlockHeight >= uint64(blockHeightLatestRemote) {
				// Wait for the next block on arweave network.
				time.Sleep(defaultBlockTime)
			}

			continue
		}

		// set the default value of RPCThreadBlocks to 1
		if s.option.RPCThreadBlocks == 0 {
			s.option.RPCThreadBlocks = 1
		}

		// Pull blocks
		blockHeightEnd := lo.Min([]uint64{
			uint64(blockHeightLatestRemote),
			s.state.BlockHeight + s.option.RPCThreadBlocks - 1,
		})

		// Pull blocks by range.
		blocks, err := s.batchPullBlocksByRange(ctx, s.state.BlockHeight, blockHeightEnd)
		if err != nil {
			return fmt.Errorf("batch pull blocks: %w", err)
		}

		// Pull transactions.
		transactionIDs := lo.FlatMap(blocks, func(block *arweave.Block, _ int) []string {
			return block.Txs
		})

		// Batch pull transactions by ids.
		transactions, err := s.batchPullTransactions(ctx, transactionIDs)
		if err != nil {
			return fmt.Errorf("batch pull transactions: %w", err)
		}

		// Filter transactions by owner.
		transactions = s.filterOwnerTransaction(transactions, filter.OwnerAddresses)

		// Pull transaction data.
		if err := s.batchPullData(ctx, transactions); err != nil {
			return fmt.Errorf("batch pull data: %w", err)
		}

		// Decode Bundle transactions group by block.
		for index, block := range blocks {
			bundleTransactions, err := s.batchPullBundleTransactions(ctx, s.GroupBundleTransactions(transactions, block))
			if err != nil {
				return fmt.Errorf("pull bundle transacctions: %w", err)
			}

			for _, bundleTransaction := range bundleTransactions {
				blocks[index].Txs = append(block.Txs, bundleTransaction.ID)
			}

			transactions = append(transactions, bundleTransactions...)
		}

		// Discard the Bundle transaction itself.
		transactions = s.discardRootBundleTransaction(transactions)

		tasks, err := s.buildTasks(blocks, transactions)
		if err != nil {
			return fmt.Errorf("build tasks: %w", err)
		}

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- tasks

		// Update state by two phase commit to avoid data inconsistency.
		s.state = s.pendingState
		s.pendingState.BlockHeight++
	}

	return nil
}

// batchPullBlocksByRange pulls blocks by range, from local state block height to remote block height.
func (s *source) batchPullBlocksByRange(ctx context.Context, blockHeightStart, blockHeightEnd uint64) ([]*arweave.Block, error) {
	zap.L().Info("begin to batch pull transactions by range", zap.Uint64("block.height.start", blockHeightStart), zap.Uint64("block.height.end", blockHeightEnd))

	// Pull blocks by block heights.
	blockHeights := lo.Map(lo.RangeWithSteps(blockHeightStart, blockHeightEnd+1, 1), func(blockHeight uint64, _ int) *big.Int {
		return new(big.Int).SetUint64(blockHeight)
	})

	blocks, err := s.batchPullBlocks(ctx, blockHeights)
	if err != nil {
		return nil, fmt.Errorf("batch pull blocks: %w", err)
	}

	return blocks, nil
}

// batchPullBlocks pulls blocks by block heights.
func (s *source) batchPullBlocks(ctx context.Context, blockHeights []*big.Int) ([]*arweave.Block, error) {
	zap.L().Info("begin to pull blocks", zap.Int("blocks", len(blockHeights)))

	resultPool := pool.NewWithResults[*arweave.Block]().
		WithContext(ctx).
		WithCancelOnError()

	for _, blockHeight := range blockHeights {
		blockHeight := blockHeight

		resultPool.Go(func(ctx context.Context) (*arweave.Block, error) {
			return s.arweaveClient.GetBlockByHeight(ctx, blockHeight.Int64())
		})
	}

	return resultPool.Wait()
}

// batchPullTransactions pulls transactions by transaction ids.
func (s *source) batchPullTransactions(ctx context.Context, transactionIDs []string) ([]*arweave.Transaction, error) {
	zap.L().Info("begin to pull transactions", zap.Int("transactions", len(transactionIDs)))

	resultPool := pool.NewWithResults[*arweave.Transaction]().
		WithContext(ctx).
		WithCancelOnError()

	for _, transactionID := range transactionIDs {
		transactionID := transactionID

		resultPool.Go(func(ctx context.Context) (*arweave.Transaction, error) {
			return s.arweaveClient.GetTransactionByID(ctx, transactionID)
		})
	}

	return resultPool.Wait()
}

// batchPullData pulls data by transactions.
func (s *source) batchPullData(ctx context.Context, transactions []*arweave.Transaction) error {
	resultPool := pool.New().
		WithContext(ctx).
		WithCancelOnError()

	for index, transaction := range transactions {
		index, transaction := index, transaction

		resultPool.Go(func(ctx context.Context) error {
			response, err := s.arweaveClient.GetTransactionData(ctx, transaction.ID)
			if err != nil {
				return fmt.Errorf("fetch transaction data: %w", err)
			}

			defer lo.Try(response.Close)

			buffer := new(bytes.Buffer)
			if _, err := io.Copy(base64.NewEncoder(base64.RawURLEncoding, buffer), response); err != nil {
				return fmt.Errorf("read and encode response: %w", err)
			}

			transactions[index].Data = buffer.String()

			return nil
		})
	}

	return resultPool.Wait()
}

// batchPullBundleTransactions pulls bundle transactions by transaction ids.
func (s *source) batchPullBundleTransactions(ctx context.Context, transactionIDs []string) ([]*arweave.Transaction, error) {
	zap.L().Info("begin to pull and filter bundle transactions", zap.Int("transactions", len(transactionIDs)))

	resultPool := pool.NewWithResults[[]*arweave.Transaction]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, transactionID := range transactionIDs {
		transactionID := transactionID

		resultPool.Go(func(ctx context.Context) ([]*arweave.Transaction, error) {
			bundleTransactions := make([]*arweave.Transaction, 0)

			response, err := s.arweaveClient.GetTransactionData(ctx, transactionID)
			if err != nil {
				return nil, fmt.Errorf("fetch transaction: %w", err)
			}

			defer lo.Try(response.Close)

			decoder := bundle.NewDecoder(response)

			header, err := decoder.DecodeHeader()
			if err != nil {
				// Ignore invalid bundle transaction.
				zap.L().Error("discard a invalid bundle transaction", zap.String("transaction_id", transactionID))

				return nil, nil
			}

			for index := 0; decoder.Next(); index++ {
				dataItemInfo := header.DataItemInfos[index]

				dataItem, err := decoder.DecodeDataItem()
				if err != nil {
					// Ignore invalid signature and data length.
					zap.L().Error("decode data item", zap.Error(err), zap.String("transaction_id", transactionID))

					return nil, nil
				}

				bundleTransaction := arweave.Transaction{
					Format: 2,
					ID:     dataItemInfo.ID,
					Owner:  dataItem.Owner,
					Tags: lo.Map(dataItem.Tags, func(tag bundle.Tag, _ int) arweave.Tag {
						return arweave.Tag{
							Name:  arweave.Base64Encode(tag.Name),
							Value: arweave.Base64Encode(tag.Value),
						}
					}),
					Target:    dataItem.Target,
					Signature: dataItem.Signature,
				}

				// TODO: Match and filter bundle transactions.

				data, err := io.ReadAll(dataItem)
				if err != nil {
					return nil, fmt.Errorf("read data item %s: %w", dataItemInfo.ID, err)
				}

				bundleTransaction.Data = arweave.Base64Encode(data)
				bundleTransaction.DataSize = strconv.Itoa(len(bundleTransaction.Data))

				bundleTransactions = append(bundleTransactions, &bundleTransaction)
			}

			return bundleTransactions, nil
		})
	}

	bundleTransactions, err := resultPool.Wait()
	if err != nil {
		return nil, fmt.Errorf("wait result pool: %w", err)
	}

	return lo.Flatten(bundleTransactions), nil
}

// GroupBundleTransactions groups bundle transactions by block.
func (s *source) GroupBundleTransactions(transactions []*arweave.Transaction, block *arweave.Block) []string {
	return lo.FilterMap(transactions, func(transaction *arweave.Transaction, _ int) (string, bool) {
		hasBundleFormatTag := lo.ContainsBy(transaction.Tags, func(tag arweave.Tag) bool {
			tagName, err := base64.RawURLEncoding.DecodeString(tag.Name)
			if err != nil {
				return false
			}

			tagValue, err := base64.RawURLEncoding.DecodeString(tag.Value)
			if err != nil {
				return false
			}

			return strings.EqualFold(string(tagName), "Bundle-Format") && strings.EqualFold(string(tagValue), "binary")
		})

		hasBundleVersionTag := lo.ContainsBy(transaction.Tags, func(tag arweave.Tag) bool {
			tagName, err := base64.RawURLEncoding.DecodeString(tag.Name)
			if err != nil {
				return false
			}

			tagValue, err := base64.RawURLEncoding.DecodeString(tag.Value)
			if err != nil {
				return false
			}

			return strings.EqualFold(string(tagName), "Bundle-Version") && strings.EqualFold(string(tagValue), "2.0.0")
		})

		if !(hasBundleFormatTag && hasBundleVersionTag) {
			return "", false
		}

		if !lo.Contains(block.Txs, transaction.ID) {
			return "", false
		}

		owner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			zap.L().Error("invalid owner of transaction", zap.String("transaction_id", transaction.ID), zap.Error(err))

			return "", false
		}

		return transaction.ID, lo.Contains(bundlrNodes, owner)
	})
}

// discardRootBundleTransaction discards the root bundle transaction.
func (s *source) discardRootBundleTransaction(transactions []*arweave.Transaction) []*arweave.Transaction {
	return lo.Filter(transactions, func(transaction *arweave.Transaction, _ int) bool {
		transactionOwner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			return false
		}

		return !lo.Contains(bundlrNodes, transactionOwner)
	})
}

// filterOwnerTransaction filters owner transactions.
func (s *source) filterOwnerTransaction(transactions []*arweave.Transaction, ownerAddress []string) []*arweave.Transaction {
	return lo.Filter(transactions, func(transaction *arweave.Transaction, _ int) bool {
		transactionOwner, err := arweave.PublicKeyToAddress(transaction.Owner)
		if err != nil {
			return false
		}

		return lo.Contains(ownerAddress, transactionOwner)
	})
}

// buildTasks builds tasks from blocks and transactions.
func (s *source) buildTasks(blocks []*arweave.Block, transactions []*arweave.Transaction) ([]engine.Task, error) {
	tasks := make([]engine.Task, 0)

	for _, transaction := range transactions {
		block, _ := lo.Find(blocks, func(block *arweave.Block) bool {
			return lo.Contains(block.Txs, transaction.ID)
		})

		tasks = append(tasks, &Task{
			Network:     s.Network(),
			Block:       *block,
			Transaction: *transaction,
		})
	}

	return tasks, nil
}

// NewSource creates a new arweave source.
func NewSource(config *config.Module, sourceFilter engine.SourceFilter, checkpoint *engine.Checkpoint) (engine.Source, error) {
	var (
		state State
		err   error
	)

	// Initialize state from checkpoint.
	if checkpoint != nil {
		if err := json.Unmarshal(checkpoint.State, &state); err != nil {
			return nil, err
		}
	}

	instance := source{
		config:       config,
		filter:       new(Filter), // Set a default filter for the source.
		state:        state,
		pendingState: state, // Default pending state is equal to the current state.
	}

	// Initialize filter.
	if sourceFilter != nil {
		var ok bool
		if instance.filter, ok = sourceFilter.(*Filter); !ok {
			return nil, fmt.Errorf("invalid source filter type %T", sourceFilter)
		}
	}

	if instance.option, err = NewOption(config.Parameters); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &instance, nil
}
