package must_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/must"
	"github.com/yylego/must/mustnum"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLogSite(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)
	saved := zaplog.ZAPS
	zaplog.ZAPS = zaplog.NewSkipZaps(zap.New(core, zap.AddCaller()))
	t.Cleanup(func() { zaplog.ZAPS = saved })
	err := errors.New("sample")
	checks := []func(){
		func() { mustnum.ZeroPositive(-1) },
		func() { mustnum.ZeroNegative(1) },
		func() { must.V0(err) },
		func() { must.V1(1, err) },
		func() { must.V2(1, 2, err) },
		func() { must.P0(err) },
		func() { must.P1[int](nil, nil) },
		func() { must.P2(new(int), (*int)(nil), nil) },
		func() { must.C0(err) },
		func() { must.C1(0, nil) },
		func() { must.C2(1, 0, nil) },
	}
	for _, check := range checks {
		require.Panics(t, check)
	}
	entries := logs.All()
	require.Len(t, entries, len(checks))
	for _, record := range entries {
		require.True(t, strings.HasSuffix(record.Caller.File, "/log_site_test.go"), record.Caller.String())
	}

	logs.TakeAll()
	require.Panics(t, func() { must.Ise(err, errors.New("expected")) })
	entries = logs.All()
	require.Len(t, entries, 1)
	require.Len(t, entries[0].Context, 2)
	require.Equal(t, "sample", entries[0].ContextMap()["given"])
	require.Equal(t, "expected", entries[0].ContextMap()["target"])
}
