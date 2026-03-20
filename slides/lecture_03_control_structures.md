# Lecture 3: Control Structures - Making Decisions 🚦
## If, Else, and Switch Statements

---

## Slide 1: Making Decisions

**Life is full of decisions:**
- If it's raining, take an umbrella ☂️
- If you're hungry, eat something 🍕
- If it's bedtime, go to sleep 😴

**Programs need to make decisions too!**

**In Go, we use:**
- `if` statements - "If this is true, do that"
- `else` statements - "Otherwise, do this"
- `switch` statements - "Choose from many options"

---

## Slide 2: The If Statement

**Basic structure:**
```go
if condition {
    // Do something
}
```

**Real example:**
```go
age := 18

if age >= 18 {
    fmt.Println("You can vote!")
}
```

**How it works:**
1. Check if `age >= 18` is true
2. If YES → run the code inside {}
3. If NO → skip the code inside {}

---

## Slide 3: Comparison Operators

**Used to compare values:**

| Operator | Meaning | Example | Result |
|----------|---------|---------|--------|
| == | Equal to | 5 == 5 | true |
| != | Not equal | 5 != 3 | true |
| < | Less than | 3 < 5 | true |
| > | Greater than | 5 > 3 | true |
| <= | Less or equal | 5 <= 5 | true |
| >= | Greater or equal | 5 >= 3 | true |

**Remember:** `=` is for assignment, `==` is for comparison!

---

## Slide 4: If-Else Statement

**Two paths - which one to take?**

```go
age := 16

if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}
```

**Flow:**
1. Check condition
2. If TRUE → run if block
3. If FALSE → run else block
4. Only ONE block runs, never both!

---

## Slide 5: If-Else If-Else

**Multiple conditions:**

```go
score := 85

if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

**Checks from top to bottom, stops at first TRUE!**

---

## Slide 6: Logical Operators

**Combine conditions:**

**AND (`&&`) - Both must be true:**
```go
if age >= 13 && age <= 19 {
    fmt.Println("You are a teenager")
}
// Both conditions must be true
```

**OR (`||`) - At least one must be true:**
```go
if day == "Saturday" || day == "Sunday" {
    fmt.Println("It's the weekend!")
}
// Either condition can be true
```

**NOT (`!`) - Flip the value:**
```go
if !isRaining {
    fmt.Println("Go outside!")
}
// Runs if isRaining is false
```

---

## Slide 7: Truth Tables

**AND (`&&`):**
| A | B | A && B |
|---|---|---------|
| true | true | true |
| true | false | false |
| false | true | false |
| false | false | false |

**OR (`||`):**
| A | B | A \|\| B |
|---|---|---------|
| true | true | true |
| true | false | true |
| false | true | true |
| false | false | false |

---

## Slide 8: Real-World Example

**Movie ticket pricing:**

```go
age := 25
isStudent := true

if age < 13 {
    fmt.Println("Child ticket: $5")
} else if age >= 65 {
    fmt.Println("Senior ticket: $6")
} else if age >= 13 && age <= 25 && isStudent {
    fmt.Println("Student ticket: $7")
} else {
    fmt.Println("Adult ticket: $10")
}
```

---

## Slide 9: The Switch Statement

**When you have many options:**

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("Almost weekend!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Regular day")
}
```

**Cleaner than many if-else statements!**

---

## Slide 10: Switch Without Expression

**Switch on conditions:**

```go
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

**Like if-else, but cleaner for many cases!**

---

## Slide 11: Switch Features

**Fallthrough (rarely used):**
```go
number := 1

switch number {
case 1:
    fmt.Println("One")
    fallthrough  // Continue to next case
case 2:
    fmt.Println("Two")
}
// Prints: One, Two
```

**No break needed!** Go automatically breaks after each case.

---

## Slide 12: Short If Statement

**Declare variable in the if:**

```go
if age := 25; age >= 18 {
    fmt.Println("Adult")
}
// 'age' only exists inside the if block
```

**Useful for:**
```go
if length := len(name); length > 10 {
    fmt.Println("Long name!")
}
```

**The variable is scoped to the if block!**

---

## Slide 13: Common Patterns

**Check if number is even or odd:**
```go
number := 7

if number % 2 == 0 {
    fmt.Println("Even")
} else {
    fmt.Println("Odd")
}
```

**Check if year is a leap year:**
```go
year := 2024

if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
    fmt.Println("Leap year!")
}
```

---

## Slide 14: Nested If Statements

**If inside an if:**

```go
isLoggedIn := true
isAdmin := true

if isLoggedIn {
    fmt.Println("Welcome!")
    if isAdmin {
        fmt.Println("Admin panel available")
    }
} else {
    fmt.Println("Please log in")
}
```

**Don't nest too deep!** It gets confusing.

---

## Slide 15: Practice Exercise 1

**Create a number guessing game:**

```go
secretNumber := 42
guess := 30

// Write code that prints:
// - "Too high!" if guess > secretNumber
// - "Too low!" if guess < secretNumber
// - "Correct!" if guess == secretNumber
```

**Bonus:** Add a message for "close" (within 5 numbers)

---

## Slide 16: Practice Exercise 2

**Day of week activity:**

```go
day := "Saturday"

