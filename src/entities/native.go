package entities

type nativeCurrency struct {
	*baseCurrency
	wrapped *token
}

// Represents the native currency of the chain on which it resides, e.g.
func NewNativeCurrency(wrapped *token, symbol, name string) Currency {
	native := &nativeCurrency{
		baseCurrency: &baseCurrency{
			isNative: true,
			isToken:  false,
			chainID:  wrapped.GetChainID(),
			decimals: wrapped.GetDecimals(),
			symbol:   symbol,
			name:     name,
		},
		wrapped: wrapped,
	}
	native.baseCurrency.currency = native

	return native
}

func (n *nativeCurrency) IsEqual(other Currency) bool {
	if other != nil {
		v, isNative := other.(*nativeCurrency)
		if isNative {
			return v.isNative && n.chainID == v.chainID
		}
	}

	return false
}

func (n *nativeCurrency) Wrapped() *token {
	return n.wrapped
}
