package utils

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

const (
	MinTick = -887272  // The minimum tick that can be used on any pool.
	MaxTick = -MinTick // The maximum tick that can be used on any pool.
)

var (
	Q32             = big.NewInt(1 << 32)
	MinSqrtRatio    = big.NewInt(4295128739)                                                          // The sqrt ratio corresponding to the minimum tick that could be used on any pool.
	MaxSqrtRatio, _ = new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10) // The sqrt ratio corresponding to the maximum tick that could be used on any pool.
)

var (
	ErrInvalidTick      = errors.New("invalid tick")
	ErrInvalidSqrtRatio = errors.New("invalid sqrt ratio")
)

func mulShift(val *big.Int, mulBy string) *big.Int {
	mulByBig, ok := new(big.Int).SetString(mulBy[2:], 16)
	if !ok {
		panic("invalid mulBy")
	}
	return new(big.Int).Rsh(new(big.Int).Mul(val, mulByBig), 128)
}

/**
 * Returns the sqrt ratio as a Q64.96 for the given tick. The sqrt ratio is computed as sqrt(1.0001)^tick
 * @param tick the tick for which to compute the sqrt ratio
 */
func GetSqrtRatioAtTick(tick int) (*big.Int, error) {
	if tick < MinTick || tick > MaxTick {
		return nil, ErrInvalidTick
	}
	absTick := tick
	if tick < 0 {
		absTick = -tick
	}
	var ratio *big.Int
	if absTick&0x1 != 0 {
		ratio, _ = new(big.Int).SetString("fffcb933bd6fad37aa2d162d1a594001", 16)
	} else {
		ratio, _ = new(big.Int).SetString("100000000000000000000000000000000", 16)
	}
	if (absTick & 0x2) != 0 {
		ratio = mulShift(ratio, "0xfff97272373d413259a46990580e213a")
	}
	if (absTick & 0x4) != 0 {
		ratio = mulShift(ratio, "0xfff2e50f5f656932ef12357cf3c7fdcc")
	}
	if (absTick & 0x8) != 0 {
		ratio = mulShift(ratio, "0xffe5caca7e10e4e61c3624eaa0941cd0")
	}
	if (absTick & 0x10) != 0 {
		ratio = mulShift(ratio, "0xffcb9843d60f6159c9db58835c926644")
	}
	if (absTick & 0x20) != 0 {
		ratio = mulShift(ratio, "0xff973b41fa98c081472e6896dfb254c0")
	}
	if (absTick & 0x40) != 0 {
		ratio = mulShift(ratio, "0xff2ea16466c96a3843ec78b326b52861")
	}
	if (absTick & 0x80) != 0 {
		ratio = mulShift(ratio, "0xfe5dee046a99a2a811c461f1969c3053")
	}
	if (absTick & 0x100) != 0 {
		ratio = mulShift(ratio, "0xfcbe86c7900a88aedcffc83b479aa3a4")
	}
	if (absTick & 0x200) != 0 {
		ratio = mulShift(ratio, "0xf987a7253ac413176f2b074cf7815e54")
	}
	if (absTick & 0x400) != 0 {
		ratio = mulShift(ratio, "0xf3392b0822b70005940c7a398e4b70f3")
	}
	if (absTick & 0x800) != 0 {
		ratio = mulShift(ratio, "0xe7159475a2c29b7443b29c7fa6e889d9")
	}
	if (absTick & 0x1000) != 0 {
		ratio = mulShift(ratio, "0xd097f3bdfd2022b8845ad8f792aa5825")
	}
	if (absTick & 0x2000) != 0 {
		ratio = mulShift(ratio, "0xa9f746462d870fdf8a65dc1f90e061e5")
	}
	if (absTick & 0x4000) != 0 {
		ratio = mulShift(ratio, "0x70d869a156d2a1b890bb3df62baf32f7")
	}
	if (absTick & 0x8000) != 0 {
		ratio = mulShift(ratio, "0x31be135f97d08fd981231505542fcfa6")
	}
	if (absTick & 0x10000) != 0 {
		ratio = mulShift(ratio, "0x9aa508b5b7a84e1c677de54f3e99bc9")
	}
	if (absTick & 0x20000) != 0 {
		ratio = mulShift(ratio, "0x5d6af8dedb81196699c329225ee604")
	}
	if (absTick & 0x40000) != 0 {
		ratio = mulShift(ratio, "0x2216e584f5fa1ea926041bedfe98")
	}
	if (absTick & 0x80000) != 0 {
		ratio = mulShift(ratio, "0x48a170391f7dc42444e8fa2")
	}
	if tick > 0 {
		ratio = new(big.Int).Div(entities.MaxUint256, ratio)
	}

	// back to Q96
	if new(big.Int).Rem(ratio, Q32).Cmp(constants.Zero) > 0 {
		return new(big.Int).Add((new(big.Int).Div(ratio, Q32)), constants.One), nil
	} else {
		return new(big.Int).Div(ratio, Q32), nil
	}
}

/**
 * Returns the tick corresponding to a given sqrt ratio, s.t. #getSqrtRatioAtTick(tick) <= sqrtRatioX96
 * and #getSqrtRatioAtTick(tick + 1) > sqrtRatioX96
 * @param sqrtRatioX96 the sqrt ratio as a Q64.96 for which to compute the tick
 */
func GetTickAtSqrtRatio(sqrtRatioX96 *big.Int) (int, error) {
	if sqrtRatioX96.Cmp(MinSqrtRatio) < 0 || sqrtRatioX96.Cmp(MaxSqrtRatio) >= 0 {
		return 0, ErrInvalidSqrtRatio
	}
	sqrtRatioX128 := new(big.Int).Lsh(sqrtRatioX96, 32)
	msb, err := MostSignificantBit(sqrtRatioX128)
	if err != nil {
		return 0, err
	}
	var r *big.Int
	if big.NewInt(msb).Cmp(big.NewInt(128)) >= 0 {
		r = new(big.Int).Rsh(sqrtRatioX128, uint(msb-127))
	} else {
		r = new(big.Int).Lsh(sqrtRatioX128, uint(127-msb))
	}

	log2 := new(big.Int).Lsh(new(big.Int).Sub(big.NewInt(msb), big.NewInt(128)), 64)

	for i := 0; i < 14; i++ {
		r = new(big.Int).Rsh(new(big.Int).Mul(r, r), 127)
		f := new(big.Int).Rsh(r, 128)
		log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, uint(63-i)))
		r = new(big.Int).Rsh(r, uint(f.Int64()))
	}

	magicSqrt10001, _ := new(big.Int).SetString("255738958999603826347141", 10)
	logSqrt10001 := new(big.Int).Mul(log2, magicSqrt10001)

	magicTickLow, _ := new(big.Int).SetString("3402992956809132418596140100660247210", 10)
	tickLow := new(big.Int).Rsh(new(big.Int).Sub(logSqrt10001, magicTickLow), 128).Int64()
	magicTickHigh, _ := new(big.Int).SetString("291339464771989622907027621153398088495", 10)
	tickHigh := new(big.Int).Rsh(new(big.Int).Add(logSqrt10001, magicTickHigh), 128).Int64()

	if tickLow == tickHigh {
		return int(tickLow), nil
	}

	sqrtRatio, err := GetSqrtRatioAtTick(int(tickHigh))
	if err != nil {
		return 0, err
	}
	if sqrtRatio.Cmp(sqrtRatioX96) <= 0 {
		return int(tickHigh), nil
	} else {
		return int(tickLow), nil
	}
}
