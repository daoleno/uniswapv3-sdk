package entities

import (
	"errors"
	"math/big"
	"sort"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrTradeHasMultipleRoutes   = errors.New("trade has multiple routes")
	ErrInvalidAmountForRoute    = errors.New("invalid amount for route")
	ErrInputCurrencyMismatch    = errors.New("input currency mismatch")
	ErrOutputCurrencyMismatch   = errors.New("output currency mismatch")
	ErrDuplicatePools           = errors.New("duplicate pools")
	ErrInvalidSlippageTolerance = errors.New("invalid slippage tolerance")
	ErrNoPools                  = errors.New("no pools")
	ErrInvalidMaxHops           = errors.New("invalid max hops")
	ErrInvalidRecursion         = errors.New("invalid recursion")
	ErrInvalidMaxSize           = errors.New("invalid max size")
	ErrMaxSizeExceeded          = errors.New("max size exceeded")
)

/**
 * Trades comparator, an extension of the input output comparator that also considers other dimensions of the trade in ranking them
 * @template TInput The input token, either Ether or an ERC-20
 * @template TOutput The output token, either Ether or an ERC-20
 * @template TTradeType The trade type, either exact input or exact output
 * @param a The first trade to compare
 * @param b The second trade to compare
 * @returns A sorted ordering for two neighboring elements in a trade array
 */
func tradeComparator(a, b *Trade) int {
	if !a.InputAmount().Currency.Equal(b.InputAmount().Currency) {
		panic(ErrInputCurrencyMismatch)
	}
	if !a.OutputAmount().Currency.Equal(b.OutputAmount().Currency) {
		panic(ErrOutputCurrencyMismatch)
	}
	if a.OutputAmount().EqualTo(b.OutputAmount().Fraction) {
		if a.InputAmount().EqualTo(b.InputAmount().Fraction) {
			// consider the number of hops since each hop costs gas
			var aHops, bHops int
			for _, swap := range a.Swaps {
				aHops += len(swap.Route.TokenPath)
			}
			for _, swap := range b.Swaps {
				bHops += len(swap.Route.TokenPath)
			}
			return aHops - bHops
		}
		// trade A requires less input than trade B, so A should come first
		if a.InputAmount().LessThan(b.InputAmount().Fraction) {
			return -1
		} else {
			return 1
		}
	} else {
		// tradeA has less output than trade B, so should come second
		if a.OutputAmount().LessThan(b.OutputAmount().Fraction) {
			return 1
		} else {
			return -1
		}
	}
}

/**
 * Represents a trade executed against a set of routes where some percentage of the input is
 * split across each route.
 *
 * Each route has its own set of pools. Pools can not be re-used across routes.
 *
 * Does not account for slippage, i.e., changes in price environment that can occur between
 * the time the trade is submitted and when it is executed.
 */
type Trade struct {
	Swaps     []*Swap            // The swaps of the trade, i.e. which routes and how much is swapped in each that make up the trade.
	TradeType entities.TradeType // The type of trade, i.e. exact input or exact output

	inputAmount    *entities.CurrencyAmount // The cached result of the input amount computation
	outputAmount   *entities.CurrencyAmount // The cached result of the output amount computation
	executionPrice *entities.Price          // The cached result of the computed execution price
	priceImpact    *entities.Percent        // The cached result of the price impact computation
}

type Swap struct {
	Route        *Route
	InputAmount  *entities.CurrencyAmount
	OutputAmount *entities.CurrencyAmount
}

/**
 * @deprecated Deprecated in favor of 'swaps' property. If the trade consists of multiple routes
 * this will return an error.
 *
 * When the trade consists of just a single route, this returns the route of the trade,
 * i.e. which pools the trade goes through.
 */
func (t *Trade) Route() (*Route, error) {
	if len(t.Swaps) != 1 {
		return nil, ErrTradeHasMultipleRoutes
	}
	return t.Swaps[0].Route, nil
}

// InputAmount the input amount for the trade assuming no slippage.
func (t *Trade) InputAmount() *entities.CurrencyAmount {
	if t.inputAmount != nil {
		return t.inputAmount
	}
	inputCurrency := t.Swaps[0].InputAmount.Currency
	totalInputFromRoutes := entities.FromRawAmount(inputCurrency, big.NewInt(0))
	for _, swap := range t.Swaps {
		totalInputFromRoutes = totalInputFromRoutes.Add(swap.InputAmount)
	}
	t.inputAmount = totalInputFromRoutes
	return t.inputAmount
}

