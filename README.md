[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/must/release.yml?branch=main&label=BUILD)](https://github.com/yylego/must/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/must)](https://pkg.go.dev/github.com/yylego/must)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/must/main.svg)](https://coveralls.io/github/yylego/must?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/must.svg)](https://github.com/yylego/must/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/must)](https://goreportcard.com/report/github.com/yylego/must)

# must

Simple assertion utilities with panic-on-failure semantics designed to reduce boilerplate code in Go.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

🎯 **Panic-on-Failure Validation**: Clean assertions with automatic panic on failure when conditions are violated
⚡ **Type-Safe Generics**: Comprehensive support with Go generics spanning assertion types
🔄 **Stack Frame Adjustment**: Precise panic location through intelligent skip configurations
🌍 **Structured Logging**: Deep integration with zap providing detailed panic context
📋 **Domain-Specific Packages**: Specialized utilities with numeric, string, slice, and map operations

---

## Installation

```bash
go get github.com/yylego/must
```

---

## Quick Start

### Example 1: Basic Assertions

```go
package main

import (
	"fmt"

	"github.com/yylego/must"
)

func main() {
	fmt.Println("=== Demo 1: Basic Assertions ===")

	// Boolean assertion
	must.True(checkCondition())
	fmt.Println("✓ Boolean check passed")

	// Validate no errors
	must.Done(performOperation())
	fmt.Println("✓ No error")

	// Non-zero value
	count := getCount()
	must.Nice(count)
	fmt.Printf("✓ Valid count: %d\n", count)

	// Values match
	must.Equals("success", getStatus())
	fmt.Println("✓ Values match")

	// Slice operations
	items := getItems()
	must.Have(items)
	must.Length(items, 3)
	must.In("banana", items)
	fmt.Printf("✓ Slice validated: %v\n", items)

	// Pointer check
	account := getAccount()
	must.Full(account)
	fmt.Printf("✓ Pointer valid: %s\n", account.Name)

	fmt.Println("\n=== All checks passed! ===")
}

type Account struct{ Name string }

func checkCondition() bool    { return true }
func performOperation() error { return nil }
func getStatus() string       { return "success" }
func getCount() int           { return 42 }
func getItems() []string      { return []string{"apple", "banana", "orange"} }
func getAccount() *Account    { return &Account{Name: "test"} }
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

---

### Example 2: Rese Package Functions

```go
package main

import (
	"fmt"

	"github.com/yylego/must"
)

func main() {
	fmt.Println("=== Demo 2: Rese Package ===")

	// V1 - single value validation
	config := must.V1(readConfig())
	fmt.Printf("✓ Config: %s\n", config)

	// V2 - two-value validation
	width, height := must.V2(getDimensions())
	fmt.Printf("✓ Dimensions: %dx%d\n", width, height)

	// P1 - non-nil data validation
	admin := must.P1(findAdmin())
	fmt.Printf("✓ Admin: %s\n", admin.Name)

	// C1 - non-zero validation
	num := must.C1(getNum())
	fmt.Printf("✓ Num: %d\n", num)

	// Combined validations
	data := getData()
	must.Full(data)
	must.Nice(data.Score)
	must.Same(data.Status, "active")
	fmt.Printf("✓ Data: score=%d, status=%s\n", data.Score, data.Status)

	fmt.Println("\n=== All checks passed! ===")
}

type Admin struct{ Name string }
type Data struct {
	Score  int
	Status string
}

func readConfig() (string, error)      { return "v1.0", nil }
func getDimensions() (int, int, error) { return 1920, 1080, nil }
func findAdmin() (*Admin, error)       { return &Admin{"Alice"}, nil }
func getNum() (int, error)             { return 123, nil }
func getData() *Data                   { return &Data{95, "active"} }
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

---

### Example 3: Advanced Specialized Packages

```go
package main

import (
	"fmt"

	"github.com/yylego/must"
	"github.com/yylego/must/mustmap"
	"github.com/yylego/must/mustnum"
	"github.com/yylego/must/mustslice"
	"github.com/yylego/must/muststrings"
)

func main() {
	fmt.Println("=== Demo 3: Advanced Packages ===")

	// Numeric validations
	score := getScore()
	mustnum.Positive(score)
	mustnum.Gt(score, 60)
	fmt.Printf("✓ Score: %d\n", score)

	// Slice validations
	tags := getTags()
	mustslice.Have(tags)
	mustslice.Contains(tags, "go")
	fmt.Printf("✓ Tags: %v\n", tags)

	// Map validations
	config := getConfig()
	mustmap.Have(config)
	timeout := mustmap.Get(config, "timeout")
	fmt.Printf("✓ Timeout: %d\n", timeout)

	// String validations
	filename := getFilename()
	muststrings.HasSuffix(filename, ".pdf")
	muststrings.Contains(filename, "report")
	fmt.Printf("✓ Filename: %s\n", filename)

	// Complex scenario
	data := getAnalytics()
	must.Full(data)
	mustmap.Have(data.Metrics)
	fmt.Printf("✓ Analytics: %d metrics\n", len(data.Metrics))

	fmt.Println("\n=== All checks passed! ===")
}

type Analytics struct {
	Metrics map[string]float64
}

func getScore() int             { return 85 }
func getTags() []string         { return []string{"go", "test"} }
func getConfig() map[string]int { return map[string]int{"timeout": 30} }
func getFilename() string       { return "report.pdf" }
func getAnalytics() *Analytics {
	return &Analytics{Metrics: map[string]float64{"score": 87.5}}
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

---

## Core Assertions

Here are the core assertions in `must`, summarized in a table:

| **Function**                 | **Description**                                            | **Example**                   | **Notes**                           |
| ---------------------------- | ---------------------------------------------------------- | ----------------------------- | ----------------------------------- |
| **`True(v bool)`**           | Panics if `v` is false.                                    | `must.True(isValid)`          | Validates if `v` is `true`.         |
| **`Done(err error)`**        | Panics if `err` is not nil.                                | `must.Done(err)`              | Ensures no error occurred.          |
| **`Must(err error)`**        | Panics if `err` is not nil.                                | `must.Must(err)`              | Same as `Done`.                     |
| **`Nice(a V)`**              | Panics if `a` is zero.                                     | `must.Nice(value)`            | Ensures `a` is non-zero.            |
| **`Zero(a V)`**              | Panics if `a` is not zero.                                 | `must.Zero(value)`            | Ensures `a` is zero.                |
| **`None(a V)`**              | Panics if `a` is non-zero.                                 | `must.None(value)`            | Ensures `a` is zero.                |
| **`Null(v any)`**            | Panics if `v` is not `nil`.                                | `must.Null(ptr)`              | Ensures `v` is `nil`.               |
| **`Full(v any)`**            | Panics if `v` is `nil`.                                    | `must.Full(value)`            | Ensures `v` is non-`nil`.           |
| **`Equals(a, b V)`**         | Panics if `a` and `b` are not the same.                    | `must.Equals(a, b)`           | Checks if `a` equals `b`.           |
| **`Same(a, b V)`**           | Panics if `a` and `b` are not the same.                    | `must.Same(a, b)`             | Alias of `Equals`.                  |
| **`SameNice(a, b V)`**       | Panics if `a` and `b` are not the same, both non-zero.     | `must.SameNice(a, b)`         | Ensures same and non-zero.          |
| **`Sane(a, b V)`**           | Panics if `a` and `b` are not the same, both non-zero.     | `must.Sane(a, b)`             | Alias of `SameNice`.                |
| **`Diff(a, b V)`**           | Panics if `a` and `b` are the same.                        | `must.Diff(a, b)`             | Ensures values mismatch.            |
| **`Different(a, b V)`**      | Panics if `a` and `b` are the same.                        | `must.Different(a, b)`        | Alias of `Diff`.                    |
| **`Is(a, b V)`**             | Panics if `a` and `b` are not the same.                    | `must.Is(a, b)`               | Alias of `Equals`.                  |
| **`Ise(err, target error)`** | Panics if `err` does not match `target` using `errors.Is`. | `must.Ise(err, targetErr)`    | Matching like `errors.Is` function. |
| **`Ok(a V)`**                | Panics if `a` is zero.                                     | `must.Ok(value)`              | Ensures `a` is non-zero.            |
| **`OK(a V)`**                | Alias of `Ok`, checks non-zero value.                      | `must.OK(value)`              | Same as `Ok`.                       |
| **`TRUE(v bool)`**           | Panics if `v` is false.                                    | `must.TRUE(isValid)`          | Alias of `True`.                    |
| **`FALSE(v bool)`**          | Panics if `v` is true.                                     | `must.FALSE(isError)`         | Ensures `v` is `false`.             |
| **`False(v bool)`**          | Panics if `v` is true.                                     | `must.False(isError)`         | Same as `FALSE`.                    |
| **`Cause(err error)`**       | Panics if `err` is nil, returns the error.                 | `must.Cause(err)`             | Ensures error is present.           |
| **`Wrong(err error)`**       | Panics if `err` is nil.                                    | `must.Wrong(err)`             | Ensures error is present.           |
| **`Have(a []T)`**            | Panics if `a` has no elements.                             | `must.Have(slice)`            | Ensures `a` is not vacant.          |
| **`Length(a []T, n int)`**   | Panics if `a` length is not `n`.                           | `must.Length(slice, 3)`       | Ensures `a` length is `n`.          |
| **`Len(a []T, n int)`**      | Alias of `Length`, ensures `a` length is `n`.              | `must.Len(slice, 3)`          | Validates `a` length.               |
| **`In(v T, a []T)`**         | Panics if `v` is not in `a`.                               | `must.In(value, slice)`       | Ensures `v` is in `a`.              |
| **`Contains(a []T, v T)`**   | Panics if `a` does not contain `v`.                        | `must.Contains(slice, value)` | Ensures `a` contains `v`.           |

### Boolean Package (`mustboolean`)

| **Function**               | **Description**                             | **Example**                     | **Notes**                            |
| -------------------------- | ------------------------------------------- | ------------------------------- | ------------------------------------ |
| **`True(v bool)`**         | Panics if `v` is false.                     | `mustboolean.True(isEnabled)`   | Validates if `v` is `true`.          |
| **`Conflict(bs ...bool)`** | Panics if multiple boolean values are true. | `mustboolean.Conflict(a, b, c)` | Ensures at most one boolean is true. |

---

## Examples

### Basic Usage Patterns

**Assert non-zero value:**

```go
value := 42
must.Nice(value) // Panics if value is zero
```

**Validate no error:**

```go
err := someFunction()
must.Done(err) // Panics if err is not nil
```

**Check slice length:**

```go
arr := []int{1, 2, 3}
must.Length(arr, 3) // Panics if length is not 3
```

### Common Validation Scenarios

**Validate map operations:**

```go
config := map[string]int{"port": 8080}
port := mustmap.Get(config, "port")
mustnum.Positive(port)
```

**String validation:**

```go
filename := "data.json"
muststrings.HasSuffix(filename, ".json")
muststrings.Contains(filename, "data")
```

**Pointer validation:**

```go
account := findAccount(id)
must.Full(account) // Panics if account is nil
```

---

## Related Projects

Explore more error handling packages in this ecosystem:

### Advanced Packages

- **[must](https://github.com/yylego/must)** - Must-style assertions with rich type support and detailed error context (this project)
- **[rese](https://github.com/yylego/rese)** - Result extraction with panic, focused on safe value unwrapping

### Foundation Packages

- **[done](https://github.com/yylego/done)** - Simple, focused error handling with method chaining
- **[sure](https://github.com/yylego/sure)** - Generates code that creates custom validation methods

Each package targets different use cases, from quick prototyping to production systems with comprehensive error handling.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/must.svg?variant=adaptive)](https://starchart.cc/yylego/must)
