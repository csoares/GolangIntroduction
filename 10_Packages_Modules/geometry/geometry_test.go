package geometry

import (
	"math"
	"testing"
)

func TestNewCircle(t *testing.T) {
	c := NewCircle(5)
	if c.Radius != 5 {
		t.Errorf("NewCircle(5).Radius = %v; want 5", c.Radius)
	}
}

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected float64
	}{
		{"radius 1", 1, math.Pi},
		{"radius 5", 5, math.Pi * 25},
		{"radius 0", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCircle(tt.radius)
			got := c.Area()
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Circle.Area() with radius %v = %v; want %v", tt.radius, got, tt.expected)
			}
		})
	}
}

func TestCircle_Circumference(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected float64
	}{
		{"radius 1", 1, 2 * math.Pi},
		{"radius 5", 5, 2 * math.Pi * 5},
		{"radius 0", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCircle(tt.radius)
			got := c.Circumference()
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Circle.Circumference() with radius %v = %v; want %v", tt.radius, got, tt.expected)
			}
		})
	}
}

func TestCircle_Diameter(t *testing.T) {
	tests := []struct {
		name     string
		radius   float64
		expected float64
	}{
		{"radius 5", 5, 10},
		{"radius 0", 0, 0},
		{"radius 2.5", 2.5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCircle(tt.radius)
			got := c.Diameter()
			if got != tt.expected {
				t.Errorf("Circle.Diameter() with radius %v = %v; want %v", tt.radius, got, tt.expected)
			}
		})
	}
}

func TestNewRectangle(t *testing.T) {
	r := NewRectangle(4, 6)
	if r.Width != 4 || r.Height != 6 {
		t.Errorf("NewRectangle(4, 6) = {width: %v, height: %v}; want {width: 4, height: 6}", r.Width, r.Height)
	}
}

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name     string
		width    float64
		height   float64
		expected float64
	}{
		{"4x6", 4, 6, 24},
		{"square", 5, 5, 25},
		{"zero width", 0, 5, 0},
		{"zero height", 5, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRectangle(tt.width, tt.height)
			got := r.Area()
			if got != tt.expected {
				t.Errorf("Rectangle.Area() with width %v, height %v = %v; want %v", tt.width, tt.height, got, tt.expected)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name     string
		width    float64
		height   float64
		expected float64
	}{
		{"4x6", 4, 6, 20},
		{"square", 5, 5, 20},
		{"zero width", 0, 5, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRectangle(tt.width, tt.height)
			got := r.Perimeter()
			if got != tt.expected {
				t.Errorf("Rectangle.Perimeter() with width %v, height %v = %v; want %v", tt.width, tt.height, got, tt.expected)
			}
		})
	}
}

func TestRectangle_IsSquare(t *testing.T) {
	tests := []struct {
		name     string
		width    float64
		height   float64
		expected bool
	}{
		{"square", 5, 5, true},
		{"rectangle", 4, 6, false},
		{"almost square", 5, 5.0001, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRectangle(tt.width, tt.height)
			got := r.IsSquare()
			if got != tt.expected {
				t.Errorf("Rectangle.IsSquare() with width %v, height %v = %v; want %v", tt.width, tt.height, got, tt.expected)
			}
		})
	}
}

func TestRectangle_Scale(t *testing.T) {
	r := NewRectangle(4, 6)
	r.Scale(2)

	if r.Width != 8 || r.Height != 12 {
		t.Errorf("After Scale(2), Rectangle = {width: %v, height: %v}; want {width: 8, height: 12}", r.Width, r.Height)
	}

	// Test scale by zero
	r2 := NewRectangle(4, 6)
	r2.Scale(0)

	if r2.Width != 0 || r2.Height != 0 {
		t.Errorf("After Scale(0), Rectangle = {width: %v, height: %v}; want {width: 0, height: 0}", r2.Width, r2.Height)
	}
}

// Benchmark tests

func BenchmarkCircle_Area(b *testing.B) {
	c := NewCircle(5)
	for i := 0; i < b.N; i++ {
		c.Area()
	}
}

func BenchmarkRectangle_Area(b *testing.B) {
	r := NewRectangle(4, 6)
	for i := 0; i < b.N; i++ {
		r.Area()
	}
}
