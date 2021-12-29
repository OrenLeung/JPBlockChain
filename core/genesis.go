package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Balances map[string]uint

type Genesis struct {
	Coinbase       string   `json:"coinbase"`
	InitalBalances Balances `json:"balances"`
	Height         uint64   `json:"height"` // Height of Genesis is always 0
	Hash      []byte   `json:hash`
}

//adapted from: https://github.com/web3coach/the-blockchain-bar/blob/c15_blockchain_forks_what_why_how/database/genesis.go#L42
func LoadGenesis(relativePath string) (Genesis, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return Genesis{}, err
	}
	path := path.Join(currentPath, relativePath)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var loadedGenesis Genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}

func (genesis *Genesis) Print() {
	colorReset := "\033[0m"

	colorRed := "\033[31m"

	fmt.Printf("CoinBase: %v\n", genesis.Coinbase)
	fmt.Printf("Height: %v\n", genesis.Height)
	fmt.Println(string(colorRed), "------Balances------", string(colorReset))
	for accountAddress, balance := range genesis.InitalBalances {
		fmt.Printf("%v: %d\n", accountAddress, balance)
	}
}