// OutputAmount the output amount for the trade assuming no slippage.
func (t *Trade) OutputAmount() *entities.CurrencyAmount {
	if t.outputAmount != nil {
		return t.outputAmount
	}
	outputCurrency := t.Swaps[0].OutputAmount.Currency
	totalOutputFromRoutes := entities.FromRawAmount(outputCurrency, big.NewInt(0))
	for _, swap := range t.Swaps {
		totalOutputFromRoutes = totalOutputFromRoutes.Add(swap.OutputAmount)
	}
	t.outputAmount = totalOutputFromRoutes
	return t.outputAmount
}

// ExecutionPrice the price expressed in terms of output amount/input amount.
func (t *Trade) ExecutionPrice() *entities.Price {
	if t.executionPrice != nil {
		return t.executionPrice
	}

	t.executionPrice = entities.NewPrice(t.inputAmount.Currency, t.outputAmount.Currency, t.inputAmount.Quotient(), t.outputAmount.Quotient())
	return t.executionPrice
}

// PriceImpact returns the percent difference between the route's mid price and the price impact
func (t *Trade) PriceImpact() (*entities.Percent, error) {
	if t.priceImpact != nil {
		return t.priceImpact, nil
	}

	spotOutputAmount := entities.FromRawAmount(t.OutputAmount().Currency, big.NewInt(0))
	for _, swap := range t.Swaps {
		midPrice, err := swap.Route.MidPrice()
		if err != nil {
			return nil, err
		}
		quotePrice, err := midPrice.Quote(swap.InputAmount)
		if err != nil {
			return nil, err
		}
		spotOutputAmount = spotOutputAmount.Add(quotePrice)
	}

	priceImpact := spotOutputAmount.Subtract(t.OutputAmount()).Divide(spotOutputAmount.Fraction)
	t.priceImpact = entities.NewPercent(priceImpact.Numerator, priceImpact.Denominator)
	return t.priceImpact, nil
}

/**
 * Constructs an exact in trade with the given amount in and route
 * @template TInput The input token, either Ether or an ERC-20
 * @template TOutput The output token, either Ether or an ERC-20
 * @param route The route of the exact in trade
 * @param amountIn The amount being passed in
 * @returns The exact in trade
 */
func ExactIn(route *Route, amountIn *entities.CurrencyAmount) (*Trade, error) {
	return FromRoute(route, amountIn, entities.ExactInput)
}

/**
 * Constructs an exact out trade with the given amount out and route
 * @template TInput The input token, either Ether or an ERC-20
 * @template TOutput The output token, either Ether or an ERC-20
 * @param route The route of the exact out trade
 * @param amountOut The amount returned by the trade
 * @returns The exact out trade
 */
func ExactOut(route *Route, amountOut *entities.CurrencyAmount) (*Trade, error) {
	return FromRoute(route, amountOut, entities.ExactOutput)
}

/**
 * Constructs a trade by simulating swaps through the given route
 * @template TInput The input token, either Ether or an ERC-20.
 * @template TOutput The output token, either Ether or an ERC-20.
 * @template TTradeType The type of the trade, either exact in or exact out.
 * @param route route to swap through
 * @param amount the amount specified, either input or output, depending on tradeType
 * @param tradeType whether the trade is an exact input or exact output swap
 * @returns The route
 */
