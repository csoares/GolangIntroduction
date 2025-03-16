package main

import "fmt"

// Stringer interface from fmt package for string representation
type Stringer interface {
	String() string
}

// Counter interface for counting operations
type Counter interface {
	Count() int
	Add(item string)
}

// StringList is a custom data structure that implements multiple interfaces
type StringList struct {
	items []string
}

// String returns a string representation of the list
func (sl *StringList) String() string {
	return fmt.Sprintf("StringList{items: %v}", sl.items)
}

// Count returns the number of items in the list
func (sl *StringList) Count() int {
	return len(sl.items)
}

// Add appends a new item to the list
func (sl *StringList) Add(item string) {
	sl.items = append(sl.items, item)
}

// DemonstrateInterfaces shows different ways to work with interfaces
func DemonstrateInterfaces(value interface{}) {
	// Type assertion example
	if stringer, ok := value.(Stringer); ok {
		fmt.Printf("As Stringer: %s\n", stringer.String())
	}

	// Type switch example
	switch v := value.(type) {
	case Counter:
		fmt.Printf("As Counter: %d items\n", v.Count())
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
