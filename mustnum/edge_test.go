package mustnum_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/must/mustnum"
)

func TestNaNComparison(t *testing.T) {
	for _, check := range []func(float64, float64){mustnum.Less[float64], mustnum.Lt[float64], mustnum.Lte[float64], mustnum.Gt[float64], mustnum.Gte[float64]} {
		require.Panics(t, func() { check(math.NaN(), 1) })
		require.Panics(t, func() { check(1, math.NaN()) })
	}
	for _, check := range []func(float64){mustnum.Positive[float64], mustnum.Negative[float64], mustnum.NonPositive[float64], mustnum.NonNegative[float64], mustnum.ZeroPositive[float64], mustnum.ZeroNegative[float64]} {
		require.Panics(t, func() { check(math.NaN()) })
	}
}

func TestNamedNumbers(t *testing.T) {
	type Amount int64
	type Rate float32
	require.NotPanics(t, func() {
		mustnum.Positive(Amount(1))
		mustnum.Less(Amount(-1), Amount(0))
		mustnum.ZeroPositive(Rate(0))
		mustnum.ZeroNegative(Rate(-1))
	})
	require.Panics(t, func() { mustnum.ZeroPositive(Amount(-1)) })
	require.Panics(t, func() { mustnum.ZeroNegative(Amount(1)) })
}
