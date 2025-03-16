package main

import (
	"calculator-app/calculator"
	"calculator-app/geometry"
	"fmt"
)

func main() {
	// Demonstrate calculator package usage
	fmt.Println("\nCalculator Package Demo:")
	fmt.Printf("Addition: %.2f\n", calculator.Add(10, 5))
	fmt.Printf("Square Root: %.2f\n", calculator.SquareRoot(16))
	fmt.Printf("Power: %.2f\n", calculator.Power(2, 3))

	// Demonstrate geometry package usage
	fmt.Println("\nGeometry Package Demo:")

	// Circle calculations
	circle := geometry.NewCircle(5)
	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Circumference: %.2f\n", circle.Circumference())

	// Rectangle calculations
	rect := geometry.NewRectangle(4, 6)
	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())
	fmt.Printf("Is Square: %v\n", rect.IsSquare())

	// Demonstrate rectangle scaling
	rect.Scale(2)
	fmt.Printf("Scaled Rectangle Area: %.2f\n", rect.Area())
}
