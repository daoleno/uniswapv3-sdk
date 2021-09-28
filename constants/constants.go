package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	PoolInitCodeHash         = "0xe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54"
	PoolInitCodeHashOptimism = "0x0c231002d0970d2126e7e00ce88c3b0e5ec8e48dac71478d56245c34ea2f9447"
)

var (
	FactoryAddress = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressZero    = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// The default factory enabled fee amounts, denominated in hundredths of bips.
type FeeAmount uint

const (
	FeeLow    FeeAmount = 500
	FeeMedium FeeAmount = 3000
	FeeHigh   FeeAmount = 10000
	FeeMax    FeeAmount = 1000000
)

// The default factory tick spacings by fee amount.
var TickSpaces = map[FeeAmount]int64{
	FeeLow:    500,
	FeeMedium: 3000,
	FeeHigh:   10000,
}

var (
	NegativeOne = big.NewInt(-1)
	Zero        = big.NewInt(0)
	One         = big.NewInt(1)

	// used in liquidity amount math
	Q96  = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
	Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)
)
