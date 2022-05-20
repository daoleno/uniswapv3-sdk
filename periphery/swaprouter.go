package periphery

import (
	_ "embed"
	"errors"
	"math/big"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

//go:embed contracts/SwapRouter.sol/SwapRouter.json
var swapRouterABI []byte

var (
	ErrTokenInDiff        = errors.New("TOKEN_IN_DIFF")
	ErrTokenOutDiff       = errors.New("TOKEN_OUT_DIFF")
	ErrNonTokenPermit     = errors.New("NON_TOKEN_PERMIT")
	ErrMultiHopPriceLimit = errors.New("MULTIHOP_PRICE_LIMIT")
)

// Options for producing the arguments to send calls to the router.
type SwapOptions struct {
	SlippageTolerance *core.Percent  // How much the execution price is allowed to move unfavorably from the trade execution price.
	Recipient         common.Address // The account that should receive the output.
	Deadline          *big.Int       // When the transaction expires, in epoch seconds.
	InputTokenPermit  *PermitOptions // The optional permit parameters for spending the input.
	SqrtPriceLimitX96 *big.Int       // The optional price limit for the trade.
	Fee               *FeeOptions    // Optional information for taking a fee on output.
}

type ExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

type ExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

type ExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

type ExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// Represents the Uniswap V3 SwapRouter

/**
 * Produces the on-chain method name to call and the hex encoded parameters to pass as arguments for a given trade.
 * @param trade to produce call parameters for
 * @param options options for the call parameters
 */
func SwapCallParameters(trades []*entities.Trade, options *SwapOptions) (*utils.MethodParameters, error) {
	abi := GetABI(swapRouterABI)
	sampleTrade := trades[0]
	tokenIn := sampleTrade.InputAmount().Currency.Wrapped()
	tokenOut := sampleTrade.OutputAmount().Currency.Wrapped()

	// All trades should have the same starting and ending token.
	for _, trade := range trades {
		if !trade.InputAmount().Currency.Wrapped().Equal(tokenIn) {
			return nil, ErrTokenInDiff
		}
		if !trade.OutputAmount().Currency.Wrapped().Equal(tokenOut) {
			return nil, ErrTokenOutDiff
		}
	}

	var calldatas [][]byte

	ZeroIn := core.FromRawAmount(trades[0].InputAmount().Currency, big.NewInt(0))
	ZeroOut := core.FromRawAmount(trades[0].OutputAmount().Currency, big.NewInt(0))

	totalAmountOut := ZeroOut
	for _, trade := range trades {
		minOut, err := trade.MinimumAmountOut(options.SlippageTolerance, nil)
		if err != nil {
			return nil, err
		}
		totalAmountOut = totalAmountOut.Add(minOut)
	}

	// flag for whether a refund needs to happen
	mustRefund := sampleTrade.InputAmount().Currency.IsNative() && sampleTrade.TradeType == core.ExactOutput
	inputIsNative := sampleTrade.InputAmount().Currency.IsNative()
	// flags for whether funds should be send first to the router
	outputIsNative := sampleTrade.OutputAmount().Currency.IsNative()
	routerMustCustody := outputIsNative || options.Fee != nil

	totalValue := ZeroIn
	if inputIsNative {
		for _, trade := range trades {
			maxIn, err := trade.MaximumAmountIn(options.SlippageTolerance, nil)
			if err != nil {
				return nil, err
			}
			totalValue = totalValue.Add(maxIn)
		}
	}

	// encode permit if necessary
	if options.InputTokenPermit != nil {
		if !sampleTrade.InputAmount().Currency.IsToken() {
			return nil, ErrNonTokenPermit
		}

		permit, err := EncodePermit(tokenIn, options.InputTokenPermit)
		if err != nil {
			return nil, err
		}
		calldatas = append(calldatas, permit)
	}

	recipient := options.Recipient
	if routerMustCustody {
		recipient = constants.AddressZero
	}

	sqrtPriceLimitX96 := big.NewInt(0)
	if options != nil && options.SqrtPriceLimitX96 != nil {
		sqrtPriceLimitX96 = options.SqrtPriceLimitX96
	}

	for _, trade := range trades {
		for _, swap := range trade.Swaps {
			amountIn, err := trade.MaximumAmountIn(options.SlippageTolerance, swap.InputAmount)
			if err != nil {
				return nil, err
			}
			amountOut, err := trade.MinimumAmountOut(options.SlippageTolerance, swap.OutputAmount)
			if err != nil {
				return nil, err
			}

			// flag for whether the trade is single hop or not
			singleHop := len(swap.Route.Pools) == 1

			if singleHop {
				if trade.TradeType == core.ExactInput {

					exactInputSingleParams := &ExactInputSingleParams{
						TokenIn:           swap.Route.TokenPath[0].Address,
						TokenOut:          swap.Route.TokenPath[1].Address,
						Fee:               big.NewInt(int64(swap.Route.Pools[0].Fee)),
						Recipient:         recipient,
						Deadline:          options.Deadline,
						AmountIn:          amountIn.Quotient(),
						AmountOutMinimum:  amountOut.Quotient(),
						SqrtPriceLimitX96: sqrtPriceLimitX96,
					}
					calldata, err := abi.Pack("exactInputSingle", exactInputSingleParams)
					if err != nil {
						return nil, err
					}
					calldatas = append(calldatas, calldata)
				} else {
					exactOutputSingleParams := &ExactOutputSingleParams{
						TokenIn:           swap.Route.TokenPath[0].Address,
						TokenOut:          swap.Route.TokenPath[1].Address,
						Fee:               big.NewInt(int64(swap.Route.Pools[0].Fee)),
						Recipient:         recipient,
						Deadline:          options.Deadline,
						AmountOut:         amountOut.Quotient(),
						AmountInMaximum:   amountIn.Quotient(),
						SqrtPriceLimitX96: sqrtPriceLimitX96,
					}
					calldata, err := abi.Pack("exactOutputSingle", exactOutputSingleParams)
					if err != nil {
						return nil, err
					}
					calldatas = append(calldatas, calldata)
				}
			} else {
				if options != nil && options.SqrtPriceLimitX96 != nil {
					return nil, ErrMultiHopPriceLimit
				}

				path, err := EncodeRouteToPath(swap.Route, trade.TradeType == core.ExactOutput)
				if err != nil {
					return nil, err
				}

				if trade.TradeType == core.ExactInput {
					exactInputParams := &ExactInputParams{
						Path:             path,
						Recipient:        recipient,
						Deadline:         options.Deadline,
						AmountIn:         amountIn.Quotient(),
						AmountOutMinimum: amountOut.Quotient(),
					}
					calldata, err := abi.Pack("exactInput", exactInputParams)
					if err != nil {
						return nil, err
					}
					calldatas = append(calldatas, calldata)
				} else {
					exactOutputParams := &ExactOutputParams{
						Path:            path,
						Recipient:       recipient,
						Deadline:        options.Deadline,
						AmountOut:       amountOut.Quotient(),
						AmountInMaximum: amountIn.Quotient(),
					}
					calldata, err := abi.Pack("exactOutput", exactOutputParams)
					if err != nil {
						return nil, err
					}
					calldatas = append(calldatas, calldata)
				}
			}
		}
	}

	// unwrap
	if routerMustCustody {
		if options.Fee != nil {
			if outputIsNative {
				calldata, err := EncodeUnwrapWETH9(totalAmountOut.Quotient(), options.Recipient, options.Fee)
				if err != nil {
					return nil, err
				}
				calldatas = append(calldatas, calldata)
			} else {
				calldata, err := EncodeSweepToken(tokenOut, totalAmountOut.Quotient(), options.Recipient, options.Fee)
				if err != nil {
					return nil, err
				}
				calldatas = append(calldatas, calldata)
			}
		} else {
			calldata, err := EncodeUnwrapWETH9(totalAmountOut.Quotient(), options.Recipient, nil)
			if err != nil {
				return nil, err
			}
			calldatas = append(calldatas, calldata)
		}
	}

	// refund
	if mustRefund {
		calldatas = append(calldatas, EncodeRefundETH())
	}
	call, err := EncodeMulticall(calldatas)
	if err != nil {
		return nil, err
	}
	return &utils.MethodParameters{
		Calldata: call,
		Value:    totalValue.Quotient(),
	}, nil
}
