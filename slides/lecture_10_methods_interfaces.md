# Lecture 10: Methods and Interfaces 🔌
## Object-Oriented Programming in Go

---

## Slide 1: What are Methods?

**Functions attached to types:**

**Regular function:**
```go
func area(r Rectangle) float64 {
    return r.Width * r.Height
}
// Called: area(rect)
```

**Method:**
```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
// Called: rect.Area()
```

**The difference:**
- Methods belong to a type
- Called with dot notation
- Can modify the receiver (with pointers)

---

## Slide 2: Defining Methods

**Syntax:**
```go
func (receiver ReceiverType) MethodName(params) returnType {
    // method body
}
```

**Example:**
```go
type Rectangle struct {
    Width, Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

---

## Slide 3: Value vs Pointer Receivers

**Value Receiver (copy):**
```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
// r is a COPY - modifications don't affect original
```

**Pointer Receiver (reference):**
```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
// r is a POINTER - modifications affect original
```

**When to use:**
- Value: Read-only, small structs
- Pointer: Modification, large structs

---

## Slide 4: Calling Methods

**Both ways work:**
```go
rect := Rectangle{Width: 10, Height: 5}

// Value method
area := rect.Area()  // 50

// Pointer method - Go automatically converts!
rect.Scale(2)  // Same as (&rect).Scale(2)

// Can also call with pointer
ptr := &rect
ptr.Area()  // Go automatically dereferences
```

**Go makes it convenient!**

---

## Slide 5: What are Interfaces?

**Think of it like a job description:**
- Job: "Shape"
- Requirements: Must have Area() and Perimeter()
- Circle qualifies ✓
- Rectangle qualifies ✓
- Car does not qualify ✗

**In code:**
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

**Any type that has these methods IS a Shape!**

---

## Slide 6: Defining Interfaces

**Interface = Set of methods:**

```go
// Define the interface
type Stringer interface {
    String() string
}

// Define a type
type Person struct {
    Name string
    Age  int
}

// Implement the interface (automatic!)
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}
```

**Person now implements Stringer!**

---

## Slide 7: Interface Satisfaction

**In Go, you don't say "I implement" - you just do!**

```go
type Shape interface {
    Area() float64
}

type Circle struct { Radius float64 }

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Circle automatically implements Shape!
var s Shape = Circle{Radius: 5}  // ✓ Works!
```

**Implicit implementation - no "implements" keyword!**

---

## Slide 8: Using Interfaces

**Write functions that accept interfaces:**

```go
// Accept ANY type that has String() method
func printString(s Stringer) {
    fmt.Println(s.String())
}

type Person struct { Name string }
func (p Person) String() string { return p.Name }

type Book struct { Title string }
func (b Book) String() string { return b.Title }

// Both work!
printString(Person{Name: "Alice"})
printString(Book{Title: "Go in Action"})
```

**Polymorphism in Go!**

---

## Slide 9: Empty Interface

**Interface with no methods - accepts anything:**

```go
// Can hold any type
var anything interface{}

