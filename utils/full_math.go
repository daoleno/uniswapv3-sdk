package utils

import (
	"math/big"

	"github.com/daoleno/uniswapv3-sdk/constants"
)

func MulDivRoundingUp(a, b, denominator *big.Int) *big.Int {
	product := new(big.Int).Mul(a, b)
	result := new(big.Int).Div(product, denominator)
	if new(big.Int).Rem(product, denominator).Cmp(big.NewInt(0)) != 0 {
		result.Add(result, constants.One)
	}
	return result
}
