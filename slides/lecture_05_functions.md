# Lecture 5: Functions - Reusable Code Blocks 🔧
## Building Blocks of Programming

---

## Slide 1: Why Functions?

**Imagine you're a chef:**
- You don't explain how to chop onions every time
- You just say "chop the onions" and everyone knows what to do

**In programming:**
- Write code once, use it many times
- Organize code into logical chunks
- Make code easier to read and understand
- Easier to find and fix bugs

**Without functions:**
```go
// Copy-paste code everywhere
// Hard to maintain!
```

**With functions:**
```go
greet()  // Simple and clear!
```

---

## Slide 2: What is a Function?

**A function is like a recipe:**
1. Has a name (what to call it)
2. Takes ingredients (parameters)
3. Does work (function body)
4. Returns a result (return value)

**Analogy:**
```
Function: Bake a Cake
- Ingredients: flour, sugar, eggs
- Process: mix, bake, cool
- Result: delicious cake!
```

**In Go:**
```go
func add(a int, b int) int {
    return a + b
}
```

---

## Slide 3: Function Structure

**Basic function:**
```go
func functionName() {
    // Code to execute
}
```

**With parameters:**
```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

**With return value:**
```go
func square(x int) int {
    return x * x
}
```

**Parts:**
- `func` - Keyword to start function
- `functionName` - What you call it
- `(parameters)` - Input values
- `return type` - What it gives back
- `{}` - Function body

---

## Slide 4: Calling Functions

**Using your functions:**

```go
package main

import "fmt"

// Define the function
func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    // Call the function
    greet("Alice")   // Output: Hello, Alice
    greet("Bob")     // Output: Hello, Bob
    greet("Charlie") // Output: Hello, Charlie
}
```

**Order matters:**
- Define functions before using them (or anywhere in package)
- Call from main() to execute

---

## Slide 5: Parameters - Input Data

**Parameters are like blanks to fill in:**

```go
func add(a int, b int) int {
    return a + b
}

// Can shorten if same type
func add(a, b int) int {
    return a + b
}
```

**Multiple different types:**
```go
func describe(name string, age int) {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Call it
describe("Alice", 25)
```

**Order matters!**

---

## Slide 6: Return Values - Output

**Functions can give back results:**

```go
func multiply(a, b int) int {
    result := a * b
    return result
}

func main() {
    answer := multiply(5, 3)
    fmt.Println(answer)  // 15
}
```

**Return exits the function:**
```go
func checkAge(age int) string {
    if age < 0 {
        return "Invalid age"  // Exit early!
    }
    return "Valid age"
}
```

---

## Slide 7: Multiple Return Values

**Go functions can return multiple things!**

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

**Very common pattern in Go!**

---

## Slide 8: Named Return Values

**Give return values names:**

```go
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // Naked return!
}

func main() {
    a, p := rectangle(5, 3)
    fmt.Println("Area:", a, "Perimeter:", p)
}
```

**Benefits:**
- Self-documenting
- Can use naked return
- Clear what function returns

---

## Slide 9: Variadic Functions

**Functions with unlimited parameters:**

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))     // 6
    fmt.Println(sum(1, 2))        // 3
    fmt.Println(sum())            // 0
}
```

**The `...` means "zero or more"**

---

## Slide 10: Passing Slices to Variadic

**Already have a slice?**

```go
func sum(numbers ...int) int {
    // ... function body
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    // Unpack slice with ...
    result := sum(nums...)
    fmt.Println(result)  // 15
}
```

**The `...` operator unpacks the slice!**

---

## Slide 11: Function Scope

**Variables only exist inside their function:**

```go
func functionA() {
    x := 10  // Only exists here
    fmt.Println(x)
}

func functionB() {
    // x doesn't exist here!
    y := 20  // Only exists here
    fmt.Println(y)
}
```

**Each function is isolated:**
- Can't access other function's variables
- Prevents accidents
- Makes code modular

---

## Slide 12: Recursion

**Function calling itself:**

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// factorial(5)
// = 5 * factorial(4)
// = 5 * 4 * factorial(3)
// = 5 * 4 * 3 * factorial(2)
// = 5 * 4 * 3 * 2 * factorial(1)
// = 5 * 4 * 3 * 2 * 1 = 120
```

**Must have a base case!** (stopping condition)

---

## Slide 13: Anonymous Functions

**Functions without names:**

```go
func main() {
    // Define and call immediately
    result := func(a, b int) int {
        return a + b
    }(3, 4)

    fmt.Println(result)  // 7
}
```

**Assign to variable:**
```go
add := func(a, b int) int {
    return a + b
}

fmt.Println(add(2, 3))  // 5
```

**Also called: function literals**

---

## Slide 14: Closures

**Functions that remember their environment:**

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter := makeCounter()

    fmt.Println(counter())  // 1
    fmt.Println(counter())  // 2
    fmt.Println(counter())  // 3
}
```

**The inner function "closes over" the count variable!**

---

## Slide 15: Deferred Functions

**Delay execution until function returns:**

```go
func readFile() {
    file := openFile("data.txt")
    defer file.Close()  // Will run at end!

    // Read file...
    // Process data...
    // Don't worry about closing!
}
```

