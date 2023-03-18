package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/mkXultra/maki_chain/x/maki/types"
)

func TestMakiMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateMaki(ctx, &types.MsgCreateMaki{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestMakiMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMaki
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateMaki{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMaki{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMaki{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateMaki(ctx, &types.MsgCreateMaki{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateMaki(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMakiMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMaki
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteMaki{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteMaki{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteMaki{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateMaki(ctx, &types.MsgCreateMaki{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteMaki(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
