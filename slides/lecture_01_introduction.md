# Lecture 1: Welcome to Go! 🚀
## Your First Steps in Programming

---

## Slide 1: What is Go?

**Go (or Golang)** is a programming language created by Google in 2009.

**Think of it like this:**
- Just like you use English to talk to people, programmers use Go to talk to computers!
- Go is like a simple, clear language that computers understand really well.

**Why is Go special?**
- It's **fast** - like a sports car 🏎️
- It's **simple** - easy to read and write
- It's **powerful** - used by big companies like Google, Netflix, and Uber

---

## Slide 2: What Can You Build with Go?

**Real-world things made with Go:**
- 🌐 Websites and web servers
- 📱 Apps that need to handle millions of users
- 🛠️ Tools that help developers work faster
- 🤖 Programs that run in the cloud

**Famous examples:**
- Docker (helps run apps everywhere)
- Kubernetes (manages computer clusters)
- Terraform (builds cloud infrastructure)

---

## Slide 3: How Computers Think

**Imagine a computer is like a very obedient robot:**
- It does EXACTLY what you tell it to do
- It never gets tired or bored
- But it needs very clear instructions!

**Programming = Writing Instructions**
```
Human Idea → Go Code → Computer Does It!
```

**Example:**
- You: "Computer, add 2 + 2"
- Computer: "4" ✅

---

## Slide 4: Setting Up Your Computer

**What you need:**
1. **Go compiler** - translates your code to computer language
2. **Text editor** - where you write your code
3. **Terminal** - where you run your programs

**Installing Go:**
1. Visit https://golang.org/dl/
2. Download for your computer (Windows/Mac/Linux)
3. Follow the installation steps
4. Open terminal and type: `go version`

**You should see something like:** `go version go1.22.3`

---

## Slide 5: Your First Program - Hello World!

**Every programmer starts here!**

**The Code:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**What does each part mean?**

---

## Slide 6: Breaking Down Hello World

```go
package main      // This file is a program that can run

import "fmt"     // Bring in the "formatting" tools

func main() {     // This is where the program starts
    fmt.Println("Hello, World!")  // Print a message
}
```

**Like a recipe:**
1. `package main` = "This is the main dish"
2. `import "fmt"` = "I need the mixing bowl"
3. `func main()` = "Start cooking here!"
4. `fmt.Println()` = "Write something on the paper"

---

## Slide 7: Running Your Program

**Step 1: Create a file**
- Open your text editor
- Create a new file called `hello.go`
- Type the Hello World code

**Step 2: Save the file**
- Save it in a folder (like `my-first-go`)

**Step 3: Run it!**
Open your terminal and type:
```bash
go run hello.go
```

**You should see:**
```
Hello, World!
```

🎉 **Congratulations! You're now a programmer!**

---

## Slide 8: Understanding the Terminal

**The terminal is like a text-based control center:**

**Common commands:**
```bash
ls          # List files in current folder
cd myfolder # Change to a different folder
pwd         # Show where you are
mkdir new   # Create a new folder
go run file.go  # Run a Go program
```

**Think of it like:**
- Walking around your computer's folders
- Giving commands by typing instead of clicking

---

## Slide 9: The Go Playground

**Don't have Go installed yet? No problem!**

**Go Playground** is a website where you can write and run Go code:
- Visit: https://play.golang.org/
- Write code in your browser
- Click "Run" to see results
- Share code with others easily

**Perfect for:**
- Learning without installing
- Testing small ideas
- Sharing examples with friends

---

## Slide 10: Code Structure - The Big Picture

**Every Go program has this structure:**

```go
package main          // Package declaration

import "package"       // Import tools you need

func main() {          // Main function (entry point)
    // Your code here
}
```

**Analogy:**
- **Package** = The room you're working in
- **Import** = Tools you bring into the room
- **Main function** = The door where work begins

---

## Slide 11: Packages Explained

**What is a package?**
- A collection of Go files working together
- Like a folder of related tools

**Two types:**
1. **Executable packages** (`package main`)
   - Can be run as programs
   - Must have a `main` function

2. **Library packages** (`package something`)
   - Provide tools for other programs
   - Can't be run directly

---

## Slide 12: Importing Packages

**The `import` statement brings in tools:**

```go
import "fmt"        // Import one package

import (            // Import multiple packages
    "fmt"
    "time"
)
```

**Common packages:**
- `fmt` - Formatting and printing
- `time` - Working with dates and times
- `math` - Math functions
- `strings` - String manipulation

---

## Slide 13: The fmt Package

**`fmt` is your best friend for printing!**

**Common functions:**
```go
fmt.Println("Hello")           // Print with newline
fmt.Print("Hello")             // Print without newline
fmt.Printf("Hello %s", "World") // Formatted print
```

**Try these:**
```go
fmt.Println("I am learning Go!")
fmt.Println("This is fun!")
fmt.Print("Same line ")
fmt.Print("continues!")
```

---

## Slide 14: Comments - Notes for Humans

**Comments are notes that computers ignore:**

```go
// This is a single-line comment

/*
This is a
multi-line comment
*/

package main

func main() {
    // This prints a message
    fmt.Println("Hello!")  // See? It works!
}
```

**Why use comments?**
- Explain what code does
- Leave notes for yourself
- Help others understand your code

---

## Slide 15: Practice Time! 📝

