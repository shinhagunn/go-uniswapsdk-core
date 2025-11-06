package utils

import (
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

// Checks a string starts with 0x, is 42 characters long and contains only hex characters after 0x
var startsWith0xLen42HexRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)

// ValidateAndParseAddress validates an address and returns the parsed (checksummed) version of that address.
// It mimics the TS version using ethers' getAddress by leveraging go-ethereum helpers.
func ValidateAndParseAddress(address string) (string, error) {
	if !common.IsHexAddress(address) {
		return "", fmt.Errorf("%s is not a valid address.", address)
	}
	// Hex() returns the EIP-55 checksummed address string (with 0x prefix)
	return common.HexToAddress(address).Hex(), nil
}

// CheckValidAddress checks if an address is valid by checking 0x prefix, length === 42 and hex encoding.
func CheckValidAddress(address string) (string, error) {
	if startsWith0xLen42HexRegex.MatchString(address) {
		return address, nil
	}
	return "", fmt.Errorf("%s is not a valid address.", address)
}
