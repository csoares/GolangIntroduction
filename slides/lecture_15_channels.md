# Lecture 15: Channels 📡
## Communication Between Goroutines

---

## Slide 1: Why Channels?

**Goroutines need to communicate:**
- Send data between goroutines 📤📥
- Synchronize execution ⏱️
- Share work 🔄

**Without channels:**
- Shared memory (complex, error-prone)
- Global variables (messy)
- Race conditions (dangerous!)

**Go philosophy:**
> "Don't communicate by sharing memory, share memory by communicating."

---

## Slide 2: What is a Channel?

**A channel is a pipe for data:**

```
Goroutine A -----> Channel -----> Goroutine B
              (data flows)
```

**Think of it like:**
- A mailbox 📬
- A conveyor belt 🏭
- A phone line ☎️

**Types of channels:**
- Unbuffered (synchronous)
- Buffered (asynchronous)

---

## Slide 3: Creating Channels

**Use `make` to create channels:**

```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel (capacity 10)
ch := make(chan int, 10)

// Channel of strings
messages := make(chan string)

// Channel of structs
users := make(chan User)
```

**Channel type: `chan ElementType`**

---

## Slide 4: Sending and Receiving

**Send data:**
```go
ch <- value  // Send value to channel
```

**Receive data:**
```go
value := <-ch  // Receive from channel
```

**Complete example:**
```go
func main() {
    messages := make(chan string)

    go func() {
        messages <- "Hello from goroutine!"
    }()

    msg := <-messages  // Wait and receive
    fmt.Println(msg)
}
```

---

## Slide 5: Unbuffered Channels

**Send and receive must happen together:**

```go
ch := make(chan int)  // No buffer

// Sender blocks until receiver is ready
go func() {
    fmt.Println("Sending...")
    ch <- 42  // Blocks here!
    fmt.Println("Sent!")
}()

time.Sleep(time.Second)
fmt.Println("Receiving...")
value := <-ch  // Unblocks sender
fmt.Println("Received:", value)
```

**Synchronization point!**

---

## Slide 6: Buffered Channels

**Send doesn't block until buffer is full:**

```go
ch := make(chan int, 3)  // Buffer size 3

ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
ch <- 3  // Doesn't block
ch <- 4  // Blocks! Buffer full

// Receive
fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
```

**Useful for:**
- Work queues
- Decoupling
- Rate limiting

---

## Slide 7: Closing Channels

**Close a channel when done:**

```go
func producer(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // No more values
}

func consumer(ch chan int) {
    for value := range ch {  // Range over channel
        fmt.Println(value)
    }
    // Loop ends when channel closed
}
```

**Only sender should close!**

---

## Slide 8: Checking if Channel Closed

**Comma ok idiom:**

```go
ch := make(chan int)
// ... send values and close ...

for {
    value, ok := <-ch
    if !ok {
        // Channel closed
        break
    }
    fmt.Println(value)
}
```

**ok is false when channel closed and empty**

---

## Slide 9: Range Over Channels

**Cleaner way to receive:**

```go
func main() {
    ch := make(chan int)

    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
        }
        close(ch)
    }()

    // Range receives until closed
    for value := range ch {
        fmt.Println(value)
    }
}
```

**Automatically exits when closed!**

---

## Slide 10: Select Statement

**Wait on multiple channels:**

```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("From ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("From ch2:", msg2)
case <-time.After(time.Second):
    fmt.Println("Timeout!")
}
```

**Like switch, but for channels!**

---

## Slide 11: Select with Default

**Non-blocking receive:**

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No message available")
}
```

**Useful for:**
- Polling
- Non-blocking operations
- Timeouts

---

## Slide 12: Directional Channels

**Restrict channel direction:**

```go
// Send-only channel
func producer(ch chan<- int) {
    ch <- 42
    // <-ch  // ❌ Compile error!
}

// Receive-only channel
func consumer(ch <-chan int) {
    value := <-ch
    // ch <- 1  // ❌ Compile error!
}

func main() {
    ch := make(chan int)
    go producer(ch)
    consumer(ch)
}
```

**Type safety!**

---

## Slide 13: Pipeline Pattern

**Chain channels together:**

```go
// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage
nums := generate(1, 2, 3, 4)
squares := square(nums)
for s := range squares {
    fmt.Println(s)
}
```

---

## Slide 14: Worker Pool Pattern

**Distribute work among workers:**

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 9 jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```

---

## Slide 15: Practice Exercise 1

**Concurrent Prime Filter:**

```go
// Stage 1: Generate numbers 2 to n
// Stage 2: Filter primes
// Stage 3: Print primes

func generate(n int) <-chan int
func filterPrimes(in <-chan int) <-chan int
func printPrimes(in <-chan int)

// Connect them in main()
```

**Use a pipeline of channels!**

---

## Slide 16: Practice Exercise 2

**Download Manager:**

```go
type Download struct {
    URL      string
    Priority int
}

func downloader(downloads <-chan Download, results chan<- string)

func main() {
    // Create download queue
    // Start multiple downloaders
    // Add downloads with priorities
    // Collect results
}
```