**Multiple defers - LIFO (Last In First Out):**
```go
func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    fmt.Println("Hello")
}
// Output: Hello, Third, Second, First
```

---

## Slide 16: Practical Example - Calculator

```go
package main

import "fmt"

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    x, y := 10.0, 5.0

    fmt.Println("Add:", add(x, y))
    fmt.Println("Subtract:", subtract(x, y))
    fmt.Println("Multiply:", multiply(x, y))

    if result, err := divide(x, y); err == nil {
        fmt.Println("Divide:", result)
    }
}
```

---

## Slide 17: Best Practices

**Function naming:**
- Use verbs: `calculateTotal`, `getUserName`, `isValid`
- Be descriptive but concise
- Start with lowercase for private, uppercase for public

**Function size:**
- Keep functions small (under 20 lines if possible)
- Do ONE thing well
- If you need to scroll, it's too long!

**Parameters:**
- Fewer is better (aim for 0-3)
- Group related parameters into structs

---

## Slide 18: Common Mistakes

**Mistake 1: Forgetting return**
```go
func add(a, b int) int {
    a + b  // ❌ No return!
}

func add(a, b int) int {
    return a + b  // ✅ Correct
}
```

**Mistake 2: Wrong return type**
```go
func getName() int {       // ❌ Should be string
    return "Alice"
}
```

**Mistake 3: Ignoring return values**
```go
result, _ := divide(10, 0)  // ⚠️ Ignoring error!
```

---

## Slide 19: Higher-Order Functions

**Functions that take or return functions:**

```go
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    multiply := func(x, y int) int { return x * y }

    fmt.Println(applyOperation(3, 4, add))      // 7
    fmt.Println(applyOperation(3, 4, multiply)) // 12
}
```

**Powerful for callbacks and customization!**

---

## Slide 20: Exercise 1 - Palindrome

**Write a function that checks if a string is a palindrome:**

```go
func isPalindrome(s string) bool {
    // Your code here
}

// Examples:
// "radar" -> true
// "hello" -> false
// "A man a plan a canal Panama" -> true (ignore spaces)
```

**Hints:**
- Compare first and last character
- Move towards middle
- Use a loop

---

## Slide 21: Exercise 2 - Fibonacci

**Write a recursive Fibonacci function:**

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21...
// fibonacci(6) = 8
```

**Bonus:**
- Write an iterative version
- Which is faster?

---

## Slide 22: Solution - Palindrome

```go
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1

    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

// With case and space insensitivity (advanced):
import "strings"

func isPalindromeAdvanced(s string) bool {
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "")
    return isPalindrome(s)
}
```

---

## Slide 23: Method vs Function

**Function:**
- Independent
- Called directly: `functionName()`

**Method:** (preview for next lecture)
- Attached to a type
- Called on an instance: `instance.method()`

```go
// Function
func area(radius float64) float64 {
    return math.Pi * radius * radius
}

// Method (coming next!)
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

---

## Slide 24: Quiz Time!

**Q1: What's the output?**
```go
func mystery(x int) int {
    if x <= 1 {
        return x
    }
    return mystery(x-1) + mystery(x-2)
}
fmt.Println(mystery(5))
```

**Q2: Fix the error:**
```go
func divide(a, b int) {
    return a / b
}
```

**Q3: How many values can a Go function return?**
A) Only 1
B) Up to 2
C) Any number

---

## Slide 25: Key Takeaways

**What you learned:**

1. ✅ Functions make code reusable
2. ✅ Parameters = input, Return = output
3. ✅ Can return multiple values
4. ✅ Named returns are self-documenting
5. ✅ Variadic functions accept any number of args
6. ✅ Anonymous functions and closures
7. ✅ defer delays execution

---

## Slide 26: Function Checklist

**When writing a function:**

- [ ] Does it have a clear, descriptive name?
- [ ] Does it do ONE thing?
- [ ] Are parameters minimal and clear?
- [ ] Are return values handled by caller?
- [ ] Is it under 20 lines?
- [ ] Is it tested?

---

## Slide 27: Real-World Example

**HTTP Handler function:**
```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")
    if userID == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    user, err := getUserFromDB(userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}
```

---

## Slide 28: Homework

**Build a "Math Toolkit":**

Create functions for:
1. `isPrime(n int) bool` - Check if number is prime
2. `gcd(a, b int) int` - Greatest common divisor
3. `lcm(a, b int) int` - Least common multiple
4. `power(base, exponent int) int` - Power function
5. `sumOfDigits(n int) int` - Sum of digits

**Bonus:**
- Write tests for each function
- Benchmark the prime function

---

## Slide 29: Advanced Challenge

**Memoization:**

Cache Fibonacci results to make it faster:

```go
func makeFibonacci() func(int) int {
    cache := make(map[int]int)

    var fib func(int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        if val, ok := cache[n]; ok {
            return val
        }
        cache[n] = fib(n-1) + fib(n-2)
        return cache[n]
    }

    return fib
}
```

---

## Slide 30: What's Next?

**Lecture 6: Arrays and Slices**
- Storing multiple values
- Fixed vs dynamic size
- Slicing and indexing
- Built-in functions

**Your code will handle collections of data! 📊**

**See you next time! 👋**
