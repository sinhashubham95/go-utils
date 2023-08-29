package maths_test

import (
	"github.com/sinhashubham95/go-utils/maths"
	"github.com/sinhashubham95/go-utils/numbers"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, math.E, maths.E)
	assert.Equal(t, math.Pi, maths.Pi)
	assert.Equal(t, math.Phi, maths.Phi)
	assert.Equal(t, math.Sqrt2, maths.Sqrt2)
	assert.Equal(t, math.SqrtE, maths.SqrtE)
	assert.Equal(t, math.SqrtPi, maths.SqrtPi)
	assert.Equal(t, math.SqrtPhi, maths.SqrtPhi)
	assert.Equal(t, math.Ln2, maths.Ln2)
	assert.Equal(t, math.Ln10, maths.Ln10)
	assert.Equal(t, math.Log2E, maths.Log2E)
	assert.Equal(t, math.Log10E, maths.Log10E)
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, maths.Abs(-1))
}

func TestACos(t *testing.T) {
	assert.Equal(t, 0.6435011087932843, maths.ACos(0.8))
}

func TestACosH(t *testing.T) {
	assert.Equal(t, 5.07513475044481, maths.ACosH(80.0))
}

func TestASin(t *testing.T) {
	assert.Equal(t, 0.9272952180016123, maths.ASin(0.8))
}

func TestASinH(t *testing.T) {
	assert.Equal(t, 5.0752128754452075, maths.ASinH(80.0))
}

func TestATan(t *testing.T) {
	assert.Equal(t, 0.6747409422235526, maths.ATan(0.8))
}

func TestATanXY(t *testing.T) {
	assert.Equal(t, 0.8960553845713439, maths.ATanXY(1.2, 1.5))
}

func TestATanH(t *testing.T) {
	assert.Equal(t, 0.10033534773107558, maths.ATanH(0.1))
}

func TestCbrt(t *testing.T) {
	assert.Equal(t, 3.0, maths.Cbrt(27.0))
}

func TestCeil(t *testing.T) {
	assert.Equal(t, 2.0, maths.Ceil(1.1))
	assert.Equal(t, 1.0, maths.Ceil(1.0))
}

func TestCopySign(t *testing.T) {
	assert.Equal(t, 1, maths.CopySign(1, 1))
	assert.Equal(t, -1, maths.CopySign(1, -1))
	assert.Equal(t, -1, maths.CopySign(-1, -1))
	assert.Equal(t, 1, maths.CopySign(-1, 1))

	assert.Panics(t, func() {
		maths.CopySign(numbers.MinInt8, 1)
	})
	assert.Panics(t, func() {
		maths.CopySign(numbers.MinInt16, 1)
	})
	assert.Panics(t, func() {
		maths.CopySign(numbers.MinInt32, 1)
	})
	assert.Panics(t, func() {
		maths.CopySign(numbers.MinInt, 1)
	})
	assert.Panics(t, func() {
		maths.CopySign(numbers.MinInt64, 1)
	})
}

func TestCos(t *testing.T) {
	assert.Equal(t, 0.6967067093471655, maths.Cos(0.8))
}

func TestCosH(t *testing.T) {
	assert.Equal(t, 2.770311192196755e+34, maths.CosH(80.0))
}

func TestDim(t *testing.T) {
	assert.Equal(t, 0, maths.Dim(2, 3))
	assert.Equal(t, 1, maths.Dim(3, 2))
	assert.Equal(t, 0, maths.Dim(2, 2))
}

func TestERF(t *testing.T) {
	assert.Equal(t, 0.8427007929497149, maths.ERF(1.0))
}

func TestERFInverse(t *testing.T) {
	assert.Equal(t, 0.08885599049425766, maths.ERFInverse(0.1))
}

func TestERFC(t *testing.T) {
	assert.Equal(t, 0.15729920705028513, maths.ERFC(1.0))
}

func TestERFCInverse(t *testing.T) {
	assert.Equal(t, 1.1630871536766738, maths.ERFCInverse(0.1))
}

func TestExp(t *testing.T) {
	assert.Equal(t, 8.0, maths.Exp(2.0, 3.0))
	assert.Equal(t, int64(8), maths.Exp[int64](2, 3))
}

func TestExpE(t *testing.T) {
	assert.Equal(t, 20.085536923187668, maths.ExpE(3.0))
}

func TestExp2(t *testing.T) {
	assert.Equal(t, 8.0, maths.Exp2(3.0))
	assert.Equal(t, int64(8), maths.Exp2[int64](3))
}

func TestFMA(t *testing.T) {
	assert.Equal(t, 6, maths.FMA(2, 2, 2))
	assert.Equal(t, 6.510000000000001, maths.FMA(2.1, 2.1, 2.1))
}

func TestFloor(t *testing.T) {
	assert.Equal(t, 1.0, maths.Floor(1.1))
	assert.Equal(t, 1.0, maths.Floor(1.0))
}

func TestFractionalExp(t *testing.T) {
	var a any
	var e int
	a, e = maths.FractionalExp(3.0)
	assert.Equal(t, 0.75, a)
	assert.Equal(t, 2, e)
	a, e = maths.FractionalExp(float32(3.0))
	assert.Equal(t, float32(0.75), a)
	assert.Equal(t, 2, e)
}

