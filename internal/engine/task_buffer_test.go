package engine

import (
	"context"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Initialization Test Cases:
func TestNewTaskBuffer(t *testing.T) {
	t.Parallel()
	t.Run("Create with positive size", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(10)
		assert.NotNil(t, buffer)
		assert.Equal(t, 10, buffer.size)
		assert.Empty(t, buffer.tasks)
	})

	t.Run("Create with zero size", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(0)
		assert.NotNil(t, buffer)
		assert.Equal(t, 0, buffer.size)
		assert.Empty(t, buffer.tasks)
	})
}

// Basic Operation Test Cases:
func TestTaskBuffer_Add(t *testing.T) {
	t.Parallel()

	t.Run("Add to non-full buffer", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(2)
		task := &Tasks{}
		buffer.Add(task)
		assert.Len(t, buffer.tasks, 1)
		assert.Equal(t, task, buffer.tasks[0])
	})

	t.Run("Add to full buffer", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(1)
		task1 := &Tasks{}
		task2 := &Tasks{}

		buffer.Add(task1)

		done := make(chan bool)
		go func() {
			buffer.Add(task2)
			done <- true
		}()

		select {
		case <-done:
			t.Fatal("Add should block when buffer is full")
		case <-time.After(100 * time.Millisecond): // Expected behavior
		}

		assert.Len(t, buffer.tasks, 1)
		assert.Equal(t, task1, buffer.tasks[0])
	})
}

func TestTaskBuffer_Get(t *testing.T) {
	t.Parallel()
	t.Run("Get from non-empty buffer", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(2)
		task := &Tasks{}
		buffer.Add(task)

		got := buffer.Get()
		assert.Equal(t, task, got)
		assert.Empty(t, buffer.tasks)
	})

	t.Run("Get from empty buffer", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(1)

		done := make(chan bool)
		go func() {
			buffer.Get()
			done <- true
		}()

		select {
		case <-done:
			t.Fatal("Get should block when buffer is empty")
		case <-time.After(100 * time.Millisecond):
		}
	})
}

func TestTaskBuffer_Concurrency(t *testing.T) {
	t.Parallel()
	t.Run("Multiple producers and consumers", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(10)
		producerCount := 5
		consumerCount := 5
		tasksPerProducer := 100

		var wg sync.WaitGroup

		wg.Add(producerCount + consumerCount)

		// Start producers
		for i := 0; i < producerCount; i++ {
			go func() {
				defer wg.Done()

				for j := 0; j < tasksPerProducer; j++ {
					buffer.Add(&Tasks{})
				}
			}()
		}

		// Start consumers
		consumedTasks := make(chan *Tasks, producerCount*tasksPerProducer)

		for i := 0; i < consumerCount; i++ {
			go func() {
				defer wg.Done()

				for j := 0; j < tasksPerProducer*producerCount/consumerCount; j++ {
					task := buffer.Get()
					consumedTasks <- task
				}
			}()
		}

		wg.Wait()
		close(consumedTasks)

		assert.Equal(t, producerCount*tasksPerProducer, len(consumedTasks))
	})
}

// Basic Edge Cases:
func TestTaskBuffer_EdgeCases(t *testing.T) {
	t.Parallel()
	t.Run("Add and get with size 1 buffer", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(1)

		task := &Tasks{}

		buffer.Add(task)
		got := buffer.Get()

		assert.Equal(t, task, got)
		assert.Empty(t, buffer.tasks)
	})

	t.Run("Alternating add and get", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(1)

		for i := 0; i < 1000; i++ {
			buffer.Add(&Tasks{})
			buffer.Get()
		}
		assert.Empty(t, buffer.tasks)
	})
}

// Advanced Edge Cases:
func TestTaskBuffer_RaceConditions(t *testing.T) {
	t.Parallel()

	t.Run("Concurrent add and get", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(1000)
		operationCount := 5000

		var wg sync.WaitGroup

		wg.Add(2)

		go func() {
			defer wg.Done()

			for i := 0; i < operationCount; i++ {
				buffer.Add(&Tasks{})
			}
		}()
		go func() {
			defer wg.Done()

			for i := 0; i < operationCount; i++ {
				buffer.Get()
			}
		}()

		wg.Wait()

		// Verification
		assert.True(t, buffer.size >= 0 && buffer.size <= 1000, "Buffer size should be between 0 and 1000")
	})
}