func FromRoute(route *Route, amount *entities.CurrencyAmount, tradeType entities.TradeType) (*Trade, error) {
	amounts := make([]*entities.CurrencyAmount, len(route.TokenPath))
	var (
		inputAmount  *entities.CurrencyAmount
		outputAmount *entities.CurrencyAmount
		err          error
	)
	if tradeType == entities.ExactInput {
		if !amount.Currency.Equal(route.Input.Currency) {
			return nil, ErrInvalidAmountForRoute
		}
		amounts[0] = amount
		for i := 0; i < len(route.TokenPath)-1; i++ {
			pool := route.Pools[i]
			outputAmount, _, err = pool.GetOutputAmount(amounts[i], nil)
			if err != nil {
				return nil, err
			}
			amounts[i+1] = outputAmount
		}
		inputAmount = entities.FromFractionalAmount(route.Input.Currency, amount.Numerator, amount.Denominator)
		outputAmount = entities.FromFractionalAmount(route.Output.Currency, amounts[len(amounts)-1].Numerator, amounts[len(amounts)-1].Denominator)
	} else {
		if !amount.Currency.Equal(route.Output.Currency) {
			return nil, ErrInvalidAmountForRoute
		}
		amounts[len(amounts)-1] = amount
		for i := len(route.TokenPath) - 1; i > 0; i-- {
			pool := route.Pools[i-1]
			inputAmount, _, err = pool.GetInputAmount(amounts[i], nil)
			if err != nil {
				return nil, err
			}
			amounts[i-1] = inputAmount
		}
		inputAmount = entities.FromFractionalAmount(route.Input.Currency, amounts[0].Numerator, amounts[0].Denominator)
		outputAmount = entities.FromFractionalAmount(route.Output.Currency, amount.Numerator, amount.Denominator)
	}
	swaps := []*Swap{{
		Route:        route,
		InputAmount:  inputAmount,
		OutputAmount: outputAmount}}

	return newTrade(swaps, tradeType)
}

type WrappedRoute struct {
	Amount *entities.CurrencyAmount
	Route  *Route
}

/**
 * Constructs a trade from routes by simulating swaps
 *
 * @template TInput The input token, either Ether or an ERC-20.
 * @template TOutput The output token, either Ether or an ERC-20.
 * @template TTradeType The type of the trade, either exact in or exact out.
 * @param routes the routes to swap through and how much of the amount should be routed through each
 * @param tradeType whether the trade is an exact input or exact output swap
 * @returns The trade
 */
func FromRoutes(wrappedRoutes []*WrappedRoute, tradeType entities.TradeType) (*Trade, error) {
	var swaps []*Swap
	for _, wrappedRoute := range wrappedRoutes {
		amounts := make([]*entities.CurrencyAmount, len(wrappedRoute.Route.TokenPath))
		var (
			inputAmount  *entities.CurrencyAmount
			outputAmount *entities.CurrencyAmount
		)
		amount := wrappedRoute.Amount
		route := wrappedRoute.Route
		if tradeType == entities.ExactInput {
			if !amount.Currency.Equal(route.Input.Currency) {
				return nil, ErrInvalidAmountForRoute
			}
			amounts[0] = entities.FromFractionalAmount(route.Input.Currency, amount.Numerator, amount.Denominator)
			for i := 0; i < len(route.TokenPath)-1; i++ {
				pool := route.Pools[i]
				outputAmount, _, err := pool.GetOutputAmount(amounts[i], nil)
				if err != nil {
					return nil, err
				}
				amounts[i+1] = outputAmount
			}
			inputAmount = entities.FromFractionalAmount(route.Input.Currency, amount.Numerator, amount.Denominator)
			outputAmount = entities.FromFractionalAmount(route.Output.Currency, amounts[len(amounts)-1].Numerator, amounts[len(amounts)-1].Denominator)
		} else {
			if !amount.Currency.Equal(route.Output.Currency) {
				return nil, ErrInvalidAmountForRoute
			}
			amounts[len(amounts)-1] = entities.FromFractionalAmount(route.Output.Currency, amount.Numerator, amount.Denominator)
			for i := len(route.TokenPath) - 1; i > 0; i-- {
				pool := route.Pools[i-1]
				inputAmount, _, err := pool.GetInputAmount(amounts[i], nil)
				if err != nil {
					return nil, err
				}
				amounts[i-1] = inputAmount
			}
			inputAmount = entities.FromFractionalAmount(route.Input.Currency, amounts[0].Numerator, amounts[0].Denominator)
			outputAmount = entities.FromFractionalAmount(route.Output.Currency, amount.Numerator, amount.Denominator)
		}
		swaps = append(swaps, &Swap{
			Route:        route,
			InputAmount:  inputAmount,
			OutputAmount: outputAmount})

	}
	return newTrade(swaps, tradeType)
}

/**
 * Creates a trade without computing the result of swapping through the route. Useful when you have simulated the trade
 * elsewhere and do not have any tick data
 * @template TInput The input token, either Ether or an ERC-20
 * @template TOutput The output token, either Ether or an ERC-20
 * @template TTradeType The type of the trade, either exact in or exact out
 * @param constructorArguments The arguments passed to the trade constructor
 * @returns The unchecked trade
 */
