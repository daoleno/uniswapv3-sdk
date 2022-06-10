package helper

import (
	"math"
	"math/big"
)

func IntWithDecimal(v uint64, decimal int) *big.Int {
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	return new(big.Int).Mul(big.NewInt(int64(v)), pow)
}

func IntDivDecimal(v *big.Int, decimal int) *big.Int {
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimal)), nil)
	return new(big.Int).Div(v, pow)
}

func FloatStringToBigInt(amount string, decimals int) *big.Int {
	fAmount, _ := new(big.Float).SetString(amount)
	fi, _ := new(big.Float).Mul(fAmount, big.NewFloat(math.Pow10(decimals))).Int(nil)
	return fi
}
