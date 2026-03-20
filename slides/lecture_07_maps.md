# Lecture 7: Maps - Key-Value Pairs 🔍
## Fast Lookups and Organized Data

---

## Slide 1: What is a Map?

**Think of a map like a real dictionary:**
- Word (key) → Definition (value) 📖
- You look up by the word, not by page number
- Fast to find what you need!

**In programming:**
- Store data as key-value pairs
- Keys are unique (like IDs)
- Values can be anything
- Super fast lookups!

**Real-world examples:**
- Phone book: Name → Phone number
- Dictionary: Word → Definition
- User database: ID → User info

---

## Slide 2: Creating Maps

**Three ways to create a map:**

```go
// Method 1: make function (most common)
scores := make(map[string]int)

// Method 2: literal with values
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}

// Method 3: empty literal
var scores map[string]int  // nil map
```

**Syntax:** `map[keyType]valueType`

---

## Slide 3: Adding and Updating

**Add key-value pairs:**

```go
scores := make(map[string]int)

// Add new entries
scores["Alice"] = 95
scores["Bob"] = 87
scores["Carol"] = 92

// Update existing
scores["Bob"] = 90  // Bob's score changes to 90

fmt.Println(scores)
// map[Alice:95 Bob:90 Carol:92]
```

**Keys must be unique!** Adding same key updates the value.

---

## Slide 4: Reading Values

**Access values by key:**

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
}

aliceScore := scores["Alice"]  // 95
bobScore := scores["Bob"]        // 87
```

**What if key doesn't exist?**
```go
daveScore := scores["Dave"]  // Returns 0 (zero value!)
```

**Important:** You get the zero value, not an error!

---

## Slide 5: Checking if Key Exists

**The "comma ok" idiom:**

```go
score, ok := scores["Alice"]
// score = 95, ok = true (key exists)

score, ok := scores["Dave"]
// score = 0, ok = false (key doesn't exist)
```

**Use it in if statements:**
```go
if score, ok := scores["Alice"]; ok {
    fmt.Println("Alice scored", score)
} else {
    fmt.Println("Alice not found")
}
```

**This is the Go way to check existence!**

---

## Slide 6: Deleting Entries

**Remove a key-value pair:**

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}

// Delete Bob's score
delete(scores, "Bob")

fmt.Println(scores)
// map[Alice:95 Carol:92]
```

**Delete is safe:**
- Works even if key doesn't exist
- No error returned
- Just does nothing

---

## Slide 7: Iterating Over Maps

**Use range (order is random!):**

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}

for name, score := range scores {
    fmt.Printf("%s scored %d\n", name, score)
}
// Output order is NOT guaranteed!
```

**Keys only:**
```go
for name := range scores {
    fmt.Println(name)
}
```

**Values only:**
```go
for _, score := range scores {
    fmt.Println(score)
}
```

---

## Slide 8: Map Size

**Get number of entries:**

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}

count := len(scores)  // 3
```

**No built-in max or min for maps!**

**Find max score:**
```go
maxScore := 0
for _, score := range scores {
    if score > maxScore {
        maxScore = score
    }
}
```

---

## Slide 9: Valid Key Types

**Keys must be comparable:**

**✅ Valid key types:**
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `string`
- `float32`, `float64`
- Pointers
- Structs with comparable fields
- Arrays

**❌ Invalid key types:**
- Slices
- Maps
- Functions

---

## Slide 10: Maps of Maps

**Nested maps:**

```go
// Students by class
classes := map[string]map[string]int{
    "Math": {
        "Alice": 95,
        "Bob":   87,
    },
    "Science": {
        "Alice": 92,
        "Bob":   88,
    },
}

// Access nested data
aliceMath := classes["Math"]["Alice"]  // 95
```

**Can go as deep as you need!**

---

## Slide 11: Maps of Slices

**Group items together:**

```go
// Group students by grade
byGrade := map[string][]string{
    "A": {"Alice", "Carol"},
    "B": {"Bob", "David"},
    "C": {"Eve"},
}

// Add to a group
byGrade["A"] = append(byGrade["A"], "Frank")

// Access
fmt.Println(byGrade["A"])  // [Alice Carol Frank]
```

**Very useful for organizing data!**

---

## Slide 12: Common Patterns

**Count occurrences:**
```go
words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
counts := make(map[string]int)

for _, word := range words {
    counts[word]++
}
// counts: map[apple:3 banana:2 cherry:1]
```

**Set (unique items):**
```go
unique := make(map[string]bool)
unique["Alice"] = true
unique["Bob"] = true

// Check if in set
if unique["Alice"] {
    fmt.Println("Alice is in the set")
}
```

---

## Slide 13: Practice Exercise 1

**Word frequency counter:**

```go
func wordFrequency(text string) map[string]int {
    // Count how many times each word appears
    // Hint: use strings.Fields() to split text
}

// Example:
// Input: "the quick brown fox jumps over the lazy dog"
// Output: map[the:2 quick:1 brown:1 fox:1 ...]
```

