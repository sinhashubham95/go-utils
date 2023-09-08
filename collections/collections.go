package collections

import (
	"fmt"
)

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

// Cardinality returns the number of occurrences of the given element in the collection.
func Cardinality[K comparable](a []K, x K) int {
	cnt := 0
	for _, v := range a {
		if v == x {
			cnt += 1
		}
	}
	return cnt
}

// CardinalityWithEquator returns the number of occurrences of the given element in the collection according to the equator.
func CardinalityWithEquator[K, V any](a []K, x V, equator func(a K, b V) bool) int {
	if equator == nil {
		panic("equator cannot be nil")
	}
	cnt := 0
	for _, v := range a {
		if equator(v, x) {
			cnt += 1
		}
	}
	return cnt
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

// Chain Combines the provided collections into a single collection.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func Chain[K any](a ...[]K) []K {
	l := 0
	for _, v := range a {
		l += len(v)
	}
	r := make([]K, 0, l)
	for _, v := range a {
		r = append(r, v...)
	}
	return r
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
	return RemoveDuplicatesFromSorted(Collate(a, b))
}

// CollateWithComparatorRemovingDuplicates merges two Collections, a and b, into a single,
// sorted List such that the natural ordering of the elements is retained
// according to the comparator provided.
// In addition, it also removes the duplicate elements from the collection.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func CollateWithComparatorRemovingDuplicates[K comparable](a, b []K, less func(x, y K) bool) []K {
	return RemoveDuplicatesFromSorted(CollateWithComparator(a, b, less))
}

// Collect returns a new Collection containing all elements of the input collection transformed by the given transformer.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func Collect[K any](a []K, transformer func(a K) K) []K {
	if transformer == nil {
		panic("transformer must not be nil")
	}
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

