// Package mustnum provides numeric-specific assertion utilities with panic-on-failure semantics
// Implements type-safe validation functions with numeric comparison and state checking
// Supports numeric types spanning integers and floating-points through generic Num constraint
// Integrates with zap structured logging to provide detailed context when assertions are not met
//
// mustnum 提供数值特定的断言工具，带 panic-on-failure 语义
// 实现类型安全的数值比较和状态检查验证函数
// 通过泛型 Num 约束支持所有整数和浮点类型
// 与 zap 结构化日志集成，当断言不满足时提供详细上下文
package mustnum

import (
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

// Num defines the constraint spanning numeric types including integers and floats
// Named types with these underlying types are accepted. Ordered comparisons reject NaN.
// Num 支持下列整数、浮点类型及其自定义类型；大小和正负断言拒绝 NaN。
type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Less validates that a is less than b. Panics if a >= b.
// Less 验证 a 小于 b。如果 a >= b 则触发 panic。
func Less[V Num](a, b V) {
	if !(a < b) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED A < B", zap.Any("a", a), zap.Any("b", b))
	}
}

// Lt validates that a is less than b. Alias of Less function. Panics if a >= b.
// Lt 验证 a 小于 b。Less 函数的别名。如果 a >= b 则触发 panic。
func Lt[V Num](a, b V) {
	if !(a < b) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED A < B", zap.Any("a", a), zap.Any("b", b))
	}
}

// Lte validates that a is less than / at most b. Panics if a > b.
// Lte 验证 a 小于或等于 b。如果 a > b 则触发 panic。
func Lte[V Num](a, b V) {
	if !(a <= b) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED A <= B", zap.Any("a", a), zap.Any("b", b))
	}
}

// Gt validates that a exceeds b. Panics if a <= b.
// Gt 验证 a 大于 b。如果 a <= b 则触发 panic。
func Gt[V Num](a, b V) {
	if !(a > b) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED A > B", zap.Any("a", a), zap.Any("b", b))
	}
}

// Gte validates that a exceeds / matches b. Panics if a < b.
// Gte 验证 a 大于或等于 b。如果 a < b 则触发 panic。
func Gte[V Num](a, b V) {
	if !(a >= b) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED A >= B", zap.Any("a", a), zap.Any("b", b))
	}
}

// Nice validates that numeric value is non-zero. Returns the value if non-zero, panics if zero.
// Nice 验证数值非零。如果非零则返回该值，如果为零则触发 panic。
func Nice[V Num](a V) V {
	if a == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
	return a
}

// Zero validates that numeric value is precise zero. Panics if non-zero.
// Zero 验证数值恰好为零。如果非零则触发 panic。
func Zero[V Num](a V) {
	if a != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS NOT ZERO(SHOULD BE ZERO)", zap.Any("a", a))
	}
}

// Positive validates that value exceeds zero. Panics if value <= 0.
// Positive 验证值严格大于零。如果值 <= 0 则触发 panic。
func Positive[V Num](v V) {
	if !(v > 0) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED V > 0", zap.Any("v", v))
	}
}

// NonNegative accepts zero and positive values. NaN fails the check.
// NonNegative 断言数值非负，接受零与正数，拒绝 NaN。
func NonNegative[V Num](v V) {
	if !(v >= 0) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED V >= 0", zap.Any("v", v))
	}
}

// ZeroPositive accepts zero and positive values, like NonNegative. NaN fails the check.
// ZeroPositive 断言数值为零或正数，与 NonNegative 等价，拒绝 NaN。
func ZeroPositive[V Num](v V) {
	if !(v >= 0) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED V >= 0", zap.Any("v", v))
	}
}

// Negative validates that value is below zero. Panics if value >= 0.
// Negative 验证值严格小于零。如果值 >= 0 则触发 panic。
func Negative[V Num](v V) {
	if !(v < 0) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED V < 0", zap.Any("v", v))
	}
}

// NonPositive accepts zero and negative values. NaN fails the check.
// NonPositive 断言数值非正，接受零与负数，拒绝 NaN。
func NonPositive[V Num](v V) {
	if !(v <= 0) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED V <= 0", zap.Any("v", v))
	}
}

// ZeroNegative accepts zero and negative values, like NonPositive. NaN fails the check.
// ZeroNegative 断言数值为零或负数，与 NonPositive 等价，拒绝 NaN。
func ZeroNegative[V Num](v V) {
	if !(v <= 0) {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED V <= 0", zap.Any("v", v))
	}
}
