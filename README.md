# Go-Lang-Cheats

Personal progress tracker and cheat sheet for the Go programming language.

---

## Table of Contents

1. [Progress Checklist](#progress-checklist)
2. [Personal Notes](#personal-notes)
3. [Configuration](#configuration)
4. [Usage & Snippets](#usage--snippets)
5. [License](#license)

---

## Progress Checklist

Track your learning with a simple checklist:

- [ ] **Basics** (syntax, variables, types)

---

## Personal Notes

Use this section to jot down your observations, tips, or any custom reminders as you progress:

<details>
<summary>Basic</summary>

- fmt is formatter  
- println auto-creates the new line  
- printf for using with vars  
- <code>:=</code> sugar syntax can be used for assigning vars directly (but not const)  
- defining int types explicitly for robust code
- the data type should be in same format when performing any  

</details>

---

## Configuration

```go
go init mod <project name> //to initialize the dir as a package
```

---

## Usage & Snippets

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go-Lang-Cheats!")
}
