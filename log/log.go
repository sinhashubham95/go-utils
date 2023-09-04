package log

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sinhashubham95/go-utils/errors"
	"io"
	"runtime/debug"
	"sync"
)

var o = sync.Once{}
var mu sync.RWMutex
var p []string

// InitLogger is used to initialize logger
func InitLogger(level Level, params []string) {
	o.Do(func() {
		zerolog.ErrorStackMarshaler = getErrorStackMarshaller()
		zerolog.SetGlobalLevel(level.zeroLogLevel())
		mu.Lock()
		defer mu.Unlock()
		log.Logger = log.With().Caller().Logger()
		p = params
	})
}

// InitLoggerWithWriter is used to initialize logger with a writer
func InitLoggerWithWriter(level Level, w io.Writer, params []string) {
	o.Do(func() {
		zerolog.ErrorStackMarshaler = getErrorStackMarshaller()
		zerolog.SetGlobalLevel(level.zeroLogLevel())
		log.Logger = zerolog.New(w).With().Caller().Timestamp().Logger()
		mu.Lock()
		defer mu.Unlock()
		p = params
	})
}

// Trace is the for trace log
func Trace(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Trace()))
}

// Debug is the for debug log
func Debug(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Debug()))
}

// Info is the for info log
func Info(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Info()))
}

// Warn is the for warn log
func Warn(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Warn()))
}

// Error is the for error log
func Error(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Error().Stack()))
}

// Panic is the for panic log
func Panic(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Panic().Stack()))
}

// Fatal is the for fatal log
func Fatal(ctx context.Context) Logger {
	return newL(withParams(ctx, log.Fatal().Stack()))
}

// ErrorWarn checks for the error object.
// In case it is corresponding to a 4XX status code, it logs it as warning.
// Otherwise, it logs it as an error.
func ErrorWarn(ctx context.Context, err error) Logger {
	if e, ok := err.(*errors.Error); ok && e.StatusCode >= 400 && e.StatusCode < 500 {
		return Warn(ctx).Err(err)
	}
	return Error(ctx).Err(err)
}

func getErrorStackMarshaller() func(err error) interface{} {
	return func(err error) interface{} {
		if err != nil {
			if e, ok := err.(*errors.Error); ok {
				return map[string]interface{}{
					CodeLogParam:    e.Code,
					MessageLogParam: e.Message,
					DetailsLogParam: e.Details,
					TraceLogParam:   e.GetTrace(),
				}
			}
		}
		return string(debug.Stack())
	}
}

func withParams(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return event
	}
	mu.RLock()
	defer mu.RUnlock()
	for _, k := range p {
		v := ctx.Value(k)
		if v != nil {
			event.Interface(k, v)
		}
	}
	return event.Ctx(ctx)
}
