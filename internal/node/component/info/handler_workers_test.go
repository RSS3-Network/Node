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
			name:     "Ready",
			statuses: []worker.Status{worker.StatusReady},
			expected: worker.StatusReady,
		},
		{
			name:     "Indexing",
			statuses: []worker.Status{worker.StatusIndexing},
			expected: worker.StatusIndexing,
		},
		{
			name:     "Unhealthy",
			statuses: []worker.Status{worker.StatusUnhealthy},
			expected: worker.StatusUnhealthy,
		},
		{
			name:     "2 Instance, Ready + Indexing",
			statuses: []worker.Status{worker.StatusReady, worker.StatusIndexing},
			expected: worker.StatusUnhealthy,
		},
		{
			name:     "2 Instance, Indexing + Indexing",
			statuses: []worker.Status{worker.StatusIndexing, worker.StatusIndexing},
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
