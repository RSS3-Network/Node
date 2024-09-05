package engine_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/rss3-network/node/internal/engine"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

var TaskBufferSize = 2000

func BenchmarkDataTransmissionPerformance(b *testing.B) {
	logger := zaptest.NewLogger(b)
	zap.ReplaceGlobals(logger)

	// Test parameters
	dataCounts := []int{1000}

	const (
		producerCount = 4
		workerCount   = 2
	)

	for _, dataCount := range dataCounts {
		b.Run(fmt.Sprintf("TaskBuffer_DataCount_%d", dataCount), func(b *testing.B) {
			benchmarkTaskBuffer(b, dataCount, producerCount, workerCount)
		})

		b.Run(fmt.Sprintf("Channel_DataCount_%d", dataCount), func(b *testing.B) {
			benchmarkChannel(b, dataCount, producerCount, workerCount)
		})
	}
}

func simulateWork() {
	n, _ := rand.Int(rand.Reader, big.NewInt(5))
	time.Sleep(time.Duration(5+n.Int64()) * time.Millisecond)
}

func benchmarkTaskBuffer(b *testing.B, dataCount, producerCount, workerCount int) {
	taskBuffer := engine.NewTaskBuffer(TaskBufferSize) // Default task buffer size: 1500

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup

		processed := make(chan *engine.Tasks, dataCount)

		startTime := time.Now()

		// Start producers (rapid data generation)
		tasksPerProducer := dataCount / producerCount
		remainingTasks := dataCount % producerCount

		for p := 0; p < producerCount; p++ {
			wg.Add(1)

			go func(producerID int) {
				defer wg.Done()

				tasksToGenerate := tasksPerProducer

				if producerID == producerCount-1 {
					tasksToGenerate += remainingTasks
				}

				for j := 0; j < tasksToGenerate; j++ {
					taskBuffer.Add(&engine.Tasks{})
				}
			}(p)
		}

		// Start workers (slower processing)
		for w := 0; w < workerCount; w++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				for {
					task := taskBuffer.Get()
					if task == nil {
						return
					}

					simulateWork()
					processed <- task
				}
			}()
		}

		// Wait for all data to be processed
		for j := 0; j < dataCount; j++ {
			<-processed
		}

		// Signal workers to stop
		for w := 0; w < workerCount; w++ {
			taskBuffer.Add(nil)
		}

		wg.Wait()

		duration := time.Since(startTime)
		b.ReportMetric(float64(duration.Milliseconds()), "ms/op")
		b.ReportMetric(float64(dataCount)/duration.Seconds(), "items/sec")
	}
}

func benchmarkChannel(b *testing.B, dataCount, producerCount, workerCount int) {
	for i := 0; i < b.N; i++ {
		taskChan := make(chan *engine.Tasks) // Unbuffered channel

		var wg sync.WaitGroup

		processed := make(chan *engine.Tasks, dataCount)

		startTime := time.Now()

		// Start producers (quick generation)
		tasksPerProducer := dataCount / producerCount
		remainingTasks := dataCount % producerCount

		wg.Add(producerCount)

		for p := 0; p < producerCount; p++ {
			go func(producerID int) {
				defer wg.Done()

				tasksToGenerate := tasksPerProducer
				if producerID == producerCount-1 {
					tasksToGenerate += remainingTasks // Last producer handles any remainder
				}

				for j := 0; j < tasksToGenerate; j++ {
					taskChan <- &engine.Tasks{}
				}
			}(p)
		}

		// Close the channel after all producers are done
		go func() {
			wg.Wait()
			close(taskChan)
		}()

		// Start workers (slower processing)
		var workerWg sync.WaitGroup

		workerWg.Add(workerCount)

		for w := 0; w < workerCount; w++ {
			go func() {
				defer workerWg.Done()

				for task := range taskChan {
					simulateWork()
					processed <- task
				}
			}()
		}

		// Wait for all data to be processed
		go func() {
			workerWg.Wait()
			close(processed)
		}()

		count := 0
		for range processed {
			count++
			if count == dataCount {
				break
			}
		}

		duration := time.Since(startTime)
		b.ReportMetric(float64(duration.Milliseconds()), "ms/op")
		b.ReportMetric(float64(dataCount)/duration.Seconds(), "items/sec")
	}
}
