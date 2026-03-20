# Lecture 14: Goroutines 🚀
## Concurrent Programming in Go

---

## Slide 1: What is Concurrency?

**Concurrency = Doing multiple things at once:**

**Real-world examples:**
- Cooking dinner while listening to music 🍳🎵
- Driving while talking to a passenger 🚗💬
- Browser loading multiple tabs 🌐

**In programming:**
- Downloading files while showing progress
- Handling multiple user requests
- Processing data in parallel

**Go makes concurrency EASY!**

---

## Slide 2: Goroutines

**A goroutine is a lightweight thread:**

**Normal function call:**
```go
function()  // Wait for it to finish
```

**Goroutine (async):**
```go
go function()  // Run in background, continue immediately!
```

**The `go` keyword starts a goroutine:**
```go
func sayHello() {
    fmt.Println("Hello!")
}

func main() {
    go sayHello()  // Starts in background
    fmt.Println("Main continues...")
    time.Sleep(time.Second)  // Wait for goroutine
}
```

---

## Slide 3: Goroutines are Cheap

**Goroutines vs OS Threads:**

| Feature | OS Thread | Goroutine |
|---------|-----------|-----------|
| Size | 1-2 MB | 2 KB |
| Startup | Slow | Fast |
| Switching | Slow | Fast |
| Max count | Thousands | Millions |

**You can have millions of goroutines!**

---

## Slide 4: Anonymous Goroutines

**Start a goroutine without a named function:**

```go
func main() {
    // Anonymous function as goroutine
    go func() {
        fmt.Println("Running in background!")
    }()

    // With parameters
    go func(msg string) {
        fmt.Println(msg)
    }("Hello from goroutine!")

    time.Sleep(time.Second)
}
```

**Very common pattern!**

---

## Slide 5: The Problem - Race Conditions

**Goroutines run independently:**

```go
func main() {
    counter := 0

    for i := 0; i < 1000; i++ {
        go func() {
            counter++  // ❌ Race condition!
        }()
    }

    time.Sleep(time.Second)
    fmt.Println(counter)  // Probably not 1000!
}
```

**Multiple goroutines accessing same data = DANGER!**

---

## Slide 6: WaitGroups

**Wait for goroutines to finish:**

```go
import "sync"

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)  // Add to counter

        go func(id int) {
            defer wg.Done()  // Decrement when done
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }

    wg.Wait()  // Wait for all
    fmt.Println("All workers done!")
}
```

---

## Slide 7: WaitGroup Pattern

**Standard pattern:**
```go
var wg sync.WaitGroup

// 1. Add before starting
wg.Add(1)
go func() {
    defer wg.Done()  // 2. Done when finished
    // Do work
}()

// 3. Wait at the end
wg.Wait()
```

**Always use defer for Done()!**

---

## Slide 8: Practical Example

**Download multiple files:**

```go
func download(url string, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Downloading %s...\n", url)
    time.Sleep(2 * time.Second)  // Simulate download
    fmt.Printf("Finished %s\n", url)
}

func main() {
    urls := []string{
        "file1.txt",
        "file2.txt",
        "file3.txt",
    }

    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go download(url, &wg)
    }

    wg.Wait()
    fmt.Println("All downloads complete!")
}
```

---

## Slide 9: Mutex - Mutual Exclusion

**Protect shared data:**

```go
import "sync"

func main() {
    var counter int
    var mu sync.Mutex
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()    // Acquire lock
            counter++
            mu.Unlock()  // Release lock
        }()
    }

    wg.Wait()
    fmt.Println(counter)  // 1000! ✅
}
```

---

## Slide 10: Mutex Pattern

**Standard pattern:**
```go
mu.Lock()
// Critical section - only one goroutine at a time
mu.Unlock()
```

**Or use defer:**
```go
mu.Lock()
defer mu.Unlock()
// Critical section
```

**RWMutex for read-heavy workloads:**
```go
var rwmu sync.RWMutex

rwmu.RLock()   // Multiple readers
// Read data
rwmu.RUnlock()

rwmu.Lock()    // Exclusive write
// Write data
rwmu.Unlock()
```

---

## Slide 11: Race Detector

**Go can detect race conditions:**

```bash
go run -race myprogram.go
```

**Output if race detected:**
```
WARNING: DATA RACE
Read at 0x00c0000b2000 by goroutine 7:
  main.main.func1()
      main.go:12 +0x3e

Previous write at 0x00c0000b2000 by goroutine 6:
  main.main.func1()
      main.go:12 +0x5a
```

