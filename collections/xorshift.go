package collections

// xorShift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
type xorShift uint64

// Next is used to get the next element of xor shift.
func (r *xorShift) Next() uint64 {
	*r ^= *r << 13
	*r ^= *r >> 17
	*r ^= *r << 5
	return uint64(*r)
}
