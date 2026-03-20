# Lecture 9: Pointers - Memory Addresses 💾
## Understanding Memory and References

---

## Slide 1: Why Pointers?

**Two scenarios:**

**Scenario 1: Giving a copy**
- You give your friend a photocopy of your homework
- They write on it
- Your original is unchanged

**Scenario 2: Giving the location**
- You tell your friend where your house is
- They go there and paint the door
- Your house actually changes!

**In programming:**
- **Values** = Copy (safe, but uses more memory)
- **Pointers** = Location (efficient, but more complex)

---

## Slide 2: What is a Pointer?

**A pointer is a variable that stores a memory address:**

```
┌─────────────┐      ┌──────────┐
│   Pointer   │──────│  Value   │
│  0x123456   │      │   42     │
└─────────────┘      └──────────┘
     Address              Data
```

**Think of it like:**
- **Pointer** = Street address
- **Value** = The actual house

**Why use them?**
- Modify original data
- Avoid copying large structs
- Share data between functions

---

## Slide 3: Pointer Syntax

**Creating pointers:**

```go
// Method 1: & operator (address of)
age := 25
agePtr := &age
// agePtr now contains the memory address of age

// Method 2: new function
agePtr := new(int)  // Creates pointer to int
*agePtr = 25

// Method 3: with struct
person := &Person{Name: "Alice"}
```

**The `&` means "address of"**
**The `*` means "value at address"**

---

## Slide 4: Dereferencing Pointers

**Access the value a pointer points to:**

```go
age := 25
agePtr := &age

// Dereference with *
value := *agePtr  // value is 25

// Modify through pointer
*agePtr = 30

fmt.Println(age)  // 30! Original changed!
```

**Think of `*` as "go to that address"**

---

## Slide 5: Pointers in Functions

**Without pointers (copy):**
```go
func doubleValue(x int) {
    x = x * 2
}

func main() {
    num := 10
    doubleValue(num)
    fmt.Println(num)  // Still 10!
}
```

**With pointers (modify original):**
```go
func doubleValue(x *int) {
    *x = *x * 2  // Dereference and modify
}

func main() {
    num := 10
    doubleValue(&num)  // Pass address
    fmt.Println(num)   // Now 20!
}
```

---

## Slide 6: Pointer to Structs

**Efficient way to modify structs:**

```go
type Person struct {
    Name string
    Age  int
}

// Without pointer - makes a copy
func birthdayBad(p Person) {
    p.Age++  // Changes copy only!
}

// With pointer - modifies original
func birthdayGood(p *Person) {
    p.Age++  // Changes original!
}

func main() {
    alice := Person{Name: "Alice", Age: 25}
    birthdayGood(&alice)
    fmt.Println(alice.Age)  // 26
}
```

---

## Slide 7: Automatic Dereferencing

**Go makes it easy with structs:**

```go
func (p *Person) Birthday() {
    p.Age++  // Same as (*p).Age++
}

// Both work:
p := &Person{Age: 25}
p.Age = 26      // Automatic dereference
(*p).Age = 27   // Explicit dereference
```

**Go automatically dereferences for you!**

---

## Slide 8: nil Pointers

**A pointer that points to nothing:**

```go
var ptr *int  // nil by default

fmt.Println(ptr == nil)  // true

// Dereferencing nil PANICS!
value := *ptr  // ❌ Runtime panic!
```

**Always check before dereferencing:**
```go
if ptr != nil {
    fmt.Println(*ptr)
}
```

**Common source of crashes!**

---

## Slide 9: When to Use Pointers

**Use pointers when:**

1. **Modifying the original:**
```go
func update(person *Person)
```

2. **Large structs** (avoid copying):
```go
func process(data *HugeStruct)
```

3. **Optional values** (can be nil):
```go
func findUser(id int) *User  // nil if not found
```

4. **Sharing data:**
```go
multipleFunctions(&sharedData)
```

---

## Slide 10: When NOT to Use Pointers

**Don't use pointers for:**

1. **Small values** (int, float, bool):
```go
// Bad
func double(x *int) int
// Good
func double(x int) int
```

2. **Slices and maps** (already references!):
```go
// Bad - slice is already a reference
func process(data *[]int)
// Good
func process(data []int)
```

3. **When you don't need to modify**

---

## Slide 11: Pointer Arithmetic?

**Go does NOT have pointer arithmetic:**

```c
// C code - valid there
int *ptr = &x;
ptr++;  // Move to next memory location
```

