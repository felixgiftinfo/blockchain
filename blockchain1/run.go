package blockchain1

import "fmt"

func Run() {
	chain := InitialBlockchain()
	chain.AddBlock(1, "First Block")
	chain.AddBlock(2, "Second Block")
	chain.AddBlock(3, "Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Nonce %v\n", block.Nonce)
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
