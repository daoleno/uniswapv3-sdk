package main

import (
	"math/big"
	"os"
	"time"

	"log"

	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	client, err := ethclient.Dial(helper.PolygonRPC)
	if err != nil {
		log.Fatal(err)
	}
	wallet := helper.InitWallet(os.Getenv("MY_PRIVATE_KEY"))
	if wallet == nil {
		log.Fatal("init wallet failed")
	}

	pool, err := helper.ConstructV3Pool(client, helper.WMATIC, helper.AMP, uint64(constants.FeeMedium))
	if err != nil {
		log.Fatal(err)
	}

	//0.01%
	slippageTolerance := coreEntities.NewPercent(big.NewInt(1), big.NewInt(1000))
	//after 5 minutes
	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)

	// single trade input
	// single-hop exact input
	r, err := entities.NewRoute([]*entities.Pool{pool}, helper.WMATIC, helper.AMP)
	if err != nil {
		log.Fatal(err)
	}

	swapValue := helper.FloatStringToBigInt("0.1", 18)
	trade, err := entities.FromRoute(r, coreEntities.FromRawAmount(helper.WMATIC, swapValue), coreEntities.ExactInput)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v %v\n", trade.Swaps[0].InputAmount.Quotient(), trade.Swaps[0].OutputAmount.Wrapped().Quotient())
	params, err := periphery.SwapCallParameters([]*entities.Trade{trade}, &periphery.SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         wallet.PublicKey,
		Deadline:          deadline,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("calldata = 0x%x\n", params.Value.String())

	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3SwapRouterV1),
		swapValue, params.Calldata, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}
