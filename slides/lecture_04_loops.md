# Lecture 4: Loops - Doing Things Over and Over 🔄
## For Loops and Iteration

---

## Slide 1: Why Loops?

**Sometimes we need to repeat actions:**
- Count from 1 to 10 📝
- Process each item in a list 📋
- Keep asking until you get the right answer ❓

**Without loops:**
```go
fmt.Println(1)
fmt.Println(2)
fmt.Println(3)
// ... this is boring and long!
```

**With loops:**
```go
for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
// Much better!
```

---

## Slide 2: The For Loop

**Go has only one loop: `for`** (but it's powerful!)

**Basic structure:**
```go
for initialization; condition; post {
    // Code to repeat
}
```

**Example - Count to 5:**
```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
// Output: 1, 2, 3, 4, 5
```

**The three parts:**
1. `i := 1` - Start at 1
2. `i <= 5` - Continue while true
3. `i++` - Add 1 after each loop

---

## Slide 3: How For Loops Work

**Step by step:**

```go
for i := 1; i <= 3; i++ {
    fmt.Println(i)
}
```

1. **Initialize:** `i` becomes 1
2. **Check:** Is 1 <= 3? YES → Run code
3. **Print:** 1
4. **Increment:** `i` becomes 2
5. **Check:** Is 2 <= 3? YES → Run code
6. **Print:** 2
7. **Increment:** `i` becomes 3
8. **Check:** Is 3 <= 3? YES → Run code
9. **Print:** 3
10. **Increment:** `i` becomes 4
11. **Check:** Is 4 <= 3? NO → Exit loop

---

## Slide 4: Loop Variations

**Countdown:**
```go
for i := 10; i >= 1; i-- {
    fmt.Println(i)
}
// Output: 10, 9, 8, ... 1
```

**Count by 2s:**
```go
for i := 0; i <= 10; i += 2 {
    fmt.Println(i)
}
// Output: 0, 2, 4, 6, 8, 10
```

**Custom step:**
```go
for i := 0; i < 100; i += 10 {
    fmt.Println(i)
}
// Output: 0, 10, 20, ... 90
```

---

## Slide 5: While-Style Loop

**Go doesn't have `while`, but you can do this:**

```go
// While condition is true
for condition {
    // Do something
}
```

**Example:**
```go
number := 1

for number <= 5 {
    fmt.Println(number)
    number++
}
```

**Just the condition, no initialization!**

---

## Slide 6: Infinite Loops

**A loop that never stops:**

```go
for {
    fmt.Println("Forever!")
}
```

**Use with `break` to exit:**
```go
number := 1

for {
    if number > 5 {
        break  // Exit the loop
    }
    fmt.Println(number)
    number++
}
```

**Useful when you don't know how many iterations!**

---

## Slide 7: Break Statement

**Exit the loop immediately:**

```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break  // Stop when i is 5
    }
    fmt.Println(i)
}
// Output: 1, 2, 3, 4
```

**Real example - Find a number:**
```go
target := 7
for i := 1; i <= 100; i++ {
    if i == target {
        fmt.Println("Found it!")
        break
    }
}
```

---

## Slide 8: Continue Statement

**Skip the rest of this iteration:**

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue  // Skip printing 3
    }
    fmt.Println(i)
}
// Output: 1, 2, 4, 5
// (3 is skipped!)
```

**Print only even numbers:**
```go
for i := 1; i <= 10; i++ {
    if i % 2 != 0 {
        continue  // Skip odd numbers
    }
    fmt.Println(i)
}
// Output: 2, 4, 6, 8, 10
```

---

## Slide 9: Nested Loops

**A loop inside a loop:**

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

**Output:**
```
i=1, j=1
i=1, j=2
i=1, j=3
i=2, j=1
...
```

**Used for:**
- Tables and grids
- Processing 2D data
- Combinations

---

## Slide 10: Multiplication Table

```go
// Create a 5x5 multiplication table
for i := 1; i <= 5; i++ {
    for j := 1; j <= 5; j++ {
        fmt.Printf("%d\t", i * j)
    }
    fmt.Println()  // New line after each row
}
```

**Output:**
```
1    2    3    4    5
2    4    6    8    10
3    6    9    12   15
...
```

---

## Slide 11: Range - Looping Over Collections

**Loop over strings, arrays, slices, maps:**

```go
word := "Hello"

for index, letter := range word {
    fmt.Printf("Index %d: %c\n", index, letter)
}
```

**Output:**
```
Index 0: H
Index 1: e
Index 2: l
Index 3: l
Index 4: o
```

**`range` gives you:**
- Index (position)
- Value (at that position)

---

## Slide 12: Range with Arrays

```go
numbers := []int{10, 20, 30, 40, 50}

for index, value := range numbers {
    fmt.Printf("numbers[%d] = %d\n", index, value)
}
```

**If you only need values:**
```go
for _, value := range numbers {
    fmt.Println(value)
}
// Use _ to ignore the index
```

**If you only need indices:**
```go
for index := range numbers {
    fmt.Println(index)
}
```

---

## Slide 13: Summing Numbers

**Classic loop example:**

```go
// Sum of 1 to 100
sum := 0

for i := 1; i <= 100; i++ {
    sum += i  // Add i to sum
}

fmt.Println("Sum:", sum)  // 5050
```

**Factorial:**
```go
n := 5
factorial := 1

for i := 2; i <= n; i++ {
    factorial *= i
}

