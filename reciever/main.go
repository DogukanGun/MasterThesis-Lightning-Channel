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
