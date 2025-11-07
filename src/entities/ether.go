package entities

import (
	"github.com/shinhagunn/go-uniswapsdk-core/src/types"
)

type Ether struct {
	*baseCurrency
}

func NewEther(chainID types.ChainID) *Ether {
	ether := &Ether{
		baseCurrency: &baseCurrency{
			isNative: true,
			isToken:  false,
			chainID:  chainID,
			decimals: 18,
			symbol:   "ETH",
			name:     "Ether",
		},
	}
	ether.baseCurrency.currency = ether

	return ether
}

func (e *Ether) IsEqual(other Currency) bool {
	if other != nil {
		v, isEther := other.(*Ether)
		if isEther {
			return v.isNative && e.chainID == v.chainID
		}
	}

	return false
}

func (e *Ether) Wrapped() *Token {
	return WETH9[e.GetChainID()]
}
