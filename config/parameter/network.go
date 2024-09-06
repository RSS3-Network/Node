package parameter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/protocol-go/schema/network"
)

// NetworkParamsData contains the network parameters
type NetworkParamsData struct {
	NetworkTolerance                   NetworkTolerance                   `json:"network_tolerance"`
	NetworkStartBlock                  NetworkStartBlock                  `json:"network_start_block"`
	NetworkCoreWorkerDiskSpacePerMonth NetworkCoreWorkerDiskSpacePerMonth `json:"network_core_worker_disk_space_per_month"`
}

// UnmarshalJSON defines a custom UnmarshalJSON method for NetworkParamsData
func (npd *NetworkParamsData) UnmarshalJSON(data []byte) error {
	networkParam := struct {
		NetworkTolerance                   map[string]uint64      `json:"network_tolerance"`
		NetworkStartBlock                  map[string]*StartBlock `json:"network_start_block"`
		NetworkCoreWorkerDiskSpacePerMonth map[string]uint        `json:"network_core_worker_disk_space_per_month"`
	}{}

	if err := json.Unmarshal(data, &networkParam); err != nil {
		return err
	}

	npd.NetworkTolerance = make(map[network.Network]uint64)
	npd.NetworkStartBlock = make(map[network.Network]*StartBlock)
	npd.NetworkCoreWorkerDiskSpacePerMonth = make(map[network.Network]uint)

	for key, value := range networkParam.NetworkTolerance {
		netValue, err := network.NetworkString(key)
		if err != nil {
			return err
		}

		npd.NetworkTolerance[netValue] = value
	}

	for key, value := range networkParam.NetworkStartBlock {
		netValue, err := network.NetworkString(key)
		if err != nil {
			return err
		}

		npd.NetworkStartBlock[netValue] = value
	}

	for key, value := range networkParam.NetworkCoreWorkerDiskSpacePerMonth {
		netValue, err := network.NetworkString(key)
		if err != nil {
			return err
		}

		npd.NetworkCoreWorkerDiskSpacePerMonth[netValue] = value
	}

	return nil
}

// PullNetworkParamsFromVSL pulls the network parameters from the VSL
func PullNetworkParamsFromVSL(networkParams *vsl.NetworkParamsCaller, epoch uint64) (string, error) {
	// Get parameters for the current epoch from networkParams
	params, err := networkParams.GetParams(&bind.CallOpts{}, epoch)
	if err != nil {
		return "", fmt.Errorf("failed to get params for epoch %d: %w", epoch, err)
	}

	if err := updateNetworkParams(params); err != nil {
		return "", fmt.Errorf("update network params: %w", err)
	}

	// Unmarshal and update network parameters
	return params, nil
}

func updateNetworkParams(params string) error {
	var paramsData NetworkParamsData

	err := json.Unmarshal([]byte(params), &paramsData)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	// Update network parameters
	CurrentNetworkTolerance = paramsData.NetworkTolerance
	CurrentNetworkStartBlock = paramsData.NetworkStartBlock
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

// GetCurrentEpoch Get the current epoch from redis cache
func GetCurrentEpoch(ctx context.Context, redisClient rueidis.Client) (int64, error) {
	if redisClient == nil {
		return 0, fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Get().Key(buildCurrentEpochCacheKey()).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		if errors.Is(err, rueidis.Nil) {
			// Key doesn't exist, return 0 or a default value
			return 0, nil
		}

		return 0, fmt.Errorf("redis result: %w", err)
	}

	// Retrieve the result as a string
	epochStr, err := result.ToString()
	if err != nil {
		return 0, fmt.Errorf("redis result to string: %w", err)
	}

	// Convert the string to an int64
	epoch, err := strconv.ParseInt(epochStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
	}

	return epoch, nil
}

// UpdateCurrentEpoch updates the current epoch in redis cache
func UpdateCurrentEpoch(ctx context.Context, redisClient rueidis.Client, epoch int64) error {
	if redisClient == nil {
		return fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Set().Key(buildCurrentEpochCacheKey()).Value(strconv.FormatInt(epoch, 10)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		if errors.Is(err, rueidis.Nil) {
			// Key doesn't exist, return 0 or a default value
			return nil
		}

		return fmt.Errorf("redis result: %w", err)
	}

	return nil
}

// GetNetworkBlockStart Get the current network block start from redis cache
func GetNetworkBlockStart(ctx context.Context, redisClient rueidis.Client, network string) (uint64, error) {
	if redisClient == nil {
		return 0, fmt.Errorf("redis client is nil")
	}

	command := redisClient.B().Get().Key(buildNetworkBlockStartCacheKey(network)).Build()

	result := redisClient.Do(ctx, command)
	if err := result.Error(); err != nil {
		if errors.Is(err, rueidis.Nil) {
			// Key doesn't exist, return 0 or a default value
			return 0, nil
		}

		return 0, fmt.Errorf("redis result: %w", err)
	}

	// Retrieve the result as a string
	blockStartStr, err := result.ToString()
	if err != nil {
		return 0, fmt.Errorf("redis result to string: %w", err)
	}

	// Convert the string to an int64
	blockStart, err := strconv.ParseInt(blockStartStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int: %w", err)
	}

	return uint64(blockStart), nil
}

// UpdateBlockStart updates the current start block in redis cache
func UpdateBlockStart(ctx context.Context, redisClient rueidis.Client, network string, blockStart int64) error {
	if redisClient == nil {
		return fmt.Errorf("redis client is nil")
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
	// Initialize vsl ethereum client.
	vslClient, err := ethereum.Dial(context.Background(), endpoint.MustGet(network.VSL))
	if err != nil {
		return nil, err
	}

	return vslClient, nil
}

func CheckParamsTask(ctx context.Context, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller) error {
	currentEpoch, err := GetCurrentEpoch(ctx, redisClient)
	if err != nil {
		return fmt.Errorf("failed to get current epoch from cache: %w", err)
	}

	_, err = PullNetworkParamsFromVSL(networkParamsCaller, uint64(currentEpoch))
	if err != nil {
		return fmt.Errorf("failed to pull network params from vsl: %w", err)
	}

	for n, blockStart := range CurrentNetworkStartBlock {
		if blockStart == nil {
			continue // Skip if the start block is not defined.
		}

		// Update the current block start for the network in Redis.
		err := UpdateBlockStart(ctx, redisClient, n.String(), blockStart.Block.Int64())
		if err != nil {
			return fmt.Errorf("update current block start: %w", err)
		}
	}

	return nil
}