```go
// Go - NOT allowed!
ptr := &x
ptr++  // ❌ Compile error!
```

**Why?** Safety! Pointer arithmetic is error-prone.

**But you can still:**
- Compare pointers (==)
- Assign pointers
- Pass to functions

---

## Slide 12: Pointers vs Values

**Comparison:**

| Aspect | Value | Pointer |
|--------|-------|---------|
| Memory | Copy | Shared |
| Safety | High | Medium |
| nil? | No | Yes |
| Modify original | No | Yes |
| Use for | Small data | Large structs |

**Default to values, use pointers when needed!**

---

## Slide 13: Practice Exercise 1

**Swap two integers:**

```go
func swap(a, b *int) {
    // Exchange values at a and b
}

func main() {
    x, y := 5, 10
    swap(&x, &y)
    fmt.Println(x, y)  // Should be 10, 5
}
```

**Hint:** Use a temporary variable

---

## Slide 14: Practice Exercise 2

**Linked List node:**

```go
type Node struct {
    Value int
    Next  *Node  // Pointer to next node
}

// Create a list: 1 -> 2 -> 3
// Print all values
// Count number of nodes
```

**This is how linked lists work!**

---

## Slide 15: Solutions

**Swap:**
```go
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

// Or without temp (Go way):
func swap(a, b *int) {
    *a, *b = *b, *a
}
```

**Linked list:**
```go
// Create list
head := &Node{Value: 1}
head.Next = &Node{Value: 2}
head.Next.Next = &Node{Value: 3}

// Print
for node := head; node != nil; node = node.Next {
    fmt.Println(node.Value)
}

// Count
count := 0
for node := head; node != nil; node = node.Next {
    count++
}
```

---

## Slide 16: Common Mistakes

**Mistake 1: Dereferencing nil**
```go
var p *Person
fmt.Println(p.Name)  // ❌ PANIC!

// Fix:
if p != nil {
    fmt.Println(p.Name)
}
```

**Mistake 2: Forgetting & when passing**
```go
func update(p *Person) { }
p := Person{}
update(p)    // ❌ Wrong!
update(&p)   // ✅ Right!
```

**Mistake 3: Returning pointer to local variable**
```go
// This is actually OK in Go!
func createPerson() *Person {
    p := Person{Name: "Alice"}
    return &p  // Go escapes to heap
}
```

---

## Slide 17: Best Practices

**Do:**
- ✅ Check for nil before dereferencing
- ✅ Use pointers for large structs
- ✅ Return pointers for optional results
- ✅ Document when function modifies via pointer

**Don't:**
- ❌ Use pointers for small values
- ❌ Use pointers for slices/maps
- ❌ Ignore nil checks
- ❌ Overuse pointers "just in case"

---

## Slide 18: Quiz Time!

**Q1: What's the output?**
```go
x := 10
p := &x
*p = 20
fmt.Println(x)
```

**Q2: Fix the bug:**
```go
func setToZero(x int) {
    x = 0
}
num := 5
setToZero(num)
fmt.Println(num)  // Still 5!
```

**Q3: What's wrong?**
```go
var p *int
fmt.Println(*p)
```

---

## Slide 19: Key Takeaways

**What you learned:**

1. ✅ `&` gets the address
2. ✅ `*` dereferences (gets value at address)
3. ✅ Pointers allow modifying original data
4. ✅ nil pointers have no value
5. ✅ Always check for nil
6. ✅ Go handles memory safely
7. ✅ Use pointers for large structs

---

## Slide 20: Pointer Cheat Sheet

| Task | Syntax |
|------|--------|
| Get address | `ptr := &x` |
| Dereference | `value := *ptr` |
| Modify via pointer | `*ptr = newValue` |
| Check nil | `if ptr != nil` |
| Pointer type | `*int`, `*Person` |
| Create pointer | `ptr := new(Type)` |

---

## Slide 21: Real-World Example

**JSON unmarshaling:**
```go
type Config struct {
    Host string
    Port int
}

func loadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err  // Return nil on error
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}
```

---

## Slide 22: Homework

**Implement a Binary Tree:**

```go
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// Methods to implement:
// Insert(value int) - Add node
// Search(value int) *TreeNode - Find node
// InOrder() []int - Return values in order
// Height() int - Return tree height
```

**Learn about tree traversal algorithms!**

---

## Slide 23: What's Next?

**Lecture 10: Methods and Interfaces**
- Methods on types
- Interface definitions
- Type satisfaction
- Polymorphism

**Your code will be flexible and extensible! 🔌**

**See you next time! 👋**