func TestGamma(t *testing.T) {
	assert.Equal(t, 1.0, maths.Gamma(2.0))
}

func TestHypotenuse(t *testing.T) {
	assert.Equal(t, 5.0, maths.Hypotenuse(3.0, 4.0))
}

func TestILogB(t *testing.T) {
	assert.Equal(t, 4.0, maths.ILogB(27.5))
}

func TestInfinity(t *testing.T) {
	assert.Equal(t, int64(math.Inf(1)), maths.Infinity[int64](1))
	assert.Equal(t, int64(math.Inf(-1)), maths.Infinity[int64](-1))
	assert.Equal(t, math.Inf(1), maths.Infinity[float64](1))
	assert.Equal(t, math.Inf(-1), maths.Infinity[float64](-1))
}

func TestIsInfinity(t *testing.T) {
	assert.True(t, maths.IsInfinity(math.Inf(1), 1))
	assert.False(t, maths.IsInfinity(math.Inf(1), -1))
	assert.False(t, maths.IsInfinity(math.Inf(-1), 1))
	assert.True(t, maths.IsInfinity(math.Inf(-1), -1))
}

func TestIsNaN(t *testing.T) {
	assert.True(t, maths.IsNaN(math.NaN()))
}

func TestJ(t *testing.T) {
	assert.Equal(t, 0.4400505857449335, maths.J(1.0, 1))
}

func TestJ0(t *testing.T) {
	assert.Equal(t, 0.7651976865579666, maths.J0(1.0))
}

func TestJ1(t *testing.T) {
	assert.Equal(t, 0.4400505857449335, maths.J1(1.0))
}

func TestLDExp(t *testing.T) {
	assert.Equal(t, 4.0, maths.LDExp(1.0, 2))
}

func TestLGamma(t *testing.T) {
	a, s := maths.LGamma(2.0)
	assert.Equal(t, 0.0, a)
	assert.Equal(t, 1, s)
}

func TestLog(t *testing.T) {
	assert.Equal(t, 4.605170185988092, maths.Log(100.0))
}

func TestLog10(t *testing.T) {
	assert.Equal(t, 2.0, maths.Log10(100.0))
}

func TestLog2(t *testing.T) {
	assert.Equal(t, 3.0, maths.Log2(8.0))
}

func TestLogB(t *testing.T) {
	assert.Equal(t, 3.0, maths.LogB(8.0))
}

func TestMin(t *testing.T) {
	assert.Panics(t, func() {
		maths.Min[int]()
	})
	assert.Equal(t, 1, maths.Min(4, 1, 3))
}

func TestMax(t *testing.T) {
	assert.Panics(t, func() {
		maths.Max[int]()
	})
	assert.Equal(t, 3, maths.Max(1, 2, 3))
}

func TestMod(t *testing.T) {
	assert.Equal(t, 1, maths.Mod(5, 2))
}

func TestModF(t *testing.T) {
	a, b := maths.ModF(5.3)
	assert.Equal(t, 5.0, a)
	assert.Equal(t, 0.2999999999999998, b)
}

func TestNaN(t *testing.T) {
	assert.True(t, maths.IsNaN(maths.NaN[float64]()))
	assert.True(t, maths.IsNaN(maths.NaN[float32]()))
}

func TestNextAfter(t *testing.T) {
	assert.Equal(t, 1.1000000000000003, maths.NextAfter(1.1, 1.11))
}

func TestNormalizeAngle(t *testing.T) {
	assert.Equal(t, 1.0, maths.NormalizeAngle(1.0, 1.0))
}

func TestReduce(t *testing.T) {
	assert.Equal(t, 0.0, maths.Reduce(7.0, 3.0, 1.0))
}

func TestRound(t *testing.T) {
	assert.Equal(t, 2.0, maths.Round(2.2))
	assert.Equal(t, 3.0, maths.Round(2.5))
	assert.Equal(t, 3.0, maths.Round(2.7))
}

func TestSignBit(t *testing.T) {
	assert.False(t, maths.SignBit(1.1))
	assert.True(t, maths.SignBit(-1.1))
	assert.False(t, maths.SignBit(0.0))
}

func TestSin(t *testing.T) {
	assert.Equal(t, 0.7173560908995227, maths.Sin(0.8))
}

func TestSinH(t *testing.T) {
	assert.Equal(t, 2.770311192196755e+34, maths.SinH(80.0))
}

func TestSqrt(t *testing.T) {
	assert.Equal(t, 2.0, maths.Sqrt[int, float64](4))
}

func TestTan(t *testing.T) {
	assert.Equal(t, 1.0296385570503641, maths.Tan(0.8))
}

func TestTanH(t *testing.T) {
	assert.Equal(t, 1.0, maths.TanH(80.0))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, 1, maths.Truncate[float64, int](1.1))
}

func TestY(t *testing.T) {
	assert.Equal(t, -0.7812128213002887, maths.Y(1.0, 1))
}

func TestY0(t *testing.T) {
	assert.Equal(t, 0.08825696421567697, maths.Y0(1.0))
}

func TestY1(t *testing.T) {
	assert.Equal(t, -0.7812128213002887, maths.Y1(1.0))
}
