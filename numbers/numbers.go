package numbers

import (
	"math"
	"strconv"
)

// floating point limit values
const (
	MaxFloat32             float32 = 0x1p127 * (1 + (1 - 0x1p-23))  // 3.40282346638528859811704183484516925440e+38
	SmallestNonZeroFloat32 float32 = 0x1p-126 * 0x1p-23             // 1.401298464324817070923729583289916131280e-45
	MaxFloat64             float64 = 0x1p1023 * (1 + (1 - 0x1p-52)) // 1.79769313486231570814527423731704356798070e+308
	SmallestNonZeroFloat64 float64 = 0x1p-1022 * 0x1p-52            // 4.9406564584124654417656879286822137236505980e-324
)

// integer limit values
const (
	IntSize          = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt    int    = 1<<(IntSize-1) - 1     // MaxInt32 or MaxInt64 depending on intSize.
	MinInt    int    = -1 << (IntSize - 1)    // MinInt32 or MinInt64 depending on intSize.
	MaxInt8   int8   = 1<<7 - 1               // 127
	MinInt8   int8   = -1 << 7                // -128
	MaxInt16  int16  = 1<<15 - 1              // 32767
	MinInt16  int16  = -1 << 15               // -32768
	MaxInt32  int32  = 1<<31 - 1              // 2147483647
	MinInt32  int32  = -1 << 31               // -2147483648
	MaxInt64  int64  = 1<<63 - 1              // 9223372036854775807
	MinInt64  int64  = -1 << 63               // -9223372036854775808
	MaxUint   uint   = 1<<IntSize - 1         // MaxUint32 or MaxUint64 depending on intSize.
	MaxUint8  uint8  = 1<<8 - 1               // 255
	MaxUint16 uint16 = 1<<16 - 1              // 65535
	MaxUint32 uint32 = 1<<32 - 1              // 4294967295
	MaxUint64 uint64 = 1<<64 - 1              // 18446744073709551615
)

// Number is the generic type for interconvertible numbers.
type Number interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 | ~float32 | ~float64 | ~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

// SNumber is the generic type for interconvertible signed numbers.
type SNumber interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 | ~float32 | ~float64
}

// Number64 is the generic type for 64-bit numbers.
type Number64 interface {
	~int64 | ~float64
}

// FloatingNumber is the generic type for floating point numbers.
type FloatingNumber interface {
	~float32 | ~float64
}

// IntegerNumber is the generic type for integers.
type IntegerNumber interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

// Compare is used to compare 2 numbers.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
func Compare[K Number](a, b K) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// Equals is used to check if the 2 numbers are equal or not
func Equals[K Number](a, b K) bool {
	return a == b
}

// HashCode is used to compute the hash values for the given number.
func HashCode[K Number](a K) int32 {
	switch any(a).(type) {
	case int, uint, int64, uint64:
		v := int64(a)
		return (int32)(v ^ (v >> 32))
	case float32:
		return int32(math.Float32bits(float32(a)))
	case float64:
		return int32(math.Float64bits(float64(a)))
	default:
		return int32(a)
	}
}

// NumberToNumber is used to convert the primitive number types to another number type.
func NumberToNumber[K Number, V Number](a K) V {
	return V(a)
}

// StringToNumber is used to convert the string to a number.
func StringToNumber[K Number](a string) (r K, err error) {
	switch any(r).(type) {
	case int8:
		t, e := strconv.ParseInt(a, 10, 8)
		err = e
		r = K(t)
	case int16:
		t, e := strconv.ParseInt(a, 10, 16)
		err = e
		r = K(t)
	case int32:
		t, e := strconv.ParseInt(a, 10, 32)
		err = e
		r = K(t)
	case int:
		t, e := strconv.ParseInt(a, 10, 0)
		err = e
		r = K(t)
	case int64:
		t, e := strconv.ParseInt(a, 10, 64)
		err = e
		r = K(t)
	case float32:
		t, e := strconv.ParseFloat(a, 32)
		err = e
		r = K(t)
	case float64:
		t, e := strconv.ParseFloat(a, 64)
		err = e
		r = K(t)
	case uint8:
		t, e := strconv.ParseUint(a, 10, 8)
		err = e
		r = K(t)
	case uint16:
		t, e := strconv.ParseUint(a, 10, 16)
		err = e
		r = K(t)
	case uint32:
		t, e := strconv.ParseUint(a, 10, 32)
		err = e
		r = K(t)
	case uint:
		t, e := strconv.ParseUint(a, 10, 0)
		err = e
		r = K(t)
	case uint64:
		t, e := strconv.ParseUint(a, 10, 64)
		err = e
		r = K(t)
	}
	return
}

// StringToNumberWithBits is used to convert the string to a number represented within the specified number of bits.
func StringToNumberWithBits[K Number64](a string, bitSize int) (r K, err error) {
	switch any(r).(type) {
	case int64:
		t, e := strconv.ParseInt(a, 10, bitSize)
		err = e
		r = K(t)
	case float64:
		t, e := strconv.ParseFloat(a, bitSize)
		err = e
		r = K(t)
	}
	return
}

// StringToIntegerNumberWithBaseAndBits is used to convert the string to a number in the given base represented
// within the specified number of bits.
func StringToIntegerNumberWithBaseAndBits[K IntegerNumber](a string, base, bitSize int) (K, error) {
	r, err := strconv.ParseInt(a, base, bitSize)
	if err != nil {
		return 0, err
	}
	return K(r), nil
}
