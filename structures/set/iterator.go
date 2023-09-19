package set

// Iterator defines an iterator over a Set, its channel can be used to range over the Set's elements.
type Iterator[T comparable] struct {
	c    <-chan T
	stop chan struct{}
}

// Close closes the Iterator, no further elements will be received on C, C will be closed.
func (i *Iterator[T]) Close() {
	close(i.stop)
	// Exhaust any remaining elements.
	for range i.c {
	}
}

// Elements returns a channel to be used to iterate over the elements of the set.
func (i *Iterator[T]) Elements() <-chan T {
	return i.c
}
