package arweave

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave/bundle"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
)

const (
	// The block time in Arweave mainnet is designed to be approximately 2 minutes.
	defaultBlockTime = 120 * time.Second
)

var _ engine.Source = (*source)(nil)

type source struct {
	config        *engine.Config
	arweaveClient arweave.Client
	state         *State
}

func (s *source) Name() string {
	//	TODO Implement it.
	return "arweave"
}

func (s *source) Chain() filter.Chain {
	return filter.ChainArweaveMainnet
}

func (s *source) Start(ctx context.Context, tasksChan chan<- []engine.Task, errorChan chan<- error) {
	if err := s.initialize(); err != nil {
		errorChan <- fmt.Errorf("initialize source: %w", err)

		return
	}

	go func() {
		errorChan <- s.pollBlocks(ctx, tasksChan)
	}()
}

func (s *source) initialize() (err error) {
	// TODO Load state from checkpoint.
	s.state = &State{
		ChainID:     filter.ChainArweaveMainnet.ID(),
		BlockNumber: 0,
	}

	if s.arweaveClient, err = arweave.NewClient(); err != nil {
		return fmt.Errorf("create arweave client: %w", err)
	}

	return nil
}

func (s *source) pollBlocks(ctx context.Context, tasksChan chan<- []engine.Task) error {
	var blockNumberLatestLocal uint64

	// Refresh the remote block number.
	blockNumberLatestRemote, err := s.arweaveClient.GetBlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("get latest block number: %w", err)
	}

	// TODO End polling at completion if target block height is specified.
	for {
		// The local block number is equal than the remote block number.
		if s.state.BlockNumber >= blockNumberLatestLocal || blockNumberLatestLocal == 0 {
			// RPC providers may incorrectly shunt the request to a lagging node.
			if uint64(blockNumberLatestRemote) <= blockNumberLatestLocal {
				zap.L().Info("waiting new block", zap.Uint64("block.number.local", s.state.BlockNumber), zap.Uint64("block.number.remote", blockNumberLatestLocal), zap.Duration("block.time", defaultBlockTime))

				time.Sleep(defaultBlockTime)
			} else {
				// TODO Need to handle block reorganization.
				blockNumberLatestLocal = uint64(blockNumberLatestRemote)
			}

			continue
		}

		// Pull blocks
		blocks, err := s.batchPullBlocksByRange(ctx, blockNumberLatestLocal, uint64(blockNumberLatestRemote))
		if err != nil {
			return fmt.Errorf("batch pull blocks: %w", err)
		}

		// TODO filter blocks by filter option.

		// Pull transactions
		transactionIDs := lo.FlatMap(blocks, func(block *types.Block, _ int) []string {
			return block.Txs
		})

		transactions, err := s.batchPullTransactions(ctx, transactionIDs)
		if err != nil {
			return fmt.Errorf("batch pull transactions: %w", err)
		}

		// TODO filter transactions by filter option.

		// pull transaction data
		// TODO add filter option
		if err := s.batchPullData(ctx, transactions); err != nil {
			return fmt.Errorf("batch pull data: %w", err)
		}

		// Filter and decode Bundle transactions group by block
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

		tasks, err := s.buildTasks(blocks, transactions)
		if err != nil {
			return fmt.Errorf("build tasks: %w", err)
		}

		// TODO It might be possible to use generics to avoid manual type assertions.
		tasksChan <- tasks

		// Update local block number.
		s.state.BlockNumber++
	}
}

func (s *source) batchPullBlocksByRange(ctx context.Context, blockHeightStart, blockHeightEnd uint64) ([]*types.Block, error) {
	zap.L().Info("begin to batch pull transactions by range", zap.Uint64("block.height.start", blockHeightStart), zap.Uint64("block.height.end", blockHeightEnd))

	// Pull blocks by block heights
	blockHeights := lo.Map(lo.RangeWithSteps(blockHeightStart, blockHeightEnd+1, 1), func(blockHeight uint64, _ int) *big.Int {
		return new(big.Int).SetUint64(blockHeight)
	})

	blocks, err := s.batchPullBlocks(ctx, blockHeights)
	if err != nil {
		return nil, fmt.Errorf("batch pull blocks: %w", err)
	}

	return blocks, nil
}

func (s *source) batchPullBlocks(ctx context.Context, blockHeights []*big.Int) ([]*types.Block, error) {
	zap.L().Info("begin to pull blocks", zap.Int("blocks", len(blockHeights)))

	resultPool := pool.NewWithResults[*types.Block]().
		WithContext(ctx).
		WithCancelOnError()

	for _, blockHeight := range blockHeights {
		blockHeight := blockHeight

		resultPool.Go(func(ctx context.Context) (*types.Block, error) {
			return s.arweaveClient.GetBlockByHeight(ctx, blockHeight.Int64())
		})
	}

	return resultPool.Wait()
}

