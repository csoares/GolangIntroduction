# File I/O in Go

This section demonstrates how to work with files in Go, covering both reading and writing operations.

## Topics Covered

1. Basic File Operations
   - Opening and closing files
   - Reading entire files
   - Writing to files
   - Using defer for cleanup

2. Advanced File Operations
   - Buffered reading and writing
   - Working with CSV files
   - File paths and manipulation
   - Reading configuration files

## Examples

### 1. Basic File Operations (main.go)
- Reading text files
- Writing text files
- Proper error handling
- Using defer statements

### 2. CSV Processing (csv_example.go)
- Reading CSV data
- Writing CSV data
- Data transformation

### 3. Configuration Files (config_example.go)
- Reading JSON configuration
- Parsing configuration data
- Applying configuration

## Best Practices

1. Always close files using defer
2. Check for errors when opening files
3. Use appropriate buffer sizes for large files
4. Handle file paths correctly for cross-platform compatibility

## Exercises

1. Create a program that reads a text file and counts words
2. Implement a log file writer
3. Create a CSV file processor that performs data transformation

Try modifying the examples and complete the exercises to better understand file operations in Go!