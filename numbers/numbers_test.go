package numbers_test

import (
	"testing"

	"github.com/sinhashubham95/go-utils/numbers"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, float32(3.4028235e+38), numbers.MaxFloat32)
	assert.Equal(t, float32(1e-45), numbers.SmallestNonZeroFloat32)
	assert.Equal(t, 1.7976931348623157e+308, numbers.MaxFloat64)
	assert.Equal(t, 5e-324, numbers.SmallestNonZeroFloat64)

	assert.Equal(t, 64, numbers.IntSize)
	assert.Equal(t, 9223372036854775807, numbers.MaxInt)
	assert.Equal(t, -9223372036854775808, numbers.MinInt)
	assert.Equal(t, int8(127), numbers.MaxInt8)
	assert.Equal(t, int8(-128), numbers.MinInt8)
	assert.Equal(t, int16(32767), numbers.MaxInt16)
	assert.Equal(t, int16(-32768), numbers.MinInt16)
	assert.Equal(t, int32(2147483647), numbers.MaxInt32)
	assert.Equal(t, int32(-2147483648), numbers.MinInt32)
	assert.Equal(t, int64(9223372036854775807), numbers.MaxInt64)
	assert.Equal(t, int64(-9223372036854775808), numbers.MinInt64)
	assert.Equal(t, uint8(0xff), numbers.MaxUint8)
	assert.Equal(t, uint16(0xffff), numbers.MaxUint16)
	assert.Equal(t, uint32(0xffffffff), numbers.MaxUint32)
	assert.Equal(t, uint(0xffffffffffffffff), numbers.MaxUint)
	assert.Equal(t, uint64(0xffffffffffffffff), numbers.MaxUint64)
}

func TestCompare(t *testing.T) {
	assert.Equal(t, 0, numbers.Compare[uint8](1, 1))
	assert.Equal(t, 0, numbers.Compare[uint16](1, 1))
	assert.Equal(t, 0, numbers.Compare[uint32](1, 1))
	assert.Equal(t, 0, numbers.Compare[uint](1, 1))
	assert.Equal(t, 0, numbers.Compare[uint64](1, 1))
	assert.Equal(t, 0, numbers.Compare[int8](1, 1))
	assert.Equal(t, 0, numbers.Compare[int16](1, 1))
	assert.Equal(t, 0, numbers.Compare[int32](1, 1))
	assert.Equal(t, 0, numbers.Compare[int](1, 1))
	assert.Equal(t, 0, numbers.Compare[int64](1, 1))
	assert.Equal(t, 0, numbers.Compare[float32](1, 1))
	assert.Equal(t, 0, numbers.Compare[float64](1, 1))

	assert.Equal(t, -1, numbers.Compare[uint8](1, 2))
	assert.Equal(t, -1, numbers.Compare[uint16](1, 2))
	assert.Equal(t, -1, numbers.Compare[uint32](1, 2))
	assert.Equal(t, -1, numbers.Compare[uint](1, 2))
	assert.Equal(t, -1, numbers.Compare[uint64](1, 2))
	assert.Equal(t, -1, numbers.Compare[int8](1, 2))
	assert.Equal(t, -1, numbers.Compare[int16](1, 2))
	assert.Equal(t, -1, numbers.Compare[int32](1, 2))
	assert.Equal(t, -1, numbers.Compare[int](1, 2))
	assert.Equal(t, -1, numbers.Compare[int64](1, 2))
	assert.Equal(t, -1, numbers.Compare[float32](1, 2))
	assert.Equal(t, -1, numbers.Compare[float64](1, 2))

	assert.Equal(t, 1, numbers.Compare[uint8](2, 1))
	assert.Equal(t, 1, numbers.Compare[uint16](2, 1))
	assert.Equal(t, 1, numbers.Compare[uint32](2, 1))
	assert.Equal(t, 1, numbers.Compare[uint](2, 1))
	assert.Equal(t, 1, numbers.Compare[uint64](2, 1))
	assert.Equal(t, 1, numbers.Compare[int8](2, 1))
	assert.Equal(t, 1, numbers.Compare[int16](2, 1))
	assert.Equal(t, 1, numbers.Compare[int32](2, 1))
	assert.Equal(t, 1, numbers.Compare[int](2, 1))
	assert.Equal(t, 1, numbers.Compare[int64](2, 1))
	assert.Equal(t, 1, numbers.Compare[float32](2, 1))
	assert.Equal(t, 1, numbers.Compare[float64](2, 1))
}

