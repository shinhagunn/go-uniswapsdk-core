package utils

import (
	"errors"
	"math"
	"math/big"
)

var (
	maxSafeInteger = big.NewInt(9007199254740991) // Number.MAX_SAFE_INTEGER
	zero           = big.NewInt(0)
	one            = big.NewInt(1)
	two            = big.NewInt(2)
)

// Sqrt computes floor(sqrt(value))
// value: the value for which to compute the square root, rounded down
func Sqrt(value *big.Int) (*big.Int, error) {
	if value.Cmp(zero) < 0 {
		return nil, errors.New("NEGATIVE")
	}

	// rely on built in sqrt if possible
	if value.Cmp(maxSafeInteger) < 0 {
		val := value.Int64()
		result := math.Floor(math.Sqrt(float64(val)))
		return big.NewInt(int64(result)), nil
	}

	z := new(big.Int).Set(value)
	x := new(big.Int).Div(value, two)
	x.Add(x, one)

	for x.Cmp(z) < 0 {
		z.Set(x)
		// x = (value/x + x) / 2
		temp := new(big.Int).Div(value, x)
		temp.Add(temp, x)
		x.Div(temp, two)
	}

	return z, nil
}