**Bonus:**
- Ignore case (convert to lowercase)
- Remove punctuation

---

## Slide 14: Practice Exercise 2

**Two-sum problem:**

```go
func twoSum(numbers []int, target int) []int {
    // Return indices of two numbers that add up to target
    // Use a map for O(n) solution
}

// Example:
// numbers: [2, 7, 11, 15], target: 9
// Output: [0, 1] (because 2 + 7 = 9)
```

**Hint:** Store complement (target - num) in map

---

## Slide 15: Solutions

**Word frequency:**
```go
import "strings"

func wordFrequency(text string) map[string]int {
    words := strings.Fields(strings.ToLower(text))
    freq := make(map[string]int)

    for _, word := range words {
        // Remove punctuation (simplified)
        word = strings.Trim(word, ".,!?;:")
        freq[word]++
    }
    return freq
}
```

**Two-sum:**
```go
func twoSum(numbers []int, target int) []int {
    seen := make(map[int]int)  // value -> index

    for i, num := range numbers {
        complement := target - num
        if j, ok := seen[complement]; ok {
            return []int{j, i}
        }
        seen[num] = i
    }
    return nil
}
```

---

## Slide 16: Nil Maps

**Be careful with nil maps:**

```go
var scores map[string]int  // nil

// Reading is OK
score := scores["Alice"]  // Returns 0

// Writing PANICS!
scores["Alice"] = 95  // ❌ Runtime panic!
```

**Always initialize before use:**
```go
scores := make(map[string]int)  // ✅ Empty but usable
scores["Alice"] = 95            // ✅ Works!
```

---

## Slide 17: Maps vs Structs

**When to use what:**

**Use Maps when:**
- Keys are dynamic/unknown at compile time
- You need to iterate over keys
- You need to check if key exists
- Keys are all the same type

**Use Structs when:**
- Fields are fixed and known
- You want type safety
- Different fields have different types
- You want methods attached

---

## Slide 18: Best Practices

**Do:**
- ✅ Always check if key exists when it matters
- ✅ Initialize maps with make()
- ✅ Use appropriate key types
- ✅ Preallocate if you know size

**Don't:**
- ❌ Rely on map iteration order
- ❌ Use slices or maps as keys
- ❌ Write to nil maps
- ❌ Assume zero value means key doesn't exist

---

## Slide 19: Common Mistakes

**Mistake 1: Assuming order**
```go
for k, v := range myMap {
    // Order is random! Don't rely on it!
}
```

**Mistake 2: Not checking existence**
```go
score := scores["Dave"]  // Might be 0 or might not exist!
// Use comma ok instead
```

**Mistake 3: Concurrent access**
```go
// Maps are NOT thread-safe!
// Use sync.RWMutex for concurrent access
```

---

## Slide 20: Quiz Time!

**Q1: What's the output?**
```go
m := map[string]int{"a": 1, "b": 2}
delete(m, "c")
fmt.Println(len(m))
```

**Q2: Fix the bug:**
```go
var scores map[string]int
scores["Alice"] = 95
```

**Q3: How do you check if "Bob" exists?**
```go
scores := map[string]int{"Alice": 95}
// ???
```

---

## Slide 21: Key Takeaways

**What you learned:**

1. ✅ Maps store key-value pairs
2. ✅ Keys must be unique and comparable
3. ✅ Use `make()` to create maps
4. ✅ Check existence with comma ok
5. ✅ Delete with `delete()` function
6. ✅ Iteration order is random
7. ✅ Nil maps panic on write

---

## Slide 22: Map Cheat Sheet

| Operation | Syntax |
|-----------|--------|
| Create | `make(map[K]V)` |
| Add/Update | `m[key] = value` |
| Read | `value := m[key]` |
| Check exists | `v, ok := m[key]` |
| Delete | `delete(m, key)` |
| Length | `len(m)` |
| Iterate | `for k, v := range m` |

---

## Slide 23: Real-World Example

**HTTP headers:**
```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // Headers are map[string][]string
    contentType := r.Header["Content-Type"]
    userAgent := r.Header["User-Agent"]

    // Set header
    w.Header()["X-Custom"] = []string{"MyValue"}
}
```

**Configuration:**
```go
config := map[string]string{
    "database_url": "localhost:5432",
    "api_key":      "secret123",
    "timeout":      "30s",
}
```

---

## Slide 24: Homework

**Build a simple phone book:**

1. Add contact (name, phone)
2. Look up contact by name
3. Delete contact
4. List all contacts
5. Search by partial name

**Bonus:**
- Save to file (JSON)
- Load from file
- Handle duplicates

---

## Slide 25: What's Next?

**Lecture 8: Structs**
- Custom data types
- Grouping related data
- Methods on structs
- Embedding

**Your code will be organized and type-safe! 🏗️**

**See you next time! 👋**
