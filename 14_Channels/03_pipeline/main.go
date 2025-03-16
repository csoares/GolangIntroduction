package main

import (
	"fmt"
)

// generator creates a channel that generates numbers from 1 to max
func generator(max int) <-chan int {
	outChan := make(chan int)
	go func() {
		for i := 1; i <= max; i++ {
			outChan <- i
		}
		close(outChan)
	}()
	return outChan
}

// square receives numbers and returns their squares
func square(in <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for num := range in {
			outChan <- num * num
		}
		close(outChan)
	}()
	return outChan
}

// filter receives numbers and only passes even numbers
func filter(in <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for num := range in {
			if num%2 == 0 {
				outChan <- num
			}
		}
		close(outChan)
	}()
	return outChan
}

func main() {
	// Create a pipeline: generate numbers -> square them -> filter even numbers
	fmt.Println("Starting pipeline: generate(1-5) -> square -> filter even numbers")

	// Stage 1: Generate numbers 1-5
	numbers := generator(5)

	// Stage 2: Square the numbers
	squares := square(numbers)

	// Stage 3: Filter even squares
	evenSquares := filter(squares)

	// Consume the final results
	fmt.Println("Results (even squares):")
	for result := range evenSquares {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
