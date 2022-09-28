package config

import (
	"strings"
)

type Chain string

const (
	// Tag: EVM compatible chains.
	Ethereum  Chain = "Ethereum"
	BNBChain  Chain = "BNBChain"
	Polygon   Chain = "Polygon"
	Arbitrum  Chain = "Arbitrum"
	Gnosis    Chain = "Gnosis"
	Avalanche Chain = "Avalanche"
	Fantom    Chain = "Fantom"
	Optimism  Chain = "Optimism"
	Cube      Chain = "Cube"
	Metis     Chain = "Metis"
	Celo      Chain = "Celo"
	KCC       Chain = "KCC"
)

func (chain Chain) String() string {
	return string(chain)
}

func GetChain(chainStr string) (Chain, bool) {
	for chain := range supportedChains {
		if strings.EqualFold(chain.String(), chainStr) {
			return chain, true
		}
	}

	return "", false
}

var (
	chainsOrdered   []Chain
	supportedChains map[Chain]bool
	chainIDs        map[uint64]Chain
	chainChainID    map[Chain]uint64
)

func init() {
	supportedChains = make(map[Chain]bool)
	// switch chain
	chainsOrdered = []Chain{
		Ethereum,
		BNBChain,
		Arbitrum,
		Polygon,
		Avalanche,
		Gnosis,
		Optimism,
		Fantom,
		Cube,
		Metis,
		Celo,
		KCC,
	}
	// switch chain
	chainIDs[1] = Ethereum
	chainIDs[10] = Optimism
	chainIDs[56] = BNBChain
	chainIDs[100] = Gnosis
	chainIDs[137] = Polygon
	chainIDs[250] = Fantom
	chainIDs[321] = KCC
	chainIDs[1088] = Metis
	chainIDs[1818] = Cube
	chainIDs[42161] = Arbitrum
	chainIDs[42220] = Celo
	chainIDs[43114] = Avalanche

	chainIDs = map[uint64]Chain{}
	for _, chain := range chainsOrdered {
		supportedChains[chain] = true
	}

	for chainID, chain := range chainIDs {
		chainChainID[chain] = chainID
	}
}
