package collections

import "math/bits"

type sortedHint int // hint for sort when choosing the pivot

type pointerInt struct {
	v int
}

func (p *pointerInt) increment() {
	p.v += 1
}

func (p *pointerInt) value() int {
	return p.v
}

const (
	unknownHint sortedHint = iota
	increasingHint
	decreasingHint
)

func getZeroValue[K any]() K {
	var v K
	return v
}

// shiftDownOrdered implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func shiftDownOrdered[E any](data []E, lo, hi, first int, less func(a, b E) bool) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(data[first+child], data[first+child+1]) {
			child++
		}
		if !less(data[first+root], data[first+child]) {
			return
		}
		data[first+root], data[first+child] = data[first+child], data[first+root]
		root = child
	}
}

func nextPowerOfTwo(length int) uint {
	return 1 << bits.Len(uint(length))
}

func factorial(n int) int {
	if n == 0 {
		return n
	}
	f := 1
	for i := 1; i <= n; i += 1 {
		f *= i
	}
	return f
}
