package utils

import (
	"math/big"
	"testing"

	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/stretchr/testify/assert"
)

func TestGetSqrtRatioAtTick(t *testing.T) {
	_, err := GetSqrtRatioAtTick(MinTick - 1)
	assert.ErrorIs(t, err, ErrInvalidTick, "tick tool small")

	_, err = GetSqrtRatioAtTick(MaxTick + 1)
	assert.ErrorIs(t, err, ErrInvalidTick, "tick tool large")

	rmax, _ := GetSqrtRatioAtTick(MinTick)
	assert.Equal(t, rmax, MinSqrtRatio, "returns the correct value for min tick")

	r0, _ := GetSqrtRatioAtTick(0)
	assert.Equal(t, r0, new(big.Int).Lsh(constants.One, 96), "returns the correct value for tick 0")

	rmin, _ := GetSqrtRatioAtTick(MaxTick)
	assert.Equal(t, rmin, MaxSqrtRatio, "returns the correct value for max tick")
}

func TestGetTickAtSqrtRatio(t *testing.T) {
	tmin, _ := GetTickAtSqrtRatio(MinSqrtRatio)
	assert.Equal(t, tmin, MinTick, "returns the correct value for sqrt ratio at min tick")

	tmax, _ := GetTickAtSqrtRatio(new(big.Int).Sub(MaxSqrtRatio, constants.One))
	assert.Equal(t, tmax, MaxTick-1, "returns the correct value for sqrt ratio at max tick")
}
