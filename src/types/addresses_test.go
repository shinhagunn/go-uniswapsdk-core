package types

import "testing"

func TestSwapRouter02Address(t *testing.T) {
	cases := []struct {
		name     string
		chainId  ChainID
		expected string
	}{
		{
			name:     "base",
			chainId:  ChainIDBase,
			expected: "0x2626664c2603336E57B271c5C0b26F421741e481",
		},
		{
			name:     "base goerli",
			chainId:  ChainIDBaseGoerli,
			expected: "0x8357227D4eDc78991Db6FDB9bD6ADE250536dE1d",
		},
		{
			name:     "avalanche",
			chainId:  ChainIDAvalanche,
			expected: "0xbb00FF08d01D300023C629E8fFfFcb65A5a578cE",
		},
		{
			name:     "bnb",
			chainId:  ChainIDBnb,
			expected: "0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2",
		},
		{
			name:     "arbitrum goerli (default fallback)",
			chainId:  ChainIDArbitrumGoerli,
			expected: "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
		},
		{
			name:     "optimism sepolia",
			chainId:  ChainIDOptimismSepolia,
			expected: "0x94cC0AaC535CCDB3C01d6787D6413C739ae12bc4",
		},
		{
			name:     "sepolia",
			chainId:  ChainIDSepolia,
			expected: "0x3bFA4769FB09eefC5a80d6E87c3B9C650f7Ae48E",
		},
		{
			name:     "blast",
			chainId:  ChainIDBlast,
			expected: "0x549FEB8c9bd4c12Ad2AB27022dA12492aC452B66",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			addr := SwapRouter02Address(tc.chainId)
			if addr != tc.expected {
				t.Fatalf("unexpected address for %s (%d): got %s, want %s", tc.name, tc.chainId, addr, tc.expected)
			}
		})
	}
}
