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

// IsComponentOnly Check if the configuration contains an RSS or AI component only
func IsComponentOnly(config *File) bool {
	return len(config.Component.Decentralized) == 0 && len(config.Component.Federated) == 0 &&
		(config.Component.RSS != nil || config.Component.AI != nil)
}

// CalculateWorkerCount returns the number of workers deployed
func CalculateWorkerCount(config *File) int {
	return len(config.Component.Decentralized) + lo.Ternary(config.Component.RSS != nil, 1, 0) + len(config.Component.Federated)
}

// CalculateComponentCount returns the number of components deployed
func CalculateComponentCount(config *File) int {
	return lo.Ternary(len(config.Component.Decentralized) > 0, 1, 0) + lo.Ternary(config.Component.RSS != nil, 1, 0) + lo.Ternary(len(config.Component.Federated) > 0, 1, 0)
}
