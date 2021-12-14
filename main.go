package main

import (
	"btc/src"
	"fmt"
	"time"
)

func main(){
	bT := time.Now()            // 开始时间
	bc := src.NewBlockChain()
	for i := 0; i < 3; i++ {
		bc.AddBlock("辛培铖转给刘德华一枚btc"+string(i))
	}
	for i,block := range bc.Blocks {
		fmt.Println("=========block num:",i)
		fmt.Printf("data: %s\n",block.Data)
		fmt.Println("Version:",block.Version)
		fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("TimeStamp: %d\n",block.TimeStamp)
		fmt.Printf("MerKelRoot: %x\n",block.MerKelRoot)
		fmt.Printf("Nonce: %d\n",block.Nonce)
		pow := src.NewProoOfWork(block)
		fmt.Printf("Isvalid: %v\n",pow.IsValid())
	}
	eT := time.Since(bT)      // 从开始到当前所消耗的时间
	fmt.Println("总共运行了: ", eT)
}