package keeper_test

import (
	"testing"

	testkeeper "github.com/mkXultra/maki_chain/testutil/keeper"
	"github.com/mkXultra/maki_chain/x/makichain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MakichainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
