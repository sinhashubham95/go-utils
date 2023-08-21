package collections_test

import (
	"github.com/sinhashubham95/go-utils/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAll(t *testing.T) {
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAll([]string{"naruto"}, "rocks"))
	assert.Equal(t, []string{"naruto", "rocks", ""}, collections.AddAll([]string{"naruto"}, "rocks", ""))
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAll([]string{"naruto", "rocks"}))
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAll(nil, "naruto", "rocks"))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAll([]int{1, 2}, 3))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAll([]int{1}, 2, 3))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAll(nil, 1, 2, 3))
	assert.Equal(t, []int{1, 2, 3, 0}, collections.AddAll(nil, 1, 2, 3, 0))
}

func TestAddAllIgnoringEmpty(t *testing.T) {
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAllIgnoringEmpty([]string{"naruto"}, "rocks"))
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAllIgnoringEmpty([]string{"naruto"}, "rocks", ""))
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAllIgnoringEmpty([]string{"naruto", "rocks"}))
	assert.Equal(t, []string{"naruto", "rocks"}, collections.AddAllIgnoringEmpty(nil, "naruto", "rocks"))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAllIgnoringEmpty([]int{1, 2}, 3))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAllIgnoringEmpty([]int{1}, 2, 3))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAllIgnoringEmpty(nil, 1, 2, 3))
	assert.Equal(t, []int{1, 2, 3}, collections.AddAllIgnoringEmpty(nil, 1, 2, 3, 0))
	assert.Equal(t, []bool{true, true, true}, collections.AddAllIgnoringEmpty(nil, true, true, false, false, true, false))
}

func TestCardinality(t *testing.T) {
	assert.Equal(t, 3, collections.Cardinality([]string{"naruto", "naruto", "naruto", "rocks"}, "naruto"))
}

func TestCardinalityWithEquator(t *testing.T) {
	assert.Equal(t, 3, collections.CardinalityWithEquator([]string{"naruto", "naruto", "naruto", "rocks"},
		"naruto", func(a, b string) bool { return a == b }))
}

func TestCardinalityMap(t *testing.T) {
	assert.Equal(t, map[string]int{"naruto": 3, "rocks": 1},
		collections.CardinalityMap([]string{"naruto", "naruto", "rocks", "naruto"}))
}

func TestChain(t *testing.T) {
	assert.Equal(t, []string{"naruto", "rocks"}, collections.Chain([]string{"naruto"}, []string{"rocks"}))
	assert.Equal(t, []string{}, collections.Chain[string]())
}
