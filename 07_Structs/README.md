# Structs in Go

This section covers structs in Go, which are composite data types that group together variables under a single name.

## What You'll Learn

- Struct declaration and initialization
- Field access and modification
- Embedded structs (composition)
- Methods with structs
- Struct tags and reflection
- Practical examples using structs

## Examples in This Section

### 1. Basic Struct Operations
- Declaring and initializing structs
- Accessing and modifying struct fields
- Comparing structs

### 2. Student Management System
- Creating a student struct
- Managing student records
- Calculating grades and averages

### 3. Geometry Calculator
- Implementing shapes using structs
- Calculating area and perimeter
- Demonstrating struct methods

## Key Concepts

### Struct Declaration
```go
type Person struct {
    Name    string
    Age     int
    Address string
}
```

### Struct Initialization
```go
// Method 1: Field names specified
person1 := Person{
    Name:    "John",
    Age:     25,
    Address: "123 Main St",
}

// Method 2: Field values in order
person2 := Person{"Jane", 30, "456 Oak Ave"}
```

### Embedded Structs
```go
type Employee struct {
    Person  // Embedded struct
    Company string
    Salary  float64
}
```

## Best Practices

1. Use meaningful field names
2. Keep structs focused and cohesive
3. Use embedding for logical composition
4. Implement methods for behavior related to the struct
5. Consider using interfaces for flexible design

## Exercises

1. Create a bank account struct with methods for deposit and withdrawal
2. Implement a library book management system using structs
3. Design a simple employee hierarchy using embedded structs

Refer to the code examples in `main.go` to see these concepts in action.