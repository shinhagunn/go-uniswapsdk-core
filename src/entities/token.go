package entities

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shinhagunn/go-uniswapsdk-core/src/types"
)

type Token struct {
	*baseCurrency
	address common.Address
}

func NewToken(chainID types.ChainID, address common.Address, decimals uint, symbol, name string) *Token {
	if decimals >= 255 {
		panic("decimals must be less than 255")
	}

	token := &Token{
		baseCurrency: &baseCurrency{
			isNative: false,
			isToken:  true,
			chainID:  chainID,
			decimals: decimals,
			symbol:   symbol,
			name:     name,
		},
		address: address,
	}
	token.baseCurrency.currency = token

	return token
}

func (t *Token) GetAddress() common.Address {
	return t.address
}

func (t *Token) IsEqual(other Currency) bool {
	if other != nil {
		v, isToken := other.(*Token)
		if isToken {
			return v.isToken && t.chainID == v.chainID && t.address.Cmp(v.address) == 0
		}
	}

	return false
}

func (t *Token) SortsBefore(other Token) (bool, error) {
	if t.chainID != other.chainID {
		return false, errors.New("token chain ID does not match")
	}

	if t.address.Cmp(other.address) == 0 {
		return false, errors.New("token address is the same")
	}

	return t.address.Cmp(other.address) < 0, nil
}

func (t *Token) Wrapped() *Token {
	return t
}