**Features:**
- Limited concurrent downloads
- Priority queue
- Progress reporting

---

## Slide 17: Solutions

**Prime filter:**
```go
func generate(n int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 2; i <= n; i++ {
            out <- i
        }
        close(out)
    }()
    return out
}

func filterPrimes(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if isPrime(n) {
                out <- n
            }
        }
        close(out)
    }()
    return out
}
```

---

## Slide 18: Common Mistakes

**Mistake 1: Sending to closed channel**
```go
close(ch)
ch <- 42  // ❌ Panic!
```

**Mistake 2: Closing received channel**
```go
func consumer(ch <-chan int) {
    close(ch)  // ❌ Compile error!
}
```

**Mistake 3: Goroutine leak**
```go
ch := make(chan int)  // Unbuffered
go func() {
    ch <- 42  // Blocks forever if no receiver!
}()
// Forgot to receive!
```

---

## Slide 19: Best Practices

**Do:**
- ✅ Close channels from sender
- ✅ Use range to receive until closed
- ✅ Use select for multiple channels
- ✅ Use buffered channels for decoupling
- ✅ Use directional types for clarity

**Don't:**
- ❌ Send to closed channels
- ❌ Close received channels
- ❌ Leave goroutines blocked forever
- ❌ Use channels for everything (overkill)

---

## Slide 20: Quiz Time!

**Q1: What does <-ch do?**
A) Send to channel
B) Receive from channel
C) Close channel

**Q2: Who should close a channel?**
A) Receiver
B) Sender
C) Anyone

**Q3: What is select used for?**
A) Switch statements
B) Multiple channel operations
C) Type selection

---

## Slide 21: Key Takeaways

**What you learned:**

1. ✅ Channels communicate between goroutines
2. ✅ `ch <- value` sends
3. ✅ `<-ch` receives
4. ✅ Buffered channels don't block immediately
5. ✅ Close channels when done
6. ✅ Range over channels
7. ✅ Select for multiple channels

---

## Slide 22: Channel Cheat Sheet

| Task | Syntax |
|------|--------|
| Create | `ch := make(chan T)` |
| Buffered | `ch := make(chan T, n)` |
| Send | `ch <- value` |
| Receive | `value := <-ch` |
| Close | `close(ch)` |
| Range | `for v := range ch` |
| Select | `select { case ... }` |

---

## Slide 23: Real-World Example

**Context cancellation:**
```go
func worker(ctx context.Context, jobs <-chan int) {
    for {
        select {
        case job := <-jobs:
            process(job)
        case <-ctx.Done():
            return  // Cancelled!
        }
    }
}

// Cancel after timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

go worker(ctx, jobs)
```

---

## Slide 24: Final Project

**Build a Concurrent Web Crawler:**

**Requirements:**
1. Start from a URL
2. Fetch pages concurrently
3. Extract links
4. Follow links (with depth limit)
5. Avoid visiting same URL twice
6. Report statistics

**Techniques:**
- Goroutines for fetching
- Channels for communication
- Mutex for visited set
- WaitGroup for synchronization
- Select for timeout

---

## Slide 25: Course Summary

**What you learned in this course:**

1. ✅ Go basics and syntax
2. ✅ Variables and types
3. ✅ Control structures
4. ✅ Loops and iteration
5. ✅ Functions and methods
6. ✅ Arrays, slices, maps
7. ✅ Structs and interfaces
8. ✅ Pointers
9. ✅ Packages and modules
10. ✅ Error handling
11. ✅ File I/O
12. ✅ Goroutines
13. ✅ Channels

---

## Slide 26: Where to Go Next?

**Continue learning:**
- Standard library deep dive
- Testing and benchmarking
- Building web servers
- Working with databases
- gRPC and microservices
- Kubernetes and cloud

**Resources:**
- gobyexample.com
- golang.org/doc
- Awesome Go (GitHub)
- Go conferences

---

## Slide 27: Final Tips

**Keep practicing:**
- Build projects
- Read others' code
- Contribute to open source
- Join Go communities

**Remember:**
- Simplicity over cleverness
- Explicit is better than implicit
- Errors are values
- Concurrency is powerful but tricky

---

## Slide 28: Thank You!

**Congratulations!** 🎉

You've completed the Go Programming Course!

**You now know:**
- How to write Go programs
- Go idioms and best practices
- Concurrent programming
- How to build real applications

**Keep coding!** 💻

**Questions?**
- Review the slides
- Practice the exercises
- Build something cool!

---

## Slide 29: Certificate of Completion

**🏆 Go Programming Course 🏆**

This certifies that you have completed
15 lectures of comprehensive Go training

**Topics mastered:**
- Go fundamentals
- Data structures
- Object-oriented patterns
- Concurrent programming
- Real-world applications

**Congratulations!**

---

## Slide 30: Final Challenge

**Build Your Best Project:**

Ideas:
- REST API server
- CLI tool
- Chat application
- Game
- Data processor
- Automation tool

**Requirements:**
- Use what you learned
- Write tests
- Handle errors
- Document your code
- Share with others!

**Good luck! 🚀**
