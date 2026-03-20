# Lecture 12: Error Handling 💪
## Robust and Reliable Code

---

## Slide 1: Why Error Handling?

**Things go wrong:**
- File not found 📄❌
- Network disconnected 🌐❌
- Invalid user input ⌨️❌
- Database unavailable 🗄️❌

**Without error handling:**
```
Program crashes! 💥
Users see confusing messages
Data gets corrupted
```

**With error handling:**
```
Graceful recovery
Clear error messages
Program keeps running
```

**Go philosophy:**
> "Errors are values. Handle them explicitly."

---

## Slide 2: The error Interface

**Errors in Go are just values:**

```go
type error interface {
    Error() string
}
```

**Any type with Error() method is an error!**

**Simplest way to create errors:**
```go
import "errors"

err := errors.New("something went wrong")
fmt.Println(err)  // "something went wrong"
```

---

## Slide 3: Returning Errors

**Go functions return errors as the last value:**

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

**Usage:**
```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

**Always check errors!**

---

## Slide 4: The nil Error

**nil means "no error":**

```go
func doSomething() error {
    // If everything is OK
    return nil
}

err := doSomething()
if err != nil {
    // Handle error
} else {
    // Success!
}
```

**Pattern:**
- Success → return result, nil
- Failure → return zero value, error

---

## Slide 5: Error Checking Patterns

**Standard pattern:**
```go
result, err := someFunction()
if err != nil {
    // Handle error
    log.Fatal(err)
    return
}
// Use result
```

**Inline (for simple cases):**
```go
if result, err := someFunction(); err != nil {
    return err
}
```

**Don't ignore errors!** ❌ `result, _ := mightFail()`

---

## Slide 6: Custom Error Types

**Create your own error types:**

```go
type ValidationError struct {
    Field string
    Issue string
}

func (v *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s",
        v.Field, v.Issue)
}

// Usage
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field: "age",
            Issue: "cannot be negative",
        }
    }
    return nil
}
```

---

## Slide 7: Error Wrapping

**Add context to errors:**

```go
import "fmt"

func readFile(path string) error {
    _, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("failed to read file: %w", err)
    }
    return nil
}
```

**The `%w` verb wraps the error!**

**Benefits:**
- Preserve original error
- Add context
- Can unwrap later

---

## Slide 8: Error Unwrapping

**Extract wrapped errors:**

```go
import "errors"

// Check if error is a specific type
var err = readFile("config.txt")

// Unwrap
if errors.Unwrap(err) == os.ErrNotExist {
    fmt.Println("File doesn't exist")
}
```

**Or use errors.Is:**
```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File doesn't exist")
}
```

**Checks the entire chain!**

---

## Slide 9: Error As

**Extract custom error types:**

```go
var err error = &ValidationError{Field: "age", Issue: "too low"}

var validationErr *ValidationError
if errors.As(err, &validationErr) {
    fmt.Println("Field:", validationErr.Field)
    fmt.Println("Issue:", validationErr.Issue)
}
```

**Useful for:**
- Getting specific error details
- Different handling for different errors
- Logging specific fields

---

## Slide 10: Sentinel Errors

**Predefined error values:**

```go
var ErrNotFound = errors.New("resource not found")
var ErrInvalidInput = errors.New("invalid input")
var ErrUnauthorized = errors.New("unauthorized")

func findUser(id int) (*User, error) {
    user := lookup(id)
    if user == nil {
        return nil, ErrNotFound
    }
    return user, nil
}

// Check
if errors.Is(err, ErrNotFound) {
    fmt.Println("User not found")
}
```

---

## Slide 11: Panic and Recover

**When things are REALLY wrong:**

```go
// Panic stops the program
func mustNotFail() {
    if somethingTerrible {
        panic("something terrible happened!")
    }
}
```

**Recover (rarely used):**
```go
func safeOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    mightPanic()
}
```

**Don't use panic for normal errors!**

---

## Slide 12: Practice Exercise 1

**Create a robust file reader:**

```go
func ReadConfig(path string) (*Config, error) {
    // Should:
    // 1. Check if file exists
    // 2. Check if file is readable
    // 3. Parse JSON
    // 4. Validate required fields
    // Return appropriate errors for each case
}
```

**Use custom errors and wrapping!**

---

## Slide 13: Practice Exercise 2

**Create a bank account with error handling:**

```go
type Account struct {
    Balance float64
}

