package main

import (
	"fmt"
	"math"
)

// Shape defines the behavior for all geometric shapes.
// Implementations must be safe for concurrent use.
type Shape interface {
	// Area returns the area of the shape in square units.
	Area() float64
	// Perimeter returns the perimeter of the shape.
	// For circles, this returns the circumference.
	Perimeter() float64
}

// Circle represents a circular shape.
type Circle struct {
	radius float64
}

// Rectangle represents a rectangular shape.
type Rectangle struct {
	width  float64
	height float64
}

// Area calculates and returns the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter calculates and returns the perimeter (circumference) of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area calculates and returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter calculates and returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// PrintShapeInfo prints the area and perimeter of any shape.
// It accepts any type that implements the Shape interface.
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
