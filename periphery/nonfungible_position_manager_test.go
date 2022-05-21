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
	tokenIDT           = big.NewInt(1)
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
			UseNative:         core.EtherOnChain(1),
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

	// succeeds for increase
	pos, err = entities.NewPosition(pool01T, big.NewInt(1), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &AddLiquidityOptions{
		IncreaseSpecificOptions: &IncreaseSpecificOptions{
			TokenID: tokenIDT,
		},
		CommonAddLiquidityOptions: &CommonAddLiquidityOptions{
			SlippageTolerance: slippageToleranceT,
			Deadline:          deadlineT,
		},
	}
	params, err = AddCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0x219f5d1700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007b", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// createPool
	pos, err = entities.NewPosition(pool01T, big.NewInt(1), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &AddLiquidityOptions{
		CommonAddLiquidityOptions: &CommonAddLiquidityOptions{
			SlippageTolerance: slippageToleranceT,
			Deadline:          deadlineT,
		},
		MintSpecificOptions: &MintSpecificOptions{
			Recipient:  recipientT,
			CreatePool: true,
		},
	}
	params, err = AddCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xac9650d80000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000008413ead562000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb8000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016488316456000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000bb8ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc4000000000000000000000000000000000000000000000000000000000000003c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000007b00000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// useNative
	pos, err = entities.NewPosition(pool1wethT, big.NewInt(1), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &AddLiquidityOptions{
		CommonAddLiquidityOptions: &CommonAddLiquidityOptions{
			SlippageTolerance: slippageToleranceT,
			Deadline:          deadlineT,
			UseNative:         core.EtherOnChain(1),
		},
		MintSpecificOptions: &MintSpecificOptions{
			Recipient: recipientT,
		},
	}
	params, err = AddCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xac9650d800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000164883164560000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000000bb8ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc4000000000000000000000000000000000000000000000000000000000000003c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000007b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000412210e8a00000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x01", utils.ToHex(params.Value))
}

func TestCollectCallParameters(t *testing.T) {
	// works
	opts := &CollectOptions{
		TokenID:               tokenIDT,
		ExpectedCurrencyOwed0: core.FromRawAmount(token0T, big.NewInt(0)),
		ExpectedCurrencyOwed1: core.FromRawAmount(token1T, big.NewInt(0)),
		Recipient:             recipientT,
	}
	params, err := CollectCallParameters(opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xfc6f78650000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// works with eth
	opts = &CollectOptions{
		TokenID:               tokenIDT,
		ExpectedCurrencyOwed0: core.FromRawAmount(token1T, big.NewInt(0)),
		ExpectedCurrencyOwed1: core.FromRawAmount(ether, big.NewInt(0)),
		ExpectedTokenOwed0:    token1T,
		ExpectedTokenOwed1:    ether,
		Recipient:             recipientT,
	}
	params, err = CollectCallParameters(opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xac9650d8000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000084fc6f78650000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004449404b7c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064df2ab5bb00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))
}

func TestRemoveCallParameters(t *testing.T) {
	// throws for 0 liquidity
	pos, err := entities.NewPosition(pool01T, big.NewInt(0), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts := &RemoveLiquidityOptions{
		TokenID:             tokenIDT,
		LiquidityPercentage: core.NewPercent(big.NewInt(1), big.NewInt(1)),
		SlippageTolerance:   slippageToleranceT,
		Deadline:            deadlineT,
		CollectOptions: &CollectOptions{
			ExpectedCurrencyOwed0: core.FromRawAmount(token0T, big.NewInt(0)),
			ExpectedCurrencyOwed1: core.FromRawAmount(token1T, big.NewInt(0)),
			Recipient:             recipientT,
		},
	}
	_, err = RemoveCallParameters(pos, opts)
	assert.Error(t, err, ErrZeroLiquidity)

	// throws for 0 liquidity from small percentage
	pos, err = entities.NewPosition(pool01T, big.NewInt(50), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &RemoveLiquidityOptions{
		TokenID:             tokenIDT,
		LiquidityPercentage: core.NewPercent(big.NewInt(1), big.NewInt(100)),
		SlippageTolerance:   slippageToleranceT,
		Deadline:            deadlineT,
		CollectOptions: &CollectOptions{
			ExpectedCurrencyOwed0: core.FromRawAmount(token0T, big.NewInt(0)),
			ExpectedCurrencyOwed1: core.FromRawAmount(token1T, big.NewInt(0)),
			Recipient:             recipientT,
		},
	}
	_, err = RemoveCallParameters(pos, opts)
	assert.Error(t, err, ErrZeroLiquidity)

	// throws for bad burn
	pos, err = entities.NewPosition(pool01T, big.NewInt(50), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &RemoveLiquidityOptions{
		TokenID:             tokenIDT,
		LiquidityPercentage: core.NewPercent(big.NewInt(99), big.NewInt(100)),
		SlippageTolerance:   slippageToleranceT,
		Deadline:            deadlineT,
		BurnToken:           true,
		CollectOptions: &CollectOptions{
			ExpectedCurrencyOwed0: core.FromRawAmount(token0T, big.NewInt(0)),
			ExpectedCurrencyOwed1: core.FromRawAmount(token1T, big.NewInt(0)),
			Recipient:             recipientT,
		},
	}
	_, err = RemoveCallParameters(pos, opts)
	assert.Error(t, err, ErrCannotBurn)

	// works
	pos, err = entities.NewPosition(pool01T, big.NewInt(100), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &RemoveLiquidityOptions{
		TokenID:             tokenIDT,
		LiquidityPercentage: core.NewPercent(big.NewInt(1), big.NewInt(1)),
		SlippageTolerance:   slippageToleranceT,
		Deadline:            deadlineT,
		CollectOptions: &CollectOptions{
			ExpectedCurrencyOwed0: core.FromRawAmount(token0T, big.NewInt(0)),
			ExpectedCurrencyOwed1: core.FromRawAmount(token1T, big.NewInt(0)),
			Recipient:             recipientT,
		},
	}
	params, err := RemoveCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xac9650d8000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000a40c49ccbe0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000084fc6f78650000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// works for partial
	pos, err = entities.NewPosition(pool01T, big.NewInt(100), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &RemoveLiquidityOptions{
		TokenID:             tokenIDT,
		LiquidityPercentage: core.NewPercent(big.NewInt(1), big.NewInt(2)),
		SlippageTolerance:   slippageToleranceT,
		Deadline:            deadlineT,
		CollectOptions: &CollectOptions{
			ExpectedCurrencyOwed0: core.FromRawAmount(token0T, big.NewInt(0)),
			ExpectedCurrencyOwed1: core.FromRawAmount(token1T, big.NewInt(0)),
			Recipient:             recipientT,
		},
	}
	params, err = RemoveCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xac9650d8000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000a40c49ccbe0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000003200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000084fc6f78650000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// works with eth
	ethAmount := core.FromRawAmount(core.EtherOnChain(1), big.NewInt(0))
	tokenAmount := core.FromRawAmount(token1T, big.NewInt(0))
	var (
		owed0Amount *core.CurrencyAmount
		owed1Amount *core.CurrencyAmount
		owed0Token  *core.Token
		owed1Token  *core.Token
	)
	if pool1wethT.Token0.Equal(token1T) {
		owed0Amount = tokenAmount
		owed1Amount = ethAmount
		owed0Token = token1T
		owed1Token = core.EtherOnChain(1).Wrapped()
	} else {
		owed0Amount = ethAmount
		owed1Amount = tokenAmount
		owed0Token = core.EtherOnChain(1).Wrapped()
		owed1Token = token1T
	}
	pos, err = entities.NewPosition(pool01T, big.NewInt(100), -constants.TickSpacings[constants.FeeMedium], constants.TickSpacings[constants.FeeMedium])
	assert.NoError(t, err)
	opts = &RemoveLiquidityOptions{
		TokenID:             tokenIDT,
		LiquidityPercentage: core.NewPercent(big.NewInt(1), big.NewInt(1)),
		SlippageTolerance:   slippageToleranceT,
		Deadline:            deadlineT,
		CollectOptions: &CollectOptions{
			ExpectedCurrencyOwed0: owed0Amount,
			ExpectedCurrencyOwed1: owed1Amount,
			ExpectedTokenOwed0:    owed0Token,
			ExpectedTokenOwed1:    owed1Token,
			Recipient:             recipientT,
		},
	}
	params, err = RemoveCallParameters(pos, opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xac9650d80000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000000a40c49ccbe0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000084fc6f78650000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004449404b7c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064df2ab5bb00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))
}

func TestSafeTransferFromParameters(t *testing.T) {
	// succeeds no data param
	opts := &SafeTransferOptions{
		Sender:    senderT,
		Recipient: recipientT,
		TokenID:   tokenIDT,
	}
	params, err := SafeTransferFromParameters(opts)
	assert.NoError(t, err)
	assert.Equal(t, "0x42842e0e000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000001", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))

	// succeeds data param
	opts = &SafeTransferOptions{
		Sender:    senderT,
		Recipient: recipientT,
		TokenID:   tokenIDT,
		Data:      common.FromHex("0x0000000000000000000000000000000000009004"),
	}
	params, err = SafeTransferFromParameters(opts)
	assert.NoError(t, err)
	assert.Equal(t, "0xb88d4fde000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000009004000000000000000000000000", hexutil.Encode(params.Calldata))
	assert.Equal(t, "0x00", utils.ToHex(params.Value))
}
