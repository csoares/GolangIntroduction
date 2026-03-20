# Golang Introduction Course

Welcome to the Golang Introduction Course! This repository contains a series of examples designed to help you learn Go programming language from the ground up.

## 🎓 HTML Slides Available!

**[📺 View Slides Online](https://csoares.github.io/GolangIntroduction/slides/)**

A complete HTML slide presentation system with 15 lectures is now available. Open the link above to view the course slides directly in your browser!

**Features:**
- 📱 Responsive design (works on mobile and desktop)
- ⌨️ Keyboard navigation (arrow keys)
- 🎨 Syntax highlighting
- 📖 ~30 slides per lecture
- 🎯 Kid-friendly explanations with examples

## Course Structure

The course is organized into numbered folders, each focusing on a specific concept. The numbering indicates the recommended learning path, starting from basic concepts and progressing to more advanced topics.

### Learning Path

1. **01_HelloWorld** - Your first Go program and basic syntax
2. **02_Variables** - Variables, constants, and data types
3. **03_ControlStructures** - If statements, loops, and switches
4. **04_Functions** - Function declarations, parameters, and return values
5. **05_Arrays_Slices** - Working with arrays and slices
6. **06_Maps** - Using maps for key-value data
7. **07_Structs** - Creating custom data types with structs
8. **08_Pointers** - Understanding pointers and memory management
9. **09_Methods_Interfaces** - Object-oriented programming in Go with interfaces
10. **10_Packages_Modules** - Organizing code with packages and modules (includes tests)
11. **11_ErrorHandling** - Handling errors effectively with custom error types
12. **12_FileIO** - Reading and writing files
13. **13_Concurrency** - Introduction to goroutines and WaitGroups
14. **14_Channels** - Communication between goroutines with pipelines
15. **15_AdvancedConcurrency** - Production-grade patterns: Fan-in/Fan-out, Rate Limiting, Context Cancellation, Semaphores

## Production-Grade Features

This codebase demonstrates production-grade Go practices:

### Concurrency Patterns
- **Context-aware operations**: All concurrent functions accept `context.Context` for cancellation and timeouts
- **Buffered channels**: Pipeline stages use buffered channels to prevent goroutine leaks
- **Proper synchronization**: WaitGroup patterns follow Go best practices
- **Race-free code**: Uses local `rand.Rand` instances instead of global `math/rand`

### Testing
- Comprehensive unit tests for `calculator` and `geometry` packages
- Table-driven tests with subtests for clarity
- Benchmark tests for performance-critical functions
- Run tests with: `go test ./...`

### Code Organization
- Build tags (`//go:build ignore`) on example mains prevent build conflicts
- Proper package documentation with godoc comments
- Interface-based design for extensibility

## Prerequisites

- Basic programming knowledge
- Go 1.22+ installed on your machine (visit [golang.org](https://golang.org/dl/) to download)

## Using the Makefile

This repository includes a Makefile to help you run examples easily:

### Basic Examples
```bash
make basic      # Run all basic examples
make hello      # Run Hello World example
make vars       # Run Variables example
make control    # Run Control Structures example
make funcs      # Run Functions example
make arrays     # Run Arrays and Slices example
make maps       # Run Maps example
make structs    # Run Structs example
make pointers   # Run Pointers example
make methods    # Run Methods and Interfaces example
```

### Intermediate Examples
```bash
make intermediate  # Run all intermediate examples
make packages      # Run Packages and Modules example
make errors        # Run Error Handling example
make fileio        # Run File I/O example
```

### Advanced Examples
```bash
make advanced    # Run all advanced examples
make concurrency # Run basic Concurrency example
make channels    # Run Channels example
make fanin       # Run Fan-in Fan-out example
make ratelimit   # Run Rate Limiting example
make context     # Run Context Cancellation example
make semaphores  # Run Semaphores example
```

### Common Operations
```bash
make all         # Run all examples
make fmt         # Format all Go files
make test        # Run all tests
make clean       # Clean build artifacts
```

## Running Tests

The `10_Packages_Modules` directory contains comprehensive tests:

```bash
cd 10_Packages_Modules
go test ./... -v          # Run all tests with verbose output
go test ./... -race       # Run tests with race detector
go test ./... -bench=.    # Run benchmarks
```

## Key Concepts Demonstrated

### Error Handling
- Custom error types implementing the `error` interface
- Error wrapping with `fmt.Errorf("%w")`
- Error checking with `errors.Is()`

### Interfaces
- Interface-based design for shapes (Area, Perimeter)
- Type assertions and type switches
- Implicit interface satisfaction

### Concurrency
- Goroutines and channels
- Buffered vs unbuffered channels
- Select statements for multiplexing
- Pipeline patterns
- Worker pools
- Rate limiting with token bucket
- Context for cancellation
- Semaphores for resource limiting

## Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

Happy coding!
