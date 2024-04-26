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
	locker    sync.RWMutex
}

func (q *quickGroup[R]) Go(f func(ctx context.Context) (R, error)) {
	if q.done.Load() {
		return
	}

	// Allocate a context and an id for the task.
	ctx, id := q.allocateContext()

	go func() {
		defer q.waitGroup.Done()

		if result, err := f(ctx); err == nil {
			if !q.done.Swap(true) { // Here is equivalent to sync.Once.
				q.result = result
				q.err = nil

				// Lock the locker for reading the cancels slice.
				q.locker.RLock()
				defer q.locker.RUnlock()

				// Cancel all other pending tasks.
				for i, cancel := range q.cancels {
					if i == id {
						continue
					}

					cancel()
				}
			}
		}
	}()
}

// allowContext returns a context that can be used to cancel slow tasks.
func (q *quickGroup[R]) allocateContext() (ctx context.Context, id int) {
	q.locker.Lock()
	defer q.locker.Unlock()

	ctx, cancel := context.WithCancel(q.ctx)

	q.cancels = append(q.cancels, cancel)
	q.waitGroup.Add(1)

	return ctx, len(q.cancels) - 1
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
