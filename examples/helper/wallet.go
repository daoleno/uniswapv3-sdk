package helper

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  common.Address
}

func (w Wallet) PubkeyStr() string {
	return w.PublicKey.String()
}

func InitWallet(privateHexKeys string) *Wallet {
	if privateHexKeys == "" {
		return nil
	}
	privateKey, err := crypto.HexToECDSA(privateHexKeys)
	if err != nil {
		return nil
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  crypto.PubkeyToAddress(privateKey.PublicKey),
	}
}
