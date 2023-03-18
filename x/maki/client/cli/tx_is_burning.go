package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mkXultra/maki_chain/x/maki/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdIsBurning() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-burning [burn-type]",
		Short: "Broadcast message IsBurning",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBurnType := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIsBurning(
				clientCtx.GetFromAddress().String(),
				argBurnType,
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
