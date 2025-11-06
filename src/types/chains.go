package types

type ChainID uint

const (
	ChainIDMainnet         = ChainID(1)
	ChainIDGoerli          = ChainID(5)
	ChainIDSepolia         = ChainID(11155111)
	ChainIDOptimism        = ChainID(10)
	ChainIDOptimismGoerli  = ChainID(420)
	ChainIDOptimismSepolia = ChainID(11155420)
	ChainIDArbitrumOne     = ChainID(42161)
	ChainIDArbitrumGoerli  = ChainID(421613)
	ChainIDArbitrumSepolia = ChainID(421614)
	ChainIDPolygon         = ChainID(137)
	ChainIDPolygonMumbai   = ChainID(80001)
	ChainIDCelo            = ChainID(42220)
	ChainIDCeloAlfajores   = ChainID(44787)
	ChainIDGnosis          = ChainID(100)
	ChainIDMoonbeam        = ChainID(1284)
	ChainIDBnb             = ChainID(56)
	ChainIDAvalanche       = ChainID(43114)
	ChainIDBaseGoerli      = ChainID(84531)
	ChainIDBase            = ChainID(8453)
	ChainIDZora            = ChainID(7777777)
	ChainIDZoraSepolia     = ChainID(999999999)
	ChainIDRootstock       = ChainID(30)
	ChainIDBlast           = ChainID(81457)
)

var SupportedChains = []ChainID{
	ChainIDMainnet,
	ChainIDOptimism,
	ChainIDOptimismGoerli,
	ChainIDOptimismSepolia,
	ChainIDArbitrumOne,
	ChainIDArbitrumGoerli,
	ChainIDArbitrumSepolia,
	ChainIDPolygon,
	ChainIDPolygonMumbai,
	ChainIDGoerli,
	ChainIDSepolia,
	ChainIDCeloAlfajores,
	ChainIDCelo,
	ChainIDBnb,
	ChainIDAvalanche,
	ChainIDBase,
	ChainIDBaseGoerli,
	ChainIDZora,
	ChainIDZoraSepolia,
	ChainIDRootstock,
	ChainIDBlast,
}

// IsSupportedChain returns true if the provided chain ID is in SupportedChains.
func IsSupportedChain(chain ChainID) bool {
	for _, c := range SupportedChains {
		if c == chain {
			return true
		}
	}
	return false
}

type NativeCurrencyName string

const (
	NativeCurrencyNameEther     = NativeCurrencyName("ETH")
	NativeCurrencyNameMatic     = NativeCurrencyName("MATIC")
	NativeCurrencyNameCelo      = NativeCurrencyName("CELO")
	NativeCurrencyNameGnosis    = NativeCurrencyName("XDAI")
	NativeCurrencyNameMoonbeam  = NativeCurrencyName("GLMR")
	NativeCurrencyNameBnb       = NativeCurrencyName("BNB")
	NativeCurrencyNameAvalanche = NativeCurrencyName("AVAX")
	NativeCurrencyNameRootstock = NativeCurrencyName("RBTC")
)
