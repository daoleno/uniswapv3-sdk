package utils

import (
	"math/big"

	"github.com/daoleno/uniswapv3-sdk/constants"
)

/**
 * Returns an imprecise maximum amount of liquidity received for a given amount of token 0.
 * This function is available to accommodate LiquidityAmounts#getLiquidityForAmount0 in the v3 periphery,
 * which could be more precise by at least 32 bits by dividing by Q64 instead of Q96 in the intermediate step,
 * and shifting the subtracted ratio left by 32 bits. This imprecise calculation will likely be replaced in a future
 * v3 router contract.
 * @param sqrtRatioAX96 The price at the lower boundary
 * @param sqrtRatioBX96 The price at the upper boundary
 * @param amount0 The token0 amount
 * @returns liquidity for amount0, imprecise
 */
func maxLiquidityForAmount0Imprecise(sqrtRatioAX96, sqrtRatioBX96, amount0 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	intermediate := new(big.Int).Div(new(big.Int).Mul(sqrtRatioAX96, sqrtRatioBX96), constants.Q96)
	return new(big.Int).Div(new(big.Int).Mul(amount0, intermediate), new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
}

/**
 * Returns a precise maximum amount of liquidity received for a given amount of token 0 by dividing by Q64 instead of Q96 in the intermediate step,
 * and shifting the subtracted ratio left by 32 bits.
 * @param sqrtRatioAX96 The price at the lower boundary
 * @param sqrtRatioBX96 The price at the upper boundary
 * @param amount0 The token0 amount
 * @returns liquidity for amount0, precise
 */
func maxLiquidityForAmount0Precise(sqrtRatioAX96, sqrtRatioBX96, amount0 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	numerator := new(big.Int).Mul(new(big.Int).Mul(amount0, sqrtRatioAX96), sqrtRatioBX96)
	denominator := new(big.Int).Mul(constants.Q96, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
	return new(big.Int).Div(numerator, denominator)
}

/**
 * Computes the maximum amount of liquidity received for a given amount of token1
 * @param sqrtRatioAX96 The price at the lower tick boundary
 * @param sqrtRatioBX96 The price at the upper tick boundary
 * @param amount1 The token1 amount
 * @returns liquidity for amount1
 */
func maxLiquidityForAmount1(sqrtRatioAX96, sqrtRatioBX96, amount1 *big.Int) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	return new(big.Int).Div(new(big.Int).Mul(amount1, constants.Q96), new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96))
}

/**
 * Computes the maximum amount of liquidity received for a given amount of token0, token1,
 * and the prices at the tick boundaries.
 * @param sqrtRatioCurrentX96 the current price
 * @param sqrtRatioAX96 price at lower boundary
 * @param sqrtRatioBX96 price at upper boundary
 * @param amount0 token0 amount
 * @param amount1 token1 amount
 * @param useFullPrecision if false, liquidity will be maximized according to what the router can calculate,
 * not what core can theoretically support
 */
func MaxLiquidityForAmounts(sqrtRatioCurrentX96 *big.Int, sqrtRatioAX96, sqrtRatioBX96 *big.Int, amount0, amount1 *big.Int, useFullPrecision bool) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}
	var maxLiquidityForAmount0 func(*big.Int, *big.Int, *big.Int) *big.Int
	if useFullPrecision {
		maxLiquidityForAmount0 = maxLiquidityForAmount0Precise
	} else {
		maxLiquidityForAmount0 = maxLiquidityForAmount0Imprecise
	}
	if sqrtRatioCurrentX96.Cmp(sqrtRatioAX96) <= 0 {
		return maxLiquidityForAmount0(sqrtRatioAX96, sqrtRatioBX96, amount0)
	} else if sqrtRatioCurrentX96.Cmp(sqrtRatioBX96) < 0 {
		liquidity0 := maxLiquidityForAmount0(sqrtRatioCurrentX96, sqrtRatioBX96, amount0)
		liquidity1 := maxLiquidityForAmount1(sqrtRatioAX96, sqrtRatioCurrentX96, amount1)
		if liquidity0.Cmp(liquidity1) < 0 {
			return liquidity0
		}
		return liquidity1

	} else {
		return maxLiquidityForAmount1(sqrtRatioAX96, sqrtRatioBX96, amount1)
	}
}
