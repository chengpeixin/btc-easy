package src

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte // 为了方便实现做的简化，实际上比特币区块不包含自身的hash值
	TimeStamp     int64
	TargetBits    int64
	Nonce         int64 // 随机值
	MerKelRoot    []byte
	Data          []byte // 区块体
}

func newBlock(data string, PrevBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: PrevBlockHash,
		// Hash
		TimeStamp:  time.Now().Unix(),
		TargetBits: TargetBits,
		MerKelRoot: []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	pow :=NewProoOfWork(block)
	nonce,hash :=pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func (block *Block)Serialize()[]byte  {
	var buffer  bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr("",err)
	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	var block Block
	err := decoder.Decode(block)
	CheckErr("",err)
	return &block
}

//func (block *Block) SetHash() {
//	tmp := [][]byte{
//		IntToByte(block.Version),
//		block.PrevBlockHash,
//		IntToByte(block.TimeStamp),
//		block.MerKelRoot,
//		IntToByte(block.Nonce),
//		block.Data,
//	}
//	data := bytes.Join(tmp, []byte{})
//	hash := sha256.Sum256(data)
//	block.Hash = hash[:]
//}

func NewGenesisBlock() *Block {
	return newBlock("genesis Block!", []byte{})
}
