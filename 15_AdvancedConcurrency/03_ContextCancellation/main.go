package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker represents a long-running task
type Worker struct {
	id int
}

// DoWork simulates a long-running task that can be cancelled
func (w *Worker) DoWork(ctx context.Context) error {
	fmt.Printf("Worker %d: Starting work\n", w.id)

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Received cancellation signal\n", w.id)
			return ctx.Err()
		default:
			// Simulate some work
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			fmt.Printf("Worker %d: Completed iteration %d\n", w.id, i)
		}
	}
}

// WorkerPool manages multiple workers with graceful shutdown
type WorkerPool struct {
	workers []*Worker
	wg      sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers
func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		workers: make([]*Worker, numWorkers),
	}

	for i := 0; i < numWorkers; i++ {
		pool.workers[i] = &Worker{id: i}
	}

	return pool
}

// Start launches all workers and manages their lifecycle
func (p *WorkerPool) Start(ctx context.Context) {
	for _, worker := range p.workers {
		p.wg.Add(1)
		go func(w *Worker) {
			defer p.wg.Done()
			if err := w.DoWork(ctx); err != nil {
				fmt.Printf("Worker %d stopped: %v\n", w.id, err)
			}
		}(worker)
	}
}

// Wait blocks until all workers have completed
func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create and start worker pool
	pool := NewWorkerPool(3)
	pool.Start(ctx)

	// Wait for either context cancellation or user interrupt
	fmt.Println("Press Ctrl+C to stop the workers...")
	<-ctx.Done()
	fmt.Println("\nShutting down workers...")

	// Wait for all workers to complete
	pool.Wait()
	fmt.Println("All workers have stopped")
}
