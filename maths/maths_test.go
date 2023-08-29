package maths_test

import (
	"github.com/sinhashubham95/go-utils/maths"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, maths.Abs(-1))
}
