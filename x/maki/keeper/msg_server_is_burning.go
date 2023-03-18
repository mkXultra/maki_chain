package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

func (k msgServer) IsBurning(goCtx context.Context, msg *types.MsgIsBurning) (*types.MsgIsBurningResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.keeper.IsBurning(ctx, msg)

	return &types.MsgIsBurningResponse{}, nil
}
