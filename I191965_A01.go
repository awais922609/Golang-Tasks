package main

import (
	"crypto/sha256"
	"fmt"
)

var lastHash string

type Block struct {
	transaction  string
	nonce        int
	previousHash string
	blockhash    string
}

func (b *Block) CalculateHash(stringToHash string) string {

	sum := sha256.Sum256([]byte(stringToHash))
	b.blockhash = string(sum[:])

	return b.blockhash

}

type BlockList struct {
	list []*Block
}

func (bc1 *BlockList) NewBlock(transaction string, nonce int, previousHash string) *Block {

	block := new(Block)
	block.transaction = transaction
	block.nonce = nonce
	block.previousHash = previousHash
	fmt.Printf("\nNew Block Hash is == %x", block.CalculateHash(transaction+string(nonce)+previousHash))
	bc1.list = append(bc1.list, block)
	lastHash = block.blockhash

	fmt.Printf("\n Last hash is == %x", lastHash)

	return block

}

func ChangeBlock(b1 *Block) {
	b1.transaction = "Elon Musk to Awais"

	fmt.Printf("\nChange Block Hash is == %x", b1.CalculateHash(b1.transaction+string(b1.nonce)+b1.previousHash))

}

func VerifyChain(ls *BlockList, lastHash string) {
	// var n int
	// n = (len(ls.list))
	for i := 0; i < len(ls.list); i++ {
		if i == 0 {
			ls.list[i].CalculateHash(ls.list[i].transaction + string(ls.list[i].nonce) + "0")
		} else {
			ls.list[i].CalculateHash(ls.list[i].transaction + string(ls.list[i].nonce) + ls.list[i-1].blockhash)

		}
	}

	if ls.list[len(ls.list)-1].blockhash == lastHash {
		fmt.Println("\n\n\t\tChain has been Verified")
	} else {
		fmt.Println("\n\n\t\tChain has been CHanged")
	}

}
func (ls *BlockList) Print() {

	for i := range ls.list {
		//fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("\nTransaction is == %s", ls.list[i].transaction)
		fmt.Printf("\nnonce is == %d", ls.list[i].nonce)
		fmt.Printf("\nPrevious hash is == %x", ls.list[i].previousHash)
		fmt.Printf("\nBlock hash is == %x", ls.list[i].blockhash)

	}
}

func main() {

	block1 := new(BlockList)

	block1.NewBlock("alice to bob", 123, "0")

	block1.NewBlock("Bob to CHarlie", 231, block1.list[len(block1.list)-1].blockhash)

	block1.NewBlock("CHarlie to Alice", 213, block1.list[len(block1.list)-1].blockhash)

	VerifyChain(block1, lastHash)

	ChangeBlock(block1.list[1])

	VerifyChain(block1, lastHash)

}
