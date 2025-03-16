// Package main is the entry point for all Go programs
package main

// Import the fmt package for formatted I/O operations
import "fmt"

// Basic function declaration
// This function takes no parameters and returns no values
func sayHello() {
	fmt.Println("Hello from a function!")
}

// Function with parameters
// This function takes two parameters of type int
func add(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// Function with return value
// This function takes two parameters and returns an int
func multiply(a int, b int) int {
	return a * b
}

// Function with multiple parameters of the same type
// When consecutive parameters have the same type, you can omit all but the last type
func subtract(a, b int) int {
	return a - b
}

// Function with multiple return values
// Go functions can return multiple values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function with named return values
// You can name the return values, which initializes them as variables
func calculateStats(numbers []int) (min, max, total int) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}

	min = numbers[0]
	max = numbers[0]
	total = 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		total += num
	}

	// A naked return returns the named return values
	return
}

// Variadic function (variable number of arguments)
// The ellipsis ... before the type name indicates that the function accepts
// zero or more arguments of that type
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as a value
// Functions are first-class citizens in Go and can be assigned to variables
func getGreetingFunction() func(string) string {
	return func(name string) string {
		return "Hello, " + name
	}
}

// Recursive function
// A function that calls itself
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Defer statement
// Defers the execution of a function until the surrounding function returns
func deferExample() {
	fmt.Println("Start of function")

	// This will be executed last, just before the function returns
	defer fmt.Println("This is deferred")

	fmt.Println("End of function")
}

// Multiple defer statements are executed in LIFO order (last in, first out)
func multipleDeferExample() {
	fmt.Println("Start of function")

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")

	fmt.Println("End of function")
}

func main() {
	// Calling a basic function
	fmt.Println("\n--- Basic Function ---")
	sayHello()

	// Calling a function with parameters
	fmt.Println("\n--- Function with Parameters ---")
	add(5, 3)

	// Calling a function with return value
	fmt.Println("\n--- Function with Return Value ---")
	result := multiply(4, 7)
	fmt.Println("Product:", result)

	// Calling a function with multiple return values
	fmt.Println("\n--- Function with Multiple Return Values ---")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	// Error handling example
	quotient, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}

	// Calling a function with named return values
	fmt.Println("\n--- Function with Named Return Values ---")
	numbers := []int{7, 2, 9, 3, 6}
	min, max, total := calculateStats(numbers)
	fmt.Printf("For numbers %v: Min=%d, Max=%d, Sum=%d\n", numbers, min, max, total)

	// Calling a variadic function
	fmt.Println("\n--- Variadic Function ---")
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 5, 10, 15, 20:", sum(5, 10, 15, 20))

	// Passing a slice to a variadic function
	nums := []int{1, 2, 3, 4, 5}
	// The ... after the slice unpacks it into individual arguments
	fmt.Println("Sum of slice:", sum(nums...))

	// Function as a value
	fmt.Println("\n--- Function as a Value ---")
	greet := getGreetingFunction()
	message := greet("Alice")
	fmt.Println(message)

	// Anonymous function (function literal)
	fmt.Println("\n--- Anonymous Function ---")
	func() {
		fmt.Println("Hello from an anonymous function!")
	}() // The parentheses () at the end call the function immediately

	// Anonymous function assigned to a variable
	doubler := func(x int) int {
		return x * 2
	}
	fmt.Println("Double of 5:", doubler(5))

	// Recursive function
	fmt.Println("\n--- Recursive Function ---")
	fmt.Println("Factorial of 5:", factorial(5))

	// Defer example
	fmt.Println("\n--- Defer Example ---")
	deferExample()

	// Multiple defer example
	fmt.Println("\n--- Multiple Defer Example ---")
	multipleDeferExample()
}
