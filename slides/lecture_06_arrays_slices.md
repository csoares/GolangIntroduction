# Lecture 6: Arrays and Slices 📊
## Storing Collections of Data

---

## Slide 1: Why Collections?

**What if you need to store:**
- Test scores for 30 students? 📝
- Names of all your friends? 👥
- Temperature readings for a week? 🌡️

**Without collections:**
```go
score1 := 85
score2 := 92
score3 := 78
// ... this is terrible!
```

**With collections:**
```go
scores := []int{85, 92, 78, 95, 88}
// Much better!
```

---

## Slide 2: Arrays vs Slices

**Go has two ways to store collections:**

**Arrays:**
- Fixed size (can't change)
- Size is part of the type
- `[5]int` is different from `[10]int`

**Slices:**
- Dynamic size (can grow/shrink)
- More flexible and common
- Built on top of arrays

**Rule of thumb:** Use slices unless you have a specific reason!

---

## Slide 3: Creating Arrays

**Fixed-size arrays:**

```go
// Declare with size
var scores [5]int  // [0, 0, 0, 0, 0]

// Initialize with values
scores := [5]int{85, 92, 78, 95, 88}

// Let Go count for you
scores := [...]int{85, 92, 78, 95, 88}  // Go counts: 5

// Partial initialization
scores := [5]int{85, 92}  // [85, 92, 0, 0, 0]
```

**Size is fixed forever!**

---

## Slide 4: Accessing Array Elements

**Use index (position) to access:**

```go
scores := [5]int{85, 92, 78, 95, 88}

first := scores[0]   // 85 (first element)
second := scores[1]  // 92 (second element)
last := scores[4]    // 88 (fifth element)
```

**Indexes start at 0!**

```
Index:    0    1    2    3    4
Value:   85   92   78   95   88
```

**Change values:**
```go
scores[2] = 90  // Change third score to 90
```

---

## Slide 5: Array Bounds

**Can't go outside the array!**

```go
scores := [3]int{10, 20, 30}

value := scores[5]   // ❌ PANIC! Index out of range
value := scores[-1]  // ❌ Won't compile!
```

**Safe access:**
```go
index := 5
if index >= 0 && index < len(scores) {
    value := scores[index]
}
```

**Length of array:**
```go
length := len(scores)  // Returns 3
```

---

## Slide 6: Iterating Over Arrays

**Use a for loop:**

```go
scores := [5]int{85, 92, 78, 95, 88}

// Method 1: By index
for i := 0; i < len(scores); i++ {
    fmt.Printf("Score %d: %d\n", i, scores[i])
}

// Method 2: Range (preferred)
for index, value := range scores {
    fmt.Printf("Score %d: %d\n", index, value)
}
```

**Range is cleaner and safer!**

---

## Slide 7: Creating Slices

**Slices are more flexible:**

```go
// From a literal
scores := []int{85, 92, 78}

// Empty slice
var empty []int  // nil slice

// Make with length
scores := make([]int, 5)  // [0, 0, 0, 0, 0]

// Make with length and capacity
scores := make([]int, 5, 10)  // Length 5, capacity 10
```

**No size in type - can grow!**

---

## Slide 8: Slice Operations

**Appending to slices:**

```go
scores := []int{85, 92}

// Add one element
scores = append(scores, 78)
// [85, 92, 78]

// Add multiple elements
scores = append(scores, 95, 88)
// [85, 92, 78, 95, 88]

// Add another slice
moreScores := []int{70, 75}
scores = append(scores, moreScores...)
// The ... unpacks the slice
```

---

## Slide 9: Slice Internals

**Slices have three parts:**
1. **Pointer** - To underlying array
2. **Length** - Number of elements
3. **Capacity** - Size of underlying array

```go
// [1, 2, 3, 4, 5]
//  ^
//  pointer
// len = 3, cap = 5

slice := []int{1, 2, 3, 4, 5}
sub := slice[1:4]  // [2, 3, 4]
```

**Capacity can be larger than length!**

---

## Slide 10: Slicing a Slice

**Create a new view into the slice:**

```go
fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}

// Syntax: slice[start:end]
some := fruits[1:3]     // ["banana", "cherry"]
firstTwo := fruits[:2]  // ["apple", "banana"]
lastTwo := fruits[3:]   // ["date", "elderberry"]
all := fruits[:]        // Copy of whole slice
```

**Important:** Slicing shares the underlying array!

```go
some[0] = "blueberry"
// fruits is now: ["apple", "blueberry", "cherry", ...]
```

---

## Slide 11: Copy Function

**Make an independent copy:**

```go
original := []int{1, 2, 3}
copy := make([]int, len(original))

copyCount := copy(copy, original)
// Returns number of elements copied
```

**Why copy?**
```go
original := []int{1, 2, 3}
view := original[:]

view[0] = 99
// original[0] is ALSO 99! They share memory.

// With copy, they're independent
```

---

## Slide 12: Delete from Slice

**No built-in delete, but we can do this:**

```go
// Delete element at index i
slice := append(slice[:i], slice[i+1:]...)

// Example:
numbers := []int{1, 2, 3, 4, 5}
// Delete element at index 2 (value 3)
numbers = append(numbers[:2], numbers[3:]...)
// Result: [1, 2, 4, 5]
```

**Order matters!**

---

## Slide 13: Insert into Slice

**Insert at specific position:**

```go
// Insert value at index i
slice = append(slice[:i], append([]int{value}, slice[i:]...)...)

// Example:
numbers := []int{1, 2, 4, 5}
// Insert 3 at index 2
numbers = append(numbers[:2], append([]int{3}, numbers[2:]...)...)
// Result: [1, 2, 3, 4, 5]
```

**Note:** This creates a temporary slice.

---

## Slide 14: Multi-Dimensional Arrays

**Arrays of arrays:**

```go
// 2D array (matrix)
var matrix [3][3]int
matrix[0][0] = 1
matrix[1][1] = 5
matrix[2][2] = 9

// Initialize
board := [3][3]string{
    {"X", "O", "X"},
    {"O", "X", "O"},
    {"X", " ", "X"},
}
```

**Can have 3D, 4D, etc.:**
```go
var cube [3][3][3]int
```

---

## Slide 15: Multi-Dimensional Slices

**Slices of slices:**

```go
// 2D slice
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access: matrix[row][column]
value := matrix[1][2]  // 6

// Iterate:
for i, row := range matrix {
    for j, value := range row {
        fmt.Printf("[%d][%d] = %d\n", i, j, value)
    }
}
```

---

## Slide 16: Common Operations

**Sum all elements:**
```go
sum := 0
for _, v := range numbers {
    sum += v
}
```

**Find max/min:**
```go
max := numbers[0]
for _, v := range numbers[1:] {
    if v > max {
        max = v
    }
}
```

**Check if contains:**
```go
found := false
for _, v := range numbers {
    if v == target {
        found = true
        break
    }
}
```

---

## Slide 17: Sorting

**Use the sort package:**

```go
import "sort"

numbers := []int{5, 2, 8, 1, 9}

sort.Ints(numbers)
// numbers is now: [1, 2, 5, 8, 9]

// Sort strings
names := []string{"Charlie", "Alice", "Bob"}
sort.Strings(names)
// ["Alice", "Bob", "Charlie"]

// Reverse order
sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
```

---

## Slide 18: Searching

**Binary search (requires sorted slice):**

```go
import "sort"

numbers := []int{1, 2, 5, 8, 9}

index := sort.Search(len(numbers), func(i int) bool {
    return numbers[i] >= 5
})
// index is 2 (where 5 is located)
```

**Simple linear search:**
```go
func contains(slice []int, target int) bool {
    for _, v := range slice {
        if v == target {
            return true
        }
    }
    return false
}
```

---

## Slide 19: Practice Exercise 1

**Reverse a slice:**

```go
func reverse(slice []int) []int {
    // Return a new slice with elements in reverse order
}

// Example:
// Input: [1, 2, 3, 4, 5]
// Output: [5, 4, 3, 2, 1]
```

**Bonus:**
- Do it in-place (modify original)
- Handle empty slices

---

## Slide 20: Practice Exercise 2

**Find duplicates:**

```go
func findDuplicates(numbers []int) []int {
    // Return a slice containing all duplicate values
    // Each duplicate appears only once in result
}

// Example:
// Input: [1, 2, 3, 2, 4, 3, 5]
// Output: [2, 3]
```

**Hints:**
- Use a map to track seen numbers
- Or sort first, then check adjacent

---

## Slide 21: Solutions

**Reverse (in-place):**
```go
func reverse(slice []int) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}
```

**Find duplicates:**
```go
func findDuplicates(numbers []int) []int {
    seen := make(map[int]bool)
    duplicates := []int{}

    for _, num := range numbers {
        if seen[num] {
            // Check if already in duplicates
            alreadyAdded := false
            for _, d := range duplicates {
                if d == num {
                    alreadyAdded = true
                    break
                }
            }
            if !alreadyAdded {
                duplicates = append(duplicates, num)
            }
        }
        seen[num] = true
    }
    return duplicates
}
```

---

## Slide 22: Slice Gotchas

**Gotcha 1: Nil vs Empty**
```go
var nilSlice []int     // nil
emptySlice := []int{}  // empty, not nil

fmt.Println(nilSlice == nil)   // true
fmt.Println(emptySlice == nil) // false

// Both have len 0, but behave differently in JSON
```

**Gotcha 2: Append to nil slice**
```go
var s []int  // nil
s = append(s, 1)  // Works! Becomes [1]
```

**Gotcha 3: Capacity growth**
```go
// Append may reallocate - capacity doubles when full
```

---

## Slide 23: Best Practices

**Do:**
- ✅ Use slices over arrays
- ✅ Preallocate capacity if you know size
- ✅ Use range for iteration
- ✅ Copy when you need independence
- ✅ Use `...` to unpack slices

**Don't:**
- ❌ Panic on index out of bounds
- ❌ Assume append doesn't reallocate
- ❌ Modify slice while ranging over it
- ❌ Return slices to internal arrays

---

## Slide 24: Quiz Time!

**Q1: What's the output?**
```go
s := []int{1, 2, 3}
t := s[1:]
t[0] = 99
fmt.Println(s)
```

**Q2: How to get the last element?**
```go
s := []int{10, 20, 30, 40}
last := ???
```

**Q3: What's the difference?**
```go
a := [5]int{1, 2, 3}
b := []int{1, 2, 3}
```

---

## Slide 25: Key Takeaways

**What you learned:**

1. ✅ Arrays have fixed size
2. ✅ Slices can grow with append
3. ✅ Indexes start at 0
4. ✅ Slicing creates views (shares memory)
5. ✅ Use copy() for independent copies
6. ✅ Range is the safest way to iterate
7. ✅ Use sort and search packages

---

## Slide 26: Arrays vs Slices Cheat Sheet

| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed | Dynamic |
| Declaration | `[5]int` | `[]int` |
| Literal | `[...]int{1,2,3}` | `[]int{1,2,3}` |
| Append | ❌ No | ✅ Yes |
| Pass to function | Copy | Reference |
| Use when... | Fixed data | Most cases |

---

## Slide 27: Real-World Example

**Reading CSV data:**

```go
func readCSV(filename string) [][]string {
    file, _ := os.Open(filename)
    defer file.Close()

    reader := csv.NewReader(file)
    var records [][]string

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        records = append(records, record)
    }

    return records
}
```

---

## Slide 28: Homework

**Implement these functions:**

1. `filter(slice []int, condition func(int) bool) []int`
   - Return elements that match condition

2. `mapSlice(slice []int, transform func(int) int) []int`
   - Apply transform to each element

3. `reduce(slice []int, reducer func(int, int) int, initial int) int`
   - Combine all elements into one value

**Example:**
```go
evens := filter(numbers, func(n int) bool { return n%2 == 0 })
doubled := mapSlice(numbers, func(n int) int { return n * 2 })
sum := reduce(numbers, func(a, b int) int { return a + b }, 0)
```

---

## Slide 29: What's Next?

**Lecture 7: Maps**
- Key-value pairs
- Fast lookups
- No order guaranteed
- Delete and check existence

**Your code will handle fast lookups! 🔍**

**See you next time! 👋**
