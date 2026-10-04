package mustnum_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/must/mustnum"
)

func TestNonPositive(t *testing.T) {
	require.NotPanics(t, func() { mustnum.NonPositive(0) })
	require.NotPanics(t, func() { mustnum.NonPositive(uint64(0)) })
	require.NotPanics(t, func() { mustnum.NonPositive(-1) })
	require.NotPanics(t, func() { mustnum.NonPositive(-0.5) })
	require.Panics(t, func() { mustnum.NonPositive(1) })
	require.Panics(t, func() { mustnum.NonPositive(uint64(1)) })
	require.Panics(t, func() { mustnum.NonPositive(0.5) })
	require.Panics(t, func() { mustnum.NonPositive(math.NaN()) })
}
