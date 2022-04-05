package periphery

import (
	"math/big"
	"testing"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestSwapCallParameters(t *testing.T) {
	pool_0_1 := makePool(token0, token1)
	pool_1_weth := makePool(token1, weth)
	// pool_0_2 := makePool(token0, token2)
	// pool_0_3 := makePool(token0, token3)
	// pool_2_3 := makePool(token2, token3)
	// pool_3_weth := makePool(token3, weth)
	// pool_1_3 := makePool(token3, token1)

	slippageTolerance := core.NewPercent(big.NewInt(1), big.NewInt(100))
	recipient := common.HexToAddress("0x0000000000000000000000000000000000000003")
	deadline := big.NewInt(123)

	// single trade input
	// single-hop exact input
	r, _ := entities.NewRoute([]*entities.Pool{pool_0_1}, token0, token1)
	trade, _ := entities.FromRoute(r, core.FromRawAmount(token0.Currency, big.NewInt(100)), core.ExactInput)
	params, err := SwapCallParameters([]*entities.Trade{trade}, token0, token1, &SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         recipient,
		Deadline:          deadline,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0x414bf389000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000007b000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000610000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// single-hop exact output
	r, _ = entities.NewRoute([]*entities.Pool{pool_0_1}, token0, token1)
	trade, _ = entities.FromRoute(r, core.FromRawAmount(token1.Currency, big.NewInt(100)), core.ExactOutput)
	params, err = SwapCallParameters([]*entities.Trade{trade}, token0, token1, &SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         recipient,
		Deadline:          deadline,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0xdb3e2198000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000007b000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000670000000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// multi-hop exact input
	r, _ = entities.NewRoute([]*entities.Pool{pool_0_1, pool_1_weth}, token0, weth)
	trade, _ = entities.FromRoute(r, core.FromRawAmount(token0.Currency, big.NewInt(100)), core.ExactInput)
	params, err = SwapCallParameters([]*entities.Trade{trade}, token0, weth, &SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         recipient,
		Deadline:          deadline,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0xc04b8d59000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000005f00000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000000001000bb80000000000000000000000000000000000000002000bb8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

}
