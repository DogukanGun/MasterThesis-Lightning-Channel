package main

import (
	"MasterThesis/cmd"
	"MasterThesis/helper"
	"MasterThesis/logger"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"google.golang.org/grpc"
	"os"
)

func main() {
	helper.SetEnv(".env")
	os.Getenv("ENV")
	grpcConn := helper.GrpcSetup(os.Getenv("POLAR_PORT"), os.Getenv("POLAR_TLS"), os.Getenv("POLAR_MACAROON"))
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logger.LogE(err)
		}
	}(grpcConn)
	lncli := lnrpc.NewLightningClient(grpcConn)

	// Create a new Cobra command
	var rootCmd = &cobra.Command{
		Use:   "lnmsg",
		Short: "Messaging in Lightning Network",
	}

	// Register commands from your `cmd` package
	sendCmd := cmd.SendCmd(&lncli, grpcConn)
	// Add flags to the sendCmd command
	sendCmd.PersistentFlags().StringP("channelID", "c", "", "ID of the channel to send message to ")
	sendCmd.PersistentFlags().StringP("message", "m", "", "Message content to send ")

	rootCmd.AddCommand(cmd.StartCmd(&lncli, grpcConn))
	rootCmd.AddCommand(cmd.StopCmd(&lncli, grpcConn))
	rootCmd.AddCommand(cmd.ChannelCmd(&lncli, grpcConn))
	rootCmd.AddCommand(sendCmd)

	// Execute the root command to start your application
	if err := rootCmd.Execute(); err != nil {
		// Handle errors
		panic(err)
	}
}