// ContainsWithEquator is used to check if x is contained in the collection according to the equator
func ContainsWithEquator[K, V any](a []K, x V, equator func(a K, b V) bool) bool {
	if equator == nil {
		panic("equator cannot be nil")
	}
	for _, v := range a {
		if equator(v, x) {
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

// Copy is used to copy the elements of the collection into a new collection maintaining the order of the elements.
func Copy[K any](a []K) []K {
	r := make([]K, len(a))
	copy(r, a)
	return r
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
	if predicate == nil {
		panic("predicate must not be nil")
	}
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
	for _, v := range a {
		if mb[v] > 0 {
			ma[v] -= 1
			mb[v] -= 1
		} else {
			r = append(r, v)
		}
	}
	for _, v := range b {
		if mb[v] > 0 {
			mb[v] -= 1
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
	return make([]K, size, capacity)
}

// EmptyIfNil is used to return an empty collection if the value is nil otherwise returns the collection itself.
func EmptyIfNil[K any](c []K) []K {
	if c == nil {
		return Empty[K]()
	}
	return c
}

// ExtractSingleton extracts the lone element of the specified Collection
func ExtractSingleton[K any](a []K) K {
	if len(a) != 1 {
		panic("can extract singleton only when collection size is 1")
	}
	return a[0]
}

// Filter filters the collection by applying a Predicate to each element.
// This function returns a completely new copy of the collection and the existing collection is not modified.
func Filter[K any](a []K, predicate func(x K) bool) []K {
	if predicate == nil {
		panic("predicate must not be nil")
	}
	r := make([]K, 0)
	for _, v := range a {
		if predicate(v) {
			r = append(r, v)
		}
	}
	return r
}

// FilterInverse filters the collection by applying a Predicate to each element.
// This function returns a completely new copy of the collection and the existing collection is not modified.
func FilterInverse[K any](a []K, predicate func(x K) bool) []K {
	if predicate == nil {
		panic("predicate must not be nil")
	}
	r := make([]K, 0)
	for _, v := range a {
		if !predicate(v) {
			r = append(r, v)
		}
	}
	return r
}

// Find finds the first element in the given iterable which matches the given predicate.
// It returns the default value of the type if the value is not found.
// This also returns a helper boolean to denote whether the element was found or not.
func Find[K any](a []K, predicate func(x K) bool) (K, bool) {
	if predicate == nil {
		panic("predicate must not be nil")
	}
	for _, v := range a {
		if predicate(v) {
			return v, true
		}
	}
	return getZeroValue[K](), false
}

// FindInverse finds the first element in the given iterable which does not match the given predicate.
// It returns the default value of the type if the value is not found.
// This also returns a helper boolean to denote whether the element was found or not.
func FindInverse[K any](a []K, predicate func(x K) bool) (K, bool) {
	if predicate == nil {
		panic("predicate must not be nil")
	}
	for _, v := range a {
		if !predicate(v) {
			return v, true
		}
	}
	return getZeroValue[K](), false
}

// FindOrDefault finds the first element in the given iterable which matches the given predicate.
// It returns the default value provided if the value is not found.
// This also returns a helper boolean to denote whether the element was found or not.
func FindOrDefault[K any](a []K, predicate func(x K) bool, defaultValue K) (K, bool) {
	if predicate == nil {
		panic("predicate must not be nil")
	}
	for _, v := range a {
		if predicate(v) {
			return v, true
		}
	}
	return defaultValue, false
}

// FindInverseOrDefault finds the first element in the given iterable which does not match the given predicate.
// It returns the default value provided if the value is not found.
// This also returns a helper boolean to denote whether the element was found or not.
func FindInverseOrDefault[K any](a []K, predicate func(x K) bool, defaultValue K) (K, bool) {
	if predicate == nil {
		panic("predicate must not be nil")
	}
	for _, v := range a {
		if !predicate(v) {
			return v, true
		}
	}
	return defaultValue, false
}

// First is used to get the first element of the collection.
// If the collection is empty or nil, then it returns the default value of the type.
// This also returns a helper boolean to know if the first value was returned from the collection or not.
func First[K any](a []K) (K, bool) {
	if len(a) == 0 {
		return getZeroValue[K](), false
	}
	return a[0], true
}

// FirstOrDefault is used to get the first element of the collection.
// If the collection is empty or nil, then it returns the default value provided.
// This also returns a helper boolean to know if the first value was returned from the collection or not.
func FirstOrDefault[K any](a []K, defaultValue K) (K, bool) {
	if len(a) == 0 {
		return defaultValue, false
	}
	return a[0], true
}

// ForEach Applies the closure to each element of the provided iterable.
func ForEach[K any](a []K, closure func(x K)) {
	if closure == nil {
		panic("closure must not be nil")
	}
	for _, v := range a {
		closure(v)
	}
}

// ForEachButLast Applies the closure to each but the last element of the provided iterable.
func ForEachButLast[K any](a []K, closure func(x K)) {
	if closure == nil {
		panic("closure must not be nil")
	}
	l := len(a)
	for i, v := range a {
		if i+1 == l {
			break
		}
		closure(v)
	}
}

// Get returns the ith element of the collection.
// If the index is out of bound, then it will return the default value of the type.
func Get[K any](a []K, i int) K {
	if i < 0 {
		return getZeroValue[K]()
	}
	l := len(a)
	if i >= l {
		return getZeroValue[K]()
	}
	return a[i]
}

// GetOrDefault returns the ith element of the collection.
// If the index is out of bound, then it will return the default value provided.
func GetOrDefault[K any](a []K, i int, defaultValue K) K {
	if i < 0 {
		return defaultValue
	}
	l := len(a)
	if i >= l {
		return defaultValue
	}
	return a[i]
}

// IndexOf returns the index of the first element in the specified collection that matches the search element.
// If the element is not found it returns -1.
func IndexOf[K comparable](a []K, x K) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

// IndexOfWithPredicate returns the index of the first element in the specified collection that matches the predicate.
// If the element is not found it returns -1.
func IndexOfWithPredicate[K any](a []K, predicate func(x K) bool) int {
	if predicate == nil {
		panic("predicate cannot be nil")
	}
	for i, v := range a {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Intersection returns a Collection containing the intersection of the given collections.
// This means the set of elements which are in both the given collections.
func Intersection[K comparable](a, b []K) []K {
	ma := CardinalityMap(a)
	mb := CardinalityMap(b)
	r := make([]K, 0)
	for _, v := range a {
		if mb[v] > 0 {
			mb[v] -= 1
			r = append(r, v)
		}
		ma[v] -= 1
	}
	return r
}

// IsEmpty returns true if the given collection is nil or does not contain any elements.
func IsEmpty[K any](c []K) bool {
	return len(c) == 0
}

// IsNotEmpty returns true if the given collection contains at least 1 element.
func IsNotEmpty[K any](c []K) bool {
	return len(c) > 0
}

// IsEqual Returns true iff the given Collections contain exactly the same elements with exactly the same cardinalities.
func IsEqual[K comparable](a, b []K) bool {
	la := len(a)
	lb := len(b)
	if la != lb {
		return false
	}
	for i := 0; i < la; i += 1 {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// IsEqualWithEquator Returns true iff the given Collections contain exactly the same elements according to the equator
// with exactly the same cardinalities.
func IsEqualWithEquator[K, V any](a []K, b []V, equator func(x K, y V) bool) bool {
	if equator == nil {
		panic("equator cannot be nil")
	}
	la := len(a)
	lb := len(b)
	if la != lb {
		return false
	}
	for i := 0; i < la; i += 1 {
		if !equator(a[i], b[i]) {
			return false
		}
	}
	return true
}

// IsProperSubCollection returns true if and only if the first collection is a proper sub-collection of the
// second collection B, that is, iff the cardinality of each element in A is less than or equal to the cardinality of
// that element in B, for each element E in A, and there is at least one element F such that the cardinality of
// f in b is strictly greater than the cardinality of F in A.
func IsProperSubCollection[K comparable](a, b []K) bool {
	ma := CardinalityMap(a)
	mb := CardinalityMap(b)
	cg := 0
	for _, v := range a {
		if mb[v] < ma[v] {
			return false
		}
		if mb[v] > ma[v] {
			cg += 1
		}
	}
	return cg > 0
}

// IsSubCollection returns true if and only if the first collection A is a proper sub-collection of the
// second collection b, that is, iff the cardinality of each element in A is less than or equal to the cardinality of
// that element in b.
func IsSubCollection[K comparable](a, b []K) bool {
	ma := CardinalityMap(a)
	mb := CardinalityMap(b)
	for _, v := range a {
		if mb[v] < ma[v] {
			return false
		}
	}
	return true
}

// LastIndexOf returns the index of the last element in the specified collection that matches the search element.
// If the element is not found it returns -1.
func LastIndexOf[K comparable](a []K, x K) int {
	l := len(a)
	for i := l - 1; i >= 0; i -= 1 {
		if a[i] == x {
			return i
		}
	}
	return -1
}

// LastIndexOfWithPredicate returns the index of the last element in the specified collection that matches the predicate.
// If the element is not found it returns -1.
func LastIndexOfWithPredicate[K any](a []K, predicate func(x K) bool) int {
	if predicate == nil {
		panic("predicate cannot be nil")
	}
	l := len(a)
	for i := l - 1; i >= 0; i -= 1 {
		if predicate(a[i]) {
			return i
		}
	}
	return -1
}

// MatchesAll answers true if a predicate is true for every element of an iterable.
// A null or empty iterable returns true.
func MatchesAll[K any](a []K, predicate func(x K) bool) bool {
	if predicate == nil {
		panic("predicate must not be nil")
	}
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
	if predicate == nil {
		panic("predicate must not be nil")
	}
	for _, v := range a {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Partition Partitions all elements from iterable into separate output collections, based on the evaluation of the given predicates.
func Partition[K any](a []K, predicates ...func(x K) bool) [][]K {
	l := len(predicates)
	r := make([][]K, l+1)
	for _, x := range a {
		assigned := false
		for i, p := range predicates {
			if p == nil {
				panic("predicate cannot be nil")
			}
			if p(x) {
				r[i] = append(r[i], x)
				assigned = true
				break
			}
		}
		if !assigned {
			r[l] = append(r[l], x)
		}
	}
	return r
}

// Permutations returns a Collection of all the permutations of the input collection.
func Permutations[K any](a []K) [][]K {
	l := len(a)
	curr := make([]K, 0)
	vis := make(map[int]bool)
	p := make([][]K, factorial(l))
	permutations(a, l, 0, curr, vis, p, &pointerInt{v: 0})
	return p
}

// PredicatedCollection returns a predicated (validating) collection backed by the given collection.
// Predicate should return true for all the elements of the collection.
func PredicatedCollection[K any](a []K, predicate func(x K) bool) {
	for _, v := range a {
		if !predicate(v) {
			panic(fmt.Sprintf("%v rejected by predicate", v))
		}
	}
}

// RemoveAll Removes the elements in remove from collection.
func RemoveAll[K comparable](a, remove []K) []K {
	m := CardinalityMap(remove)
	r := make([]K, 0)
	for _, v := range a {
		if m[v] == 0 {
			r = append(r, v)
		}
	}
	return r
}

// RemoveAllWithEquator Removes the elements in remove from collection according to the equator
func RemoveAllWithEquator[K any](a, remove []K, equator func(x, y K) bool) []K {
	if equator == nil {
		panic("equator cannot be nil")
	}
	r := make([]K, 0)
	for _, v := range a {
		if !ContainsWithEquator(remove, v, equator) {
			r = append(r, v)
		}
	}
	return r
}

// RemoveDuplicatesFromSorted is used to remove duplicates from the sorted collection given.
// This function returns a new collection in itself and the existing collection is not affected.
func RemoveDuplicatesFromSorted[K comparable](a []K) []K {
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

// RetainAll Retains the elements in retain from collection.
func RetainAll[K comparable](a, retain []K) []K {
	m := CardinalityMap(retain)
	r := make([]K, 0)
	for _, v := range a {
		if m[v] > 0 {
			r = append(r, v)
		}
	}
	return r
}

// RetainAllWithEquator Retains the elements in retain from collection according to the equator
func RetainAllWithEquator[K any](a, retain []K, equator func(x, y K) bool) []K {
	if equator == nil {
		panic("equator cannot be nil")
	}
	r := make([]K, 0)
	for _, v := range a {
		if ContainsWithEquator(retain, v, equator) {
			r = append(r, v)
		}
	}
	return r
}

// Reverse is used to reverse the order of the given collection.
// This returns a new collection without affecting the existing collection.
func Reverse[K any](a []K) []K {
	l := len(a)
	r := make([]K, len(a))
	for i := 0; i < l; i += 1 {
		r[i] = a[l-1-i]
	}
	return r
}

// Select Selects all elements from input collection which match the given predicate into an output collection.
// This returns a new collection without affecting the existing collection.
func Select[K any](a []K, predicate func(x K) bool) []K {
	if predicate == nil {
		panic("predicate cannot be nil")
	}
	r := make([]K, 0)
	for _, v := range a {
		if predicate(v) {
			r = append(r, v)
		}
	}
	return r
}

// SelectRejected Selects all elements from input collection which do not match the given predicate into an output collection.
// This returns a new collection without affecting the existing collection.
func SelectRejected[K any](a []K, predicate func(x K) bool) []K {
	if predicate == nil {
		panic("predicate cannot be nil")
	}
	r := make([]K, 0)
	for _, v := range a {
		if !predicate(v) {
			r = append(r, v)
		}
	}
	return r
}

// Sort sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
// This method modifies the existing collection.
func Sort[K ordered](a []K) {
	sortWithLess(a, func(x, y K) bool { return x < y })
}

// SortWithLess sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
// This method modifies the existing collection.
func SortWithLess[K any](a []K, less func(x, y K) bool) {
	sortWithLess(a, less)
}

// Subtract Returns a new Collection containing a - b.
// This function returns a completely new copy of the collection and none of the existing collections are modified.
func Subtract[K comparable](a, b []K) []K {
	m := CardinalityMap(b)
	r := make([]K, 0)
	for _, v := range a {
		c := m[v]
		if c > 0 {
			m[v] -= 1
		} else {
			r = append(r, v)
		}
	}
	return r
}

// Transform transforms the collection by applying a Transformer to each element.
// This returns a new collection without affecting the existing collection.
func Transform[K, L any](a []K, transformer func(x K) L) []L {
	if transformer == nil {
		panic("transformer cannot be nil")
	}
	r := make([]L, len(a))
	for i, v := range a {
		r[i] = transformer(v)
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

func permutations[K any](a []K, l, ind int, curr []K, vis map[int]bool, p [][]K, rind *pointerInt) {
	if ind == l {
		if len(curr) == 0 {
			return
		}
		p[rind.value()] = curr
		rind.increment()
		return
	}
	for i := 0; i < l; i += 1 {
		if vis[i] {
			continue
		}
		vis[i] = true
		curr = append(curr, a[i])
		permutations(a, l, ind+1, curr, vis, p, rind)
		curr = curr[:ind]
		vis[i] = false
	}
}
