package utils

import (
	"math/big"
	"testing"

	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/stretchr/testify/assert"
)

func TestEncodeSqrtRatioX96(t *testing.T) {
	assert.Equal(t, EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(1)), constants.Q96, "1/1")

	r0, _ := new(big.Int).SetString("792281625142643375935439503360", 10)
	assert.Equal(t, EncodeSqrtRatioX96(big.NewInt(100), big.NewInt(1)), r0, 10, "100/1")

	r1, _ := new(big.Int).SetString("7922816251426433759354395033", 10)
	assert.Equal(t, EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(100)), r1, 10, "1/100")

	r2, _ := new(big.Int).SetString("45742400955009932534161870629", 10)
	assert.Equal(t, EncodeSqrtRatioX96(big.NewInt(111), big.NewInt(333)), r2, 10, "111/333")

	r3, _ := new(big.Int).SetString("137227202865029797602485611888", 10)
	assert.Equal(t, EncodeSqrtRatioX96(big.NewInt(333), big.NewInt(111)), r3, 10, "333/111")
}
