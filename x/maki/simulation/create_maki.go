package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/mkXultra/maki_chain/x/maki/keeper"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

func SimulateMsgCreateMaki(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateMaki{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateMaki simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateMaki simulation not implemented"), nil, nil
	}
}
