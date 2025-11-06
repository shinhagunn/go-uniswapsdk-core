package entities

import (
	"errors"
	"math/big"
)

var (
	ErrDifferentCurrencies = errors.New("different currencies")
)

type Price struct {
	*fraction
	baseCurrency  Currency
	quoteCurrency Currency
	scalar        *fraction
}

func NewPrice(baseCurrency, quoteCurrency Currency, denominator, numerator *big.Int) *Price {
	return &Price{
		fraction:      NewFraction(numerator, denominator),
		baseCurrency:  baseCurrency,
		quoteCurrency: quoteCurrency,
		scalar: NewFraction(
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(baseCurrency.GetDecimals())), nil),
			new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(quoteCurrency.GetDecimals())), nil)),
	}
}

func (p *Price) Invert() *Price {
	return NewPrice(p.quoteCurrency, p.baseCurrency, p.fraction.denominator, p.fraction.numerator)
}

func (p *Price) Multiply(other *Price) (*Price, error) {
	if !other.baseCurrency.IsEqual(p.quoteCurrency) {
		return nil, ErrDifferentCurrencies
	}

	fraction := p.fraction.Multiply(other.fraction)
	return NewPrice(p.baseCurrency, other.quoteCurrency, fraction.denominator, fraction.numerator), nil
}

func (p *Price) Quote(currencyAmount *CurrencyAmount) (*CurrencyAmount, error) {
	if !currencyAmount.currency.IsEqual(p.baseCurrency) {
		return nil, ErrDifferentCurrencies
	}

	multiplied := p.fraction.Multiply(currencyAmount.fraction)
	return newCurrencyAmount(p.quoteCurrency, multiplied.numerator, multiplied.denominator), nil
}

func (p *Price) adjustedForDecimals() *fraction {
	return p.fraction.Multiply(p.scalar)
}

func (p *Price) ToSignificant(significantDigits int32) (string, error) {
	return p.adjustedForDecimals().ToSignificant(significantDigits)
}

func (p *Price) ToFixed(decimalPlaces int32) (string, error) {
	return p.adjustedForDecimals().ToFixed(decimalPlaces)
}
