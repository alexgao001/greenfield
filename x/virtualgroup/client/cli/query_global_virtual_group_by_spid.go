package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/bnb-chain/greenfield/x/virtualgroup/types"
)

var _ = strconv.Itoa(0)

func CmdGlobalVirtualGroupBySPID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "global-virtual-group-by-spid",
		Short: "query all global virtual groups of the storage provider.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGlobalVirtualGroupBySPIDRequest{}

			res, err := queryClient.GlobalVirtualGroupBySPID(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}