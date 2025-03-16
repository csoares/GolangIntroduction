# Maps in Go

This section covers maps in Go, which are powerful built-in data structures that associate keys with values (key-value pairs).

## What You'll Learn

- Map declaration and initialization
- Adding and removing elements
- Accessing map elements
- Checking for key existence
- Iterating through maps
- Common map operations and patterns

## Examples in the Code

The code examples demonstrate:

1. Basic map operations
2. Student grade tracker implementation
3. Word frequency counter
4. Map with custom types

## Key Concepts

### Map Declaration
```go
// Method 1: Using make
scores := make(map[string]int)

// Method 2: Map literal
grades := map[string]int{"Alice": 95, "Bob": 87}
```

### Key Operations

1. Adding/Updating elements:
```go
scores["John"] = 85
```

2. Deleting elements:
```go
delete(scores, "John")
```

3. Checking existence:
```go
score, exists := scores["John"]
if exists {
    fmt.Println("John's score:", score)
}
```

### Important Notes

- Maps are reference types
- Keys must be comparable types
- Map operations are not guaranteed to happen in any specific order
- Maps are not safe for concurrent use (use sync.Map for concurrent access)

## Practice Exercises

1. Create a phone book application using maps
2. Implement a simple cache using maps
3. Count character frequencies in a string

Try modifying the example code to experiment with different map operations and use cases!