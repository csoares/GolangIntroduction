# Lecture 2: Variables and Types 📦
## Storing Information in Your Programs

---

## Slide 1: What is a Variable?

**Think of a variable like a labeled box:**
- 📦 You put something inside it
- 🏷️ It has a name (label)
- 🔍 You can look inside to see what's there
- 🔄 You can change what's inside

**In real life:**
- Box labeled "Age" contains "25"
- Box labeled "Name" contains "Alice"

**In Go:**
```go
var age = 25
var name = "Alice"
```

---

## Slide 2: Why Do We Need Variables?

**Without variables:**
```go
fmt.Println(2 + 3)        // What does this mean?
fmt.Println("Hello")      // Who are we greeting?
```

**With variables:**
```go
var score = 2 + 3
var greeting = "Hello"
fmt.Println(score)        // Clear! Shows a score
fmt.Println(greeting)     // Clear! Shows a greeting
```

**Variables make code readable and reusable!**

---

## Slide 3: Creating Variables - Three Ways

**Way 1: var keyword (explicit)**
```go
var age int = 25
var name string = "Bob"
```

**Way 2: var keyword (inferred)**
```go
var age = 25        // Go knows this is an int
var name = "Bob"    // Go knows this is a string
```

**Way 3: Short declaration (preferred)**
```go
age := 25           // Use := for new variables
name := "Bob"
```

---

## Slide 4: Understanding Types

**Types tell the computer what kind of data to expect:**

| Type | Description | Examples |
|------|-------------|----------|
| `int` | Whole numbers | 42, -7, 0 |
| `float64` | Decimal numbers | 3.14, -0.5 |
| `string` | Text | "Hello", "Go" |
| `bool` | True or false | true, false |

**Like different shaped boxes:**
- Round box for balls (integers)
- Rectangular box for books (strings)

---

## Slide 5: Numbers - int and float64

**Integers (int) - Whole numbers:**
```go
var apples int = 5
var temperature int = -10
var score int = 0

// Or use short declaration:
players := 11
year := 2024
```

**Floats (float64) - Decimal numbers:**
```go
var pi float64 = 3.14159
var price float64 = 19.99
var height float64 = 1.75

// Short declaration:
temperature := 98.6
```

---

## Slide 6: Strings - Text Data

**Strings hold text (words, sentences, symbols):**

```go
var greeting string = "Hello, World!"
var name string = "Alice"
var emoji string = "🎉"

// Short declaration:
message := "Welcome to Go!"
```

**String rules:**
- Must be in double quotes: "text"
- Can contain letters, numbers, spaces, symbols
- Single quotes are for something else (runes)

---

## Slide 7: Booleans - True or False

**Booleans are like light switches:**
- ON (true) or OFF (false)
- Only two possible values!

```go
var isReady bool = true
var isRaining bool = false

// Short declaration:
isHappy := true
hasFinished := false
```

**Used for:**
- Yes/No questions
- Feature flags
- Game states (game over?)

---

## Slide 8: Zero Values

**When you create a variable without a value, Go gives it a default:**

```go
var count int       // Starts as 0
var price float64   // Starts as 0.0
var name string     // Starts as "" (empty)
var active bool     // Starts as false
```

**Think of it like:**
- A new counter starts at 0
- A blank piece of paper is empty
- A switch starts in the OFF position

**Go never leaves variables uninitialized!**

---

## Slide 9: Multiple Variables

**Declare multiple variables at once:**

```go
// Same type
var x, y, z int = 1, 2, 3

// Different types (use var block)
var (
    name   string = "Alice"
    age    int    = 25
    isCool bool   = true
)

// Short declaration (multiple)
name, age := "Alice", 25
```

**Clean and organized!**

---

## Slide 10: Variable Naming Rules

**Rules (must follow):**
1. Start with a letter or underscore
2. Can contain letters, numbers, underscores
3. Case-sensitive (Age ≠ age)
4. Can't use Go keywords (var, func, etc.)

**Examples:**
```go
userName    // ✅ Good
user_name   // ✅ Good
_userName   // ✅ Good (but unusual)
1stPlace    // ❌ Can't start with number
var         // ❌ Reserved word
```

---

## Slide 11: Naming Conventions

**Go's style guide (best practices):**

