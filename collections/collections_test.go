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
	assert.Equal(t, []int{1, 5, 9, 2, 3}, collections.Disjunction([]int{1, 5, 9, 6, 7}, []int{2, 6, 7, 3}))
	assert.Equal(t, []int{1, 1}, collections.Disjunction([]int{1, 2, 1, 3, 1, 2}, []int{1, 2, 2, 3}))
	assert.Equal(t, []int{}, collections.Disjunction[int](nil, nil))
	assert.Equal(t, []int{}, collections.Disjunction([]int{1, 2, 3}, []int{1, 2, 3}))
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, []int{}, collections.Empty[int]())
}

func TestEmptyBySize(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0}, collections.EmptyBySize[int](3))
}

func TestEmptyBySizeAndCapacity(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0}, collections.EmptyBySizeAndCapacity[int](3, 6))
}

func TestEmptyIfNil(t *testing.T) {
	assert.Equal(t, []int{}, collections.EmptyIfNil[int](nil))
	assert.Equal(t, []int{1, 2, 3}, collections.EmptyIfNil[int]([]int{1, 2, 3}))
}

func TestExtractSingleton(t *testing.T) {
	assert.Panics(t, func() {
		collections.ExtractSingleton([]int{1, 2, 3})
	})
	assert.Equal(t, 1, collections.ExtractSingleton([]int{1}))
}

func TestFilter(t *testing.T) {
	assert.Panics(t, func() {
		collections.Filter([]int{1, 2, 3}, nil)
	})
	assert.Equal(t, []int{1, 1}, collections.Filter([]int{1, 5, 9, 2, 1, 3}, func(x int) bool { return x == 1 }))
}

func TestFilterInverse(t *testing.T) {
	assert.Panics(t, func() {
		collections.FilterInverse([]int{1, 2, 3}, nil)
	})
	assert.Equal(t, []int{5, 9, 2, 3}, collections.FilterInverse([]int{1, 5, 9, 2, 1, 3}, func(x int) bool { return x == 1 }))
}

func TestFind(t *testing.T) {
	assert.Panics(t, func() {
		collections.Find([]int{1, 2, 3}, nil)
	})
	x, b := collections.Find([]int{1, 2, 3}, func(x int) bool { return x == 1 })
	assert.Equal(t, 1, x)
	assert.True(t, b)
	x, b = collections.Find([]int{1, 2, 3}, func(x int) bool { return x == 9 })
	assert.Equal(t, 0, x)
	assert.False(t, b)
	x, b = collections.Find(nil, func(x int) bool { return x == 9 })
	assert.Equal(t, 0, x)
	assert.False(t, b)
}

func TestFindInverse(t *testing.T) {
	assert.Panics(t, func() {
		collections.FindInverse([]int{1, 2, 3}, nil)
	})
	x, b := collections.FindInverse([]int{1, 2, 3}, func(x int) bool { return x == 1 })
	assert.Equal(t, 2, x)
	assert.True(t, b)
	x, b = collections.FindInverse([]int{1, 2, 3}, func(x int) bool { return x == 1 || x == 2 || x == 3 })
	assert.Equal(t, 0, x)
	assert.False(t, b)
	x, b = collections.FindInverse(nil, func(x int) bool { return x == 9 })
	assert.Equal(t, 0, x)
	assert.False(t, b)
}

func TestFindOrDefault(t *testing.T) {
	assert.Panics(t, func() {
		collections.FindOrDefault([]int{1, 2, 3}, nil, 9)
	})
	x, b := collections.FindOrDefault([]int{1, 2, 3}, func(x int) bool { return x == 1 }, 5)
	assert.Equal(t, 1, x)
	assert.True(t, b)
	x, b = collections.FindOrDefault([]int{1, 2, 3}, func(x int) bool { return x == 9 }, 5)
	assert.Equal(t, 5, x)
	assert.False(t, b)
	x, b = collections.FindOrDefault(nil, func(x int) bool { return x == 9 }, 5)
	assert.Equal(t, 5, x)
	assert.False(t, b)
}

func TestFindInverseOrDefault(t *testing.T) {
	assert.Panics(t, func() {
		collections.FindInverseOrDefault([]int{1, 2, 3}, nil, 9)
	})
	x, b := collections.FindInverseOrDefault([]int{1, 2, 3}, func(x int) bool { return x == 1 }, 5)
	assert.Equal(t, 2, x)
	assert.True(t, b)
	x, b = collections.FindInverseOrDefault([]int{1, 2, 3}, func(x int) bool { return x == 1 || x == 2 || x == 3 }, 5)
	assert.Equal(t, 5, x)
	assert.False(t, b)
	x, b = collections.FindInverseOrDefault(nil, func(x int) bool { return x == 9 }, 5)
	assert.Equal(t, 5, x)
	assert.False(t, b)
}

