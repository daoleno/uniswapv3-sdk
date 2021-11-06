package utils

import "math/big"

/**
 * Returns the sqrt ratio as a Q64.96 corresponding to a given ratio of amount1 and amount0
 * @param amount1 The numerator amount i.e., the amount of token1
 * @param amount0 The denominator amount i.e., the amount of token0
 * @returns The sqrt ratio
 */
func EncodeSqrtRatioX96(amount1 uint64, amount0 uint64) *big.Int {
	numerator := new(big.Int).Lsh(big.NewInt(int64(amount1)), 192)
	denominator := big.NewInt(int64(amount0))
	ratioX192 := new(big.Int).Div(numerator, denominator)
	return new(big.Int).Sqrt(ratioX192)
}
