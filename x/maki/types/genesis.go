package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MakiList: []Maki{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in maki
	makiIdMap := make(map[uint64]bool)
	makiCount := gs.GetMakiCount()
	for _, elem := range gs.MakiList {
		if _, ok := makiIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for maki")
		}
		if elem.Id >= makiCount {
			return fmt.Errorf("maki id should be lower or equal than the last id")
		}
		makiIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
