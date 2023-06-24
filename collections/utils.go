package collections

import "math/bits"

func getZeroValue[K comparable]() K {
	var v K
	return v
}

// shiftDownOrdered implements the heap property on data[lo:hi].
// first is an offset into the array where the root of the heap lies.
func shiftDownOrdered[E ordered](data []E, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && (data[first+child] < data[first+child+1]) {
			child++
		}
		if !(data[first+root] < data[first+child]) {
			return
		}
		data[first+root], data[first+child] = data[first+child], data[first+root]
		root = child
	}
}

func nextPowerOfTwo(length int) uint {
	return 1 << bits.Len(uint(length))
}
