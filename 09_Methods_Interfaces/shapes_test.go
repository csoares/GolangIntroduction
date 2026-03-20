package main

import (
	"math"
	"testing"
)

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
			c := Circle{radius: tt.radius}
			got := c.Area()
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Circle.Area() = %v; want %v", got, tt.expected)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
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
			c := Circle{radius: tt.radius}
			got := c.Perimeter()
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Circle.Perimeter() = %v; want %v", got, tt.expected)
			}
		})
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
			r := Rectangle{width: tt.width, height: tt.height}
			got := r.Area()
			if got != tt.expected {
				t.Errorf("Rectangle.Area() = %v; want %v", got, tt.expected)
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
			r := Rectangle{width: tt.width, height: tt.height}
			got := r.Perimeter()
			if got != tt.expected {
				t.Errorf("Rectangle.Perimeter() = %v; want %v", got, tt.expected)
			}
		})
	}
}

// Test the Shape interface
func TestShapeInterface(t *testing.T) {
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
	}

	for i, shape := range shapes {
		area := shape.Area()
		perim := shape.Perimeter()

		if area <= 0 {
			t.Errorf("Shape %d: Area() = %v; want > 0", i, area)
		}
		if perim <= 0 {
			t.Errorf("Shape %d: Perimeter() = %v; want > 0", i, perim)
		}
	}
}
