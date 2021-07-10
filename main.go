package main

import (
	"fmt"

	"github.com/michael_cho77/go-michael-coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second BLock")
	chain.AddBlock("Third dBLock")
	chain.AddBlock("Fourth BLock")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
