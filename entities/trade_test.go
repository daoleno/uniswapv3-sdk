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
	Ether  = entities.EtherOnChain(1).Wrapped()
	token0 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "t0", "token0")
	token1 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "t1", "token1")
	token2 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000003"), 18, "t2", "token2")
	token3 = entities.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000004"), 18, "t3", "token3")

	pool_0_1 = v2StylePool(
		token0,
		token1,
		entities.FromRawAmount(token0.Currency, big.NewInt(100000)),
		entities.FromRawAmount(token1.Currency, big.NewInt(100000)),
		constants.FeeMedium,
	)
	pool_0_2 = v2StylePool(
		token0,
		token2,
		entities.FromRawAmount(token0.Currency, big.NewInt(100000)),
		entities.FromRawAmount(token2.Currency, big.NewInt(110000)),
		constants.FeeMedium,
	)
	pool_0_3 = v2StylePool(
		token0,
		token3,
		entities.FromRawAmount(token0.Currency, big.NewInt(100000)),
		entities.FromRawAmount(token3.Currency, big.NewInt(90000)),
		constants.FeeMedium,
	)
	pool_1_2 = v2StylePool(
		token1,
		token2,
		entities.FromRawAmount(token1.Currency, big.NewInt(120000)),
		entities.FromRawAmount(token2.Currency, big.NewInt(100000)),
		constants.FeeMedium,
	)
	pool_1_3 = v2StylePool(
		token1,
		token3,
		entities.FromRawAmount(token1.Currency, big.NewInt(120000)),
		entities.FromRawAmount(token3.Currency, big.NewInt(130000)),
		constants.FeeMedium,
	)
	pool_weth_0 = v2StylePool(
		entities.WETH9[1],
		token0,
		entities.FromRawAmount(entities.WETH9[1].Currency, big.NewInt(100000)),
		entities.FromRawAmount(token0.Currency, big.NewInt(100000)),
		constants.FeeMedium,
	)
	pool_weth_1 = v2StylePool(
		entities.WETH9[1],
		token1,
		entities.FromRawAmount(entities.WETH9[1].Currency, big.NewInt(100000)),
		entities.FromRawAmount(token1.Currency, big.NewInt(100000)),
		constants.FeeMedium,
	)
	pool_weth_2 = v2StylePool(
		entities.WETH9[1],
		token2,
		entities.FromRawAmount(entities.WETH9[1].Currency, big.NewInt(100000)),
		entities.FromRawAmount(token2.Currency, big.NewInt(100000)),
		constants.FeeMedium,
	)
)

func v2StylePool(token0, token1 *entities.Token, reserve0, reserve1 *entities.CurrencyAmount, feeAmount constants.FeeAmount) *Pool {
	sqrtRatioX96 := utils.EncodeSqrtRatioX96(reserve1.Quotient(), reserve0.Quotient())
	liquidity := new(big.Int).Sqrt(new(big.Int).Mul(reserve0.Quotient(), reserve1.Quotient()))
	ticks := []Tick{
		{
			Index:          NearestUsableTick(utils.MinTick, constants.TickSpaces[feeAmount]),
			LiquidityNet:   liquidity,
			LiquidityGross: liquidity,
		},
		{
			Index:          NearestUsableTick(utils.MaxTick, constants.TickSpaces[feeAmount]),
			LiquidityNet:   new(big.Int).Mul(liquidity, big.NewInt(-1)),
			LiquidityGross: liquidity,
		},
	}
	s, err := utils.GetTickAtSqrtRatio(sqrtRatioX96)
	if err != nil {
		panic(err)
	}
	p, err := NewTickListDataProvider(ticks, constants.TickSpaces[feeAmount])
	if err != nil {
		panic(err)
	}
	pool, err := NewPool(token0, token1, feeAmount, sqrtRatioX96, liquidity, s, p)
	if err != nil {
		panic(err)
	}
	return pool
}

