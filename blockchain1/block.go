package blockchain1

import (
	"bytes"
	"crypto/sha256"
	"log"

	"github.com/felixgiftinfo/fg-blockchain/common/utils"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int64
}

func (bk *Block) SetHash() {
	nonceByteArray, nonceError := utils.GetByteArray(bk.Nonce)
	if nonceError != nil {
		log.Panic(nonceError)
	}
	join_bytes := bytes.Join([][]byte{bk.Data, bk.PreviousHash, nonceByteArray}, []byte{})
	hashBytes := sha256.Sum256(join_bytes)
	bk.Hash = hashBytes[:]
}

func (chain BlockChain) GetPreviousBlock() *Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

func (chain *BlockChain) AddBlock(nonce int64, data string) {
	prvBk := chain.GetPreviousBlock()
	newBk := CreateBlock(nonce, data, prvBk.Hash)
	chain.Blocks = append(chain.Blocks, newBk)
}

func CreateBlock(nonce int64, data string, previousHash []byte) *Block {
	block := &Block{Data: []byte(data), PreviousHash: previousHash, Nonce: nonce}
	block.SetHash()
	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock(0, "Genesis", []byte{})
}

func InitialBlockchain() *BlockChain {
	return &BlockChain{Blocks: []*Block{CreateGenesisBlock()}}
}
