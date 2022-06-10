package main

import (
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/examples/contract"
	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//mint a new liquidity
func mintOrAdd(client *ethclient.Client, wallet *helper.Wallet, tokenID *big.Int) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	pool, err := helper.ConstructV3Pool(client, helper.WMATIC, helper.AMP, uint64(constants.FeeMedium))
	if err != nil {
		log.Fatal(err)
	}

	//0.1 MATIC
	amount0 := helper.IntWithDecimal(1, 17)
	amount1 := helper.FloatStringToBigInt("5", 18)
	pos, err := entities.FromAmounts(pool, -43260, 29400, amount0, amount1, false)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	onePercent := coreEntities.NewPercent(big.NewInt(1), big.NewInt(100))
	log.Println(pos.MintAmountsWithSlippage(onePercent))

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)

	var opts *periphery.AddLiquidityOptions
	if tokenID == nil {
		//mint a new liquidity position
		opts = &periphery.AddLiquidityOptions{
			CommonAddLiquidityOptions: &periphery.CommonAddLiquidityOptions{
				SlippageTolerance: onePercent,
				Deadline:          deadline,
			},
			MintSpecificOptions: &periphery.MintSpecificOptions{
				Recipient:  wallet.PublicKey,
				CreatePool: true,
			},
		}
	} else {
		//add liquidity to an existing pool
		opts = &periphery.AddLiquidityOptions{
			IncreaseSpecificOptions: &periphery.IncreaseSpecificOptions{
				TokenID: tokenID,
			},
			CommonAddLiquidityOptions: &periphery.CommonAddLiquidityOptions{
				SlippageTolerance: onePercent,
				Deadline:          deadline,
			},
		}

	}
	params, err := periphery.AddCallParameters(pos, opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("0x%x  value=%s\n", params.Calldata, params.Value.String())

	//matic is a native token, so we need to set the actually value to transfer
	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3NFTPositionManager),
		amount0, params.Calldata, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}

func remove(client *ethclient.Client, wallet *helper.Wallet, tokenID *big.Int) {
	//our pool is the fee medium pool
	pool, err := helper.ConstructV3Pool(client, helper.WMATIC, helper.AMP, uint64(constants.FeeMedium))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("liquidity= ", pool.Liquidity)

	posManager, err := contract.NewUniswapv3NFTPositionManager(common.HexToAddress(helper.ContractV3NFTPositionManager), client)
	if err != nil {
		log.Fatal(err)
	}
	contractPos, err := posManager.Positions(nil, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	percent25 := coreEntities.NewPercent(big.NewInt(1), big.NewInt(25))
	fullPercent := coreEntities.NewPercent(contractPos.Liquidity, big.NewInt(1))
	removingLiquidity := fullPercent.Multiply(percent25)

	pos, err := entities.NewPosition(pool, removingLiquidity.Quotient(),
		int(contractPos.TickLower.Int64()),
		int(contractPos.TickUpper.Int64()),
	)
	if err != nil {
		log.Fatal(err)
	}

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)
	opts := &periphery.RemoveLiquidityOptions{
		TokenID:             tokenID,
		LiquidityPercentage: percent25,
		SlippageTolerance:   coreEntities.NewPercent(big.NewInt(1), big.NewInt(100)), //%1  ,
		Deadline:            deadline,
		CollectOptions: &periphery.CollectOptions{
			ExpectedCurrencyOwed0: coreEntities.FromRawAmount(helper.AMP, big.NewInt(0)),
			ExpectedCurrencyOwed1: coreEntities.FromRawAmount(helper.WMATIC, big.NewInt(0)),
			Recipient:             wallet.PublicKey,
		},
	}
	params, err := periphery.RemoveCallParameters(pos, opts)
	if err != nil {
		log.Fatal(err)
	}

	//matic is a native token, so we need to set the actually value to transfer
	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3NFTPositionManager),
		big.NewInt(0), params.Calldata, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}

func burn(client *ethclient.Client, wallet *helper.Wallet, tokenID *big.Int) {
	ABI, _ := abi.JSON(strings.NewReader(contract.Uniswapv3NFTPositionManagerABI))
	out, err := ABI.Pack("burn", tokenID)
	if err != nil {
		log.Fatal(err)
	}

	//matic is a native token, so we need to set the actually value to transfer
	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3NFTPositionManager),
		big.NewInt(0), out, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
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
	_ = client
	_ = wallet
	//mintOrAdd(client, wallet)   //it will create a new NFT ID
	//remove(client, wallet, nftTokenID)
	//burn(client, wallet, nftTokenID) //remove the liquidity
}
