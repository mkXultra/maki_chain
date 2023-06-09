package maki

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/mkXultra/maki_chain/testutil/sample"
	makisimulation "github.com/mkXultra/maki_chain/x/maki/simulation"
	"github.com/mkXultra/maki_chain/x/maki/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = makisimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgMintToken = "op_weight_msg_mint_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintToken int = 100

	opWeightMsgBurnToken = "op_weight_msg_burn_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnToken int = 100

	opWeightMsgIsBurning = "op_weight_msg_is_burning"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIsBurning int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	opWeightMsgCreateMaki = "op_weight_msg_maki"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMaki int = 100

	opWeightMsgUpdateMaki = "op_weight_msg_maki"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMaki int = 100

	opWeightMsgDeleteMaki = "op_weight_msg_maki"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMaki int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	makiGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MakiList: []types.Maki{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MakiCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&makiGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMintToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintToken, &weightMsgMintToken, nil,
		func(_ *rand.Rand) {
			weightMsgMintToken = defaultWeightMsgMintToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintToken,
		makisimulation.SimulateMsgMintToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurnToken, &weightMsgBurnToken, nil,
		func(_ *rand.Rand) {
			weightMsgBurnToken = defaultWeightMsgBurnToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnToken,
		makisimulation.SimulateMsgBurnToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgIsBurning int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgIsBurning, &weightMsgIsBurning, nil,
		func(_ *rand.Rand) {
			weightMsgIsBurning = defaultWeightMsgIsBurning
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIsBurning,
		makisimulation.SimulateMsgIsBurning(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		makisimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMaki int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMaki, &weightMsgCreateMaki, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMaki = defaultWeightMsgCreateMaki
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMaki,
		makisimulation.SimulateMsgCreateMaki(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMaki int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMaki, &weightMsgUpdateMaki, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMaki = defaultWeightMsgUpdateMaki
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMaki,
		makisimulation.SimulateMsgUpdateMaki(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMaki int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMaki, &weightMsgDeleteMaki, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMaki = defaultWeightMsgDeleteMaki
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMaki,
		makisimulation.SimulateMsgDeleteMaki(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
