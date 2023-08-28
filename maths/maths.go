package maths

import (
	"github.com/sinhashubham95/go-utils/numbers"
	"math"
)

// mathematical constants
const (
	E       = 2.71828182845904523536028747135266249775724709369995957496696763  // https://oeis.org/A001113
	Pi      = 3.14159265358979323846264338327950288419716939937510582097494459  // https://oeis.org/A000796
	Phi     = 1.61803398874989484820458683436563811772030917980576286213544862  // https://oeis.org/A001622
	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974  // https://oeis.org/A002193
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931  // https://oeis.org/A019774
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779  // https://oeis.org/A002161
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038  // https://oeis.org/A139339
	Ln2     = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
	Log2E   = 1 / Ln2
	Ln10    = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
	Log10E  = 1 / Ln10
)

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[K numbers.SNumber](a K) K {
	return K(math.Abs(float64(a)))
}

// ACos returns the arc-cosine, in radians, of x.
//
// Special case is:
//	ACos(x) = NaN if x < -1 or x > 1
func ACos[K numbers.FloatingNumber](a K) K {
	return K(math.Acos(float64(a)))
}

// ACosH returns the inverse hyperbolic cosine of x.
//
// Special cases are:
//	ACosH(+Inf) = +Inf
//	ACosH(x) = NaN if x < 1
//	ACosH(NaN) = NaN
func ACosH[K numbers.FloatingNumber](a K) K {
	return K(math.Acosh(float64(a)))
}

// ASin returns the arc-sine, in radians, of x.
//
// Special cases are:
//	ASin(±0) = ±0
//	ASin(x) = NaN if x < -1 or x > 1
func ASin[K numbers.FloatingNumber](a K) K {
	return K(math.Asin(float64(a)))
}

// ASinH returns the inverse hyperbolic sine of x.
//
// Special cases are:
//	ASinH(±0) = ±0
//	ASinH(±Inf) = ±Inf
//	ASinH(NaN) = NaN
func ASinH[K numbers.FloatingNumber](a K) K {
	return K(math.Asinh(float64(a)))
}

// ATan returns the arc-tangent, in radians, of x.
//
// Special cases are:
//      ATan(±0) = ±0
//      ATan(±Inf) = ±Pi/2
func ATan[K numbers.FloatingNumber](a K) K {
	return K(math.Atan(float64(a)))
}

// ATanXY returns the arc tangent of y/x, using
// the signs of the two to determine the quadrant
// of the return value.
//
// Special cases are (in order):
//	ATanXY(y, NaN) = NaN
//	ATanXY(NaN, x) = NaN
//	ATanXY(+0, x>=0) = +0
//	ATanXY(-0, x>=0) = -0
//	ATanXY(+0, x<=-0) = +Pi
//	ATanXY(-0, x<=-0) = -Pi
//	ATanXY(y>0, 0) = +Pi/2
//	ATanXY(y<0, 0) = -Pi/2
//	ATanXY(+Inf, +Inf) = +Pi/4
//	ATanXY(-Inf, +Inf) = -Pi/4
//	ATanXY(+Inf, -Inf) = 3Pi/4
//	ATanXY(-Inf, -Inf) = -3Pi/4
//	ATanXY(y, +Inf) = 0
//	ATanXY(y>0, -Inf) = +Pi
//	ATanXY(y<0, -Inf) = -Pi
//	ATanXY(+Inf, x) = +Pi/2
//	ATanXY(-Inf, x) = -Pi/2
func ATanXY[K numbers.FloatingNumber](x, y K) K {
	return K(math.Atan2(float64(y), float64(x)))
}

// ATanH returns the inverse hyperbolic tangent of x.
//
// Special cases are:
//	ATanH(1) = +Inf
//	ATanH(±0) = ±0
//	ATanH(-1) = -Inf
//	ATanH(x) = NaN if x < -1 or x > 1
//	ATanH(NaN) = NaN
func ATanH[K numbers.FloatingNumber](a K) K {
	return K(math.Atanh(float64(a)))
}

// Cbrt returns the cube root of x.
//
// Special cases are:
//	Cbrt(±0) = ±0
//	Cbrt(±Inf) = ±Inf
//	Cbrt(NaN) = NaN
func Cbrt[K numbers.FloatingNumber](a K) K {
	return K(math.Cbrt(float64(a)))
}

// Ceil returns the least integer value greater than or equal to x.
//
// Special cases are:
//	Ceil(±0) = ±0
//	Ceil(±Inf) = ±Inf
//	Ceil(NaN) = NaN
func Ceil[K numbers.FloatingNumber](a K) K {
	return K(math.Ceil(float64(a)))
}

// CopySign returns the magnitude of the first argument with the sign of the second argument.
func CopySign[K numbers.SNumber](magnitude, sign K) K {
	if (magnitude < 0 && sign >= 0) || (magnitude >= 0 && sign < 0) {
		switch any(magnitude).(type) {
		case int8:
			if magnitude == math.MinInt8 {
				panic("overflow")
			}
		case int16:
			if magnitude == math.MinInt16 {
				panic("overflow")
			}
		case int32:
			if magnitude == math.MinInt32 {
				panic("overflow")
			}
		case int:
			if magnitude == math.MinInt {
				panic("overflow")
			}
		case int64:
			if magnitude == math.MinInt64 {
				panic("overflow")
			}
		}
		return -magnitude
	}
	return magnitude
}

