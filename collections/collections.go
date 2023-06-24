package collections

// AddAll adds all elements given to the given collection.
func AddAll[K any](c []K, e ...K) []K {
	return append(c, e...)
}

// AddAllIgnoringEmpty adds all elements whose value is not the zero value of their type to the given collection.
func AddAllIgnoringEmpty[K comparable](c []K, e ...K) []K {
	for _, v := range e {
		if v == getZeroValue[K]() {
			continue
		}
		c = append(c, v)
	}
	return c
}

// CardinalityMap returns a Map mapping each unique element in the given Collection to an Integer
// representing the number of occurrences of that element in the Collection.
func CardinalityMap[K comparable](a []K) map[K]int {
	m := make(map[K]int)
	for _, v := range a {
		m[v] += 1
	}
	return m
}

// Collate merges two Collections, a and b, into a single, sorted List such that the natural ordering of the elements is retained.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func Collate[K ordered](a, b []K) []K {
	return CollateWithComparator(a, b, func(x, y K) bool { return x < y })
}

// CollateWithComparator merges two Collections, a and b, into a single,
// sorted List such that the natural ordering of the elements is retained
// according to the comparator provided.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func CollateWithComparator[K any](a, b []K, less func(x, y K) bool) []K {
	la := len(a)
	lb := len(b)

	if la == 0 && lb == 0 {
		return Empty[K]()
	}
	if la == 0 {
		return b
	}
	if lb == 0 {
		return a
	}

	// sort the individual arrays
	sortWithLess(a, less)
	sortWithLess(b, less)

	// form the results and the indices
	r := make([]K, len(a)+len(b))

	i, j, k := 0, 0, 0
	for {
		if i >= la && j >= lb {
			break
		}
		if i >= la {
			r[k] = b[j]
			j += 1
			k += 1
			continue
		}
		if j >= lb {
			r[k] = a[i]
			i += 1
			k += 1
			continue
		}
		if less(a[i], b[j]) {
			r[k] = a[i]
			i += 1
		} else {
			r[k] = b[j]
			j += 1
		}
		k += 1
	}

	return r
}

// CollateRemovingDuplicates merges two Collections, a and b, into a single,
// sorted List such that the natural ordering of the elements is retained.
// In addition, it also removes the duplicate elements from the collection.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func CollateRemovingDuplicates[K ordered](a, b []K) []K {
	return RemoveDuplicates(Collate(a, b))
}

// CollateWithComparatorRemovingDuplicates merges two Collections, a and b, into a single,
// sorted List such that the natural ordering of the elements is retained
// according to the comparator provided.
// In addition, it also removes the duplicate elements from the collection.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func CollateWithComparatorRemovingDuplicates[K comparable](a, b []K, less func(x, y K) bool) []K {
	return RemoveDuplicates(CollateWithComparator(a, b, less))
}

// Collect returns a new Collection containing all elements of the input collection transformed by the given transformer.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func Collect[K any](a []K, transformer func(a K) K) []K {
	r := make([]K, len(a))
	for i, v := range a {
		r[i] = transformer(v)
	}
	return r
}

// Contains is used to check if x is contained in the collection a
func Contains[K comparable](a []K, x K) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

// ContainsAll returns true if all the elements of collection b are also contained in collection a
func ContainsAll[K comparable](a, b []K) bool {
	m := make(map[K]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if !m[v] {
			return false
		}
	}
	return true
}

// ContainsAny returns true if any of the elements of collection b are contained in collection a
func ContainsAny[K comparable](a, b []K) bool {
	m := make(map[K]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			return true
		}
	}
	return false
}

// Count is used to find the count of x in the collection a
func Count[K comparable](a []K, x K) int {
	cnt := 0
	for _, v := range a {
		if v == x {
			cnt += 1
		}
	}
	return cnt
}

// CountMatches Counts the number of elements in the input iterable that match the predicate.
// A null or empty iterable matches no elements.
func CountMatches[K any](a []K, predicate func(x K) bool) int {
	cnt := 0
	for _, v := range a {
		if predicate(v) {
			cnt += 1
		}
	}
	return cnt
}

// Disjunction returns a Collection containing the exclusive disjunction (symmetric difference) of the given collections.
// This means the set of elements which are in either one of the collections but not in both.
func Disjunction[K comparable](a, b []K) []K {
	ma := CardinalityMap(a)
	mb := CardinalityMap(b)
	r := make([]K, 0)
	for v, c := range ma {
		x := c - mb[v]
		if x < 0 {
			x = -x
		}
		for i := 0; i < (c - mb[v]); i += 1 {
			r = append(r, v)
		}
	}
	for v, c := range mb {
		if ma[v] >= 0 {
			continue
		}
		for i := 0; i < c; i += 1 {
			r = append(r, v)
		}
	}
	return r
}

// Empty is used to return an empty collection of the required type.
func Empty[K any]() []K {
	return make([]K, 0)
}

// EmptyBySize is used to return an empty collection of the required type with the specified size.
func EmptyBySize[K any](size int) []K {
	return make([]K, size)
}

// EmptyBySizeAndCapacity is used to return an empty collection of the required type with the specified size and capacity.
func EmptyBySizeAndCapacity[K any](size, capacity int) []K {
	return make([]K, size)
}

// EmptyIfNil is used to return an empty collection if the value is nil otherwise returns the collection itself.
func EmptyIfNil[K any](c []K) []K {
	if c == nil {
		return Empty[K]()
	}
	return c
}

// IsEmpty returns true if the given collection is nil or does not contain any elements.
func IsEmpty[K any](c []K) bool {
	return len(c) == 0
}

// MatchesAll answers true if a predicate is true for every element of an iterable.
// A null or empty iterable returns true.
func MatchesAll[K any](a []K, predicate func(x K) bool) bool {
	for _, v := range a {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// MatchesAny Answers true if a predicate is true for any element of the iterable.
// A null or empty iterable returns false.
func MatchesAny[K any](a []K, predicate func(x K) bool) bool {
	for _, v := range a {
		if predicate(v) {
			return true
		}
	}
	return false
}

// RemoveDuplicates is used to remove duplicates from the collection given.
// This function returns a new collection in itself and the existing collection is not affected.
func RemoveDuplicates[K comparable](a []K) []K {
	l := len(a)
	if l == 0 {
		return Empty[K]()
	}
	r := make([]K, 0)
	last := a[0]
	r = append(r, last)
	for _, v := range a {
		if v == last {
			continue
		}
		last = v
		r = append(r, last)
	}
	return r
}

// Union is used to do the union of the 2 provided collections.
// The cardinality of each element in the returned collection will be equal to the maximum of that element in the given 2 collections.
func Union[K comparable](a, b []K) []K {
	if IsEmpty(a) {
		return b
	}
	if IsEmpty(b) {
		return a
	}
	c := CardinalityMap(a)
	r := append(make([]K, 0, len(a)), a...)
	for _, v := range b {
		if c[v] <= 0 {
			r = append(r, v)
		}
		c[v] -= 1
	}
	return r
}
