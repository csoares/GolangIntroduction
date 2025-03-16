package main

import (
	"fmt"
	"math"
)

// Shape interface defines the behavior for all shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle represents a circular shape
type Circle struct {
	radius float64
}

// Rectangle represents a rectangular shape
type Rectangle struct {
	width  float64
	height float64
}

// Area calculates and returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter calculates and returns the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area calculates and returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter calculates and returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// PrintShapeInfo prints the area and perimeter of any shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
