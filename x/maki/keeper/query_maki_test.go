package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/mkXultra/maki_chain/testutil/keeper"
	"github.com/mkXultra/maki_chain/testutil/nullify"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

func TestMakiQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MakiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMaki(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMakiRequest
		response *types.QueryGetMakiResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMakiRequest{Id: msgs[0].Id},
			response: &types.QueryGetMakiResponse{Maki: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetMakiRequest{Id: msgs[1].Id},
			response: &types.QueryGetMakiResponse{Maki: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetMakiRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Maki(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMakiQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MakiKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMaki(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMakiRequest {
		return &types.QueryAllMakiRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MakiAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Maki), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Maki),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MakiAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Maki), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Maki),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MakiAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Maki),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MakiAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
