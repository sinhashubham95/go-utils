package set_test

import (
	"github.com/sinhashubham95/go-utils/structures/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	s := set.New[int]()
	assert.True(t, s.Add(1))
	assert.False(t, s.Add(1))
	assert.Equal(t, 2, s.Append(1, 2, 3))
	assert.Zero(t, s.Append(1, 2, 3))
	s.Clear()
	assert.Zero(t, s.Length())
	assert.True(t, s.Add(1))
	o := s.Clone()
	assert.Equal(t, 1, o.Length())
	assert.False(t, o.Add(1))
	assert.Equal(t, []int{1}, s.Collection())
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(1, 2))
	assert.True(t, s.ContainsAny(1))
	assert.True(t, s.ContainsAny(1, 2))
	assert.False(t, s.ContainsAny(2))
	d := s.Difference(o)
	assert.Zero(t, d.Length())
	assert.True(t, s.Equal(o))
	s.Append(2, 3)
	d = s.Difference(o)
	assert.Equal(t, 2, d.Length())
	assert.False(t, s.Equal(o))
	i := s.Intersection(o)
	assert.Equal(t, 1, i.Length())
	assert.True(t, o.Equal(i))
	s.Remove(2)
	assert.Equal(t, 2, s.Length())
	s.RemoveAll(3)
	assert.Equal(t, 1, s.Length())
	assert.True(t, o.Add(2))
	u := s.Union(o)
	assert.Equal(t, 2, u.Length())
	assert.False(t, s.Equal(u))
	assert.True(t, o.Equal(u))
	assert.True(t, s.Add(3))
	assert.False(t, s.Equal(o))
	assert.True(t, o.Add(4))
	i = s.Intersection(o)
	assert.Equal(t, 1, i.Length())
}

func TestSetIterator(t *testing.T) {
	s := set.New[int]()
	assert.Equal(t, 3, s.Append(1, 2, 3))
	i := s.Iterator()
	defer i.Close()
	for v := range i.Elements() {
		assert.True(t, s.Contains(v))
	}
	ai := s.Iterator()
	ai.Close()
}
