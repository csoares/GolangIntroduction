# Control Structures in Go

This example demonstrates how to use control structures in Go, including if statements, loops, and switches.

## What You'll Learn

- How to use if statements and if-else chains
- Different types of loops in Go
- How to use break and continue statements
- How to use the range keyword with loops
- How to use switch statements
- Type switching in Go

## Running the Example

To run this example, navigate to this directory in your terminal and run:

```
go run main.go
```

## Key Concepts

1. **If Statements**: Go's if statements are similar to other languages but don't require parentheses around the condition.

2. **For Loops**: Go has only one looping construct, the `for` loop, but it can be used in multiple ways:
   - Traditional three-component loop
   - As a while loop
   - As an infinite loop with break
   - With the range keyword for collections

3. **Switch Statements**: Go's switch is more flexible than in many other languages:
   - Cases don't fall through by default (no need for break)
   - Multiple values can be specified in a single case
   - Can be used without an expression (like an if-else chain)
   - The `fallthrough` keyword can be used to explicitly fall through to the next case

4. **Type Switch**: A special kind of switch used with interfaces to determine the type of a value.

## Try It Yourself

1. Modify the if statement example to check different conditions
2. Create a loop that counts down from 10 to 1
3. Write a switch statement that checks the day of the week and outputs a custom message for each day