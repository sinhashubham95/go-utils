package pair_test

import (
	"github.com/sinhashubham95/go-utils/structures/pair"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPair(t *testing.T) {
	p := pair.New[int, int](0, 0)
	assert.Zero(t, p.GetFirst())
	assert.Zero(t, p.GetSecond())
	p.SetFirst(1)
	p.SetSecond(2)
	assert.Equal(t, 1, p.GetFirst())
	assert.Equal(t, 2, p.GetSecond())
}

func TestPairWithBuilder(t *testing.T) {
	p := pair.Builder[int, int]().Build()
	assert.Zero(t, p.GetFirst())
	assert.Zero(t, p.GetSecond())
	p = pair.Builder[int, int]().First(1).Second(2).Build()
	assert.Equal(t, 1, p.GetFirst())
	assert.Equal(t, 2, p.GetSecond())
	b := &pair.B[int, int]{}
	b.First(1)
	b.Second(2)
	p = b.Build()
	assert.Equal(t, 1, p.GetFirst())
	assert.Equal(t, 2, p.GetSecond())
	b = &pair.B[int, int]{}
	b.Second(2)
	b.First(1)
	p = b.Build()
	assert.Equal(t, 1, p.GetFirst())
	assert.Equal(t, 2, p.GetSecond())
	b = &pair.B[int, int]{}
	p = b.Build()
	assert.Zero(t, p.GetFirst())
	assert.Zero(t, p.GetSecond())
}

func TestPairFromCollection(t *testing.T) {
	p := pair.NewFromCollection[int](nil)
	assert.Nil(t, p)
	p = pair.NewFromCollection([]int{1, 2})
	assert.Equal(t, 1, p.GetFirst())
	assert.Equal(t, 2, p.GetSecond())
}
