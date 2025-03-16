# Go Functions

This section covers various aspects of functions in Go programming language. Functions are fundamental building blocks in Go that help organize code into reusable and manageable pieces.

## Topics Covered

### 1. Basic Functions
- Simple function declaration and calling
- Functions without parameters or return values
```go
func sayHello() {
    fmt.Println("Hello from a function!")
}
```

### 2. Functions with Parameters
- Passing data to functions
- Parameter type declarations
```go
func add(a int, b int) {
    fmt.Println("Sum:", a+b)
}
```

### 3. Functions with Return Values
- Single return value
- Return type declaration
```go
func multiply(a int, b int) int {
    return a * b
}
```

### 4. Multiple Parameters of Same Type
- Shorthand parameter type declaration
```go
func subtract(a, b int) int {
    return a - b
}
```

### 5. Multiple Return Values
- Returning multiple values
- Error handling pattern
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### 6. Named Return Values
- Pre-declaring return variable names
- Naked return statement
```go
func calculateStats(numbers []int) (min, max, total int)
```

### 7. Variadic Functions
- Variable number of arguments
- Using ellipsis (...)
```go
func sum(numbers ...int) int
```

### 8. Functions as Values
- Functions as first-class citizens
- Function types
- Anonymous functions
```go
func getGreetingFunction() func(string) string
```

### 9. Recursive Functions
- Functions that call themselves
- Base case and recursive case
```go
func factorial(n int) int
```

### 10. Defer Statement
- Delaying execution until function returns
- LIFO (Last In, First Out) order for multiple defers
```go
func deferExample()
```

## Key Points

1. Functions in Go are declared using the `func` keyword
2. Parameters require explicit type declarations
3. Return types must be specified if the function returns a value
4. Multiple return values are supported and commonly used for error handling
5. Variadic functions can accept a variable number of arguments
6. Functions are first-class citizens and can be assigned to variables
7. Anonymous functions can be declared and used inline
8. Deferred function calls are executed in LIFO order

## Best Practices

1. Keep functions focused and single-purpose
2. Use meaningful function and parameter names
3. Document complex functions with comments
4. Return errors instead of using panic
5. Use named return values when they make the code clearer
6. Consider using defer for cleanup operations

## Examples

Check the `main.go` file in this directory for complete examples of each function type and their usage.