func CreateUncheckedTrade(route *Route, inputAmount, outputAmount *entities.CurrencyAmount, tradeType entities.TradeType) (*Trade, error) {
	swaps := []*Swap{{
		Route:        route,
		InputAmount:  inputAmount,
		OutputAmount: outputAmount}}
	return newTrade(swaps, tradeType)
}

/**
 * Creates a trade without computing the result of swapping through the routes. Useful when you have simulated the trade
 * elsewhere and do not have any tick data
 * @template TInput The input token, either Ether or an ERC-20
 * @template TOutput The output token, either Ether or an ERC-20
 * @template TTradeType The type of the trade, either exact in or exact out
 * @param constructorArguments The arguments passed to the trade constructor
 * @returns The unchecked trade
 */
func CreateUncheckedTradeWithMultipleRoutes(routes []*Swap, tradeType entities.TradeType) (*Trade, error) {
	return newTrade(routes, tradeType)
}

/**
 * Construct a trade by passing in the pre-computed property values
 * @param routes The routes through which the trade occurs
 * @param tradeType The type of trade, exact input or exact output
 */
func newTrade(routes []*Swap, tradeType entities.TradeType) (*Trade, error) {
	inputCurrency := routes[0].InputAmount.Currency
	outputCurrency := routes[0].OutputAmount.Currency
	for _, route := range routes {
		if !inputCurrency.Equal(route.Route.Input.Currency) {
			return nil, ErrInputCurrencyMismatch
		}
		if !outputCurrency.Equal(route.Route.Output.Currency) {
			return nil, ErrOutputCurrencyMismatch
		}
	}

	var numPools int
	for _, route := range routes {
		numPools += len(route.Route.Pools)
	}

	var poolAddressSet = make(map[common.Address]bool)
	for _, route := range routes {
		for _, pool := range route.Route.Pools {
			addr, err := GetAddress(pool.Token0, pool.Token1, pool.Fee, "")
			if err != nil {
				return nil, err
			}
			poolAddressSet[addr] = true
		}
	}

	if numPools != len(poolAddressSet) {
		return nil, ErrDuplicatePools
	}

	return &Trade{
		Swaps:     routes,
		TradeType: tradeType,
	}, nil
}

/**
 * Get the minimum amount that must be received from this trade for the given slippage tolerance
 * @param slippageTolerance The tolerance of unfavorable slippage from the execution price of this trade
 * @returns The amount out
 */
func (t *Trade) MinimumAmountOut(slippageTolerance *entities.Percent) (*entities.CurrencyAmount, error) {
	if slippageTolerance.LessThan(constants.PercentZero) {
		return nil, ErrInvalidSlippageTolerance
	}
	if t.TradeType == entities.ExactOutput {
		return t.OutputAmount(), nil
	} else {
		slippageAdjustedAmountOut := entities.NewFraction(big.NewInt(1), big.NewInt(1)).
			Add(slippageTolerance.Fraction).
			Invert().
			Multiply(t.OutputAmount().Fraction).Quotient()
		return entities.FromRawAmount(t.OutputAmount().Currency, slippageAdjustedAmountOut), nil
	}
}

/**
 * Get the maximum amount in that can be spent via this trade for the given slippage tolerance
 * @param slippageTolerance The tolerance of unfavorable slippage from the execution price of this trade
 * @returns The amount in
 */
func (t *Trade) MaximumAmountIn(slippageTolerance *entities.Percent) (*entities.CurrencyAmount, error) {
	if slippageTolerance.LessThan(constants.PercentZero) {
		return nil, ErrInvalidSlippageTolerance
	}
	if t.TradeType == entities.ExactInput {
		return t.InputAmount(), nil
	} else {
		slippageAdjustedAmountIn := entities.NewFraction(big.NewInt(1), big.NewInt(1)).
			Add(slippageTolerance.Fraction).
			Multiply(t.InputAmount().Fraction).Quotient()
		return entities.FromRawAmount(t.InputAmount().Currency, slippageAdjustedAmountIn), nil
	}
}

/**
 * Return the execution price after accounting for slippage tolerance
 * @param slippageTolerance the allowed tolerated slippage
 * @returns The execution price
 */
