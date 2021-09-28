package utils

import (
	"math/big"
)

var MaxFee = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)

// type SwapMath struct{}

// func ComputeSwapStep(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, amountRemaining *big.Int, feePips constants.FeeAmount) (sqrtRatioNextX96, amountIn, amountOut, feeAmount *big.Int) {
// 	zeroForOne := sqrtRatioCurrentX96.Cmp(sqrtRatioTargetX96) >= 0
// 	exactIn := amountRemaining.Cmp(constants.Zero) >= 0

// 	if exactIn {
// 		amountRemainingLessFee := new(big.Int).Div(new(big.Int).Mul(amountRemaining, new(big.Int).Sub(MaxFee, big.NewInt(int64(feePips)))), MaxFee)
// 		if zeroForOne {
// 			GetAmount
// 		}
// 	}

// }
