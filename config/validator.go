package config

import (
	"fmt"

	"github.com/samber/lo"
)

// HasOneWorker Check if there is at least one worker deployed
func HasOneWorker(config *File) error {
	if (CalculateWorkerCount(config)) == 0 {
		return fmt.Errorf("at least one worker must be deployed")
	}

	return nil
}

// IsRSSOrAIComponentOnly Check if the configuration contains an RSS or AI component only
func IsRSSOrAIComponentOnly(config *File) bool {
	return len(config.Component.Decentralized) == 0 && len(config.Component.Federated) == 0 &&
		(config.Component.RSS != nil || config.Component.AI != nil)
}

// CalculateWorkerCount returns the number of workers deployed
func CalculateWorkerCount(config *File) int {
	return len(config.Component.Decentralized) + lo.Ternary(config.Component.RSS != nil, 1, 0) + len(config.Component.Federated) + lo.Ternary(config.Component.AI != nil, 1, 0)
}
