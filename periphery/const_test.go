package periphery

import (
	"math/big"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ether  = core.EtherOnChain(1).Wrapped()
	token0 = core.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "t0", "token0")
	token1 = core.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "t1", "token1")
	token2 = core.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000003"), 18, "t2", "token2")

	weth = core.WETH9[1]

	pool_0_1_medium, _ = entities.NewPool(token0, token1, constants.FeeMedium, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)
	pool_1_2_low, _    = entities.NewPool(token1, token2, constants.FeeLow, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)
	pool_0_weth, _     = entities.NewPool(token0, weth, constants.FeeMedium, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)
	pool_1_weth, _     = entities.NewPool(token1, weth, constants.FeeMedium, utils.EncodeSqrtRatioX96(constants.One, constants.One), big.NewInt(0), 0, nil)

	route_0_1, _   = entities.NewRoute([]*entities.Pool{pool_0_1_medium}, token0, token1)
	route_0_1_2, _ = entities.NewRoute([]*entities.Pool{pool_0_1_medium, pool_1_2_low}, token0, token2)

	route_0_weth, _   = entities.NewRoute([]*entities.Pool{pool_0_weth}, token0, ether)
	route_0_1_weth, _ = entities.NewRoute([]*entities.Pool{pool_0_1_medium, pool_1_weth}, token0, ether)
	route_weth_0, _   = entities.NewRoute([]*entities.Pool{pool_0_weth}, ether, token0)
	route_weth_0_1, _ = entities.NewRoute([]*entities.Pool{pool_0_weth, pool_0_1_medium}, ether, token1)
)
