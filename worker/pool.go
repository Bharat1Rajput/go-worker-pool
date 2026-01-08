package worker

import (
	"context"
	"fmt"
	"sync"

	"github.com/Bharat1Rajput/go-worker-pool/job"
)

// Pool manages a fixed number of workers and job distribution
type Pool struct {
	NumWorkers int
	JobChan    chan job.Job
	Wg         *sync.WaitGroup
}

// NewPool creates a new worker pool with the specified configuration
// Buffered channel provides backpressure control
func NewPool(numWorkers int, jobQueueSize int) *Pool {
	return &Pool{
		NumWorkers: numWorkers,
		JobChan:    make(chan job.Job, jobQueueSize),
		Wg:         &sync.WaitGroup{},
	}
}

// Start initializes all workers and begins processing
// Each worker is a goroutine that listens on the shared job channel
func (p *Pool) Start(ctx context.Context) {
	fmt.Printf("Starting worker pool with %d workers\n", p.NumWorkers)

	// Launch fixed number of worker goroutines
	for i := 1; i <= p.NumWorkers; i++ {
		p.Wg.Add(1)
		w := &Worker{
			ID:      i,
			JobChan: p.JobChan,
			Wg:      p.Wg,
		}
		go w.Start(ctx)
	}
}

// Submit adds a job to the queue
// Blocks if the queue is full (backpressure)
func (p *Pool) Submit(j job.Job) {
	p.JobChan <- j
}

// Shutdown closes the job channel and waits for all workers to finish
func (p *Pool) Shutdown() {
	fmt.Println("Closing job channel...")
	close(p.JobChan)
	fmt.Println("Waiting for workers to complete...")
	p.Wg.Wait()
	fmt.Println("All workers stopped")
}
