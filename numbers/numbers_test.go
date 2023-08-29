package numbers_test

import (
	"github.com/sinhashubham95/go-utils/numbers"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
