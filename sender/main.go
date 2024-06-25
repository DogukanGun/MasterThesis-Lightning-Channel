package sender

import (
	"MasterThesis/logger"
	"MasterThesis/recorder"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/lightningnetwork/lnd/lnrpc"
)

func SendMessage(message string, channelID string, lncli lnrpc.LightningClient) {
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
	//channelID := "03fea3149e0afff6b948299b247eb3995c54105fa69da6a9d9dd425beb43df3342"
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
		Data: []byte(message),
	}
	resMulti, err := lncli.SendCustomMessage(context.TODO(), &sendCustomMessageRequest)
	if err != nil {
		logger.LogE(err.Error())
	}
	logger.LogS(resMulti.String())
	logger.LogI("Message has been sent")
	db := recorder.Connect()
	recorder.Save(db, "Messages", map[string]interface{}{
		"Message": message,
		"Type":    "SEND",
		"Peer":    peerPublicKeyBytes,
	})
}
