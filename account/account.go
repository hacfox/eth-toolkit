package account

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/hacfox/eth-toolkit/utils/log"
)

func CreateAddress() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]
	log.Info(privateKeyStr)

	publicKey := privateKey.Public()
	publicECDSA := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicECDSA)
	publicKeyStr := hexutil.Encode(publicKeyBytes)[4:]
	log.Info(publicKeyStr)

	address := crypto.PubkeyToAddress(*publicECDSA).Hex()
	log.Info(address)
}
