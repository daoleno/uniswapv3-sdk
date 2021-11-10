package entities

import (
	"errors"

	"github.com/daoleno/uniswap-sdk-core/entities"
)

var (
	ErrRouteNoPools      = errors.New("route must have at least one pool")
	ErrAllOnSameChain    = errors.New("all pools must be on the same chain")
	ErrInputNotInvolved  = errors.New("input token not involved in route")
	ErrOutputNotInvolved = errors.New("output token not involved in route")
	ErrPathNotContinuous = errors.New("path not continuous")
)

// Route represents a list of pools through which a swap can occur
type Route struct {
	Pools     []*Pool
	TokenPath []*entities.Token
	Input     *entities.Token
	Output    *entities.Token

	midPrice *entities.Price
}

/**
 * Creates an instance of route.
 * @param pools An array of `Pool` objects, ordered by the route the swap will take
 * @param input The input token
 * @param output The output token
 */
func NewRoute(pools []*Pool, input, output *entities.Token) (*Route, error) {
	if len(pools) == 0 {
		return nil, ErrRouteNoPools
	}
	chainID := pools[0].ChainID()
	for _, p := range pools {
		if p.ChainID() != chainID {
			return nil, ErrAllOnSameChain
		}
	}

	if !pools[0].InvolvesToken(input) {
		return nil, ErrInputNotInvolved
	}
	if !pools[len(pools)-1].InvolvesToken(output) {
		return nil, ErrOutputNotInvolved
	}

	// Normalizes token0-token1 order and selects the next token/fee step to add to the path
	tokenPath := []*entities.Token{input}
	for i, p := range pools {
		currentInputToken := tokenPath[i]
		if !(currentInputToken.Equals(p.Token0) || currentInputToken.Equals(p.Token1)) {
			return nil, ErrPathNotContinuous
		}
		var nextToken *entities.Token
		if currentInputToken.Equals(p.Token0) {
			nextToken = p.Token1
		} else {
			nextToken = p.Token0
		}
		tokenPath = append(tokenPath, nextToken)
	}

	if output == nil {
		output = tokenPath[len(tokenPath)-1]
	}
	return &Route{
		Pools:     pools,
		TokenPath: tokenPath,
		Input:     input,
		Output:    output,
	}, nil
}

func (r *Route) ChainID() uint {
	return r.Pools[0].ChainID()
}

// MidPrice Returns the mid price of the route
func (r *Route) MidPrice() (*entities.Price, error) {
	if r.midPrice != nil {
		return r.midPrice, nil
	}
	var (
		nextInput *entities.Token
		price     *entities.Price
	)
	if r.Pools[0].Token0.Equals(r.Input) {
		nextInput = r.Pools[0].Token1
		price = r.Pools[0].Token0Price()
	} else {
		nextInput = r.Pools[0].Token0
		price = r.Pools[0].Token1Price()
	}
	price, err := reducePrice(nextInput, price, r.Pools[1:])
	if err != nil {
		return nil, err
	}
	r.midPrice = entities.NewPrice(r.Input.Currency, r.Output.Currency, price.Denominator, price.Numerator)
	return r.midPrice, nil
}

// reducePrice reduces the price of the route by the given amount
func reducePrice(nextInput *entities.Token, price *entities.Price, pools []*Pool) (*entities.Price, error) {
	var err error
	for _, p := range pools {
		if nextInput.Equals(p.Token0) {
			nextInput = p.Token1
			price, err = price.Multiply(p.Token0Price())
			if err != nil {
				return nil, err
			}
		} else {
			nextInput = p.Token0
			price, err = price.Multiply(p.Token1Price())
			if err != nil {
				return nil, err
			}
		}
	}
	return price, nil
}
