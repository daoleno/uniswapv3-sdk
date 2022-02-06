package periphery

import (
	_ "embed"
	"encoding/json"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed contracts/interfaces/ISelfPermit.sol/ISelfPermit.json
var selfpermitABI []byte

type StandardPermitArguments struct {
	V        uint
	R        string
	S        string
	Amount   *big.Int
	Deadline *big.Int
}

type AllowedPermitArguments struct {
	V      uint
	R      string
	S      string
	Nonce  *big.Int
	Expiry *big.Int
}

func getSelfPermitABI() abi.ABI {
	var wabi WrappedABI
	err := json.Unmarshal(selfpermitABI, &wabi)
	if err != nil {
		panic(err)
	}
	return wabi.ABI
}

func EncodePermit(token *entities.Token, standardPermitOptions *StandardPermitArguments, allowedPermitOptions *AllowedPermitArguments) ([]byte, error) {
	abi := getSelfPermitABI()
	if standardPermitOptions != nil {
		return abi.Pack("selfPermit", token.Address, standardPermitOptions.Amount, standardPermitOptions.Deadline, standardPermitOptions.V, standardPermitOptions.R, standardPermitOptions.S)
	}

	return abi.Pack("selfPermitAllowed", token.Address, allowedPermitOptions.Nonce, allowedPermitOptions.Expiry, allowedPermitOptions.V, allowedPermitOptions.R, allowedPermitOptions.S)
}

func EncodeStandardPermit(token *entities.Token, options *StandardPermitArguments) ([]byte, error) {
	return EncodePermit(token, options, nil)
}

func EncodeAllowedPermit(token *entities.Token, options *AllowedPermitArguments) ([]byte, error) {
	return EncodePermit(token, nil, options)
}
