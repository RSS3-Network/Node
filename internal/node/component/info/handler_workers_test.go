package info

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
			name:     "AtLeastOneStatusUnhealthy(Ready)",
			statuses: []worker.Status{worker.StatusReady, worker.StatusUnhealthy, worker.StatusReady},
			expected: worker.StatusUnhealthy,
		},
		{
			name:     "AtLeastOneStatusUnhealthy(Indexing)",
			statuses: []worker.Status{worker.StatusIndexing, worker.StatusUnhealthy, worker.StatusIndexing},
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
