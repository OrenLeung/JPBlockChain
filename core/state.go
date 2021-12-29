package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type State struct {
	CurrentBalances map[string]uint `json:"balances"`
	Height          uint            `json:height`
	LastBlockHash   []byte          `json:lastblockhash`
}

func LoadState(pathGenesis string, pathTransactions string) (State, error) {
	var blockCount uint

	currentPath, err := os.Getwd()
	if err != nil {
		return State{}, err
	}

	genesis, err := LoadGenesis(pathGenesis)
	//blockCount increases since we read genesis
	blockCount += 1

	if err != nil {
		return State{}, err
	}

	balances := make(map[string]uint)
	for accountAddress, balance := range genesis.InitalBalances {
		balances[accountAddress] = balance
	}

	genesis.Print()

	parentBlockHash := genesis.Hash

	state := State{balances, 0, parentBlockHash}

	file, err := os.Open(path.Join(currentPath, pathTransactions))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentBlock Block
		err = json.Unmarshal(scanner.Bytes(), &currentBlock)
		if err != nil {
			return State{}, err
		}

		if string(currentBlock.Header.ParentHash) != string(parentBlockHash) {
			return State{}, fmt.Errorf("Block %d's parent hash doesnt match", currentBlock.Header.Height)
		}

		if currentBlock.Header.Height != blockCount {
			return State{}, fmt.Errorf("Block Height Expected %d, but Given %d", blockCount, currentBlock.Header.Height)
		}

		for _, transaction := range currentBlock.Transactions {
			err := state.Apply(transaction)

			if err != nil {
				return State{}, err
			}
		}

		parentBlockHash = currentBlock.Header.Hash
		blockCount++ // increment block height after reading the block
	}

	state.Height = blockCount - 1
	state.LastBlockHash = parentBlockHash

	return state, nil
}

func (state *State) AddBlock(block Block) error {
	state.Height++
	state.LastBlockHash = block.Header.Hash

	for _, transaction := range block.Transactions {
		err := state.Apply(transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (state *State) Apply(transaction Transaction) error {
	if state.CurrentBalances[transaction.SendFrom] >= transaction.Amount {
		state.CurrentBalances[transaction.SendFrom] -= transaction.Amount
		state.CurrentBalances[transaction.SendTo] += transaction.Amount
		return nil
	} else {
		return fmt.Errorf("%v doesn't have enough balance to send amount %d", transaction.SendFrom, transaction.Amount)
	}
}

func (state *State) Validate(transaction Transaction) bool {
	return state.CurrentBalances[transaction.SendFrom] >= transaction.Amount
}

func (state *State) Persist(relativePath string) error {
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return err
	}

	directory, err := os.Getwd()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(directory, relativePath), stateJSON, 0644)
	if err != nil {
		panic(err)
	}

	return nil
}

func (state *State) Print() {
	colorReset := "\033[0m"

	colorRed := "\033[31m"

	fmt.Println(string(colorRed), "------Balances------", string(colorReset))
	for accountAddress, balance := range state.CurrentBalances {
		fmt.Printf("%v: %d\n", accountAddress, balance)
	}
}
