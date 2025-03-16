// Package calculator provides basic arithmetic operations
package calculator

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
// If b is 0, it returns 0 to avoid division by zero
func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
