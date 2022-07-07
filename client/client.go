package client

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client
var once sync.Once

func initClient() {
	c, err := ethclient.Dial("xxx")
	if err != nil {
		log.Fatal(err)
	}

	client = c
}

func GetEthClient() *ethclient.Client {
	once.Do(initClient)
	return client
}
