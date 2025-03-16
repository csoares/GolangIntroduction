package main

import (
	"fmt"
	"time"
)

// ping sends a message through a channel
func ping(ch chan string) {
	fmt.Println("Sending ping...")
	ch <- "ping" // Send message through channel
}

// pong receives a message and sends back a response
func pong(ch1, ch2 chan string) {
	msg := <-ch1 // Receive message from first channel
	fmt.Printf("Received: %s\n", msg)
	time.Sleep(time.Second) // Simulate some processing
	fmt.Println("Sending pong...")
	ch2 <- "pong" // Send response through second channel
}

func main() {
	// Create unbuffered channels
	ping_channel := make(chan string)
	pong_channel := make(chan string)

	// Start goroutines
	go ping(ping_channel)
	go pong(ping_channel, pong_channel)

	// Wait for and print the final response
	response := <-pong_channel
	fmt.Printf("Final response: %s\n", response)

	// Demonstrate channel closing
	fmt.Println("\nDemonstrating channel closing:")
	numbers := make(chan int, 3)

	// Send values
	numbers <- 1
	numbers <- 2
	numbers <- 3
	close(numbers) // Close the channel

	// Read all values until channel is closed
	for num := range numbers {
		fmt.Printf("Received number: %d\n", num)
	}

	// Demonstrate checking if channel is closed
	val, ok := <-numbers
	if !ok {
		fmt.Println("Channel is closed, received zero value:", val)
	}
}
