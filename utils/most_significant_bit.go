package utils

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
)

var ErrInvalidInput = errors.New("invalid input")

func MostSignificantBit(x *big.Int) (int64, error) {
	if x.Cmp(constants.Zero) <= 0 {
		return 0, ErrInvalidInput
	}
	if x.Cmp(entities.MaxUint256) > 0 {
		return 0, ErrInvalidInput
	}
	var msb int64
	for _, power := range []int64{128, 64, 32, 16, 8, 4, 2, 1} {
		min := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(power)), nil)
		if x.Cmp(min) >= 0 {
			x = new(big.Int).Rsh(x, uint(power))
			msb += power
		}
	}
	return msb, nil
}
