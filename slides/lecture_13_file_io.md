# Lecture 13: File I/O 📁
## Reading and Writing Files

---

## Slide 1: Why File I/O?

**Programs need to:**
- Read configuration files ⚙️
- Save user data 💾
- Process CSV/JSON data 📊
- Log events 📝
- Read input from users ⌨️

**Without files:**
- Data disappears when program ends!
- Can't share data between runs
- Can't work with existing data

**Go makes file I/O simple!**

---

## Slide 2: The os Package

**Basic file operations:**

```go
import "os"

// Read entire file
data, err := os.ReadFile("file.txt")
// Returns []byte and error

// Write entire file
err := os.WriteFile("output.txt", []byte("Hello"), 0644)
// Returns error
```

**Quick and easy for small files!**

---

## Slide 3: Reading Files

**Read entire file at once:**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("hello.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Convert bytes to string
    content := string(data)
    fmt.Println(content)
}
```

**Good for:** Small files that fit in memory

---

## Slide 4: Writing Files

**Write entire file at once:**

```go
package main

import "os"

func main() {
    message := []byte("Hello, World!")

    // Write with permissions
    err := os.WriteFile("output.txt", message, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("File written!")
}
```

**Permission 0644:**
- Owner: read + write
- Others: read only

---

## Slide 5: Open and Close

**For more control, open a file:**

```go
file, err := os.Open("file.txt")  // Read-only
if err != nil {
    log.Fatal(err)
}
defer file.Close()  // Always close!

// Or for writing
file, err := os.Create("new.txt")  // Create/truncate
file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
```

**defer ensures file closes even if error!**

---

## Slide 6: Reading Line by Line

**For large files:**

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("large.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNum := 1

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("Line %d: %s\n", lineNum, line)
        lineNum++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
```

---

## Slide 7: Writing Line by Line

**Buffered writing:**

```go
package main

import (
    "bufio"
    "os"
)

func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)

    writer.WriteString("Line 1\n")
    writer.WriteString("Line 2\n")
    writer.WriteString("Line 3\n")

    writer.Flush()  // Don't forget!
}
```

---

## Slide 8: File Information

**Get file stats:**

```go
info, err := os.Stat("file.txt")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())        // bytes
fmt.Println("Permissions:", info.Mode())
fmt.Println("Modified:", info.ModTime())
fmt.Println("Is Directory:", info.IsDir())
```

**Check if file exists:**
```go
if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
    fmt.Println("File doesn't exist")
}
```

---

## Slide 9: Directory Operations

**Working with directories:**

```go
// Create directory
err := os.Mkdir("newfolder", 0755)

// Create all parent directories
err := os.MkdirAll("path/to/folder", 0755)

// Read directory
entries, err := os.ReadDir("myfolder")
for _, entry := range entries {
    fmt.Println(entry.Name())
    fmt.Println(entry.IsDir())
}

// Remove
os.Remove("file.txt")       // Single file
os.RemoveAll("folder")      // Directory and contents
```

---

## Slide 10: Path Operations

**Work with file paths:**

```go
import "path/filepath"

// Join paths (works on all OS)
fullPath := filepath.Join("folder", "subfolder", "file.txt")
// Windows: folder\subfolder\file.txt
// Unix: folder/subfolder/file.txt

// Get extension
ext := filepath.Ext("image.png")  // ".png"

// Get base name
base := filepath.Base("/path/to/file.txt")  // "file.txt"

// Get directory
dir := filepath.Dir("/path/to/file.txt")  // "/path/to"
```

---

## Slide 11: CSV Files

**Reading CSV:**

```go
import "encoding/csv"

file, _ := os.Open("data.csv")
defer file.Close()

reader := csv.NewReader(file)

// Read all records
records, _ := reader.ReadAll()

for _, record := range records {
    fmt.Println("Name:", record[0])
    fmt.Println("Age:", record[1])
}
```

**CSV = Comma Separated Values**
```
Name,Age,City
Alice,25,NYC
Bob,30,LA
```

---

## Slide 12: Writing CSV

```go
import "encoding/csv"

file, _ := os.Create("output.csv")
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

// Write header
writer.Write([]string{"Name", "Age", "City"})

// Write data
writer.Write([]string{"Alice", "25", "NYC"})
writer.Write([]string{"Bob", "30", "LA"})
```

**Always flush the writer!**

---

## Slide 13: JSON Files

**Reading JSON:**

```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Read file
data, _ := os.ReadFile("person.json")

// Parse JSON
var person Person
json.Unmarshal(data, &person)

fmt.Println(person.Name)
```

---

## Slide 14: Writing JSON

```go
import "encoding/json"

person := Person{Name: "Alice", Age: 25}

// Marshal to JSON
data, _ := json.MarshalIndent(person, "", "    ")

// Write to file
os.WriteFile("person.json", data, 0644)
```

**Output:**
```json
{
    "name": "Alice",
    "age": 25
}
```

---

## Slide 15: Practice Exercise 1

**Create a File Copy Program:**

```go
func CopyFile(src, dst string) error {
    // Should:
    // 1. Open source file
    // 2. Create destination file
    // 3. Copy contents
    // 4. Close both files
    // 5. Handle all errors
}
```

**Features:**
- Show progress (bytes copied)
- Preserve permissions
- Handle large files efficiently

---

## Slide 16: Practice Exercise 2

**Create a Contact Manager:**

**Store contacts in CSV:**
```csv
Name,Email,Phone
Alice,alice@example.com,555-1234
Bob,bob@example.com,555-5678
```

**Features:**
- Add contact
- List all contacts
- Search by name
- Delete contact
- Export to JSON

---

## Slide 17: Solutions

**File copy:**
```go
func CopyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    return err
}
```

**Use io.Copy for efficiency!**

---

## Slide 18: Common Mistakes

**Mistake 1: Not closing files**
```go
file, _ := os.Open("file.txt")
// ❌ File stays open!
```

**Mistake 2: Not checking errors**
```go
data, _ := os.ReadFile("file.txt")  // Dangerous!
```

**Mistake 3: Wrong paths**
```go
// Hardcoded paths don't work on all systems
os.ReadFile("C:\\Users\\name\\file.txt")  // Windows only

// Use filepath.Join instead
filepath.Join("folder", "file.txt")
```

---

## Slide 19: Best Practices

**Do:**
- ✅ Always use defer to close files
- ✅ Check every error
- ✅ Use filepath.Join for paths
- ✅ Use buffered I/O for large files
- ✅ Set appropriate permissions

**Don't:**
- ❌ Ignore file close
- ❌ Hardcode paths
- ❌ Read entire huge files
- ❌ Forget to flush writers

---

## Slide 20: Quiz Time!

**Q1: How do you check if file exists?**
A) os.Exists()
B) os.Stat() + os.IsNotExist()
C) os.FileExists()