**Exercise 1: Modify Hello World**
Change the message to print your name:
```go
fmt.Println("Hello, [Your Name]!")
```

**Exercise 2: Multiple Lines**
Print three different messages on separate lines.

**Exercise 3: Using Print**
Use `fmt.Print` to print "Go is " and "awesome!" on the same line.

---

## Slide 16: Common Mistakes

**Mistake 1: Forgetting quotes**
```go
fmt.Println(Hello)     // ❌ Wrong!
fmt.Println("Hello")  // ✅ Right!
```

**Mistake 2: Missing parentheses**
```go
fmt.Println "Hello"    // ❌ Wrong!
fmt.Println("Hello")  // ✅ Right!
```

**Mistake 3: Forgetting package main**
```go
// No package declaration!
func main() { }        // ❌ Won't work!
```

---

## Slide 17: Reading Error Messages

**When Go finds a problem, it tells you:**

```
hello.go:5:2: undefined: fmt
```

**What this means:**
- `hello.go` - The file with the problem
- `:5` - Line 5
- `:2` - Position 2 on that line
- `undefined: fmt` - You forgot to import "fmt"!

**Don't panic!** Errors are just the computer asking for clarification.

---

## Slide 18: Building vs Running

**Two ways to execute Go code:**

**1. Run (for testing):**
```bash
go run hello.go
```
- Compiles and runs immediately
- Good for development

**2. Build (for sharing):**
```bash
go build hello.go
./hello        # Run the compiled program
```
- Creates an executable file
- Can run without Go installed

---

## Slide 19: Your Project Structure

**Organize your learning:**
```
my-go-learning/
├── lecture-01/
│   └── hello.go
├── lecture-02/
│   └── variables.go
└── lecture-03/
    └── control.go
```

**Create a new folder for each lecture!**

---

## Slide 20: Go Formatting - gofmt

**Go has a built-in style guide!**

**Format your code automatically:**
```bash
gofmt -w hello.go    # Format a file
go fmt ./...          # Format all files
```

**Why format?**
- Makes code look consistent
- Easier to read
- Professional standard

**Your editor can auto-format on save!**

---

## Slide 21: Key Takeaways

**What you learned today:**

1. ✅ Go is a simple, fast programming language
2. ✅ Programs start with `package main` and `func main()`
3. ✅ Use `fmt.Println()` to print messages
4. ✅ Run programs with `go run filename.go`
5. ✅ Comments help explain your code
6. ✅ Errors are helpful messages, not scary!

---

## Slide 22: Vocabulary Review

**New words:**
- **Compiler** - Translates Go to computer language
- **Package** - A group of related Go files
- **Import** - Bring in external tools
- **Function** - A block of code that does something
- **String** - Text in quotes (like "Hello")
- **Terminal** - Command-line interface

---

## Slide 23: Homework 🏠

**Create a program that prints:**
```
My name is [Your Name]
I am learning Go programming
Today is [Current Day]
```

**Bonus:**
Add a simple ASCII art:
```
  /\\_/\\
 ( o.o )
  > ^ <
```

---

## Slide 24: Quick Review Questions

**Q1: What package do executable programs use?**
A) package hello
B) package main ✅
C) package run

**Q2: How do you print text in Go?**
A) print("Hello")
B) fmt.Println("Hello") ✅
C) console.log("Hello")

**Q3: What does // mean?**
A) Division
B) A comment ✅
C) Import statement

---

## Slide 25: What's Next?

**In Lecture 2, you'll learn:**
- Variables (storing information)
- Different types of data (numbers, text, true/false)
- Constants (values that never change)
- How to do math in Go!

**Get ready to make your programs smarter! 🧠**

---

## Slide 26: Resources

**Official:**
- golang.org - Official website
- tour.golang.org - Interactive tour
- play.golang.org - Online playground

**Community:**
- Go by Example - gobyexample.com
- Reddit: r/golang
- Discord: Gophers server

---

## Slide 27: Troubleshooting Tips

**Problem: "command not found: go"**
- Solution: Go isn't installed or not in PATH

**Problem: "undefined: fmt"**
- Solution: Add `import "fmt"`

**Problem: "syntax error"**
- Solution: Check for missing braces, parentheses, or quotes

**When stuck:**
1. Read the error message carefully
2. Check line numbers
3. Compare with working examples

---

## Slide 28: Programming Mindset

**Remember:**
- 🐛 Bugs are normal - everyone makes mistakes
- 🔍 Debugging is problem-solving
- 📚 Reading code is as important as writing it
- 🎯 Start simple, then add complexity
- 💪 Practice makes perfect!

**Famous quote:**
"The only way to learn a new programming language is by writing programs in it." - Dennis Ritchie

---

## Slide 29: Today's Challenge

**Create a "About Me" program:**

```go
package main

import "fmt"

func main() {
    fmt.Println("===== ABOUT ME =====")
    fmt.Println("Name: Your Name")
    fmt.Println("Age: Your Age")
    fmt.Println("Hobby: Your Hobby")
    fmt.Println("Favorite Language: Go!")
    fmt.Println("===================")
}
```

**Run it and show your friends!**

---

## Slide 30: See You Next Time! 👋

**You did it!**
- Wrote your first Go program
- Learned how Go works
- Set up your development environment

**Questions?**
- Review the slides
- Try the exercises
- Experiment with the code

**Next Lecture:** Variables and Types - Making your programs remember things!

**Keep coding! 🚀**
