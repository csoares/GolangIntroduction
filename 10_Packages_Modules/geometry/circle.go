// Package geometry provides geometric calculations
package geometry

import "math"

// Pi is the mathematical constant π
const Pi = math.Pi

// Circle represents a circle shape
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle with the given radius
func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

// Circumference calculates the circumference of the circle
func (c *Circle) Circumference() float64 {
	return 2 * Pi * c.Radius
}

// Diameter calculates the diameter of the circle
func (c *Circle) Diameter() float64 {
	return 2 * c.Radius
}
