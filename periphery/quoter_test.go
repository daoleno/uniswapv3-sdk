package periphery

import (
	"math/big"
	"testing"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestEncodeRouteToPath(t *testing.T) {
	// packs them for exact input single hop
	p, _ := EncodeRouteToPath(route_0_1, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002", hexutil.Encode(p))

	// packs them correctly for exact output single hop
	p, _ = EncodeRouteToPath(route_0_1, true)
	assert.Equal(t, "0x0000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// packs them correctly for multihop exact input
	p, _ = EncodeRouteToPath(route_0_1_2, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb800000000000000000000000000000000000000020001f40000000000000000000000000000000000000003", hexutil.Encode(p))

	// packs them correctly for multihop exact output
	p, _ = EncodeRouteToPath(route_0_1_2, true)
	assert.Equal(t, "0x00000000000000000000000000000000000000030001f40000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// wraps ether input for exact input single hop
	p, _ = EncodeRouteToPath(route_weth_0, false)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// wraps ether input for exact output single hop
	p, _ = EncodeRouteToPath(route_weth_0, true)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether input for exact input multihop
	p, _ = EncodeRouteToPath(route_weth_0_1, false)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002", hexutil.Encode(p))

	// wraps ether input for exact output multihop
	p, _ = EncodeRouteToPath(route_weth_0_1, true)
	assert.Equal(t, "0x0000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether output for exact input single hop
	p, _ = EncodeRouteToPath(route_0_weth, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether output for exact output single hop
	p, _ = EncodeRouteToPath(route_0_weth, true)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000001", hexutil.Encode(p))

	// wraps ether output for exact input multihop
	p, _ = EncodeRouteToPath(route_0_1_weth, false)
	assert.Equal(t, "0x0000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", hexutil.Encode(p))

	// wraps ether output for exact output multihop
	p, _ = EncodeRouteToPath(route_0_1_weth, true)
	assert.Equal(t, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001", hexutil.Encode(p))
}

func TestQuoteCallParameters(t *testing.T) {
	feeAmount := constants.FeeMedium
	sqrtRatioX96 := utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(1))
	liquidity := big.NewInt(1_000_000)
	WETH := core.WETH9[1]
	tick, _ := utils.GetTickAtSqrtRatio(sqrtRatioX96)
	ticks := []entities.Tick{
		{
			Index:          entities.NearestUsableTick(utils.MinTick, constants.TickSpacings[feeAmount]),
			LiquidityNet:   liquidity,
			LiquidityGross: liquidity,
		},
		{
			Index:          entities.NearestUsableTick(utils.MaxTick, constants.TickSpacings[feeAmount]),
			LiquidityNet:   new(big.Int).Mul(liquidity, constants.NegativeOne),
			LiquidityGross: liquidity,
		},
	}

	p, _ := entities.NewTickListDataProvider(ticks, constants.TickSpacings[feeAmount])
	pool_0_1, _ := entities.NewPool(token0, token1, feeAmount, sqrtRatioX96, liquidity, tick, p)
	pool_1_weth, _ := entities.NewPool(token1, WETH, feeAmount, sqrtRatioX96, liquidity, tick, p)

	// single trade input
	// single-hop exact input
	r, _ := entities.NewRoute([]*entities.Pool{pool_0_1}, token0, token1)
	trade, _ := entities.FromRoute(r, core.FromRawAmount(token0.Currency, big.NewInt(100)), core.ExactInput)
	params, err := QuoteCallParameters(trade.Swaps[0].Route, trade.InputAmount(), trade.TradeType, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0xf7729d43000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb800000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// single-hop exact output
	r, _ = entities.NewRoute([]*entities.Pool{pool_0_1}, token0, token1)
	trade, _ = entities.FromRoute(r, core.FromRawAmount(token1.Currency, big.NewInt(100)), core.ExactOutput)
	params, err = QuoteCallParameters(trade.Swaps[0].Route, trade.OutputAmount(), trade.TradeType, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0x30d07f21000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb800000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// multi-hop exact input
	r, _ = entities.NewRoute([]*entities.Pool{pool_0_1, pool_1_weth}, token0, WETH)
	trade, _ = entities.FromRoute(r, core.FromRawAmount(token0.Currency, big.NewInt(100)), core.ExactInput)
	route, _ := trade.Route()
	params, err = QuoteCallParameters(route, trade.InputAmount(), trade.TradeType, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0xcdca17530000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// multi-hop exact output
	r, _ = entities.NewRoute([]*entities.Pool{pool_0_1, pool_1_weth}, token0, WETH)
	trade, _ = entities.FromRoute(r, core.FromRawAmount(WETH.Currency, big.NewInt(100)), core.ExactOutput)
	route, _ = trade.Route()
	params, err = QuoteCallParameters(route, trade.OutputAmount(), trade.TradeType, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0x2f80bb1d000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000042c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb80000000000000000000000000000000000000002000bb80000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// sqrtPriceLimitX96
	r, _ = entities.NewRoute([]*entities.Pool{pool_0_1}, token0, token1)
	trade, _ = entities.FromRoute(r, core.FromRawAmount(token0.Currency, big.NewInt(100)), core.ExactInput)
	route, _ = trade.Route()
	params, err = QuoteCallParameters(route, trade.InputAmount(), trade.TradeType, &QuoteOptions{SqrtPriceLimitX96: new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0xf7729d43000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb800000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000100000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))
}
