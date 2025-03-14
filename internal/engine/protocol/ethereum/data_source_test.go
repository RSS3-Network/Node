package ethereum_test

import (
	"context"
	"sync"
	"testing"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

var initializeOnce sync.Once

func initialize(t *testing.T) {
	initializeOnce.Do(func() {
		zap.ReplaceGlobals(zaptest.NewLogger(t))
	})
}

func TestSource(t *testing.T) {
	t.Parallel()

	initialize(t)

	type arguments struct {
		config *config.Module
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      require.ValueAssertionFunc
		wantError require.ErrorAssertionFunc
	}{
		// TODO Implement a solution to configure custom block number ranges for protocol.
		// {
		// 	name: "From block number 15537393 to 15537398",
		// 	want: func(t require.TestingT, actual interface{}, msgAndArgs ...interface{}) {
		// 		tasks, ok := actual.([]engine.Task)
		// 		require.True(t, ok, "invalid tasks type: %T", actual)
		//
		// 		require.Len(t, tasks, 5, "invalid tasks length")
		// 	},
		// },
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			instance, err := ethereum.NewSource(testcase.arguments.config, nil, nil, nil)
			require.NoError(t, err, "new ethereum dataSource")

			var (
				tasksChan = make(chan *engine.Tasks)
				errorChan = make(chan error)
			)

			instance.Start(context.Background(), tasksChan, errorChan)

			for {
				select {
				case tasks := <-tasksChan:
					for _, task := range tasks.Tasks {
						t.Logf("Task %s", task.ID())
					}
				case err := <-errorChan:
					require.NoError(t, err)

					return
				}
			}
		})
	}
}
