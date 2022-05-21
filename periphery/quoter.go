package periphery

import (
	_ "embed"
	"errors"
	"fmt"
	"math/big"
	"reflect"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

//go:embed contracts/lens/Quoter.sol/Quoter.json
var quoterABI []byte

var ErrMultihopPriceLimit = errors.New("MULTIHOP_PRICE_LIMIT")

// Optional arguments to send to the quoter.
type QuoteOptions struct {
	SqrtPriceLimitX96 *big.Int // The optional price limit for the trade.
}

/**
 * Represents the Uniswap V3 QuoterV1 contract with a method for returning the formatted
 * calldata needed to call the quoter contract.
 */

/**
 * Produces the on-chain method name of the appropriate function within QuoterV2,
 * and the relevant hex encoded parameters.
 * @template TInput The input token, either Ether or an ERC-20
 * @template TOutput The output token, either Ether or an ERC-20
 * @param route The swap route, a list of pools through which a swap can occur
 * @param amount The amount of the quote, either an amount in, or an amount out
 * @param tradeType The trade type, either exact input or exact output
 * @returns The formatted calldata
 */
func QuoteCallParameters(
	route *entities.Route,
	amount *core.CurrencyAmount,
	tradeType core.TradeType,
	options *QuoteOptions,
) (*utils.MethodParameters, error) {
	singleHop := len(route.Pools) == 1
	quoteAmount := amount.Quotient()
	abi := GetABI(quoterABI)
	var (
		calldata []byte
		err      error
	)
	sqrtPriceLimitX96 := big.NewInt(0)
	if options != nil {
		sqrtPriceLimitX96 = options.SqrtPriceLimitX96
	}

	if singleHop {
		if tradeType == core.ExactInput {
			calldata, err = abi.Pack("quoteExactInputSingle", route.TokenPath[0].Address, route.TokenPath[1].Address, big.NewInt(int64(route.Pools[0].Fee)), quoteAmount, sqrtPriceLimitX96)
		} else {
			calldata, err = abi.Pack("quoteExactOutputSingle", route.TokenPath[0].Address, route.TokenPath[1].Address, big.NewInt(int64(route.Pools[0].Fee)), quoteAmount, sqrtPriceLimitX96)
		}
		if err != nil {
			return nil, err
		}
	} else {
		if options != nil && sqrtPriceLimitX96 != big.NewInt(0) {
			return nil, ErrMultihopPriceLimit
		}
		path, err := EncodeRouteToPath(route, tradeType == core.ExactOutput)
		if err != nil {
			return nil, err
		}
		if tradeType == core.ExactInput {
			calldata, err = abi.Pack("quoteExactInput", path, quoteAmount)
		} else {
			calldata, err = abi.Pack("quoteExactOutput", path, quoteAmount)
		}
		if err != nil {
			return nil, err
		}
	}
	return &utils.MethodParameters{
		Calldata: calldata,
		Value:    big.NewInt(0),
	}, nil
}

/**
 * Converts a route to a hex encoded path
 * @param route the v3 path to convert to an encoded path
 * @param exactOutput whether the route should be encoded in reverse, for making exact output swaps
 */
func EncodeRouteToPath(route *entities.Route, exactOutput bool) ([]byte, error) {
	var (
		inputToken = route.Input.Wrapped()

		types []string
		path  []interface{}

		addressTy = "address"
		uint24Ty  = "uint24"
	)

	for i, pool := range route.Pools {
		var outputToken *core.Token
		if pool.Token0.Equal(inputToken) {
			outputToken = pool.Token1
		} else {
			outputToken = pool.Token0
		}
		if i == 0 {
			types = []string{addressTy, uint24Ty, addressTy}
			path = []interface{}{inputToken.Address, uint64(pool.Fee), outputToken.Address}
		} else {
			types = append(types, uint24Ty, addressTy)
			path = append(path, uint64(pool.Fee), outputToken.Address)
		}
		inputToken = outputToken
	}

	if exactOutput {
		reverse(types)
		reverse(path)
	}

	// tight pack the path
	var packedPath [][]byte
	for i, t := range types {
		switch t {
		case addressTy:
			packedPath = append(packedPath, path[i].(common.Address).Bytes())
		case uint24Ty:
			packedPath = append(packedPath, common.LeftPadBytes(PutUint24(path[i].(uint64)), 24/8))
		default:
			return nil, fmt.Errorf("unknown type %s", t)
		}

	}

	var pack []byte
	for _, p := range packedPath {
		pack = append(pack, p...)
	}
	return pack, nil
}

// PutUint24 put bigendian uint24
func PutUint24(i uint64) []byte {
	b := make([]byte, 3)
	b[0] = byte(i >> 16)
	b[1] = byte(i >> 8)
	b[2] = byte(i)
	return b
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
