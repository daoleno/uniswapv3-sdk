package helper

import (
	"context"
	"log"
	"math/big"
	"testing"

	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ETH_RPC_URL = "https://rpc.ankr.com/eth"

func TestGetPoolAddress(t *testing.T) {

	client, err := ethclient.Dial(ETH_RPC_URL)
	if err != nil {
		t.FailNow()
	}

	poolAddr := common.HexToAddress("0x5777d92f208679db4b9778590fa3cab3ac9e2168")
	dai := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	fee := new(big.Int).SetUint64(100)

	result1, err := GetPoolAddress(client, dai, usdc, fee)
	if err != nil {
		t.FailNow()
	}

	result2, err := GetPoolAddress(client, dai, usdc, fee)
	if err != nil {
		t.FailNow()
	}
	if poolAddr.Hash().Big().Cmp(result1.Hash().Big()) != 0 {
		t.FailNow()
	}

	if poolAddr.Hash().Big().Cmp(result2.Hash().Big()) != 0 {
		t.FailNow()
	}
}

func TestConstructV3Pool(t *testing.T) {

	client, err := ethclient.Dial(ETH_RPC_URL)
	if err != nil {
		t.FailNow()
	}
	cid, _ := client.ChainID(context.Background())

	daiAddr := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	usdcAddr := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	dai := coreEntities.NewToken(uint(cid.Uint64()), daiAddr, 18, "dai", "")
	usdc := coreEntities.NewToken(uint(cid.Uint64()), usdcAddr, 6, "usdc", "")
	fee := uint64(100)

	_, err = ConstructV3Pool(client, dai, usdc, fee)
	if err != nil {
		log.Fatalf("%s", err)
		t.FailNow()
	}
}