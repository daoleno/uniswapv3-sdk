package entities

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrFeeTooHigh               = errors.New("Fee too high")
	ErrInvalidSqrtRatioX96      = errors.New("Invalid sqrtRatioX96")
	ErrTokenNotInvolved         = errors.New("Token not involved in pool")
	ErrSqrtPriceLimitX96TooLow  = errors.New("SqrtPriceLimitX96 too low")
	ErrSqrtPriceLimitX96TooHigh = errors.New("SqrtPriceLimitX96 too high")
)

type StepComputations struct {
	sqrtPriceStartX96 *big.Int
	tickNext          int
	initialized       bool
	sqrtPriceNextX96  *big.Int
	amountIn          *big.Int
	amountOut         *big.Int
	feeAmount         *big.Int
}

// Represents a V3 pool
type Pool struct {
	Token0           *entities.Token
	Token1           *entities.Token
	Fee              constants.FeeAmount
	SqrtRatioX96     *big.Int
	Liquidity        *big.Int
	TickCurrent      int
	TickDataProvider TickDataProvider

	token0Price *entities.Price
	token1Price *entities.Price
}

func GetAddress(tokenA, tokenB *entities.Token, fee constants.FeeAmount, initCodeHashManualOverride string) (common.Address, error) {
	return utils.ComputePoolAddress(constants.FactoryAddress, tokenA, tokenB, fee, initCodeHashManualOverride)
}

/**
 * Construct a pool
 * @param tokenA One of the tokens in the pool
 * @param tokenB The other token in the pool
 * @param fee The fee in hundredths of a bips of the input amount of every swap that is collected by the pool
 * @param sqrtRatioX96 The sqrt of the current ratio of amounts of token1 to token0
 * @param liquidity The current value of in range liquidity
 * @param tickCurrent The current tick of the pool
 * @param ticks The current state of the pool ticks or a data provider that can return tick data
 */
func NewPool(tokenA, tokenB *entities.Token, fee constants.FeeAmount, sqrtRatioX96 *big.Int, liquidity *big.Int, tickCurrent int, ticks TickDataProvider) (*Pool, error) {
	if fee >= constants.FeeMax {
		return nil, ErrFeeTooHigh
	}

	tickCurrentSqrtRatioX96, err := utils.GetSqrtRatioAtTick(tickCurrent)
	if err != nil {
		return nil, err
	}
	nextTickSqrtRatioX96, err := utils.GetSqrtRatioAtTick(tickCurrent + 1)
	if err != nil {
		return nil, err
	}

	if sqrtRatioX96.Cmp(tickCurrentSqrtRatioX96) < 0 || sqrtRatioX96.Cmp(nextTickSqrtRatioX96) > 0 {
		return nil, ErrInvalidSqrtRatioX96
	}
	token0 := tokenA
	token1 := tokenB
	isSorted, err := tokenA.SortsBefore(tokenB)
	if err != nil {
		return nil, err
	}
	if !isSorted {
		token0 = tokenB
		token1 = tokenA
	}

	return &Pool{
		Token0:           token0,
		Token1:           token1,
		Fee:              fee,
		SqrtRatioX96:     sqrtRatioX96,
		Liquidity:        liquidity,
		TickCurrent:      tickCurrent,
		TickDataProvider: ticks, // TODO: new tick data provider
	}, nil
}

/**
 * Returns true if the token is either token0 or token1
 * @param token The token to check
 * @returns True if token is either token0 or token
 */
func (p *Pool) InvolvesToken(token *entities.Token) bool {
	return p.Token0.Equal(token) || p.Token1.Equal(token)
}

// Token0Price returns the current mid price of the pool in terms of token0, i.e. the ratio of token1 over token0
func (p *Pool) Token0Price() *entities.Price {
	if p.token0Price != nil {
		return p.token0Price
	}
	p.token0Price = entities.NewPrice(p.Token0, p.Token1, constants.Q192, new(big.Int).Mul(p.SqrtRatioX96, p.SqrtRatioX96))
	return p.token0Price
}

// Token1Price returns the current mid price of the pool in terms of token1, i.e. the ratio of token0 over token1
func (p *Pool) Token1Price() *entities.Price {
	if p.token1Price != nil {
		return p.token1Price
	}
	p.token1Price = entities.NewPrice(p.Token1, p.Token0, new(big.Int).Mul(p.SqrtRatioX96, p.SqrtRatioX96), constants.Q192)
	return p.token1Price
}

/**
 * Return the price of the given token in terms of the other token in the pool.
 * @param token The token to return price of
 * @returns The price of the given token, in terms of the other.
 */
