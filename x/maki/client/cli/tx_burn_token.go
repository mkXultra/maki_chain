package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mkXultra/maki_chain/x/maki/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBurnToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-token [burn-type] [amount]",
		Short: "Broadcast message BurnToken",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBurnType := args[0]
			argAmount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnToken(
				clientCtx.GetFromAddress().String(),
				argBurnType,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
