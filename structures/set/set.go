package set

// Set is used to handle use cases for the set data structure.
type Set[T comparable] map[T]struct{}

// New is used to create a new set.
func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// Add adds an element to the set.
// It returns whether the item was added.
func (s Set[T]) Add(val T) bool {
	l := len(s)
	s[val] = struct{}{}
	return l != len(s)
}

// Append multiple elements to the set.
// It returns the number of elements added.
func (s Set[T]) Append(v ...T) int {
	l := len(s)
	for _, val := range v {
		(s)[val] = struct{}{}
	}
	return len(s) - l
}

// Clear removes all elements from the set, leaving the empty set.
func (s Set[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

// Clone returns a clone of the set using the same implementation, duplicating all keys.
func (s Set[T]) Clone() Set[T] {
	clone := make(map[T]struct{}, len(s))
	for v := range s {
		clone[v] = struct{}{}
	}
	return clone
}

// Collection returns the elements of the set as a collection.
func (s Set[T]) Collection() []T {
	r := make([]T, 0, s.Length())
	for v := range s {
		r = append(r, v)
	}
	return r
}

// Contains returns whether the given items are all in the set.
func (s Set[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := s[val]; !ok {
			return false
		}
	}
	return true
}

// ContainsAny returns whether at least one of the given items are in the set.
func (s Set[T]) ContainsAny(v ...T) bool {
	for _, val := range v {
		if _, ok := s[val]; ok {
			return true
		}
	}
	return false
}

// Difference returns the difference between this set
// and other. The returned set will contain
// all elements of this set that are not also
// elements of the second set.
//
// Note that the argument to Difference
// must be of the same type as the receiver
// of the method. Otherwise, Difference will
// panic.
func (s Set[T]) Difference(o Set[T]) Set[T] {
	diff := New[T]()
	for v := range s {
		if _, ok := o[v]; !ok {
			diff[v] = struct{}{}
		}
	}
	return diff
}

// Equal determines if two sets are equal to each
// other. If they have the same cardinality
// and contain the same elements, they are
// considered equal. The order in which
// the elements were added is irrelevant.
//
// Note that the argument to Equal must be
// of the same type as the receiver of the
// method. Otherwise, Equal will panic.
func (s Set[T]) Equal(o Set[T]) bool {
	if s.Length() != o.Length() {
		return false
	}
	for v := range s {
		if !o.Contains(v) {
			return false
		}
	}
	return true
}

// Intersection returns a new set containing only the elements
// that exist only in both sets.
//
// Note that the argument to Intersect
// must be of the same type as the receiver
// of the method. Otherwise, Intersect will
// panic.
func (s Set[T]) Intersection(o Set[T]) Set[T] {
	intersection := make(map[T]struct{})
	// loop over smaller set
	if s.Length() < o.Length() {
		for v := range s {
			if o.Contains(v) {
				intersection[v] = struct{}{}
			}
		}
	} else {
		for v := range o {
			if s.Contains(v) {
				intersection[v] = struct{}{}
			}
		}
	}
	return intersection
}

// Iterator returns an Iterator object that can be used to range over the set.
func (s Set[T]) Iterator() *Iterator[T] {
	i := &Iterator[T]{
		c:    make(chan T),
		stop: make(chan struct{}),
	}
	go func() {
	L:
		for v := range s {
			select {
			case <-i.stop:
				break L
			case i.c <- v:
			}
		}
		close(i.c)
	}()
	return i
}

// Length is used to find the number of elements in the set.
func (s Set[T]) Length() int {
	return len(s)
}

// Remove removes a single element from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// RemoveAll removes multiple elements from the set.
func (s Set[T]) RemoveAll(i ...T) {
	for _, v := range i {
		delete(s, v)
	}
}

// Union returns a new set with all elements in both sets.
//
// Note that the argument to Union must be of the
// same type as the receiver of the method.
// Otherwise, Union will panic.
func (s Set[T]) Union(o Set[T]) Set[T] {
	n := s.Length()
	if o.Length() > n {
		n = o.Length()
	}
	r := make(Set[T], n)
	for v := range s {
		r[v] = struct{}{}
	}
	for v := range o {
		r[v] = struct{}{}
	}
	return r
}
