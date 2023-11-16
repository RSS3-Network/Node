package syncx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQuickGroup(t *testing.T) {
	t.Parallel()

	task := func(ctx context.Context, i time.Duration) (time.Duration, string) {
		ticker := time.NewTicker(i)

		select {
		case <-ctx.Done():
			return i, "cancel"
		case <-ticker.C:
			return i, "done"
		}
	}

	quickGroup := NewQuickGroup[time.Duration](context.Background())

	for i := 5; i > 0; i-- {
		duration := time.Duration(i) * time.Second

		quickGroup.Go(func(ctx context.Context) (time.Duration, error) {
			duration, status := task(ctx, duration)
			t.Log(duration, status)

			if status == "done" {
				return duration, nil
			}

			return duration, fmt.Errorf("cancelled")
		})
	}

	id, err := quickGroup.Wait()
	require.Equal(t, id, time.Duration(1)*time.Second)
	require.NoError(t, err)
}
