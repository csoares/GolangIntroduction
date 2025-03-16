// Package main is the entry point for all Go programs
package main

// Import the fmt package for formatted I/O operations
import (
	"fmt"
	"time"
)

func main() {
	// If Statements
	fmt.Println("\n--- If Statements ---")

	// Basic if statement
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If with a short statement (initialize and check in one line)
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D or below")
	}
	// Note: 'score' is only in scope within the if-else block

	// For Loops
	fmt.Println("\n--- For Loops ---")

	// Basic for loop (similar to C/Java)
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop as a while loop
	fmt.Println("For as while:")
	count := 0
	for count < 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// Infinite loop with break
	fmt.Println("Infinite loop with break:")
	sum := 0
	for {
		sum++
		if sum > 5 {
			break // Exit the loop
		}
		fmt.Printf("%d ", sum)
	}
	fmt.Println()

	// For loop with continue
	fmt.Println("For loop with continue:")
	for i := 0; i < 10; i++ {
		// Skip even numbers
		if i%2 == 0 {
			continue // Skip the rest of this iteration
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop with range (for arrays, slices, maps, strings)
	fmt.Println("For loop with range:")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// If you only need the value, use _ for the index
	fmt.Println("Range with only values:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// If you only need the index, omit the second variable
	fmt.Println("Range with only indices:")
	for index := range numbers {
		fmt.Printf("%d ", index)
	}
	fmt.Println()

	// Switch Statements
	fmt.Println("\n--- Switch Statements ---")

	// Basic switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("End of work week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with no expression (acts like if-else chain)
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Switch with fallthrough
	n := 4
	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
		fallthrough // Execution will continue to the next case
	case 5:
		fmt.Println("Five or Four")
	default:
		fmt.Println("Other number")
	}

	// Type switch (used with interfaces)
	fmt.Println("\n--- Type Switch ---")
	// Declare an interface variable that can hold values of different types
	var x interface{}

	// Assign a string value to x
	x = "Hello"

	// Type switch to determine the type of x
	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Printf("x is an integer: %d\n", v)
	case float64:
		fmt.Printf("x is a float: %f\n", v)
	case string:
		fmt.Printf("x is a string: %s\n", v)
	case bool:
		fmt.Printf("x is a boolean: %t\n", v)
	default:
		fmt.Printf("x is of an unknown type: %T\n", v)
	}

	// Try with different types
	x = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("x is an integer: %d\n", v)
	case string:
		fmt.Printf("x is a string: %s\n", v)
	default:
		fmt.Printf("x is of type: %T\n", v)
	}

	x = true
	switch v := x.(type) {
	case bool:
		fmt.Printf("x is a boolean: %t\n", v)
	default:
		fmt.Printf("x is of type: %T\n", v)
	}
}
