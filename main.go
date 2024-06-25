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
	grpcConn := helper.GrpcSetup(os.Getenv("POLAR_PORT"), os.Getenv("POLAR_TLS"), os.Getenv("POLAR_MACAROON"))
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logger.LogE(err)
		}
	}(grpcConn)
	clientGrpcConn := helper.GrpcSetup(os.Getenv("POLAR_CLIENT_PORT"), os.Getenv("POLAR_CLIENT_TLS"), os.Getenv("POLAR_CLIENT_MACAROON"))
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logger.LogE(err)
		}
	}(clientGrpcConn)
	lncli := lnrpc.NewLightningClient(grpcConn)
	//clientLncli := lnrpc.NewLightningClient(clientGrpcConn)

	// Create a new Cobra command
	var rootCmd = &cobra.Command{
		Use:   "lnmsg",
		Short: "Messaging in Lightning Network",
	}

	// Register commands from your `cmd` package
	sendCmd := cmd.SendCmd(lncli)
	closeCmd := cmd.StopCmd(lncli)

	// Add flags to the sendCmd command
	sendCmd.PersistentFlags().StringP("channelID", "c", "", "ID of the channel to send message to ")
	sendCmd.PersistentFlags().StringP("message", "m", "", "Message content to send ")
	closeCmd.PersistentFlags().StringP("txid", "t", "", "Funding transaction id ")

	rootCmd.AddCommand(cmd.StartCmd())
	rootCmd.AddCommand(closeCmd)
	rootCmd.AddCommand(cmd.ChannelCmd(lncli))
	rootCmd.AddCommand(sendCmd)

	// Execute the root command to start your application
	if err := rootCmd.Execute(); err != nil {
		// Handle errors
		panic(err)
	}
}
