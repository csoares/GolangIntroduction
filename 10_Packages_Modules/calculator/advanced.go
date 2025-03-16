// Package calculator provides advanced arithmetic operations
package calculator

import "math"

// Power returns a raised to the power of b
func Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// SquareRoot returns the square root of a number
// If a is negative, it returns 0
func SquareRoot(a float64) float64 {
	if a < 0 {
		return 0
	}
	return math.Sqrt(a)
}

// Absolute returns the absolute value of a number
func Absolute(a float64) float64 {
	return math.Abs(a)
}

// Round returns the nearest integer, rounding half away from zero
func Round(a float64) float64 {
	return math.Round(a)
}
