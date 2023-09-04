package log

import (
	"fmt"
	"github.com/rs/zerolog"
	"sync"
)

// Logger is the interface to the logging possibilities.
type Logger interface {
	Err(err error) Logger

	// Msg sends the *Event with msg added as the message field if not empty.
	// NOTICE: once this method is called, the Logger should be disposed.
	// Calling Msg twice can have unexpected result.
	Msg(msg string)

	// Msgf sends the event with formatted msg added as the message field if not empty.
	// NOTICE: once this method is called, the *Event should be disposed.
	// Calling Msgf twice can have unexpected result.
	Msgf(format string, v ...interface{})

	// Send is equivalent to calling Msg("").
	// NOTICE: once this method is called, the Logger should be disposed.
	Send()

	// Str adds the field key with val as a string to the Logger context.
	Str(key, val string) Logger

	// Strs adds the field key with values as a []string to the *Event context.
	Strs(key string, values []string) Logger

	// Stringer adds the field key with val.String() (or null if val is nil)
	// to the Logger context.
	Stringer(key string, val fmt.Stringer) Logger

	// Stringers adds the field key with values where each individual val
	// is used as val.String() (or null if val is empty) to the Logger
	// context.
	Stringers(key string, values []fmt.Stringer) Logger
}

type l struct {
	e *zerolog.Event
}

var lPool = &sync.Pool{
	New: func() interface{} {
		return &l{}
	},
}

func newL(e *zerolog.Event) *l {
	x := lPool.Get().(*l)
	x.e = e
	return x
}

func disposeL(x *l) {
	if x.e != nil {
		x.e = nil
	}
	lPool.Put(x)
}

func (x *l) Err(err error) Logger {
	x.e = x.e.Err(err)
	return x
}

func (x *l) Msg(msg string) {
	x.e.Msg(msg)
	disposeL(x)
}

func (x *l) Msgf(format string, v ...interface{}) {
	x.e.Msgf(format, v...)
	disposeL(x)
}

func (x *l) Send() {
	x.e.Send()
	disposeL(x)
}

func (x *l) Str(key, val string) Logger {
	x.e.Str(key, val)
	return x
}

func (x *l) Strs(key string, values []string) Logger {
	x.e.Strs(key, values)
	return x
}

func (x *l) Stringer(key string, val fmt.Stringer) Logger {
	x.e.Stringer(key, val)
	return x
}

func (x *l) Stringers(key string, values []fmt.Stringer) Logger {
	x.e.Stringers(key, values)
	return x
}
