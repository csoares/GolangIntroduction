package main

import (
	"fmt"
	"time"
)

// producer sends numbers to the channel
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
	}
	close(ch)
}

// consumer receives numbers from the channel
func consumer(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Printf("Consuming: %d\n", num)
		time.Sleep(time.Second) // Simulate processing time
	}
	done <- true
}

func main() {
	// Create a buffered channel with capacity 3
	numbers := make(chan int, 3)
	done := make(chan bool)

	fmt.Println("Using buffered channel with capacity 3")

	// Start producer and consumer
	go producer(numbers)
	go consumer(numbers, done)

	// Wait for consumer to finish
	<-done

	// Demonstrate channel capacity
	fmt.Println("\nDemonstrating buffer overflow protection:")
	buffered := make(chan int, 2)

	// Fill the buffer
	buffered <- 1
	fmt.Println("Added 1 to buffer")
	buffered <- 2
	fmt.Println("Added 2 to buffer")

	// This would block if we didn't receive in a separate goroutine
	go func() {
		time.Sleep(time.Second)
		x := <-buffered
		fmt.Printf("Received %d from buffer\n", x)
	}()

	// Now we can send another value
	fmt.Println("Sending 3 to buffer...")
	buffered <- 3
	fmt.Println("Successfully sent 3 to buffer")

	// Give time for the last receive operation
	time.Sleep(time.Second * 2)
}