func TestEquals(t *testing.T) {
	assert.True(t, numbers.Equals[uint8](1, 1))
	assert.True(t, numbers.Equals[uint16](1, 1))
	assert.True(t, numbers.Equals[uint32](1, 1))
	assert.True(t, numbers.Equals[uint](1, 1))
	assert.True(t, numbers.Equals[uint64](1, 1))
	assert.True(t, numbers.Equals[int8](1, 1))
	assert.True(t, numbers.Equals[int16](1, 1))
	assert.True(t, numbers.Equals[int32](1, 1))
	assert.True(t, numbers.Equals[int](1, 1))
	assert.True(t, numbers.Equals[int64](1, 1))
	assert.True(t, numbers.Equals[float32](1, 1))
	assert.True(t, numbers.Equals[float64](1, 1))

	assert.False(t, numbers.Equals[uint8](1, 2))
	assert.False(t, numbers.Equals[uint16](1, 2))
	assert.False(t, numbers.Equals[uint32](1, 2))
	assert.False(t, numbers.Equals[uint](1, 2))
	assert.False(t, numbers.Equals[uint64](1, 2))
	assert.False(t, numbers.Equals[int8](1, 2))
	assert.False(t, numbers.Equals[int16](1, 2))
	assert.False(t, numbers.Equals[int32](1, 2))
	assert.False(t, numbers.Equals[int](1, 2))
	assert.False(t, numbers.Equals[int64](1, 2))
	assert.False(t, numbers.Equals[float32](1, 2))
	assert.False(t, numbers.Equals[float64](1, 2))
}

func TestHashCode(t *testing.T) {
	assert.Equal(t, int32(1), numbers.HashCode[uint8](1))
	assert.Equal(t, int32(1), numbers.HashCode[uint16](1))
	assert.Equal(t, int32(1), numbers.HashCode[uint32](1))
	assert.Equal(t, int32(1), numbers.HashCode[uint](1))
	assert.Equal(t, int32(1), numbers.HashCode[uint64](1))
	assert.Equal(t, int32(1), numbers.HashCode[int8](1))
	assert.Equal(t, int32(1), numbers.HashCode[int16](1))
	assert.Equal(t, int32(1), numbers.HashCode[int32](1))
	assert.Equal(t, int32(1), numbers.HashCode[int](1))
	assert.Equal(t, int32(1), numbers.HashCode[int64](1))
	assert.Equal(t, int32(1071225242), numbers.HashCode[float32](1.7))
	assert.Equal(t, int32(858993459), numbers.HashCode[float64](1.7))
}

func TestNumberToNumber(t *testing.T) {
	assert.Equal(t, uint8(1), numbers.NumberToNumber[uint32, uint8](1))
}

func TestStringToNumber(t *testing.T) {
	var a any
	var err error
	a, err = numbers.StringToNumber[uint8]("1")
	assert.Equal(t, uint8(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[uint8]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[uint16]("1")
	assert.Equal(t, uint16(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[uint16]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[uint32]("1")
	assert.Equal(t, uint32(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[uint32]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[uint]("1")
	assert.Equal(t, uint(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[uint]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[uint64]("1")
	assert.Equal(t, uint64(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[uint64]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[int8]("1")
	assert.Equal(t, int8(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[int8]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[int16]("1")
	assert.Equal(t, int16(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[int16]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[int32]("1")
	assert.Equal(t, int32(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[int32]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[int]("1")
	assert.Equal(t, 1, a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[int]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[int64]("1")
	assert.Equal(t, int64(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[int64]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[float32]("1.5")
	assert.Equal(t, float32(1.5), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[float32]("abc")
	assert.Error(t, err)
	a, err = numbers.StringToNumber[float64]("1.5")
	assert.Equal(t, 1.5, a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumber[float64]("abc")
	assert.Error(t, err)
}

func TestStringToNumberWithBits(t *testing.T) {
	var a any
	var err error
	a, err = numbers.StringToNumberWithBits[int64]("1", 64)
	assert.Equal(t, int64(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumberWithBits[int64]("abc", 64)
	assert.Error(t, err)
	a, err = numbers.StringToNumberWithBits[float64]("1.5", 64)
	assert.Equal(t, 1.5, a)
	assert.NoError(t, err)
	_, err = numbers.StringToNumberWithBits[float64]("abc", 64)
	assert.Error(t, err)
}

func TestStringToIntegerNumberWithBaseAndBits(t *testing.T) {
	var a any
	var err error
	a, err = numbers.StringToIntegerNumberWithBaseAndBits[int64]("1", 10, 64)
	assert.Equal(t, int64(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToIntegerNumberWithBaseAndBits[int64]("abc", 10, 64)
	assert.Error(t, err)
	a, err = numbers.StringToIntegerNumberWithBaseAndBits[int8]("1", 10, 64)
	assert.Equal(t, int8(1), a)
	assert.NoError(t, err)
	_, err = numbers.StringToIntegerNumberWithBaseAndBits[int8]("abc", 10, 64)
	assert.Error(t, err)
}

func TestIntegerNumberToString(t *testing.T) {
	assert.Equal(t, "1", numbers.IntegerNumberToString(1))
}

func TestFloatingNumberToString(t *testing.T) {
	assert.Equal(t, "1.10", numbers.FloatingNumberToString(float32(1.1), 'f', 2))
	assert.Equal(t, "1.10", numbers.FloatingNumberToString(1.1, 'f', 2))
}
