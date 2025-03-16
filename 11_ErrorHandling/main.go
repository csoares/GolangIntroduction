package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Custom error type for validation
type ValidationError struct {
	Field string
	Issue string
}

// Implement the error interface
func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", v.Field, v.Issue)
}

// Function demonstrating basic error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Function demonstrating custom error types
func validateUsername(username string) error {
	if len(username) < 3 {
		return &ValidationError{"username", "length must be at least 3 characters"}
	}
	if strings.Contains(username, " ") {
		return &ValidationError{"username", "spaces are not allowed"}
	}
	return nil
}

// Function demonstrating error wrapping
func readConfig(filename string) error {
	_, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	return nil
}

func main() {
	// Example 1: Basic error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Error in division: %v\n", err)
	} else {
		fmt.Printf("Division result: %f\n", result)
	}

	// Example 2: Custom error types
	username := "jo"
	if err := validateUsername(username); err != nil {
		if validErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Validation error: %v\n", validErr)
		}
	}

	// Example 3: Error wrapping
	err = readConfig("nonexistent.conf")
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		// Check if the underlying error is a file not found error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The config file does not exist")
		}
	}
}
