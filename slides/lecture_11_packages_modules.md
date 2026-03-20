# Lecture 11: Packages and Modules 📦
## Organizing and Reusing Code

---

## Slide 1: Why Organize Code?

**Imagine your room:**
- Without organization: Everything in one pile! 😱
- With organization: Clothes in closet, books on shelf, tools in garage

**Code organization:**
- Easier to find things
- Reusable components
- Team collaboration
- Clear boundaries

**In Go:**
- **Packages** - Group related files
- **Modules** - Manage dependencies

---

## Slide 2: What is a Package?

**A package is a collection of Go files:**

```
calculator/
├── basic.go       // package calculator
├── advanced.go    // package calculator
└── calculator_test.go  // package calculator
```

**All files in a folder = same package**

**Purpose:**
- Organize code logically
- Control visibility (exported/unexported)
- Reuse across projects

---

## Slide 3: Package Declaration

**Every Go file starts with a package:**

```go
// calculator/basic.go
package calculator

func Add(a, b float64) float64 {
    return a + b
}
```

```go
// calculator/advanced.go
package calculator

import "math"

func Power(a, b float64) float64 {
    return math.Pow(a, b)
}
```

**Files in same package can access each other freely!**

---

## Slide 4: Package Naming

**Rules:**
- Lowercase only
- No underscores (usually)
- Short but clear
- Matches folder name

**Good names:**
- `calculator`
- `http`
- `json`
- `strings`

**Bad names:**
- `MyPackage` (uppercase)
- `calculator_package` (underscore)
- `pkg` (too vague)

---

## Slide 5: Exported vs Unexported

**Capital = Exported (public):**
```go
package calculator

func Add(a, b float64) float64 {  // Exported!
    return a + b
}

func subtract(a, b float64) float64 {  // Unexported
    return a - b
}
```

**Usage from another package:**
```go
import "calculator"

calculator.Add(2, 3)       // ✅ Works!
calculator.subtract(5, 2)  // ❌ Won't compile!
```

---

## Slide 6: Creating Your Package

**Step 1: Create folder:**
```bash
mkdir calculator
cd calculator
```

**Step 2: Create files:**
```go
// basic.go
package calculator

func Add(a, b float64) float64 {
    return a + b
}

func Subtract(a, b float64) float64 {
    return a - b
}
```

**Step 3: Use in main:**
```go
package main

import "mymodule/calculator"

func main() {
    result := calculator.Add(2, 3)
}
```

---

## Slide 7: Importing Packages

**Import paths:**

```go
// Standard library (no module needed)
import "fmt"
import "strings"
import "net/http"

// Your module
import "github.com/user/project/calculator"

// Third-party
import "github.com/gin-gonic/gin"
```

**Multiple imports:**
```go
import (
    "fmt"
    "strings"
    "mymodule/calculator"
)
```

---

## Slide 8: What is a Module?

**A module is a collection of packages:**
- Has a name (module path)
- Has a version
- Tracks dependencies

**Defined in go.mod:**
```
module github.com/user/myproject

go 1.22

require (
    github.com/some/library v1.2.3
)
```

**Creating a module:**
```bash
go mod init github.com/user/myproject
```

---

## Slide 9: go.mod File

**Example go.mod:**
```go
module golang-introduction

go 1.22

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/stretchr/testify v1.8.4
)

require (
    github.com/davecgh/go-spew v1.1.1 // indirect
    github.com/pmezard/go-difflib v1.0.0 // indirect
    gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

**Parts:**
- Module name
- Go version
- Direct dependencies
- Indirect dependencies

---

## Slide 10: go.sum File

**go.sum = Checksums for security:**
```
github.com/gin-gonic/gin v1.9.1 h1:Qat...Xyz=
github.com/gin-gonic/gin v1.9.1/go.mod h1:AbB...=
```

**Purpose:**
- Verify downloaded code
- Prevent tampering
- Reproducible builds

**Never edit manually!**

---

## Slide 11: Working with Modules

**Common commands:**

```bash
# Initialize module
go mod init myproject

# Download dependencies
go mod download

# Add dependency (automatic when importing)
go get github.com/gin-gonic/gin

# Update dependency
go get -u github.com/gin-gonic/gin

# Tidy up dependencies
go mod tidy

# Vendor dependencies (copy to local)
go mod vendor
```

---

## Slide 12: Importing Local Packages

**Project structure:**
```
myproject/
├── go.mod
├── main.go
└── calculator/
    ├── basic.go
    └── advanced.go
```

**In main.go:**
```go
package main

import "myproject/calculator"

func main() {
    result := calculator.Add(2, 3)
}
```

**Module name + package path:**
```
myproject/calculator
^^^^^^^   ^^^^^^^^^^
module    package
```

---

## Slide 13: Third-Party Packages

**Finding packages:**
- pkg.go.dev - Official package registry
- GitHub - Most packages hosted there

**Installing:**
```bash
go get github.com/gin-gonic/gin
```

**Automatically:**
- Downloads code
- Updates go.mod
- Updates go.sum
- Makes available in imports

---

## Slide 14: Package Documentation

**Document with comments:**

```go
// Package calculator provides basic arithmetic operations.
// It includes addition, subtraction, multiplication, and division.
package calculator

