package makichain_test

import (
	"testing"

	keepertest "github.com/mkXultra/maki_chain/testutil/keeper"
	"github.com/mkXultra/maki_chain/testutil/nullify"
	"github.com/mkXultra/maki_chain/x/makichain"
	"github.com/mkXultra/maki_chain/x/makichain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MakichainKeeper(t)
	makichain.InitGenesis(ctx, *k, genesisState)
	got := makichain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
