package stack_test

import (
	"testing"

	"github.com/sinhashubham95/go-utils/structures/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack.New[int]()
	v, b := s.Pop()
	assert.Zero(t, v)
	assert.False(t, b)
	v, b = s.Peek()
	assert.Zero(t, v)
	assert.False(t, b)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	v, b = s.Pop()
	assert.Equal(t, 3, v)
	assert.True(t, b)
	v, b = s.Peek()
	assert.Equal(t, 2, v)
	assert.True(t, b)
	assert.Equal(t, 2, s.Length())
}