// Add returns the sum of two numbers.
func Add(a, b float64) float64 {
    return a + b
}

// Divide returns the quotient of a and b.
// Returns an error if b is zero.
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

**Start with "Package name ..."**

---

## Slide 15: Internal Packages

**Hide implementation details:**

```
myproject/
├── api/
│   └── public.go       # Public API
├── internal/
│   └── helper/
│       └── private.go  # Can only be used within myproject
└── main.go
```

**Anything in `internal/` can't be imported by other modules!**

**Good for:**
- Implementation details
- Utilities
- Testing helpers

---

## Slide 16: Practice Exercise 1

**Create a Math Package:**

**Structure:**
```
mathpkg/
├── go.mod
├── math.go
├── stats/
│   └── stats.go
└── geom/
    └── geometry.go
```

**Implement:**
- Basic operations (Add, Sub, Mul, Div)
- Statistics (Mean, Median, Mode)
- Geometry (Area, Perimeter)

**Create a main program that uses all!**

---

## Slide 17: Practice Exercise 2

**Create a String Utilities Package:**

```go
package strutil

// Reverse a string
func Reverse(s string) string

// Check if palindrome
func IsPalindrome(s string) bool

// Count words
func WordCount(s string) int

// Truncate with ellipsis
func Truncate(s string, maxLen int) string
```

**Write tests for each function!**

---

## Slide 18: Solutions

**Basic package structure:**
```go
// mathpkg/basic.go
package math

func Add(a, b float64) float64 {
    return a + b
}

// mathpkg/stats/stats.go
package stats

func Mean(numbers []float64) float64 {
    sum := 0.0
    for _, n := range numbers {
        sum += n
    }
    return sum / float64(len(numbers))
}
```

**Usage:**
```go
import (
    "myproject/math"
    "myproject/math/stats"
)
```

---

## Slide 19: Versioning

**Semantic versioning:**
```
v1.2.3
│ │ │
│ │ └── Patch (bug fixes)
│ └──── Minor (new features, backward compatible)
└────── Major (breaking changes)
```

**In go.mod:**
```go
require (
    github.com/lib/pq v1.10.9
    github.com/redis/go-redis/v9 v9.0.5
)
```

**Major versions:**
- v0.x.x - Unstable
- v1.x.x - Stable
- v2+ - In module path: `/v2`

---

## Slide 20: Common Mistakes

**Mistake 1: Import cycle**
```go
// package a imports b
// package b imports a
// ❌ Compile error!
```

**Mistake 2: Wrong import path**
```go
import "calculator"  // ❌ Missing module prefix!
import "myproject/calculator"  // ✅ Correct
```

**Mistake 3: Not initializing module**
```bash
go run main.go  # ❌ No go.mod!
go mod init myproject  # ✅ First
```

---

## Slide 21: Best Practices

**Do:**
- ✅ Keep packages focused (single responsibility)
- ✅ Document exported items
- ✅ Use internal for implementation
- ✅ Version your modules
- ✅ Use go mod tidy

**Don't:**
- ❌ Create circular imports
- ❌ Put everything in one package
- ❌ Export everything (be selective)
- ❌ Forget to commit go.sum

---

## Slide 22: Quiz Time!

**Q1: Which is exported?**
```go
func calculate()   // A
func Calculate()   // B
```

**Q2: What file defines a module?**
A) package.go
B) go.mod
C) module.go

**Q3: How do you add a dependency?**
A) go add
B) go get
C) go install

---

## Slide 23: Key Takeaways

**What you learned:**

1. ✅ Packages group related Go files
2. ✅ Capitalize to export, lowercase to hide
3. ✅ Modules manage dependencies
4. ✅ go.mod defines module
5. ✅ go.sum ensures security
6. ✅ Use go get to add packages
7. ✅ Document with comments

---

## Slide 24: Module Cheat Sheet

| Task | Command |
|------|---------|
| Init module | `go mod init name` |
| Download deps | `go mod download` |
| Add package | `go get url` |
| Update all | `go get -u ./...` |
| Tidy | `go mod tidy` |
| Vendor | `go mod vendor` |

---

## Slide 25: Real-World Example

**Project structure:**
```
myapp/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── db/
│   ├── auth/
│   └── cache/
├── pkg/
│   ├── models/
│   └── utils/
├── go.mod
└── README.md
```

**Standard Go layout!**

---

## Slide 26: Homework

**Build a CLI Toolkit:**

Create a module with multiple packages:
- `text` - Text manipulation
- `math` - Math utilities
- `time` - Time helpers
- `file` - File operations

**Requirements:**
- Each package has 3+ functions
- Each function has tests
- Document all exported functions
- Create a CLI that uses all packages

---

## Slide 27: What's Next?

**Lecture 12: Error Handling**
- Proper error handling
- Custom error types
- Error wrapping
- Best practices

**Your code will be robust! 💪**

**See you next time! 👋**
