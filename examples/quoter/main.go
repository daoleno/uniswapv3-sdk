package main

import (
	"fmt"
	"math/big"

	"log"

	"github.com/daoleno/uniswapv3-sdk/examples/contract"
	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	client, err := ethclient.Dial(helper.PolygonRPC)
	if err != nil {
		panic(err)
	}
	quoterContract, err := contract.NewUniswapv3Quoter(common.HexToAddress(helper.ContractV3Quoter), client)
	if err != nil {
		panic(err)
	}

	token0 := common.HexToAddress(helper.WMaticAddr)
	token1 := common.HexToAddress(helper.AmpAddr)
	fee := big.NewInt(3000)
	amountIn := helper.FloatStringToBigInt("1.00", 18)
	sqrtPriceLimitX96 := big.NewInt(0)

	var out []interface{}
	rawCaller := &contract.Uniswapv3QuoterRaw{Contract: quoterContract}
	err = rawCaller.Call(nil, &out, "quoteExactInputSingle", token0, token1,
		fee, amountIn, sqrtPriceLimitX96)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("amountOut: ", out[0].(*big.Int).String())
}
