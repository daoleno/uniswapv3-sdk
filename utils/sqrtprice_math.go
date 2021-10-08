package utils

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

var (
	ErrSqrtPriceLessThanZero = errors.New("sqrt price less than zero")
	ErrLiquidityLessThanZero = errors.New("liquidity less than zero")
	ErrInvariant             = errors.New("invariant violation")
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

func GetNextSqrtPriceFromInput(sqrtPX96, liquidity, amountIn *big.Int, zeroForOne bool) (*big.Int, error) {
	if sqrtPX96.Cmp(constants.Zero) <= 0 {
		return nil, ErrSqrtPriceLessThanZero
	}
	if liquidity.Cmp(constants.Zero) <= 0 {
		return nil, ErrLiquidityLessThanZero
	}
	if zeroForOne {
		return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountIn, true)
	}
	return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountIn, true)
}

func GetNextSqrtPriceFromOutput(sqrtPX96, liquidity, amountOut *big.Int, zeroForOne bool) (*big.Int, error) {
	if sqrtPX96.Cmp(constants.Zero) <= 0 {
		return nil, ErrSqrtPriceLessThanZero
	}
	if liquidity.Cmp(constants.Zero) <= 0 {
		return nil, ErrLiquidityLessThanZero
	}
	if zeroForOne {
		return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountOut, false)
	}
	return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountOut, false)
}

func getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amount *big.Int, add bool) (*big.Int, error) {
	if amount.Cmp(constants.Zero) == 0 {
		return sqrtPX96, nil
	}

	numerator1 := new(big.Int).Lsh(liquidity, 96)
	if add {
		product := multiplyIn256(amount, sqrtPX96)
		if new(big.Int).Div(product, amount).Cmp(sqrtPX96) == 0 {
			denominator := addIn256(numerator1, product)
			if denominator.Cmp(numerator1) >= 0 {
				return MulDivRoundingUp(numerator1, sqrtPX96, denominator), nil
			}
		}
		return MulDivRoundingUp(numerator1, constants.One, new(big.Int).Add(new(big.Int).Div(numerator1, sqrtPX96), amount)), nil
	} else {
		product := multiplyIn256(amount, sqrtPX96)
		if new(big.Int).Div(product, amount).Cmp(sqrtPX96) != 0 {
			return nil, ErrInvariant
		}
		if numerator1.Cmp(product) <= 0 {
			return nil, ErrInvariant
		}
		denominator := new(big.Int).Sub(numerator1, product)
		return MulDivRoundingUp(numerator1, sqrtPX96, denominator), nil
	}
}

func getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amount *big.Int, add bool) (*big.Int, error) {
	if add {
		var quotient *big.Int
		if amount.Cmp(MaxUint160) <= 0 {
			quotient = new(big.Int).Div(new(big.Int).Lsh(amount, 96), liquidity)
		} else {
			quotient = new(big.Int).Div(new(big.Int).Mul(amount, constants.Q96), liquidity)
		}
		return new(big.Int).Add(sqrtPX96, quotient), nil
	}

	quotient := MulDivRoundingUp(amount, constants.Q96, liquidity)
	if sqrtPX96.Cmp(quotient) <= 0 {
		return nil, ErrInvariant
	}
	return new(big.Int).Sub(sqrtPX96, quotient), nil
}
