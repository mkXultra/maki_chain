package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMaki = "create_maki"
	TypeMsgUpdateMaki = "update_maki"
	TypeMsgDeleteMaki = "delete_maki"
)

var _ sdk.Msg = &MsgCreateMaki{}

func NewMsgCreateMaki(creator string, unit sdk.Coin, expiredHour uint64) *MsgCreateMaki {
	return &MsgCreateMaki{
		Creator:     creator,
		Unit:        unit,
		ExpiredHour: expiredHour,
	}
}

func (msg *MsgCreateMaki) Route() string {
	return RouterKey
}

func (msg *MsgCreateMaki) Type() string {
	return TypeMsgCreateMaki
}

func (msg *MsgCreateMaki) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMaki) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMaki) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMaki{}

func NewMsgUpdateMaki(creator string, id uint64, unit sdk.Coin, expiredHour uint64) *MsgUpdateMaki {
	return &MsgUpdateMaki{
		Id:          id,
		Creator:     creator,
		Unit:        unit,
		ExpiredHour: expiredHour,
	}
}

func (msg *MsgUpdateMaki) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMaki) Type() string {
	return TypeMsgUpdateMaki
}

func (msg *MsgUpdateMaki) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMaki) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMaki) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMaki{}

func NewMsgDeleteMaki(creator string, id uint64) *MsgDeleteMaki {
	return &MsgDeleteMaki{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMaki) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMaki) Type() string {
	return TypeMsgDeleteMaki
}

func (msg *MsgDeleteMaki) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMaki) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMaki) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
