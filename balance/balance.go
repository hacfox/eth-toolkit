package balance

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hacfox/eth-toolkit/rpc"
	"github.com/hacfox/eth-toolkit/utils/log"
)

func GetBalance(address string) *big.Float {
	account := common.HexToAddress(address)
	c := rpc.GetEthClient()
	ba, err := c.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	b := new(big.Float)
	b.SetString(ba.String())
	ethValue := new(big.Float).Quo(b, big.NewFloat(math.Pow10(18)))
	return ethValue
}
