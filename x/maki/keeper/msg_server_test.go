package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mkXultra/maki_chain/testutil/keeper"
	"github.com/mkXultra/maki_chain/x/maki/keeper"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MakiKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