**CamelCase for variables:**
```go
userName      // ✅ Good
user_name     // ❌ Not Go style
username      // ✅ Okay if one word
HTTPRequest   // ✅ Acronyms uppercase
```

**Descriptive names:**
```go
n             // ❌ Too short
count         // ⚠️ Okay
studentCount  // ✅ Clear!
```

**Be clear but concise!**

---

## Slide 12: Changing Variable Values

**Variables can change (unless they're constants):**

```go
score := 0
fmt.Println(score)  // Prints: 0

score = 100         // Change the value
fmt.Println(score)  // Prints: 100

score = 50          // Change again
fmt.Println(score)  // Prints: 50
```

**Notice:**
- `:=` for first assignment (declaration)
- `=` for re-assignment (updating)

---

## Slide 13: Constants - Variables That Never Change

**Constants are like your birthday:**
- Set once, never changes
- Use `const` instead of `var`

```go
const pi = 3.14159
const daysInWeek = 7
const appName = "MyApp"

// Multiple constants
const (
    Monday    = 1
    Tuesday   = 2
    Wednesday = 3
)
```

**Good for:**
- Mathematical values (pi, e)
- Configuration values
- Days of the week

---

## Slide 14: Math Operations

**Go can do math with numbers:**

```go
// Basic operations
sum := 5 + 3       // Addition: 8
diff := 10 - 4     // Subtraction: 6
product := 4 * 7   // Multiplication: 28
quotient := 20 / 4 // Division: 5

// With variables
a, b := 10, 3
result := a + b    // 13
```

**Order of operations:**
```go
result := 2 + 3 * 4  // 14 (not 20!)
// Multiplication happens first
```

---

## Slide 15: Division and Remainder

**Integer division:**
```go
result := 7 / 3    // Result: 2 (not 2.33!)
// Integers truncate (cut off) decimals
```

**Remainder (modulo):**
```go
remainder := 7 % 3  // Result: 1
// 7 divided by 3 is 2 with remainder 1
```

**Float division:**
```go
result := 7.0 / 3.0  // Result: 2.333...
// Use floats for decimal results
```

---

## Slide 16: String Operations

**Concatenation (joining strings):**
```go
firstName := "John"
lastName := "Doe"
fullName := firstName + " " + lastName
// Result: "John Doe"
```

**Can't mix types:**
```go
age := 25
message := "I am " + age + " years old"  // ❌ ERROR!

// Must convert number to string first
```

**String length:**
```go
name := "Hello"
length := len(name)  // Result: 5
```

---

## Slide 17: Type Conversion

**Go is strict about types - you must convert:**

```go
// Converting int to float64
age := 25
ageFloat := float64(age)  // 25.0

// Converting float64 to int
price := 19.99
priceInt := int(price)     // 19 (truncates!)

// Converting number to string (requires strconv package)
import "strconv"
age := 25
ageString := strconv.Itoa(age)  // "25"
```

---

## Slide 18: Printing Variables

**Different ways to print:**

```go
name := "Alice"
age := 25

// Simple print
fmt.Println("Name:", name, "Age:", age)
// Output: Name: Alice Age: 25

// Formatted print (Printf)
fmt.Printf("Name: %s, Age: %d\n", name, age)
// %s = string, %d = integer

// Shorthand with Println
fmt.Println("Name:", name)
```

---

## Slide 19: Format Verbs

**Special codes for Printf:**

| Verb | Type | Example |
|------|------|---------|
| %s | string | "Hello" |
| %d | integer | 42 |
| %f | float | 3.14 |
| %t | boolean | true |
| %v | any value | (generic) |
| %T | type of value | int |
| %% | literal % | % |

```go
fmt.Printf("Pi is %f\n", 3.14159)
fmt.Printf("I have %d apples\n", 5)
```

---

## Slide 20: Practice Exercise 1

**Create a program with:**
- Your name (string)
- Your age (int)
- Your height in meters (float64)
- Whether you like pizza (bool)

**Print them all:**
```
Name: [Your Name]
Age: [Your Age]
Height: [Your Height]m
Likes Pizza: [true/false]
```

**Try both Println and Printf!**

---

## Slide 21: Practice Exercise 2

**Calculate your age in different units:**

```go
ageInYears := 25
ageInMonths := ageInYears * 12
ageInDays := ageInYears * 365

fmt.Println("Age in years:", ageInYears)
fmt.Println("Age in months:", ageInMonths)
fmt.Println("Age in days:", ageInDays)
```

**Bonus:** Calculate age in hours!

---

## Slide 22: Common Mistakes

**Mistake 1: Using undeclared variables**
```go
fmt.Println(score)  // ❌ score doesn't exist!
score := 100        // ✅ Declare first
```

**Mistake 2: Wrong type assignment**
```go
var age int = "25"  // ❌ Can't put string in int!
var age int = 25    // ✅ Correct
```

**Mistake 3: Using := for existing variables**
```go
score := 100
score := 200        // ❌ Use = not :=
score = 200         // ✅ Correct
```

---

## Slide 23: Variable Scope

**Scope = Where a variable can be used:**

```go
package main

import "fmt"

var globalVar = "I exist everywhere"

func main() {
    var localVar = "I only exist in main()"
    fmt.Println(globalVar)   // ✅ Works
    fmt.Println(localVar)    // ✅ Works
}

// fmt.Println(localVar)     // ❌ Won't work here!
```

**Rule:** Variables only exist inside the `{}` braces where they're declared.

---

## Slide 24: Shadowing

**When a local variable "hides" a global one:**

```go
var name = "Global"

func main() {
    fmt.Println(name)        // Prints: Global

    name := "Local"          // Creates NEW variable
    fmt.Println(name)        // Prints: Local

    // The global 'name' is still there, just hidden!
}
```

**Be careful!** This can cause confusion. Use different names.

---

## Slide 25: Type Inference Deep Dive

**Go is smart about types:**

```go
// Go infers the smallest type that fits
x := 42           // int
y := 3.14         // float64
z := "Hello"      // string

// But sometimes you need to be specific
var bigNumber int64 = 10000000000
var smallNumber int8 = 100
```

**Integer types:** int8, int16, int32, int64, uint (unsigned)
**Float types:** float32, float64

**Usually just use `int` and `float64`**

---

## Slide 26: Real-World Example

**A simple shopping cart:**

```go
package main

import "fmt"

func main() {
    // Product information
    productName := "Super Widget"
    price := 29.99
    quantity := 3
    inStock := true

    // Calculate total
    total := price * float64(quantity)

    // Display receipt
    fmt.Println("=== RECEIPT ===")
    fmt.Println("Product:", productName)
    fmt.Printf("Price: $%.2f\n", price)
    fmt.Println("Quantity:", quantity)
    fmt.Printf("Total: $%.2f\n", total)
    fmt.Println("In Stock:", inStock)
}
```

---

## Slide 27: Debugging Tips

**When variables don't work:**

1. **Check the type:**
   ```go
   fmt.Printf("Type: %T, Value: %v\n", variable, variable)
   ```

2. **Print intermediate values:**
   ```go
   a := 5
   b := 10
   fmt.Println("a:", a, "b:", b)  // Check before using
   sum := a + b
   ```

3. **Use descriptive variable names**

4. **Initialize variables before using**

---

## Slide 28: Best Practices Summary

**Do:**
- ✅ Use descriptive names (`studentCount` not `n`)
- ✅ Use short declaration `:=` when possible
- ✅ Group related variables
- ✅ Use constants for values that never change
- ✅ Convert types explicitly

**Don't:**
- ❌ Use single-letter names (except in loops)
- ❌ Mix types without conversion
- ❌ Reuse variables for different purposes
- ❌ Use global variables unnecessarily

---

## Slide 29: Quick Quiz

**Q1: What's the output?**
```go
x := 10
y := 3
fmt.Println(x / y)
```
A) 3.33  B) 3  C) 3.0

**Q2: Which is correct?**
A) `var 1stPlace = "Gold"`
B) `var firstPlace = "Gold"`
C) `var first place = "Gold"`

**Q3: What type is `isReady := true`?**
A) string  B) int  C) bool

---

## Slide 30: What's Next?

**In Lecture 3, you'll learn:**
- Making decisions with if/else
- Comparing values
- Multiple conditions
- The switch statement

**Your programs will become smart! 🧠**

**Homework:**
Create a "Temperature Converter" that converts Celsius to Fahrenheit!

**Formula:** `F = (C × 9/5) + 32`

See you next time! 👋