func (t *Trade) WorstExecutionPrice(slippageTolerance *entities.Percent) (*entities.Price, error) {
	maxAmountIn, err := t.MaximumAmountIn(slippageTolerance)
	if err != nil {
		return nil, err
	}
	minAmountOut, err := t.MinimumAmountOut(slippageTolerance)
	if err != nil {
		return nil, err
	}
	return entities.NewPrice(t.InputAmount().Currency, t.OutputAmount().Currency, maxAmountIn.Quotient(), minAmountOut.Quotient()), nil
}

type BestTradeOptions struct {
	MaxNumResults int // how many results to return
	MaxHops       int // the maximum number of hops a trade should contain
}

/**
 * Given a list of pools, and a fixed amount in, returns the top `maxNumResults` trades that go from an input token
 * amount to an output token, making at most `maxHops` hops.
 * Note this does not consider aggregation, as routes are linear. It's possible a better route exists by splitting
 * the amount in among multiple routes.
 * @param pools the pools to consider in finding the best trade
 * @param nextAmountIn exact amount of input currency to spend
 * @param currencyOut the desired currency out
 * @param maxNumResults maximum number of results to return
 * @param maxHops maximum number of hops a returned trade can make, e.g. 1 hop goes through a single pool
 * @param currentPools used in recursion; the current list of pools
 * @param currencyAmountIn used in recursion; the original value of the currencyAmountIn parameter
 * @param bestTrades used in recursion; the current list of best trades
 * @returns The exact in trade
 */
//  TODO: Merge Token and CurrencyAmount
func BestTradeExactIn(pools []*Pool, currencyAmountIn *entities.CurrencyAmount, tokenIn *entities.Token, tokenOut *entities.Token, opts *BestTradeOptions, currentPools []*Pool, nextAmountIn *entities.CurrencyAmount, bestTrades []*Trade) ([]*Trade, error) {
	if len(pools) <= 0 {
		return nil, ErrNoPools
	}
	if opts == nil {
		opts = &BestTradeOptions{MaxNumResults: 3, MaxHops: 3}
	}

	if nextAmountIn == nil {
		nextAmountIn = currencyAmountIn
	}
	if opts.MaxHops <= 0 {
		return nil, ErrInvalidMaxHops
	}
	if !(currencyAmountIn.EqualTo(nextAmountIn.Fraction) || len(currentPools) > 0) {
		return nil, ErrInvalidRecursion
	}

	amountIn := nextAmountIn
	for i := 0; i < len(pools); i++ {
		pool := pools[i]
		//  pool irrelevant
		if !pool.Token0.Equal(amountIn.Currency) && !pool.Token1.Equal(amountIn.Currency) {
			continue
		}
		amountOut, _, err := pool.GetOutputAmount(amountIn, nil)
		if err != nil {
			// TODO
			// input too low
			//  if (error.isInsufficientInputAmountError) {
			// 	continue
			//       }
			return nil, err
		}
		// we have arrived at the output token, so this is the final trade of one of the paths
		if amountOut.Currency.IsToken && amountOut.Currency.Equal(tokenOut.Currency) {
			r, err := NewRoute(append(currentPools, pool), tokenIn, tokenOut)
			if err != nil {
				return nil, err
			}
			trade, err := FromRoute(r, currencyAmountIn, entities.ExactInput)
			if err != nil {
				return nil, err
			}
			bestTrades, err = sortedInsert(bestTrades, trade, opts.MaxNumResults, tradeComparator)
			if err != nil {
				return nil, err
			}
		} else if opts.MaxHops > 1 && len(pools) > 1 {
			var poolsExcludingThisPool []*Pool
			poolsExcludingThisPool = append(poolsExcludingThisPool, pools[:i]...)
			poolsExcludingThisPool = append(poolsExcludingThisPool, pools[i+1:]...)

			// otherwise, consider all the other paths that lead from this token as long as we have not exceeded maxHops
			bestTrades, err = BestTradeExactIn(poolsExcludingThisPool, currencyAmountIn, tokenIn, tokenOut, &BestTradeOptions{MaxNumResults: opts.MaxNumResults, MaxHops: opts.MaxHops - 1}, append(currentPools, pool), amountOut, bestTrades)
			if err != nil {
				return nil, err
			}
		}
	}
	return bestTrades, nil
}

