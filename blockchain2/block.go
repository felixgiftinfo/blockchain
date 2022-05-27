package blockchain2

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

func (chain BlockChain) GetPreviousBlock() *Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

func (chain *BlockChain) AddBlock(data string) {
	prvBk := chain.GetPreviousBlock()
	newBk := CreateBlock(data, prvBk.Hash)
	chain.Blocks = append(chain.Blocks, newBk)
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{Data: []byte(data), PreviousHash: previousHash, Nonce: 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitialBlockchain() *BlockChain {
	return &BlockChain{Blocks: []*Block{CreateGenesisBlock()}}
}
