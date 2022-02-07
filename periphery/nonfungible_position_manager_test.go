package periphery

import (
	"math/big"
	"testing"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

var (
	token0T = core.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000001"), 18, "t0", "token0")
	token1T = core.NewToken(1, common.HexToAddress("0x0000000000000000000000000000000000000002"), 18, "t1", "token1")

	feeT = constants.FeeMedium

	pool01T, _    = entities.NewPool(token0T, token1T, feeT, utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(1)), big.NewInt(0), 0, nil)
	pool1wethT, _ = entities.NewPool(token1T, core.WETH9[1], feeT, utils.EncodeSqrtRatioX96(big.NewInt(1), big.NewInt(1)), big.NewInt(0), 0, nil)

	recipientT         = common.HexToAddress("0x0000000000000000000000000000000000000003")
	senderT            = common.HexToAddress("0x0000000000000000000000000000000000000004")
	tokenIDT           = 1
	slippageToleranceT = core.NewPercent(big.NewInt(1), big.NewInt(100))
	deadlineT          = big.NewInt(123)
)

func TestCreateCallParameters(t *testing.T) {
	params, err := CreateCallParameters(pool01T)
	assert.NoError(t, err)
	assert.Equal(t, "0x13ead562000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000001000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))
}

func TestAddCallParameters(t *testing.T) {
	// throws if liquidity is 0
	pos, err := entities.NewPosition(pool01T, big.NewInt(0), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts := &AddLiquidityOptions{
		MintSpecificOptions: &MintSpecificOptions{
			Recipient: recipientT,
		},
		CommonAddLiquidityOptions: &CommonAddLiquidityOptions{
			SlippageTolerance: slippageToleranceT,
			Deadline:          deadlineT,
		},
	}
	_, err = AddCallParameters(pos, opts)
	assert.ErrorIs(t, err, ErrZeroLiquidity)

	// throws if pool does not involve ether and useNative is true
	pos, err = entities.NewPosition(pool01T, big.NewInt(1), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &AddLiquidityOptions{
		MintSpecificOptions: &MintSpecificOptions{
			Recipient: recipientT,
		},
		CommonAddLiquidityOptions: &CommonAddLiquidityOptions{
			SlippageTolerance: slippageToleranceT,
			Deadline:          deadlineT,
			UseNative:         core.EtherOnChain(1).Currency,
			NativeToken:       core.WETH9[1],
		},
	}
	_, err = AddCallParameters(pos, opts)
	assert.ErrorIs(t, err, ErrNoWETH)

	// succeeds for mint
	pos, err = entities.NewPosition(pool01T, big.NewInt(1), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &AddLiquidityOptions{
		MintSpecificOptions: &MintSpecificOptions{
			Recipient: recipientT,
		},
		CommonAddLiquidityOptions: &CommonAddLiquidityOptions{
			SlippageTolerance: slippageToleranceT,
			Deadline:          deadlineT,
		},
	}
	params, err := AddCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0x88316456000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb8ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc4000000000000000000000000000000000000000000000000000000000000003c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000007b", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))
}
