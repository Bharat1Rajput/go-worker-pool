package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/Bharat1Rajput/go-worker-pool/job"
)

// Worker represents a single worker that processes jobs
type Worker struct {
	ID      int
	JobChan <-chan job.Job
	Wg      *sync.WaitGroup
}

// Start begins the worker's processing loop
// It listens for jobs on the job channel and processes them
// Uses select to handle graceful shutdown via context cancellation
func (w *Worker) Start(ctx context.Context) {
	defer w.Wg.Done()
	fmt.Printf("Worker %d started\n", w.ID)

	for {
		select {
		case <-ctx.Done():
			// Context cancelled - shutdown signal received
			fmt.Printf("Worker %d shutting down gracefully\n", w.ID)
			return
		case j, ok := <-w.JobChan:
			if !ok {
				// Channel closed - no more jobs
				fmt.Printf("Worker %d: job channel closed\n", w.ID)
				return
			}
			// Process the job
			j.Process()
		}
	}
}