**Q2: What does defer do?**
A) Deletes file
B) Delays execution
C) Opens file

**Q3: Which is best for large files?**
A) os.ReadFile
B) bufio.Scanner
C) ioutil.ReadAll

---

## Slide 21: Key Takeaways

**What you learned:**

1. ✅ os.ReadFile for small files
2. ✅ bufio.Scanner for large files
3. ✅ Always defer file.Close()
4. ✅ Check all errors
5. ✅ Use filepath.Join
6. ✅ CSV with encoding/csv
7. ✅ JSON with encoding/json

---

## Slide 22: File I/O Cheat Sheet

| Task | Method |
|------|--------|
| Read file | `os.ReadFile(path)` |
| Write file | `os.WriteFile(path, data, perm)` |
| Open file | `os.Open(path)` |
| Create file | `os.Create(path)` |
| Read lines | `bufio.Scanner` |
| Join paths | `filepath.Join(...)` |
| Check exists | `os.Stat()` + `os.IsNotExist()` |

---

## Slide 23: Real-World Example

**Configuration loader:**
```go
func LoadConfig(path string) (*Config, error) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return nil, fmt.Errorf("config not found: %s", path)
    }

    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }

    return &cfg, nil
}
```

---

## Slide 24: Homework

**Create a Log Analyzer:**

**Parse log files:**
```
[2024-01-15 10:23:45] INFO: Application started
[2024-01-15 10:24:12] ERROR: Database connection failed
[2024-01-15 10:24:15] INFO: Retrying connection...
```

**Features:**
1. Count errors, warnings, info
2. Find most common error
3. Filter by date range
4. Generate summary report
5. Export to CSV

---

## Slide 25: What's Next?

**Lecture 14: Goroutines**
- Concurrent programming
- Running tasks in parallel
- WaitGroups
- Basic concurrency patterns

**Your code will do multiple things at once! 🚀**

**See you next time! 👋**
