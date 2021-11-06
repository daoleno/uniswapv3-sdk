package utils

import (
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

/**
 * Computes a pool address
 * @param factoryAddress The Uniswap V3 factory address
 * @param tokenA The first token of the pair, irrespective of sort order
 * @param tokenB The second token of the pair, irrespective of sort order
 * @param fee The fee tier of the pool
 * @returns The pool address
 */
func ComputePoolAddress(factoryAddress common.Address, tokenA *entities.Token, tokenB *entities.Token, fee constants.FeeAmount, initCodeHashManualOverride string) (common.Address, error) {
	isSorted, err := tokenA.SortsBefore(tokenB)
	if err != nil {
		return common.Address{}, err
	}
	var (
		token0 *entities.Token
		token1 *entities.Token
	)
	if isSorted {
		token0 = tokenA
		token1 = tokenB
	} else {
		token0 = tokenB
		token1 = tokenA
	}
	return getCreate2Address(factoryAddress, token0.Address, token1.Address, fee, initCodeHashManualOverride), nil
}

func getCreate2Address(factoyAddress, addressA, addressB common.Address, fee constants.FeeAmount, initCodeHashManualOverride string) common.Address {
	var salt [32]byte
	copy(salt[:], crypto.Keccak256(abiEncode(addressA, addressB, fee)))

	if initCodeHashManualOverride != "" {
		crypto.CreateAddress2(factoyAddress, salt, common.FromHex(initCodeHashManualOverride))
	}
	return crypto.CreateAddress2(factoyAddress, salt, common.FromHex(constants.PoolInitCodeHash))
}

func abiEncode(addressA, addressB common.Address, fee constants.FeeAmount) []byte {
	addressTy, _ := abi.NewType("address", "address", nil)
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)

	arguments := abi.Arguments{{Type: addressTy}, {Type: addressTy}, {Type: uint256Ty}}

	bytes, _ := arguments.Pack(
		addressA,
		addressB,
		big.NewInt(int64(fee)),
	)
	return bytes
}
