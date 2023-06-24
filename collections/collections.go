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
	c := make(map[K]int)
	for _, v := range a {
		c[v] += 1
	}
	r := append(make([]K, 0, len(a)), a...)
	for _, v := range b {
		if c[v] <= 0 {
			r = append(r, v)
		}
		c[v] -= 1
	}
	return r
}
