# Lecture 8: Structs - Custom Data Types 🏗️
## Building Your Own Data Structures

---

## Slide 1: Why Structs?

**What if you want to store related data?**

**Without structs:**
```go
person1Name := "Alice"
person1Age := 25
person1Email := "alice@example.com"
// Variables are scattered!
```

**With structs:**
```go
type Person struct {
    Name  string
    Age   int
    Email string
}

alice := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
// Organized and clear!
```

**Structs = A way to bundle related data together!**

---

## Slide 2: Defining a Struct

**Create your own type:**

```go
type Person struct {
    Name   string
    Age    int
    Email  string
    Active bool
}

// Another example
type Rectangle struct {
    Width  float64
    Height float64
}
```

**Parts:**
- `type` - Create a new type
- `Person` - Name of your type
- `struct` - It's a structure
- `{ ... }` - Fields inside

---

## Slide 3: Creating Struct Instances

**Different ways to create:**

```go
// Method 1: Field names (clear and flexible)
alice := Person{
    Name:  "Alice",
    Age:   25,
    Email: "alice@example.com",
}

// Method 2: Order matters (must match struct definition)
bob := Person{"Bob", 30, "bob@example.com", true}

// Method 3: Zero values
carol := Person{}  // All fields are zero values

// Method 4: Partial
dave := Person{Name: "Dave", Age: 35}
// Email and Active are zero values ("", 0, false)
```

---

## Slide 4: Accessing Fields

**Use dot (.) to access:**

```go
alice := Person{
    Name: "Alice",
    Age:  25,
}

// Read
fmt.Println(alice.Name)  // "Alice"
fmt.Println(alice.Age)   // 25

// Write
alice.Age = 26
fmt.Println(alice.Age)   // 26

// Update
alice.Email = "alice@work.com"
```

**Simple and intuitive!**

---

## Slide 5: Structs in Slices

**Store multiple people:**

```go
type Person struct {
    Name string
    Age  int
}

// Slice of structs
people := []Person{
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 30},
    {Name: "Carol", Age: 35},
}

// Iterate
for _, person := range people {
    fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}
```

**Very common pattern!**

---

## Slide 6: Anonymous Structs

**Structs without names:**

```go
// One-time use
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  25,
}
```

**Useful for:**
- JSON unmarshaling
- Quick data grouping
- Testing

**But usually better to define a named type!**

---

## Slide 7: Nested Structs

**Structs inside structs:**

```go
type Address struct {
    Street  string
    City    string
    Country string
    ZipCode string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // Nested!
}

// Create
alice := Person{
    Name: "Alice",
    Age:  25,
    Address: Address{
        Street:  "123 Main St",
        City:    "New York",
        Country: "USA",
        ZipCode: "10001",
    },
}
```

---

## Slide 8: Accessing Nested Fields

```go
alice := Person{
    Name: "Alice",
    Address: Address{
        City: "New York",
    },
}

// Access nested fields
fmt.Println(alice.Address.City)      // "New York"
fmt.Println(alice.Address.Country)   // "USA"

// Update nested field
alice.Address.City = "Boston"
```

**Chain the dots!**

---

## Slide 9: Struct Methods (Preview)

**Functions attached to structs:**

```go
type Rectangle struct {
    Width, Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())   // 50
rect.Scale(2)
fmt.Println(rect.Area())   // 200
```

**Full lecture on methods coming soon!**

---

## Slide 10: Exported vs Unexported

**Capitalization matters:**

```go
type Person struct {
    Name    string     // Exported (public) - starts with capital
    age     int        // Unexported (private) - starts with lowercase
    email   string     // Unexported
}

// In same package: can access all
// In other package: can only access Name
```

**Go's access control is simple:**
- Capital letter = Exported (public)
- Lowercase = Unexported (private)

---

## Slide 11: Struct Tags

**Add metadata to fields:**

```go
type Person struct {
    Name  string `json:"name" db:"full_name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}
```

**Used for:**
- JSON encoding/decoding
- Database mapping
- Validation
- Configuration

**The backticks `` ` `` contain the tag!**

---

## Slide 12: JSON Encoding/Decoding

**Working with JSON:**

```go
import "encoding/json"

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

// Encode to JSON
alice := Person{Name: "Alice", Age: 25, Email: "alice@example.com"}
data, _ := json.Marshal(alice)
// {"name":"Alice","age":25,"email":"alice@example.com"}

// Decode from JSON
jsonData := `{"name":"Bob","age":30}`
var bob Person
json.Unmarshal([]byte(jsonData), &bob)
```

---

## Slide 13: Comparing Structs

**Structs can be compared if all fields are comparable:**

```go
type Point struct {
    X, Y int
}

p1 := Point{X: 1, Y: 2}
p2 := Point{X: 1, Y: 2}
p3 := Point{X: 3, Y: 4}

fmt.Println(p1 == p2)  // true
fmt.Println(p1 == p3)  // false
```

**Can't compare if they contain:**
- Slices
- Maps
- Functions

---

## Slide 14: Practice Exercise 1

