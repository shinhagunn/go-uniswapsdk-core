package entities

import (
	"errors"
	"math/big"

	"github.com/shinhagunn/go-uniswapsdk-core/src/types"
	"github.com/shopspring/decimal"
)

type CurrencyAmount struct {
	*fraction
	currency     Currency
	decimalScale *big.Int
}

func newCurrencyAmount(currency Currency, numerator, denominator *big.Int) *CurrencyAmount {
	fraction := NewFraction(numerator, denominator)

	if fraction.Quotient().Cmp(types.MaxUint256) > 0 {
		panic("Currency amount exceeds maximum value(uint256)")
	}

	return &CurrencyAmount{
		fraction:     fraction,
		currency:     currency,
		decimalScale: new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(currency.GetDecimals())), nil),
	}
}

func FromRawAmount(currency Currency, rawAmount *big.Int) *CurrencyAmount {
	return newCurrencyAmount(currency, rawAmount, big.NewInt(1))
}

func FromFractionalAmount(currency Currency, numerator, denominator *big.Int) *CurrencyAmount {
	return newCurrencyAmount(currency, numerator, denominator)
}

func (ca *CurrencyAmount) GetFraction() *fraction {
	return ca.fraction
}

func (ca *CurrencyAmount) Add(other *CurrencyAmount) (*CurrencyAmount, error) {
	if !ca.currency.IsEqual(other.currency) {
		return nil, errors.New("currencies not the same")
	}

	added := ca.fraction.Add(other.fraction)
	return FromFractionalAmount(ca.currency, added.numerator, added.denominator), nil
}

func (ca *CurrencyAmount) Subtract(other *CurrencyAmount) (*CurrencyAmount, error) {
	if !ca.currency.IsEqual(other.currency) {
		return nil, errors.New("currencies not the same")
	}

	subtracted := ca.fraction.Subtract(other.fraction)
	return FromFractionalAmount(ca.currency, subtracted.numerator, subtracted.denominator), nil
}

func (ca *CurrencyAmount) Multiply(other *fraction) *CurrencyAmount {
	multiplied := ca.fraction.Multiply(other)
	return FromFractionalAmount(ca.currency, multiplied.numerator, multiplied.denominator)
}

func (ca *CurrencyAmount) Divide(other *fraction) *CurrencyAmount {
	divided := ca.fraction.Divide(other)
	return FromFractionalAmount(ca.currency, divided.numerator, divided.denominator)
}

func (ca *CurrencyAmount) ToSignificant(significantDigits int32) (string, error) {
	return ca.fraction.Divide(NewFraction(ca.decimalScale, big.NewInt(1))).ToSignificant(significantDigits)
}

func (ca *CurrencyAmount) ToFixed(decimalPlaces int32) (string, error) {
	if uint(decimalPlaces) > ca.currency.GetDecimals() {
		panic("Decimal places exceeds currency decimals")
	}

	return ca.fraction.Divide(NewFraction(ca.decimalScale, big.NewInt(1))).ToFixed(decimalPlaces)
}

func (ca *CurrencyAmount) ToExact() string {
	return decimal.NewFromBigInt(ca.Quotient(), 0).Div(decimal.NewFromBigInt(ca.decimalScale, 0)).String()
}

func (ca *CurrencyAmount) Wrapped() *CurrencyAmount {
	if ca.currency.IsToken() {
		return ca
	}

	return newCurrencyAmount(ca.currency.Wrapped(), ca.numerator, ca.denominator)
}
