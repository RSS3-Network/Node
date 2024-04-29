package pool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/protocol-go/schema/network"
)

type ContractType string

// Contract types.
const (
	ContractTypePool  ContractType = "pool"
	ContractTypeToken ContractType = "token"
	ContractTypeGauge ContractType = "gauge"
)

// Response is a curve pool.
type Response[T any] struct {
	Success   bool  `json:"success"`
	Data      T     `json:"data"`
	Timestamp int64 `json:"generatedTimeMs"`
}

type GetPoolData struct {
	PoolData []Pool `json:"poolData"`
}

// Pool is a curve pool.
type Pool struct {
	Network                       network.Network `json:"-"`
	Name                          string          `json:"name"`
	Address                       common.Address  `json:"address"`
	LiquidityProviderTokenAddress common.Address  `json:"lpTokenAddress"`
	GaugeAddress                  common.Address  `json:"gaugeAddress"`
}
