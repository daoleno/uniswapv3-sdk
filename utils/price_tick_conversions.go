package utils

import (
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

/**
 * Returns a price object corresponding to the input tick and the base/quote token
 * Inputs must be tokens because the address order is used to interpret the price represented by the tick
 * @param baseToken the base token of the price
 * @param quoteToken the quote token of the price
 * @param tick the tick for which to return the price
 */
func TickToPrice(baseToken *entities.Token, quoteToken *entities.Token, tick int) (*entities.Price, error) {
	sqrtRatioX96, err := GetSqrtRatioAtTick(tick)
	if err != nil {
		return nil, err
	}
	ratioX192 := new(big.Int).Mul(sqrtRatioX96, sqrtRatioX96)

	sorted, err := baseToken.SortsBefore(quoteToken)
	if err != nil {
		return nil, err
	}
	if sorted {
		return entities.NewPrice(baseToken, quoteToken, constants.Q192, ratioX192), nil
	}
	return entities.NewPrice(baseToken, quoteToken, ratioX192, constants.Q192), nil
}

/**
 * Returns the first tick for which the given price is greater than or equal to the tick price
 * @param price for which to return the closest tick that represents a price less than or equal to the input price,
 * i.e. the price of the returned tick is less than or equal to the input price
 */
func PriceToClosestTick(price *entities.Price, baseToken, quoteToken *entities.Token) (int, error) {
	sorted, err := baseToken.SortsBefore(quoteToken)
	if err != nil {
		return 0, err
	}
	var sqrtRatioX96 *big.Int
	if sorted {
		sqrtRatioX96 = EncodeSqrtRatioX96(price.Numerator, price.Denominator)
	} else {
		sqrtRatioX96 = EncodeSqrtRatioX96(price.Denominator, price.Numerator)
	}
	tick, err := GetTickAtSqrtRatio(sqrtRatioX96)
	if err != nil {
		return 0, err
	}
	nextTickPrice, err := TickToPrice(baseToken, quoteToken, tick+1)
	if err != nil {
		return 0, err
	}
	if sorted {
		if !price.LessThan(nextTickPrice.Fraction) {
			tick++
		}
	} else {
		if !price.GreaterThan(nextTickPrice.Fraction) {
			tick++
		}
	}
	return tick, nil
}
