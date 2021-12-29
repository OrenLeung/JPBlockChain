package core

import "fmt"

type Transaction struct {
	SendFrom string `json:"sendFrom"`
	SendTo   string `json:"sendTo"`
	Amount   uint   `json:"amount"`
}

func (t *Transaction) Print() {
	fmt.Printf("%v -> %v: %d\n", t.SendFrom, t.SendTo, t.Amount)
}
