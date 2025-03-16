# Packages and Modules in Go

This section demonstrates how to organize code using Go packages and modules, covering essential concepts of Go's package system.

## Key Concepts Covered

1. Package Organization
   - Creating and using custom packages
   - Package naming conventions
   - Package-level variables and functions

2. Module Management
   - Module initialization with go.mod
   - Managing dependencies
   - Version control in Go modules

3. Project Structure
   - Standard project layout
   - Internal vs external packages
   - Visibility rules (exported vs unexported)

## Project Structure

```
10_Packages_Modules/
├── calculator/           # Package demonstrating basic arithmetic operations
│   ├── basic.go         # Basic arithmetic operations
│   └── advanced.go      # Advanced mathematical operations
├── geometry/            # Package for geometric calculations
│   ├── circle.go        # Circle-related calculations
│   └── rectangle.go     # Rectangle-related calculations
├── main.go              # Main application demonstrating package usage
└── go.mod              # Module definition and dependency management
```

## Running the Example

1. Initialize the module:
   ```bash
   go mod init calculator-app
   ```

2. Run the main program:
   ```bash
   go run main.go
   ```

## Key Takeaways

- Packages help organize and modularize code
- Modules manage dependencies and versioning
- Proper naming and organization improve code maintainability
- Understanding visibility rules is crucial for package design