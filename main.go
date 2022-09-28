package main

import (
	"fmt"
	"math/big"

	"github.com/hacfox/eth-toolkit/rpc"
)

func main() {
	amount, err := rpc.Allowance(
		"1", // chain
		"",  // address
		"",  // token address
		"",  // contract address
	)
	if err != nil {
		fmt.Printf("get error. %+v", err)
		return
	}

	if amount.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Approved: %s", amount)
		return
	} else {
		fmt.Println("unapproved")
		return
	}
}
