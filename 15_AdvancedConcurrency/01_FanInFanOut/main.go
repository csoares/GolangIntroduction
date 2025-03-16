package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// simulate a time-consuming task
func processData(id int, data int) int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return data * 2
}

// fanOut distributes work to multiple goroutines
func fanOut(input []int, numWorkers int) []<-chan int {
	var channels []<-chan int
	for i := 0; i < numWorkers; i++ {
		ch := make(chan int)
		go func(workerID int, ch chan int) {
			defer close(ch)
			// each worker processes a subset of the input
			for j := workerID; j < len(input); j += numWorkers {
				result := processData(workerID, input[j])
				ch <- result
			}
		}(i, ch)
		channels = append(channels, ch)
	}
	return channels
}

// fanIn combines multiple channels into a single channel
func fanIn(channels []<-chan int) <-chan int {
	var wg sync.WaitGroup
	mergedCh := make(chan int)

	// start a goroutine for each input channel
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				mergedCh <- val
			}
		}(ch)
	}

	// close mergedCh when all input channels are done
	go func() {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func main() {
	// sample input data
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 3

	fmt.Printf("Processing %d items with %d workers\n", len(input), numWorkers)

	// distribute work across workers
	channels := fanOut(input, numWorkers)

	// combine results
	results := fanIn(channels)

	// collect and print results
	processedData := make([]int, 0, len(input))
	for result := range results {
		processedData = append(processedData, result)
	}

	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Output: %v\n", processedData)
}
