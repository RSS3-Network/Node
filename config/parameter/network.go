package parameter

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
)

// NetworkParamsData contains the network parameters
type NetworkParamsData struct {
	NetworkTolerance                   NetworkTolerance                   `json:"NetworkTolerance"`
	NetworkStartBlock                  NetworkStartBlock                  `json:"NetworkStartBlock"`
	NetworkCoreWorkerDiskSpacePerMonth NetworkCoreWorkerDiskSpacePerMonth `json:"NetworkCoreWorkerDiskSpacePerMonth"`
}

// PullNetworkParamsFromVSL pulls the network parameters from the VSL
func PullNetworkParamsFromVSL(networkParams *vsl.NetworkParamsCaller) error {
	params, err := networkParams.GetParams(&bind.CallOpts{})
	if err != nil {
		return err
	}

	// unmarshal params to update CurrentNetworkTolerance, CurrentNetworkStartBlock and CurrentNetworkCoreWorkerDiskSpacePerMonth
	var paramsData NetworkParamsData

	err = json.Unmarshal([]byte(params), &paramsData)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	// CurrentNetworkTolerance should be updated each epoch from vsl
	CurrentNetworkTolerance = paramsData.NetworkTolerance

	// CurrentNetworkStartBlock should be updated each epoch from vsl
	CurrentNetworkStartBlock = paramsData.NetworkStartBlock

	// CurrentNetworkCoreWorkerDiskSpacePerMonth the disk space required for the network's core worker to store a month worth of data
	// The data is calculated based on the average disk space usage during 2024 Q1.
	// Actually usage may vary depending on the network's activity.
	CurrentNetworkCoreWorkerDiskSpacePerMonth = paramsData.NetworkCoreWorkerDiskSpacePerMonth

	return nil
}

// GetCurrentEpochFromVSL Get the current epoch from VSL blockchain
func GetCurrentEpochFromVSL(settlement *vsl.SettlementCaller) (int64, error) {
	epoch, err := settlement.CurrentEpoch(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return epoch.Int64(), nil
}

// GetCurrentEpochFromCache Get the current epoch from redis cache
func GetCurrentEpochFromCache(ctx context.Context, redisClient rueidis.Client) int64 {
	if redisClient == nil {
		return 0
	}

	command := redisClient.B().Get().Key(buildCurrentEpochCacheKey()).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return 0
	}

	// Retrieve the result as a string
	epochStr, err := result.ToString()
	if err != nil {
		return 0
	}

	// Convert the string to an int64
	epoch, err := strconv.ParseInt(epochStr, 10, 64)
	if err != nil {
		return 0
	}

	return epoch
}

// UpdateCurrentEpoch updates the current epoch in redis cache
func UpdateCurrentEpoch(ctx context.Context, redisClient rueidis.Client, epoch int64) error {
	if redisClient == nil {
		return nil
	}

	command := redisClient.B().Set().Key(buildCurrentEpochCacheKey()).Value(strconv.FormatInt(epoch, 10)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// GetCurrentNetworkBlockStartFromCache Get the current network block start from redis cache
func GetCurrentNetworkBlockStartFromCache(ctx context.Context, redisClient rueidis.Client, network string) uint64 {
	if redisClient == nil {
		return 0
	}

	command := redisClient.B().Get().Key(buildNetworkBlockStartCacheKey(network)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		fmt.Println("err", err)
		return 0
	}

	// Retrieve the result as a string
	blockStartStr, err := result.ToString()
	if err != nil {
		return 0
	}

	// Convert the string to an int64
	blockStart, err := strconv.ParseInt(blockStartStr, 10, 64)
	if err != nil {
		return 0
	}

	return uint64(blockStart)
}

// UpdateCurrentBlockStart updates the current start block in redis cache
func UpdateCurrentBlockStart(ctx context.Context, redisClient rueidis.Client, network string, blockStart int64) error {
	if redisClient == nil {
		return nil
	}

	command := redisClient.B().Set().Key(buildNetworkBlockStartCacheKey(network)).Value(strconv.FormatInt(blockStart, 10)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// buildWorkerIDStatusCacheKey builds the cache key for the epoch
func buildCurrentEpochCacheKey() string {
	return "vsl:settlement:epoch"
}

// buildNetworkBlockStartCacheKey builds the cache key for the network block start
func buildNetworkBlockStartCacheKey(network string) string {
	return fmt.Sprintf("source:network:start:%s", strings.ToLower(network))
}

// InitVSLClient initializes the VSL client
func InitVSLClient() (ethereum.Client, error) {
	// TODO should use vsl rpc url in prod
	// vslEndpoint := endpoint.MustGet(network.VSL)
	vslEndpoint := "https://rpc.testnet.rss3.io"

	// Initialize vsl ethereum client.
	vslClient, err := ethereum.Dial(context.Background(), vslEndpoint)
	if err != nil {
		return nil, err
	}

	return vslClient, nil
}

func CheckParamsTask(ctx context.Context, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	localEpoch := GetCurrentEpochFromCache(ctx, redisClient)

	remoteEpoch, err := GetCurrentEpochFromVSL(settlementCaller)
	if err != nil {
		return fmt.Errorf("failed to get epoch from vsl: %w", err)
	}

	if remoteEpoch > localEpoch {
		err = PullNetworkParamsFromVSL(networkParamsCaller)
		if err != nil {
			return fmt.Errorf("failed to pull network params from vsl: %w", err)
		}

		for network, blockStart := range CurrentNetworkStartBlock {
			if blockStart == nil {
				continue // Skip if the start block is not defined.
			}

			// Convert big.Int to int64; safe as long as the value fits in int64.
			blockStartInt64 := blockStart.Int64()

			// Update the current block start for the network in Redis.
			err := UpdateCurrentBlockStart(ctx, redisClient, network, blockStartInt64)
			if err != nil {
				return fmt.Errorf("update current block start: %w", err)
			}
		}

		err = UpdateCurrentEpoch(ctx, redisClient, remoteEpoch)
		if err != nil {
			return fmt.Errorf("failed to update epoch in cache: %w", err)
		}
	}

	return nil
}
