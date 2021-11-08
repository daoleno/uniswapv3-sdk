package entities

import (
	"math"

	"github.com/daoleno/uniswapv3-sdk/utils"
)

/**
 * Returns the closest tick that is nearest a given tick and usable for the given tick spacing
 * @param tick the target tick
 * @param tickSpacing the spacing of the pool
 */
func NearestUsableTick(tick int, tickSpacing int) int {
	if tickSpacing <= 0 {
		panic("tickSpacing must be greater than 0")
	}
	if !(tick >= utils.MinTick && tick <= utils.MaxTick) {
		panic("tick exceeds bounds")
	}

	rounded := Round(float64(tick)/float64(tickSpacing)) * float64(tickSpacing)
	if rounded < utils.MinTick {
		return int(rounded) + tickSpacing
	}
	if rounded > utils.MaxTick {
		return int(rounded) - tickSpacing
	}
	return int(rounded)
}

// Round like javascript Math.round
// Note that this differs from many languages' round() functions, which often round this case to the next integer away from zero, instead giving a different result in the case of negative numbers with a fractional part of exactly 0.5.
// For example, -1.5 rounds to -2, but -1.5 rounds to -1.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/round#description
func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}
