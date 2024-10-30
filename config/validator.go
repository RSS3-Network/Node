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

// IsRSSComponentOnly Check if the configuration contains an RSS component only
func IsRSSComponentOnly(config *File) bool {
	return len(config.Component.Decentralized) == 0 && config.Component.RSS != nil && len(config.Component.Federated) == 0
}

func CalculateWorkerCount(config *File) int {
	return len(config.Component.Decentralized) + lo.Ternary(config.Component.RSS != nil, 1, 0) + len(config.Component.Federated)
}