// Return errors for:
// - Withdrawing more than balance
// - Negative deposits
// - Negative withdrawals

func (a *Account) Withdraw(amount float64) error
func (a *Account) Deposit(amount float64) error
func (a *Account) Transfer(to *Account, amount float64) error
```

---

## Slide 14: Solutions

**File reader:**
```go
type ConfigError struct {
    Path string
    Err  error
}

func (c *ConfigError) Error() string {
    return fmt.Sprintf("config error for %s: %v", c.Path, c.Err)
}

func ReadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, &ConfigError{Path: path, Err: err}
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("invalid JSON: %w", err)
    }

    return &cfg, nil
}
```

---

## Slide 15: Solutions (cont.)

**Bank account:**
```go
var (
    ErrInsufficientFunds = errors.New("insufficient funds")
    ErrNegativeAmount    = errors.New("amount cannot be negative")
)

func (a *Account) Withdraw(amount float64) error {
    if amount < 0 {
        return ErrNegativeAmount
    }
    if amount > a.Balance {
        return ErrInsufficientFunds
    }
    a.Balance -= amount
    return nil
}
```

---

## Slide 16: Common Mistakes

**Mistake 1: Ignoring errors**
```go
file, _ := os.Open("file.txt")  // ❌ Dangerous!
```

**Mistake 2: Using panic for control flow**
```go
if user == nil {
    panic("user not found")  // ❌ Don't do this!
}
```

**Mistake 3: Not wrapping errors**
```go
return err  // ❌ Lost context!
return fmt.Errorf("reading config: %w", err)  // ✅ Better
```

---

## Slide 17: Best Practices

**Do:**
- ✅ Check every error
- ✅ Add context with wrapping
- ✅ Use custom types for program logic
- ✅ Use sentinel errors for checks
- ✅ Return errors, don't log and continue

**Don't:**
- ❌ Ignore errors with _
- ❌ Use panic for expected errors
- ❌ Create errors inside hot paths
- ❌ Return generic errors for everything

---

## Slide 18: Quiz Time!

**Q1: How do you wrap an error?**
A) fmt.Errorf("message", err)
B) fmt.Errorf("message: %w", err)
C) errors.Wrap(err, "message")

**Q2: What does nil error mean?**
A) Unknown error
B) No error
C) Severe error

**Q3: When should you panic?**
A) File not found
B) Invalid input
C) Programming bug

---

## Slide 19: Key Takeaways

**What you learned:**

1. ✅ Errors are values, handle them!
2. ✅ Return errors as last value
3. ✅ Check errors immediately
4. ✅ Wrap with %w for context
5. ✅ Use errors.Is and errors.As
6. ✅ Create custom error types
7. ✅ Don't panic for normal errors

---

## Slide 20: Error Cheat Sheet

| Task | How |
|------|-----|
| Create error | `errors.New("msg")` |
| Wrap error | `fmt.Errorf("ctx: %w", err)` |
| Check specific | `errors.Is(err, target)` |
| Extract type | `errors.As(err, &custom)` |
| Custom type | Implement `Error() string` |
| Sentinel | `var Err = errors.New(...)` |

---

## Slide 21: Real-World Example

**HTTP handler:**
```go
func GetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id", http.StatusBadRequest)
        return
    }

    user, err := db.GetUser(id)
    if err != nil {
        if errors.Is(err, ErrNotFound) {
            http.Error(w, "user not found", http.StatusNotFound)
            return
        }
        log.Printf("database error: %v", err)
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}
```

---

## Slide 22: Homework

**Build a Robust CLI Tool:**

Create a file processor that:
1. Reads a CSV file
2. Validates each row
3. Transforms data
4. Writes JSON output

**Error handling requirements:**
- Custom error types for each failure
- Detailed error messages with line numbers
- Continue processing on row errors
- Summary report at the end

---

## Slide 23: What's Next?

**Lecture 13: File I/O**
- Reading files
- Writing files
- Working with directories
- CSV and JSON

**Your code will interact with the world! 📁**

**See you next time! 👋**
