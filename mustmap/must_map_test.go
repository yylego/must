// Package mustmap_test provides comprehensive testing of mustmap assertion package
// Tests include map matching, difference validation, length checking, and item lookup
// Checks each assertion functions with both success and failure cases
//
// mustmap_test 为 mustmap 断言包提供全面的测试
// 测试涵盖 map 相等性、差异验证、长度检查和键存在性验证
// 使用成功和失败案例验证所有断言函数
package mustmap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/must/mustmap"
)

// TestEquals tests matching maps
// Validates Equals passes when maps have same content and panics on mismatch
//
// TestEquals 测试 map 相等断言
// 验证 Equals 在 map 内容相同时通过，在 map 不同时 panic
func TestEquals(t *testing.T) {
	mustmap.Equals(map[int]string{
		1: "a",
		2: "b",
	}, map[int]string{
		2: "b",
		1: "a",
	})

	require.Panics(t, func() {
		mustmap.Equals(map[string]int{
			"a": 1,
		}, map[string]int{
			"b": 2,
		})
	})
}

// TestDiff tests map difference assertion
// Validates Diff passes on mismatch and panics when maps are same
//
// TestDiff 测试 map 差异断言
// 验证 Diff 在 map 不同时通过，在 map 相同时 panic
func TestDiff(t *testing.T) {
	mustmap.Diff(map[string]int{
		"a": 1,
	}, map[string]int{
		"b": 2,
	})

	require.Panics(t, func() {
		mustmap.Diff(map[int]string{
			1: "a",
			2: "b",
		}, map[int]string{
			2: "b",
			1: "a",
		})
	})
}

// TestDifferent tests map difference assertion
// Validates Different passes on mismatch and panics when maps are same
//
// TestDifferent 测试 map 差异断言
// 验证 Different 在 map 不同时通过，在 map 相同时 panic
func TestDifferent(t *testing.T) {
	mustmap.Different(map[string]int{
		"a": 1,
	}, map[string]int{
		"b": 2,
	})

	require.Panics(t, func() {
		mustmap.Different(map[int]string{
			1: "a",
			2: "b",
		}, map[int]string{
			2: "b",
			1: "a",
		})
	})
}

// TestHave tests maps with items
// Validates Have passes with items and panics without items
//
// TestHave 测试 map 非空断言
// 验证 Have 在 map 非空时通过，在 map 为空时 panic
func TestHave(t *testing.T) {
	mustmap.Have(map[int]string{
		1: "value1",
		2: "value2",
	})

	require.Panics(t, func() {
		mustmap.Have(map[string]int{})
	})
}

// TestNice tests maps with items and returns the map
// Validates Nice returns maps with items and panics without items
//
// TestNice 测试非空 map 断言并返回
// 验证 Nice 返回非空 map 并在空 map 时 panic
func TestNice(t *testing.T) {
	require.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, mustmap.Nice(map[string]int{"c": 3, "b": 2, "a": 1}))

	require.Panics(t, func() {
		mustmap.Nice(map[string]int{})
	})
}

// TestZero tests vacant maps
// Validates Zero passes without items and panics with items
//
// TestZero 测试空 map 断言
// 验证 Zero 在空 map 时通过，在非空 map 时 panic
func TestZero(t *testing.T) {
	mustmap.Zero(map[int]string{})
	mustmap.Zero(map[string]int{})
	mustmap.Zero((map[int]string)(nil))

	require.Panics(t, func() {
		mustmap.Zero(map[int]string{1: "a"})
	})

	require.Panics(t, func() {
		mustmap.Zero(map[string]int{"a": 1})
	})

	require.Panics(t, func() {
		mustmap.Zero(map[string]int{"a": 1, "b": 2})
	})
}

// TestNone tests vacant maps (alias of Zero)
// Validates None passes without items and panics with items
//
// TestNone 测试空 map 断言（Zero 的别名）
// 验证 None 在空 map 时通过，在非空 map 时 panic
func TestNone(t *testing.T) {
	mustmap.None(map[int]int{})

	require.Panics(t, func() {
		mustmap.None(map[string]int{"a": 1, "b": 2})
	})
}

// TestLength tests map length assertion
// Validates Length passes when map length matches expected value and panics on mismatch
//
// TestLength 测试 map 长度断言
// 验证 Length 在 map 长度匹配期望值时通过，在不匹配时 panic
func TestLength(t *testing.T) {
	mustmap.Length(map[int]string{
		1: "value1",
		2: "value2",
	}, 2)

	require.Panics(t, func() {
		mustmap.Length(map[int]string{
			1: "value1",
		}, 2)
	})

	mustmap.Length(map[string]int{}, 0)

	require.Panics(t, func() {
		mustmap.Length(map[string]int{"a": 1, "b": 2}, 3)
	})
}

// TestLen tests map length assertion (short form)
// Validates Len passes when map length matches expected value and panics on mismatch
//
// TestLen 测试 map 长度断言（简短形式）
// 验证 Len 在 map 长度匹配期望值时通过，在不匹配时 panic
func TestLen(t *testing.T) {
	mustmap.Len(map[int]string{
		1: "value1",
		2: "value2",
	}, 2)

	require.Panics(t, func() {
		mustmap.Len(map[int]string{
			1: "value1",
		}, 2)
	})

	mustmap.Len(map[string]int{}, 0)

	require.Panics(t, func() {
		mustmap.Len(map[string]int{"a": 1, "b": 2}, 3)
	})
}

// TestGet tests map item lookup
// Validates Get returns the value when found and panics when absent
//
// TestGet 测试 map 键存在性并返回值
// 验证 Get 在键存在时返回值，在键不存在时 panic
func TestGet(t *testing.T) {
	value := mustmap.Get(map[string]int{"a": 1, "b": 2}, "a")
	require.Equal(t, 1, value)

	require.Panics(t, func() {
		mustmap.Get(map[string]int{"a": 1, "b": 2}, "c")
	})
}

func TestGetSlice(t *testing.T) {
	values := map[string][]int{"present": {1, 2}, "nil": nil}
	require.Equal(t, []int{1, 2}, mustmap.Get(values, "present"))
	require.Nil(t, mustmap.Get(values, "nil"))
	require.Panics(t, func() { mustmap.Get(values, "absent") })
}
