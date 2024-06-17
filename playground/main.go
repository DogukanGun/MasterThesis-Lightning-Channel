package main

import (
	"MasterThesis/helper"
	"MasterThesis/logger"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"os"
)

func main() {
	/*lndClient := helper.GetDefaultNode()
	helper.ConnectNode(&lndClient)
	info, err := lndClient.GetInfo()
	if err == nil {
		infoJson, _ := json.Marshal(info)
		logger.LogI(infoJson)
	}
	balance, err := lndClient.GetWalletBalance()
	if err != nil {
		logger.LogE(err)
	}
	logger.LogI(balance)
	channel := helper.CreateChannel(
		&lndClient,
		"03fea3149e0afff6b948299b247eb3995c54105fa69da6a9d9dd425beb43df3342@172.31.0.5:9735",
		250000,
	)
	openStatus, err := channel.Recv()
	logger.LogI(err)
	logger.LogI(openStatus.GetChanOpen())*/
	grpcConn := helper.GrpcSetup(os.Getenv("POLAR_PORT"), os.Getenv("POLAR_TLS"), os.Getenv("POLAR_MACAROON"))
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
	lncli := lnrpc.NewLightningClient(grpcConn)
	getInfoRequest := lnrpc.GetInfoRequest{}
	info, err := lncli.GetInfo(context.TODO(), &getInfoRequest)
	if err != nil {
		fmt.Println(err)
	}
	logger.LogI(info)

	listChannelRequest := lnrpc.ListChannelsRequest{}
	res, err := lncli.ListChannels(context.TODO(), &listChannelRequest)
	if err != nil {
		logger.LogE(err.Error())
	}
	logger.LogS(res.String())
	channelID := "03fea3149e0afff6b948299b247eb3995c54105fa69da6a9d9dd425beb43df3342"
	peerPublicKeyBytes, err := hex.DecodeString(channelID)
	if err != nil {
		logger.LogE("error decoding peer public key: %v", err)
	}
	// Ensure the peerPublicKeyBytes is 33 bytes long
	if len(peerPublicKeyBytes) != 33 {
		logger.LogE("peer public key is not 33 bytes long")
	}
	sendCustomMessageRequest := lnrpc.SendCustomMessageRequest{
		Peer: peerPublicKeyBytes,
		Type: 33768,
		Data: []byte("Hiiii"), //look at the data field, check in low level what are the limits
	}
	//transaction receipts
	//what is published
	//section 1.2 must be section 2. background and add business process management
	//start with how bitcoin works more detail in techinal side
	resMulti, err := lncli.SendCustomMessage(context.TODO(), &sendCustomMessageRequest)
	if err != nil {
		logger.LogE(err.Error())
	}
	logger.LogS(resMulti.String())

	logger.LogI("Message has been sent")
}
