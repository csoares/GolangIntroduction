# Error Handling in Go

This section demonstrates various error handling patterns and best practices in Go.

## Concepts Covered

1. Basic Error Handling
   - Using the `error` interface
   - Creating simple errors with `errors.New`
   - Error checking patterns

2. Custom Error Types
   - Creating custom error types
   - Implementing the `error` interface
   - Type assertions for error handling

3. Error Wrapping
   - Using `fmt.Errorf` with `%w` verb
   - Error unwrapping
   - Using `errors.Is` for error comparison

## Examples in `main.go`

1. **Division Function**: Demonstrates basic error creation and handling
   - Returns an error for division by zero
   - Shows how to check and handle errors

2. **Username Validation**: Shows custom error types
   - Custom `ValidationError` type
   - Type assertion for detailed error handling
   - Multiple validation rules

3. **Config Reading**: Illustrates error wrapping
   - Wraps file operation errors
   - Shows how to check for specific error types
   - Demonstrates error chain inspection

## Running the Examples

To run the examples:

```bash
go run main.go
```

You'll see different error scenarios and how they're handled in Go.

## Best Practices

1. Always check returned errors
2. Use custom error types for domain-specific errors
3. Wrap errors to add context while preserving the original error
4. Use type assertions and error unwrapping when needed
5. Keep error messages clear and actionable