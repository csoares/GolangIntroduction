package main

import (
	"fmt"
	"time"
)

// Task represents a job to be done
type Task struct {
	ID     int
	Result int
}

// worker processes tasks from the tasks channel and sends results to the results channel
func worker(id int, tasks <-chan Task, results chan<- Task) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		time.Sleep(time.Second)   // Simulate work
		task.Result = task.ID * 2 // Simple computation
		results <- task
	}
}

func main() {
	// Create channels for tasks and results
	tasks := make(chan Task, 10)
	results := make(chan Task, 10)

	// Start 3 workers
	numWorkers := 3
	fmt.Printf("Starting %d workers\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, tasks, results)
	}

	// Send 5 tasks
	numTasks := 5
	fmt.Printf("Sending %d tasks\n", numTasks)
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks)

	// Collect results
	fmt.Println("\nCollecting results:")
	for i := 1; i <= numTasks; i++ {
		result := <-results
		fmt.Printf("Task %d result: %d\n", result.ID, result.Result)
	}
}
