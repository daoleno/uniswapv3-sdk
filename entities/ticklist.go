package entities

import (
	"errors"
	"math"
	"math/big"
)

var (
	ErrZeroTickSpacing    = errors.New("tick spacing must be greater than 0")
	ErrInvalidTickSpacing = errors.New("invalid tick spacing")
	ErrZeroNet            = errors.New("tick net delta must be zero")
	ErrSorted             = errors.New("ticks must be sorted")
)

func ValidateList(ticks []Tick, tickSpacing int) error {
	if tickSpacing <= 0 {
		return ErrZeroTickSpacing
	}

	// ensure ticks are spaced appropriately
	for _, t := range ticks {
		if t.Index%tickSpacing != 0 {
			return ErrInvalidTickSpacing
		}
	}

	// ensure tick liquidity deltas sum to 0
	sum := big.NewInt(0)
	for _, tick := range ticks {
		sum.Add(sum, tick.LiquidityNet)
	}
	if sum.Cmp(big.NewInt(0)) != 0 {
		return ErrZeroNet
	}

	if !isTicksSorted(ticks) {
		return ErrSorted
	}

	return nil
}

func IsBelowSmallest(ticks []Tick, tick int) bool {
	if len(ticks) == 0 {
		panic("empty tick list")
	}
	return tick < ticks[0].Index
}

func IsAtOrAboveLargest(ticks []Tick, tick int) bool {
	if len(ticks) == 0 {
		panic("empty tick list")
	}
	return tick >= ticks[len(ticks)-1].Index
}

func GetTick(ticks []Tick, index int) Tick {
	tick := ticks[binarySearch(ticks, index)]
	if tick.Index != index {
		panic("index is not contained in ticks")
	}
	return tick
}

func NextInitializedTick(ticks []Tick, tick int, lte bool) Tick {
	if lte {
		if IsBelowSmallest(ticks, tick) {
			panic("below smallest")
		}
		if IsAtOrAboveLargest(ticks, tick) {
			return ticks[len(ticks)-1]
		}
		index := binarySearch(ticks, tick)
		return ticks[index]
	} else {
		if IsAtOrAboveLargest(ticks, tick) {
			panic("at or above largest")
		}
		if IsBelowSmallest(ticks, tick) {
			return ticks[0]
		}
		index := binarySearch(ticks, tick)
		return ticks[index+1]
	}
}

func NextInitializedTickWithinOneWord(ticks []Tick, tick int, lte bool, tickSpacing int) (int, bool) {
	compressed := math.Floor(float64(tick) / float64(tickSpacing)) // matches rounding in the code

	if lte {
		wordPos := int(compressed) >> 8
		minimum := (wordPos << 8) * tickSpacing
		if IsBelowSmallest(ticks, tick) {
			return minimum, false
		}
		index := NextInitializedTick(ticks, tick, lte).Index
		nextInitializedTick := math.Max(float64(minimum), float64(index))
		return int(nextInitializedTick), int(nextInitializedTick) == index
	} else {
		wordPos := int(compressed+1) >> 8
		maximum := ((wordPos+1)<<8)*tickSpacing - 1
		if IsAtOrAboveLargest(ticks, tick) {
			return maximum, false
		}
		index := NextInitializedTick(ticks, tick, lte).Index
		nextInitializedTick := math.Min(float64(maximum), float64(index))
		return int(nextInitializedTick), int(nextInitializedTick) == index
	}
}

// utils

func isTicksSorted(ticks []Tick) bool {
	for i := 0; i < len(ticks)-1; i++ {
		if ticks[i].Index > ticks[i+1].Index {
			return false
		}
	}
	return true
}

/**
 * Finds the largest tick in the list of ticks that is less than or equal to tick
 * @param ticks list of ticks
 * @param tick tick to find the largest tick that is less than or equal to tick
 * @private
 */
func binarySearch(ticks []Tick, tick int) int {
	if IsBelowSmallest(ticks, tick) {
		panic("tick is below smallest tick")
	}

	// binary search
	start := 0
	end := len(ticks) - 1
	for start <= end {
		mid := (start + end) / 2
		if ticks[mid].Index == tick {
			return mid
		} else if ticks[mid].Index < tick {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// if we get here, we didn't find a tick that is less than or equal to tick
	// so we return the index of the tick that is closest to tick
	if ticks[start].Index < tick {
		return start
	} else {
		return start - 1
	}
}
