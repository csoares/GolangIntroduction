# Methods and Interfaces in Go

This section covers object-oriented programming concepts in Go through methods and interfaces.

## Key Concepts

### Methods
- Methods are functions with a special receiver argument
- Can be defined on both value and pointer receivers
- Enable behavior similar to classes in other languages

### Interfaces
- Define behavior through method signatures
- Implemented implicitly (no "implements" keyword needed)
- Enable polymorphism and flexible code design

## Examples in This Section

1. **Shape Calculator**
   - Demonstrates interface implementation
   - Shows polymorphism in action
   - Includes multiple shapes (Circle, Rectangle)

2. **Custom Data Structure**
   - Implements multiple interfaces
   - Shows method receivers (value vs pointer)
   - Demonstrates type assertions

3. **Empty Interface Usage**
   - Working with interface{}
   - Type assertions and type switches
   - Practical use cases

## Best Practices
- Choose receiver type appropriately (value vs pointer)
- Keep interfaces small and focused
- Use composition over inheritance
- Leverage type assertions safely

## Practice Exercises
1. Add a new shape to the shape calculator
2. Implement a custom interface
3. Create a type that implements multiple interfaces