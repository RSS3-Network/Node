package ethereum

import "github.com/ethereum/go-ethereum/common"

type Filter struct {
	LogAddresses []common.Address `yaml:"log_addresses"`
	LogTopics    []common.Hash    `yaml:"log_topics"`
}
