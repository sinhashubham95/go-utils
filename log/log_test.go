// test case has been written with the same package as the code
// to be able to mock sync.Once.
package log

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/sinhashubham95/go-utils/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"sync"
	"testing"
)

func TestInitLoggerTraceLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.TraceLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.TraceLevel, zerolog.GlobalLevel())
}

func TestInitLoggerDebugLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, nil)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
}

func TestInitLoggerInfoLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(InfoLevel, nil)
	assert.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel())
}

func TestInitLoggerWarnLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(WarnLevel, nil)
	assert.Equal(t, zerolog.WarnLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.WarnLevel, zerolog.GlobalLevel())
}

func TestInitLoggerErrorLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(ErrorLevel, nil)
	assert.Equal(t, zerolog.ErrorLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.ErrorLevel, zerolog.GlobalLevel())
}

func TestInitLoggerFatalLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(FatalLevel, nil)
	assert.Equal(t, zerolog.FatalLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.FatalLevel, zerolog.GlobalLevel())
}

func TestInitLoggerPanicLevel(t *testing.T) {
	defer resetOnce()
	InitLogger(PanicLevel, nil)
	assert.Equal(t, zerolog.PanicLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.PanicLevel, zerolog.GlobalLevel())
}

func TestInitLoggerIncorrectLevel(t *testing.T) {
	defer resetOnce()
	InitLogger("naruto", nil)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
	InitLogger(TraceLevel, nil)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
}

func TestInitLoggerWithNilWriter(t *testing.T) {
	defer resetOnce()
	InitLoggerWithWriter(DebugLevel, nil, nil)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
}

func TestInitLoggerWithWriter(t *testing.T) {
	defer resetOnce()
	fi, err := os.OpenFile(fmt.Sprintf("%s/test.log", t.TempDir()),
		os.O_CREATE|os.O_APPEND, os.ModePerm)
	assert.NoError(t, err)
	defer func(f *os.File) {
		e := f.Close()
		assert.NoError(t, e)
	}(fi)
	InitLoggerWithWriter(DebugLevel, fi, nil)
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
}

func TestInitLoggerWithParams(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, []string{"naruto", "rocks"})
	assert.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel())
	assert.Equal(t, []string{"naruto", "rocks"}, p)
}

func TestLoggerMethods(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, []string{"naruto", "rocks"})
	Trace(nil)
	Debug(nil)
	Info(context.Background())
	ctx := context.WithValue(context.Background(), "naruto", "rocks")
	Warn(ctx)
	Error(ctx)
	Fatal(ctx)
	Panic(ctx)
}

func TestLoggerErrorWarn(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, []string{"naruto", "rocks"})
	ErrorWarn(nil, &errors.Error{StatusCode: http.StatusUnauthorized})
	ErrorWarn(nil, nil)
}

func TestLoggerCapabilities(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, []string{"naruto", "rocks"})
	Error(nil).Err(&errors.Error{Code: "naruto"}).Stack().Msg("naruto rocks")
	Error(nil).Msgf("naruto rocks")
	Error(nil).Send()
	Debug(nil).Bool("naruto", false).Send()
	Debug(nil).Bools("naruto", []bool{false}).Send()

	Debug(nil).Bytes("naruto", []byte("naruto rocks")).Send()

	Debug(nil).Errs("naruto", nil).Send()

	Debug(nil).Float32("naruto", 0).Send()
	Debug(nil).Floats32("naruto", nil).Send()
	Debug(nil).Float64("naruto", 0).Send()
	Debug(nil).Floats64("naruto", nil).Send()

	Debug(nil).Hex("naruto", nil).Send()

	Debug(nil).Int("naruto", 0).Send()
	Debug(nil).Ints("naruto", nil).Send()
	Debug(nil).Int8("naruto", 0).Send()
	Debug(nil).Ints8("naruto", nil).Send()
	Debug(nil).Int16("naruto", 0).Send()
	Debug(nil).Ints16("naruto", nil).Send()
	Debug(nil).Int32("naruto", 0).Send()
	Debug(nil).Ints32("naruto", nil).Send()
	Debug(nil).Int64("naruto", 0).Send()
	Debug(nil).Ints64("naruto", nil).Send()

	Debug(nil).RawJSON("naruto", []byte("naruto rocks"))

	Debug(nil).Str("naruto", "").Send()
	Debug(nil).Strs("naruto", nil).Send()

	Debug(nil).Stringer("naruto", nil).Send()
	Debug(nil).Stringers("naruto", nil).Send()

	Debug(nil).Uint("naruto", 0).Send()
	Debug(nil).Uints("naruto", nil).Send()
	Debug(nil).Uint8("naruto", 0).Send()
	Debug(nil).Uints8("naruto", nil).Send()
	Debug(nil).Uint16("naruto", 0).Send()
	Debug(nil).Uints16("naruto", nil).Send()
	Debug(nil).Uint32("naruto", 0).Send()
	Debug(nil).Uints32("naruto", nil).Send()
	Debug(nil).Uint64("naruto", 0).Send()
	Debug(nil).Uints64("naruto", nil).Send()
}

func resetOnce() {
	o = &sync.Once{}
}
