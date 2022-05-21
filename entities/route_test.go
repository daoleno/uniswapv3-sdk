package entities

import (
	"math/big"
	"testing"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	rEther  = entities.EtherOnChain(1).Wrapped()
	rtoken0 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "t0", "token0")
	rtoken1 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "t1", "token1")
	rtoken2 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000003"), 18, "t2", "token2")
	rweth   = entities.WETH9[1]

	rpool_0_1, _    = NewPool(rtoken0, rtoken1, constants.FeeMedium, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)
	rpool_0_weth, _ = NewPool(rtoken0, rweth, constants.FeeMedium, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)
	rpool_1_weth, _ = NewPool(rtoken1, rweth, constants.FeeMedium, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)
)

func TestPath(t *testing.T) {
	// constructs a path from the tokens
	route, err := NewRoute([]*Pool{rpool_0_1}, rtoken0, rtoken1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, route.Pools, []*Pool{rpool_0_1})
	assert.Equal(t, route.TokenPath, []*entities.Token{rtoken0, rtoken1})
	assert.Equal(t, route.Input, rtoken0)
	assert.Equal(t, route.Output, rtoken1)
	assert.Equal(t, route.ChainID(), uint(1))

	_, err = NewRoute([]*Pool{rpool_0_1}, rweth, rtoken1)
	assert.ErrorIs(t, err, ErrInputNotInvolved, "should fail if the input is not in the first pool")

	_, err = NewRoute([]*Pool{rpool_0_1}, rtoken0, rweth)
	assert.ErrorIs(t, err, ErrOutputNotInvolved, "should fail if the output is not in the last pool")
}

func TestSameInputOutput(t *testing.T) {
	// can have a token as both input and output
	route, err := NewRoute([]*Pool{rpool_0_weth, rpool_0_1, rpool_1_weth}, rweth, rweth)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, route.Pools, []*Pool{rpool_0_weth, rpool_0_1, rpool_1_weth})
	assert.Equal(t, route.Input, rweth)
	assert.Equal(t, route.Output, rweth)
}

func TestEtherInput(t *testing.T) {
	// supports ether input
	route, err := NewRoute([]*Pool{rpool_0_weth}, rEther, rtoken0)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, route.Pools, []*Pool{rpool_0_weth})
	assert.Equal(t, route.Input, rEther)
	assert.Equal(t, route.Output, rtoken0)
}

func TestEtherOutput(t *testing.T) {
	// supports ether output
	route, err := NewRoute([]*Pool{rpool_0_weth}, rtoken0, rEther)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, route.Pools, []*Pool{rpool_0_weth})
	assert.Equal(t, route.Input, rtoken0)
	assert.Equal(t, route.Output, rEther)
}

func TestMidPrice(t *testing.T) {
	r, _ := utils.GetTickAtSqrtRatio(utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(5)))
	pool_0_1, _ := NewPool(rtoken0, rtoken1, constants.FeeMedium, utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(5)), big.NewInt(0), r, nil)

	r, _ = utils.GetTickAtSqrtRatio(utils.EncodeSqrtRatioX96(big.NewInt(15), big.NewInt(30)))
	pool_1_2, _ := NewPool(rtoken1, rtoken2, constants.FeeMedium, utils.EncodeSqrtRatioX96(big.NewInt(15), big.NewInt(30)), big.NewInt(0), r, nil)

	r, _ = utils.GetTickAtSqrtRatio(utils.EncodeSqrtRatioX96(big.NewInt(3), big.NewInt(1)))
	pool_0_weth, _ := NewPool(rtoken0, rweth, constants.FeeMedium, utils.EncodeSqrtRatioX96(big.NewInt(3), big.NewInt(1)), big.NewInt(0), r, nil)

	r, _ = utils.GetTickAtSqrtRatio(utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(7)))
	pool_1_weth, _ := NewPool(rtoken1, rweth, constants.FeeMedium, utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(7)), big.NewInt(0), r, nil)

	// correct for 0 -> 1
	route, _ := NewRoute([]*Pool{pool_0_1}, rtoken0, rtoken1)
	price, _ := route.MidPrice()
	assert.True(t, price.BaseCurrency.Equal(rtoken0))
	assert.True(t, price.QuoteCurrency.Equal(rtoken1))

	// correct for 1 -> 0
	route, _ = NewRoute([]*Pool{pool_0_1}, rtoken1, rtoken0)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToFixed(4), "5.0000")
	assert.True(t, price.BaseCurrency.Equal(rtoken1))
	assert.True(t, price.QuoteCurrency.Equal(rtoken0))

	// correct for 0 -> 1 -> 2
	route, _ = NewRoute([]*Pool{pool_0_1, pool_1_2}, rtoken0, rtoken2)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToFixed(4), "0.1000")
	assert.True(t, price.BaseCurrency.Equal(rtoken0))
	assert.True(t, price.QuoteCurrency.Equal(rtoken2))

	// correct for 2 -> 1 -> 0
	route, _ = NewRoute([]*Pool{pool_1_2, pool_0_1}, rtoken2, rtoken0)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToFixed(4), "10.0000")
	assert.True(t, price.BaseCurrency.Equal(rtoken2))
	assert.True(t, price.QuoteCurrency.Equal(rtoken0))

	// correct for ether -> 0
	route, _ = NewRoute([]*Pool{pool_0_weth}, rEther, rtoken0)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToFixed(4), "0.3333")
	assert.True(t, price.BaseCurrency.Equal(rEther))
	assert.True(t, price.QuoteCurrency.Equal(rtoken0))

	// correct for 1 -> weth
	route, _ = NewRoute([]*Pool{pool_1_weth}, rtoken1, rweth)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToFixed(4), "0.1429")
	assert.True(t, price.BaseCurrency.Equal(rtoken1))
	assert.True(t, price.QuoteCurrency.Equal(rweth))

	// correct for ether -> 0 -> 1 -> weth
	route, _ = NewRoute([]*Pool{pool_0_weth, pool_0_1, pool_1_weth}, rEther, rweth)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToSignificant(4), "0.009524")
	assert.True(t, price.BaseCurrency.Equal(rEther))
	assert.True(t, price.QuoteCurrency.Equal(rweth))

	// correct for weth -> 0 -> 1 -> ether
	route, _ = NewRoute([]*Pool{pool_0_weth, pool_0_1, pool_1_weth}, rweth, rEther)
	price, _ = route.MidPrice()
	assert.Equal(t, price.ToSignificant(4), "0.009524")
	assert.True(t, price.BaseCurrency.Equal(rweth))
	assert.True(t, price.QuoteCurrency.Equal(rEther))
}
