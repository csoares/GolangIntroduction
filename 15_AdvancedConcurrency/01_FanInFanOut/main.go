//go:build ignore

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// rng is a local random number generator for this package
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// processData simulates a time-consuming task with context support
func processData(ctx context.Context, id int, data int) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		// Create a timeout for this specific operation
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		case <-time.After(time.Duration(rng.Intn(1000)) * time.Millisecond):
			return data * 2, nil
		}
	}
}

// fanOut distributes work to multiple goroutines with context support
func fanOut(ctx context.Context, input []int, numWorkers int) []<-chan int {
	var channels []<-chan int
	for i := range numWorkers {
		ch := make(chan int, len(input)/numWorkers+1) // Buffered to prevent blocking
		go func(workerID int, ch chan int) {
			defer close(ch)
			// each worker processes a subset of the input
			for j := workerID; j < len(input); j += numWorkers {
				select {
				case <-ctx.Done():
					return
				default:
					result, err := processData(ctx, workerID, input[j])
					if err != nil {
						return
					}
					select {
					case ch <- result:
					case <-ctx.Done():
						return
					}
				}
			}
		}(i, ch)
		channels = append(channels, ch)
	}
	return channels
}

// fanIn combines multiple channels into a single channel with context support
func fanIn(ctx context.Context, channels []<-chan int) <-chan int {
	var wg sync.WaitGroup
	mergedCh := make(chan int, len(channels)*2) // Buffered to reduce blocking

	// start a goroutine for each input channel
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for {
				select {
				case val, ok := <-c:
					if !ok {
						return
					}
					select {
					case mergedCh <- val:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
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
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// sample input data
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 3

	fmt.Printf("Processing %d items with %d workers\n", len(input), numWorkers)

	// distribute work across workers
	channels := fanOut(ctx, input, numWorkers)

	// combine results
	results := fanIn(ctx, channels)

	// collect and print results
	processedData := make([]int, 0, len(input))
	for result := range results {
		processedData = append(processedData, result)
	}

	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Output: %v\n", processedData)
}
