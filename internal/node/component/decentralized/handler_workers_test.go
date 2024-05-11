package decentralized

import (
	"testing"

	"github.com/rss3-network/node/schema/worker"
	"github.com/stretchr/testify/assert"
)

func TestDetermineFinalStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		statuses []worker.Status
		expected worker.Status
	}{
		{
			name:     "AllStatusesReady",
			statuses: []worker.Status{worker.StatusReady, worker.StatusReady, worker.StatusReady},
			expected: worker.StatusReady,
		},
		{
			name:     "AtLeastOneStatusIndexing",
			statuses: []worker.Status{worker.StatusReady, worker.StatusIndexing, worker.StatusReady},
			expected: worker.StatusIndexing,
		},
		{
			name:     "AtLeastOneStatusUnhealthy",
			statuses: []worker.Status{worker.StatusReady, worker.StatusUnhealthy, worker.StatusReady},
			expected: worker.StatusUnhealthy,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := determineFinalStatus(tc.statuses)
			assert.Equal(t, tc.expected, result)
		})
	}
}
