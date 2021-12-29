package core

import (
	"fmt"
)

type BlockHeader struct {
	Height          uint   `json:"height"`
	Hash            []byte `json:"hash"`
	ParentHash      []byte `json:"parenthash"`
	Difficulty      uint   `json:"difficulty"`
	Nonce           string `json:"nonce"`
	TransactionRoot []byte `json:"transactionroot"`
	Coinbase        string `json:"coinbase"` // address of the miner
	Reward          uint   `json:"reward"`
}

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

func CreateBlock(height uint, parentHash []byte, difficulty uint, coinbase string, reward uint) Block {
	return Block{Header: BlockHeader{Height: height, ParentHash: parentHash, Difficulty: difficulty, Coinbase: coinbase, Reward: reward}}
}

func (b *Block) String() string {
	s := fmt.Sprintf("block height is %d", b.Header.Height)

	return s
}

func (b *Block) AddTransaction(s *State, tx Transaction) error {
	if s.Validate(tx) && len(b.Transactions) < 100 {
		b.Transactions = append(b.Transactions, tx)
		return nil
	} else {
		return fmt.Errorf("%v doesn't have enough balance to send amount %d", tx.SendFrom, tx.Amount)
	}
}
