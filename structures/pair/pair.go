package pair

// Pair is used to handle use cases where a pair of values have to be handled.
// For example, coordinates x and y.
type Pair[K, V any] struct {
	first  K
	second V
}

// B is used to handle the builder use case.
type B[K, V any] struct {
	p *Pair[K, V]
}

// New is used to create a new object for the struct.
func New[K, V any](f K, s V) *Pair[K, V] {
	return &Pair[K, V]{
		first:  f,
		second: s,
	}
}

// NewFromCollection is used to create a new Pair from the given collection.
// The collection should contain exactly 2 elements, first element of which becomes first and second element becomes second.
// If the collection contains more or less than 2 elements, then it returns NIL Pair.
func NewFromCollection[K any](l []K) *Pair[K, K] {
	if len(l) != 2 {
		return nil
	}
	return New(l[0], l[1])
}

// Builder is used to create a builder for creating Pair type.
func Builder[K, V any]() *B[K, V] {
	return &B[K, V]{
		p: &Pair[K, V]{},
	}
}

// Build is used to get the built instance of the Pair type.
func (b *B[K, V]) Build() *Pair[K, V] {
	if b.p == nil {
		b.p = new(Pair[K, V])
	}
	return b.p
}

// First is used to assign value to the first value in the pair.
func (b *B[K, V]) First(f K) *B[K, V] {
	if b.p == nil {
		b.p = new(Pair[K, V])
	}
	b.p.first = f
	return b
}

// Second is used to assign value to the second value in the pair.
func (b *B[K, V]) Second(s V) *B[K, V] {
	if b.p == nil {
		b.p = new(Pair[K, V])
	}
	b.p.second = s
	return b
}

// GetFirst is used to get the first value.
func (p *Pair[K, _]) GetFirst() K {
	return p.first
}

// GetSecond is used to get the second value.
func (p *Pair[_, V]) GetSecond() V {
	return p.second
}

// SetFirst is used to get the first value.
func (p *Pair[K, _]) SetFirst(f K) {
	p.first = f
}

// SetSecond is used to get the second value.
func (p *Pair[_, V]) SetSecond(s V) {
	p.second = s
}
