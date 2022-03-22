package periphery

import (
	_ "embed"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed contracts/interfaces/ISelfPermit.sol/ISelfPermit.json
var selfpermitABI []byte

var (
	ErrInvalidOptions = errors.New("invalid options")
)

type StandardPermitArguments struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Amount   *big.Int
	Deadline *big.Int
}

type AllowedPermitArguments struct {
	V      uint8
	R      [32]byte
	S      [32]byte
	Nonce  *big.Int
	Expiry *big.Int
}

type PermitOptions struct {
	*StandardPermitArguments
	*AllowedPermitArguments
}

func getSelfPermitABI() abi.ABI {
	var wabi WrappedABI
	err := json.Unmarshal(selfpermitABI, &wabi)
	if err != nil {
		panic(err)
	}
	return wabi.ABI
}

func EncodePermit(token *entities.Token, options *PermitOptions) ([]byte, error) {
	if options == nil {
		return nil, ErrInvalidOptions
	}

	if options.StandardPermitArguments != nil {
		return EncodeStandardPermit(token, options.StandardPermitArguments)
	}

	if options.AllowedPermitArguments != nil {
		return EncodeAllowedPermit(token, options.AllowedPermitArguments)
	}

	return nil, ErrInvalidOptions
}

func EncodeStandardPermit(token *entities.Token, options *StandardPermitArguments) ([]byte, error) {
	abi := getSelfPermitABI()
	return abi.Pack("selfPermit", token.Address, options.Amount, options.Deadline, options.V, options.R, options.S)
}

func EncodeAllowedPermit(token *entities.Token, options *AllowedPermitArguments) ([]byte, error) {
	abi := getSelfPermitABI()
	return abi.Pack("selfPermitAllowed", token.Address, options.Nonce, options.Expiry, options.V, options.R, options.S)
}
