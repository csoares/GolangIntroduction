package main

import (
	"fmt"
	"strings"
)

// Student represents a student with a name and their grades
type Student struct {
	Name   string
	Grades map[string]int
}

func main() {
	// Example 1: Basic Map Operations
	fmt.Println("Example 1: Basic Map Operations")
	// Initialize a map using make
	scores := make(map[string]int)

	// Add elements
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92

	// Access and print elements
	fmt.Printf("Alice's score: %d\n", scores["Alice"])

	// Check if key exists
	score, exists := scores["David"]
	if !exists {
		fmt.Println("David's score not found")
	} else {
		fmt.Printf("David's score: %d\n", score)
	}

	// Delete an element
	delete(scores, "Bob")

	// Iterate through the map
	fmt.Println("\nAll scores:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Example 2: Student Grade Tracker
	fmt.Println("\nExample 2: Student Grade Tracker")
	// Initialize a student with grades
	student := Student{
		Name: "Alice",
		Grades: map[string]int{
			"Math":    95,
			"Science": 88,
			"English": 92,
		},
	}

	// Add a new grade
	student.Grades["History"] = 89

	// Calculate average grade
	total := 0
	for _, grade := range student.Grades {
		total += grade
	}
	average := float64(total) / float64(len(student.Grades))

	fmt.Printf("%s's grades:\n", student.Name)
	for subject, grade := range student.Grades {
		fmt.Printf("%s: %d\n", subject, grade)
	}
	fmt.Printf("Average grade: %.2f\n", average)

	// Example 3: Word Frequency Counter
	fmt.Println("\nExample 3: Word Frequency Counter")
	text := "the quick brown fox jumps over the lazy dog the quick brown fox"
	words := strings.Fields(text)

	// Count word frequencies
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	// Print word frequencies
	fmt.Println("Word frequencies:")
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