// Use switch to print:
// - Monday-Friday: "Go to work/school"
// - Saturday-Sunday: "Weekend fun!"
// - Invalid day: "That's not a day!"
```

---

## Slide 17: Common Mistakes

**Mistake 1: Assignment instead of comparison**
```go
if x = 5 {      // ❌ Assignment!
if x == 5 {     // ✅ Comparison
```

**Mistake 2: Missing braces**
```go
if x > 5
    fmt.Println("Big")  // ❌ Needs {}

if x > 5 {
    fmt.Println("Big")  // ✅ Correct
}
```

**Mistake 3: Unreachable code**
```go
if true {
    return
}
fmt.Println("Hi")  // ❌ Never runs!
```

---

## Slide 18: Boolean Variables

**Store conditions in variables:**

```go
isWeekend := day == "Saturday" || day == "Sunday"
isHoliday := day == "Christmas"

if isWeekend || isHoliday {
    fmt.Println("No work today!")
}
```

**Makes code readable!**

---

## Slide 19: Ternary? No!

**Other languages have:**
```javascript
// JavaScript
result = (age >= 18) ? "Adult" : "Minor"
```

**Go doesn't have ternary!** Use if-else:
```go
var result string
if age >= 18 {
    result = "Adult"
} else {
    result = "Minor"
}
```

**Go values readability over brevity!**

---

## Slide 20: Real Example - Login System

```go
username := "alice"
password := "secret123"

storedUsername := "alice"
storedPassword := "secret123"
isActive := true

if username == storedUsername && password == storedPassword {
    if isActive {
        fmt.Println("Login successful!")
    } else {
        fmt.Println("Account disabled")
    }
} else {
    fmt.Println("Invalid credentials")
}
```

---

## Slide 21: Type Assertions with Switch

**Check what type something is:**

```go
var value interface{} = "Hello"

switch v := value.(type) {
case string:
    fmt.Printf("String: %s\n", v)
case int:
    fmt.Printf("Integer: %d\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

**Advanced topic - we'll cover interfaces later!**

---

## Slide 22: Best Practices

**1. Keep it simple:**
```go
// Good
if user.IsActive {
    return user
}
return nil

// Overly complex
if user.IsActive == true { ... }
```

**2. Early returns:**
```go
func checkAge(age int) {
    if age < 0 {
        return  // Exit early
    }
    // Continue with valid age
}
```

**3. Use switch for multiple options**

---

## Slide 23: Debugging If Statements

**Add print statements to see what's happening:**

```go
age := 15
isStudent := true

fmt.Println("Debug: age =", age)
fmt.Println("Debug: isStudent =", isStudent)

if age >= 18 {
    fmt.Println("Adult")
} else if isStudent {
    fmt.Println("Student")
}
```

**Check your conditions:**
```go
fmt.Println("Condition:", age >= 18)  // See if true/false
```

---

## Slide 24: Exercise Solutions

**Number guessing game:**
```go
secretNumber := 42
guess := 30

if guess > secretNumber {
    fmt.Println("Too high!")
} else if guess < secretNumber {
    fmt.Println("Too low!")
} else {
    fmt.Println("Correct!")
}

// Bonus: Close guess
if diff := secretNumber - guess; diff >= -5 && diff <= 5 {
    fmt.Println("Close!")
}
```

---

## Slide 25: Exercise Solutions (cont.)

**Day of week:**
```go
day := "Saturday"

switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Go to work/school")
case "Saturday", "Sunday":
    fmt.Println("Weekend fun!")
default:
    fmt.Println("That's not a day!")
}
```

---

## Slide 26: Quiz Time!

**Q1: What's the output?**
```go
x := 10
if x > 5 {
    fmt.Println("A")
} else if x > 3 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}
```

**Q2: Which is true?**
```go
a, b := 5, 10
// A) a > b && b > 0
// B) a < b || a > 100
// C) !(a == b)
```

**Q3: Fix the bug:**
```go
if x = 10 {
    fmt.Println("Ten")
}
```

---

## Slide 27: Key Takeaways

**What you learned:**

1. ✅ `if` checks conditions
2. ✅ `else` provides alternative
3. ✅ `else if` handles multiple cases
4. ✅ Comparison operators: ==, !=, <, >, <=, >=
5. ✅ Logical operators: &&, ||, !
6. ✅ `switch` is cleaner for many options
7. ✅ Go doesn't have ternary operator

---

## Slide 28: Common Use Cases

**Input validation:**
```go
if age < 0 || age > 150 {
    fmt.Println("Invalid age")
}
```

**Feature flags:**
```go
if featureEnabled {
    showNewFeature()
}
```

**Access control:**
```go
if user.IsAdmin || user.HasPermission("read") {
    showData()
}
```

---

## Slide 29: Homework

**Create a "Smart Calculator":**

1. Ask for two numbers
2. Ask for an operation (+, -, *, /)
3. Use if-else or switch to perform the operation
4. Handle division by zero!
5. Handle invalid operations

**Example:**
```
Enter first number: 10
Enter second number: 5
Enter operation: /
Result: 2
```

---

## Slide 30: What's Next?

**Lecture 4: Loops!**
- Doing things over and over
- For loops
- Range loops
- Break and continue

**Your programs will do repetitive tasks easily! 🔄**

**See you next time! 👋**
