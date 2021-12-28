package core

type GenesisAlloc map[string]uint

type Genesis struct {
	Coinbase string
	Alloc    GenesisAlloc
	Number   uint64
}