func (p *Pool) PriceOf(token *entities.Token) (*entities.Price, error) {
	if !p.InvolvesToken(token) {
		return nil, ErrTokenNotInvolved
	}
	if p.Token0.Equal(token) {
		return p.Token0Price(), nil
	}
	return p.Token1Price(), nil
}

// ChainId returns the chain ID of the tokens in the pool.
func (p *Pool) ChainID() uint {
	return p.Token0.ChainId()
}

/**
 * Given an input amount of a token, return the computed output amount, and a pool with state updated after the trade
 * @param inputAmount The input amount for which to quote the output amount
 * @param sqrtPriceLimitX96 The Q64.96 sqrt price limit
 * @returns The output amount and the pool with updated state
 */
func (p *Pool) GetOutputAmount(inputAmount *entities.CurrencyAmount, sqrtPriceLimitX96 *big.Int) (*entities.CurrencyAmount, *Pool, error) {
	if !(inputAmount.Currency.IsToken() && p.InvolvesToken(inputAmount.Currency.Wrapped())) {
		return nil, nil, ErrTokenNotInvolved
	}
	zeroForOne := inputAmount.Currency.Equal(p.Token0)
	outputAmount, sqrtRatioX96, liquidity, tickCurrent, err := p.swap(zeroForOne, inputAmount.Quotient(), sqrtPriceLimitX96)
	if err != nil {
		return nil, nil, err
	}
	var outputToken *entities.Token
	if zeroForOne {
		outputToken = p.Token1
	} else {
		outputToken = p.Token0
	}
	pool, err := NewPool(p.Token0, p.Token1, p.Fee, sqrtRatioX96, liquidity, tickCurrent, p.TickDataProvider)
	if err != nil {
		return nil, nil, err
	}
	return entities.FromRawAmount(outputToken, new(big.Int).Mul(outputAmount, constants.NegativeOne)), pool, nil
}

/**
 * Given a desired output amount of a token, return the computed input amount and a pool with state updated after the trade
 * @param outputAmount the output amount for which to quote the input amount
 * @param sqrtPriceLimitX96 The Q64.96 sqrt price limit. If zero for one, the price cannot be less than this value after the swap. If one for zero, the price cannot be greater than this value after the swap
 * @returns The input amount and the pool with updated state
 */
func (p *Pool) GetInputAmount(outputAmount *entities.CurrencyAmount, sqrtPriceLimitX96 *big.Int) (*entities.CurrencyAmount, *Pool, error) {
	if !(outputAmount.Currency.IsToken() && p.InvolvesToken(outputAmount.Currency.Wrapped())) {
		return nil, nil, ErrTokenNotInvolved
	}
	zeroForOne := outputAmount.Currency.Equal(p.Token1)
	inputAmount, sqrtRatioX96, liquidity, tickCurrent, err := p.swap(zeroForOne, new(big.Int).Mul(outputAmount.Quotient(), constants.NegativeOne), sqrtPriceLimitX96)
	if err != nil {
		return nil, nil, err
	}
	var inputToken *entities.Token
	if zeroForOne {
		inputToken = p.Token0
	} else {
		inputToken = p.Token1
	}
	pool, err := NewPool(p.Token0, p.Token1, p.Fee, sqrtRatioX96, liquidity, tickCurrent, p.TickDataProvider)
	if err != nil {
		return nil, nil, err
	}
	return entities.FromRawAmount(inputToken, inputAmount), pool, nil
}

/**
 * Executes a swap
 * @param zeroForOne Whether the amount in is token0 or token1
 * @param amountSpecified The amount of the swap, which implicitly configures the swap as exact input (positive), or exact output (negative)
 * @param sqrtPriceLimitX96 The Q64.96 sqrt price limit. If zero for one, the price cannot be less than this value after the swap. If one for zero, the price cannot be greater than this value after the swap
 * @returns amountCalculated
 * @returns sqrtRatioX96
 * @returns liquidity
 * @returns tickCurrent
 */
