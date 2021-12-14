package src

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const TargetBits = 24

type ProofOWork struct {
	Block     *Block
	TargetBit *big.Int
}

func NewProoOfWork(block *Block) *ProofOWork {
	var IntTarget = big.NewInt(1)
	IntTarget.Lsh(IntTarget, uint(256-TargetBits))
	return &ProofOWork{
		Block:     block,
		TargetBit: IntTarget,
	}
}

func (pow *ProofOWork) PrepareRawData(nonce int64) []byte {
	block := pow.Block
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.MerKelRoot,
		IntToByte(nonce),
		IntToByte(TargetBits),
		block.Data,
	}
	data := bytes.Join(tmp, []byte{})
	return data
}

func (pow *ProofOWork) Run() (int64, []byte) {
	var nonce int64
	var hash [32]byte
	var HashInt big.Int
	fmt.Println("开始挖矿...")
	fmt.Printf("target hash: %x\n",pow.TargetBit.Bytes())
	for nonce < math.MaxInt64 {
		data := pow.PrepareRawData(nonce)
		hash = sha256.Sum256(data)
		// hash字符串转换为int
		HashInt.SetBytes(hash[:])
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if HashInt.Cmp(pow.TargetBit) == -1 {
			fmt.Printf("Found Hash : %x\n", hash)
			break
		} else {
			//fmt.Printf("current Hash: %x\n",hash)
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOWork) IsValid() bool {
	data :=pow.PrepareRawData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	var IntHash big.Int
	IntHash.SetBytes(hash[:])
	return IntHash.Cmp(pow.TargetBit) == -1
}
