package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mkXultra/maki_chain/x/maki/types"
)

func (k Keeper) BurnToken(ctx sdk.Context, msg *types.MsgBurnToken) error {
	return nil
}

func (k Keeper) IsBurning(ctx sdk.Context, msg *types.MsgIsBurning) error {
	return nil
}
