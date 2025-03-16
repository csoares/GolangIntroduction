// Package main is the entry point for all Go programs
package main

// Import the fmt package which contains functions for formatted I/O operations
import "fmt"

// The main function is the entry point of the program
// Every executable Go program must have a main function in the main package
func main() {
	// Println is a function from the fmt package that prints a line to standard output
	// It automatically adds a newline character at the end
	fmt.Println("Hello, World!")

	// You can also use Printf for formatted output
	fmt.Printf("Welcome to %s programming!\n", "Go")

	// Note: In Go, unused imports or variables will cause compilation errors
	// This encourages clean code practices
}