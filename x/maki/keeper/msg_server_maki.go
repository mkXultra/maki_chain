package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

func (k msgServer) CreateMaki(goCtx context.Context, msg *types.MsgCreateMaki) (*types.MsgCreateMakiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var maki = types.Maki{
		Creator:     msg.Creator,
		Unit:        msg.Unit,
		ExpiredHour: msg.ExpiredHour,
	}

	id := k.keeper.AppendMaki(
		ctx,
		maki,
	)

	return &types.MsgCreateMakiResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMaki(goCtx context.Context, msg *types.MsgUpdateMaki) (*types.MsgUpdateMakiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var maki = types.Maki{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Unit:        msg.Unit,
		ExpiredHour: msg.ExpiredHour,
	}

	// Checks that the element exists
	val, found := k.keeper.GetMaki(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.keeper.SetMaki(ctx, maki)

	return &types.MsgUpdateMakiResponse{}, nil
}

func (k msgServer) DeleteMaki(goCtx context.Context, msg *types.MsgDeleteMaki) (*types.MsgDeleteMakiResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.keeper.GetMaki(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.keeper.RemoveMaki(ctx, msg.Id)

	return &types.MsgDeleteMakiResponse{}, nil
}
