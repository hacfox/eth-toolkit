package hashutil

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/hacfox/eth-toolkit/utils/convert"
	"github.com/inwecrypto/sha3"
	"golang.org/x/crypto/ripemd160"
)

func Sha3(data []byte) []byte {
	hash := sha3.NewKeccak256()
	hash.Write(data)

	buf := hash.Sum(nil)

	return buf
}

// Hash160 returns hash160 of input data bytes.
func Hash160(data []byte) []byte {
	return Ripemd160(Sha256(data))
}

// Hash256 returns hash256 of input data bytes.
func Hash256(data []byte) []byte {
	return Sha256(Sha256(data))
}

// Sha256 returns sha256 of input data bytes.
func Sha256(data []byte) []byte {
	sha256H := sha256.New()
	sha256H.Reset()
	sha256H.Write(data)
	return sha256H.Sum(nil)
}

// Ripemd160 returns RIPEMD-160 hash bytes.
func Ripemd160(data []byte) []byte {
	ripemd160H := ripemd160.New()
	ripemd160H.Reset()
	ripemd160H.Write(data)
	return ripemd160H.Sum(nil)
}

// Checksum returns the checksum for a given piece of data
// using sha256 twice as the hash algorithm.
func Checksum(data []byte) []byte {
	hash := Hash256(data)
	return hash[:4]
}

func Sha3Sig(funcName string) string {
	sigBytes := Sha3([]byte(funcName))
	sig := "0x" + hex.EncodeToString(sigBytes)

	return sig
}

func Sha3Sig4Bytes(funcName string) string {
	sigBytes := Sha3([]byte(funcName))[:4]
	sig := hex.EncodeToString(sigBytes)

	return sig
}

func Sha3Sig4BytesWithPadding(funcName string) string {
	sigBytes := Sha3([]byte(funcName))[:4]
	sig := hex.EncodeToString(sigBytes)
	sig = "0x" + convert.ZeroPeddingRight(sig)

	return sig
}
