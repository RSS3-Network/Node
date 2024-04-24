package sync

import (
	"context"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestQuickGroup(t *testing.T) {
	t.Parallel()

	quickGroup := NewQuickGroup[time.Duration](context.Background())

	for duration := time.Second; duration > 0; duration -= 100 * time.Millisecond {
		duration := duration

		task := func(ctx context.Context) (time.Duration, error) {
			timer := time.NewTimer(duration)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				t.Logf("Task %s: %s", duration, context.Canceled)

				return duration, context.Canceled
			case <-timer.C:
				t.Logf("Task %s: %s", duration, "done")

				return duration, nil
			}
		}

		quickGroup.Go(task)
	}

	result, err := quickGroup.Wait()
	require.NoError(t, err)
	require.Equal(t, result, 100*time.Millisecond)
}

// TestQuickGroupDataRace checks if any data race exists for QuickGroup.
func TestQuickGroupDataRace(t *testing.T) {
	t.Parallel()

	quickGroup := NewQuickGroup[bool](context.Background())

	for range lo.Range(10000) {
		task := func(ctx context.Context) (bool, error) {
			timer := time.NewTimer(time.Second)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				return false, context.Canceled
			case <-timer.C:
				return true, nil
			}
		}

		quickGroup.Go(task)
	}

	result, err := quickGroup.Wait()
	require.NoError(t, err)
	require.Equal(t, result, true)
}
