package main

import (
	"fmt"
)

func main() {
	// Demonstrate custom data structures
	fmt.Println("=== Custom Data Structures ===\n")
	// Create and use our custom data structure
	list := &StringList{}
	list.Add("first")
	list.Add("second")

	// Demonstrate interface implementations
	fmt.Println("Demonstrating StringList interfaces:")
	DemonstrateInterfaces(list)

	// Demonstrate direct usage
	fmt.Printf("\nDirect usage - Count: %d\n", list.Count())

	// Demonstrate shapes
	fmt.Println("\n=== Shapes ===\n")
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}

	fmt.Println("Circle:")
	PrintShapeInfo(circle)

	fmt.Println("\nRectangle:")
	PrintShapeInfo(rectangle)
}