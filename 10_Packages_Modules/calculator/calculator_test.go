package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"with zero", 5, 0, 5},
		{"decimals", 2.5, 3.5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -2, -3, 1},
		{"mixed signs", -2, 3, -5},
		{"with zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 6},
		{"negative numbers", -2, -3, 6},
		{"mixed signs", -2, 3, -6},
		{"with zero", 5, 0, 0},
		{"with one", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 6, 3, 2},
		{"negative numbers", -6, -3, 2},
		{"mixed signs", -6, 3, -2},
		{"divide by zero", 5, 0, 0},
		{"decimals", 5, 2, 2.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Divide(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive exponent", 2, 3, 8},
		{"zero exponent", 5, 0, 1},
		{"negative exponent", 2, -1, 0.5},
		{"fractional base", 4, 0.5, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Power(tt.a, tt.b)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("Power(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"perfect square", 16, 4},
		{"zero", 0, 0},
		{"one", 1, 1},
		{"decimal", 2, math.Sqrt(2)},
		{"negative input", -1, 0}, // Current behavior returns 0 for negative
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquareRoot(tt.input)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("SquareRoot(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestAbsolute(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Absolute(tt.input)
			if got != tt.expected {
				t.Errorf("Absolute(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"round up", 2.6, 3},
		{"round down", 2.4, 2},
		{"round half up", 2.5, 3},
		{"negative", -2.5, -3},
		{"whole number", 2.0, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Round(tt.input)
			if got != tt.expected {
				t.Errorf("Round(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// Benchmark tests

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(123.456, 789.012)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(123.456, 789.012)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(2.5, 10)
	}
}
