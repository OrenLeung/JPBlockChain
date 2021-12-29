package main

import (
	"fmt"
	"os"

	"github.com/OrenLeung/JPBlockChain/core"
)

func main() {

	state, err := core.LoadState("config/genesis.json", "config/blocks.json")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state.Print()

	state.Persist("config/state.json")

}
