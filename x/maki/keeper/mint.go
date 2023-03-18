package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mkXultra/maki_chain/x/maki/types"
)

func (k Keeper) MintToken(ctx sdk.Context, msg *types.MsgMintToken) error {
	return nil
}
