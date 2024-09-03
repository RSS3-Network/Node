package near_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/internal/engine/source/near"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSource(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		config         *config.Module
		sourceFilter   engine.DataSourceFilter
		checkpoint     *engine.Checkpoint
		expectedError  string
		expectedTasks  int
		expectedHeight uint64
	}{
		{
			name: "Success",
			config: &config.Module{
				Network: network.Near,
				Endpoint: config.Endpoint{
					URL: "https://archival-rpc.mainnet.near.org",
				},
				Parameters: &config.Parameters{
					"block_start": 127170738,
				},
			},
			sourceFilter:   &near.Filter{},
			checkpoint:     nil,
			expectedTasks:  129,
			expectedHeight: 127170740,
		},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			dataSource, err := near.NewSource(testCase.config, testCase.sourceFilter, testCase.checkpoint, nil)
			if testCase.expectedError != "" {
				assert.EqualError(t, err, testCase.expectedError)
				return
			}

			require.NoError(t, err)

			tasksChan := make(chan *engine.Tasks)
			errorChan := make(chan error)

			go dataSource.Start(ctx, tasksChan, errorChan)

			var tasks []engine.Task

			var lastError error

			for {
				select {
				case receivedTasks := <-tasksChan:
					tasks = append(tasks, receivedTasks.Tasks...)
					if len(tasks) >= testCase.expectedTasks {
						goto done
					}
				case err := <-errorChan:
					lastError = err
					goto done
				case <-ctx.Done():
					goto done
				}
			}

		done:
			require.NoError(t, lastError)
			assert.Len(t, tasks, testCase.expectedTasks)

			var state near.State
			err = json.Unmarshal(dataSource.State(), &state)
			require.NoError(t, err)
			assert.Equal(t, testCase.expectedHeight, state.BlockHeight)
		})
	}
}