// Cos returns the cosine of the radian argument x.
//
// Special cases are:
//	Cos(±Inf) = NaN
//	Cos(NaN) = NaN
func Cos[K numbers.FloatingNumber](a K) K {
	return K(math.Cos(float64(a)))
}

// CosH returns the hyperbolic cosine of x.
//
// Special cases are:
//	CosH(±0) = 1
//	CosH(±Inf) = +Inf
//	CosH(NaN) = NaN
func CosH[K numbers.FloatingNumber](a K) K {
	return K(math.Cosh(float64(a)))
}

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim[K numbers.SNumber](a, b K) K {
	return Max(a-b, 0)
}

// ERF returns the error function of x.
//
// Special cases are:
//	Erf(+Inf) = 1
//	Erf(-Inf) = -1
//	Erf(NaN) = NaN
func ERF[K numbers.FloatingNumber](a K) K {
	return K(math.Erf(float64(a)))
}

// ERFInverse returns the inverse error function of x.
//
// Special cases are:
//	ERFInverse(1) = +Inf
//	ERFInverse(-1) = -Inf
//	ERFInverse(x) = NaN if x < -1 or x > 1
//	ERFInverse(NaN) = NaN
func ERFInverse[K numbers.FloatingNumber](a K) K {
	return K(math.Erfinv(float64(a)))
}

// ERFC returns the complementary error function of x.
//
// Special cases are:
//	ERFC(+Inf) = 0
//	ERFC(-Inf) = 2
//	ERFC(NaN) = NaN
func ERFC[K numbers.FloatingNumber](a K) K {
	return K(math.Erfc(float64(a)))
}

// ERFCInverse returns the inverse of Erfc(x).
//
// Special cases are:
//	ERFCInverse(0) = +Inf
//	ERFCInverse(2) = -Inf
//	ERFCInverse(x) = NaN if x < 0 or x > 2
//	ERFCInverse(NaN) = NaN
func ERFCInverse[K numbers.FloatingNumber](a K) K {
	return K(math.Erfcinv(float64(a)))
}

// Exp returns x**y, the base-x exponential of y.
//
// Special cases are (in order):
//	Exp(x, ±0) = 1 for any x
//	Exp(1, y) = 1 for any y
//	Exp(x, 1) = x for any x
//	Exp(NaN, y) = NaN
//	Exp(x, NaN) = NaN
//	Exp(±0, y) = ±Inf for y an odd integer < 0
//	Exp(±0, -Inf) = +Inf
//	Exp(±0, +Inf) = +0
//	Exp(±0, y) = +Inf for finite y < 0 and not an odd integer
//	Exp(±0, y) = ±0 for y an odd integer > 0
//	Exp(±0, y) = +0 for finite y > 0 and not an odd integer
//	Exp(-1, ±Inf) = 1
//	Exp(x, +Inf) = +Inf for |x| > 1
//	Exp(x, -Inf) = +0 for |x| > 1
//	Exp(x, +Inf) = +0 for |x| < 1
//	Exp(x, -Inf) = +Inf for |x| < 1
//	Exp(+Inf, y) = +Inf for y > 0
//	Exp(+Inf, y) = +0 for y < 0
//	Exp(-Inf, y) = Exp(-0, -y)
//	Exp(x, y) = NaN for finite x < 0 and finite non-integer y
func Exp[K numbers.Number64](a, b K) K {
	return K(math.Pow(float64(a), float64(b)))
}

// ExpE returns e**x, the base-e exponential of x.
//
// Special cases are:
//	ExpE(+Inf) = +Inf
//	ExpE(NaN) = NaN
// Very large values overflow to 0 or +Inf.
// Very small values underflow to 1.
func ExpE[K numbers.Number64](a K) K {
	return K(math.Exp(float64(a)))
}

// Exp2 returns 2**x, the base-2 exponential of x.
//
// Special cases are the same as ExpE.
func Exp2[K numbers.Number64](a K) K {
	return K(math.Exp2(float64(a)))
}

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor[K numbers.FloatingNumber](a K) K {
	return K(math.Floor(float64(a)))
}

// Min is used to return the smallest of the 2 numbers
func Min[K numbers.Number](a ...K) K {
	if len(a) == 0 {
		panic("no elements provided to find the minimum")
	}
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}

// Max is used to return the greatest of the 2 numbers
func Max[K numbers.Number](a ...K) K {
	if len(a) == 0 {
		panic("no elements provided to find the maximum")
	}
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

// NormalizeAngle normalizes an angle in a 2*pi wide interval around a center value.
func NormalizeAngle[K numbers.FloatingNumber](a, center K) K {
	return a - 6.283185307179586*Floor((a+Pi-center)/6.283185307179586)
}

// Reduce ...
func Reduce[K numbers.FloatingNumber](a, period, offset K) K {
	p := Abs(period)
	return a - p*Floor((a-offset)/p) - offset
}
