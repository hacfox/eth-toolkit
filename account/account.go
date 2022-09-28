package account

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateAddress() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]
	log.Println(privateKeyStr)

	publicKey := privateKey.Public()
	publicECDSA := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicECDSA)
	publicKeyStr := hexutil.Encode(publicKeyBytes)[4:]
	log.Println(publicKeyStr)

	address := crypto.PubkeyToAddress(*publicECDSA).Hex()
	log.Println(address)
}
