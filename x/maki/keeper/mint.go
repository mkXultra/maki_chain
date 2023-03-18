package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mkXultra/maki_chain/x/maki/types"
)

func (k Keeper) MintToken(ctx sdk.Context, msg *types.MsgMintToken) error {
	acc, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	originalDenom := fmt.Sprintf("%s/%s", msg.Creator, msg.Amount.Denom)
	originalCoin := sdk.NewCoin(originalDenom, msg.Amount.Amount)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(originalCoin))
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acc, sdk.NewCoins(originalCoin))
	if err != nil {
		return err
	}

	return nil
}