func TestFromRoute(t *testing.T) {
	// can be constructed with ETHER as input'
	r, _ := NewRoute([]*Pool{pool_weth_0}, Ether, token0)
	trade, _ := FromRoute(r, entities.FromRawAmount(Ether.Currency, big.NewInt(10000)), entities.ExactInput)
	assert.Equal(t, trade.InputAmount().Currency, Ether.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, token0.Currency)

	// can be constructed with ETHER as input for exact output
	r, _ = NewRoute([]*Pool{pool_weth_0}, Ether, token0)
	trade, _ = FromRoute(r, entities.FromRawAmount(token0.Currency, big.NewInt(10000)), entities.ExactOutput)
	assert.Equal(t, trade.InputAmount().Currency, Ether.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, token0.Currency)

	// can be constructed with ETHER as output
	r, _ = NewRoute([]*Pool{pool_weth_0}, token0, Ether)
	trade, err := FromRoute(r, entities.FromRawAmount(Ether.Currency, big.NewInt(10000)), entities.ExactOutput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, trade.InputAmount().Currency, token0.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, Ether.Currency)

	// can be constructed with ETHER as output for exact input
	r, _ = NewRoute([]*Pool{pool_weth_0}, token0, Ether)
	trade, err = FromRoute(r, entities.FromRawAmount(token0.Currency, big.NewInt(10000)), entities.ExactInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, trade.InputAmount().Currency, token0.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, Ether.Currency)
}

func TestFromRoutes(t *testing.T) {
	// can be constructed with ETHER as input with multiple routes
	r, _ := NewRoute([]*Pool{pool_weth_0}, Ether, token0)
	trade, _ := FromRoutes([]*WrappedRoute{{Amount: entities.FromRawAmount(Ether.Currency, big.NewInt(10000)), Route: r}}, entities.ExactInput)
	assert.Equal(t, trade.InputAmount().Currency, Ether.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, token0.Currency)

	// can be constructed with ETHER as input for exact output with multiple routes
	r0, _ := NewRoute([]*Pool{pool_weth_0}, Ether, token0)
	r1, _ := NewRoute([]*Pool{pool_weth_1, pool_0_1}, Ether, token0)
	trade, err := FromRoutes([]*WrappedRoute{
		{Amount: entities.FromRawAmount(token0.Currency, big.NewInt(3000)), Route: r0},
		{Amount: entities.FromRawAmount(token0.Currency, big.NewInt(7000)), Route: r1},
	}, entities.ExactOutput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, trade.InputAmount().Currency, Ether.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, token0.Currency)

	// can be constructed with ETHER as output with multiple routes
	r0, _ = NewRoute([]*Pool{pool_weth_0}, token0, Ether)
	r1, _ = NewRoute([]*Pool{pool_0_1, pool_weth_1}, token0, Ether)
	trade, err = FromRoutes([]*WrappedRoute{
		{Amount: entities.FromRawAmount(Ether.Currency, big.NewInt(4000)), Route: r0},
		{Amount: entities.FromRawAmount(Ether.Currency, big.NewInt(6000)), Route: r1},
	}, entities.ExactOutput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, trade.InputAmount().Currency, token0.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, Ether.Currency)

	// can be constructed with ETHER as output for exact input with multiple routes
	r0, _ = NewRoute([]*Pool{pool_weth_0}, token0, Ether)
	r1, _ = NewRoute([]*Pool{pool_0_1, pool_weth_1}, token0, Ether)
	trade, err = FromRoutes([]*WrappedRoute{
		{Amount: entities.FromRawAmount(token0.Currency, big.NewInt(3000)), Route: r0},
		{Amount: entities.FromRawAmount(token0.Currency, big.NewInt(7000)), Route: r1},
	}, entities.ExactInput)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, trade.InputAmount().Currency, token0.Currency)
	assert.Equal(t, trade.OutputAmount().Currency, Ether.Currency)

	// errors if pools are re-used between routes
	r0, _ = NewRoute([]*Pool{pool_0_1, pool_weth_1}, token0, Ether)
	r1, _ = NewRoute([]*Pool{pool_0_1, pool_1_2, pool_weth_2}, token0, Ether)
	_, err = FromRoutes([]*WrappedRoute{
		{Amount: entities.FromRawAmount(token0.Currency, big.NewInt(4500)), Route: r0},
		{Amount: entities.FromRawAmount(token0.Currency, big.NewInt(5500)), Route: r1},
	}, entities.ExactInput)
	assert.ErrorIs(t, err, ErrDuplicatePools)
}
