// Package main is the entry point for all Go programs
package main

// Import the fmt package for formatted I/O operations
import "fmt"

// The main function is the entry point of the program
func main() {
	// Variable Declaration and Initialization

	// Method 1: Declare a variable with explicit type
	var age int
	age = 25
	fmt.Println("Age:", age)

	// Method 2: Declare and initialize in one line
	var name string = "John Doe"
	fmt.Println("Name:", name)

	// Method 3: Type inference (Go infers the type from the value)
	var salary = 50000.50 // Go infers this as float64
	fmt.Println("Salary:", salary)

	// Method 4: Short declaration (most common in function bodies)
	city := "New York" // Equivalent to: var city string = "New York"
	fmt.Println("City:", city)

	// Multiple variable declaration
	var (
		isEmployed     bool = true
		yearsOfService int  = 3
	)
	fmt.Println("Employed:", isEmployed)
	fmt.Println("Years of Service:", yearsOfService)

	// Multiple short declarations in one line
	country, code := "USA", 1
	fmt.Println("Country:", country, "Code:", code)

	// Constants
	const Pi = 3.14159
	const (
		MaxConnections = 100
		AppName        = "GoApp"
	)
	fmt.Println("Pi:", Pi)
	fmt.Println("Max Connections:", MaxConnections)
	fmt.Println("App Name:", AppName)

	// Basic Data Types

	// Numeric Types
	var integerNum int = 42                  // Platform dependent size (32 or 64 bit)
	var int8Num int8 = 127                   // 8-bit integer (-128 to 127)
	var int16Num int16 = 32767               // 16-bit integer
	var int32Num int32 = 2147483647          // 32-bit integer
	var int64Num int64 = 9223372036854775807 // 64-bit integer

	var uintNum uint = 42    // Unsigned integer
	var uint8Num uint8 = 255 // 8-bit unsigned integer (0 to 255)

	var floatNum float32 = 3.14               // 32-bit floating point
	var doubleNum float64 = 3.141592653589793 // 64-bit floating point (double precision)

	var complexNum complex64 = 1 + 2i // Complex number with float32 real and imaginary parts

	fmt.Println("Integer Types:", integerNum, int8Num, int16Num, int32Num, int64Num)
	fmt.Println("Unsigned Integer Types:", uintNum, uint8Num)
	fmt.Println("Floating Point Types:", floatNum, doubleNum)
	fmt.Println("Complex Number:", complexNum)

	// Boolean Type
	var isTrue bool = true
	var isFalse bool = false
	fmt.Println("Boolean Values:", isTrue, isFalse)

	// String Type
	var message string = "Hello, Go!"
	fmt.Println("String:", message)

	// Raw string literal (preserves line breaks and escape characters)
	var multiLineStr = `This is a multi-line
string in Go.
Escape characters like \n are not interpreted.`
	fmt.Println("Multi-line String:")
	fmt.Println(multiLineStr)

	// Zero values (default values when variables are declared but not initialized)
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string
	fmt.Println("Zero Values:", defaultInt, defaultFloat, defaultBool, "'"+defaultString+"'")

	// Type conversion (explicit conversion is required in Go)
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println("Type Conversion:", i, f, u)

	// Type aliases
	type Celsius float64
	type Fahrenheit float64

	var tempC Celsius = 25
	var tempF Fahrenheit = 77
	fmt.Println("Temperature in Celsius:", tempC)
	fmt.Println("Temperature in Fahrenheit:", tempF)
}
