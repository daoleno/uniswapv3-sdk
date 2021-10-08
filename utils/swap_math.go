package utils

import (
	"math/big"

	"github.com/daoleno/uniswapv3-sdk/constants"
)

var MaxFee = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

func ComputeSwapStep(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, amountRemaining *big.Int, feePips constants.FeeAmount) (sqrtRatioNextX96, amountIn, amountOut, feeAmount *big.Int, err error) {
	zeroForOne := sqrtRatioCurrentX96.Cmp(sqrtRatioTargetX96) >= 0
	exactIn := amountRemaining.Cmp(constants.Zero) >= 0

	if exactIn {
		amountRemainingLessFee := new(big.Int).Div(new(big.Int).Mul(amountRemaining, new(big.Int).Sub(MaxFee, big.NewInt(int64(feePips)))), MaxFee)
		if zeroForOne {
			amountIn = GetAmount0Delta(sqrtRatioTargetX96, sqrtRatioCurrentX96, liquidity, true)
		} else {
			amountIn = GetAmount1Delta(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, true)
		}
		if amountRemainingLessFee.Cmp(amountIn) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96, err = GetNextSqrtPriceFromInput(sqrtRatioCurrentX96, liquidity, amountRemainingLessFee, zeroForOne)
			if err != nil {
				return
			}
		}
	} else {
		if zeroForOne {
			amountOut = GetAmount1Delta(sqrtRatioTargetX96, sqrtRatioCurrentX96, liquidity, false)
		} else {
			amountOut = GetAmount0Delta(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, false)
		}
		if new(big.Int).Mul(amountRemaining, constants.NegativeOne).Cmp(amountOut) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96, err = GetNextSqrtPriceFromOutput(sqrtRatioCurrentX96, liquidity, new(big.Int).Mul(amountRemaining, constants.NegativeOne), zeroForOne)
			if err != nil {
				return
			}
		}
	}

	max := sqrtRatioTargetX96.Cmp(sqrtRatioNextX96) == 0

	if zeroForOne {
		if !(max && exactIn) {
			amountIn = GetAmount0Delta(sqrtRatioNextX96, sqrtRatioCurrentX96, liquidity, true)
		}
		if !(max && !exactIn) {
			amountOut = GetAmount1Delta(sqrtRatioNextX96, sqrtRatioCurrentX96, liquidity, false)
		}
	} else {
		if !(max && exactIn) {
			amountIn = GetAmount1Delta(sqrtRatioCurrentX96, sqrtRatioNextX96, liquidity, true)
		}
		if !(max && !exactIn) {
			amountOut = GetAmount0Delta(sqrtRatioCurrentX96, sqrtRatioNextX96, liquidity, false)
		}
	}

	if !exactIn && amountOut.Cmp(new(big.Int).Mul(amountRemaining, constants.NegativeOne)) > 0 {
		amountOut = new(big.Int).Mul(amountRemaining, constants.NegativeOne)
	}

	if exactIn && sqrtRatioNextX96.Cmp(sqrtRatioTargetX96) != 0 {
		// we didn't reach the target, so take the remainder of the maximum input as fee
		feeAmount = new(big.Int).Sub(amountRemaining, amountIn)
	} else {
		feeAmount = MulDivRoundingUp(amountIn, big.NewInt(int64(feePips)), new(big.Int).Sub(MaxFee, big.NewInt(int64(feePips))))
	}

	return
}