fmt.Println(n, "! =", factorial)  // 120
```

---

## Slide 14: Finding Max/Min

```go
numbers := []int{45, 12, 89, 33, 67, 5}

max := numbers[0]
min := numbers[0]

for _, num := range numbers[1:] {
    if num > max {
        max = num
    }
    if num < min {
        min = num
    }
}

fmt.Println("Max:", max)  // 89
fmt.Println("Min:", min)  // 5
```

---

## Slide 15: Practice Exercise 1

**FizzBuzz - Classic programming problem:**

Print numbers 1 to 100, but:
- If divisible by 3, print "Fizz" instead
- If divisible by 5, print "Buzz" instead
- If divisible by both, print "FizzBuzz"

```go
for i := 1; i <= 100; i++ {
    // Your code here
}
```

**Expected output:**
```
1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, ...
```

---

## Slide 16: Practice Exercise 2

**Print a pyramid:**

```
    *
   ***
  *****
 *******
*********
```

**Hints:**
- Use nested loops
- Outer loop for rows
- Inner loops for spaces and stars

```go
height := 5
for i := 1; i <= height; i++ {
    // Print spaces
    // Print stars
    // New line
}
```

---

## Slide 17: Common Mistakes

**Mistake 1: Infinite loop**
```go
for i := 0; i < 10; i-- {  // ❌ i decreases!
    // Never ends (or goes negative forever)
}

for i := 0; i < 10; i++ {   // ✅ i increases
```

**Mistake 2: Off-by-one error**
```go
for i := 0; i <= 10; i++ {  // Runs 11 times (0-10)
for i := 0; i < 10; i++ {   // Runs 10 times (0-9)
```

**Mistake 3: Modifying loop variable**
```go
for i := 0; i < 10; i++ {
    i += 2  // ❌ Messes up the loop!
}
```

---

## Slide 18: Loop Best Practices

**Use meaningful variable names:**
```go
// Not great
for i := 0; i < len(students); i++ { }

// Better
for studentIndex := 0; studentIndex < len(students); studentIndex++ { }

// Best (with range)
for _, student := range students { }
```

**Keep loops short:**
- If it's getting long, extract code to a function

---

## Slide 19: Real-World Example - Password Generator

```go
import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyz"
password := ""
length := 10

for i := 0; i < length; i++ {
    randomIndex := rand.Intn(len(letters))
    password += string(letters[randomIndex])
}

fmt.Println("Password:", password)
```

---

## Slide 20: Exercise Solutions

**FizzBuzz:**
```go
for i := 1; i <= 100; i++ {
    if i % 3 == 0 && i % 5 == 0 {
        fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
        fmt.Println("Fizz")
    } else if i % 5 == 0 {
        fmt.Println("Buzz")
    } else {
        fmt.Println(i)
    }
}
```

**Important:** Check divisibility by both FIRST!

---

## Slide 21: Exercise Solutions (cont.)

**Pyramid:**
```go
height := 5

for i := 1; i <= height; i++ {
    // Print spaces
    for j := 1; j <= height-i; j++ {
        fmt.Print(" ")
    }
    // Print stars
    for k := 1; k <= 2*i-1; k++ {
        fmt.Print("*")
    }
    fmt.Println()
}
```

---

## Slide 22: Advanced - Labels

**Breaking out of nested loops:**

```go
outer:
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            if i*j > 50 {
                break outer  // Break out of BOTH loops
            }
        }
    }
```

**Rarely needed, but useful!**

---

## Slide 23: Quiz Time!

**Q1: How many times does this run?**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
A) 4 times  B) 5 times  C) 6 times

**Q2: What's the output?**
```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Print(i)
}
```

**Q3: Convert to while-style:**
```go
for i := 0; i < 10; i++ { }
```

---

## Slide 24: Key Takeaways

**What you learned:**

1. ✅ `for init; condition; post` - Classic loop
2. ✅ `for condition` - While-style loop
3. ✅ `for` - Infinite loop
4. ✅ `break` - Exit loop early
5. ✅ `continue` - Skip to next iteration
6. ✅ `range` - Loop over collections
7. ✅ Nested loops for 2D data

---

## Slide 25: When to Use What

| Situation | Solution |
|-----------|----------|
| Count X times | `for i := 0; i < X; i++` |
| Until condition | `for condition { }` |
| Don't know when to stop | `for { if cond break }` |
| Array/slice elements | `for _, v := range` |
| Array/slice indices | `for i := range` |
| Both index and value | `for i, v := range` |

---

## Slide 26: Common Patterns

**Validate input:**
```go
for {
    fmt.Print("Enter a positive number: ")
    // read input
    if number > 0 {
        break
    }
    fmt.Println("Must be positive!")
}
```

**Retry logic:**
```go
for attempts := 3; attempts > 0; attempts-- {
    if tryConnect() {
        break
    }
}
```

---

## Slide 27: Homework

**Create a "Number Guessing Game":**

1. Generate a random number (1-100)
2. Ask the user to guess
3. Say "Too high!" or "Too low!"
4. Count number of attempts
5. After correct guess, ask to play again

**Features:**
- Limit to 10 attempts
- Track best score (lowest attempts)
- Validate input (1-100 only)

---

## Slide 28: What's Next?

**Lecture 5: Functions**
- Reusable blocks of code
- Parameters and return values
- Named return values
- Variadic functions
- Closures

**Your code will become organized and reusable! 🔧**

**See you next time! 👋**
