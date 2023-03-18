package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

func (k msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintTokenResponse{}, nil
}
