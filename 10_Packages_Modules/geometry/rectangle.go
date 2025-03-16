// Package geometry provides geometric calculations
package geometry

// Rectangle represents a rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle creates a new Rectangle with the given dimensions
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// IsSquare checks if the rectangle is a square
func (r *Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// Scale scales the rectangle by a given factor
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
