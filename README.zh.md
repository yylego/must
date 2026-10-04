<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/must/release.yml?branch=main&label=BUILD)](https://github.com/yylego/must/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/must)](https://pkg.go.dev/github.com/yylego/must)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/must/main.svg)](https://coveralls.io/github/yylego/must?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/must.svg)](https://github.com/yylego/must/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/must)](https://goreportcard.com/report/github.com/yylego/must)
<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# must

简洁的断言工具包，提供 panic-on-failure 语义，减少 Go 代码样板。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

## 核心特性

- 🎯 **断言检查**：检查失败时记录上下文并触发 panic。
- ⚡ **Go 泛型**：通过类型参数支持多种输入类型。
- 🔄 **栈帧调整**：通过 skip 配置定位调用代码。
- 🌍 **结构化日志**：使用 Zap 记录 panic 上下文。
- 📋 **领域专用包**：提供数值、字符串、切片和映射断言。

---

## 使用边界

- 断言失败会记录上下文并触发 panic，适用于条件不满足时必须中断当前操作的场景；可预期的失败应通过返回错误处理。
- 日志可能包含输入值。敏感值使用 `mustsecret` 提供的断言，避免把值写入日志。
- `Null` 和 `Full` 接收指针（`*T`）；字符串长度按字节计算，不是 Unicode 字符数。
- 数值大小和正负断言拒绝 NaN。`Num` 支持其列出的整数、浮点类型及以这些类型为底层类型的自定义类型。
- `mustmap.Get` 不限制值的类型，只检查键是否存在；键存在且值为 nil 时正常返回。

## 安装

```bash
go get github.com/yylego/must
```

---

## 快速入门

### 示例 1: 基础断言

```go
package main

import (
	"fmt"

	"github.com/yylego/must"
)

func main() {
	fmt.Println("=== Demo 1: 基础断言 ===")

	// 布尔断言
	must.True(checkCondition())
	fmt.Println("✓ 布尔检查通过")

	// 验证无错误
	must.Done(performOperation())
	fmt.Println("✓ 无错误")

	// 非零值
	count := getCount()
	must.Nice(count)
	fmt.Printf("✓ 有效计数: %d\n", count)

	// 值匹配验证
	must.Equals("success", getStatus())
	fmt.Println("✓ 值匹配")

	// 切片操作
	items := getItems()
	must.Have(items)
	must.Length(items, 3)
	must.In("banana", items)
	fmt.Printf("✓ 切片验证通过: %v\n", items)

	// 指针检查
	account := getAccount()
	must.Full(account)
	fmt.Printf("✓ 指针有效: %s\n", account.Name)

	fmt.Println("\n=== 所有检查通过！ ===")
}

type Account struct{ Name string }

func checkCondition() bool    { return true }
func performOperation() error { return nil }
func getStatus() string       { return "success" }
func getCount() int           { return 42 }
func getItems() []string      { return []string{"apple", "banana", "orange"} }
func getAccount() *Account    { return &Account{Name: "test"} }
```

⬆️ **源码：** [源码](internal/demos/demo1x/main.go)

---

### 示例 2: 结果提取函数

```go
package main

import (
	"fmt"

	"github.com/yylego/must"
)

func main() {
	fmt.Println("=== Demo 2: 结果提取函数 ===")

	// V1 - 单值验证
	config := must.V1(readConfig())
	fmt.Printf("✓ 配置: %s\n", config)

	// V2 - 双值验证
	width, height := must.V2(getDimensions())
	fmt.Printf("✓ 尺寸: %dx%d\n", width, height)

	// P1 - 非空数据验证
	admin := must.P1(findAdmin())
	fmt.Printf("✓ 管理员: %s\n", admin.Name)

	// C1 - 非零验证
	num := must.C1(getNum())
	fmt.Printf("✓ 数值: %d\n", num)

	// 组合验证
	data := getData()
	must.Full(data)
	must.Nice(data.Score)
	must.Same(data.Status, "active")
	fmt.Printf("✓ 数据: score=%d, status=%s\n", data.Score, data.Status)

	fmt.Println("\n=== 所有检查通过！ ===")
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

⬆️ **源码：** [源码](internal/demos/demo2x/main.go)

---

### 示例 3: 高级专用包

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
	fmt.Println("=== Demo 3: 高级包 ===")

	// 数值验证
	score := getScore()
	mustnum.Positive(score)
	mustnum.NonNegative(0) // 接受零，拒绝负数及 NaN
	mustnum.NonPositive(0) // 接受零与负数，拒绝 NaN
	mustnum.ZeroPositive(0) // 与 NonNegative 等价
	mustnum.ZeroNegative(0) // 与 NonPositive 等价
	mustnum.Gt(score, 60)
	fmt.Printf("✓ 分数: %d\n", score)

	// 切片验证
	tags := getTags()
	mustslice.Have(tags)
	mustslice.Contains(tags, "go")
	fmt.Printf("✓ 标签: %v\n", tags)

	// Map 验证
	config := getConfig()
	mustmap.Have(config)
	timeout := mustmap.Get(config, "timeout")
	fmt.Printf("✓ 超时: %d\n", timeout)

	// 字符串验证
	filename := getFilename()
	muststrings.HasSuffix(filename, ".pdf")
	muststrings.Contains(filename, "report")
	fmt.Printf("✓ 文件名: %s\n", filename)

	// 复杂场景
	data := getAnalytics()
	must.Full(data)
	mustmap.Have(data.Metrics)
	fmt.Printf("✓ 分析数据: %d 个指标\n", len(data.Metrics))

	fmt.Println("\n=== 所有检查通过！ ===")
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

⬆️ **源码：** [源码](internal/demos/demo3x/main.go)

---

## 核心断言

以下是 `must` 中的核心断言函数，概述如下：

| **函数**                     | **描述**                                                      | **示例**                      | **备注**                        |
| ---------------------------- | ------------------------------------------------------------- | ----------------------------- | ------------------------------- |
| **`True(v bool)`**           | 如果 `v` 为 `false`，触发 panic。                             | `must.True(isValid)`          | 验证 `v` 是否为 `true`。        |
| **`Done(err error)`**        | 如果 `err` 不为 `nil`，触发 panic。                           | `must.Done(err)`              | 确保没有错误发生。              |
| **`Must(err error)`**        | 如果 `err` 不为 `nil`，触发 panic。                           | `must.Must(err)`              | 类似于 `Done`。                 |
| **`Nice(a V)`**              | 如果 `a` 为零，触发 panic。                                   | `must.Nice(value)`            | 确保 `a` 非零。                 |
| **`Zero(a V)`**              | 如果 `a` 不是零，触发 panic。                                 | `must.Zero(value)`            | 确保 `a` 为零。                 |
| **`None(a V)`**              | 如果 `a` 非零，触发 panic。                                   | `must.None(value)`            | 确保 `a` 为零。                 |
| **`Null[T any](v *T)`**      | 如果指针 `v` 不为 `nil`，触发 panic。                         | `must.Null(data)`             | 检查指针为空。                  |
| **`Full[T any](v *T) *T`**   | 如果指针 `v` 为 `nil`，触发 panic；否则返回该指针。           | `must.Full(value)`            | 检查指针非空。                  |
| **`Equals(a, b V)`**         | 如果 `a` 和 `b` 不相等，触发 panic。                          | `must.Equals(a, b)`           | 检查 `a` 是否等于 `b`。         |
| **`Same(a, b V)`**           | 如果 `a` 和 `b` 不相等，触发 panic。                          | `must.Same(a, b)`             | `Equals` 的别名。               |
| **`SameNice(a, b V)`**       | 如果 `a` 和 `b` 不相等或为零，触发 panic。                    | `must.SameNice(a, b)`         | 确保相等且非零。                |
| **`Sane(a, b V)`**           | 如果 `a` 和 `b` 不相等或为零，触发 panic。                    | `must.Sane(a, b)`             | `SameNice` 的别名。             |
| **`Diff(a, b V)`**           | 如果 `a` 和 `b` 相等，触发 panic。                            | `must.Diff(a, b)`             | 确保值不同。                    |
| **`Different(a, b V)`**      | 如果 `a` 和 `b` 相等，触发 panic。                            | `must.Different(a, b)`        | `Diff` 的别名。                 |
| **`Is(a, b V)`**             | 如果 `a` 和 `b` 不相等，触发 panic。                          | `must.Is(a, b)`               | `Equals` 的别名。               |
| **`Ise(err, target error)`** | 如果 `err` 不与 `target` 匹配，触发 panic，使用 `errors.Is`。 | `must.Ise(err, targetErr)`    | 类似于 `errors.Is` 的错误匹配。 |
| **`Ok(a V)`**                | 如果 `a` 为零，触发 panic。                                   | `must.Ok(value)`              | 确保 `a` 非零。                 |
| **`OK(a V)`**                | `Ok` 的别名，检查值是否非零。                                 | `must.OK(value)`              | 与 `Ok` 相同。                  |
| **`TRUE(v bool)`**           | 如果 `v` 为 `false`，触发 panic。                             | `must.TRUE(isValid)`          | `True` 的别名。                 |
| **`FALSE(v bool)`**          | 如果 `v` 为 `true`，触发 panic。                              | `must.FALSE(failed)`          | 确保 `v` 为 `false`。           |
| **`False(v bool)`**          | 如果 `v` 为 `true`，触发 panic。                              | `must.False(failed)`          | 与 `FALSE` 相同。               |
| **`Cause(err error)`**       | 如果 `err` 为 `nil`，触发 panic，返回该错误。                 | `must.Cause(err)`             | 确保错误存在。                  |
| **`Wrong(err error)`**       | 如果 `err` 为 `nil`，触发 panic。                             | `must.Wrong(err)`             | 确保错误存在。                  |
| **`Have(a []T)`**            | 如果 `a` 为空，触发 panic。                                   | `must.Have(slice)`            | 确保 `a` 不为空。               |
| **`Length(a []T, n int)`**   | 如果 `a` 的长度不是 `n`，触发 panic。                         | `must.Length(slice, 3)`       | 确保 `a` 的长度是 `n`。         |
| **`Len(a []T, n int)`**      | `Length` 的别名，确保 `a` 的长度是 `n`。                      | `must.Len(slice, 3)`          | 验证 `a` 的长度。               |
| **`In(v T, a []T)`**         | 如果 `v` 不在 `a` 中，触发 panic。                            | `must.In(value, slice)`       | 确保 `v` 在 `a` 中。            |
| **`Contains(a []T, v T)`**   | 如果 `a` 不包含 `v`，触发 panic。                             | `must.Contains(slice, value)` | 确保 `a` 包含 `v`。             |

### 布尔包 (`mustboolean`)

| **函数**                   | **描述**                              | **示例**                        | **备注**                      |
| -------------------------- | ------------------------------------- | ------------------------------- | ----------------------------- |
| **`True(v bool)`**         | 如果 `v` 为 `false`，触发 panic。     | `mustboolean.True(isEnabled)`   | 验证 `v` 是否为 `true`。      |
| **`Conflict(bs ...bool)`** | 如果多个布尔值为 `true`，触发 panic。 | `mustboolean.Conflict(a, b, c)` | 确保最多一个布尔值为 `true`。 |

---

## 使用示例

### 基础使用模式

**断言非零值：**

```go
value := 42
must.Nice(value) // 如果值为零则 panic
```

**验证无错误：**

```go
err := someFunction()
must.Done(err) // 如果 err 非 nil 则 panic
```

**检查切片长度：**

```go
items := []int{1, 2, 3}
must.Length(items, 3) // 如果长度不是 3 则 panic
```

### 常见验证场景

**Map 操作验证：**

```go
config := map[string]int{"port": 8080}
port := mustmap.Get(config, "port")
mustnum.Positive(port)
```

**字符串验证：**

```go
filename := "data.json"
muststrings.HasSuffix(filename, ".json")
muststrings.Contains(filename, "data")
```

**指针验证：**

```go
account := findAccount(id)
must.Full(account) // 如果 account 为 nil 则 panic
```

---

## 相关项目

探索此生态系统中的更多错误处理包：

### 高级包

- **[must](https://github.com/yylego/must)** - Must 风格断言，提供丰富的类型支持和详细的错误上下文（本项目）
- **[rese](https://github.com/yylego/rese)** - 带 panic 的结果提取，专注于安全的值解包

### 基础包

- **[done](https://github.com/yylego/done)** - 简单专注的错误处理，支持方法链式调用
- **[sure](https://github.com/yylego/sure)** - 生成代码以创建自定义验证方法

每个包针对不同的使用场景，从快速原型开发到具有全面错误处理的生产系统。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->

---

<!-- TEMPLATE (ZH) BEGIN: GITHUB STARS -->

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yylego/must.svg?variant=adaptive)](https://starchart.cc/yylego/must)
<!-- TEMPLATE (ZH) CLOSE: GITHUB STARS -->
