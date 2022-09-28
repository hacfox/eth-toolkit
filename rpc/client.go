package rpc

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hacfox/eth-toolkit/abis"
	"github.com/hacfox/eth-toolkit/utils/log"
)

var client *ethclient.Client
var once sync.Once

// var clientMap = make(map[chain]*ethclient.Client)

func initClient() {
	c, err := ethclient.Dial("https://bsc-dataseed1.ninicoin.io")
	if err != nil {
		log.Fatal(err)
	}

	client = c
}

func GetEthClient() *ethclient.Client {
	once.Do(initClient)
	return client
}

func Allowance(chainID, address, tokenAddress, spender string) (*big.Int, error) {
	token, err := abis.NewErc20Caller(common.HexToAddress(tokenAddress), GetEthClient())
	if err != nil {
		return nil, err
	}

	return token.Allowance(
		&bind.CallOpts{},
		common.HexToAddress(address),
		common.HexToAddress(spender),
	)
}
