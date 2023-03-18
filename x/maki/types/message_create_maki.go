package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMaki = "create_maki"

var _ sdk.Msg = &MsgCreateMaki{}

func NewMsgCreateMaki(creator string, unit sdk.Coin, expiredHour int32) *MsgCreateMaki {
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
