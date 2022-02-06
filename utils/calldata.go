package utils

import "math/big"

type MethodParameters struct {
	Calldata []byte   // The hex encoded calldata to perform the given operation
	Value    *big.Int // The amount of ether (wei) to send in hex
}
