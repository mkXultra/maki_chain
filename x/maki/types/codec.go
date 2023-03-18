package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintToken{}, "maki/MintToken", nil)
	cdc.RegisterConcrete(&MsgBurnToken{}, "maki/BurnToken", nil)
	cdc.RegisterConcrete(&MsgIsBurning{}, "maki/IsBurning", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "maki/Swap", nil)
	cdc.RegisterConcrete(&MsgCreateMaki{}, "maki/CreateMaki", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIsBurning{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwap{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMaki{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
