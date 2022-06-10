// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IApproveAndCallIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallIncreaseLiquidityParams struct {
	Token0     common.Address
	Token1     common.Address
	TokenId    *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
}

// IApproveAndCallMintParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallMintParams struct {
	Token0     common.Address
	Token1     common.Address
	Fee        *big.Int
	TickLower  *big.Int
	TickUpper  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Recipient  common.Address
}

// IV3SwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IV3SwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IV3SwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IV3SwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// Uniswapv3RouterV2ABI is the input ABI used to generate the binding from.
const Uniswapv3RouterV2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factoryV3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callPositionManager\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"paths\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getApprovalType\",\"outputs\":[{\"internalType\":\"enumIApproveAndCall.ApprovalType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"}],\"internalType\":\"structIApproveAndCall.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIApproveAndCall.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"previousBlockhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"wrapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Uniswapv3RouterV2 is an auto generated Go binding around an Ethereum contract.
type Uniswapv3RouterV2 struct {
	Uniswapv3RouterV2Caller     // Read-only binding to the contract
	Uniswapv3RouterV2Transactor // Write-only binding to the contract
	Uniswapv3RouterV2Filterer   // Log filterer for contract events
}

// Uniswapv3RouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3RouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3RouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3RouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3RouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3RouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3RouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3RouterV2Session struct {
	Contract     *Uniswapv3RouterV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Uniswapv3RouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3RouterV2CallerSession struct {
	Contract *Uniswapv3RouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Uniswapv3RouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3RouterV2TransactorSession struct {
	Contract     *Uniswapv3RouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Uniswapv3RouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3RouterV2Raw struct {
	Contract *Uniswapv3RouterV2 // Generic contract binding to access the raw methods on
}

// Uniswapv3RouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3RouterV2CallerRaw struct {
	Contract *Uniswapv3RouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3RouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3RouterV2TransactorRaw struct {
	Contract *Uniswapv3RouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3RouterV2 creates a new instance of Uniswapv3RouterV2, bound to a specific deployed contract.
func NewUniswapv3RouterV2(address common.Address, backend bind.ContractBackend) (*Uniswapv3RouterV2, error) {
	contract, err := bindUniswapv3RouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3RouterV2{Uniswapv3RouterV2Caller: Uniswapv3RouterV2Caller{contract: contract}, Uniswapv3RouterV2Transactor: Uniswapv3RouterV2Transactor{contract: contract}, Uniswapv3RouterV2Filterer: Uniswapv3RouterV2Filterer{contract: contract}}, nil
}

// NewUniswapv3RouterV2Caller creates a new read-only instance of Uniswapv3RouterV2, bound to a specific deployed contract.
func NewUniswapv3RouterV2Caller(address common.Address, caller bind.ContractCaller) (*Uniswapv3RouterV2Caller, error) {
	contract, err := bindUniswapv3RouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3RouterV2Caller{contract: contract}, nil
}

// NewUniswapv3RouterV2Transactor creates a new write-only instance of Uniswapv3RouterV2, bound to a specific deployed contract.
func NewUniswapv3RouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3RouterV2Transactor, error) {
	contract, err := bindUniswapv3RouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3RouterV2Transactor{contract: contract}, nil
}

// NewUniswapv3RouterV2Filterer creates a new log filterer instance of Uniswapv3RouterV2, bound to a specific deployed contract.
func NewUniswapv3RouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3RouterV2Filterer, error) {
	contract, err := bindUniswapv3RouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3RouterV2Filterer{contract: contract}, nil
}

// bindUniswapv3RouterV2 binds a generic wrapper to an already deployed contract.
func bindUniswapv3RouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswapv3RouterV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3RouterV2.Contract.Uniswapv3RouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Uniswapv3RouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Uniswapv3RouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3RouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Caller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3RouterV2.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) WETH9() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.WETH9(&_Uniswapv3RouterV2.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerSession) WETH9() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.WETH9(&_Uniswapv3RouterV2.CallOpts)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Caller) CheckOracleSlippage(opts *bind.CallOpts, paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _Uniswapv3RouterV2.contract.Call(opts, &out, "checkOracleSlippage", paths, amounts, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _Uniswapv3RouterV2.Contract.CheckOracleSlippage(&_Uniswapv3RouterV2.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerSession) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _Uniswapv3RouterV2.Contract.CheckOracleSlippage(&_Uniswapv3RouterV2.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Caller) CheckOracleSlippage0(opts *bind.CallOpts, path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _Uniswapv3RouterV2.contract.Call(opts, &out, "checkOracleSlippage0", path, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _Uniswapv3RouterV2.Contract.CheckOracleSlippage0(&_Uniswapv3RouterV2.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerSession) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _Uniswapv3RouterV2.Contract.CheckOracleSlippage0(&_Uniswapv3RouterV2.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3RouterV2.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Factory() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.Factory(&_Uniswapv3RouterV2.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerSession) Factory() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.Factory(&_Uniswapv3RouterV2.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Caller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3RouterV2.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) FactoryV2() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.FactoryV2(&_Uniswapv3RouterV2.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerSession) FactoryV2() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.FactoryV2(&_Uniswapv3RouterV2.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Caller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3RouterV2.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) PositionManager() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.PositionManager(&_Uniswapv3RouterV2.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2CallerSession) PositionManager() (common.Address, error) {
	return _Uniswapv3RouterV2.Contract.PositionManager(&_Uniswapv3RouterV2.CallOpts)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ApproveMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "approveMax", token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveMax(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveMax(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ApproveMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "approveMaxMinusOne", token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveMaxMinusOne(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveMaxMinusOne(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ApproveZeroThenMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "approveZeroThenMax", token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveZeroThenMax(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveZeroThenMax(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ApproveZeroThenMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "approveZeroThenMaxMinusOne", token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveZeroThenMaxMinusOne(&_Uniswapv3RouterV2.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ApproveZeroThenMaxMinusOne(&_Uniswapv3RouterV2.TransactOpts, token)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) CallPositionManager(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "callPositionManager", data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.CallPositionManager(&_Uniswapv3RouterV2.TransactOpts, data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.CallPositionManager(&_Uniswapv3RouterV2.TransactOpts, data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ExactInput(opts *bind.TransactOpts, params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactInput(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactInput(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ExactInputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactInputSingle(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactInputSingle(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ExactOutput(opts *bind.TransactOpts, params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactOutput(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactOutput(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) ExactOutputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactOutputSingle(&_Uniswapv3RouterV2.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.ExactOutputSingle(&_Uniswapv3RouterV2.TransactOpts, params)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) GetApprovalType(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "getApprovalType", token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.GetApprovalType(&_Uniswapv3RouterV2.TransactOpts, token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.GetApprovalType(&_Uniswapv3RouterV2.TransactOpts, token, amount)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) IncreaseLiquidity(opts *bind.TransactOpts, params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.IncreaseLiquidity(&_Uniswapv3RouterV2.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.IncreaseLiquidity(&_Uniswapv3RouterV2.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) Mint(opts *bind.TransactOpts, params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Mint(&_Uniswapv3RouterV2.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Mint(&_Uniswapv3RouterV2.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) Multicall(opts *bind.TransactOpts, previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "multicall", previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Multicall(&_Uniswapv3RouterV2.TransactOpts, previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Multicall(&_Uniswapv3RouterV2.TransactOpts, previousBlockhash, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) Multicall0(opts *bind.TransactOpts, deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "multicall0", deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Multicall0(&_Uniswapv3RouterV2.TransactOpts, deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Multicall0(&_Uniswapv3RouterV2.TransactOpts, deadline, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) Multicall1(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "multicall1", data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Multicall1(&_Uniswapv3RouterV2.TransactOpts, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Multicall1(&_Uniswapv3RouterV2.TransactOpts, data)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) Pull(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "pull", token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Pull(&_Uniswapv3RouterV2.TransactOpts, token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Pull(&_Uniswapv3RouterV2.TransactOpts, token, value)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) RefundETH() (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.RefundETH(&_Uniswapv3RouterV2.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) RefundETH() (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.RefundETH(&_Uniswapv3RouterV2.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermit(&_Uniswapv3RouterV2.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermit(&_Uniswapv3RouterV2.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermitAllowed(&_Uniswapv3RouterV2.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermitAllowed(&_Uniswapv3RouterV2.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermitAllowedIfNecessary(&_Uniswapv3RouterV2.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermitAllowedIfNecessary(&_Uniswapv3RouterV2.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermitIfNecessary(&_Uniswapv3RouterV2.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SelfPermitIfNecessary(&_Uniswapv3RouterV2.TransactOpts, token, value, deadline, v, r, s)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SwapExactTokensForTokens(&_Uniswapv3RouterV2.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SwapExactTokensForTokens(&_Uniswapv3RouterV2.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SwapTokensForExactTokens(&_Uniswapv3RouterV2.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SwapTokensForExactTokens(&_Uniswapv3RouterV2.TransactOpts, amountOut, amountInMax, path, to)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepToken(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepToken(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SweepToken0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "sweepToken0", token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepToken0(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepToken0(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepTokenWithFee(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepTokenWithFee(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) SweepTokenWithFee0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepTokenWithFee0(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.SweepTokenWithFee0(&_Uniswapv3RouterV2.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UniswapV3SwapCallback(&_Uniswapv3RouterV2.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UniswapV3SwapCallback(&_Uniswapv3RouterV2.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH9(&_Uniswapv3RouterV2.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH9(&_Uniswapv3RouterV2.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) UnwrapWETH90(opts *bind.TransactOpts, amountMinimum *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "unwrapWETH90", amountMinimum)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) UnwrapWETH90(amountMinimum *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH90(&_Uniswapv3RouterV2.TransactOpts, amountMinimum)
}

// UnwrapWETH90 is a paid mutator transaction binding the contract method 0x49616997.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) UnwrapWETH90(amountMinimum *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH90(&_Uniswapv3RouterV2.TransactOpts, amountMinimum)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH9WithFee(&_Uniswapv3RouterV2.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH9WithFee(&_Uniswapv3RouterV2.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) UnwrapWETH9WithFee0(opts *bind.TransactOpts, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH9WithFee0(&_Uniswapv3RouterV2.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.UnwrapWETH9WithFee0(&_Uniswapv3RouterV2.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) WrapETH(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.Transact(opts, "wrapETH", value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.WrapETH(&_Uniswapv3RouterV2.TransactOpts, value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.WrapETH(&_Uniswapv3RouterV2.TransactOpts, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3RouterV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2Session) Receive() (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Receive(&_Uniswapv3RouterV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Uniswapv3RouterV2 *Uniswapv3RouterV2TransactorSession) Receive() (*types.Transaction, error) {
	return _Uniswapv3RouterV2.Contract.Receive(&_Uniswapv3RouterV2.TransactOpts)
}
