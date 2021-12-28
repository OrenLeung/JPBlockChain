package core

import (
	"fmt"
)

type BlockHeader struct {
	Height          uint
	ParentHash      []byte
	Difficulty      uint
	Nonce           uint
	TransactionHash []byte
	Coinbase        []byte // address of the miner
	Reward          uint
}

type Block struct {
	header       BlockHeader
	transactions []Transaction
}

func (b *Block) String() string {
	// Calling Sprintf() function
	s := fmt.Sprintf("this is block %d", b.header.Height)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	return s
}

func (b *Block) addTransaction(tx Transaction) {
	b.transactions = append(b.transactions, tx)
}
