package main

import (
	"MasterThesis/helper"
	"MasterThesis/logger"
	"context"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"os"
)

func main() {
	grpcConn := helper.GrpcSetup(os.Getenv("POLAR_CLIENT_PORT"), os.Getenv("POLAR_CLIENT_TLS"), os.Getenv("POLAR_CLIENT_MACAROON"))
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logger.LogE(err)
		}
	}(grpcConn)
	/*
		Comment out here when Polar is used
		conn := lnrpc.NewWalletUnlockerClient(grpcConn)
		unlockReq := lnrpc.UnlockWalletRequest{
			WalletPassword: []byte("dogukan1"),
			RecoveryWindow: 10000,
			ChannelBackups: nil,
		}
		_, err := conn.UnlockWallet(context.Background(), &unlockReq)
		if err != nil {
			if strings.Contains(err.Error(), "unknown service lnrpc.WalletUnlocker") || strings.Contains(err.Error(), "wallet already unlocked") {
				return
			}
			logger.LogE(err.Error())
		}*/
	lnci := lnrpc.NewLightningClient(grpcConn)
	subscribeMessageRequest := lnrpc.SubscribeCustomMessagesRequest{}
	client, err := lnci.SubscribeCustomMessages(context.Background(), &subscribeMessageRequest)
	if err != nil {
		logger.LogE(err.Error())
	}
	for {
		message, err := client.Recv()
		if err != nil {
			logger.LogE(err)
		} else {
			logger.LogI(string(message.Data))
		}

	}
}
