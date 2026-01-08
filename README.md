# Concurrent Job Scheduler – Go Worker Pool

A clean, interview-ready implementation of an asynchronous worker pool in Go demonstrating concurrency patterns, job coordination, backpressure handling, and graceful shutdown.

## Features

- **Asynchronous worker pool** using goroutines, channels, and wait groups
- **Job coordination** with buffered channels for backpressure control
- **Graceful shutdown** using context cancellation
- **OS signal handling** (SIGINT/SIGTERM) for clean termination
- **Zero external dependencies** - pure standard library

## Project Structure

```
go-worker-pool/
├── cmd/
│   └── main.go           # Entry point with job submission logic
├── job/
│   └── job.go            # Job struct and processing logic
├── worker/
│   ├── worker.go         # Individual worker implementation
│   └── pool.go           # Worker pool management
├── internal/
│   └── shutdown.go       # Signal handling for graceful shutdown
├── go.mod
└── README.md
```

## Key Concepts Demonstrated

### 1. Goroutines and Channels
- Each worker runs in its own goroutine
- Workers share a buffered job channel for communication
- Channel acts as a thread-safe queue

### 2. Backpressure
- Buffered channel (`jobQueueSize`) limits queue depth
- `Submit()` blocks when queue is full, preventing memory overflow

### 3. Synchronization with WaitGroup
- `sync.WaitGroup` tracks active workers
- `Wg.Add(1)` before starting each worker
- `Wg.Done()` when worker exits
- `Wg.Wait()` blocks until all workers finish

### 4. Context for Cancellation
- `context.Context` propagates shutdown signal
- `select` statement in worker listens for both jobs and context cancellation
- Enables graceful shutdown without killing goroutines abruptly

### 5. Select Statement
```go
select {
case <-ctx.Done():
    // Handle shutdown
case j := <-w.JobChan:
    // Process job
}
```

## How to Run

```bash
# Initialize module (update module path in all files)
go mod init github.com/Bharat1Rajput/go-worker-pool

# Run the program
go run cmd/main.go

# Test graceful shutdown (press Ctrl+C)
```

## Expected Output

```
Starting worker pool with 3 workers
Worker 1 started
Worker 2 started
Worker 3 started
Submitting Job ID: 1
Worker 1 processing Job ID: 1, Payload: Task-1
Submitting Job ID: 2
Worker 2 processing Job ID: 2, Payload: Task-2
...
Worker 1 completed Job ID: 1
Closing job channel...
Waiting for workers to complete...
Worker 1 shutting down gracefully
Worker 2 shutting down gracefully
Worker 3 shutting down gracefully
All workers stopped
Main: shutdown complete
```

## Interview Talking Points

### Architecture Decisions
- **Fixed-size pool**: Prevents resource exhaustion
- **Buffered channel**: Decouples producers from consumers
- **Shared channel**: Simple work distribution (workers pull from same queue)

### Concurrency Primitives
- **Goroutines**: Lightweight threads managed by Go runtime
- **Channels**: Safe communication between goroutines
- **WaitGroup**: Synchronization barrier for goroutine completion
- **Context**: Cancellation signal propagation

### Graceful Shutdown Flow
1. OS signal received (SIGINT/SIGTERM)
2. Context cancelled via `cancel()`
3. Workers detect cancellation in `select` statement
4. Workers exit their loops and call `Wg.Done()`
5. Pool waits for all workers via `Wg.Wait()`
6. Main function exits cleanly

### Backpressure Handling
- When job queue is full, `Submit()` blocks
- Prevents unlimited memory growth
- Natural flow control mechanism

## Potential Extensions (Not Implemented)

If asked in an interview about improvements:
- Add job priorities with multiple channels
- Implement job retry logic with exponential backoff
- Add metrics (jobs processed, success/failure rates)
- Support dynamic worker scaling
- Add job timeouts using context.WithTimeout
- Implement result collection channel

## Configuration

Modify in `cmd/main.go`:
```go
numWorkers := 3      // Number of concurrent workers
jobQueueSize := 10   // Buffer size for job channel
```

## License

MIT License - Free for personal and educational use