func TestFirst(t *testing.T) {
	x, b := collections.First([]int{1, 2, 3})
	assert.Equal(t, 1, x)
	assert.True(t, b)
	x, b = collections.First[int](nil)
	assert.Equal(t, 0, x)
	assert.False(t, b)
	x, b = collections.First([]int{})
	assert.Equal(t, 0, x)
	assert.False(t, b)
}

func TestFirstOrDefault(t *testing.T) {
	x, b := collections.FirstOrDefault([]int{1, 2, 3}, 5)
	assert.Equal(t, 1, x)
	assert.True(t, b)
	x, b = collections.FirstOrDefault[int](nil, 5)
	assert.Equal(t, 5, x)
	assert.False(t, b)
	x, b = collections.FirstOrDefault([]int{}, 5)
	assert.Equal(t, 5, x)
	assert.False(t, b)
}

func TestForEach(t *testing.T) {
	a := 0
	collections.ForEach([]int{1, 2, 3}, func(x int) { a += x })
	assert.Equal(t, 6, a)
	collections.ForEach(nil, func(x int) { a += x })
	assert.Equal(t, 6, a)
	assert.Panics(t, func() {
		collections.ForEach[int](nil, nil)
	})
}

func TestForEachButLast(t *testing.T) {
	a := 0
	collections.ForEachButLast([]int{1, 2, 3}, func(x int) { a += x })
	assert.Equal(t, 3, a)
	collections.ForEachButLast(nil, func(x int) { a += x })
	assert.Equal(t, 3, a)
	assert.Panics(t, func() {
		collections.ForEachButLast[int](nil, nil)
	})
}

func TestGet(t *testing.T) {
	assert.Equal(t, 3, collections.Get([]int{1, 2, 3}, 2))
	assert.Equal(t, 0, collections.Get([]int{1, 2, 3}, 7))
	assert.Equal(t, 0, collections.Get([]int{1, 2, 3}, -1))
}

func TestGetOrDefault(t *testing.T) {
	assert.Equal(t, 3, collections.GetOrDefault([]int{1, 2, 3}, 2, 5))
	assert.Equal(t, 5, collections.GetOrDefault([]int{1, 2, 3}, 7, 5))
	assert.Equal(t, 5, collections.GetOrDefault([]int{1, 2, 3}, -1, 5))
}

func TestIndexOf(t *testing.T) {
	assert.Equal(t, 2, collections.IndexOf([]int{1, 2, 3}, 3))
	assert.Equal(t, -1, collections.IndexOf([]int{1, 2, 3}, 5))
}

func TestIndexOfWithPredicate(t *testing.T) {
	assert.Panics(t, func() {
		collections.IndexOfWithPredicate[int](nil, nil)
	})
	assert.Equal(t, 2, collections.IndexOfWithPredicate([]int{1, 2, 3}, func(x int) bool { return x == 3 }))
	assert.Equal(t, -1, collections.IndexOfWithPredicate([]int{1, 2, 3}, func(x int) bool { return x == 5 }))
}

func TestIntersection(t *testing.T) {
	assert.Equal(t, []int{1, 2, 1}, collections.Intersection([]int{1, 5, 9, 2, 1, 3}, []int{1, 1, 2}))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, collections.IsEmpty[int](nil))
	assert.True(t, collections.IsEmpty([]int{}))
	assert.False(t, collections.IsEmpty([]int{1, 2, 3}))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, collections.IsNotEmpty[int](nil))
	assert.False(t, collections.IsNotEmpty([]int{}))
	assert.True(t, collections.IsNotEmpty([]int{1, 2, 3}))
}