anything = 42
anything = "hello"
anything = true
anything = []int{1, 2, 3}
```

**Use cases:**
```go
func printAny(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

printAny(42)
printAny("hello")
printAny([]int{1, 2, 3})
```

---

## Slide 10: Type Assertions

**Extract concrete type from interface:**

```go
var s Shape = Circle{Radius: 5}

// Assert it's a Circle
circle := s.(Circle)
fmt.Println(circle.Radius)  // Works!

// Safe assertion
circle, ok := s.(Circle)
if ok {
    fmt.Println("It's a circle with radius", circle.Radius)
}
```

**Wrong assertion PANICS:**
```go
rect := s.(Rectangle)  // ❌ Panic if s is not Rectangle!
```

---

## Slide 11: Type Switches

**Check multiple types:**

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

describe(42)       // Integer: 42
describe("hello")  // String: hello
describe(true)     // Boolean: true
```

---

## Slide 12: Common Interfaces

**Built-in interfaces:**

```go
// Stringer - for String() method
type Stringer interface {
    String() string
}

// Reader - for reading data
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer - for writing data
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Error - for errors
type Error interface {
    Error() string
}
```

**These are used everywhere!**

---

## Slide 13: Interface Composition

**Build bigger interfaces from smaller ones:**

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Composed interface
type ReadWriter interface {
    Reader
    Writer
}

// ReadWriter requires BOTH Read and Write methods
```

**Like building with Lego blocks!**

---

## Slide 14: Practice Exercise 1

**Create a Payment interface:**

```go
type Payment interface {
    Process(amount float64) error
    Refund(amount float64) error
}

// Implement for:
// CreditCard
// PayPal
// Crypto

// Function:
// func ProcessPayment(p Payment, amount float64) error
```

**Each implementation should print how it's processing.**

---

## Slide 15: Practice Exercise 2

**Create a Logger interface:**

```go
type Logger interface {
    Log(message string)
}

// Implement:
// ConsoleLogger - prints to stdout
// FileLogger - writes to file
// MultiLogger - logs to multiple loggers
```

**Usage:**
```go
logger := MultiLogger{
    loggers: []Logger{
        ConsoleLogger{},
        FileLogger{Filename: "app.log"},
    },
}
logger.Log("Application started")
```

---

## Slide 16: Solutions

**Payment interface:**
```go
type CreditCard struct { Number string }

func (c CreditCard) Process(amount float64) error {
    fmt.Printf("Charging $%.2f to credit card ending in %s\n",
        amount, c.Number[len(c.Number)-4:])
    return nil
}

func (c CreditCard) Refund(amount float64) error {
    fmt.Printf("Refunding $%.2f to credit card\n", amount)
    return nil
}

func ProcessPayment(p Payment, amount float64) error {
    return p.Process(amount)
}
```

---

## Slide 17: Solutions (cont.)

**MultiLogger:**
```go
type MultiLogger struct {
    loggers []Logger
}

func (m MultiLogger) Log(message string) {
    for _, logger := range m.loggers {
        logger.Log(message)
    }
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
    fmt.Println("[CONSOLE]", message)
}
```

---

## Slide 18: Best Practices

**Do:**
- ✅ Keep interfaces small (1-3 methods)
- ✅ Use descriptive names
- ✅ Accept interfaces, return concrete types
- ✅ Define interfaces where they're used

**Don't:**
- ❌ Create giant interfaces with 10+ methods
- ❌ Export interfaces that aren't needed
- ❌ Force types to implement interfaces

**The Go proverb:** "The bigger the interface, the weaker the abstraction"

---

## Slide 19: Interface Values

**What's inside an interface value?**

```go
var s Shape = Circle{Radius: 5}
```

**Interface contains:**
1. **Type** - Concrete type (Circle)
2. **Value** - The actual value

```go
var s Shape
type: nil, value: nil  // nil interface

s = Circle{Radius: 5}
type: Circle, value: Circle{Radius: 5}

// nil pointer can be tricky!
var c *Circle = nil
s = c  // type: *Circle, value: nil
s != nil  // TRUE! (type is set)
```

---

## Slide 20: Common Mistakes

**Mistake 1: Nil pointer in interface**
```go
func GetCircle() *Circle { return nil }

var s Shape = GetCircle()
if s != nil {
    s.Area()  // ❌ Panic! Value is nil!
}
```

**Mistake 2: Wrong method signature**
```go
type Shape interface {
    Area() float64
}

type Circle struct{}

func (c Circle) area() float64 {  // ❌ Lowercase!
    return 3.14
}
// Circle doesn't implement Shape!
```

---

## Slide 21: Quiz Time!

**Q1: Does Rectangle implement Shape?**
```go
type Shape interface { Area() float64 }
type Rectangle struct{}
func (r Rectangle) Area() int { return 0 }
```

**Q2: What's the output?**
```go
var i interface{} = 42
v, ok := i.(string)
fmt.Println(v, ok)
```

**Q3: Which receiver should modify?**
A) func (r Rectangle) Update()
B) func (r *Rectangle) Update()

---

## Slide 22: Key Takeaways

**What you learned:**

1. ✅ Methods are functions with receivers
2. ✅ Pointer receivers can modify
3. ✅ Interfaces define behavior contracts
4. ✅ Types implement interfaces implicitly
5. ✅ Accept interfaces, return concrete types
6. ✅ Keep interfaces small and focused
7. ✅ Use type assertions and type switches

---

## Slide 23: Interface Cheat Sheet

| Task | Syntax |
|------|--------|
| Define interface | `type I interface { Method() }` |
| Implement | Just write the methods! |
| Use interface | `func f(i I)` |
| Type assert | `v := i.(Concrete)` |
| Safe assert | `v, ok := i.(Concrete)` |
| Type switch | `switch v := i.(type)` |

---

## Slide 24: Real-World Example

**HTTP handler interface:**
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}

// Use
http.Handle("/", MyHandler{})
http.ListenAndServe(":8080", nil)
```

---

## Slide 25: Homework

**Build a Notification System:**

```go
type Notifier interface {
    Send(message string, recipient string) error
}

// Implement:
// EmailNotifier
// SMSNotifier
// PushNotifier

// NotificationManager that can:
// - Register notifiers
// - Send to all registered notifiers
// - Send to specific type
```

**Bonus:** Add retry logic and error handling

---

## Slide 26: What's Next?

**Lecture 11: Packages and Modules**
- Organizing code
- Importing packages
- Creating modules
- Versioning

**Your code will be modular and reusable! 📦**

**See you next time! 👋**
