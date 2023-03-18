package maki

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mkXultra/maki_chain/x/maki/keeper"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the maki
	for _, elem := range genState.MakiList {
		k.SetMaki(ctx, elem)
	}

	// Set maki count
	k.SetMakiCount(ctx, genState.MakiCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MakiList = k.GetAllMaki(ctx)
	genesis.MakiCount = k.GetMakiCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
