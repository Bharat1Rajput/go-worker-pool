package job

import (
	"fmt"
	"time"
)

// Job represents a unit of work to be processed by workers
type Job struct {
	ID      int
	Payload string
}

// Process simulates processing the job
// In a real scenario, this would contain actual business logic
func (j Job) Process() {
	fmt.Printf("Worker processing Job ID: %d, Payload: %s\n", j.ID, j.Payload)
	// Simulate work taking time
	time.Sleep(time.Second * 2)
	fmt.Printf("Worker completed Job ID: %d\n", j.ID)
}
