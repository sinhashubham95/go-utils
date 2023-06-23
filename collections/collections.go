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

func getZeroValue[K comparable]() K {
	var v K
	return v
}
