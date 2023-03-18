package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mkXultra/maki_chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMaki_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMaki
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMaki{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMaki{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateMaki_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMaki
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMaki{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMaki{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteMaki_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMaki
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMaki{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMaki{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
