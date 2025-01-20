package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v8/modules/core/04-channel/client/utils"
	"github.com/igor-sikachyna/leaderboard/x/leaderboard/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// CmdSendIbcTopRank() returns the IbcTopRank send packet command.
// This command does not use AutoCLI because it gives a better UX to do not.
func CmdSendIbcTopRank() *cobra.Command {
	flagPacketTimeoutTimestamp := "packet-timeout-timestamp"

	cmd := &cobra.Command{
		Use:   "send-ibc-top-rank [src-port] [src-channel] [player-id] [rank] [score]",
		Short: "Send a ibcTopRank over IBC",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argPlayerId := args[2]
			argRank, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argScore, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendIbcTopRank(creator, srcPort, srcChannel, timeoutTimestamp, argPlayerId, argRank, argScore)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
