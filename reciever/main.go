package reciever

import (
	"MasterThesis/helper"
	"MasterThesis/logger"
	"MasterThesis/recorder"
	"context"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"os"
)

func SubscribeMessages() {
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
			db := recorder.Connect()
			recorder.Save(db, "Messages", map[string]interface{}{
				"Message": message.Data,
				"Type":    "RECEIVE",
				"Peer":    string(message.Peer),
			})
		}

	}
}
