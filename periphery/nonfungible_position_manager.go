package periphery

import (
	"encoding/json"
	"errors"
	"math/big"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

//go:embed contracts/interfaces/INonfungiblePositionManager.sol/INonfungiblePositionManager.json
var nonFungiblePositionManagerABI []byte

var (
	ErrZeroLiquidity = errors.New("zero liquidity")
	ErrNoWETH        = errors.New("no WETH")
)

func getNonFungiblePositionManagerABI() abi.ABI {
	var wabi WrappedABI
	err := json.Unmarshal(nonFungiblePositionManagerABI, &wabi)
	if err != nil {
		panic(err)
	}
	return wabi.ABI
}

var MaxUint128 = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil), big.NewInt(1))

type MintSpecificOptions struct {
	Recipient  common.Address // The account that should receive the minted NFT
	CreatePool bool           // Creates pool if not initialized before mint
}

type IncreaseSpecificOptions struct {
	TokenID *big.Int // Indicates the ID of the position to increase liquidity for
}

//  Options for producing the calldata to add liquidity
type CommonAddLiquidityOptions struct {
	SlippageTolerance *core.Percent  // How much the pool price is allowed to move
	Deadline          *big.Int       // When the transaction expires, in epoch seconds
	UseNative         *core.Currency // Whether to spend ether. If true, one of the pool tokens must be WETH, by default false
	NativeToken       *core.Token    // TODO: merge this with UseNative

	Token0Permit *PermitOptions // The optional permit parameters for spending token0
	Token1Permit *PermitOptions // The optional permit parameters for spending token1
}

type MintOptions struct {
	*CommonAddLiquidityOptions
	*MintSpecificOptions
}

type IncreaseOptions struct {
	*CommonAddLiquidityOptions
	*IncreaseSpecificOptions
}

type AddLiquidityOptions struct {
	*CommonAddLiquidityOptions
	*MintSpecificOptions
	*IncreaseSpecificOptions
}

type SafeTransferOptions struct {
	Sender    string   // The account sending the NFT
	Recipient string   // The account that should receive the NFT
	TokenID   *big.Int //  The id of the token being sent
	Data      string   // The optional parameter that passes data to the `onERC721Received` call for the staker
}

type CollectOptions struct {
	TokenID               *big.Int             // Indicates the ID of the position to collect for
	ExpectedCurrencyOwed0 *core.CurrencyAmount // Expected value of tokensOwed0, including as-of-yet-unaccounted-for fees/liquidity value to be burned
	ExpectedCurrencyOwed1 *core.CurrencyAmount // Expected value of tokensOwed1, including as-of-yet-unaccounted-for fees/liquidity value to be burned
	ExpectedTokenOwed0    *core.Token          // TODO: merge this with Currency
	ExpectedTokenOwed1    *core.Token          // TODO: merge this with Currency
	Recipient             common.Address       // The account that should receive the tokens
}

type NFTPermitOptions struct {
	V        uint
	R        string
	S        string
	Deadline *big.Int
	Spender  string
}

// Options for producing the calldata to exit a position
type RemoveLiquidityOptions struct {
	TokenID             *big.Int          // The ID of the token to exit
	LiquidityPercentage *core.Percent     // The percentage of position liquidity to exit
	SlippageTolerance   *core.Percent     // How much the pool price is allowed to move
	Deadline            *big.Int          // When the transaction expires, in epoch seconds.
	BurnToken           bool              // Whether the NFT should be burned if the entire position is being exited, by default false
	Permit              *NFTPermitOptions // The optional permit of the token ID being exited, in case the exit transaction is being sent by an account that does not own the NFT
	CollectOptions      *CollectOptions   // Parameters to be passed on to collect
}

type NonFungiblePositionManager struct{}

func encodeCreate(pool *entities.Pool) ([]byte, error) {
	abi := getNonFungiblePositionManagerABI()
	return abi.Pack("createAndInitializePoolIfNecessary", pool.Token0.Address, pool.Token1.Address, pool.Fee, pool.SqrtRatioX96)
}

func CreateCallParameters(pool *entities.Pool) (*utils.MethodParameters, error) {
	calldata, err := encodeCreate(pool)
	if err != nil {
		return nil, err
	}
	return &utils.MethodParameters{
		Calldata: calldata,
		Value:    constants.Zero,
	}, nil
}

