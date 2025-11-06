package entities

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shinhagunn/go-uniswapsdk-core/src/types"
)

var WETH9 = map[types.ChainID]*token{
	1:        NewToken(1, common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), 18, "WETH", "Wrapped Ether"),
	11155111: NewToken(11155111, common.HexToAddress("0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14"), 18, "WETH", "Wrapped Ether"),
	3:        NewToken(3, common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"), 18, "WETH", "Wrapped Ether"),
	4:        NewToken(4, common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"), 18, "WETH", "Wrapped Ether"),
	5:        NewToken(5, common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"), 18, "WETH", "Wrapped Ether"),
	42:       NewToken(42, common.HexToAddress("0xd0A1E359811322d97991E03f863a0C30C2cF029C"), 18, "WETH", "Wrapped Ether"),

	10:       NewToken(10, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	69:       NewToken(69, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	11155420: NewToken(11155420, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),

	42161:  NewToken(42161, common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"), 18, "WETH", "Wrapped Ether"),
	421611: NewToken(421611, common.HexToAddress("0xB47e6A5f8b33b3F17603C83a0535A9dcD7E32681"), 18, "WETH", "Wrapped Ether"),
	421614: NewToken(421614, common.HexToAddress("0x980B62Da83eFf3D4576C647993b0c1D7faf17c73"), 18, "WETH", "Wrapped Ether"),

	8453:  NewToken(8453, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	84532: NewToken(84532, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),

	56:      NewToken(56, common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"), 18, "WBNB", "Wrapped BNB"),
	137:     NewToken(137, common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"), 18, "WMATIC", "Wrapped MATIC"),
	43114:   NewToken(43114, common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"), 18, "WAVAX", "Wrapped AVAX"),
	7777777: NewToken(7777777, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	81457:   NewToken(81457, common.HexToAddress("0x4300000000000000000000000000000000000004"), 18, "WETH", "Wrapped Ether"),
	324:     NewToken(324, common.HexToAddress("0x5AEa5775959fBC2557Cc8789bC1bf90A239D9a91"), 18, "WETH", "Wrapped Ether"),
	480:     NewToken(480, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	1301:    NewToken(1301, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	130:     NewToken(130, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	10143:   NewToken(10143, common.HexToAddress("0x760AfE86e5de5fa0Ee542fc7B7B713e1c5425701"), 18, "WMON", "Wrapped Monad"),
	1868:    NewToken(1868, common.HexToAddress("0x4200000000000000000000000000000000000006"), 18, "WETH", "Wrapped Ether"),
	143:     NewToken(143, common.HexToAddress("0x3bd359C1119dA7Da1D913D1C4D2B7c461115433A"), 18, "WMON", "Wrapped Monad"),
}