func (p *Pool) swap(zeroForOne bool, amountSpecified, sqrtPriceLimitX96 *big.Int) (amountCalCulated *big.Int, sqrtRatioX96 *big.Int, liquidity *big.Int, tickCurrent int, err error) {
	if sqrtPriceLimitX96 == nil {
		if zeroForOne {
			sqrtPriceLimitX96 = new(big.Int).Add(utils.MinSqrtRatio, constants.One)
		} else {
			sqrtPriceLimitX96 = new(big.Int).Sub(utils.MaxSqrtRatio, constants.One)
		}
	}

	if zeroForOne {
		if sqrtPriceLimitX96.Cmp(utils.MinSqrtRatio) <= 0 {
			return nil, nil, nil, 0, ErrSqrtPriceLimitX96TooLow
		}
		if sqrtPriceLimitX96.Cmp(p.SqrtRatioX96) >= 0 {
			return nil, nil, nil, 0, ErrSqrtPriceLimitX96TooHigh
		}
	} else {
		if sqrtPriceLimitX96.Cmp(utils.MaxSqrtRatio) >= 0 {
			return nil, nil, nil, 0, ErrSqrtPriceLimitX96TooHigh
		}
		if sqrtPriceLimitX96.Cmp(p.SqrtRatioX96) <= 0 {
			return nil, nil, nil, 0, ErrSqrtPriceLimitX96TooLow
		}
	}

	exactInput := amountSpecified.Cmp(constants.Zero) >= 0

	// keep track of swap state

	state := struct {
		amountSpecifiedRemaining *big.Int
		amountCalculated         *big.Int
		sqrtPriceX96             *big.Int
		tick                     int
		liquidity                *big.Int
	}{
		amountSpecifiedRemaining: amountSpecified,
		amountCalculated:         constants.Zero,
		sqrtPriceX96:             p.SqrtRatioX96,
		tick:                     p.TickCurrent,
		liquidity:                p.Liquidity,
	}

	// start swap while loop
	for state.amountSpecifiedRemaining.Cmp(constants.Zero) != 0 && state.sqrtPriceX96.Cmp(sqrtPriceLimitX96) != 0 {
		var step StepComputations
		step.sqrtPriceStartX96 = state.sqrtPriceX96

		// because each iteration of the while loop rounds, we can't optimize this code (relative to the smart contract)
		// by simply traversing to the next available tick, we instead need to exactly replicate
		// tickBitmap.nextInitializedTickWithinOneWord
		step.tickNext, step.initialized = p.TickDataProvider.NextInitializedTickWithinOneWord(state.tick, zeroForOne, p.tickSpacing())

		if step.tickNext < utils.MinTick {
			step.tickNext = utils.MinTick
		} else if step.tickNext > utils.MaxTick {
			step.tickNext = utils.MaxTick
		}

		step.sqrtPriceNextX96, err = utils.GetSqrtRatioAtTick(step.tickNext)
		if err != nil {
			return nil, nil, nil, 0, err
		}
		var targetValue *big.Int
		if zeroForOne {
			if step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) < 0 {
				targetValue = sqrtPriceLimitX96
			} else {
				targetValue = step.sqrtPriceNextX96
			}
		} else {
			if step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) > 0 {
				targetValue = sqrtPriceLimitX96
			} else {
				targetValue = step.sqrtPriceNextX96
			}
		}

		state.sqrtPriceX96, step.amountIn, step.amountOut, step.feeAmount, err = utils.ComputeSwapStep(state.sqrtPriceX96, targetValue, state.liquidity, state.amountSpecifiedRemaining, p.Fee)
		if err != nil {
			return nil, nil, nil, 0, err
		}

		if exactInput {
			state.amountSpecifiedRemaining = new(big.Int).Sub(state.amountSpecifiedRemaining, new(big.Int).Add(step.amountIn, step.feeAmount))
			state.amountCalculated = new(big.Int).Sub(state.amountCalculated, step.amountOut)
		} else {
			state.amountSpecifiedRemaining = new(big.Int).Add(state.amountSpecifiedRemaining, step.amountOut)
			state.amountCalculated = new(big.Int).Add(state.amountCalculated, new(big.Int).Add(step.amountIn, step.feeAmount))
		}

		// TODO
		if state.sqrtPriceX96.Cmp(step.sqrtPriceNextX96) == 0 {
			// if the tick is initialized, run the tick transition
			if step.initialized {
				liquidityNet := p.TickDataProvider.GetTick(step.tickNext).LiquidityNet
				// if we're moving leftward, we interpret liquidityNet as the opposite sign
				// safe because liquidityNet cannot be type(int128).min
				if zeroForOne {
					liquidityNet = new(big.Int).Mul(liquidityNet, constants.NegativeOne)
				}
				state.liquidity = utils.AddDelta(state.liquidity, liquidityNet)
			}
			if zeroForOne {
				state.tick = step.tickNext - 1
			} else {
				state.tick = step.tickNext
			}
		} else if state.sqrtPriceX96.Cmp(step.sqrtPriceStartX96) != 0 {
			// recompute unless we're on a lower tick boundary (i.e. already transitioned ticks), and haven't moved
			state.tick, err = utils.GetTickAtSqrtRatio(state.sqrtPriceX96)
			if err != nil {
				return nil, nil, nil, 0, err
			}
		}
	}
	return state.amountCalculated, state.sqrtPriceX96, state.liquidity, state.tick, nil
}

func (p *Pool) tickSpacing() int {
	return constants.TickSpacings[p.Fee]
}