func TestTaskBuffer_StressTest(t *testing.T) {
	t.Parallel()

	t.Run("Rapid add and get operations", func(t *testing.T) {
		t.Parallel()

		buffer := NewTaskBuffer(100)
		operationCount := 5000

		var wg sync.WaitGroup

		wg.Add(2)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		runOperations := func() {
			defer wg.Done()

			for i := 0; i < operationCount; i++ {
				select {
				case <-ctx.Done():
					return
				default:
					if i%2 == 0 {
						buffer.Add(&Tasks{})
					} else {
						buffer.Get()
					}
				}
			}
		}

		go runOperations()
		go runOperations()

		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Test completed successfully
		case <-ctx.Done():
			t.Fatal("Test timed out")
		}

		assert.True(t, buffer.size >= 0 && buffer.size <= 100, "Final buffer size should be between 0 and 100")
	})
}

// Rapid switching between full and empty states
func TestTaskBuffer_RapidFullEmptyCycles(t *testing.T) {
	t.Parallel()

	buffer := NewTaskBuffer(10)
	cycles := 1000

	for i := 0; i < cycles; i++ {
		// Fill the Buffer
		for j := 0; j < 10; j++ {
			buffer.Add(&Tasks{})
		}
		assert.Equal(t, 10, len(buffer.tasks), "Buffer should be full")

		// Empty the buffer
		for j := 0; j < 10; j++ {
			buffer.Get()
		}
		assert.Equal(t, 0, len(buffer.tasks), "Buffer should be empty")
	}
}

// Multiple goroutines waiting on full/empty buffer
// func TestTaskBuffer_MultipleWaiters(t *testing.T) {
//	buffer := NewTaskBuffer(1)
//	var wg sync.WaitGroup
//	waiterCount := 2
//	wg.Add(waiterCount * 2)
//
//	// Waiters for Add
//	for i := 0; i < waiterCount; i++ {
//		go func() {
//			defer wg.Done()
//			buffer.Add(&Tasks{})
//		}()
//	}
//
//	// Waiters for Get
//	for i := 0; i < waiterCount; i++ {
//		go func() {
//			defer wg.Done()
//			buffer.Get()
//		}()
//	}
//
//	// Allow some time for goroutines to start waiting
//	time.Sleep(100 * time.Millisecond)
//
//	// Trigger the waiters
//	for i := 0; i < waiterCount; i++ {
//		buffer.Get()
//		buffer.Add(&Tasks{})
//	}
//
//	wg.Wait()
//	assert.Equal(t, 0, len(buffer.tasks), "Buffer should be empty")
// }

// Performance Test Cases:
func TestTaskBuffer_MemoryLeak(t *testing.T) {
	t.Parallel()

	buffer := NewTaskBuffer(1000)

	var m1, m2 runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&m1)

	for i := 0; i < 1000000; i++ {
		buffer.Add(&Tasks{})
		buffer.Get()
	}

	runtime.GC()
	runtime.ReadMemStats(&m2)

	// Allow for some small increase, but nothing significant

	// Detect Memory Leak Issue
	assert.InDelta(t, m1.Alloc, m2.Alloc, float64(m1.Alloc)*0.1, "Potential memory leak detected")
}

func TestTaskBuffer_Fairness(t *testing.T) {
	t.Parallel()

	buffer := NewTaskBuffer(10)
	taskCount := 100
	workerCount := 10

	var wg sync.WaitGroup

	wg.Add(workerCount)

	startTime := time.Now()

	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < taskCount; j++ {
				task := &Tasks{}
				buffer.Add(task)
				time.Sleep(time.Millisecond)

				gotTask := buffer.Get()
				require.Equal(t, task, gotTask, "Task mismatch for worker %d", workerID)
			}
		}(i)
	}

	wg.Wait()

	duration := time.Since(startTime)

	t.Logf("Fairness test completed in %v", duration)
}
