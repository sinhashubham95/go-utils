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

func TestCollate(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5, 9}, collections.Collate([]int{1, 5, 9}, []int{2, 3}))
	assert.Equal(t, []int{1, 1, 1, 1, 2, 3, 5, 5, 5, 9}, collections.Collate([]int{1, 5, 1, 9, 1, 5}, []int{2, 1, 5, 3}))
}

func TestCollateWithComparator(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5, 9}, collections.CollateWithComparator([]int{1, 5, 9}, []int{2, 3}, func(x, y int) bool { return x < y }))
	assert.Equal(t, []int{1, 1, 1, 1, 2, 3, 5, 5, 5, 9}, collections.CollateWithComparator([]int{1, 5, 1, 9, 1, 5},
		[]int{2, 1, 5, 3}, func(x, y int) bool { return x < y }))
}

func TestCollateRemovingDuplicates(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5, 9}, collections.CollateRemovingDuplicates([]int{1, 5, 9}, []int{2, 3}))
	assert.Equal(t, []int{1, 2, 3, 5, 9}, collections.CollateRemovingDuplicates([]int{1, 5, 1, 9, 1, 5}, []int{2, 1, 5, 3}))
}

func TestCollateWithComparatorRemovingDuplicates(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5, 9}, collections.CollateWithComparatorRemovingDuplicates([]int{1, 5, 9}, []int{2, 3},
		func(x, y int) bool { return x < y }))
	assert.Equal(t, []int{1, 2, 3, 5, 9}, collections.CollateWithComparatorRemovingDuplicates([]int{1, 5, 1, 9, 1, 5},
		[]int{2, 1, 5, 3}, func(x, y int) bool { return x < y }))
}

func TestCollect(t *testing.T) {
	assert.Equal(t, []int{4, 5, 6}, collections.Collect([]int{1, 2, 3}, func(a int) int { return a + 3 }))
	assert.Panics(t, func() { collections.Collect[int](nil, nil) })
	assert.Equal(t, []int{}, collections.Collect(nil, func(a int) int { return a + 3 }))
}

func TestContains(t *testing.T) {
	assert.True(t, collections.Contains([]int{1, 2, 3}, 1))
	assert.True(t, collections.Contains([]int{1, 2, 3}, 3))
	assert.False(t, collections.Contains([]int{1, 2, 3}, 5))
}

func TestContainsWithEquator(t *testing.T) {
	assert.True(t, collections.ContainsWithEquator([]int{1, 2, 3}, 1, func(a, b int) bool { return a == b }))
	assert.True(t, collections.ContainsWithEquator([]int{1, 2, 3}, 3, func(a, b int) bool { return a == b }))
	assert.False(t, collections.ContainsWithEquator([]int{1, 2, 3}, 5, func(a, b int) bool { return a == b }))
}

func TestContainsAll(t *testing.T) {
	assert.True(t, collections.ContainsAll([]int{1, 2, 3}, []int{1, 2}))
	assert.False(t, collections.ContainsAll([]int{1, 2, 3}, []int{1, 5}))
}

func TestContainsAny(t *testing.T) {
	assert.True(t, collections.ContainsAny([]int{1, 2, 3}, []int{1, 2}))
	assert.True(t, collections.ContainsAny([]int{1, 2, 3}, []int{1, 5}))
	assert.False(t, collections.ContainsAny([]int{1, 2, 3}, []int{6, 5}))
}

func TestCopy(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, collections.Copy([]int{1, 2, 3}))
	assert.Equal(t, []int{}, collections.Copy[int](nil))
}

func TestCount(t *testing.T) {
	assert.Equal(t, 2, collections.Count([]int{1, 5, 9, 1, 2, 3}, 1))
	assert.Equal(t, 0, collections.Count([]int{1, 5, 9, 1, 2, 3}, 11))
}

func TestCountMatches(t *testing.T) {
	assert.Equal(t, 2, collections.CountMatches([]int{1, 5, 9, 1, 2, 3}, func(x int) bool { return x == 1 }))
	assert.Equal(t, 0, collections.CountMatches([]int{1, 5, 9, 1, 2, 3}, func(x int) bool { return x == 11 }))
}

func TestDisjunction(t *testing.T) {
	//assert.Equal(t, []int{1, 5, 9, 2, 3}, collections.Disjunction([]int{1, 5, 9, 6, 7}, []int{2, 6, 7, 3}))
	assert.Equal(t, []int{}, collections.Disjunction[int](nil, nil))
	assert.Equal(t, []int{}, collections.Disjunction([]int{1, 2, 3}, []int{1, 2, 3}))
}
