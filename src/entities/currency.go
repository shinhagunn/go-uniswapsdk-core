package entities

import (
	"fmt"

	"github.com/shinhagunn/go-uniswapsdk-core/src/types"
)

// Currency is any fungible financial instrument, including Ether, all ERC20 tokens, and other chain-native currencies
type Currency interface {
	IsNative() bool
	IsToken() bool
	GetChainID() types.ChainID
	GetDecimals() uint
	GetSymbol() string
	GetName() string
	IsEqual(other Currency) bool
	Wrapped() *Token
}

// BaseCurrency is an abstract struct, do not use it directly
type baseCurrency struct {
	currency Currency
	isNative bool
	isToken  bool
	chainID  types.ChainID
	decimals uint
	symbol   string
	name     string
}

func NewBaseCurrency(chainID types.ChainID, decimals uint, symbol, name string) (*baseCurrency, error) {
	if !types.IsSupportedChain(chainID) {
		return nil, fmt.Errorf("unsupported chain: %d", chainID)
	}

	if decimals < 0 || decimals > 254 {
		return nil, fmt.Errorf("invalid decimals: %d (must be >= 0 and <= 254)", decimals)
	}

	return &baseCurrency{
		chainID:  chainID,
		decimals: decimals,
		symbol:   symbol,
		name:     name,
	}, nil
}

func (c *baseCurrency) IsNative() bool {
	return c.isNative
}

func (c *baseCurrency) IsToken() bool {
	return c.isToken
}

func (c *baseCurrency) GetChainID() types.ChainID {
	return c.chainID
}

func (c *baseCurrency) GetDecimals() uint {
	return c.decimals
}

func (c *baseCurrency) GetSymbol() string {
	return c.symbol
}

func (c *baseCurrency) GetName() string {
	return c.name
}

func (c *baseCurrency) IsEqual(other Currency) bool {
	panic("Equal method has to be overridden")
}

func (c *baseCurrency) Wrapped() *Token {
	panic("Wrapped method has to be overridden")
}
