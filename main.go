package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcs = []string{
	"https://bsc.rpc.blxrbdn.com",
}

const (
	prvKey = "c0dae90ed265c4507128ed20da3736cadf6b2ebbfd297cf2d0e2238ca2cbb773"
)

func run(rpc string) {
	_, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Printf("Dial: %v, rpc: %v\n", err, rpc)
		return
	}
	privateKey, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		panic(err)
	}
	crypto.PubkeyToAddress(privateKey.PublicKey)
}

func main() {
	run(rpcs[0])
}
