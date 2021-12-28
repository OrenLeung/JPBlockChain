package core

type Transaction struct {
	sendFrom string
	sendTo   string
	amount   uint
	gas      uint
	data     []byte
}
