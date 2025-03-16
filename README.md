# Golang Introduction Course

Welcome to the Golang Introduction Course! This repository contains a series of examples designed to help you learn Go programming language from the ground up.

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
9. **09_Methods_Interfaces** - Object-oriented programming in Go
10. **10_Packages_Modules** - Organizing code with packages and modules
11. **11_ErrorHandling** - Handling errors effectively
12. **12_FileIO** - Reading and writing files
13. **13_Concurrency** - Introduction to goroutines
14. **14_Channels** - Communication between goroutines
15. **15_AdvancedConcurrency** - Advanced patterns with goroutines and channels

## How to Use This Course

1. Follow the folders in numerical order
2. Read the comments in each file to understand the concepts
3. Try to modify the examples to experiment with the concepts
4. Complete the exercises in each section if provided

## Prerequisites

- Basic programming knowledge
- Go installed on your machine (visit [golang.org](https://golang.org/dl/) to download)

## Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)

## Using the Makefile

This repository includes a Makefile to help you run examples easily. Here are the available commands:

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

Happy coding!