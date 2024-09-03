package engine

import (
	"sync"
)

// TaskBuffer represents a fixed-size buffer, it operates as a sliding-window buffer(FIFO)
type TaskBuffer struct {
	tasks    []*Tasks
	size     int
	mu       sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

// NewTaskBuffer creates a new TaskBuffer buffer with the specified size
func NewTaskBuffer(size int) *TaskBuffer {
	sw := &TaskBuffer{
		tasks: make([]*Tasks, 0, size),
		size:  size,
	}
	sw.notEmpty = sync.NewCond(&sw.mu)
	sw.notFull = sync.NewCond(&sw.mu)

	return sw
}

// Add adds a new task to the sliding window, wait if the buffer is full
func (sw *TaskBuffer) Add(task *Tasks) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	for len(sw.tasks) >= sw.size {
		sw.notFull.Wait()
	}

	sw.tasks = append(sw.tasks, task)
	sw.notEmpty.Signal()
}

// Get retrieves the oldest task from the sliding window and removes it
func (sw *TaskBuffer) Get() *Tasks {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	for len(sw.tasks) == 0 {
		sw.notEmpty.Wait()
	}

	task := sw.tasks[0]
	sw.tasks = sw.tasks[1:]
	sw.notFull.Signal()

	return task
}
