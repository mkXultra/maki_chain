package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mkXultra/maki_chain/testutil/keeper"
	"github.com/mkXultra/maki_chain/testutil/nullify"
	"github.com/mkXultra/maki_chain/x/maki/keeper"
	"github.com/mkXultra/maki_chain/x/maki/types"
	"github.com/stretchr/testify/require"
)

func createNMaki(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Maki {
	items := make([]types.Maki, n)
	for i := range items {
		items[i].Id = keeper.AppendMaki(ctx, items[i])
	}
	return items
}

func TestMakiGet(t *testing.T) {
	keeper, ctx := keepertest.MakiKeeper(t)
	items := createNMaki(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMaki(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMakiRemove(t *testing.T) {
	keeper, ctx := keepertest.MakiKeeper(t)
	items := createNMaki(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMaki(ctx, item.Id)
		_, found := keeper.GetMaki(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMakiGetAll(t *testing.T) {
	keeper, ctx := keepertest.MakiKeeper(t)
	items := createNMaki(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMaki(ctx)),
	)
}

func TestMakiCount(t *testing.T) {
	keeper, ctx := keepertest.MakiKeeper(t)
	items := createNMaki(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMakiCount(ctx))
}
