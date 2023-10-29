package log

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog"
)

// Logger is the interface to the logging possibilities.
type Logger interface {
	// Msg sends the Logger with msg added as the message field if not empty.
	// NOTICE: once this method is called, the Logger should be disposed.
	// Calling Msg twice can have unexpected result.
	Msg(msg string)

	// Msgf sends the event with formatted msg added as the message field if not empty.
	// NOTICE: once this method is called, the Logger should be disposed.
	// Calling Msgf twice can have unexpected result.
	Msgf(format string, v ...interface{})

	// Send is equivalent to calling Msg("").
	// NOTICE: once this method is called, the Logger should be disposed.
	Send()

	// Bool adds the field key with val as a bool to the Logger context.
	Bool(key string, b bool) Logger

	// Bools adds the field key with val as a []bool to the Logger context.
	Bools(key string, b []bool) Logger

	// Bytes adds the field key with val as a string to the Logger context.
	// Runes outside of normal ASCII ranges will be hex-encoded in the resulting
	// JSON.
	Bytes(key string, val []byte) Logger

	// Err adds the field "error" with serialized err to the Logger context.
	// If err is nil, no field is added.
	Err(err error) Logger

	// Errs adds the field key with errs as an array of serialized errors to the
	// Logger context.
	Errs(key string, errs []error) Logger

	// Float32 adds the field key with f as a float32 to the Logger context.
	Float32(key string, f float32) Logger

	// Floats32 adds the field key with f as a []float32 to the Logger context.
	Floats32(key string, f []float32) Logger

	// Float64 adds the field key with f as a float64 to the Logger context.
	Float64(key string, f float64) Logger

	// Floats64 adds the field key with f as a []float64 to the Logger context.
	Floats64(key string, f []float64) Logger

	// Hex adds the field key with val as a hex string to the Logger context.
	Hex(key string, val []byte) Logger

	// Int adds the field key with i as a int to the Logger context.
	Int(key string, i int) Logger

	// Ints adds the field key with i as a []int to the Logger context.
	Ints(key string, i []int) Logger

	// Int8 adds the field key with i as a int8 to the Logger context.
	Int8(key string, i int8) Logger

	// Ints8 adds the field key with i as a []int8 to the Logger context.
	Ints8(key string, i []int8) Logger

	// Int16 adds the field key with i as a int16 to the Logger context.
	Int16(key string, i int16) Logger

	// Ints16 adds the field key with i as a []int16 to the Logger context.
	Ints16(key string, i []int16) Logger

	// Int32 adds the field key with i as a int32 to the Logger context.
	Int32(key string, i int32) Logger

	// Ints32 adds the field key with i as a []int32 to the Logger context.
	Ints32(key string, i []int32) Logger

	// Int64 adds the field key with i as a int64 to the Logger context.
	Int64(key string, i int64) Logger

	// Ints64 adds the field key with i as a []int64 to the Logger context.
	Ints64(key string, i []int64) Logger

	// RawJSON adds already encoded JSON to the log line under key.
	// No sanity check is performed on b; it must not contain carriage returns and
	// be valid JSON.
	RawJSON(key string, b []byte) Logger

	// Stack enables stack trace printing for the error passed to Err().
	Stack() Logger

	// Str adds the field key with val as a string to the Logger context.
	Str(key, val string) Logger

	// Strs adds the field key with values as a []string to the Logger context.
	Strs(key string, values []string) Logger

	// Stringer adds the field key with val.String() (or null if val is nil)
	// to the Logger context.
	Stringer(key string, val fmt.Stringer) Logger

	// Stringers adds the field key with values where each individual val
	// is used as val.String() (or null if val is empty) to the Logger
	// context.
	Stringers(key string, values []fmt.Stringer) Logger

	// Uint adds the field key with i as a uint to the Logger context.
	Uint(key string, i uint) Logger

	// Uints adds the field key with i as a []uint to the Logger context.
	Uints(key string, i []uint) Logger

	// Uint8 adds the field key with i as a uint8 to the Logger context.
	Uint8(key string, i uint8) Logger

	// Uints8 adds the field key with i as a []uint8 to the Logger context.
	Uints8(key string, i []uint8) Logger

	// Uint16 adds the field key with i as a uint16 to the Logger context.
	Uint16(key string, i uint16) Logger

	// Uints16 adds the field key with i as a []int16 to the Logger context.
	Uints16(key string, i []uint16) Logger

	// Uint32 adds the field key with i as a uint32 to the Logger context.
	Uint32(key string, i uint32) Logger

	// Uints32 adds the field key with i as a []int32 to the Logger context.
	Uints32(key string, i []uint32) Logger

	// Uint64 adds the field key with i as a uint64 to the Logger context.
	Uint64(key string, i uint64) Logger

	// Uints64 adds the field key with i as a []int64 to the Logger context.
	Uints64(key string, i []uint64) Logger

	// Interface adds the field key with value marshaled using reflection.
	Interface(key string, i interface{}) Logger

	// Any adds the field key with value marshaled using reflection.
	Any(key string, a any) Logger
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

func (x *l) Bool(key string, b bool) Logger {
	x.e.Bool(key, b)
	return x
}

func (x *l) Bools(key string, b []bool) Logger {
	x.e.Bools(key, b)
	return x
}

func (x *l) Bytes(key string, val []byte) Logger {
	x.e.Bytes(key, val)
	return x
}

func (x *l) Err(err error) Logger {
	x.e.Err(err)
	return x
}

func (x *l) Errs(key string, errs []error) Logger {
	x.e.Errs(key, errs)
	return x
}

func (x *l) Float32(key string, f float32) Logger {
	x.e.Float32(key, f)
	return x
}

func (x *l) Floats32(key string, f []float32) Logger {
	x.e.Floats32(key, f)
	return x
}

func (x *l) Float64(key string, f float64) Logger {
	x.e.Float64(key, f)
	return x
}

func (x *l) Floats64(key string, f []float64) Logger {
	x.e.Floats64(key, f)
	return x
}

func (x *l) Hex(key string, val []byte) Logger {
	x.e.Hex(key, val)
	return x
}

func (x *l) Int(key string, i int) Logger {
	x.e.Int(key, i)
	return x
}

func (x *l) Ints(key string, i []int) Logger {
	x.e.Ints(key, i)
	return x
}

func (x *l) Int8(key string, i int8) Logger {
	x.e.Int8(key, i)
	return x
}

func (x *l) Ints8(key string, i []int8) Logger {
	x.e.Ints8(key, i)
	return x
}

func (x *l) Int16(key string, i int16) Logger {
	x.e.Int16(key, i)
	return x
}

func (x *l) Ints16(key string, i []int16) Logger {
	x.e.Ints16(key, i)
	return x
}

func (x *l) Int32(key string, i int32) Logger {
	x.e.Int32(key, i)
	return x
}

func (x *l) Ints32(key string, i []int32) Logger {
	x.e.Ints32(key, i)
	return x
}

func (x *l) Int64(key string, i int64) Logger {
	x.e.Int64(key, i)
	return x
}

func (x *l) Ints64(key string, i []int64) Logger {
	x.e.Ints64(key, i)
	return x
}

func (x *l) RawJSON(key string, b []byte) Logger {
	x.e.RawJSON(key, b)
	return x
}

func (x *l) Stack() Logger {
	x.e.Stack()
	return x
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

func (x *l) Uint(key string, i uint) Logger {
	x.e.Uint(key, i)
	return x
}

func (x *l) Uints(key string, i []uint) Logger {
	x.e.Uints(key, i)
	return x
}

func (x *l) Uint8(key string, i uint8) Logger {
	x.e.Uint8(key, i)
	return x
}

func (x *l) Uints8(key string, i []uint8) Logger {
	x.e.Uints8(key, i)
	return x
}

func (x *l) Uint16(key string, i uint16) Logger {
	x.e.Uint16(key, i)
	return x
}

func (x *l) Uints16(key string, i []uint16) Logger {
	x.e.Uints16(key, i)
	return x
}

func (x *l) Uint32(key string, i uint32) Logger {
	x.e.Uint32(key, i)
	return x
}

func (x *l) Uints32(key string, i []uint32) Logger {
	x.e.Uints32(key, i)
	return x
}

func (x *l) Uint64(key string, i uint64) Logger {
	x.e.Uint64(key, i)
	return x
}

func (x *l) Uints64(key string, i []uint64) Logger {
	x.e.Uints64(key, i)
	return x
}

func (x *l) Interface(key string, i interface{}) Logger {
	x.e.Interface(key, i)
	return x
}

func (x *l) Any(key string, a any) Logger {
	x.e.Any(key, a)
	return x
}
