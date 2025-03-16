package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Semaphore implements a counting semaphore
type Semaphore struct {
	permits int
	ch      chan struct{}
}

// NewSemaphore creates a new semaphore with the given number of permits
func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{
		permits: permits,
		ch:      make(chan struct{}, permits),
	}
}

// Acquire acquires a permit from the semaphore
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release releases a permit back to the semaphore
func (s *Semaphore) Release() {
	<-s.ch
}

// Resource represents a shared resource
type Resource struct {
	id int
}

// UseResource simulates using a shared resource
func (r *Resource) UseResource() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

// ResourcePool manages a pool of resources with semaphore-controlled access
type ResourcePool struct {
	resources []*Resource
	sem       *Semaphore
}

// NewResourcePool creates a new resource pool
func NewResourcePool(numResources int) *ResourcePool {
	pool := &ResourcePool{
		resources: make([]*Resource, numResources),
		sem:       NewSemaphore(numResources),
	}

	for i := 0; i < numResources; i++ {
		pool.resources[i] = &Resource{id: i}
	}

	return pool
}

// UseResource simulates a worker using a resource from the pool
func (p *ResourcePool) UseResource(workerID int) {
	p.sem.Acquire()
	defer p.sem.Release()

	// Simulate resource selection
	resourceID := rand.Intn(len(p.resources))
	resource := p.resources[resourceID]

	fmt.Printf("Worker %d acquired resource %d\n", workerID, resource.id)
	resource.UseResource()
	fmt.Printf("Worker %d released resource %d\n", workerID, resource.id)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create a resource pool with 3 resources
	pool := NewResourcePool(3)

	// Simulate 5 workers trying to use resources
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Each worker tries to use resources multiple times
			for j := 0; j < 3; j++ {
				pool.UseResource(id)
				// Wait a bit before next attempt
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}
