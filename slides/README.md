# Go Programming Course - HTML Slides

This directory contains a complete HTML-based slide presentation system for teaching Go programming.

## 📁 Files

- `index.html` - Main slide viewer (open this in a browser)
- `lecture_01_introduction.md` through `lecture_15_channels.md` - 15 lecture slide decks

## 🚀 How to Use

### Option 1: Open Directly (Simplest)
1. Open `index.html` in any modern web browser
2. Select a lecture from the dropdown
3. Navigate with arrow keys or buttons

### Option 2: Local Server (Recommended for development)
```bash
cd slides
python3 -m http.server 8000
# or
npx serve .
```

Then open http://localhost:8000

### Option 3: From the Repository Root
If you have the whole repo:
```bash
# From project root
open slides/index.html
```

## 🎓 Lecture Overview

| # | Title | Duration | Topics |
|---|-------|----------|--------|
| 01 | Welcome to Go! | 1.5h | Setup, Hello World, basic syntax |
| 02 | Variables and Types | 1.5h | Variables, constants, types, scope |
| 03 | Control Structures | 1.5h | If/else, switch, comparisons |
| 04 | Loops | 1.5h | For loops, range, break/continue |
| 05 | Functions | 1.5h | Parameters, returns, closures |
| 06 | Arrays and Slices | 1.5h | Fixed/dynamic arrays, slicing |
| 07 | Maps | 1.5h | Key-value pairs, lookups |
| 08 | Structs | 1.5h | Custom types, methods |
| 09 | Pointers | 1.5h | Memory addresses, dereferencing |
| 10 | Methods and Interfaces | 1.5h | OOP in Go, polymorphism |
| 11 | Packages and Modules | 1.5h | Code organization, imports |
| 12 | Error Handling | 1.5h | Robust error handling |
| 13 | File I/O | 1.5h | Reading/writing files, CSV, JSON |
| 14 | Goroutines | 1.5h | Concurrent programming |
| 15 | Channels | 1.5h | Communication between goroutines |

## 🎮 Navigation

- **⬅️ ➡️ Arrow Keys** - Navigate slides
- **Space/Enter** - Next slide
- **Dropdown Menu** - Switch lectures
- **Previous/Next Buttons** - Navigate within lecture

## ✨ Features

- 📱 **Responsive** - Works on mobile and desktop
- 🎨 **Syntax Highlighting** - Code blocks are styled
- 📊 **Tables** - Rendered nicely
- 🔗 **Internal Links** - Jump between sections
- ⌨️ **Keyboard Navigation** - Arrow keys work
- 💾 **No Dependencies** - Pure HTML/JS, no build step

## 📝 Slide Format

Slides are written in Markdown and separated by `---`:

```markdown
# Lecture Title

## Slide 1: Topic

Content here...

---

## Slide 2: Next Topic

More content...
```

## 🎨 Customization

Edit `index.html` to change:
- Colors (CSS variables)
- Font sizes
- Slide dimensions
- Animation effects

## 🐛 Troubleshooting

**Slides not loading?**
- Make sure you're accessing via HTTP (not file://)
- Use a local server (Option 2 above)
- Check browser console for errors

**Formatting issues?**
- The markdown parser is simple but handles most common cases
- Complex tables might need manual HTML

## 📚 Additional Resources

- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)

## 🎓 Teaching Tips

1. **Before Class**: Open the HTML file and test navigation
2. **During Class**: Use arrow keys to navigate smoothly
3. **Code Examples**: Click on code blocks to highlight them
4. **Exercises**: Pause on exercise slides for students to work
5. **Homework**: End each lecture with the homework slide

## 📄 License

These slides are part of the GolangIntroduction course.
