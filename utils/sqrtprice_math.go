package utils

import (
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

var MaxUint160 = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil), constants.One)

func multiplyIn256(x, y *big.Int) *big.Int {
	product := new(big.Int).Mul(x, y)
	return new(big.Int).And(product, entities.MaxUint256)
}

func addIn256(x, y *big.Int) *big.Int {
	sum := new(big.Int).Add(x, y)
	return new(big.Int).And(sum, entities.MaxUint256)
}

func GetAmount0Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int, roundUp bool) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) >= 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	numerator1 := new(big.Int).Lsh(liquidity, 96)
	numerator2 := new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)

	if roundUp {
		return MulDivRoundingUp(MulDivRoundingUp(numerator1, numerator2, sqrtRatioBX96), constants.One, sqrtRatioAX96)
	}
	return new(big.Int).Div(new(big.Int).Div(new(big.Int).Mul(numerator1, numerator2), sqrtRatioBX96), sqrtRatioAX96)
}

func GetAmount1Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *big.Int, roundUp bool) *big.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) >= 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	if roundUp {
		return MulDivRoundingUp(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96), constants.Q96)
	}
	return new(big.Int).Div(new(big.Int).Mul(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)), constants.Q96)
}
