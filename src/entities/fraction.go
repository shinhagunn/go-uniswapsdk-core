package entities

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

var (
	oneInt = big.NewInt(1)
	twoInt = big.NewInt(2)
	tenInt = big.NewInt(10)
)

type fraction struct {
	numerator   *big.Int
	denominator *big.Int
}

func NewFraction(numerator *big.Int, denominator *big.Int) *fraction {
	return &fraction{
		numerator:   numerator,
		denominator: denominator,
	}
}

func (f *fraction) GetNumerator() *big.Int {
	return f.numerator
}

func (f *fraction) GetDenominator() *big.Int {
	return f.denominator
}

func (f *fraction) Quotient() *big.Int {
	return new(big.Int).Div(f.numerator, f.denominator)
}

func (f *fraction) Remainder() *fraction {
	return NewFraction(new(big.Int).Rem(f.numerator, f.denominator), f.denominator)
}

func (f *fraction) Invert() *fraction {
	return NewFraction(f.denominator, f.numerator)
}

func (f *fraction) Add(other *fraction) *fraction {
	if f.denominator.Cmp(other.denominator) == 0 {
		return NewFraction(
			new(big.Int).Add(f.numerator, other.numerator),
			f.denominator,
		)
	}

	return NewFraction(
		new(big.Int).Add(
			new(big.Int).Mul(f.numerator, other.denominator),
			new(big.Int).Mul(other.numerator, f.denominator),
		),
		new(big.Int).Mul(f.denominator, other.denominator),
	)
}

func (f *fraction) Subtract(other *fraction) *fraction {
	if f.denominator.Cmp(other.denominator) == 0 {
		return NewFraction(
			new(big.Int).Sub(f.numerator, other.numerator),
			f.denominator,
		)
	}
	return NewFraction(
		new(big.Int).Sub(
			new(big.Int).Mul(f.numerator, other.denominator),
			new(big.Int).Mul(other.numerator, f.denominator),
		),
		new(big.Int).Mul(f.denominator, other.denominator),
	)
}

func (f *fraction) LessThan(other *fraction) bool {
	return new(big.Int).Mul(f.numerator, other.denominator).Cmp(
		new(big.Int).Mul(other.numerator, f.denominator),
	) < 0
}

func (f *fraction) EqualTo(other *fraction) bool {
	return new(big.Int).Mul(f.numerator, other.denominator).Cmp(
		new(big.Int).Mul(other.numerator, f.denominator),
	) == 0
}

func (f *fraction) GreaterThan(other *fraction) bool {
	return new(big.Int).Mul(f.numerator, other.denominator).Cmp(
		new(big.Int).Mul(other.numerator, f.denominator),
	) > 0
}

func (f *fraction) Multiply(other *fraction) *fraction {
	return NewFraction(
		new(big.Int).Mul(f.numerator, other.numerator),
		new(big.Int).Mul(f.denominator, other.denominator),
	)
}

func (f *fraction) Divide(other *fraction) *fraction {
	return NewFraction(
		new(big.Int).Mul(f.numerator, other.denominator),
		new(big.Int).Mul(f.denominator, other.numerator),
	)
}

func (f *fraction) ToSignificant(significantDigits int32) (string, error) {
	if significantDigits <= 0 {
		return "", errors.New("significantDigits must be greater than 0")
	}

	return roundToSignificantFigures(f, significantDigits).String(), nil
}

func (f *fraction) ToFixed(decimalPlaces int32) (string, error) {
	if decimalPlaces < 0 {
		return "", errors.New("decimalPlaces must be greater than 0")
	}

	return decimal.NewFromBigInt(f.numerator, 0).Div(decimal.NewFromBigInt(f.denominator, 0)).StringFixed(decimalPlaces), nil
}

func roundToSignificantFigures(f *fraction, significantDigits int32) decimal.Decimal {
	if significantDigits <= 0 {
		return decimal.Zero
	}

	d := decimal.NewFromBigInt(f.numerator, 0).Div(decimal.NewFromBigInt(f.denominator, 0))
	twoMant := d.Mul(decimal.NewFromFloat(math.Pow10(decimal.DivisionPrecision))).BigInt()
	twoMant.Abs(twoMant)
	twoMant.Mul(twoMant, twoInt)
	upper := big.NewInt(int64(significantDigits))
	upper.Exp(tenInt, upper, nil)
	upper.Mul(upper, twoInt)
	upper.Sub(upper, oneInt)

	m := int64(0)
	for twoMant.Cmp(upper) >= 0 {
		upper.Mul(upper, tenInt)
		m++
	}

	if int64(d.Exponent())+m > int64(math.MaxInt32) {
		panic(fmt.Sprintf("exponent %d overflows an int32", int64(d.Exponent())+m))
	}

	return d.Round(-d.Exponent() - int32(m))
}