/**
 * similar to the above method but instead targets a fixed output amount
 * given a list of pools, and a fixed amount out, returns the top `maxNumResults` trades that go from an input token
 * to an output token amount, making at most `maxHops` hops
 * note this does not consider aggregation, as routes are linear. it's possible a better route exists by splitting
 * the amount in among multiple routes.
 * @param pools the pools to consider in finding the best trade
 * @param currencyIn the currency to spend
 * @param currencyAmountOut the desired currency amount out
 * @param nextAmountOut the exact amount of currency out
 * @param maxNumResults maximum number of results to return
 * @param maxHops maximum number of hops a returned trade can make, e.g. 1 hop goes through a single pool
 * @param currentPools used in recursion; the current list of pools
 * @param bestTrades used in recursion; the current list of best trades
 * @returns The exact out trade
 */
func BestTradeExactOut(pools []*Pool, tokenIn *entities.Token, currencyAmountOut *entities.CurrencyAmount, tokenOut *entities.Token, opts *BestTradeOptions, currentPools []*Pool, nextAmountOut *entities.CurrencyAmount, bestTrades []*Trade) ([]*Trade, error) {
	if len(pools) <= 0 {
		return nil, ErrNoPools
	}
	if opts == nil {
		opts = &BestTradeOptions{MaxNumResults: 3, MaxHops: 3}
	}

	if nextAmountOut == nil {
		nextAmountOut = currencyAmountOut
	}
	if opts.MaxHops <= 0 {
		return nil, ErrInvalidMaxHops
	}
	if !(currencyAmountOut.EqualTo(nextAmountOut.Fraction) || len(currentPools) > 0) {
		return nil, ErrInvalidRecursion
	}

	amountOut := nextAmountOut

	for i := 0; i < len(pools); i++ {
		pool := pools[i]
		// pool irrelevant
		if !pool.Token0.Equal(amountOut.Currency) && !pool.Token1.Equal(amountOut.Currency) {
			continue
		}
		amountIn, _, err := pool.GetInputAmount(amountOut, nil)
		if err != nil {
			// TODO
			// not enough liquidity in this pool
			// if error.isInsufficientReservesError {
			// 	continue
			// }
			return nil, err
		}
		// we have arrived at the input token, so this is the final trade of one of the paths
		if amountIn.Currency.Equal(tokenIn.Currency) {
			r, err := NewRoute(append([]*Pool{pool}, currentPools...), tokenIn, tokenOut)
			if err != nil {
				return nil, err
			}
			trade, err := FromRoute(r, currencyAmountOut, entities.ExactOutput)
			if err != nil {
				return nil, err
			}
			bestTrades, err = sortedInsert(bestTrades, trade, opts.MaxNumResults, tradeComparator)
			if err != nil {
				return nil, err
			}
		} else if opts.MaxHops > 1 && len(pools) > 1 {
			var poolsExcludingThisPool []*Pool
			poolsExcludingThisPool = append(poolsExcludingThisPool, pools[:i]...)
			poolsExcludingThisPool = append(poolsExcludingThisPool, pools[i+1:]...)

			// otherwise, consider all the other paths that arrive at this token as long as we have not exceeded maxHops
			bestTrades, err = BestTradeExactOut(poolsExcludingThisPool, tokenIn, currencyAmountOut, tokenOut, &BestTradeOptions{MaxNumResults: opts.MaxNumResults, MaxHops: opts.MaxHops - 1}, append([]*Pool{pool}, currentPools...), amountIn, bestTrades)
			if err != nil {
				return nil, err
			}
		}
	}
	return bestTrades, nil
}

// sortedInsert given an array of items sorted by `comparator`, insert an item into its sort index and constrain the size to
// `maxSize` by removing the last item
func sortedInsert(items []*Trade, add *Trade, maxSize int, comparator func(a, b *Trade) int) ([]*Trade, error) {
	if maxSize <= 0 {
		return nil, ErrInvalidMaxSize
	}
	if len(items) > maxSize {
		return nil, ErrMaxSizeExceeded
	}

	isFull := len(items) == maxSize

	if isFull && comparator(items[maxSize-1], add) <= 0 {
		return []*Trade{add}, nil
	}

	i := sort.Search(len(items), func(i int) bool {
		return comparator(items[i], add) > 0
	})
	items = append(items, nil)
	copy(items[i+1:], items[i:])
	items[i] = add
	if isFull {
		return items[:maxSize], nil
	}
	return items, nil
}
