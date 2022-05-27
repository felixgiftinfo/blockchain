package blockchain2

import (
	"fmt"
	"strconv"
)

func Run() {
	chain := InitialBlockchain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Nonce %v\n", block.Nonce)
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
