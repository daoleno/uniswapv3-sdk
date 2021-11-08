package entities

import "math/big"

type Tick struct {
	Index          int
	LiquidityGross *big.Int
	LiquidityNet   *big.Int
}

// Provides information about ticks
type TickDataProvider interface {
	/**
	 * Return information corresponding to a specific tick
	 * @param tick the tick to load
	 */
	GetTick(tick int) Tick

	/**
	 * Return the next tick that is initialized within a single word
	 * @param tick The current tick
	 * @param lte Whether the next tick should be lte the current tick
	 * @param tickSpacing The tick spacing of the pool
	 */
	NextInitializedTickWithinOneWord(tick int, lte bool, tickSpacing int) (int, bool)
}
