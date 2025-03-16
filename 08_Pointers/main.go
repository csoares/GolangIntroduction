package main

import (
	"fmt"
)

// Person demonstrates struct with pointer receiver methods
type Person struct {
	Name string
	Age  int
}

// UpdateAge demonstrates a pointer receiver method
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge // Directly modifies the struct's field
}

// DisplayInfo demonstrates a value receiver method
func (p Person) DisplayInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// UpdateValue demonstrates passing by value vs reference
func UpdateValue(val *int) {
	*val = 100 // Modifies the value at the pointer's address
}

// SwapValues demonstrates using pointers to swap values
func SwapValues(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	// Basic pointer declaration and dereferencing
	fmt.Println("1. Basic Pointer Operations:")
	x := 42
	ptr := &x // & operator takes the address of x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of ptr: %p\n", ptr)
	fmt.Printf("Value at ptr's address: %d\n\n", *ptr) // * operator dereferences ptr

	// Modifying value through pointer
	*ptr = 100
	fmt.Printf("Modified x through ptr: %d\n\n", x)

	// Pointer to struct
	fmt.Println("2. Struct Pointers:")
	person := &Person{Name: "Alice", Age: 25}
	fmt.Printf("Initial person: %+v\n", *person)

	// Using pointer receiver method
	person.UpdateAge(26) // Go automatically dereferences
	fmt.Printf("After age update: %+v\n\n", *person)

	// Passing by reference example
	fmt.Println("3. Passing by Reference:")
	num := 42
	fmt.Printf("Before update: %d\n", num)
	UpdateValue(&num)
	fmt.Printf("After update: %d\n\n", num)

	// Swapping values using pointers
	fmt.Println("4. Swapping Values:")
	a, b := 10, 20
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	SwapValues(&a, &b)
	fmt.Printf("After swap: a = %d, b = %d\n\n", a, b)

	// Nil pointer safety
	fmt.Println("5. Nil Pointer Safety:")
	var nilPtr *int
	fmt.Printf("Nil pointer value: %v\n", nilPtr)

	// Go prevents dereferencing nil pointers at compile-time when possible
	// Uncommenting the next line would cause a panic:
	// fmt.Println(*nilPtr)

	// Array pointers
	fmt.Println("6. Array Pointers:")
	numbers := [3]int{1, 2, 3}
	arrPtr := &numbers

	// Accessing array elements through pointer
	fmt.Printf("First element: %d\n", (*arrPtr)[0])
	// Go also allows implicit dereferencing
	fmt.Printf("Second element: %d\n", arrPtr[1])
}
