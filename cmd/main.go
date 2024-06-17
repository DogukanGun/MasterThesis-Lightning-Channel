package cmd

import (
	"MasterThesis/channel"
	"MasterThesis/sender"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func StartCmd(lncli *lnrpc.LightningClient, grpcConn *grpc.ClientConn) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start message receiver",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
func ChannelCmd(lncli *lnrpc.LightningClient, grpcConn *grpc.ClientConn) *cobra.Command {
	return &cobra.Command{
		Use:   "channel",
		Short: "Creates channel with --destination",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			channel.OpenChannel(args[0], grpcConn)
		},
	}
}

func SendCmd(lncli *lnrpc.LightningClient, grpcConn *grpc.ClientConn) *cobra.Command {
	return &cobra.Command{
		Use:   "send",
		Short: "Send [message] to [channelID]",
		Run: func(cmd *cobra.Command, args []string) {
			channelID, err := cmd.Flags().GetString("channelID")
			if err != nil {
				fmt.Printf("Please put channelID flag")
				return
			}
			message, err := cmd.Flags().GetString("message")
			if err != nil {
				fmt.Printf("Please put message flag")
				return
			}
			sender.SendMessage(message, channelID, grpcConn)
		},
	}
}

func FundChannel(lncli *lnrpc.LightningClient, grpcConn *grpc.ClientConn) *cobra.Command {
	return &cobra.Command{
		Use:   "fund",
		Short: "Fund the channel",
		Run: func(cmd *cobra.Command, args []string) {
			//TODO write process here
		},
	}
}

func CloseChannel(lncli *lnrpc.LightningClient, grpcConn *grpc.ClientConn) *cobra.Command {
	return &cobra.Command{
		Use:   "fund",
		Short: "Fund the channel",
		Run: func(cmd *cobra.Command, args []string) {
			//TODO write process here
		},
	}
}

func StopCmd(lncli *lnrpc.LightningClient, grpcConn *grpc.ClientConn) *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stops the message receiver",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
