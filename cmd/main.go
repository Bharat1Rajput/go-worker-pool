package main

import (
	"fmt"
	"time"

	"github.com/Bharat1Rajput/go-worker-pool/internal"
	"github.com/Bharat1Rajput/go-worker-pool/job"
	"github.com/Bharat1Rajput/go-worker-pool/worker"
)

func main() {
	// Setup graceful shutdown with context cancellation
	ctx, cancel := internal.SetupSignalHandler()
	defer cancel()

	// Configuration
	numWorkers := 3
	jobQueueSize := 10 // Buffered channel for backpressure control

	// Create and start the worker pool
	pool := worker.NewPool(numWorkers, jobQueueSize)
	pool.Start(ctx)

	// Simulate submitting jobs
	go func() {
		for i := 1; i <= 15; i++ {
			j := job.Job{
				ID:      i,
				Payload: fmt.Sprintf("Task-%d", i),
			}
			fmt.Printf("Submitting Job ID: %d\n", j.ID)
			pool.Submit(j)
			time.Sleep(time.Millisecond * 500)
		}
		// After submitting all jobs, initiate shutdown
		pool.Shutdown()
		cancel()
	}()

	// Wait for context cancellation (either from signal or after jobs complete)
	<-ctx.Done()
	fmt.Println("Main: shutdown complete")
}
