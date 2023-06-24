package collections

import "math/bits"

// insertionSortOrdered sorts data[a:b] using insertion sort.
func insertionSortOrdered[E ordered](data []E, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && (data[j] < data[j-1]); j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

// heapSortOrdered sorts data[a:b] using heap sort.
func heapSortOrdered[E ordered](data []E, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with the greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		shiftDownOrdered(data, i, hi, first)
	}

	// Pop elements, the largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data[first], data[first+i] = data[first+i], data[first]
		shiftDownOrdered(data, lo, i, first)
	}
}

// breakPatternsOrdered scatters some elements around in an attempt to break some patterns
// that might cause imbalanced partitions in quicksort.
func breakPatternsOrdered[E ordered](data []E, a, b int) {
	length := b - a
	if length >= 8 {
		random := xorShift(length)
		modulus := nextPowerOfTwo(length)

		for idx := a + (length/4)*2 - 1; idx <= a+(length/4)*2+1; idx++ {
			other := int(uint(random.Next()) & (modulus - 1))
			if other >= length {
				other -= length
			}
			data[idx], data[a+other] = data[a+other], data[idx]
		}
	}
}

// order2Ordered returns x,y where data[x] <= data[y], where x,y=a,b or x,y=b,a.
func order2Ordered[E ordered](data []E, a, b int, swaps *int) (int, int) {
	if data[b] < data[a] {
		*swaps++
		return b, a
	}
	return a, b
}

// medianOrdered returns x where data[x] is the median of data[a],data[b],data[c], where x is a, b, or c.
func medianOrdered[E ordered](data []E, a, b, c int, swaps *int) int {
	a, b = order2Ordered(data, a, b, swaps)
	b, c = order2Ordered(data, b, c, swaps)
	a, b = order2Ordered(data, a, b, swaps)
	return b
}

// medianAdjacentOrdered finds the median of data[a - 1], data[a], data[a + 1] and stores the index into a.
func medianAdjacentOrdered[E ordered](data []E, a int, swaps *int) int {
	return medianOrdered(data, a-1, a, a+1, swaps)
}

// choosePivotOrdered chooses a pivot in data[a:b].
//
// [0,8): chooses a static pivot.
// [8,shortestNinth): uses the simple median-of-three method.
// [shortestNinth,∞): uses the TUKEY NINTH method.
func choosePivotOrdered[E ordered](data []E, a, b int) (pivot int, hint sortedHint) {
	const (
		shortestNinth = 50
		maxSwaps      = 4 * 3
	)

	l := b - a

	var (
		swaps int
		i     = a + l/4*1
		j     = a + l/4*2
		k     = a + l/4*3
	)

	if l >= 8 {
		if l >= shortestNinth {
			// Tukey ninth method, the idea came from Rust's implementation.
			i = medianAdjacentOrdered(data, i, &swaps)
			j = medianAdjacentOrdered(data, j, &swaps)
			k = medianAdjacentOrdered(data, k, &swaps)
		}
		// Find the median among i, j, k and stores it into j.
		j = medianOrdered(data, i, j, k, &swaps)
	}

	switch swaps {
	case 0:
		return j, increasingHint
	case maxSwaps:
		return j, decreasingHint
	default:
		return j, unknownHint
	}
}

