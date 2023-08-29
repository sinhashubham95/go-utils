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
			if int8(magnitude) == numbers.MinInt8 {
				panic("overflow")
			}
		case int16:
			if int16(magnitude) == numbers.MinInt16 {
				panic("overflow")
			}
		case int32:
			if int32(magnitude) == numbers.MinInt32 {
				panic("overflow")
			}
		case int:
			if int(magnitude) == numbers.MinInt {
				panic("overflow")
			}
		case int64:
			if int64(magnitude) == numbers.MinInt64 {
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

// Dim returns the maximum of a-b or 0.
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
func ExpE(a float64) float64 {
	return math.Exp(a)
}

// Exp2 returns 2**x, the base-2 exponential of x.
//
// Special cases are the same as ExpE.
func Exp2[K numbers.Number64](a K) K {
	return K(math.Exp2(float64(a)))
}

// FMA returns x * y + z, computed with only one rounding.
// (That is, FMA returns the fused multiply-add of x, y, and z.)
func FMA[K numbers.Number](x, y, z K) K {
	return K(math.FMA(float64(x), float64(y), float64(z)))
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

// FractionalExp breaks f into a normalized fraction and an integral power of two.
// It returns fraction and exp satisfying f == frac × 2**exp,
// with the absolute value of fraction in the interval [½, 1).
//
// Special cases are:
//	FractionalExp(±0) = ±0, 0
//	FractionalExp(±Inf) = ±Inf, 0
//	FractionalExp(NaN) = NaN, 0
func FractionalExp[K numbers.FloatingNumber](a K) (fraction K, exp int) {
	fr, e := math.Frexp(float64(a))
	return K(fr), e
}

// Gamma returns the Gamma function of x.
//
// Special cases are:
//	Gamma(+Inf) = +Inf
//	Gamma(+0) = +Inf
//	Gamma(-0) = -Inf
//	Gamma(x) = NaN for integer x < 0
//	Gamma(-Inf) = NaN
//	Gamma(NaN) = NaN
func Gamma[K numbers.FloatingNumber](a K) K {
	return K(math.Gamma(float64(a)))
}

// Hypotenuse returns square-root(p*p + q*q), taking care to avoid unnecessary overflow and underflow.
//
// Special cases are:
//	Hypot(±Inf, q) = +Inf
//	Hypot(p, ±Inf) = +Inf
//	Hypot(NaN, q) = NaN
//	Hypot(p, NaN) = NaN
func Hypotenuse[K numbers.FloatingNumber](x, y K) K {
	return K(math.Hypot(float64(x), float64(y)))
}

// ILogB returns the binary exponent of x as an integer.
//
// Special cases are:
//	ILogB(±Inf) = MaxInt32
//	ILogB(0) = MinInt32
//	ILogB(NaN) = MaxInt32
func ILogB[K numbers.FloatingNumber](a K) K {
	return K(math.Ilogb(float64(a)))
}

// Infinity returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Infinity[K numbers.Number64](sign K) K {
	if sign >= 0 {
		return K(math.Inf(1))
	}
	return K(math.Inf(-1))
}

// IsInfinity reports whether a is an infinity, according to sign.
// If sign > 0, IsInfinity reports whether a is positive infinity.
// If sign < 0, IsInfinity reports whether a is negative infinity.
// If sign == 0, IsInfinity reports whether a is either infinity.
func IsInfinity[K numbers.Number64](a, sign K) bool {
	return a == Infinity(sign)
}

// IsNaN reports whether a is an IEEE 754 ``not-a-number'' value.
func IsNaN[K numbers.Number](a K) bool {
	// IEEE 754 says that only NaNs satisfy f != f.
	// To avoid the floating-point hardware, could use:
	//	x := Float64bits(f);
	//	return uint32(x>>shift)&mask == mask && x != uvinf && x != uvneginf
	return a != a
}

// J returns the order-n Bessel function of the first kind.
//
// Special cases are:
//	J(±Inf) = 0
//	J(0) = 1
//	J(NaN) = NaN
func J[K numbers.FloatingNumber](a K, n int) K {
	return K(math.Jn(n, float64(a)))
}

// J0 returns the order-zero Bessel function of the first kind.
//
// Special cases are:
//	J0(±Inf) = 0
//	J0(0) = 1
//	J0(NaN) = NaN
func J0[K numbers.FloatingNumber](a K) K {
	return K(math.J0(float64(a)))
}

// J1 returns the order-one Bessel function of the first kind.
//
// Special cases are:
//	J1(±Inf) = 0
//	J1(0) = 1
//	J1(NaN) = NaN
func J1[K numbers.FloatingNumber](a K) K {
	return K(math.J1(float64(a)))
}

// LDExp is the inverse of FractionalExp.
// It returns fraction × 2**exp.
//
// Special cases are:
//	LDExp(±0, exp) = ±0
//	LDExp(±Inf, exp) = ±Inf
//	LDExp(NaN, exp) = NaN
func LDExp[K numbers.FloatingNumber](fraction K, exp int) K {
	return K(math.Ldexp(float64(fraction), exp))
}

// LGamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).
//
// Special cases are:
//	LGamma(+Inf) = +Inf
//	LGamma(0) = +Inf
//	LGamma(-integer) = +Inf
//	LGamma(-Inf) = -Inf
//	LGamma(NaN) = NaN
func LGamma[K numbers.FloatingNumber](a K) (l K, sign int) {
	lg, sign := math.Lgamma(float64(a))
	return K(lg), sign
}

// Log returns the natural logarithm of x.
//
// Special cases are:
//	Log(+Inf) = +Inf
//	Log(0) = -Inf
//	Log(x < 0) = NaN
//	Log(NaN) = NaN
func Log[K numbers.FloatingNumber](a K) K {
	return K(math.Log(float64(a)))
}

// Log10 returns the decimal logarithm of x.
// The special cases are the same as for Log.
func Log10[K numbers.FloatingNumber](a K) K {
	return K(math.Log10(float64(a)))
}

// Log2 returns the binary logarithm of x.
// The special cases are the same as for Log.
func Log2[K numbers.FloatingNumber](a K) K {
	return K(math.Log2(float64(a)))
}

// LogB returns the binary exponent of x.
//
// Special cases are:
//	LogB(±Inf) = +Inf
//	LogB(0) = -Inf
//	LogB(NaN) = NaN
func LogB[K numbers.FloatingNumber](a K) K {
	return K(math.Logb(float64(a)))
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

// Mod returns the floating-point remainder of x/y.
// The magnitude of the result is less than y and its
// sign agrees with that of x.
//
// Special cases are:
//	Mod(±Inf, y) = NaN
//	Mod(NaN, y) = NaN
//	Mod(x, 0) = NaN
//	Mod(x, ±Inf) = x
//	Mod(x, NaN) = NaN
func Mod[K numbers.Number](a, b K) K {
	return K(math.Mod(float64(a), float64(b)))
}

// ModF returns integer and fractional floating-point numbers
// that sum to f. Both values have the same sign as f.
//
// Special cases are:
//	ModF(±Inf) = ±Inf, NaN
//	ModF(NaN) = NaN, NaN
func ModF[K numbers.FloatingNumber](a K) (K, K) {
	x, y := math.Modf(float64(a))
	return K(x), K(y)
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN[K numbers.Number64]() K {
	return K(math.NaN())
}

// NextAfter returns the next representable float64 value after a towards b.
func NextAfter[K numbers.SNumber](a, b K) K {
	return K(math.Nextafter(float64(a), float64(b)))
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

// Round returns the nearest integer, rounding half away from zero.
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round[K numbers.FloatingNumber](a K) K {
	return K(math.Round(float64(a)))
}

// SignBit reports whether x is negative or negative zero.
func SignBit[K numbers.SNumber](a K) bool {
	return math.Signbit(float64(a))
}

// Sin returns the sine of the radian argument x.
//
// Special cases are:
//	Sin(±0) = ±0
//	Sin(±Inf) = NaN
//	Sin(NaN) = NaN
func Sin[K numbers.FloatingNumber](a K) K {
	return K(math.Sin(float64(a)))
}

// SinH returns the hyperbolic sine of x.
//
// Special cases are:
//	Sinh(±0) = ±0
//	Sinh(±Inf) = ±Inf
//	Sinh(NaN) = NaN
func SinH[K numbers.FloatingNumber](a K) K {
	return K(math.Sinh(float64(a)))
}

// Sqrt returns the square root of a.
//
// Special cases are:
//	Sqrt(+Inf) = +Inf
//	Sqrt(±0) = ±0
//	Sqrt(x < 0) = NaN
//	Sqrt(NaN) = NaN
func Sqrt[K numbers.Number, V numbers.FloatingNumber](a K) V {
	return V(math.Sqrt(float64(a)))
}

// Tan returns the tangent of the radian argument x.
//
// Special cases are:
//	Tan(±0) = ±0
//	Tan(±Inf) = NaN
//	Tan(NaN) = NaN
func Tan[K numbers.FloatingNumber](a K) K {
	return K(math.Tan(float64(a)))
}

// TanH returns the hyperbolic tangent of x.
//
// Special cases are:
//	TanH(±0) = ±0
//	TanH(±Inf) = ±1
//	TanH(NaN) = NaN
func TanH[K numbers.FloatingNumber](a K) K {
	return K(math.Tanh(float64(a)))
}

// Truncate returns the integer value of x.
//
// Special cases are:
//	Truncate(±0) = ±0
//	Truncate(±Inf) = ±Inf
//	Truncate(NaN) = NaN
func Truncate[K numbers.FloatingNumber, V numbers.IntegerNumber](a K) V {
	return V(math.Trunc(float64(a)))
}

// Y returns the order-n Bessel function of the second kind.
//
// Special cases are:
//	Yn(n, +Inf) = 0
//	Yn(n ≥ 0, 0) = -Inf
//	Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
//	Yn(n, x < 0) = NaN
//	Yn(n, NaN) = NaN
func Y[K numbers.FloatingNumber](a K, n int) K {
	return K(math.Yn(n, float64(a)))
}

// Y0 returns the order-zero Bessel function of the second kind.
//
// Special cases are:
//	Y0(+Inf) = 0
//	Y0(0) = -Inf
//	Y0(x < 0) = NaN
//	Y0(NaN) = NaN
func Y0[K numbers.FloatingNumber](a K) K {
	return K(math.Y0(float64(a)))
}

// Y1 returns the order-one Bessel function of the second kind.
//
// Special cases are:
//	Y1(+Inf) = 0
//	Y1(0) = -Inf
//	Y1(x < 0) = NaN
//	Y1(NaN) = NaN
func Y1[K numbers.FloatingNumber](a K) K {
	return K(math.Y1(float64(a)))
}
