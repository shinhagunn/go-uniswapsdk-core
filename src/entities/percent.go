package entities

import "math/big"

var OneHundred = NewFraction(big.NewInt(100), big.NewInt(1))

type Percent struct {
	*fraction
}

func NewPercent(numerator, denominator *big.Int) *Percent {
	return &Percent{NewFraction(numerator, denominator)}
}

func toPercent(fraction *fraction) *Percent {
	return NewPercent(fraction.numerator, fraction.denominator)
}

func (p *Percent) Add(other *Percent) *Percent {
	return toPercent(p.fraction.Add(other.fraction))
}

func (p *Percent) Subtract(other *Percent) *Percent {
	return toPercent(p.fraction.Subtract(other.fraction))
}

func (p *Percent) Multiply(other *Percent) *Percent {
	return toPercent(p.fraction.Multiply(other.fraction))
}

func (p *Percent) Divide(other *Percent) *Percent {
	return toPercent(p.fraction.Divide(other.fraction))
}

func (p *Percent) ToSignificant(significantDigits int32) (string, error) {
	return p.fraction.Multiply(OneHundred).ToSignificant(significantDigits)
}

func (p *Percent) ToFixed(decimalPlaces int32) (string, error) {
	return p.fraction.Multiply(OneHundred).ToFixed(decimalPlaces)
}