**Create a Student Management System:**

```go
type Student struct {
    ID       int
    Name     string
    Grades   []int
    Average  float64
}

// Implement:
// 1. Create students
// 2. Calculate average grade
// 3. Find student by ID
// 4. List students with average > 80
```

**Bonus:**
- Add a Class struct that contains students
- Calculate class average

---

## Slide 15: Practice Exercise 2

**Create a Point struct with methods:**

```go
type Point struct {
    X, Y float64
}

// Distance to another point
func (p Point) Distance(other Point) float64

// Move by delta
func (p *Point) Move(dx, dy float64)

// Check if same point
func (p Point) Equals(other Point) bool
```

**Test:**
```go
p1 := Point{X: 0, Y: 0}
p2 := Point{X: 3, Y: 4}
fmt.Println(p1.Distance(p2))  // Should be 5
```

---

## Slide 16: Solutions

**Student average:**
```go
func (s *Student) CalculateAverage() {
    if len(s.Grades) == 0 {
        s.Average = 0
        return
    }
    sum := 0
    for _, grade := range s.Grades {
        sum += grade
    }
    s.Average = float64(sum) / float64(len(s.Grades))
}

func FindStudentByID(students []Student, id int) *Student {
    for i := range students {
        if students[i].ID == id {
            return &students[i]
        }
    }
    return nil
}
```

---

## Slide 17: Solutions (cont.)

**Point methods:**
```go
import "math"

func (p Point) Distance(other Point) float64 {
    dx := p.X - other.X
    dy := p.Y - other.Y
    return math.Sqrt(dx*dx + dy*dy)
}

func (p *Point) Move(dx, dy float64) {
    p.X += dx
    p.Y += dy
}

func (p Point) Equals(other Point) bool {
    return p.X == other.X && p.Y == other.Y
}
```

---

## Slide 18: Common Mistakes

**Mistake 1: Expecting copy to update original**
```go
func updateAge(person Person) {
    person.Age = 30  // Changes copy only!
}

// Use pointer instead:
func updateAge(person *Person) {
    person.Age = 30  // Changes original
}
```

**Mistake 2: Zero value confusion**
```go
var p Person
if p == nil { }  // ❌ Structs can't be nil!
// Only pointers to structs can be nil
```

**Mistake 3: Comparing structs with slices**
```go
type Person struct {
    Name   string
    Grades []int
}
p1 == p2  // ❌ Compile error! Can't compare slices
```

---

## Slide 19: Best Practices

**Do:**
- ✅ Group related fields in structs
- ✅ Use clear, descriptive field names
- ✅ Export what should be public (capitalize)
- ✅ Add JSON tags for API structs
- ✅ Use pointers for large structs in functions

**Don't:**
- ❌ Create giant structs with 20+ fields
- ❌ Mix concerns in one struct
- ❌ Export everything
- ❌ Forget to document exported types

---

## Slide 20: Quiz Time!

**Q1: Is this valid?**
```go
type Thing struct {
    a int
    B string
}
```

**Q2: How do you access the City?**
```go
type Address struct { City string }
type Person struct { Address Address }
p := Person{Address: Address{City: "NYC"}}
// ???
```

**Q3: What's the difference?**
```go
func (p Person) Method1()
func (p *Person) Method2()
```

---

## Slide 21: Key Takeaways

**What you learned:**

1. ✅ Structs bundle related data
2. ✅ Define with `type Name struct`
3. ✅ Create with field names or positionally
4. ✅ Access with dot notation
5. ✅ Can nest structs
6. ✅ Capital = exported, lowercase = private
7. ✅ Use tags for metadata

---

## Slide 22: Struct Cheat Sheet

| Task | Syntax |
|------|--------|
| Define | `type T struct { ... }` |
| Create | `T{Field: value}` |
| Create (positional) | `T{value1, value2}` |
| Access | `instance.Field` |
| Nested access | `instance.Inner.Field` |
| Method | `func (r T) Method() {}` |
| Pointer method | `func (r *T) Method() {}` |

---

## Slide 23: Real-World Example

**Database model:**
```go
type User struct {
    ID        int       `json:"id" db:"id"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"-" db:"password"`  // Hidden in JSON
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Post struct {
    ID       int    `json:"id"`
    UserID   int    `json:"user_id"`
    Title    string `json:"title"`
    Content  string `json:"content"`
    Author   User   `json:"author,omitempty"`
}
```

---

## Slide 24: Homework

**Build a Library Management System:**

Types needed:
- `Book` (ID, Title, Author, ISBN, Available)
- `Member` (ID, Name, Email, BorrowedBooks)
- `Library` (Books, Members)

Features:
1. Add/remove books
2. Register members
3. Borrow/return books
4. List available books
5. List member's borrowed books
6. Search books by title/author

**Bonus:** Save/load from JSON file

---

## Slide 25: What's Next?

**Lecture 9: Pointers**
- Memory addresses
- Dereferencing
- Pointer receivers
- nil pointers

**Your code will be memory-efficient! 💾**

**See you next time! 👋**
