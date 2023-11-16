package sync

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrorNoResult = errors.New("no result")

type QuickGroup[R any] interface {
	Go(f func(ctx context.Context) (R, error))
	Wait() (R, error)
}

var _ QuickGroup[any] = (*quickGroup[any])(nil)

type quickGroup[R any] struct {
	waitGroup sync.WaitGroup
	ctx       context.Context
	cancels   []context.CancelFunc
	result    R
	err       error
	done      atomic.Bool
	locker    sync.Mutex
}

func (q *quickGroup[R]) Go(f func(ctx context.Context) (R, error)) {
	if q.done.Load() {
		return
	}

	// Ensure that the cancel slice is thread-safe.
	q.locker.Lock()

	q.waitGroup.Add(1)

	// Create a context to later cancel slow tasks.
	ctx, cancel := context.WithCancel(q.ctx)
	q.cancels = append(q.cancels, cancel)
	index := len(q.cancels) - 1

	q.locker.Unlock()

	go func() {
		defer q.waitGroup.Done()

		if result, err := f(ctx); err == nil {
			if !q.done.Swap(true) { // Here is equivalent to sync.Once.
				q.result = result
				q.err = nil

				// Cancel all other pending tasks.
				for i, cancel := range q.cancels {
					if i == index {
						continue
					}

					cancel()
				}
			}
		}
	}()
}

func (q *quickGroup[R]) Wait() (result R, err error) {
	q.waitGroup.Wait()

	return q.result, q.err
}

func NewQuickGroup[R any](ctx context.Context) QuickGroup[R] {
	instance := &quickGroup[R]{
		ctx: ctx,
		err: ErrorNoResult, // Default error.
	}

	return instance
}