// partialInsertionSortOrdered partially sorts a slice, returns true if the slice is sorted at the end.
func partialInsertionSortOrdered[E ordered](data []E, a, b int) bool {
	const (
		maxSteps         = 5  // maximum number of adjacent out-of-order pairs that will get shifted
		shortestShifting = 50 // don't shift any elements on short arrays
	)
	i := a + 1
	for j := 0; j < maxSteps; j++ {
		for i < b && !(data[i] < data[i-1]) {
			i++
		}

		if i == b {
			return true
		}

		if b-a < shortestShifting {
			return false
		}

		data[i], data[i-1] = data[i-1], data[i]

		// Shift the smaller one to the left.
		if i-a >= 2 {
			for j := i - 1; j >= 1; j-- {
				if !(data[j] < data[j-1]) {
					break
				}
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
		// Shift the greater one to the right.
		if b-i >= 2 {
			for j := i + 1; j < b; j++ {
				if !(data[j] < data[j-1]) {
					break
				}
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
	return false
}

// partitionEqualOrdered partitions data[a:b] into elements equal to data[pivot] followed by elements greater than data[pivot].
// It assumed that data[a:b] does not contain elements smaller than the data[pivot].
func partitionEqualOrdered[E ordered](data []E, a, b, pivot int) (newPivot int) {
	data[a], data[pivot] = data[pivot], data[a]
	i, j := a+1, b-1 // i and j are inclusive of the elements remaining to be partitioned

	for {
		for i <= j && !(data[a] < data[i]) {
			i++
		}
		for i <= j && (data[a] < data[j]) {
			j--
		}
		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return i
}

// partitionOrdered does one quicksort partition.
// Let p = data[pivot]
// Moves elements in data[a:b] around, so that data[i]<p and data[j]>=p for i<newPivot and j>newPivot.
// On return, data[new-pivot] = p
func partitionOrdered[E ordered](data []E, a, b, pivot int) (newPivot int, alreadyPartitioned bool) {
	data[a], data[pivot] = data[pivot], data[a]
	i, j := a+1, b-1 // i and j are inclusive of the elements remaining to be partitioned

	for i <= j && (data[i] < data[a]) {
		i++
	}
	for i <= j && !(data[j] < data[a]) {
		j--
	}
	if i > j {
		data[j], data[a] = data[a], data[j]
		return j, true
	}
	data[i], data[j] = data[j], data[i]
	i++
	j--

	for {
		for i <= j && (data[i] < data[a]) {
			i++
		}
		for i <= j && !(data[j] < data[a]) {
			j--
		}
		if i > j {
			break
		}
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	data[j], data[a] = data[a], data[j]
	return j, false
}

func reverseRangeOrdered[E ordered](data []E, a, b int) {
	i := a
	j := b - 1
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

func pdqSortOrdered[E ordered](data []E, a, b, limit int) {
	const maxInsertion = 12

	var (
		wasBalanced    = true // whether the last partitioning was reasonably balanced
		wasPartitioned = true // whether the slice was already partitioned
	)

	for {
		length := b - a

		if length <= maxInsertion {
			insertionSortOrdered(data, a, b)
			return
		}

		// Fall back to heapsort if too many bad choices were made.
		if limit == 0 {
			heapSortOrdered(data, a, b)
			return
		}

		// If the last partitioning was imbalanced, we need to breaking patterns.
		if !wasBalanced {
			breakPatternsOrdered(data, a, b)
			limit--
		}

		pivot, hint := choosePivotOrdered(data, a, b)
		if hint == decreasingHint {
			reverseRangeOrdered(data, a, b)
			// The chosen pivot was pivot-a elements after the start of the array.
			// After reversing, it is pivot-a elements before the end of the array.
			// The idea came from Rust's implementation.
			pivot = (b - 1) - (pivot - a)
			hint = increasingHint
		}

		// The slice is likely already sorted.
		if wasBalanced && wasPartitioned && hint == increasingHint {
			if partialInsertionSortOrdered(data, a, b) {
				return
			}
		}

		// Probably the slice contains many duplicate elements, partition the slice into
		// elements equal to and elements greater than the pivot.
		if a > 0 && !(data[a-1] < data[pivot]) {
			mid := partitionEqualOrdered(data, a, b, pivot)
			a = mid
			continue
		}

		mid, alreadyPartitioned := partitionOrdered(data, a, b, pivot)
		wasPartitioned = alreadyPartitioned

		leftLen, rightLen := mid-a, b-mid
		balanceThreshold := length / 8
		if leftLen < rightLen {
			wasBalanced = leftLen >= balanceThreshold
			pdqSortOrdered(data, a, mid, limit)
			a = mid + 1
		} else {
			wasBalanced = rightLen >= balanceThreshold
			pdqSortOrdered(data, mid+1, b, limit)
			b = mid
		}
	}
}

func sort[K ordered](x []K) {
	n := len(x)
	pdqSortOrdered(x, 0, n, bits.Len(uint(n)))
}
