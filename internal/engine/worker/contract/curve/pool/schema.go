package pool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/protocol-go/schema/filter"
)

type ContractType string

const (
	ContractTypePool  ContractType = "pool"
	ContractTypeToken ContractType = "token"
	ContractTypeGauge ContractType = "gauge"
)

type Response[T any] struct {
	Success   bool  `json:"success"`
	Data      T     `json:"data"`
	Timestamp int64 `json:"generatedTimeMs"`
}

type GetPoolData struct {
	PoolData []Pool `json:"poolData"`
}

type Pool struct {
	Network                       filter.Network `json:"-"`
	Name                          string         `json:"name"`
	Address                       common.Address `json:"address"`
	LiquidityProviderTokenAddress common.Address `json:"lpTokenAddress"`
	GaugeAddress                  common.Address `json:"gaugeAddress"`
}
