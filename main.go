package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hacfox/eth-toolkit/contract"
)

func main() {
	c := contract.EthRPC{
		RPCURL: "",
	}

	contractAddr := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	checkSumed := common.HexToAddress(contractAddr).Hex()
	tokenName := c.GetTokenName(checkSumed)
	tokenSymbol := c.GetTokenSymbol(checkSumed)
	tokenDecimals := c.GetTokenDecimals(checkSumed)

	fmt.Println(tokenName)
	fmt.Println(tokenSymbol)
	fmt.Println(tokenDecimals)
}
