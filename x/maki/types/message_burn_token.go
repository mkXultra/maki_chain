package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnToken = "burn_token"

var _ sdk.Msg = &MsgBurnToken{}

func NewMsgBurnToken(creator string, burnType string, amount sdk.Coin) *MsgBurnToken {
	return &MsgBurnToken{
		Creator:  creator,
		BurnType: burnType,
		Amount:   amount,
	}
}

func (msg *MsgBurnToken) Route() string {
	return RouterKey
}

func (msg *MsgBurnToken) Type() string {
	return TypeMsgBurnToken
}

func (msg *MsgBurnToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
