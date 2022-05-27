package blockchain2

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/felixgiftinfo/fg-blockchain/common/utils"
)

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(bk *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{bk, target}
	return pow
}

func (pow *ProofOfWork) GetHashData(nonce int) []byte {

	nonceByteArray, err := utils.GetByteArray(int64(nonce))
	if err != nil {
		log.Panic(err)
	}

	difficultyByteArray, err := utils.GetByteArray(int64(Difficulty))
	if err != nil {
		log.Panic(err)
	}

	data := bytes.Join(
		[][]byte{
			pow.Block.PreviousHash,
			pow.Block.Data,
			nonceByteArray,
			difficultyByteArray,
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.GetHashData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.GetHashData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
