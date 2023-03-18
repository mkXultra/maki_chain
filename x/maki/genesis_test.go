package maki_test

import (
	"testing"

	keepertest "github.com/mkXultra/maki_chain/testutil/keeper"
	"github.com/mkXultra/maki_chain/testutil/nullify"
	"github.com/mkXultra/maki_chain/x/maki"
	"github.com/mkXultra/maki_chain/x/maki/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MakiList: []types.Maki{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MakiCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MakiKeeper(t)
	maki.InitGenesis(ctx, *k, genesisState)
	got := maki.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MakiList, got.MakiList)
	require.Equal(t, genesisState.MakiCount, got.MakiCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
