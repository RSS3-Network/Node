package pool

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
)

// Registry is a registry for the curve pools.
type Registry interface {
	Refresh(ctx context.Context) error
	Validate(ctx context.Context, chain network.Network, contractType ContractType, address common.Address) (*Pool, error)
}

var _ Registry = (*registry)(nil)

// client is a client for the curve registry.
type registry struct {
	redisClient rueidis.Client
	httpClient  httpx.Client
	mu          sync.Mutex
}

// Refresh fetches the pools from the curve endpoint and stores them in the redis.
func (r *registry) Refresh(ctx context.Context) error {
	r.mu.Lock()         // Lock the mutex
	defer r.mu.Unlock() // Unlock the mutex when the function returns

	const Endpoint = "https://api.curve.fi/"

	// Networks to fetch pools for.
	networks := []network.Network{
		network.Arbitrum,
		network.Avalanche,
		network.Ethereum,
		network.Gnosis,
		network.Optimism,
		network.Polygon,
	}

	// Registry IDs to fetch pools for.
	registryIDs := []string{
		"main",
		"crypto",
		"factory",
		"factory-crypto",
	}

	// Fetch pools from the curve endpoint.
	resultPool := pool.NewWithResults[[]Pool]().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	// iterate over networks and registry IDs.
	for _, network := range networks {
		network := network

		// iterate over registry IDs.
		for _, registryID := range registryIDs {
			registryID := registryID

			resultPool.Go(func(ctx context.Context) ([]Pool, error) {
				readCloser, err := r.httpClient.Fetch(ctx, fmt.Sprintf("%sapi/getPools/%s/%s", Endpoint, r.formatNetwork(network), registryID))
				if err != nil {
					return nil, fmt.Errorf("fetch request: %w", err)
				}
				defer lo.Try(readCloser.Close)

				var result Response[GetPoolData]

				if err := json.NewDecoder(readCloser).Decode(&result); err != nil {
					return nil, fmt.Errorf("decode json: %w", err)
				}

				for index, curvePool := range result.Data.PoolData {
					curvePool.Network = network

					result.Data.PoolData[index] = curvePool
				}

				return result.Data.PoolData, nil
			})
		}
	}

	result, err := resultPool.Wait()
	if err != nil {
		return fmt.Errorf("wait: %w", err)
	}

	curvePools := lo.Flatten(result)

	commands := make([]rueidis.Completed, 0, len(curvePools))

	// Store the pools in the redis.
	for _, curvePool := range curvePools {
		keys := []string{
			r.formatRedisKey(curvePool.Network, ContractTypePool, curvePool.Address),
			r.formatRedisKey(curvePool.Network, ContractTypeToken, curvePool.LiquidityProviderTokenAddress),
			r.formatRedisKey(curvePool.Network, ContractTypeGauge, curvePool.GaugeAddress),
		}

		for _, key := range keys {
			command := r.redisClient.B().Set().Key(key).Value(curvePool.Name).Build()

			commands = append(commands, command)
		}
	}

	// Execute the redis commands.
	for _, redisResult := range r.redisClient.DoMulti(ctx, commands...) {
		if err := redisResult.Error(); err != nil {
			return fmt.Errorf("redis result: %w", err)
		}
	}

	return nil
}

// Validate checks if the pool is valid.
func (r *registry) Validate(ctx context.Context, network network.Network, contractType ContractType, address common.Address) (*Pool, error) {
	command := r.redisClient.B().Get().Key(r.formatRedisKey(network, contractType, address)).Build()

	result := r.redisClient.Do(ctx, command)

	if err := result.Error(); err != nil {
		return nil, fmt.Errorf("redis result: %w", err)
	}

	curvePool := Pool{
		Network: network,
		Address: address,
	}

	var err error
	if curvePool.Name, err = result.ToString(); err != nil {
		return nil, fmt.Errorf("to string: %w", err)
	}

	return &curvePool, nil
}

// formatRedisKey formats the redis key.
func (r *registry) formatRedisKey(network network.Network, contractType ContractType, address common.Address) string {
	return fmt.Sprintf("curve:%s:%s:%s", r.formatNetwork(network), contractType, address)
}

// formatNetwork handling special network name in curve endpoint like xdai.
func (r *registry) formatNetwork(n network.Network) string {
	switch n {
	case network.Gnosis:
		return "xdai"
	case network.Avalanche:
		return "avalanche"
	default:
		return n.String()
	}
}

// NewRegistry creates a new registry.
func NewRegistry(redisClient rueidis.Client, httpClient httpx.Client) Registry {
	return &registry{
		redisClient: redisClient,
		httpClient:  httpClient,
	}
}