func TestIsEqual(t *testing.T) {
	assert.True(t, collections.IsEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.False(t, collections.IsEqual([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}))
	assert.True(t, collections.IsEqual[int](nil, nil))
}

func TestIsEqualWithEquator(t *testing.T) {
	assert.True(t, collections.IsEqualWithEquator([]int{1, 2, 3}, []int{1, 2, 3}, func(x, y int) bool { return x == y }))
	assert.False(t, collections.IsEqualWithEquator([]int{1, 2, 3}, []int{1, 2, 3, 4, 5}, func(x, y int) bool { return x == y }))
	assert.True(t, collections.IsEqualWithEquator(nil, nil, func(x, y int) bool { return x == y }))
	assert.Panics(t, func() {
		collections.IsEqualWithEquator[int](nil, nil, nil)
	})
}

func TestIsSubCollection(t *testing.T) {
	assert.True(t, collections.IsSubCollection([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.True(t, collections.IsSubCollection([]int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}))
	assert.False(t, collections.IsSubCollection([]int{1, 2, 3}, []int{1, 2}))
	assert.True(t, collections.IsSubCollection[int](nil, nil))
	assert.True(t, collections.IsSubCollection(nil, []int{1, 2, 3}))
	assert.False(t, collections.IsSubCollection([]int{1, 2, 3}, nil))
	assert.True(t, collections.IsSubCollection[int](nil, []int{}))
	assert.True(t, collections.IsSubCollection[int]([]int{}, nil))
}

func TestIsProperSubCollection(t *testing.T) {
	assert.False(t, collections.IsProperSubCollection([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.True(t, collections.IsProperSubCollection([]int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}))
	assert.False(t, collections.IsProperSubCollection([]int{1, 2, 3}, []int{1, 2}))
	assert.False(t, collections.IsProperSubCollection[int](nil, nil))
	assert.False(t, collections.IsProperSubCollection(nil, []int{1, 2, 3}))
	assert.False(t, collections.IsProperSubCollection([]int{1, 2, 3}, nil))
	assert.False(t, collections.IsProperSubCollection[int](nil, []int{}))
	assert.False(t, collections.IsProperSubCollection[int]([]int{}, nil))
}

func TestLastIndexOf(t *testing.T) {
	assert.Equal(t, 1, collections.LastIndexOf([]int{1, 3, 2}, 3))
	assert.Equal(t, -1, collections.LastIndexOf([]int{1, 3, 2, 3}, 5))
	assert.Equal(t, 3, collections.LastIndexOf([]int{1, 3, 2, 3}, 3))
}

func TestLastIndexOfWithPredicate(t *testing.T) {
	assert.Panics(t, func() {
		collections.LastIndexOfWithPredicate[int](nil, nil)
	})
	assert.Equal(t, 2, collections.LastIndexOfWithPredicate([]int{1, 2, 3}, func(x int) bool { return x == 3 }))
	assert.Equal(t, -1, collections.LastIndexOfWithPredicate([]int{1, 2, 3}, func(x int) bool { return x == 5 }))
	assert.Equal(t, 3, collections.LastIndexOfWithPredicate([]int{1, 2, 3, 3}, func(x int) bool { return x == 3 }))
}

func TestMatchesAll(t *testing.T) {
	assert.True(t, collections.MatchesAll([]int{1, 1, 1}, func(x int) bool { return x == 1 }))
	assert.False(t, collections.MatchesAll([]int{1, 2, 3}, func(x int) bool { return x == 1 }))
	assert.True(t, collections.MatchesAll(nil, func(x int) bool { return x == 1 }))
	assert.Panics(t, func() {
		collections.MatchesAll[int](nil, nil)
	})
}

func TestMatchesAny(t *testing.T) {
	assert.True(t, collections.MatchesAny([]int{1, 1, 1}, func(x int) bool { return x == 1 }))
	assert.True(t, collections.MatchesAny([]int{1, 2, 3}, func(x int) bool { return x == 1 }))
	assert.False(t, collections.MatchesAny(nil, func(x int) bool { return x == 1 }))
	assert.False(t, collections.MatchesAny([]int{5, 9, 2}, func(x int) bool { return x == 1 }))
	assert.Panics(t, func() {
		collections.MatchesAny[int](nil, nil)
	})
}

func TestPartition(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}}, collections.Partition([]int{1, 2, 3}))
	assert.Equal(t, [][]int{{1, 2, 3}, {5, 9}, nil}, collections.Partition([]int{1, 5, 9, 2, 3},
		func(x int) bool { return x == 1 || x == 2 || x == 3 }, func(x int) bool { return x == 5 || x == 9 || x == 2 }))
	assert.Equal(t, [][]int{{1, 2, 3}, {5, 9}, {10}}, collections.Partition([]int{1, 5, 9, 2, 3, 10},
		func(x int) bool { return x == 1 || x == 2 || x == 3 }, func(x int) bool { return x == 5 || x == 9 || x == 2 }))
	assert.Panics(t, func() {
		collections.Partition([]int{1, 2, 3}, nil)
	})
}

func TestPermutations(t *testing.T) {

}

func TestPredicatedCollection(t *testing.T) {

}

func TestRemoveAll(t *testing.T) {

}

func TestRemoveAllWithEquator(t *testing.T) {

}

func TestRemoveDuplicates(t *testing.T) {

}

func TestRetainAll(t *testing.T) {

}

func TestRetainAllWithEquator(t *testing.T) {

}

func TestReverse(t *testing.T) {

}
