package keeper

import (
	"github.com/mkXultra/maki_chain/x/maki/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

// func (k msgServer) MintToken(c context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	err := k.keeper.MintToken(ctx, msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgMintTokenResponse{}, nil
// }

// func (k msgServer) BurnToken(c context.Context, msg *types.MsgBurnToken) (*types.MsgBurnTokenResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	err := k.keeper.BurnToken(ctx, msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgBurnTokenResponse{}, nil
// }

// func (k msgServer) IsBurning(c context.Context, msg *types.MsgIsBurning) (*types.MsgIsBurningResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	err := k.keeper.IsBurning(ctx, msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgIsBurningResponse{}, nil
// }

// func (k msgServer) Swap(c context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	err := k.keeper.Swap(ctx, msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &types.MsgSwapResponse{}, nil
// }
