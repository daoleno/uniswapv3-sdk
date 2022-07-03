# Uniswap V3 SDK

[![API Reference](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667)](https://pkg.go.dev/github.com/daoleno/uniswapv3-sdk)
[![Test](https://github.com/daoleno/uniswapv3-sdk/actions/workflows/test.yml/badge.svg)](https://github.com/daoleno/uniswapv3-sdk/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/daoleno/uniswapv3-sdk)](https://goreportcard.com/report/github.com/daoleno/uniswapv3-sdk)

🛠 A Go SDK for building applications on top of Uniswap V3

## Installation

```sh
go get github.com/daoleno/uniswapv3-sdk
```

## Usage

The following example shows how to create a pool, and get the inputAmount

```go
package main

import (
	"fmt"
	"math/big"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	USDC     = core.NewToken(1, common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), 6, "USDC", "USD Coin")
	DAI      = core.NewToken(1, common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), 18, "DAI", "Dai Stablecoin")
	OneEther = big.NewInt(1e18)
)

func main() {
	// create demo ticks
	ticks := []entities.Tick{
		{
			Index:          entities.NearestUsableTick(utils.MinTick, constants.TickSpacings[constants.FeeLow]),
			LiquidityNet:   OneEther,
			LiquidityGross: OneEther,
		},
		{
			Index:          entities.NearestUsableTick(utils.MaxTick, constants.TickSpacings[constants.FeeLow]),
			LiquidityNet:   new(big.Int).Mul(OneEther, constants.NegativeOne),
			LiquidityGross: OneEther,
		},
	}

	// create tick data provider
	p, err := entities.NewTickListDataProvider(ticks, constants.TickSpacings[constants.FeeLow])
	if err != nil {
		panic(err)
	}

	// new pool
	pool, err := entities.NewPool(USDC, DAI, constants.FeeLow, utils.EncodeSqrtRatioX96(constants.One, constants.One), OneEther, 0, p)
	if err != nil {
		panic(err)
	}

	// USDC -> DAI
	outputAmount := core.FromRawAmount(DAI, big.NewInt(98))
	inputAmount, _, err := pool.GetInputAmount(outputAmount, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(inputAmount.ToSignificant(4))
}
```

[More Examples](./examples/README.md)
