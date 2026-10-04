package mustnum_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/must/mustnum"
)

func TestNonNegative(t *testing.T) {
	require.NotPanics(t, func() { mustnum.NonNegative(0) })
	require.NotPanics(t, func() { mustnum.NonNegative(uint64(1)) })
	require.NotPanics(t, func() { mustnum.NonNegative(0.5) })
	require.Panics(t, func() { mustnum.NonNegative(-1) })
	require.Panics(t, func() { mustnum.NonNegative(math.NaN()) })
}