func AddCallParameters(position *entities.Position, opts *AddLiquidityOptions) (*utils.MethodParameters, error) {
	if position.Liquidity.Cmp(constants.Zero) <= 0 {
		return nil, ErrZeroLiquidity
	}

	var calldatas [][]byte

	// get amounts
	amount0Desired, amount1Desired, err := position.MintAmounts()
	if err != nil {
		return nil, err
	}

	// adjust for slippage
	amount0Min, amount1Min, err := position.MintAmountsWithSlippage(opts.SlippageTolerance)
	if err != nil {
		return nil, err
	}

	// create pool if needed
	if opts.MintSpecificOptions != nil && opts.MintSpecificOptions.CreatePool {
		calldata, err := encodeCreate(position.Pool)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, calldata)
	}

	// permits if necessary
	if opts.Token0Permit != nil {
		calldata, err := EncodePermit(position.Pool.Token0, opts.Token0Permit)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, calldata)
	}
	if opts.Token1Permit != nil {
		calldata, err := EncodePermit(position.Pool.Token1, opts.Token1Permit)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, calldata)
	}

	abi := getNonFungiblePositionManagerABI()

	// mint
	if opts.MintSpecificOptions != nil {
		calldata, err := abi.Pack("mint",
			position.Pool.Token0.Address,
			position.Pool.Token1.Address,
			position.Pool.Fee,
			position.TickLower,
			position.TickUpper,
			amount0Desired,
			amount1Desired,
			amount0Min,
			amount1Min,
			opts.Recipient,
			opts.Deadline,
		)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, calldata)
	}

	// increase
	if opts.IncreaseSpecificOptions != nil {
		calldata, err := abi.Pack("increaseLiquidity",
			opts.TokenID,
			amount0Desired,
			amount1Desired,
			amount0Min,
			amount1Min,
			opts.Deadline,
		)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, calldata)
	}

	value := constants.Zero
	if opts.UseNative != nil {
		if !position.Pool.Token0.Equals(opts.NativeToken) && !position.Pool.Token1.Equals(opts.NativeToken) {
			return nil, ErrNoWETH
		}

		if position.Pool.Token0.Equals(opts.NativeToken) {
			value = amount0Desired
		} else {
			value = amount1Desired
		}

		// we only need to refund if we're actually sending ETH
		if value.Cmp(constants.Zero) > 0 {
			calldatas = append(calldatas, EncodeRefundETH())
		}
	}

	datas, err := EncodeMulticall(calldatas)
	if err != nil {
		return nil, err
	}

	return &utils.MethodParameters{
		Calldata: datas,
		Value:    value,
	}, nil
}

func encodeCollect(opts *CollectOptions) ([][]byte, error) {
	var calldatas [][]byte

	involvesETH := opts.ExpectedCurrencyOwed0.Currency.IsNative || opts.ExpectedCurrencyOwed1.Currency.IsNative

	// collect
	abi := getNonFungiblePositionManagerABI()
	collectRecipent := opts.Recipient
	if involvesETH {
		collectRecipent = constants.AddressZero
	}
	calldata, err := abi.Pack("collect", opts.TokenID, collectRecipent, MaxUint128, MaxUint128)
	if err != nil {
		return nil, err
	}
	calldatas = append(calldatas, calldata)

	if involvesETH {
		var (
			ethAmount   *big.Int
			token       *core.Token
			tokenAmount *big.Int
		)
		if opts.ExpectedCurrencyOwed0.Currency.IsNative {
			ethAmount = opts.ExpectedCurrencyOwed0.Quotient()
			token = opts.ExpectedTokenOwed1
			tokenAmount = opts.ExpectedCurrencyOwed1.Quotient()
		} else {
			ethAmount = opts.ExpectedCurrencyOwed1.Quotient()
			token = opts.ExpectedTokenOwed0
			tokenAmount = opts.ExpectedCurrencyOwed0.Quotient()
		}

		weth9data, err := EncodeUnwrapWETH9(ethAmount, opts.Recipient, nil)
		if err != nil {
			return nil, err
		}
		sweepdata, err := EncodeSweepToken(token, tokenAmount, opts.Recipient, nil)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, weth9data, sweepdata)
	}

	return calldatas, nil
}

func CollectCallParameters(opts *CollectOptions) (*utils.MethodParameters, error) {
	calldatas, err := encodeCollect(opts)
	if err != nil {
		return nil, err
	}

	data, err := EncodeMulticall(calldatas)
	if err != nil {
		return nil, err
	}
	return &utils.MethodParameters{
		Calldata: data,
		Value:    constants.Zero,
	}, nil
}

/**
 * Produces the calldata for completely or partially exiting a position
 * @param position The position to exit
 * @param options Additional information necessary for generating the calldata
 * @returns The call parameters
 */
// func RemoveCallParameters(position *entities.Position, opts *RemoveLiquidityOptions) (*utils.MethodParameters, error) {
// 	var calldatas [][]byte

// 	// construct a partial position with a percentage of liquidity
// 	partialPosition, err := entities.NewPosition(
// 		position.Pool,
// 		opts.LiquidityPercentage.Multiply(core.NewPercent(position.Liquidity, big.NewInt(1))).Quotient(),
// 		position.TickLower,
// 		position.TickUpper,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if partialPosition.Liquidity.Cmp(constants.Zero) <= 0 {
// 		return nil, ErrZeroLiquidity
// 	}

// 	// slippage-adjusted underlying amounts
// 	amount0Min, amount1Min, err := partialPosition.BurnAmountsWithSlippage(opts.SlippageTolerance)
// 	if err != nil {
// 		return nil, err
// 	}

// 	abi := getNonFungiblePositionManagerABI()
// 	if opts.Permit != nil {
// 		calldata, err := abi.Pack("permit", opts.Permit.Spender, opts.TokenID, opts.Permit.Deadline, opts.Permit.V, opts.Permit.R, opts.Permit.S)
// 		if err != nil {
// 			return nil, err
// 		}
// 		calldatas = append(calldatas, calldata)
// 	}

// 	// remove liquidity
// 	calldata, err := abi.Pack("decreaseLiquidity", opts.TokenID, partialPosition.Liquidity, amount0Min, amount1Min, opts.Deadline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	calldatas = append(calldatas, calldata)

// }
