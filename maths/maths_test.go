package maths_test

import (
	"github.com/sinhashubham95/go-utils/maths"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, maths.Abs(-1))
}

func TestACos(t *testing.T) {
	assert.Equal(t, 0.6435011087932843, maths.ACos(0.8))
}

func TestACosH(t *testing.T) {
	assert.Equal(t, 5.07513475044481, maths.ACosH(80.0))
}

func TestASin(t *testing.T) {
	assert.Equal(t, 0.9272952180016123, maths.ASin(0.8))
}

func TestASinH(t *testing.T) {
	assert.Equal(t, 5.0752128754452075, maths.ASinH(80.0))
}

func TestATan(t *testing.T) {
	assert.Equal(t, 0.6747409422235526, maths.ATan(0.8))
}

func TestATanXY(t *testing.T) {
	assert.Equal(t, 0.8960553845713439, maths.ATanXY(1.2, 1.5))
}

func TestATanH(t *testing.T) {
	assert.Equal(t, 0.10033534773107558, maths.ATanH(0.1))
}
