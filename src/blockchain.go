package src

import (
	"github.com/boltdb/bolt"
)

type BlockChain struct {
	//Blocks []*Block
	db *bolt.DB
	lastHash []byte
}
const dbfile = "blockChainDb.db"
const blockBucket = "block"
const lastHash = "lastHash"
func NewBlockChain() *BlockChain {
	// path,mode,options
	db,err := bolt.Open(dbfile,0600,nil)
	CheckErr("newBlockChain",err)
	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil{
			// 读取
			lastHash = bucket.Get([]byte(lastHash))
		} else {
			// 创建，写数据
			genesis := NewGenesisBlock()
			bucket,err := tx.CreateBucket([]byte(blockBucket))
			CheckErr("newBlockChain",err)
			err = bucket.Put(genesis.Hash,genesis.Serialize())
			CheckErr("newBlockChain",err)
			err = bucket.Put([]byte(lastHash),genesis.Hash)
			CheckErr("newBlockChain",err)
			lastHash = genesis.Hash
		}
		return nil
	})
	//db.View()
	return &BlockChain{db,lastHash}
	//return &BlockChain{
	//	Blocks: []*Block{
	//		NewGenesisBlock(),
	//	},
	//}
}

func (bc *BlockChain) AddBlock(data string) {
	var prevBlockHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		lastHash := bucket.Get([]byte(lastHash))
		prevBlockHash = lastHash
		return nil
	})
	CheckErr("",err)
	block := newBlock(data,prevBlockHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		err := bucket.Put(block.Hash,block.Serialize())
		CheckErr("",err)
		err = bucket.Put([]byte(lastHash),block.Hash)
		CheckErr("",err)
		bc.lastHash = block.Hash
		return nil
	})
	CheckErr("",err)
}