**Always use -race during development!**

---

## Slide 12: Practice Exercise 1

**Concurrent Prime Number Finder:**

```go
func isPrime(n int) bool {
    // Check if n is prime
}

func main() {
    numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

    // Check each number concurrently
    // Print results as they complete
    // Use WaitGroup to wait for all
}
```

**Bonus:** Count total primes found

---

## Slide 13: Practice Exercise 2

**Concurrent Web Scraper:**

```go
type Result struct {
    URL      string
    Status   int
    Duration time.Duration
}

func fetch(url string) Result {
    // Fetch URL and return Result
}

func main() {
    urls := []string{
        "https://google.com",
        "https://github.com",
        "https://stackoverflow.com",
    }

    // Fetch all URLs concurrently
    // Print results
}
```

---

## Slide 14: Solutions

**Prime finder:**
```go
func main() {
    numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
    var wg sync.WaitGroup
    var mu sync.Mutex
    primeCount := 0

    for _, n := range numbers {
        wg.Add(1)
        go func(num int) {
            defer wg.Done()
            if isPrime(num) {
                mu.Lock()
                primeCount++
                mu.Unlock()
                fmt.Printf("%d is prime\n", num)
            }
        }(n)
    }

    wg.Wait()
    fmt.Println("Total primes:", primeCount)
}
```

---

## Slide 15: Common Mistakes

**Mistake 1: Loop variable capture**
```go
for i := 0; i < 10; i++ {
    go func() {
        fmt.Println(i)  // ❌ All print same value!
    }()
}

// Fix: Pass as parameter
for i := 0; i < 10; i++ {
    go func(n int) {
        fmt.Println(n)  // ✅ Each has its own n
    }(i)
}
```

**Mistake 2: Forgetting WaitGroup**
```go
go doWork()  // Program exits before goroutine runs!
```

---

## Slide 16: Best Practices

**Do:**
- ✅ Use WaitGroup to wait for completion
- ✅ Protect shared data with Mutex
- ✅ Pass loop variables as parameters
- ✅ Use -race flag during development
- ✅ Keep goroutines small and focused

**Don't:**
- ❌ Share memory without synchronization
- ❌ Create unlimited goroutines
- ❌ Ignore race detector warnings
- ❌ Use sleep for synchronization

---

## Slide 17: Quiz Time!

**Q1: How do you start a goroutine?**
A) start function()
B) go function()
C) async function()

**Q2: What does WaitGroup do?**
A) Stops goroutines
B) Waits for goroutines to finish
C) Creates goroutines

**Q3: What protects shared data?**
A) Channel
B) Mutex
C) Array

---

## Slide 18: Key Takeaways

**What you learned:**

1. ✅ `go` starts a goroutine
2. ✅ Goroutines are lightweight
3. ✅ Use WaitGroup to wait
4. ✅ Use Mutex to protect data
5. ✅ Use -race to detect races
6. ✅ Pass loop vars as parameters
7. ✅ Avoid shared state when possible

---

## Slide 19: Goroutine Cheat Sheet

| Task | Syntax |
|------|--------|
| Start goroutine | `go function()` |
| Wait for group | `var wg sync.WaitGroup` |
| Add to group | `wg.Add(n)` |
| Mark done | `defer wg.Done()` |
| Wait | `wg.Wait()` |
| Lock mutex | `mu.Lock()` |
| Unlock mutex | `mu.Unlock()` |

---

## Slide 20: Real-World Example

**HTTP server handler:**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Each request runs in its own goroutine!
    // Go handles this automatically

    go logRequest(r)  // Log asynchronously

    // Process request
    data := processRequest(r)
    json.NewEncoder(w).Encode(data)
}
```

**Go HTTP servers are concurrent by default!**

---

## Slide 21: Homework

**Build a Concurrent Image Processor:**

**Features:**
1. Load multiple images concurrently
2. Apply filters (resize, grayscale)
3. Save processed images
4. Show progress bar
5. Report total time

**Requirements:**
- Use goroutines for parallel processing
- Use WaitGroup to wait
- Limit concurrent workers (e.g., max 4)
- Handle errors gracefully

---

## Slide 22: What's Next?

**Lecture 15: Channels**
- Communication between goroutines
- Channel operations
- Select statement
- Pipeline patterns

**Your goroutines will communicate! 📡**

**See you next time! 👋**
