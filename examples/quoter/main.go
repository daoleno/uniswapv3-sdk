package main

import (
	"fmt"
	"math/big"

	"github.com/daoleno/uniswapv3-sdk/examples/quoter/uniswapv3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("YOUR-RPC-ENDPOINT")
	if err != nil {
		panic(err)
	}
	quoterAddress := common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")
	quoterContract, err := uniswapv3.NewQuoter(quoterAddress, client)
	if err != nil {
		panic(err)
	}

	token0 := common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6")
	token1 := common.HexToAddress("0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C")
	fee := big.NewInt(3000)
	amountIn := big.NewInt(1000000000000000000)
	sqrtPriceLimitX96 := big.NewInt(0)

	amountOut, err := quoterContract.QuoterCaller.QuoteExactInputSingle(&bind.CallOpts{}, token0, token1, fee, amountIn, sqrtPriceLimitX96)
	if err != nil {
		panic(err)
	}

	fmt.Println("amountOut: ", amountOut)
}
