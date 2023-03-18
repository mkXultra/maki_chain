package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/mkXultra/maki_chain/x/maki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MakiAll(goCtx context.Context, req *types.QueryAllMakiRequest) (*types.QueryAllMakiResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var makis []types.Maki
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	makiStore := prefix.NewStore(store, types.KeyPrefix(types.MakiKey))

	pageRes, err := query.Paginate(makiStore, req.Pagination, func(key []byte, value []byte) error {
		var maki types.Maki
		if err := k.cdc.Unmarshal(value, &maki); err != nil {
			return err
		}

		makis = append(makis, maki)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMakiResponse{Maki: makis, Pagination: pageRes}, nil
}

func (k Keeper) Maki(goCtx context.Context, req *types.QueryGetMakiRequest) (*types.QueryGetMakiResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	maki, found := k.GetMaki(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMakiResponse{Maki: maki}, nil
}
