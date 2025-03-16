package main

import (
	"fmt"
	"sync"
	"time"
)

// simulateWork simulates a time-consuming task
func simulateWork(id int, duration time.Duration) {
	fmt.Printf("Task %d starting\n", id)
	time.Sleep(duration)
	fmt.Printf("Task %d completed\n", id)
}

// basicGoroutineExample demonstrates simple goroutine creation
func basicGoroutineExample() {
	fmt.Println("\n=== Basic Goroutine Example ===")

	// Starting a goroutine
	go simulateWork(1, time.Millisecond*500)

	// Main thread continues execution
	fmt.Println("Main thread continues...")

	// Wait a bit to see the goroutine complete
	time.Sleep(time.Second)
}

// waitGroupExample shows how to synchronize multiple goroutines
func waitGroupExample() {
	fmt.Println("\n=== WaitGroup Example ===")

	var wg sync.WaitGroup

	// Launch multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			simulateWork(id, time.Millisecond*300)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All tasks completed")
}

// concurrentNumberProcessing demonstrates practical use of goroutines
func concurrentNumberProcessing() {
	fmt.Println("\n=== Concurrent Number Processing ===")

	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Process numbers concurrently
	for _, n := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// Simulate processing by calculating square
			time.Sleep(time.Millisecond * 200)
			results <- num * num
		}(n)
	}

	// Close results channel when all processing is done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results
	fmt.Println("Processing results:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

func main() {
	fmt.Println("Starting concurrency examples...")

	// Run examples
	basicGoroutineExample()
	waitGroupExample()
	concurrentNumberProcessing()

	fmt.Println("\nAll examples completed")
}