func (s *source) batchPullTransactions(ctx context.Context, transactionIDs []string) ([]*types.Transaction, error) {
	zap.L().Info("begin to pull transactions", zap.Int("transactions", len(transactionIDs)))

	resultPool := pool.NewWithResults[*types.Transaction]().
		WithContext(ctx).
		WithCancelOnError()

	for _, transactionID := range transactionIDs {
		transactionID := transactionID

		resultPool.Go(func(ctx context.Context) (*types.Transaction, error) {
			return s.arweaveClient.GetTransactionByID(ctx, transactionID)
		})
	}

	return resultPool.Wait()
}

func (s *source) batchPullData(ctx context.Context, transactions []*types.Transaction) error {
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

func (s *source) batchPullBundleTransactions(ctx context.Context, transactionIDs []string) ([]*types.Transaction, error) {
	zap.L().Info("begin to pull and filter bundle transactions", zap.Int("transactions", len(transactionIDs)))

	resultPool := pool.NewWithResults[[]*types.Transaction]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	for _, transactionID := range transactionIDs {
		transactionID := transactionID

		resultPool.Go(func(ctx context.Context) ([]*types.Transaction, error) {
			bundleTransactions := make([]*types.Transaction, 0)

			response, err := s.arweaveClient.GetTransactionData(ctx, transactionID)
			if err != nil {
				return nil, fmt.Errorf("fetch transaction: %w", err)
			}

			defer lo.Try(response.Close)

			decoder := bundle.NewDecoder(response)

			header, err := decoder.DecodeHeader()
			if err != nil {
				// Ignore invalid bundle transaction
				zap.L().Error("discard a invalid bundle transaction", zap.String("transaction_id", transactionID))

				return nil, nil
			}

			for index := 0; decoder.Next(); index++ {
				dataItemInfo := header.DataItemInfos[index]

				dataItem, err := decoder.DecodeDataItem()
				if err != nil {
					// Ignore invalid signature and data length
					zap.L().Error("decode data item", zap.Error(err), zap.String("transaction_id", transactionID))

					return nil, nil
				}

				bundleTransaction := types.Transaction{
					Format: 2,
					ID:     dataItemInfo.ID,
					Owner:  dataItem.Owner,
					Tags: lo.Map(dataItem.Tags, func(tag bundle.Tag, _ int) types.Tag {
						return types.Tag{
							Name:  utils.Base64Encode(tag.Name),
							Value: utils.Base64Encode(tag.Value),
						}
					}),
					Target:    dataItem.Target,
					Signature: dataItem.Signature,
				}

				data, err := io.ReadAll(dataItem)
				if err != nil {
					return nil, fmt.Errorf("read data item %s: %w", dataItemInfo.ID, err)
				}

				bundleTransaction.Data = utils.Base64Encode(data)
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

func (s *source) GroupBundleTransactions(transactions []*types.Transaction, block *types.Block) []string {
	return lo.FilterMap(transactions, func(transaction *types.Transaction, _ int) (string, bool) {
		hasBundleFormatTag := lo.ContainsBy(transaction.Tags, func(tag types.Tag) bool {
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

		hasBundleVersionTag := lo.ContainsBy(transaction.Tags, func(tag types.Tag) bool {
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

		owner, err := utils.OwnerToAddress(transaction.Owner)
		if err != nil {
			zap.L().Error("invalid owner of transaction", zap.String("transaction_id", transaction.ID), zap.Error(err))

			return "", false
		}

		bundlrNodes := []string{
			"OXcT1sVRSA5eGwt2k6Yuz8-3e3g9WJi5uSE99CWqsBs", // Bundlr Node 1
			"ZE0N-8P9gXkhtK-07PQu9d8me5tGDxa_i4Mee5RzVYg", // Bundlr Node 2
		}

		return transaction.ID, lo.Contains(bundlrNodes, owner)
	})
}

func (s *source) buildTasks(blocks []*types.Block, transactions []*types.Transaction) ([]engine.Task, error) {
	tasks := make([]engine.Task, 0)

	for _, transaction := range transactions {
		block, _ := lo.Find(blocks, func(block *types.Block) bool {
			return lo.Contains(block.Txs, transaction.ID)
		})

		tasks = append(tasks, Task{
			Chain:       filter.ChainArweave(s.state.ChainID),
			Block:       *block,
			Transaction: *transaction,
		})
	}

	return tasks, nil
}

func NewSource(config *engine.Config) (engine.Source, error) {
	instance := source{
		config: config,
	}

	return &instance, nil
}
