//go:build ignore

package main

import (
	"context"
	"fmt"
)

// generator creates a channel that generates numbers from 1 to max
func generator(ctx context.Context, max int) <-chan int {
	outChan := make(chan int, max) // Buffered to prevent blocking
	go func() {
		defer close(outChan)
		for i := 1; i <= max; i++ {
			select {
			case outChan <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return outChan
}

// square receives numbers and returns their squares
func square(ctx context.Context, in <-chan int) <-chan int {
	outChan := make(chan int, cap(in)) // Buffered to prevent blocking
	go func() {
		defer close(outChan)
		for num := range in {
			select {
			case outChan <- num * num:
			case <-ctx.Done():
				return
			}
		}
	}()
	return outChan
}

// filter receives numbers and only passes even numbers
func filter(ctx context.Context, in <-chan int) <-chan int {
	outChan := make(chan int, cap(in)) // Buffered to prevent blocking
	go func() {
		defer close(outChan)
		for num := range in {
			if num%2 == 0 {
				select {
				case outChan <- num:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return outChan
}

func main() {
	ctx := context.Background()

	// Create a pipeline: generate numbers -> square them -> filter even numbers
	fmt.Println("Starting pipeline: generate(1-5) -> square -> filter even numbers")

	// Stage 1: Generate numbers 1-5
	numbers := generator(ctx, 5)

	// Stage 2: Square the numbers
	squares := square(ctx, numbers)

	// Stage 3: Filter even squares
	evenSquares := filter(ctx, squares)

	// Consume the final results
	fmt.Println("Results (even squares):")
	for result := range evenSquares {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
