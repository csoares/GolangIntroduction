# Arrays and Slices in Go

This example demonstrates how to work with arrays and slices in Go, two fundamental data structures for storing collections of values.

## What You'll Learn

- How to declare and initialize arrays
- How to access and modify array elements
- The difference between arrays and slices
- How to create and manipulate slices
- Slicing operations
- How to use built-in functions like `make`, `len`, and `cap`

## Running the Example

To run this example, navigate to this directory in your terminal and run:

```
go run main.go
```

## Key Concepts

1. **Arrays**: Fixed-size collections of elements of the same type
   - Have a fixed length that cannot be changed after declaration
   - Are value types (copying an array creates a new copy of all elements)
   - Can be single or multi-dimensional

2. **Slices**: Dynamic, flexible views into arrays
   - Have a dynamic length that can grow or shrink
   - Are reference types (copying a slice creates a reference to the same underlying array)
   - Created using array slicing, slice literals, or the `make` function
   - Have both a length and a capacity

3. **Slice Operations**:
   - Slicing: `slice[start:end]`
   - Appending: `append(slice, elements...)`
   - Copying: `copy(destSlice, sourceSlice)`

## Try It Yourself

1. Create an array and a slice, then compare their behaviors when copied
2. Use the `append` function to add elements to a slice
3. Create a slice with `make` and experiment with different length and capacity values
4. Practice different slicing operations