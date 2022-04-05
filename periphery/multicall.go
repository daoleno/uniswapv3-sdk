package periphery

import (
	_ "embed"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed contracts/interfaces/IMulticall.sol/IMulticall.json
var multicallABI []byte

type WrappedABI struct {
	ABI abi.ABI `json:"abi"`
}

func EncodeMulticall(calldatas [][]byte) ([]byte, error) {
	if len(calldatas) == 1 {
		return calldatas[0], nil
	}
	abi := GetABI(multicallABI)
	b, err := abi.Pack("multicall", calldatas)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetABI(abi []byte) abi.ABI {
	var wabi WrappedABI
	err := json.Unmarshal(abi, &wabi)
	if err != nil {
		panic(err)
	}
	return wabi.ABI
}
