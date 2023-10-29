// test case has been written with the same package as the code
// to be able to mock sync.Once.
package log

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/sinhashubham95/go-utils/errors"
	"github.com/stretchr/testify/assert"
)

type cKey string

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
	Trace(context.Background())
	Debug(context.Background())
	Info(context.Background())
	var a cKey = "naruto"
	ctx := context.WithValue(context.Background(), a, "rocks")
	Warn(ctx)
	Error(ctx)
	Fatal(ctx)
	Panic(ctx)
}

func TestLoggerErrorWarn(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, []string{"naruto", "rocks"})
	ErrorWarn(context.Background(), &errors.Error{StatusCode: http.StatusUnauthorized})
	ErrorWarn(context.Background(), nil)
}

func TestLoggerCapabilities(t *testing.T) {
	defer resetOnce()
	InitLogger(DebugLevel, []string{"naruto", "rocks"})
	Error(context.Background()).Err(&errors.Error{Code: "naruto"}).Stack().Msg("naruto rocks")
	Error(context.Background()).Msgf("naruto rocks")
	Error(context.Background()).Send()
	Debug(context.Background()).Bool("naruto", false).Send()
	Debug(context.Background()).Bools("naruto", []bool{false}).Send()

	Debug(context.Background()).Bytes("naruto", []byte("naruto rocks")).Send()

	Debug(context.Background()).Errs("naruto", nil).Send()

	Debug(context.Background()).Float32("naruto", 0).Send()
	Debug(context.Background()).Floats32("naruto", nil).Send()
	Debug(context.Background()).Float64("naruto", 0).Send()
	Debug(context.Background()).Floats64("naruto", nil).Send()

	Debug(context.Background()).Hex("naruto", nil).Send()

	Debug(context.Background()).Int("naruto", 0).Send()
	Debug(context.Background()).Ints("naruto", nil).Send()
	Debug(context.Background()).Int8("naruto", 0).Send()
	Debug(context.Background()).Ints8("naruto", nil).Send()
	Debug(context.Background()).Int16("naruto", 0).Send()
	Debug(context.Background()).Ints16("naruto", nil).Send()
	Debug(context.Background()).Int32("naruto", 0).Send()
	Debug(context.Background()).Ints32("naruto", nil).Send()
	Debug(context.Background()).Int64("naruto", 0).Send()
	Debug(context.Background()).Ints64("naruto", nil).Send()

	Debug(context.Background()).RawJSON("naruto", []byte("naruto rocks"))

	Debug(context.Background()).Str("naruto", "").Send()
	Debug(context.Background()).Strs("naruto", nil).Send()

	Debug(context.Background()).Stringer("naruto", nil).Send()
	Debug(context.Background()).Stringers("naruto", nil).Send()

	Debug(context.Background()).Uint("naruto", 0).Send()
	Debug(context.Background()).Uints("naruto", nil).Send()
	Debug(context.Background()).Uint8("naruto", 0).Send()
	Debug(context.Background()).Uints8("naruto", nil).Send()
	Debug(context.Background()).Uint16("naruto", 0).Send()
	Debug(context.Background()).Uints16("naruto", nil).Send()
	Debug(context.Background()).Uint32("naruto", 0).Send()
	Debug(context.Background()).Uints32("naruto", nil).Send()
	Debug(context.Background()).Uint64("naruto", 0).Send()
	Debug(context.Background()).Uints64("naruto", nil).Send()

	Debug(context.Background()).Interface("naruto", nil).Send()
	Debug(context.Background()).Any("naruto", nil).Send()
}

func resetOnce() {
	o = &sync.Once{}
}
