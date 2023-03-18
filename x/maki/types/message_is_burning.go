package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIsBurning = "is_burning"

var _ sdk.Msg = &MsgIsBurning{}

func NewMsgIsBurning(creator string, burnType string) *MsgIsBurning {
	return &MsgIsBurning{
		Creator:  creator,
		BurnType: burnType,
	}
}

func (msg *MsgIsBurning) Route() string {
	return RouterKey
}

func (msg *MsgIsBurning) Type() string {
	return TypeMsgIsBurning
}

func (msg *MsgIsBurning) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIsBurning) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIsBurning) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
