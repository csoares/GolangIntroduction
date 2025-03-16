// Package main is the entry point for all Go programs
package main

// Import the fmt package for formatted I/O operations
import "fmt"

func main() {
	// Arrays
	fmt.Println("\n--- Arrays ---")

	// Array declaration with explicit size
	// Arrays in Go have a fixed size that cannot be changed
	var numbers [5]int // Declares an array of 5 integers, initialized to zero values
	fmt.Println("Empty array:", numbers)

	// Assigning values to array elements
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Filled array:", numbers)

	// Array declaration with initialization
	fruits := [3]string{"apple", "banana", "cherry"}
	fmt.Println("Fruits array:", fruits)

	// Let the compiler count the array size
	colors := [...]string{"red", "green", "blue", "yellow"}
	fmt.Println("Colors array:", colors)
	fmt.Println("Number of colors:", len(colors))

	// Accessing array elements
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Last color:", colors[len(colors)-1])

	// Arrays are value types in Go
	// When you assign an array to another variable, it creates a copy
	colors2 := colors
	colors2[0] = "purple" // This doesn't affect the original array
	fmt.Println("Original colors:", colors)
	fmt.Println("Modified colors2:", colors2)

	// Multi-dimensional arrays
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Matrix[1][2]:", matrix[1][2]) // Accessing element at row 1, column 2

	// Iterating over arrays
	fmt.Println("Iterating over fruits:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("fruits[%d] = %s\n", i, fruits[i])
	}

	// Using range to iterate
	fmt.Println("Iterating over colors using range:")
	for index, value := range colors {
		fmt.Printf("colors[%d] = %s\n", index, value)
	}

	// Slices
	fmt.Println("\n--- Slices ---")

	// Slice declaration
	// Slices are more flexible than arrays and are used more often
	var scores []int // Declares a nil slice (no underlying array yet)
	fmt.Println("Nil slice:", scores, "Length:", len(scores), "Capacity:", cap(scores))

	// Check if a slice is nil
	if scores == nil {
		fmt.Println("scores is nil")
	}

	// Creating a slice with make
	// make([]T, length, capacity)
	scores = make([]int, 5) // Length 5, capacity 5
	fmt.Println("Slice created with make:", scores, "Length:", len(scores), "Capacity:", cap(scores))

	// Creating a slice with literal values
	vegetables := []string{"carrot", "broccoli", "spinach"}
	fmt.Println("Vegetables slice:", vegetables, "Length:", len(vegetables), "Capacity:", cap(vegetables))

	// Creating a slice from an array
	fruitSlice := fruits[:] // Slice of the entire array
	fmt.Println("Fruit slice from array:", fruitSlice)

	// Slices are reference types
	// When you assign a slice to another variable, both refer to the same underlying array
	veggies := vegetables
	veggies[0] = "cucumber" // This affects the original slice
	fmt.Println("Original vegetables:", vegetables)
	fmt.Println("Modified veggies:", veggies)

	// Slicing operations
	numbers = [5]int{10, 20, 30, 40, 50} // Reset numbers array

	// Slice from index 1 to 3 (exclusive)
	slice1 := numbers[1:3] // [20, 30]
	fmt.Println("slice1 (numbers[1:3]):", slice1)

	// Slice from beginning to index 3 (exclusive)
	slice2 := numbers[:3] // [10, 20, 30]
	fmt.Println("slice2 (numbers[:3]):", slice2)

	// Slice from index 2 to end
	slice3 := numbers[2:] // [30, 40, 50]
	fmt.Println("slice3 (numbers[2:]):", slice3)

	// Append operations
	// When appending to a slice, if the capacity is insufficient,
	// Go will automatically create a new underlying array with more capacity
	fmt.Println("\n--- Append Operations ---")

	// Create a slice with initial capacity
	numbers2 := make([]int, 0, 3)
	fmt.Printf("Initial slice - Length: %d, Capacity: %d\n", len(numbers2), cap(numbers2))

	// Append elements and observe capacity changes
	for i := 1; i <= 5; i++ {
		numbers2 = append(numbers2, i)
		fmt.Printf("After appending %d - Length: %d, Capacity: %d\n", i, len(numbers2), cap(numbers2))
	}

	// Append multiple elements at once
	numbers2 = append(numbers2, 6, 7, 8)
	fmt.Println("After appending multiple elements:", numbers2)

	// Append one slice to another
	extraNumbers := []int{9, 10}
	numbers2 = append(numbers2, extraNumbers...)
	fmt.Println("After appending another slice:", numbers2)

	// Copy slices
	fmt.Println("\n--- Copy Operations ---")

	// Create a destination slice
	dest := make([]int, len(numbers2))
	copied := copy(dest, numbers2)
	fmt.Printf("Copied %d elements\n", copied)
	fmt.Println("Source slice:", numbers2)
	fmt.Println("Destination slice:", dest)

	// Demonstrate that copies are independent
	dest[0] = 999
	fmt.Println("After modifying destination:")
	fmt.Println("Source slice:", numbers2)
	fmt.Println("Modified destination slice:", dest)
}
